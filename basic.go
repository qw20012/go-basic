// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Basic utility functions.
package basic

import "github.com/qw20012/go-basic/ref"

// Get the exact value of any.
func FromAny[T any](value any) T {
	return value.(T)
}

// Make sure any type is created. Create by reflect if it is not there.
func NewIfEmpty[T any](value T) T {
	if !ref.IsNil(value) {
		return value
	}

	newValue := ref.ReflectNew(value)
	//creatd := newValue.Interface().(T)
	creatd := ref.GetValue[T](newValue)
	return creatd
}
