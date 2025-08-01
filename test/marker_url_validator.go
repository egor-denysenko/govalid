// Code generated by govalid; DO NOT EDIT.
package test

import (
	"errors"
	"github.com/sivchari/govalid/validation/validationhelper"
)

var (
	// ErrNilURL is returned when the URL is nil.
	ErrNilURL = errors.New("input URL is nil")

	// ErrURLURLURLValidation is the error returned when the field is not a valid URL.
	ErrURLURLURLValidation = errors.New("field URLURL must be a valid URL")
)

func ValidateURL(t *URL) error {
	if t == nil {
		return ErrNilURL
	}

	if !validationhelper.IsValidURL(t.URL) {
		return ErrURLURLURLValidation
	}

	return nil
}
