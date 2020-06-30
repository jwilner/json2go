// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
)

// Bar is generated from https://example.com/testdata/generate/multiline_description/foo/bar.json
// Bar gives you some dumb info
//
// Wheee
type Bar struct {
	Blah *string `json:"blah,omitempty"`
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/multiline_description/foo/bar.json
func (m *Bar) Validate() error {
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