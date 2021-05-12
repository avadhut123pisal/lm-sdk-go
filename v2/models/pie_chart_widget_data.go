// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PieChartWidgetData pie chart widget data
// swagger:model PieChartWidgetData
type PieChartWidgetData struct {
	titleField string

	// data
	// Read Only: true
	Data []*PieChartData `json:"data,omitempty"`

	// group remaining as others
	// Read Only: true
	GroupRemainingAsOthers *bool `json:"groupRemainingAsOthers,omitempty"`

	// hide zero percent slices
	HideZeroPercentSlices bool `json:"hideZeroPercentSlices,omitempty"`

	// max slices can be shown
	// Read Only: true
	MaxSlicesCanBeShown int32 `json:"maxSlicesCanBeShown,omitempty"`
}

// Title gets the title of this subtype
func (m *PieChartWidgetData) Title() string {
	return m.titleField
}

// SetTitle sets the title of this subtype
func (m *PieChartWidgetData) SetTitle(val string) {
	m.titleField = val
}

// Type gets the type of this subtype
func (m *PieChartWidgetData) Type() string {
	return "pieChart"
}

// SetType sets the type of this subtype
func (m *PieChartWidgetData) SetType(val string) {

}

// Data gets the data of this subtype

// GroupRemainingAsOthers gets the group remaining as others of this subtype

// HideZeroPercentSlices gets the hide zero percent slices of this subtype

// MaxSlicesCanBeShown gets the max slices can be shown of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PieChartWidgetData) UnmarshalJSON(raw []byte) error {
	var data struct {

		// data
		// Read Only: true
		Data []*PieChartData `json:"data,omitempty"`

		// group remaining as others
		// Read Only: true
		GroupRemainingAsOthers *bool `json:"groupRemainingAsOthers,omitempty"`

		// hide zero percent slices
		HideZeroPercentSlices bool `json:"hideZeroPercentSlices,omitempty"`

		// max slices can be shown
		// Read Only: true
		MaxSlicesCanBeShown int32 `json:"maxSlicesCanBeShown,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Title string `json:"title,omitempty"`

		Type string `json:"type,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PieChartWidgetData

	result.titleField = base.Title

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Data = data.Data

	result.GroupRemainingAsOthers = data.GroupRemainingAsOthers

	result.HideZeroPercentSlices = data.HideZeroPercentSlices

	result.MaxSlicesCanBeShown = data.MaxSlicesCanBeShown

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PieChartWidgetData) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// data
		// Read Only: true
		Data []*PieChartData `json:"data,omitempty"`

		// group remaining as others
		// Read Only: true
		GroupRemainingAsOthers *bool `json:"groupRemainingAsOthers,omitempty"`

		// hide zero percent slices
		HideZeroPercentSlices bool `json:"hideZeroPercentSlices,omitempty"`

		// max slices can be shown
		// Read Only: true
		MaxSlicesCanBeShown int32 `json:"maxSlicesCanBeShown,omitempty"`
	}{

		Data: m.Data,

		GroupRemainingAsOthers: m.GroupRemainingAsOthers,

		HideZeroPercentSlices: m.HideZeroPercentSlices,

		MaxSlicesCanBeShown: m.MaxSlicesCanBeShown,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Title string `json:"title,omitempty"`

		Type string `json:"type,omitempty"`
	}{

		Title: m.Title(),

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this pie chart widget data
func (m *PieChartWidgetData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PieChartWidgetData) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PieChartWidgetData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PieChartWidgetData) UnmarshalBinary(b []byte) error {
	var res PieChartWidgetData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
