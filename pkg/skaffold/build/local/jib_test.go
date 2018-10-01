/*
Copyright 2018 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package local

import (
	"testing"
)

func TestBuildGradle(t *testing.T) {
	t.Error("unimplemented")
}

func TestGenerateJibImageRef_simple_noProject(t *testing.T) {
	imageName := generateJibImageRef("simple", "")
	assertEquals(t, "jibsimple", imageName)
}

func TestGenerateJibImageRef_simple_withProject(t *testing.T) {
	imageName := generateJibImageRef("simple", "project")
	assertEquals(t, "jibsimple_project", imageName)
}

func TestGenerateJibImageRef_complex(t *testing.T) {
	imageName := generateJibImageRef("complex/workspace", "project")
	assertEquals(t, "jib__965ec099f720d3ccc9c038c21ea4a598c9632883", imageName)
}

func assertEquals(t *testing.T, expected, computed string) {
	if expected != computed {
		t.Errorf("Expected '%s': '%s'", expected, computed)
	}
}

