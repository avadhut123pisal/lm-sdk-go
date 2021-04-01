// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WMICollectorAttribute w m i collector attribute
//
// swagger:model WMICollectorAttribute
type WMICollectorAttribute struct {

	// ip
	// Read Only: true
	IP string `json:"ip,omitempty"`

	// method inputs
	MethodInputs string `json:"methodInputs,omitempty"`

	// method name
	MethodName string `json:"methodName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// query class
	QueryClass string `json:"queryClass,omitempty"`

	// query index
	QueryIndex string `json:"queryIndex,omitempty"`

	// query value
	QueryValue string `json:"queryValue,omitempty"`

	// target path
	TargetPath string `json:"targetPath,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Name gets the name of this subtype
func (m *WMICollectorAttribute) Name() string {
	return "wmi"
}

// SetName sets the name of this subtype
func (m *WMICollectorAttribute) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *WMICollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {

		// ip
		// Read Only: true
		IP string `json:"ip,omitempty"`

		// method inputs
		MethodInputs string `json:"methodInputs,omitempty"`

		// method name
		MethodName string `json:"methodName,omitempty"`

		// namespace
		Namespace string `json:"namespace,omitempty"`

		// query class
		QueryClass string `json:"queryClass,omitempty"`

		// query index
		QueryIndex string `json:"queryIndex,omitempty"`

		// query value
		QueryValue string `json:"queryValue,omitempty"`

		// target path
		TargetPath string `json:"targetPath,omitempty"`

		// type
		Type string `json:"type,omitempty"`
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

	var result WMICollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.IP = data.IP
	result.MethodInputs = data.MethodInputs
	result.MethodName = data.MethodName
	result.Namespace = data.Namespace
	result.QueryClass = data.QueryClass
	result.QueryIndex = data.QueryIndex
	result.QueryValue = data.QueryValue
	result.TargetPath = data.TargetPath
	result.Type = data.Type

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m WMICollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// ip
		// Read Only: true
		IP string `json:"ip,omitempty"`

		// method inputs
		MethodInputs string `json:"methodInputs,omitempty"`

		// method name
		MethodName string `json:"methodName,omitempty"`

		// namespace
		Namespace string `json:"namespace,omitempty"`

		// query class
		QueryClass string `json:"queryClass,omitempty"`

		// query index
		QueryIndex string `json:"queryIndex,omitempty"`

		// query value
		QueryValue string `json:"queryValue,omitempty"`

		// target path
		TargetPath string `json:"targetPath,omitempty"`

		// type
		Type string `json:"type,omitempty"`
	}{

		IP: m.IP,

		MethodInputs: m.MethodInputs,

		MethodName: m.MethodName,

		Namespace: m.Namespace,

		QueryClass: m.QueryClass,

		QueryIndex: m.QueryIndex,

		QueryValue: m.QueryValue,

		TargetPath: m.TargetPath,

		Type: m.Type,
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

// Validate validates this w m i collector attribute
func (m *WMICollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this w m i collector attribute based on the context it is used
func (m *WMICollectorAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WMICollectorAttribute) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ip", "body", string(m.IP)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WMICollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WMICollectorAttribute) UnmarshalBinary(b []byte) error {
	var res WMICollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}