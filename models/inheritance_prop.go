// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InheritanceProp inheritance prop
// swagger:model InheritanceProp
type InheritanceProp struct {

	// fullpath
	// Read Only: true
	Fullpath string `json:"fullpath,omitempty"`

	// id
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// type
	// Read Only: true
	Type string `json:"type,omitempty"`

	// value
	// Read Only: true
	Value string `json:"value,omitempty"`
}

// Validate validates this inheritance prop
func (m *InheritanceProp) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InheritanceProp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InheritanceProp) UnmarshalBinary(b []byte) error {
	var res InheritanceProp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}