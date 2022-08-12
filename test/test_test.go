package test

import (
	"testing"
)

func TestEqual(t *testing.T) {
	if !Equal(t, 123, 123) {
		t.Fatalf("TestEqual failed")
	}
}

func TestNotEqual(t *testing.T) {
	if !NotEqual(t, 123, 12) {
		t.Fatalf("TestNotEqual failed")
	}
}

func TestTrue(t *testing.T) {
	if !True(t, true) {
		t.Fatalf("TestTrue failed")
	}
}

func TestFalse(t *testing.T) {
	if !False(t, false) {
		t.Fatalf("TestFalse failed")
	}
}
