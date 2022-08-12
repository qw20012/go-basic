// Original from https://github.com/gookit/goutil/tree/master/dump
// Package dump like fmt.Println but more pretty and beautiful print Go values.
package dump

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/qw20012/go-basic/color"
)

// These flags define which print caller information
const (
	Fnopos = 1 << iota // no position
	Ffunc
	Ffile
	Ffname
	Fline
)

var (
	// valid flag for print caller info
	callerFlags = []int{Ffunc, Ffile, Ffname, Fline}
	// default theme
	defaultTheme = Theme{
		"caller": "magenta",
		"field":  "green", // field name color of the map, struct.
		"value":  "normal",
		// special type
		"msType":  "green", // for keywords map, struct type
		"lenTip":  "gray",  // tips comments for string, slice, map len
		"string":  "green",
		"integer": "lightBlue",
	}

	// std dumper
	std = NewDumper(os.Stdout, 3)
	// no location dumper.
	std2 = NewWithOptions(func(opts *Options) {
		opts.Output = os.Stdout
		opts.ShowFlag = Fnopos
	})
)

// Theme color code/tag map for dump
type Theme map[string]string

func (ct Theme) caller(s string) string  { return color.MagentaString(s) } //{ return ct.wrap("caller", s) }
func (ct Theme) field(s string) string   { return color.GreenString(s) }   //{ return ct.wrap("field", s) }
func (ct Theme) value(s string) string   { return s }                      //{ return ct.wrap("value", s) }
func (ct Theme) msType(s string) string  { return color.GreenString(s) }   //{ return ct.wrap("msType", s) }
func (ct Theme) lenTip(s string) string  { return color.YellowString(s) }  //{ return ct.wrap("lenTip", s) }
func (ct Theme) string(s string) string  { return color.GreenString(s) }   //{ return ct.wrap("string", s) }
func (ct Theme) integer(s string) string { return color.BlueString(s) }    //{ return ct.wrap("integer", s) }

// WrapTag wrap a tag for a string "<tag>content</>"
func WrapTag(s string, tag string) string {
	if s == "" || tag == "" {
		return s
	}
	return fmt.Sprintf("<%s>%s</>", tag, s)
}

// wrap color tag.
func (ct Theme) wrap(key string, s string) string {
	if tag := ct[key]; tag != "" {
		return WrapTag(s, tag)
	}
	return s
}

// Std dumper
func Std() *Dumper {
	return std
}

// Reset std dumper
func Reset() {
	std = NewDumper(os.Stdout, 3)
}

// Config std dumper
func Config(fn func(opts *Options)) {
	std.WithOptions(fn)
}

// V like fmt.Println, but the output is clearer and more beautiful
func V(vs ...interface{}) {
	std.Dump(vs...)
}

// P like fmt.Println, but the output is clearer and more beautiful
func P(vs ...interface{}) {
	std.Print(vs...)
}

// Print like fmt.Println, but the output is clearer and more beautiful
func Print(vs ...interface{}) {
	std.Print(vs...)
}

// Println like fmt.Println, but the output is clearer and more beautiful
func Println(vs ...interface{}) {
	std.Println(vs...)
}

// Fprint like fmt.Println, but the output is clearer and more beautiful
func Fprint(w io.Writer, vs ...interface{}) {
	std.Fprint(w, vs...)
}

// Format like fmt.Println, but the output is clearer and more beautiful
func Format(vs ...interface{}) string {
	w := &bytes.Buffer{}

	std2.Fprint(w, vs...)
	return w.String()
}

// NoLoc dump vars data, without location.
func NoLoc(vs ...interface{}) {
	std2.Println(vs...)
}

// Clear dump clear data, without location.
func Clear(vs ...interface{}) {
	std2.Println(vs...)
}
