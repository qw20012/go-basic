package testing

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"reflect"

	"strings"
	"testing"
	"time"

	"github.com/qw20012/go-basic"
	dif "github.com/qw20012/go-basic/diff"
	"github.com/qw20012/go-basic/ref"
	"github.com/qw20012/go-basic/str"
)

func isFunction(arg any) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}

// validateEqualArgs checks whether provided arguments can be safely used in the
// Equal/NotEqual functions.
func validateEqualArgs(expected, actual any) error {
	if expected == nil && actual == nil {
		return nil
	}

	if isFunction(expected) || isFunction(actual) {
		return errors.New("cannot take func type as argument")
	}
	return nil
}

// Aligns the provided message so that all lines after the first line start at the same location as the first line.
// Assumes that the first line starts at the correct location (after carriage return, tab, label, spacer and tab).
// The longestLabelLen parameter specifies the length of the longest label in the output (required becaues this is the
// basis on which the alignment occurs).
func indentMessageLines(message string, longestLabelLen int) string {
	outBuf := new(bytes.Buffer)

	for i, scanner := 0, bufio.NewScanner(strings.NewReader(message)); scanner.Scan(); i++ {
		// no need to align first line because it starts at the correct location (after the label)
		if i != 0 {
			// append alignLen+1 spaces to align with "{{longestLabel}}:" before adding tab
			outBuf.WriteString("\n\t" + strings.Repeat(" ", longestLabelLen+1) + "\t")
		}
		outBuf.WriteString(scanner.Text())
	}

	return outBuf.String()
}

// labeledOutput returns a string consisting of the provided labeledContent. Each labeled output is appended in the following manner:
//
//   \t{{label}}:{{align_spaces}}\t{{content}}\n
//
// The initial carriage return is required to undo/erase any padding added by testing.T.Errorf. The "\t{{label}}:" is for the label.
// If a label is shorter than the longest label provided, padding spaces are added to make all the labels match in length. Once this
// alignment is achieved, "\t{{content}}\n" is added for the output.
//
// If the content of the labeledOutput contains line breaks, the subsequent lines are aligned so that they start at the same location as the first line.
func labeledOutput(content map[string]string) string {
	longestLabel := 0
	for k := range content {
		if len(k) > longestLabel {
			longestLabel = len(k)
		}
	}

	var output string
	for k, v := range content {
		output += "\t" + k + ":" + strings.Repeat(" ", longestLabel-len(k)) + "\t" +
			indentMessageLines(v, longestLabel) + "\n"
	}
	return output
}

func messageFromMsgAndArgs(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return str.Empty
	}
	if len(msgAndArgs) == 1 {
		return str.From(msgAndArgs[0])
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return str.Empty
}

// Fail reports a failure through
func Fail(t *testing.T, failureMessage string, msgAndArgs ...any) bool {
	content := map[string]string{
		"Error Trace": strings.Join(basic.CallerInfo(), "\n\t\t\t"),
		"Error":       failureMessage,
	}

	message := messageFromMsgAndArgs(msgAndArgs...)
	if len(message) > 0 {
		content["Messages"] = message
	}

	t.Errorf("\n%s", labeledOutput(content))

	return false
}

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding the
// bufio.MaxScanTokenSize max line length that the go testing framework imposes.
func truncatingFormat(data any) string {
	value := fmt.Sprintf("%#v", data)
	max := bufio.MaxScanTokenSize - 100 // Give us some space the type info too if needed.
	if len(value) > max {
		value = value[0:max] + "<... truncated>"
	}
	return value
}

// formatUnequalValues takes two values of arbitrary types and returns string
// representations appropriate to be presented to the user.
//
// If the values are not of like type, the returned strings will be prefixed
// with the type name, and the value will be enclosed in parenthesis similar
// to a type conversion in the Go grammar.
func formatUnequalValues(expected, actual any) (e, a string) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		return fmt.Sprintf("%T(%s)", expected, truncatingFormat(expected)),
			fmt.Sprintf("%T(%s)", actual, truncatingFormat(actual))
	}
	switch expected.(type) {
	case time.Duration:
		return fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)
	}
	return truncatingFormat(expected), truncatingFormat(actual)
}

// diff returns a diff of both values as long as both are of the same type and
// are a struct, map, slice, array or string. Otherwise it returns an empty string.
func diff(expected interface{}, actual interface{}) string {
	if expected == nil || actual == nil {
		return ""
	}

	et, ek, _ := ref.Parse(expected)
	at, _, _ := ref.Parse(actual)

	if et != at {
		return ""
	}

	if ek != reflect.Struct && ek != reflect.Map && ek != reflect.Slice && ek != reflect.Array && ek != reflect.String {
		return ""
	}

	var e, a string

	switch et {
	case reflect.TypeOf(""):
		e = reflect.ValueOf(expected).String()
		a = reflect.ValueOf(actual).String()
	//case reflect.TypeOf(time.Time{}):
	//e = spewConfigStringerEnabled.Sdump(expected)
	//a = spewConfigStringerEnabled.Sdump(actual)
	default:
		//e = spewConfig.Sdump(expected)
		js, _ := json.MarshalIndent(expected, "", "    ")
		e = string(js)
		//e = fmt.Sprintf("%#v", e)
		//fmt.Println(expected)
		//a = spewConfig.Sdump(actual)
		js, _ = json.MarshalIndent(actual, "", "    ")
		a = string(js)
		//a = fmt.Sprintf("%#v", actual)
		//fmt.Println(a)
	}
	d, _ := dif.GetUnifiedDiffString(dif.UnifiedDiff{
		A:        dif.SplitLines(e),
		B:        dif.SplitLines(a),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	})

	return "\n\nDiff:\n" + d
}

// Equal asserts that two objects are equal.
//
//    assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equal(t *testing.T, expected, actual any, msgAndArgs ...any) bool {

	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, err), msgAndArgs...)
	}

	if !basic.ObjectsAreEqual(expected, actual) {
		diff := diff(expected, actual)
		expected, actual = formatUnequalValues(expected, actual)
		return Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s%s", expected, actual, diff), msgAndArgs...)
	}

	return true

}

// NotEqual asserts that the specified values are NOT equal.
//
//    assert.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqual(t *testing.T, expected, actual any, msgAndArgs ...any) bool {
	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("Invalid operation: %#v != %#v (%s)",
			expected, actual, err), msgAndArgs...)
	}

	if basic.ObjectsAreEqual(expected, actual) {
		return Fail(t, fmt.Sprintf("Should not be: %#v\n", actual), msgAndArgs...)
	}

	return true

}

// NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, err)
func NotNil(t *testing.T, object any, msgAndArgs ...any) bool {
	if !ref.IsNil(object) {
		return true
	}
	return Fail(t, "Expected value not to be nil.", msgAndArgs...)
}

// Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
func Nil(t *testing.T, object any, msgAndArgs ...any) bool {
	if ref.IsNil(object) {
		return true
	}

	return Fail(t, fmt.Sprintf("Expected nil, but got: %#v", object), msgAndArgs...)
}
