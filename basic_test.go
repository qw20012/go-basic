package basic

import (
	"testing"
)

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
