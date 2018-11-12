package main

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/TheThingsIndustries/protoc-gen-fieldmask/internal/extensions"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
)

const FileHeader = `// Code generated by protoc-gen-fieldmask. DO NOT EDIT.`

const copyUtil = `func deepCopy(dst, src interface{}) {
	copyRecursive(reflect.ValueOf(dst), reflect.ValueOf(src))
}

// NOTE: The following block is sligthly modified https://github.com/mohae/deepcopy/tree/c48cc78d482608239f6c4c92a4abd87eb8761c90

// The MIT License (MIT)
//
// Copyright (c) 2014 Joel
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// deepCopier for delegating copy process to type
type deepCopier interface {
	DeepCopy() interface{}
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(cpy, original reflect.Value) {
	// check for implement deepcopy.deepCopier
	if original.CanInterface() {
		if copier, ok := original.Interface().(deepCopier); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// handle according to original's Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to.
		originalValue := original.Elem()

		// if  it isn't valid, return.
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(cpy.Elem(), originalValue)

	case reflect.Interface:
		// If this is a nil, don't do anything
		if original.IsNil() {
			return
		}
		// Get the value for the interface, not the pointer.
		originalValue := original.Elem()

		// Get the value by calling Elem().
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(copyValue, originalValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		t, ok := original.Interface().(time.Time)
		if ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		// Go through each field of the struct and copy it.
		for i := 0; i < original.NumField(); i++ {
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(cpy.Field(i), original.Field(i))
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}
		// Make a new slice and copy each element.
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(cpy.Index(i), original.Index(i))
		}

	case reflect.Map:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, originalKey := range original.MapKeys() {
			originalValue := original.MapIndex(originalKey)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(copyValue, originalValue)
			copyKey := reflect.New(originalKey.Type()).Elem()
			copyRecursive(copyKey, originalKey)
			cpy.SetMapIndex(copyKey, copyValue)
		}

	default:
		cpy.Set(original)
	}
}`

const (
	protoAnyType       = ".google.protobuf.Any"
	protoDurationType  = ".google.protobuf.Duration"
	protoFieldMaskType = ".google.protobuf.FieldMask"
	protoStructType    = ".google.protobuf.Struct"
	protoTimestampType = ".google.protobuf.Timestamp"
)

type unknownTypeError string

func (e unknownTypeError) Error() string {
	return fmt.Sprintf("message of proto type '%s' is unknown", string(e))
}

type unsupportedTypeError string

func (e unsupportedTypeError) Error() string {
	return fmt.Sprintf("fields of proto type '%s' are unsupported", string(e))
}

type recursionError struct {
	file  string
	field string
}

func (e recursionError) Error() string {
	return fmt.Sprintf("field '%s' defined at %s is recursive", e.field, e.file)
}

func eqCopyOp(dst, src string) string {
	return fmt.Sprintf("%s = %s", dst, src)
}

func deepCopyOp(dst, src string) string {
	return fmt.Sprintf("deepCopy(&%s, &%s)", dst, src)
}

func init() {
	log.SetFlags(0)
}

func appendPaths(paths []string, prefix string, md *protokit.Descriptor, mds map[string]*protokit.Descriptor, seen map[string]struct{}) ([]string, error) {
	if seen == nil {
		seen = map[string]struct{}{}
	}

	for _, fd := range md.GetMessageFields() {
		if _, ok := seen[fd.GetFullName()]; ok {
			return nil, recursionError{
				file:  fd.GetFile().GetName(),
				field: fd.GetFullName(),
			}
		}
		seen[fd.GetFullName()] = struct{}{}

		fp := fd.GetName()
		if prefix != "" {
			fp = fmt.Sprintf("%s.%s", prefix, fp)
		}

		if fd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			paths = append(paths, fp)
			delete(seen, fd.GetFullName())
			continue
		}
		if fd.GetType() != descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			paths = append(paths, fp)
			delete(seen, fd.GetFullName())
			continue
		}

		fmd, ok := mds[fd.GetTypeName()]
		if !ok {
			switch fd.GetTypeName() {
			case protoTimestampType, protoFieldMaskType, protoDurationType, protoStructType, protoAnyType:
				paths = append(paths, fp)
				delete(seen, fd.GetFullName())
				continue
			}
			return nil, unknownTypeError(fd.GetTypeName())
		}

		if len(fmd.GetMessageFields()) == 0 {
			paths = append(paths, fp)
			delete(seen, fd.GetFullName())
			continue
		}

		var err error
		paths, err = appendPaths(paths, fp, fmd, mds, seen)
		if err != nil {
			return nil, err
		}
		delete(seen, fd.GetFullName())
	}
	return paths, nil
}

