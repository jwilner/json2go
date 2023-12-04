// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
)

// Bar is generated from https://example.com/testdata/generate/omitempty_array/foo/bar.json
// Bar gives you null value
type Bar struct {
	Slice BarSlice `json:"slice,omitempty"`
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/omitempty_array/foo/bar.json
func (m *Bar) Validate() error {
	if err := m.Slice.Validate(); err != nil {
		if err, ok := err.(valErr); ok {
			return &validationError{
				errType:  err.ErrType(),
				message:  err.Message(),
				path:     append([]interface{}{"Slice"}, err.Path()...),
				jsonPath: append([]interface{}{"slice"}, err.JSONPath()...),
			}
		}
		return err
	}
	return nil
}

// BarSlice is generated from https://example.com/testdata/generate/omitempty_array/foo/bar.json#/properties/slice
// literally nothing interesting
type BarSlice []interface{}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/omitempty_array/foo/bar.json#/properties/slice
func (m BarSlice) Validate() error {
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