package arr

import (
	"testing"
)

func TestContains(t *testing.T) {
	// Test Array
	var aArray = [3]int{1, 2, 3}
	if !Contains(aArray[:], 2) {
		t.Fatalf("TestContains failed")
	}
}
