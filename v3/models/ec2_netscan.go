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

// Ec2Netscan ec2 netscan
//
// swagger:model Ec2Netscan
type Ec2Netscan struct {
	collectorField *int32

	collectorDescriptionField string

	collectorGroupField int32

	collectorGroupNameField string

	creatorField string

	descriptionField string

	duplicateField *ExcludeDuplicateIps

	groupField string

	idField int32

	ignoreSystemIPsDuplicatesField bool

	nameField *string

	nextStartField string

	nextStartEpochField int64

	nsgIdField int32

	scheduleField *RestSchedule

	versionField int32

	// Which IP the EC2 instance should be monitored with for nec2 scans: private or public
	// Example: private
	// Required: true
	Accessibility *string `json:"accessibility"`

	// The credentials to be used for the scan
	// Required: true
	Credentials *EC2NetscanPolicyCredential `json:"credentials"`

	// ddr
	Ddr *Ec2DDR `json:"ddr,omitempty"`

	// How dead EC2 instances should be handled for nec2 scans. Must be Manually
	// Example: MANUALLY
	DeadOperation string `json:"deadOperation,omitempty"`

	// ports
	Ports *NetscanPorts `json:"ports,omitempty"`
}

// Collector gets the collector of this subtype
func (m *Ec2Netscan) Collector() *int32 {
	return m.collectorField
}

// SetCollector sets the collector of this subtype
func (m *Ec2Netscan) SetCollector(val *int32) {
	m.collectorField = val
}

// CollectorDescription gets the collector description of this subtype
func (m *Ec2Netscan) CollectorDescription() string {
	return m.collectorDescriptionField
}

// SetCollectorDescription sets the collector description of this subtype
func (m *Ec2Netscan) SetCollectorDescription(val string) {
	m.collectorDescriptionField = val
}

// CollectorGroup gets the collector group of this subtype
func (m *Ec2Netscan) CollectorGroup() int32 {
	return m.collectorGroupField
}

// SetCollectorGroup sets the collector group of this subtype
func (m *Ec2Netscan) SetCollectorGroup(val int32) {
	m.collectorGroupField = val
}

// CollectorGroupName gets the collector group name of this subtype
func (m *Ec2Netscan) CollectorGroupName() string {
	return m.collectorGroupNameField
}

// SetCollectorGroupName sets the collector group name of this subtype
func (m *Ec2Netscan) SetCollectorGroupName(val string) {
	m.collectorGroupNameField = val
}

// Creator gets the creator of this subtype
func (m *Ec2Netscan) Creator() string {
	return m.creatorField
}

// SetCreator sets the creator of this subtype
func (m *Ec2Netscan) SetCreator(val string) {
	m.creatorField = val
}

// Description gets the description of this subtype
func (m *Ec2Netscan) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *Ec2Netscan) SetDescription(val string) {
	m.descriptionField = val
}

// Duplicate gets the duplicate of this subtype
func (m *Ec2Netscan) Duplicate() *ExcludeDuplicateIps {
	return m.duplicateField
}

// SetDuplicate sets the duplicate of this subtype
func (m *Ec2Netscan) SetDuplicate(val *ExcludeDuplicateIps) {
	m.duplicateField = val
}

// Group gets the group of this subtype
func (m *Ec2Netscan) Group() string {
	return m.groupField
}

// SetGroup sets the group of this subtype
func (m *Ec2Netscan) SetGroup(val string) {
	m.groupField = val
}

// ID gets the id of this subtype
func (m *Ec2Netscan) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *Ec2Netscan) SetID(val int32) {
	m.idField = val
}

// IgnoreSystemIPsDuplicates gets the ignore system i ps duplicates of this subtype
func (m *Ec2Netscan) IgnoreSystemIPsDuplicates() bool {
	return m.ignoreSystemIPsDuplicatesField
}

// SetIgnoreSystemIPsDuplicates sets the ignore system i ps duplicates of this subtype
func (m *Ec2Netscan) SetIgnoreSystemIPsDuplicates(val bool) {
	m.ignoreSystemIPsDuplicatesField = val
}

// Method gets the method of this subtype
func (m *Ec2Netscan) Method() string {
	return "nec2"
}

// SetMethod sets the method of this subtype
func (m *Ec2Netscan) SetMethod(val string) {
}

