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

// NewGetDashboardGroupByIDParams creates a new GetDashboardGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardGroupByIDParams() *GetDashboardGroupByIDParams {
	return &GetDashboardGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardGroupByIDParamsWithTimeout creates a new GetDashboardGroupByIDParams object
// with the ability to set a timeout on a request.
func NewGetDashboardGroupByIDParamsWithTimeout(timeout time.Duration) *GetDashboardGroupByIDParams {
	return &GetDashboardGroupByIDParams{
		timeout: timeout,
	}
}

// NewGetDashboardGroupByIDParamsWithContext creates a new GetDashboardGroupByIDParams object
// with the ability to set a context for a request.
func NewGetDashboardGroupByIDParamsWithContext(ctx context.Context) *GetDashboardGroupByIDParams {
	return &GetDashboardGroupByIDParams{
		Context: ctx,
	}
}

// NewGetDashboardGroupByIDParamsWithHTTPClient creates a new GetDashboardGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardGroupByIDParamsWithHTTPClient(client *http.Client) *GetDashboardGroupByIDParams {
	return &GetDashboardGroupByIDParams{
		HTTPClient: client,
	}
}

/* GetDashboardGroupByIDParams contains all the parameters to send to the API endpoint
   for the get dashboard group by Id operation.

   Typically these are written to a http.Request.
*/
type GetDashboardGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// Fields.
	Fields *string

	// Format.
	//
	// Default: "json"
	Format *string

	// ID.
	//
	// Format: int32
	ID int32

	// Template.
	Template *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardGroupByIDParams) WithDefaults() *GetDashboardGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")

		formatDefault = string("json")

		templateDefault = bool(false)
	)

	val := GetDashboardGroupByIDParams{
		UserAgent: &userAgentDefault,
		Format:    &formatDefault,
		Template:  &templateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) WithTimeout(timeout time.Duration) *GetDashboardGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) WithContext(ctx context.Context) *GetDashboardGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) WithHTTPClient(client *http.Client) *GetDashboardGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) WithUserAgent(userAgent *string) *GetDashboardGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) WithFields(fields *string) *GetDashboardGroupByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFormat adds the format to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) WithFormat(format *string) *GetDashboardGroupByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) WithID(id int32) *GetDashboardGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WithTemplate adds the template to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) WithTemplate(template *bool) *GetDashboardGroupByIDParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the get dashboard group by Id params
func (o *GetDashboardGroupByIDParams) SetTemplate(template *bool) {
	o.Template = template
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Template != nil {

		// query param template
		var qrTemplate bool

		if o.Template != nil {
			qrTemplate = *o.Template
		}
		qTemplate := swag.FormatBool(qrTemplate)
		if qTemplate != "" {

			if err := r.SetQueryParam("template", qTemplate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
