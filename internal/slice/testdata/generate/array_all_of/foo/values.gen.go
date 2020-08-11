// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
)

// Bar is generated from https://example.com/testdata/generate/array_all_of/foo/bar.json
// Bar gives you some dumb info
type Bar struct {
	_ []byte
	BarAllOf0
	BarAllOf1
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/array_all_of/foo/bar.json
func (m *Bar) Validate() error {
	if err := m.BarAllOf0.Validate(); err != nil {
		return err
	}
	if err := m.BarAllOf1.Validate(); err != nil {
		return err
	}
	return nil
}

// BarAllOf0 is generated from https://example.com/testdata/generate/array_all_of/foo/bar.json#/allOf/0
type BarAllOf0 struct {
	_  []byte
	ID *int64 `json:"id,omitempty"`
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/array_all_of/foo/bar.json#/allOf/0
func (m *BarAllOf0) Validate() error {
	return nil
}

// BarAllOf1 is generated from https://example.com/testdata/generate/array_all_of/foo/bar.json#/allOf/1
type BarAllOf1 struct {
	_    []byte
	Name *string `json:"name,omitempty"`
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/array_all_of/foo/bar.json#/allOf/1
func (m *BarAllOf1) Validate() error {
	return nil
}

// Barz is generated from https://example.com/testdata/generate/array_all_of/foo/barz.json
// Barz gives you lots of dumb info
type Barz []*Bar

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/array_all_of/foo/barz.json
func (m Barz) Validate() error {
	for i := range m {
		if err := m[i].Validate(); err != nil {
			if err, ok := err.(valErr); ok {
				return &validationError{
					errType:  err.ErrType(),
					message:  err.Message(),
					path:     append([]interface{}{i}, err.Path()...),
					jsonPath: append([]interface{}{i}, err.JSONPath()...),
				}
			}
			return err
		}
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
