// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata"
	"github.com/kr/pretty"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

var regenerate = flag.Bool("regenerate", false, "regenerate golden files")

// When the environment variable RUN_AS_PROTOC_GEN_GO is set, we skip running
// tests and instead act as protoc-gen-fieldmask. This allows the test binary to
// pass itself to protoc.
func init() {
	if os.Getenv("IN_TEST") != "" {
		main()
		os.Exit(0)
	}
}

func runProtoc(t *testing.T, args ...string) {
	protocBin := "protoc"
	if v := os.Getenv("PROTOC"); v != "" {
		protocBin = v
	}

	cmd := exec.Command(
		strings.Fields(protocBin)[0],
		append(append(strings.Fields(protocBin)[1:], fmt.Sprintf("--plugin=protoc-gen-fieldmask=%s", os.Args[0])), args...)...,
	)
	cmd.Env = append(os.Environ(), "IN_TEST=1")
	t.Logf(`Running '%s'...`, strings.Join(cmd.Args, " "))

	out, err := cmd.CombinedOutput()
	t.Logf(`Output:
%s`,
		string(out),
	)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}
}

func TestGolden(t *testing.T) {
	workDir, err := ioutil.TempDir(os.Getenv("WORKDIR"), "fieldmask-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(workDir)

	var paths []string
	if err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".proto") {
			return nil
		}
		paths = append(paths, path)
		return nil
	}); err != nil {
		t.Errorf("Failed to walk `testdata`: %s", err)
		t.FailNow()
	}

	runProtoc(t, append([]string{
		"-Ivendor",
		"-Itestdata",
		fmt.Sprintf("--fieldmask_out=%s", workDir),
		fmt.Sprintf("--gogo_out=%s", workDir),
	}, paths...)...)

	if err := filepath.Walk(workDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("Failed to read generated file at `%s`: %s", path, err)
			return nil
		}

		goldenPath := filepath.Join(".", strings.TrimPrefix(path, filepath.Join(workDir, "github.com", "TheThingsIndustries", "protoc-gen-fieldmask")))
		if *regenerate {
			if err := ioutil.WriteFile(goldenPath, b, 0666); err != nil {
				t.Errorf("Failed to write golden file at `%s`: %s", goldenPath, err)
			}
			return nil
		}

		golden, err := ioutil.ReadFile(goldenPath)
		if err != nil {
			t.Errorf("Failed to read golden file at `%s`: %s", goldenPath, err)
			return nil
		}

		if len(pretty.Diff(golden, b)) > 0 {
			pretty.Ldiff(t, golden, b)
			t.Fail()
		}
		return nil

	}); err != nil {
		t.Errorf("Failed to walk `%s`: %s", workDir, err)
		t.FailNow()
	}
}

func TestFieldMaskPaths(t *testing.T) {
	pb := &testdata.Test{}
	assertions.New(t).So(pb.FieldMaskPaths(), should.Resemble, []string{
		"a.a.a",
		"a.a.b",
		"a.a.c",
		"a.a.d",
		"a.b",
		"a.c",
		"a.d",
		"a.e",
		"b.a.a",
		"b.a.b",
		"b.a.c",
		"b.a.d",
		"b.b",
		"b.c",
		"b.d",
		"b.e",
		"c.a.a",
		"c.a.b",
		"c.a.c",
		"c.a.d",
		"c.b",
		"c.c",
		"c.d",
		"c.e",
	})
}

func TestSetFields(t *testing.T) {
	a := assertions.New(t)

	pb := &testdata.Test{
		CustomName: &testdata.Test_TestNested{
			A: &testdata.Test_TestNested_TestNestedNested{},
		},
	}
	pb.SetFields(&testdata.Test{
		A: &testdata.Test_TestNested{
			B: []byte{1, 2, 3},
		},
		CustomName: &testdata.Test_TestNested{
			B: []byte{1, 2, 4},
		},
	}, "a.b")
	a.So(pb, should.Resemble, &testdata.Test{
		A: &testdata.Test_TestNested{
			B: []byte{1, 2, 3},
		},
		CustomName: &testdata.Test_TestNested{
			A: &testdata.Test_TestNested_TestNestedNested{},
		},
	})
}
