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

// NewAddReportGroupParams creates a new AddReportGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddReportGroupParams() *AddReportGroupParams {
	return &AddReportGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddReportGroupParamsWithTimeout creates a new AddReportGroupParams object
// with the ability to set a timeout on a request.
func NewAddReportGroupParamsWithTimeout(timeout time.Duration) *AddReportGroupParams {
	return &AddReportGroupParams{
		timeout: timeout,
	}
}

// NewAddReportGroupParamsWithContext creates a new AddReportGroupParams object
// with the ability to set a context for a request.
func NewAddReportGroupParamsWithContext(ctx context.Context) *AddReportGroupParams {
	return &AddReportGroupParams{
		Context: ctx,
	}
}

// NewAddReportGroupParamsWithHTTPClient creates a new AddReportGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddReportGroupParamsWithHTTPClient(client *http.Client) *AddReportGroupParams {
	return &AddReportGroupParams{
		HTTPClient: client,
	}
}

/* AddReportGroupParams contains all the parameters to send to the API endpoint
   for the add report group operation.

   Typically these are written to a http.Request.
*/
type AddReportGroupParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// Body.
	Body *models.ReportGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add report group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddReportGroupParams) WithDefaults() *AddReportGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add report group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddReportGroupParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := AddReportGroupParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add report group params
func (o *AddReportGroupParams) WithTimeout(timeout time.Duration) *AddReportGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add report group params
func (o *AddReportGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add report group params
func (o *AddReportGroupParams) WithContext(ctx context.Context) *AddReportGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add report group params
func (o *AddReportGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add report group params
func (o *AddReportGroupParams) WithHTTPClient(client *http.Client) *AddReportGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add report group params
func (o *AddReportGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add report group params
func (o *AddReportGroupParams) WithUserAgent(userAgent *string) *AddReportGroupParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add report group params
func (o *AddReportGroupParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add report group params
func (o *AddReportGroupParams) WithBody(body *models.ReportGroup) *AddReportGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add report group params
func (o *AddReportGroupParams) SetBody(body *models.ReportGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddReportGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
