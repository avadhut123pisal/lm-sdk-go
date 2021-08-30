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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewAddDashboardParams creates a new AddDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDashboardParams() *AddDashboardParams {
	return &AddDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDashboardParamsWithTimeout creates a new AddDashboardParams object
// with the ability to set a timeout on a request.
func NewAddDashboardParamsWithTimeout(timeout time.Duration) *AddDashboardParams {
	return &AddDashboardParams{
		timeout: timeout,
	}
}

// NewAddDashboardParamsWithContext creates a new AddDashboardParams object
// with the ability to set a context for a request.
func NewAddDashboardParamsWithContext(ctx context.Context) *AddDashboardParams {
	return &AddDashboardParams{
		Context: ctx,
	}
}

// NewAddDashboardParamsWithHTTPClient creates a new AddDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDashboardParamsWithHTTPClient(client *http.Client) *AddDashboardParams {
	return &AddDashboardParams{
		HTTPClient: client,
	}
}

/* AddDashboardParams contains all the parameters to send to the API endpoint
   for the add dashboard operation.

   Typically these are written to a http.Request.
*/
type AddDashboardParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// Body.
	Body *models.Dashboard

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDashboardParams) WithDefaults() *AddDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDashboardParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := AddDashboardParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add dashboard params
func (o *AddDashboardParams) WithTimeout(timeout time.Duration) *AddDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add dashboard params
func (o *AddDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add dashboard params
func (o *AddDashboardParams) WithContext(ctx context.Context) *AddDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add dashboard params
func (o *AddDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add dashboard params
func (o *AddDashboardParams) WithHTTPClient(client *http.Client) *AddDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add dashboard params
func (o *AddDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add dashboard params
func (o *AddDashboardParams) WithUserAgent(userAgent *string) *AddDashboardParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add dashboard params
func (o *AddDashboardParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add dashboard params
func (o *AddDashboardParams) WithBody(body *models.Dashboard) *AddDashboardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add dashboard params
func (o *AddDashboardParams) SetBody(body *models.Dashboard) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
