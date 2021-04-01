// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BigNumberData big number data
//
// swagger:model BigNumberData
type BigNumberData struct {

	// bottom label
	// Read Only: true
	BottomLabel string `json:"bottomLabel,omitempty"`

	// color level
	// Read Only: true
	ColorLevel int32 `json:"colorLevel,omitempty"`

	// error message
	// Read Only: true
	ErrorMessage string `json:"errorMessage,omitempty"`

	// position
	// Read Only: true
	Position int32 `json:"position,omitempty"`

	// right label
	// Read Only: true
	RightLabel string `json:"rightLabel,omitempty"`

	// rounding
	// Read Only: true
	Rounding int32 `json:"rounding,omitempty"`

	// use comma separators
	// Required: true
	UseCommaSeparators *bool `json:"useCommaSeparators"`

	// value
	// Read Only: true
	Value float64 `json:"value,omitempty"`
}

// Validate validates this big number data
func (m *BigNumberData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUseCommaSeparators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BigNumberData) validateUseCommaSeparators(formats strfmt.Registry) error {

	if err := validate.Required("useCommaSeparators", "body", m.UseCommaSeparators); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this big number data based on the context it is used
func (m *BigNumberData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBottomLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateColorLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRightLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRounding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BigNumberData) contextValidateBottomLabel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "bottomLabel", "body", string(m.BottomLabel)); err != nil {
		return err
	}

	return nil
}

func (m *BigNumberData) contextValidateColorLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "colorLevel", "body", int32(m.ColorLevel)); err != nil {
		return err
	}

	return nil
}

func (m *BigNumberData) contextValidateErrorMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "errorMessage", "body", string(m.ErrorMessage)); err != nil {
		return err
	}

	return nil
}

func (m *BigNumberData) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "position", "body", int32(m.Position)); err != nil {
		return err
	}

	return nil
}

func (m *BigNumberData) contextValidateRightLabel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rightLabel", "body", string(m.RightLabel)); err != nil {
		return err
	}

	return nil
}

func (m *BigNumberData) contextValidateRounding(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rounding", "body", int32(m.Rounding)); err != nil {
		return err
	}

	return nil
}

func (m *BigNumberData) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", float64(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BigNumberData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BigNumberData) UnmarshalBinary(b []byte) error {
	var res BigNumberData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}