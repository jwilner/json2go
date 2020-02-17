// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
)

// Bar contains some info
// generated from https://example.com/testdata/generate/map_field/foo/bar.json
type Bar struct {
	Baz map[string]interface{} `json:"baz,omitempty"`
}

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
