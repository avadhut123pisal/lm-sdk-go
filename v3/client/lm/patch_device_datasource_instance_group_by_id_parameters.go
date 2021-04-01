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

// NewPatchDeviceDatasourceInstanceGroupByIDParams creates a new PatchDeviceDatasourceInstanceGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDeviceDatasourceInstanceGroupByIDParams() *PatchDeviceDatasourceInstanceGroupByIDParams {
	return &PatchDeviceDatasourceInstanceGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDeviceDatasourceInstanceGroupByIDParamsWithTimeout creates a new PatchDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a timeout on a request.
func NewPatchDeviceDatasourceInstanceGroupByIDParamsWithTimeout(timeout time.Duration) *PatchDeviceDatasourceInstanceGroupByIDParams {
	return &PatchDeviceDatasourceInstanceGroupByIDParams{
		timeout: timeout,
	}
}

// NewPatchDeviceDatasourceInstanceGroupByIDParamsWithContext creates a new PatchDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a context for a request.
func NewPatchDeviceDatasourceInstanceGroupByIDParamsWithContext(ctx context.Context) *PatchDeviceDatasourceInstanceGroupByIDParams {
	return &PatchDeviceDatasourceInstanceGroupByIDParams{
		Context: ctx,
	}
}

// NewPatchDeviceDatasourceInstanceGroupByIDParamsWithHTTPClient creates a new PatchDeviceDatasourceInstanceGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDeviceDatasourceInstanceGroupByIDParamsWithHTTPClient(client *http.Client) *PatchDeviceDatasourceInstanceGroupByIDParams {
	return &PatchDeviceDatasourceInstanceGroupByIDParams{
		HTTPClient: client,
	}
}

/* PatchDeviceDatasourceInstanceGroupByIDParams contains all the parameters to send to the API endpoint
   for the patch device datasource instance group by Id operation.

   Typically these are written to a http.Request.
*/
type PatchDeviceDatasourceInstanceGroupByIDParams struct {

	// Body.
	Body *models.DeviceDataSourceInstanceGroup

	/* DeviceDsID.

	   The device-datasource ID you'd like to add an instance group for

	   Format: int32
	*/
	DeviceDsID int32

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch device datasource instance group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WithDefaults() *PatchDeviceDatasourceInstanceGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch device datasource instance group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WithTimeout(timeout time.Duration) *PatchDeviceDatasourceInstanceGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WithContext(ctx context.Context) *PatchDeviceDatasourceInstanceGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WithHTTPClient(client *http.Client) *PatchDeviceDatasourceInstanceGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WithBody(body *models.DeviceDataSourceInstanceGroup) *PatchDeviceDatasourceInstanceGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) SetBody(body *models.DeviceDataSourceInstanceGroup) {
	o.Body = body
}

// WithDeviceDsID adds the deviceDsID to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WithDeviceDsID(deviceDsID int32) *PatchDeviceDatasourceInstanceGroupByIDParams {
	o.SetDeviceDsID(deviceDsID)
	return o
}

// SetDeviceDsID adds the deviceDsId to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) SetDeviceDsID(deviceDsID int32) {
	o.DeviceDsID = deviceDsID
}

// WithDeviceID adds the deviceID to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WithDeviceID(deviceID int32) *PatchDeviceDatasourceInstanceGroupByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithID adds the id to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WithID(id int32) *PatchDeviceDatasourceInstanceGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch device datasource instance group by Id params
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDeviceDatasourceInstanceGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceDsId
	if err := r.SetPathParam("deviceDsId", swag.FormatInt32(o.DeviceDsID)); err != nil {
		return err
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
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