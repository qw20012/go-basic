package basic

import (
	"testing"
)

func TestIsNil(t *testing.T) {
	// Test Array
	var emptyArray [1]int
	if IsNil(emptyArray) {
		t.Fatalf("GetOrCreate with emtpy slice failed")
	}

	// Test Struct
	type Geek struct {
		A int `tag1:"First Tag" tag2:"Second Tag"`
		B string
	}
	var p Geek
	if IsNil(p) {
		t.Fatalf("GetOrCreate with emtpy struct pointer ")
	}

}

func TestGetOrCreate(t *testing.T) {
	// Test map
	var emptyAnyMap map[string]any
	fromEmptyAnyMap := NewIfEmpty(emptyAnyMap)
	if fromEmptyAnyMap == nil {
		t.Fatalf("GetOrCreate with emtpy map failed")
	}
	fromEmptyAnyMap["key"] = 1

	var emptyIntMap map[string]int
	fromEmptyIntMap := NewIfEmpty(emptyIntMap)
	if fromEmptyIntMap == nil {
		t.Fatalf("GetOrCreate with emtpy int map failed")
	}
	fromEmptyIntMap["key"] = 1

	notEmptyMap := make(map[string]int)
	fromNotEmptyMap := NewIfEmpty(notEmptyMap)
	if fromNotEmptyMap == nil {
		t.Fatalf("GetOrCreate with not emtpy map failed")
	}
	fromNotEmptyMap["key"] = 1

	// Test Slice
	var emptySlice []int
	fromEmptySlice := NewIfEmpty(emptySlice)
	if fromEmptySlice == nil {
		t.Fatalf("GetOrCreate with emtpy slice failed")
	}
	if len(append(fromEmptySlice, 1)) != 1 {
		t.Fatalf("GetOrCreate with emtpy slice failed")
	}

	notEmptySlice := []int{1}
	fromNotEmptySlice := NewIfEmpty(notEmptySlice)
	if fromNotEmptySlice == nil {
		t.Fatalf("GetOrCreate with emtpy slice failed")
	}
	if len(append(fromNotEmptySlice, 1)) != 2 {
		t.Fatalf("GetOrCreate with emtpy slice failed")
	}
}
