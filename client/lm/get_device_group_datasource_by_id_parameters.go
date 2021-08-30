// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetDeviceGroupDatasourceByIDParams creates a new GetDeviceGroupDatasourceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceGroupDatasourceByIDParams() *GetDeviceGroupDatasourceByIDParams {
	return &GetDeviceGroupDatasourceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupDatasourceByIDParamsWithTimeout creates a new GetDeviceGroupDatasourceByIDParams object
// with the ability to set a timeout on a request.
func NewGetDeviceGroupDatasourceByIDParamsWithTimeout(timeout time.Duration) *GetDeviceGroupDatasourceByIDParams {
	return &GetDeviceGroupDatasourceByIDParams{
		timeout: timeout,
	}
}

// NewGetDeviceGroupDatasourceByIDParamsWithContext creates a new GetDeviceGroupDatasourceByIDParams object
// with the ability to set a context for a request.
func NewGetDeviceGroupDatasourceByIDParamsWithContext(ctx context.Context) *GetDeviceGroupDatasourceByIDParams {
	return &GetDeviceGroupDatasourceByIDParams{
		Context: ctx,
	}
}

// NewGetDeviceGroupDatasourceByIDParamsWithHTTPClient creates a new GetDeviceGroupDatasourceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceGroupDatasourceByIDParamsWithHTTPClient(client *http.Client) *GetDeviceGroupDatasourceByIDParams {
	return &GetDeviceGroupDatasourceByIDParams{
		HTTPClient: client,
	}
}

/* GetDeviceGroupDatasourceByIDParams contains all the parameters to send to the API endpoint
   for the get device group datasource by Id operation.

   Typically these are written to a http.Request.
*/
type GetDeviceGroupDatasourceByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// DeviceGroupID.
	//
	// Format: int32
	DeviceGroupID int32

	// Fields.
	Fields *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device group datasource by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupDatasourceByIDParams) WithDefaults() *GetDeviceGroupDatasourceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device group datasource by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupDatasourceByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := GetDeviceGroupDatasourceByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) WithTimeout(timeout time.Duration) *GetDeviceGroupDatasourceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) WithContext(ctx context.Context) *GetDeviceGroupDatasourceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) WithHTTPClient(client *http.Client) *GetDeviceGroupDatasourceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) WithUserAgent(userAgent *string) *GetDeviceGroupDatasourceByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceGroupID adds the deviceGroupID to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) WithDeviceGroupID(deviceGroupID int32) *GetDeviceGroupDatasourceByIDParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithFields adds the fields to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) WithFields(fields *string) *GetDeviceGroupDatasourceByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) WithID(id int32) *GetDeviceGroupDatasourceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device group datasource by Id params
func (o *GetDeviceGroupDatasourceByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupDatasourceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserAgent != nil {

		// header param User-Agent
		if err := r.SetHeaderParam("User-Agent", *o.UserAgent); err != nil {
			return err
		}
	}

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
