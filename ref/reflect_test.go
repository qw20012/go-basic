package ref

import (
	"reflect"
	"testing"
)

func TestGetValue(t *testing.T) {
	if GetValue[int](reflect.ValueOf(1)) != 1 {
		t.Fatalf("TestGetValue with int type failed")
	}
}
func TestIsNil(t *testing.T) {
	// Test Array
	var emptyArray [1]int
	if IsNil(emptyArray) {
		t.Fatalf("TestIsNil with emtpy slice failed")
	}

	// Test Struct
	type Geek struct {
		A int `tag1:"First Tag" tag2:"Second Tag"`
		B string
	}
	var p Geek
	if IsNil(p) {
		t.Fatalf("TestIsNil with emtpy struct ")
	}

	ptr := &p
	if IsNil(ptr) {
		t.Fatalf("TestIsNil with emtpy struct pointer ")
	}
}

func TestIsZero(t *testing.T) {
	// Test Array
	var emptyArray [1]int
	if !IsZero(emptyArray) {
		t.Fatalf("TestIsZero with emtpy slice failed")
	}

	// Test Struct
	type Geek struct {
		A int `tag1:"First Tag" tag2:"Second Tag"`
		B string
	}
	var p Geek
	if !IsZero(p) {
		t.Fatalf("TestIsZero with emtpy struct")
	}

	ptr := &p
	if !IsZero(ptr) {
		t.Fatalf("TestIsZero with emtpy struct pointer ")
	}
}
