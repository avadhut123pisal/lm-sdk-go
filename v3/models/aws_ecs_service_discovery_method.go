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
)

// AwsEcsServiceDiscoveryMethod aws ecs service discovery method
//
// swagger:model AwsEcsServiceDiscoveryMethod
type AwsEcsServiceDiscoveryMethod struct {
	AwsEcsServiceDiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *AwsEcsServiceDiscoveryMethod) Name() string {
	return "ad_awsecsservice"
}

// SetName sets the name of this subtype
func (m *AwsEcsServiceDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsEcsServiceDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		AwsEcsServiceDiscoveryMethodAllOf1
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

	var result AwsEcsServiceDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}
	result.AwsEcsServiceDiscoveryMethodAllOf1 = data.AwsEcsServiceDiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsEcsServiceDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AwsEcsServiceDiscoveryMethodAllOf1
	}{

		AwsEcsServiceDiscoveryMethodAllOf1: m.AwsEcsServiceDiscoveryMethodAllOf1,
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

// Validate validates this aws ecs service discovery method
func (m *AwsEcsServiceDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsEcsServiceDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this aws ecs service discovery method based on the context it is used
func (m *AwsEcsServiceDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsEcsServiceDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsEcsServiceDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsEcsServiceDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res AwsEcsServiceDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsEcsServiceDiscoveryMethodAllOf1 aws ecs service discovery method all of1
//
// swagger:model AwsEcsServiceDiscoveryMethodAllOf1
type AwsEcsServiceDiscoveryMethodAllOf1 interface{}