// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EscalatingChain escalating chain
// swagger:model EscalatingChain
type EscalatingChain struct {

	// cc destinations
	CcDestinations []*Recipient `json:"ccDestinations,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// destinations
	// Required: true
	Destinations []*Chain `json:"destinations"`

	// enable throttling
	EnableThrottling bool `json:"enableThrottling,omitempty"`

	// id
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// in alerting
	// Read Only: true
	InAlerting *bool `json:"inAlerting,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// throttling alerts
	ThrottlingAlerts int32 `json:"throttlingAlerts,omitempty"`

	// throttling period
	ThrottlingPeriod int32 `json:"throttlingPeriod,omitempty"`
}

// Validate validates this escalating chain
func (m *EscalatingChain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCcDestinations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EscalatingChain) validateCcDestinations(formats strfmt.Registry) error {
	if swag.IsZero(m.CcDestinations) { // not required
		return nil
	}

	for i := 0; i < len(m.CcDestinations); i++ {
		if swag.IsZero(m.CcDestinations[i]) { // not required
			continue
		}

		if m.CcDestinations[i] != nil {
			if err := m.CcDestinations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ccDestinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EscalatingChain) validateDestinations(formats strfmt.Registry) error {
	if err := validate.Required("destinations", "body", m.Destinations); err != nil {
		return err
	}

	for i := 0; i < len(m.Destinations); i++ {
		if swag.IsZero(m.Destinations[i]) { // not required
			continue
		}

		if m.Destinations[i] != nil {
			if err := m.Destinations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EscalatingChain) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EscalatingChain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EscalatingChain) UnmarshalBinary(b []byte) error {
	var res EscalatingChain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
