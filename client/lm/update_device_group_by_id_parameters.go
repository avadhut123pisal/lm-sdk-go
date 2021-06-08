// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	models "github.com/logicmonitor/lm-sdk-go/models"
	"golang.org/x/net/context"
)

// NewUpdateDeviceGroupByIDParams creates a new UpdateDeviceGroupByIDParams object
// with the default values initialized.
func NewUpdateDeviceGroupByIDParams() *UpdateDeviceGroupByIDParams {
	var ()
	return &UpdateDeviceGroupByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceGroupByIDParamsWithTimeout creates a new UpdateDeviceGroupByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeviceGroupByIDParamsWithTimeout(timeout time.Duration) *UpdateDeviceGroupByIDParams {
	var ()
	return &UpdateDeviceGroupByIDParams{

		timeout: timeout,
	}
}

// NewUpdateDeviceGroupByIDParamsWithContext creates a new UpdateDeviceGroupByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeviceGroupByIDParamsWithContext(ctx context.Context) *UpdateDeviceGroupByIDParams {
	var ()
	return &UpdateDeviceGroupByIDParams{

		Context: ctx,
	}
}

// NewUpdateDeviceGroupByIDParamsWithHTTPClient creates a new UpdateDeviceGroupByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeviceGroupByIDParamsWithHTTPClient(client *http.Client) *UpdateDeviceGroupByIDParams {
	var ()
	return &UpdateDeviceGroupByIDParams{
		HTTPClient: client,
	}
}

/*UpdateDeviceGroupByIDParams contains all the parameters to send to the API endpoint
for the update device group by Id operation typically these are written to a http.Request
*/
type UpdateDeviceGroupByIDParams struct {

	/*Body*/
	Body *models.DeviceGroup
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithTimeout(timeout time.Duration) *UpdateDeviceGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithContext(ctx context.Context) *UpdateDeviceGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithHTTPClient(client *http.Client) *UpdateDeviceGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithBody(body *models.DeviceGroup) *UpdateDeviceGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetBody(body *models.DeviceGroup) {
	o.Body = body
}

// WithID adds the id to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithID(id int32) *UpdateDeviceGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
