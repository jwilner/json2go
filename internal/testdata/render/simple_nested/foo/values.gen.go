// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
)

type A struct {
	B `json:"b,omitempty"`
}

func (m *A) Validate() error {
	if err := m.B.Validate(); err != nil {
		return err
	}
	return nil
}

type validationError struct {
	errType, jsonField, field, message string
}

func (e *validationError) ErrType() string {
	return e.errType
}

func (e *validationError) JSONField() string {
	return e.jsonField
}

func (e *validationError) Field() string {
	return e.field
}

func (e *validationError) Message() string {
	return e.message
}

func (e *validationError) Error() string {
	return fmt.Sprintf("%v: %v", e.field, e.message)
}