func fieldTypeName(fd *protokit.FieldDescriptor) (goType string) {
	goType = fd.GetTypeName()
	if i := strings.LastIndex(goType, "."); i > 0 {
		goType = goType[i+1:]
	}

	protoType := fd.GetTypeName()[1:]

	for parent := fd.GetMessage(); parent != nil; parent = parent.GetParent() {
		for _, smd := range parent.GetMessages() {
			if protoType == smd.GetFullName() {
				goType = fmt.Sprintf("%s_%s", parent.GetName(), goType)
				if i := strings.LastIndex(goType, "."); i > 0 {
					goType = goType[i+1:]
				}
				protoType = parent.GetFullName()
			}
		}
	}
	return goType
}

func enumTypeName(fd *protokit.FieldDescriptor) (goType string) {
	goType = fd.GetTypeName()
	if i := strings.LastIndex(goType, "."); i > 0 {
		goType = goType[i+1:]
	}

	protoType := fd.GetTypeName()[1:]

	for parent := fd.GetMessage(); parent != nil; parent = parent.GetParent() {
		for _, sed := range parent.GetEnums() {
			if protoType == sed.GetFullName() {
				goType = fmt.Sprintf("%s_%s", parent.GetName(), goType)
				if i := strings.LastIndex(goType, "."); i > 0 {
					goType = goType[i+1:]
				}
				protoType = parent.GetFullName()
			}
		}
	}
	return goType
}

var importPathReplacer = strings.NewReplacer(
	".", "_",
	"/", "_",
	"-", "_",
)

