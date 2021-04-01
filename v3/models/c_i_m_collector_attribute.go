// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CIMCollectorAttribute c i m collector attribute
//
// swagger:model CIMCollectorAttribute
type CIMCollectorAttribute struct {

	// fields
	Fields []*DataSourceAttribute `json:"fields,omitempty"`

	// ip
	// Read Only: true
	IP string `json:"ip,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// query class
	QueryClass string `json:"queryClass,omitempty"`

	// query index
	QueryIndex string `json:"queryIndex,omitempty"`

	// query value
	QueryValue string `json:"queryValue,omitempty"`
}

// Name gets the name of this subtype
func (m *CIMCollectorAttribute) Name() string {
	return "cim"
}

// SetName sets the name of this subtype
func (m *CIMCollectorAttribute) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CIMCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {

		// fields
		Fields []*DataSourceAttribute `json:"fields,omitempty"`

		// ip
		// Read Only: true
		IP string `json:"ip,omitempty"`

		// namespace
		Namespace string `json:"namespace,omitempty"`

		// query class
		QueryClass string `json:"queryClass,omitempty"`

		// query index
		QueryIndex string `json:"queryIndex,omitempty"`

		// query value
		QueryValue string `json:"queryValue,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result CIMCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.Fields = data.Fields
	result.IP = data.IP
	result.Namespace = data.Namespace
	result.QueryClass = data.QueryClass
	result.QueryIndex = data.QueryIndex
	result.QueryValue = data.QueryValue

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CIMCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// fields
		Fields []*DataSourceAttribute `json:"fields,omitempty"`

		// ip
		// Read Only: true
		IP string `json:"ip,omitempty"`

		// namespace
		Namespace string `json:"namespace,omitempty"`

		// query class
		QueryClass string `json:"queryClass,omitempty"`

		// query index
		QueryIndex string `json:"queryIndex,omitempty"`

		// query value
		QueryValue string `json:"queryValue,omitempty"`
	}{

		Fields: m.Fields,

		IP: m.IP,

		Namespace: m.Namespace,

		QueryClass: m.QueryClass,

		QueryIndex: m.QueryIndex,

		QueryValue: m.QueryValue,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this c i m collector attribute
func (m *CIMCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CIMCollectorAttribute) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	for i := 0; i < len(m.Fields); i++ {
		if swag.IsZero(m.Fields[i]) { // not required
			continue
		}

		if m.Fields[i] != nil {
			if err := m.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this c i m collector attribute based on the context it is used
func (m *CIMCollectorAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CIMCollectorAttribute) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fields); i++ {

		if m.Fields[i] != nil {
			if err := m.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CIMCollectorAttribute) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ip", "body", string(m.IP)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CIMCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CIMCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res CIMCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}