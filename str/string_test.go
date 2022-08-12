package str

import (
	"testing"

	"github.com/qw20012/go-basic/test"
)

func TestIsEmpty(t *testing.T) {
	test.True(t, IsEmpty(""))
	test.False(t, IsEmpty("abc"))
}

func TestIsNotEmpty(t *testing.T) {
	test.False(t, IsNotEmpty(""))
	test.True(t, IsNotEmpty("abc"))
}

func TestContact(t *testing.T) {
	test.Equal(t, Empty, Contact())
	test.Equal(t, Empty, Contact(""))
	test.Equal(t, "abc", Contact("abc"))
	test.Equal(t, "abc1", Contact("abc", 1))
}

func TestFrom(t *testing.T) {
	test.Equal(t, "abc", From("abc"))
	test.Equal(t, "1", From(1))
	test.Equal(t, "1.123", From(1.123))
	test.Equal(t, "true", From(true))
}

func TestFormat(t *testing.T) {
	test.Equal(t, "", Format("", "name", "value"))
	test.Equal(t, "abc {name}", Format("abc {name}", "", "value"))
	test.Equal(t, "abc ", Format("abc {name}", "name", ""))
	test.Equal(t, "abc 1", Format("abc {name}", "name", 1))
	test.Equal(t, "abc 1 {name}", Format("abc %v {name}", "n", 1))
}

func TestFormats(t *testing.T) {
	test.Equal(t, "abc", Formats("abc", nil))
	test.Equal(t, "abc", Formats("abc", make(map[string]any)))

	diffTypeValue := map[string]any{
		"a": "Dog",
		"b": 1,
	}
	test.Equal(t, "Dog1c", Formats("{a}{ b }c", diffTypeValue))
}

func TestRepeatRune(t *testing.T) {
	tests := []struct {
		want  []rune
		give  rune
		times int
	}{
		{[]rune("bbb"), 'b', 3},
		{[]rune("..."), '.', 3},
		{[]rune("  "), ' ', 2},
	}

	for _, tt := range tests {
		test.Equal(t, tt.want, RepeatRune(tt.give, tt.times))
	}
}
func TestPadding(t *testing.T) {
	tests := []struct {
		want, give, pad string
		len             int
		pos             bool
	}{
		{"ab000", "ab", "0", 5, true},
		{"000ab", "ab", "0", 5, false},
		{"ab012", "ab012", "0", 4, false},
		{"ab   ", "ab", "", 5, true},
		{"   ab", "ab", "", 5, false},
	}

	for _, tt := range tests {
		test.Equal(t, tt.want, Padding(tt.give, tt.pad, tt.len, tt.pos))
		if tt.pos {
			test.Equal(t, tt.want, PadRight(tt.give, tt.pad, tt.len))
		} else {
			test.Equal(t, tt.want, PadLeft(tt.give, tt.pad, tt.len))
		}
	}
}
