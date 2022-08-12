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

func TestNewIfEmpty(t *testing.T) {
	// Test map
	var emptyAnyMap map[string]any
	fromEmptyAnyMap := NewIfEmpty(emptyAnyMap)
	if fromEmptyAnyMap == nil {
		t.Fatalf("NewIfEmpty with emtpy map failed")
	}
	fromEmptyAnyMap["key"] = 1

	var emptyIntMap map[string]int
	fromEmptyIntMap := NewIfEmpty(emptyIntMap)
	if fromEmptyIntMap == nil {
		t.Fatalf("NewIfEmpty with emtpy int map failed")
	}
	fromEmptyIntMap["key"] = 1

	notEmptyMap := make(map[string]int)
	fromNotEmptyMap := NewIfEmpty(notEmptyMap)
	if fromNotEmptyMap == nil {
		t.Fatalf("NewIfEmpty with not emtpy map failed")
	}
	fromNotEmptyMap["key"] = 1

	// Test Slice
	var emptySlice []int
	fromEmptySlice := NewIfEmpty(emptySlice)
	if fromEmptySlice == nil {
		t.Fatalf("NewIfEmpty with emtpy slice failed")
	}
	if len(append(fromEmptySlice, 1)) != 1 {
		t.Fatalf("NewIfEmpty with emtpy slice failed")
	}

	notEmptySlice := []int{1}
	fromNotEmptySlice := NewIfEmpty(notEmptySlice)
	if fromNotEmptySlice == nil {
		t.Fatalf("NewIfEmpty with emtpy slice failed")
	}
	if len(append(fromNotEmptySlice, 1)) != 2 {
		t.Fatalf("NewIfEmpty with emtpy slice failed")
	}
}
