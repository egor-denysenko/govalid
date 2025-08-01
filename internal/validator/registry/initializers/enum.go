// Code generated by generate-validators; DO NOT EDIT.
package initializers

import (
	"github.com/sivchari/govalid/internal/markers"
	"github.com/sivchari/govalid/internal/validator/registry"
	"github.com/sivchari/govalid/internal/validator/rules"
)

// EnumInitializer implements ValidatorInitializer for the enum validator.
type EnumInitializer struct{}

// Marker returns the marker identifier for the enum validator.
func (e EnumInitializer) Marker() string {
	return markers.GoValidMarkerEnum
}

// Init initializes the enum validator factory.
func (e EnumInitializer) Init() registry.ValidatorFactory {
	return rules.ValidateEnum
}
