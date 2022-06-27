package str

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.Fatal("IsEmpty failed")
	}

	if IsEmpty("abc") {
		t.Fatal("IsEmpty failed " + "abc")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty("") {
		t.Fatal("IsNotEmpty failed")
	}

	if !IsNotEmpty("abc") {
		t.Fatal("IsNotEmpty failed" + "abc")
	}
}

func TestContact(t *testing.T) {
	nilToTest := Contact()
	if nilToTest != "" {
		t.Fatal("TestContact failed")
	}

	empty := Contact("")
	if empty != "" {
		t.Fatal("TestContact failed")
	}

	one := Contact("abc")
	if one != "abc" {
		t.Fatal("TestContact failed " + "abc")
	}

	twoDiffType := Contact("abc", 1)
	if twoDiffType != "abc1" {
		t.Fatal("TestContact failed " + "abc, 1")
	}
}

func TestFrom(t *testing.T) {
	str := From("abc")
	if str != "abc" {
		t.Fatal("From failed " + "abc")
	}

	integer := From(1)
	if integer != "1" {
		t.Fatal("From failed " + "1")
	}

	f := From(1.123)
	if f != "1.123" {
		t.Fatal("From failed " + "1.123")
	}

	b := From(true)
	if b != "true" {
		t.Fatal("From failed " + "true")
	}
}

func TestFormat(t *testing.T) {
	emptySource := Format("", "name", "value")
	if emptySource != "" {
		t.Fatal("TestFormat failed " + "name" + "value")
	}

	emptyName := Format("abc {name}", "", "value")
	if emptyName != "abc {name}" {
		t.Fatal("TestFormat failed " + "value")
	}

	empatyValue := Format("abc {name}", "name", "")
	if empatyValue != "abc " {
		t.Fatal("TestFormat failed " + "name")
	}

	diffTypeValue := Format("abc {name}", "name", 1)
	if diffTypeValue != "abc 1" {
		t.Fatal("TestFormat failed " + "abc, 1 ")
	}
}

func TestFormats(t *testing.T) {

	nilMapStr := Formats("abc", nil)
	if nilMapStr != "abc" {
		t.Fatal("TestFormats failed " + "abc")
	}

	emptyMap := Formats("abc", make(map[string]any))
	if emptyMap != "abc" {
		t.Fatal("TestFormats failed " + "abc")
	}

	diffTypeValue := map[string]any{
		"a": "Dog",
		"b": 1,
	}
	strFromMap := Formats("{a}{ b }c", diffTypeValue)
	if strFromMap != "Dog1c" {
		t.Fatal("TestFormats failed " + "Dog1c")
	}
}
