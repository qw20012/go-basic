// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Basic utility functions.
package basic

import (
	"reflect"
)

// Identify whether the any type is nil.
func IsNil(value any) bool {
	if value == nil {
		return true
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(value).IsNil()
	}

	return false
}

// Make sure any type is created. Create by reflect if it is not there.
func NewIfEmpty[T any](value T) T {
	if !IsNil(value) {
		return value
	}

	mapReflect := DynamicNew(value)
	creatd := mapReflect.Interface().(T)
	return creatd
}

// Todo: Would complete this function in future.
func DynamicNew(value any) reflect.Value {
	ty := reflect.TypeOf(value)
	switch ty.Kind() {
	case reflect.Interface:
		return reflect.MakeMap(ty).Elem()
	case reflect.Ptr:
		return reflect.New(ty.Elem()).Elem()
	case reflect.Map:
		return reflect.MakeMap(ty)
	case reflect.Struct:
		return reflect.New(ty.Elem()).Elem()
	case reflect.Slice:
		return reflect.MakeSlice(ty, 0, 0)
	}

	return reflect.ValueOf(value)
}
