// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
	"regexp"
)

// Bar is generated from https://example.com/testdata/generate/map_schema/foo/bar.json
// Bar contains some info
type Bar struct {
	_   []byte
	Baz BarBaz `json:"baz,omitempty"`
	Biz BarBiz `json:"biz,omitempty"`
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/map_schema/foo/bar.json
func (m *Bar) Validate() error {
	return nil
}

// BarBizAdditionalProperties is generated from https://example.com/testdata/generate/map_schema/foo/bar.json#/properties/biz/additionalProperties
type BarBizAdditionalProperties struct {
	_  []byte
	ID *int64 `json:"id,omitempty"`
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/map_schema/foo/bar.json#/properties/biz/additionalProperties
func (m *BarBizAdditionalProperties) Validate() error {
	return nil
}

// BarBaz is generated from https://example.com/testdata/generate/map_schema/foo/bar.json#/properties/baz
type BarBaz map[string]string

var (
	barBazPattern = regexp.MustCompile(`^abc`)
)

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/map_schema/foo/bar.json#/properties/baz
func (m BarBaz) Validate() error {
	if len(m) < 3 {
		return &validationError{
			errType: "min_properties",
			message: "minimum of 3 properties",
		}
	}
	if len(m) > 10 {
		return &validationError{
			errType: "max_properties",
			message: "maximum of 10 properties",
		}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	strings.Sort(keys)
	for k := range keys {
		v := m[k]

		if !barBazPattern.MatchString(v) {
			return &validationError{
				path:     []interface{}{k},
				jsonPath: []interface{}{k},
				errType:  "pattern",
				message:  fmt.Sprintf(`must match '^abc' but got %q`, v),
			}
		}
	}
	return nil
}

// BarBiz is generated from https://example.com/testdata/generate/map_schema/foo/bar.json#/properties/biz
type BarBiz map[string]BarBizAdditionalProperties

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/map_schema/foo/bar.json#/properties/biz
func (m BarBiz) Validate() error {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	strings.Sort(keys)
	for k := range keys {
		v := m[k]

		if err := v.Validate(); err != nil {
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
