package err

import (
	"errors"
	"testing"
)

func create(t *testing.T) *BqError {
	bqError := New("test", "test error")
	if bqError == nil || bqError.Id() != "test" {
		t.Fatal("Test New failed")
	}
	return bqError
}

func TestError(t *testing.T) {
	params := make(map[string]any)
	params["p1"] = 1
	params["p2"] = "2"
	bqError := New("test", "Test error with {p1} and {p2}", errors.New("Nested error"), params)
	if bqError == nil || bqError.Id() != "test" {
		t.Fatal("Test New failed")
	}

	if bqError.Error() != `Test error with 1 and 2
Nested error` {
		t.Fatal("TestError failed")

	}
}

func TestWrap(t *testing.T) {
	bqError := create(t)
	bqError.Wrap("Wrap ")
	if bqError.Message() != "Wrap test error" {
		t.Fatal("TestWrap failed")
	}
}

func WithParameter(t *testing.T) {
	bqError := create(t)
	bqError.Wrap("Wrap {p}")
	bqError.WithParameter("p", 1)
	if bqError.Error() != `Wrap 1 test error` {
		t.Fatal("WithParameter failed")
	}
}
