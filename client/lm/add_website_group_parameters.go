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

// NewAddWebsiteGroupParams creates a new AddWebsiteGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddWebsiteGroupParams() *AddWebsiteGroupParams {
	return &AddWebsiteGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddWebsiteGroupParamsWithTimeout creates a new AddWebsiteGroupParams object
// with the ability to set a timeout on a request.
func NewAddWebsiteGroupParamsWithTimeout(timeout time.Duration) *AddWebsiteGroupParams {
	return &AddWebsiteGroupParams{
		timeout: timeout,
	}
}

// NewAddWebsiteGroupParamsWithContext creates a new AddWebsiteGroupParams object
// with the ability to set a context for a request.
func NewAddWebsiteGroupParamsWithContext(ctx context.Context) *AddWebsiteGroupParams {
	return &AddWebsiteGroupParams{
		Context: ctx,
	}
}

// NewAddWebsiteGroupParamsWithHTTPClient creates a new AddWebsiteGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddWebsiteGroupParamsWithHTTPClient(client *http.Client) *AddWebsiteGroupParams {
	return &AddWebsiteGroupParams{
		HTTPClient: client,
	}
}

/* AddWebsiteGroupParams contains all the parameters to send to the API endpoint
   for the add website group operation.

   Typically these are written to a http.Request.
*/
type AddWebsiteGroupParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// Body.
	Body *models.WebsiteGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add website group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddWebsiteGroupParams) WithDefaults() *AddWebsiteGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add website group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddWebsiteGroupParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := AddWebsiteGroupParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add website group params
func (o *AddWebsiteGroupParams) WithTimeout(timeout time.Duration) *AddWebsiteGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add website group params
func (o *AddWebsiteGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add website group params
func (o *AddWebsiteGroupParams) WithContext(ctx context.Context) *AddWebsiteGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add website group params
func (o *AddWebsiteGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add website group params
func (o *AddWebsiteGroupParams) WithHTTPClient(client *http.Client) *AddWebsiteGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add website group params
func (o *AddWebsiteGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add website group params
func (o *AddWebsiteGroupParams) WithUserAgent(userAgent *string) *AddWebsiteGroupParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add website group params
func (o *AddWebsiteGroupParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add website group params
func (o *AddWebsiteGroupParams) WithBody(body *models.WebsiteGroup) *AddWebsiteGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add website group params
func (o *AddWebsiteGroupParams) SetBody(body *models.WebsiteGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddWebsiteGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