// Name gets the name of this subtype
func (m *Ec2Netscan) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *Ec2Netscan) SetName(val *string) {
	m.nameField = val
}

// NextStart gets the next start of this subtype
func (m *Ec2Netscan) NextStart() string {
	return m.nextStartField
}

// SetNextStart sets the next start of this subtype
func (m *Ec2Netscan) SetNextStart(val string) {
	m.nextStartField = val
}

// NextStartEpoch gets the next start epoch of this subtype
func (m *Ec2Netscan) NextStartEpoch() int64 {
	return m.nextStartEpochField
}

// SetNextStartEpoch sets the next start epoch of this subtype
func (m *Ec2Netscan) SetNextStartEpoch(val int64) {
	m.nextStartEpochField = val
}

// NsgID gets the nsg Id of this subtype
func (m *Ec2Netscan) NsgID() int32 {
	return m.nsgIdField
}

// SetNsgID sets the nsg Id of this subtype
func (m *Ec2Netscan) SetNsgID(val int32) {
	m.nsgIdField = val
}

// Schedule gets the schedule of this subtype
func (m *Ec2Netscan) Schedule() *RestSchedule {
	return m.scheduleField
}

// SetSchedule sets the schedule of this subtype
func (m *Ec2Netscan) SetSchedule(val *RestSchedule) {
	m.scheduleField = val
}

// Version gets the version of this subtype
func (m *Ec2Netscan) Version() int32 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *Ec2Netscan) SetVersion(val int32) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Ec2Netscan) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Which IP the EC2 instance should be monitored with for nec2 scans: private or public
		// Example: private
		// Required: true
		Accessibility *string `json:"accessibility"`

		// The credentials to be used for the scan
		// Required: true
		Credentials *EC2NetscanPolicyCredential `json:"credentials"`

		// ddr
		Ddr *Ec2DDR `json:"ddr,omitempty"`

		// How dead EC2 instances should be handled for nec2 scans. Must be Manually
		// Example: MANUALLY
		DeadOperation string `json:"deadOperation,omitempty"`

		// ports
		Ports *NetscanPorts `json:"ports,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Collector *int32 `json:"collector"`

		CollectorDescription string `json:"collectorDescription,omitempty"`

		CollectorGroup int32 `json:"collectorGroup,omitempty"`

		CollectorGroupName string `json:"collectorGroupName,omitempty"`

		Creator string `json:"creator,omitempty"`

		Description string `json:"description,omitempty"`

		Duplicate *ExcludeDuplicateIps `json:"duplicate"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		IgnoreSystemIPsDuplicates bool `json:"ignoreSystemIPsDuplicates,omitempty"`

		Method string `json:"method"`

		Name *string `json:"name"`

		NextStart string `json:"nextStart,omitempty"`

		NextStartEpoch int64 `json:"nextStartEpoch,omitempty"`

		NsgID int32 `json:"nsgId,omitempty"`

		Schedule *RestSchedule `json:"schedule,omitempty"`

		Version int32 `json:"version,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result Ec2Netscan

	result.collectorField = base.Collector

	result.collectorDescriptionField = base.CollectorDescription

	result.collectorGroupField = base.CollectorGroup

	result.collectorGroupNameField = base.CollectorGroupName

	result.creatorField = base.Creator

	result.descriptionField = base.Description

	result.duplicateField = base.Duplicate

	result.groupField = base.Group

	result.idField = base.ID

	result.ignoreSystemIPsDuplicatesField = base.IgnoreSystemIPsDuplicates

	if base.Method != result.Method() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid method value: %q", base.Method)
	}
	result.nameField = base.Name

	result.nextStartField = base.NextStart

	result.nextStartEpochField = base.NextStartEpoch

	result.nsgIdField = base.NsgID

	result.scheduleField = base.Schedule

	result.versionField = base.Version

	result.Accessibility = data.Accessibility
	result.Credentials = data.Credentials
	result.Ddr = data.Ddr
	result.DeadOperation = data.DeadOperation
	result.Ports = data.Ports

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Ec2Netscan) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Which IP the EC2 instance should be monitored with for nec2 scans: private or public
		// Example: private
		// Required: true
		Accessibility *string `json:"accessibility"`

		// The credentials to be used for the scan
		// Required: true
		Credentials *EC2NetscanPolicyCredential `json:"credentials"`

		// ddr
		Ddr *Ec2DDR `json:"ddr,omitempty"`

		// How dead EC2 instances should be handled for nec2 scans. Must be Manually
		// Example: MANUALLY
		DeadOperation string `json:"deadOperation,omitempty"`

		// ports
		Ports *NetscanPorts `json:"ports,omitempty"`
	}{

		Accessibility: m.Accessibility,

		Credentials: m.Credentials,

		Ddr: m.Ddr,

		DeadOperation: m.DeadOperation,

		Ports: m.Ports,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Collector *int32 `json:"collector"`

		CollectorDescription string `json:"collectorDescription,omitempty"`

		CollectorGroup int32 `json:"collectorGroup,omitempty"`

		CollectorGroupName string `json:"collectorGroupName,omitempty"`

		Creator string `json:"creator,omitempty"`

		Description string `json:"description,omitempty"`

		Duplicate *ExcludeDuplicateIps `json:"duplicate"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		IgnoreSystemIPsDuplicates bool `json:"ignoreSystemIPsDuplicates,omitempty"`

		Method string `json:"method"`

		Name *string `json:"name"`

		NextStart string `json:"nextStart,omitempty"`

		NextStartEpoch int64 `json:"nextStartEpoch,omitempty"`

		NsgID int32 `json:"nsgId,omitempty"`

		Schedule *RestSchedule `json:"schedule,omitempty"`

		Version int32 `json:"version,omitempty"`
	}{

		Collector: m.Collector(),

		CollectorDescription: m.CollectorDescription(),

		CollectorGroup: m.CollectorGroup(),

		CollectorGroupName: m.CollectorGroupName(),

		Creator: m.Creator(),

		Description: m.Description(),

		Duplicate: m.Duplicate(),

		Group: m.Group(),

		ID: m.ID(),

		IgnoreSystemIPsDuplicates: m.IgnoreSystemIPsDuplicates(),

		Method: m.Method(),

		Name: m.Name(),

		NextStart: m.NextStart(),

		NextStartEpoch: m.NextStartEpoch(),

		NsgID: m.NsgID(),

		Schedule: m.Schedule(),

		Version: m.Version(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this ec2 netscan
func (m *Ec2Netscan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuplicate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessibility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDdr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ec2Netscan) validateCollector(formats strfmt.Registry) error {

	if err := validate.Required("collector", "body", m.Collector()); err != nil {
		return err
	}

	return nil
}

func (m *Ec2Netscan) validateDuplicate(formats strfmt.Registry) error {

	if err := validate.Required("duplicate", "body", m.Duplicate()); err != nil {
		return err
	}

	if m.Duplicate() != nil {
		if err := m.Duplicate().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duplicate")
			}
			return err
		}
	}

	return nil
}

