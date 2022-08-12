package testing

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

func TestNil(t *testing.T) {
	if !Nil(t, nil) {
		t.Fatalf("TestNil failed")
	}
}

func TestNotNil(t *testing.T) {
	if !NotNil(t, 1) {
		t.Fatalf("TestNotNil failed")
	}
}

func TestDiff(t *testing.T) {
	expected :=
		`

Diff:
--- Expected
+++ Actual
@@ -1,3 +1,3 @@
 {
-    "Foo": "hello"
+    "Foo": "bar"
 }
`
	actual := diff(
		struct{ Foo string }{"hello"},
		struct{ Foo string }{"bar"},
	)
	//fmt.Println(actual)
	//Equal(t, expected, actual)
	Equal(t, expected, actual, "hello %s", "world")
}