func buildMethods(buf *strings.Builder, md *protokit.Descriptor, mds map[string]*protokit.Descriptor) (map[string]string, error) {
	imports := map[string]string{
		"types": "github.com/gogo/protobuf/types",
	}

	paths, err := appendPaths(make([]string, 0, len(md.GetMessageFields())), "", md, mds, nil)
	if err != nil {
		// TODO: Return error here
		log.Printf("Failed to traverse `%s`: %s, skipping...", md.GetFullName(), err)
		return nil, nil
	}

	fmt.Fprintf(buf, `
var _%sFieldPaths = [...]string{%s}

func (*%s) FieldMaskPaths() []string {`,
		md.GetName(), `"`+strings.Join(paths, `", "`)+`"`,
		md.GetName(),
	)

	sort.Strings(paths)
	if len(paths) == 0 {
		fmt.Fprintf(buf, `
	return nil`,
		)
	} else {
		fmt.Fprintf(buf, `
	ret := make([]string, len(_%sFieldPaths))
	copy(ret, _%sFieldPaths[:])
	return ret`,
			md.GetName(),
			md.GetName(),
		)
	}

	fmt.Fprintf(buf, `
}

func (dst *%s) SetFields(src *%s, mask *types.FieldMask) {
	for _, path := range mask.GetPaths() {
		switch path {`,
		md.GetName(), md.GetName(),
	)

	for _, p := range paths {
		fmt.Fprintf(buf, `
		case "%s":
			var nilPath bool`,
			p)

		sp := strings.Split(p, ".")

		srcPath := "src"
		fm := md
		for i := 0; i < len(sp)-1; i++ {
			fd := fm.GetMessageField(sp[i])
			if fd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
				panic(errors.New("Fieldmask for repeated field generated"))
			}

			goType := fieldTypeName(fd)

			goName := generator.CamelCase(fd.GetName())
			if v, ok := fd.OptionExtensions["gogoproto.customname"].(*string); ok {
				goName = *v
			}
			if v, ok := fd.OptionExtensions["gogoproto.embed"].(*bool); ok && *v {
				goName = goType
			}

			if fd.OneofIndex != nil {
				oneOfType := fmt.Sprintf("%s_%s", fm.GetName(), goName)
				if fm.GetMessage(goName) != nil {
					oneOfType = fmt.Sprintf("%s_", oneOfType)
				}

				srcPath = fmt.Sprintf("%s.Get%s()", srcPath, goName)

				fmt.Fprintf(buf, `
			nilPath = nilPath || %s == nil`,
					srcPath,
				)

			} else {
				srcPath = fmt.Sprintf("%s.%s", srcPath, goName)
			}

			var ok bool
			fm, ok = mds[fd.GetTypeName()]
			if !ok {
				return nil, unknownTypeError(fd.GetTypeName())
			}

			if v, ok := fd.OptionExtensions["gogoproto.nullable"].(*bool); ok && !*v {
				continue
			}

			fmt.Fprintf(buf, `
			nilPath = nilPath || %s == nil`,
				srcPath,
			)
		}

		fmt.Fprintln(buf)

		dstPath := "dst"
		fm = md
		for i := 0; i < len(sp)-1; i++ {
			fd := fm.GetMessageField(sp[i])
			if fd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
				panic(errors.New("Fieldmask for repeated field generated"))
			}

			goType := fieldTypeName(fd)

			goName := generator.CamelCase(fd.GetName())
			if v, ok := fd.OptionExtensions["gogoproto.customname"].(*string); ok {
				goName = *v
			}
			if v, ok := fd.OptionExtensions["gogoproto.embed"].(*bool); ok && *v {
				goName = goType
			}

			if fd.OneofIndex != nil {
				oneOfType := fmt.Sprintf("%s_%s", fm.GetName(), goName)
				if fm.GetMessage(goName) != nil {
					oneOfType = fmt.Sprintf("%s_", oneOfType)
				}
				oneOfName := generator.CamelCase(fm.GetOneofDecl()[fd.GetOneofIndex()].GetName())

				dstPath = fmt.Sprintf("%s.%s", dstPath, oneOfName)

				fmt.Fprintf(buf, `
			switch {
			case %s != nil && nilPath:
			case %s == nil && nilPath:
				continue
			case %s == nil:
				%s = &%s{}
			}
`,
					dstPath,
					dstPath,
					dstPath,
					dstPath, oneOfType,
				)

				dstPath = fmt.Sprintf("%s.(*%s).%s", dstPath, oneOfType, goName)

			} else {
				dstPath = fmt.Sprintf("%s.%s", dstPath, goName)
			}

			var ok bool
			fm, ok = mds[fd.GetTypeName()]
			if !ok {
				return nil, unknownTypeError(fd.GetTypeName())
			}

			if v, ok := fd.OptionExtensions["gogoproto.nullable"].(*bool); ok && !*v {
				continue
			}

			fmt.Fprintf(buf, `
			switch {
			case %s != nil && nilPath:
			case %s == nil && nilPath:
				continue
			case %s == nil:
				%s = &%s{}
			}
`,
				dstPath,
				dstPath,
				dstPath,
				dstPath, goType,
			)
		}

		fd := fm.GetMessageField(sp[len(sp)-1])

		copyOp := eqCopyOp
		var goType string
		switch fd.GetType() {
		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			goType = "bool"

		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			goType = "float64"

		case descriptor.FieldDescriptorProto_TYPE_ENUM:
			goType = enumTypeName(fd)

		case descriptor.FieldDescriptorProto_TYPE_FIXED32:
			return nil, unsupportedTypeError(fd.GetType().String())

		case descriptor.FieldDescriptorProto_TYPE_FIXED64:
			return nil, unsupportedTypeError(fd.GetType().String())

		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			goType = "float32"

		case descriptor.FieldDescriptorProto_TYPE_INT32:
			goType = "int32"

		case descriptor.FieldDescriptorProto_TYPE_INT64:
			goType = "int64"

		case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
			return nil, unsupportedTypeError(fd.GetType().String())

		case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
			return nil, unsupportedTypeError(fd.GetType().String())

		case descriptor.FieldDescriptorProto_TYPE_SINT32:
			return nil, unsupportedTypeError(fd.GetType().String())

		case descriptor.FieldDescriptorProto_TYPE_SINT64:
			return nil, unsupportedTypeError(fd.GetType().String())

		case descriptor.FieldDescriptorProto_TYPE_STRING:
			goType = "string"

		case descriptor.FieldDescriptorProto_TYPE_UINT32:
			goType = "uint32"

		case descriptor.FieldDescriptorProto_TYPE_UINT64:
			goType = "uint64"

		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			goType = "[]byte"
			copyOp = func(dst, src string) string {
				return fmt.Sprintf(`%s = make([]byte, len(%s))
copy(%s, %s)`,
					dst, src,
					dst, src,
				)
			}

		case descriptor.FieldDescriptorProto_TYPE_GROUP:
			return nil, unsupportedTypeError(fd.GetType().String())

		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			goType = fd.GetTypeName()
			switch goType {
			case protoTimestampType:
				if v, ok := fd.OptionExtensions["gogoproto.stdtime"].(*bool); ok && *v {
					imports["time"] = "time"
					goType = "time.Time"
					copyOp = func(dst, src string) string {
						return fmt.Sprintf("%s = time.Unix(0, %s.UnixNano()).UTC()", dst, src)
					}
					break
				}
				return nil, unsupportedTypeError(fd.GetType().String())

			case protoDurationType:
				if v, ok := fd.OptionExtensions["gogoproto.stdduration"].(*bool); ok && *v {
					imports["time"] = "time"
					goType = "time.Duration"
					break
				}
				return nil, unsupportedTypeError(fd.GetType().String())

			case protoAnyType:
				goType = "types.Any"
				// TODO: Implement non-reflective copying
				copyOp = deepCopyOp

			case protoStructType:
				goType = "types.Struct"
				// TODO: Implement non-reflective copying
				copyOp = deepCopyOp

			case protoFieldMaskType:
				goType = "types.FieldMask"
				copyOp = func(dst, src string) string {
					return fmt.Sprintf(`%s.Paths = make([]string, len(%s.Paths))
copy(%s.Paths, %s.Paths)`,
						dst, src,
						dst, src,
					)
				}

			default:
				// NOTE: Message has no fields, so we can copy with `=`
				goType = fieldTypeName(fd)
			}

		default:
			return nil, unsupportedTypeError(fd.GetType().String())
		}
		var isNullable bool
		if v, ok := fd.OptionExtensions["gogoproto.customtype"].(*string); ok {
			isNullable = true
			copyOp = deepCopyOp
			if i := strings.LastIndex(*v, "."); i >= 0 {
				pkgPath := (*v)[:i]
				typeName := (*v)[i+1:]
				pkgAlias := importPathReplacer.Replace((*v)[:i])

				goType = fmt.Sprintf("%s.%s", pkgAlias, typeName)
				if imported, ok := imports[pkgAlias]; ok && imported != pkgPath {
					panic(fmt.Errorf("Import name clash at `%s`. Imported `%s` and `%s`", pkgAlias, pkgPath, imported))
				}
				imports[pkgAlias] = pkgPath
			} else {
				goType = *v
			}
		}

		isNullable = isNullable || fd.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE || fd.GetType() == descriptor.FieldDescriptorProto_TYPE_GROUP
		if v, ok := fd.OptionExtensions["gogoproto.nullable"].(*bool); ok {
			isNullable = *v
		}

		goName := generator.CamelCase(fd.GetName())
		if v, ok := fd.OptionExtensions["gogoproto.customname"].(*string); ok {
			goName = *v
		}
		if v, ok := fd.OptionExtensions["gogoproto.embed"].(*bool); ok && *v {
			goName = goType
		}

		if fd.OneofIndex != nil {
			oneOfType := fmt.Sprintf("%s_%s", fm.GetName(), goName)
			if fm.GetMessage(goName) != nil {
				oneOfType = fmt.Sprintf("%s_", oneOfType)
			}
			oneOfName := generator.CamelCase(fm.GetOneofDecl()[fd.GetOneofIndex()].GetName())

			srcPath = fmt.Sprintf("%s.Get%s()", srcPath, goName)
			dstPath = fmt.Sprintf("%s.%s", dstPath, oneOfName)

			fmt.Fprintf(buf, `
			switch {
			case %s != nil && nilPath:
			case %s == nil && nilPath:
				continue
			case %s == nil:
				%s = &%s{}
			}
`,
				dstPath,
				dstPath,
				dstPath,
				dstPath, oneOfType,
			)

			dstPath = fmt.Sprintf("%s.(*%s).%s", dstPath, oneOfType, goName)

		} else {
			srcPath = fmt.Sprintf("%s.%s", srcPath, goName)
			dstPath = fmt.Sprintf("%s.%s", dstPath, goName)
		}

		isRepeated := fd.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED
		if isNullable || isRepeated {
			fmt.Fprintf(buf, `
			if nilPath || %s == nil {
				%s = nil
				continue
			}`,
				srcPath,
				dstPath,
			)
		} else {
			fmt.Fprintf(buf, `
			if nilPath {
				var v %s
				%s = v
				continue
			}`,
				goType,
				dstPath,
			)
		}

		if isRepeated {
			if strings.HasSuffix(fd.GetTypeName(), "Entry") && mds[fd.GetFullName()] == nil {
				// fd is a map field, however obtaining its type is non-trivial, hence we resort to reflection
				fmt.Fprintf(buf, `
			%s`, deepCopyOp(dstPath, srcPath))
				continue
			}

			if isNullable {
				goType = fmt.Sprintf("[]*%s", goType)
			} else {
				goType = fmt.Sprintf("[]%s", goType)
			}

			var copyStr string
			for _, l := range strings.Split(copyOp(fmt.Sprintf("%s[%s]", dstPath, "i"), "v"), "\n") {
				copyStr += fmt.Sprintf(`
				%s`,
					l,
				)
			}

			fmt.Fprintf(buf, `
			%s = make(%s, len(%s))
			for i, v := range %s {%s
			}`,
				dstPath, goType, srcPath,
				srcPath, copyStr,
			)

			continue
		}

		if isNullable {
			fmt.Fprintf(buf, `
			var v %s
			%s = &v`,
				goType,
				dstPath,
			)

			srcPath = fmt.Sprintf("(*%s)", srcPath)
			dstPath = fmt.Sprintf("(*%s)", dstPath)
		}

		var copyStr string
		for _, l := range strings.Split(copyOp(dstPath, srcPath), "\n") {
			copyStr += fmt.Sprintf(`
			%s`,
				l,
			)
		}
		fmt.Fprint(buf, copyStr)
	}
	fmt.Fprint(buf, `
		}
	}
}`,
	)
	return imports, nil
}