func (m *Ec2Netscan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *Ec2Netscan) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule()) { // not required
		return nil
	}

	if m.Schedule() != nil {
		if err := m.Schedule().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *Ec2Netscan) validateAccessibility(formats strfmt.Registry) error {

	if err := validate.Required("accessibility", "body", m.Accessibility); err != nil {
		return err
	}

	return nil
}

func (m *Ec2Netscan) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Ec2Netscan) validateDdr(formats strfmt.Registry) error {

	if swag.IsZero(m.Ddr) { // not required
		return nil
	}

	if m.Ddr != nil {
		if err := m.Ddr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ddr")
			}
			return err
		}
	}

	return nil
}

func (m *Ec2Netscan) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	if m.Ports != nil {
		if err := m.Ports.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ports")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ec2 netscan based on the context it is used
func (m *Ec2Netscan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDuplicate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDdr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ec2Netscan) contextValidateDuplicate(ctx context.Context, formats strfmt.Registry) error {

	if m.Duplicate() != nil {
		if err := m.Duplicate().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duplicate")
			}
			return err
		}
	}

	return nil
}

func (m *Ec2Netscan) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule() != nil {
		if err := m.Schedule().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *Ec2Netscan) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *Ec2Netscan) contextValidateDdr(ctx context.Context, formats strfmt.Registry) error {

	if m.Ddr != nil {
		if err := m.Ddr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ddr")
			}
			return err
		}
	}

	return nil
}

func (m *Ec2Netscan) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	if m.Ports != nil {
		if err := m.Ports.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ports")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Ec2Netscan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ec2Netscan) UnmarshalBinary(b []byte) error {
	var res Ec2Netscan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}