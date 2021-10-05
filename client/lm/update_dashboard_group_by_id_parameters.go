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

// NewUpdateDashboardGroupByIDParams creates a new UpdateDashboardGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDashboardGroupByIDParams() *UpdateDashboardGroupByIDParams {
	return &UpdateDashboardGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDashboardGroupByIDParamsWithTimeout creates a new UpdateDashboardGroupByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDashboardGroupByIDParamsWithTimeout(timeout time.Duration) *UpdateDashboardGroupByIDParams {
	return &UpdateDashboardGroupByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDashboardGroupByIDParamsWithContext creates a new UpdateDashboardGroupByIDParams object
// with the ability to set a context for a request.
func NewUpdateDashboardGroupByIDParamsWithContext(ctx context.Context) *UpdateDashboardGroupByIDParams {
	return &UpdateDashboardGroupByIDParams{
		Context: ctx,
	}
}

// NewUpdateDashboardGroupByIDParamsWithHTTPClient creates a new UpdateDashboardGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDashboardGroupByIDParamsWithHTTPClient(client *http.Client) *UpdateDashboardGroupByIDParams {
	return &UpdateDashboardGroupByIDParams{
		HTTPClient: client,
	}
}

/* UpdateDashboardGroupByIDParams contains all the parameters to send to the API endpoint
   for the update dashboard group by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateDashboardGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1"
	UserAgent *string

	// Body.
	Body *models.DashboardGroup

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update dashboard group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardGroupByIDParams) WithDefaults() *UpdateDashboardGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update dashboard group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1")
	)

	val := UpdateDashboardGroupByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) WithTimeout(timeout time.Duration) *UpdateDashboardGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) WithContext(ctx context.Context) *UpdateDashboardGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) WithHTTPClient(client *http.Client) *UpdateDashboardGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) WithUserAgent(userAgent *string) *UpdateDashboardGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) WithBody(body *models.DashboardGroup) *UpdateDashboardGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) SetBody(body *models.DashboardGroup) {
	o.Body = body
}

// WithID adds the id to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) WithID(id int32) *UpdateDashboardGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update dashboard group by Id params
func (o *UpdateDashboardGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDashboardGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
