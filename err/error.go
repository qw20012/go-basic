// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Help error handling gracefully.
package err

import (
	"github.com/qw20012/go-basic/ref"
	"github.com/qw20012/go-basic/str"
)

type BqError struct {
	id         string
	message    string
	parameters map[string]any
	cause      error
}

// Stack trace error message.
func (e *BqError) Error() string {
	msg := e.Message()

	if e.cause != nil {
		msg = str.Contact(msg, "\n", e.cause.Error())
	}

	return msg
}

// New BqError pointer with given id and message.
// Parameter cause and parameters are optional.
func New(id string, params ...any) *BqError {
	err := &BqError{
		id: id,
	}

	for _, param := range params {
		switch param := param.(type) {
		case string:
			err.message = param
		case error:
			err.cause = param
		case map[string]any:
			err.parameters = param
		default:
		}
	}
	return err
}

// Wrap message with given prefix.
func (e *BqError) Wrap(prefix string) *BqError {
	e.message = str.Contact(prefix, e.message)
	return e
}

// Set parameter in build pattern.
func (e *BqError) WithParameter(name string, value any) *BqError {
	e.parameters = ref.NewIfEmpty(e.parameters)
	e.parameters[name] = value
	return e
}

// Get id field.
func (e *BqError) Id() string {
	return e.id
}

// Get cause field.
func (e *BqError) Cause() error {
	return e.cause
}

// Get formated message with parameters.
func (e *BqError) Message() string {
	msg := e.message
	if len(e.parameters) > 0 && str.IsNotEmpty(msg) {
		msg = str.Formats(msg, e.parameters)
	}

	return msg
}
