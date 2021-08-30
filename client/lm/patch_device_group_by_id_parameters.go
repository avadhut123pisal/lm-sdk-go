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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchDeviceGroupByIDParams creates a new PatchDeviceGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDeviceGroupByIDParams() *PatchDeviceGroupByIDParams {
	return &PatchDeviceGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDeviceGroupByIDParamsWithTimeout creates a new PatchDeviceGroupByIDParams object
// with the ability to set a timeout on a request.
func NewPatchDeviceGroupByIDParamsWithTimeout(timeout time.Duration) *PatchDeviceGroupByIDParams {
	return &PatchDeviceGroupByIDParams{
		timeout: timeout,
	}
}

// NewPatchDeviceGroupByIDParamsWithContext creates a new PatchDeviceGroupByIDParams object
// with the ability to set a context for a request.
func NewPatchDeviceGroupByIDParamsWithContext(ctx context.Context) *PatchDeviceGroupByIDParams {
	return &PatchDeviceGroupByIDParams{
		Context: ctx,
	}
}

// NewPatchDeviceGroupByIDParamsWithHTTPClient creates a new PatchDeviceGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDeviceGroupByIDParamsWithHTTPClient(client *http.Client) *PatchDeviceGroupByIDParams {
	return &PatchDeviceGroupByIDParams{
		HTTPClient: client,
	}
}

/* PatchDeviceGroupByIDParams contains all the parameters to send to the API endpoint
   for the patch device group by Id operation.

   Typically these are written to a http.Request.
*/
type PatchDeviceGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// Body.
	Body *models.DeviceGroup

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch device group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceGroupByIDParams) WithDefaults() *PatchDeviceGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch device group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := PatchDeviceGroupByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithTimeout(timeout time.Duration) *PatchDeviceGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithContext(ctx context.Context) *PatchDeviceGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithHTTPClient(client *http.Client) *PatchDeviceGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithUserAgent(userAgent *string) *PatchDeviceGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithBody(body *models.DeviceGroup) *PatchDeviceGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetBody(body *models.DeviceGroup) {
	o.Body = body
}

// WithID adds the id to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithID(id int32) *PatchDeviceGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDeviceGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
