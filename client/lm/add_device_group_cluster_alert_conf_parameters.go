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

// NewAddDeviceGroupClusterAlertConfParams creates a new AddDeviceGroupClusterAlertConfParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDeviceGroupClusterAlertConfParams() *AddDeviceGroupClusterAlertConfParams {
	return &AddDeviceGroupClusterAlertConfParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDeviceGroupClusterAlertConfParamsWithTimeout creates a new AddDeviceGroupClusterAlertConfParams object
// with the ability to set a timeout on a request.
func NewAddDeviceGroupClusterAlertConfParamsWithTimeout(timeout time.Duration) *AddDeviceGroupClusterAlertConfParams {
	return &AddDeviceGroupClusterAlertConfParams{
		timeout: timeout,
	}
}

// NewAddDeviceGroupClusterAlertConfParamsWithContext creates a new AddDeviceGroupClusterAlertConfParams object
// with the ability to set a context for a request.
func NewAddDeviceGroupClusterAlertConfParamsWithContext(ctx context.Context) *AddDeviceGroupClusterAlertConfParams {
	return &AddDeviceGroupClusterAlertConfParams{
		Context: ctx,
	}
}

// NewAddDeviceGroupClusterAlertConfParamsWithHTTPClient creates a new AddDeviceGroupClusterAlertConfParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDeviceGroupClusterAlertConfParamsWithHTTPClient(client *http.Client) *AddDeviceGroupClusterAlertConfParams {
	return &AddDeviceGroupClusterAlertConfParams{
		HTTPClient: client,
	}
}

/* AddDeviceGroupClusterAlertConfParams contains all the parameters to send to the API endpoint
   for the add device group cluster alert conf operation.

   Typically these are written to a http.Request.
*/
type AddDeviceGroupClusterAlertConfParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// Body.
	Body *models.DeviceClusterAlertConfig

	// DeviceGroupID.
	//
	// Format: int32
	DeviceGroupID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add device group cluster alert conf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDeviceGroupClusterAlertConfParams) WithDefaults() *AddDeviceGroupClusterAlertConfParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add device group cluster alert conf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDeviceGroupClusterAlertConfParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := AddDeviceGroupClusterAlertConfParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithTimeout(timeout time.Duration) *AddDeviceGroupClusterAlertConfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithContext(ctx context.Context) *AddDeviceGroupClusterAlertConfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithHTTPClient(client *http.Client) *AddDeviceGroupClusterAlertConfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithUserAgent(userAgent *string) *AddDeviceGroupClusterAlertConfParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithBody(body *models.DeviceClusterAlertConfig) *AddDeviceGroupClusterAlertConfParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetBody(body *models.DeviceClusterAlertConfig) {
	o.Body = body
}

// WithDeviceGroupID adds the deviceGroupID to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithDeviceGroupID(deviceGroupID int32) *AddDeviceGroupClusterAlertConfParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *AddDeviceGroupClusterAlertConfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