type plugin struct{}

func (p plugin) Generate(in *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	resp := &plugin_go.CodeGeneratorResponse{}

	fds := protokit.ParseCodeGenRequest(in)

	mds := map[string]*protokit.Descriptor{}
	for _, fd := range fds {
		for _, md := range fd.GetMessages() {
			for _, smd := range append(md.GetMessages(), md) {
				k := fmt.Sprintf(".%s", smd.GetFullName())
				if _, ok := mds[k]; ok {
					return nil, fmt.Errorf("Message name clash at `%s`", k)
				}
				mds[k] = smd
			}
		}
	}

	dirs := map[string]struct{}{}
	for _, fd := range fds {
		if len(fd.GetMessages()) == 0 {
			continue
		}

		dirName := fd.Options.GetGoPackage()
		if dirName == "" {
			dirName = filepath.Dir(fd.GetName())
		}
		dirs[dirName] = struct{}{}

		fileName := filepath.Join(dirName, fmt.Sprintf("%s.pb.fm.go", strings.TrimSuffix(filepath.Base(fd.GetName()), filepath.Ext(fd.GetName()))))

		imports := map[string]string{}
		buf := &strings.Builder{}
		for _, md := range fd.GetMessages() {
			if v, ok := md.OptionExtensions["fieldmask.enable"].(*bool); ok && !*v {
				continue
			}

			mBuf := &strings.Builder{}
			mImports, err := buildMethods(mBuf, md, mds)
			if err != nil {
				return nil, err
			}

			for name, pkg := range mImports {
				if v, ok := imports[name]; ok && v != pkg {
					return nil, fmt.Errorf("Import name clash at `%s`. Imported `%s` and `%s`", name, pkg, v)
				}
				imports[name] = pkg
			}

			if mBuf.Len() == 0 {
				continue
			}
			fmt.Fprintf(buf, `
%s`,
				mBuf.String())
		}

		if buf.Len() == 0 {
			continue
		}

		var importString string
		switch len(imports) {
		case 0:
		case 1:
			for name, pkg := range imports {
				importString = fmt.Sprintf(`
import %s "%s"`, name, pkg)
			}
		default:
			importLines := make([]string, 0, len(imports))
			for name, pkg := range imports {
				importLines = append(importLines, fmt.Sprintf(`	%s "%s"`, name, pkg))
			}
			sort.Slice(importLines, func(i, j int) bool {
				return strings.Fields(importLines[i])[1] < strings.Fields(importLines[j])[1]
			})
			importString = fmt.Sprintf(`
import (
%s
)`,
				strings.Join(importLines, "\n"))
		}

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name: proto.String(fileName),
			Content: proto.String(fmt.Sprintf(`%s

package %s
%s%s
`,
				FileHeader,
				filepath.Base(dirName),
				importString,
				buf.String(),
			)),
		})
	}

	for dirName := range dirs {
		pkgName := filepath.Base(dirName)
		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name: proto.String(filepath.Join(dirName, fmt.Sprintf("%s.pb.util.fm.go", pkgName))),
			Content: proto.String(fmt.Sprintf(`%s

package %s

import (
	"reflect"
	"time"
)

%s
`,
				FileHeader,
				pkgName,
				copyUtil,
			)),
		})
	}

	return resp, nil
}

func main() {
	if err := protokit.RunPlugin(plugin{}); err != nil {
		log.Fatalf("Failed to run plugin: %s", err)
	}
}
