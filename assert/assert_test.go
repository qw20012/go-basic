package assert

import (
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, 123, 123)
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 123, 12)
}

func TestTrue(t *testing.T) {
	True(t, true)
}

func TestFalse(t *testing.T) {
	False(t, false)
}
