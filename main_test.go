package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	if err := run("unknown", "fake", 0); err != nil {
		t.Fatalf(`TestMain, %v, want nil`, err)
	}
}
