// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
)

// Bar gives you some dumb info
// generated from https://example.com/testdata/generate/double_nested/foo/bar.json
type Bar struct {
	Foo Foo `json:"foo,omitempty"`
}

func (m *Bar) Validate() error {
	if err := m.Foo.Validate(); err != nil {
		if err, ok := err.(valErr); ok {
			return &validationError{
				errType:  err.ErrType(),
				message:  err.Message(),
				path:     append([]interface{}{"Foo"}, err.Path()...),
				jsonPath: append([]interface{}{"foo"}, err.JSONPath()...),
			}
		}
		return err
	}
	return nil
}

// generated from https://example.com/testdata/generate/double_nested/foo/baz.json
type Baz struct {
	Name *string `json:"name,omitempty"`
}

func (m *Baz) Validate() error {
	return nil
}

// generated from https://example.com/testdata/generate/double_nested/foo/bar.json#/properties/foo
type Foo struct {
	Baz Baz `json:"baz,omitempty"`
}

func (m *Foo) Validate() error {
	if err := m.Baz.Validate(); err != nil {
		if err, ok := err.(valErr); ok {
			return &validationError{
				errType:  err.ErrType(),
				message:  err.Message(),
				path:     append([]interface{}{"Baz"}, err.Path()...),
				jsonPath: append([]interface{}{"baz"}, err.JSONPath()...),
			}
		}
		return err
	}
	return nil
}

type valErr interface {
	ErrType() string
	JSONPath() []interface{}
	Path() []interface{}
	Message() string
}

type validationError struct {
	errType, message string
	jsonPath, path   []interface{}
}

func (e *validationError) ErrType() string {
	return e.errType
}

func (e *validationError) JSONPath() []interface{} {
	return e.jsonPath
}

func (e *validationError) Path() []interface{} {
	return e.path
}

func (e *validationError) Message() string {
	return e.message
}

func (e *validationError) Error() string {
	return fmt.Sprintf("%v: %v", e.path, e.message)
}

var _ valErr = new(validationError)
