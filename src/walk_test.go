package main

import (
	"testing"
)

func TestWalkInsideDirectories(t *testing.T) {

	tests := []struct {
		path       string
		shouldFail bool
	}{
		{"tests/basic-tests", false},
		{"tests/no-rights-directory", true},
		{"there-is-no-directory-here", true},
	}

	for _, test := range tests {
		t.Log("Testing this path:", test.path)
		err := WalkInsideProject(test.path)
		if test.shouldFail {
			if err == nil {
				t.Error("This test should fail, here's what happen :", err)
			}
		} else {
			if err != nil {
				t.Error("This test shouldn't fail, here's what happen : ", err)
			}
		}
	}
}
