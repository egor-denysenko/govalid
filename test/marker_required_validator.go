// Code generated by govalid; DO NOT EDIT.
package test

import (
	"errors"
)

var (
	// ErrNilRequired is returned when the Required is nil.
	ErrNilRequired = errors.New("input Required is nil")

	// ErrRequiredNameRequiredValidation is returned when the RequiredName is required but not provided.
	ErrRequiredNameRequiredValidation = errors.New("field RequiredName is required")

	// ErrRequiredAgeRequiredValidation is returned when the RequiredAge is required but not provided.
	ErrRequiredAgeRequiredValidation = errors.New("field RequiredAge is required")

	// ErrRequiredItemsRequiredValidation is returned when the RequiredItems is required but not provided.
	ErrRequiredItemsRequiredValidation = errors.New("field RequiredItems is required")
)

func ValidateRequired(t *Required) error {
	if t == nil {
		return ErrNilRequired
	}

	if t.Name == "" {
		return ErrRequiredNameRequiredValidation
	}

	if t.Age == 0 {
		return ErrRequiredAgeRequiredValidation
	}

	if t.Items == nil {
		return ErrRequiredItemsRequiredValidation
	}

	return nil
}
