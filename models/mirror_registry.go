// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MirrorRegistry mirror registry
//
// swagger:model mirror_registry
type MirrorRegistry struct {

	// the original registry location
	Location string `json:"location,omitempty"`

	// the mirror regsitry location
	MirrorLocation string `json:"mirror_location,omitempty"`

	// prefix for choosing this specific mirror
	Prefix string `json:"prefix,omitempty"`
}

// Validate validates this mirror registry
func (m *MirrorRegistry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MirrorRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MirrorRegistry) UnmarshalBinary(b []byte) error {
	var res MirrorRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
