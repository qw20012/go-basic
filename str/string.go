// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Utility to manipulate string.
package str

import (
	"fmt"
	"regexp"
	"strings"
)

// Represents the emptry string.
const Empty = ""

// Identify whether the source string is empty.
func IsEmpty(source string) bool {
	return source == Empty
}

// Identify whether the source string is not empty.
func IsNotEmpty(source string) bool {
	return !IsEmpty(source)
}

// Contact the sources from any type.
func Contact(sources ...any) string {
	count := len(sources)
	if count <= 0 {
		return Empty
	}

	format := strings.Repeat("%v", count)
	return fmt.Sprintf(format, sources...)
}

// Convert to string from any type.
func From(value any) string {
	return fmt.Sprintf("%v", value)
}

// Format source string that instead given name in curly brackets by given value.
//
// Example:
//     Format("abc {name}", "name", 1) -> "abc 1"
func Format(source string, name string, value any) string {
	if IsEmpty(source) {
		return source
	}
	pattern := Contact("{\\s*", name, "\\s*}")
	r, _ := regexp.Compile(pattern)

	matched := r.FindString(source)
	if IsNotEmpty(matched) {
		return strings.Replace(source, matched, From(value), -1)
	}

	if strings.Contains(source, "%") {
		return fmt.Sprintf(source, value)
	}

	return source
}

// Format source string by calling Format functon.
// See also Format.
func Formats(source string, params map[string]any) string {
	for key, val := range params {
		source = Format(source, key, val)
	}
	return source
}

/*************************************************************
 * String repeat operation
 *************************************************************/

// RepeatRune repeat a rune char.
func RepeatRune(char rune, count int) (chars []rune) {
	for i := 0; i < count; i++ {
		chars = append(chars, char)
	}
	return
}

// RepeatBytes repeat a byte char.
func RepeatBytes(char byte, times int) (chars []byte) {
	for i := 0; i < times; i++ {
		chars = append(chars, char)
	}
	return
}

/*************************************************************
 * String padding operation
 *************************************************************/

// Padding a string.
func Padding(s, pad string, length int, isRight bool) string {
	diff := len(s) - length
	if diff >= 0 { // do not need padding.
		return s
	}

	if pad == "" || pad == " " {
		mark := ""
		if isRight { // to right
			mark = "-"
		}

		// padding left: "%7s", padding right: "%-7s"
		tpl := fmt.Sprintf("%s%d", mark, length)
		return fmt.Sprintf(`%`+tpl+`s`, s)
	}

	if isRight { // to right
		return s + strings.Repeat(pad, -diff)
	}

	return strings.Repeat(pad, -diff) + s
}

// PadLeft a string.
func PadLeft(s, pad string, length int) string {
	return Padding(s, pad, length, false)
}

// PadRight a string.
func PadRight(s, pad string, length int) string {
	return Padding(s, pad, length, true)
}

// WrapTag for given string.
func WrapTag(s, tag string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("<%s>%s</%s>", tag, s, tag)
}

/*
// Any formats any value as a string.
func Any(value interface{}) string {
    return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
    switch v.Kind() {
    case reflect.Invalid:
        return "invalid"
    case reflect.Int, reflect.Int8, reflect.Int16,
        reflect.Int32, reflect.Int64:
        return strconv.FormatInt(v.Int(), 10)
    case reflect.Uint, reflect.Uint8, reflect.Uint16,
        reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return strconv.FormatUint(v.Uint(), 10)
    // ...floating-point and complex cases omitted for brevity...
    case reflect.Bool:
        return strconv.FormatBool(v.Bool())
    case reflect.String:
        return strconv.Quote(v.String())
    case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
        return v.Type().String() + " 0x" +
            strconv.FormatUint(uint64(v.Pointer()), 16)
    default: // reflect.Array, reflect.Struct, reflect.Interface
        return v.Type().String() + " value"
    }
}
*/
