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

// NewPatchAlertRuleByIDParams creates a new PatchAlertRuleByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAlertRuleByIDParams() *PatchAlertRuleByIDParams {
	return &PatchAlertRuleByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAlertRuleByIDParamsWithTimeout creates a new PatchAlertRuleByIDParams object
// with the ability to set a timeout on a request.
func NewPatchAlertRuleByIDParamsWithTimeout(timeout time.Duration) *PatchAlertRuleByIDParams {
	return &PatchAlertRuleByIDParams{
		timeout: timeout,
	}
}

// NewPatchAlertRuleByIDParamsWithContext creates a new PatchAlertRuleByIDParams object
// with the ability to set a context for a request.
func NewPatchAlertRuleByIDParamsWithContext(ctx context.Context) *PatchAlertRuleByIDParams {
	return &PatchAlertRuleByIDParams{
		Context: ctx,
	}
}

// NewPatchAlertRuleByIDParamsWithHTTPClient creates a new PatchAlertRuleByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAlertRuleByIDParamsWithHTTPClient(client *http.Client) *PatchAlertRuleByIDParams {
	return &PatchAlertRuleByIDParams{
		HTTPClient: client,
	}
}

/* PatchAlertRuleByIDParams contains all the parameters to send to the API endpoint
   for the patch alert rule by Id operation.

   Typically these are written to a http.Request.
*/
type PatchAlertRuleByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// Body.
	Body *models.AlertRule

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch alert rule by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAlertRuleByIDParams) WithDefaults() *PatchAlertRuleByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch alert rule by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAlertRuleByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := PatchAlertRuleByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) WithTimeout(timeout time.Duration) *PatchAlertRuleByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) WithContext(ctx context.Context) *PatchAlertRuleByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) WithHTTPClient(client *http.Client) *PatchAlertRuleByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) WithUserAgent(userAgent *string) *PatchAlertRuleByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) WithBody(body *models.AlertRule) *PatchAlertRuleByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) SetBody(body *models.AlertRule) {
	o.Body = body
}

// WithID adds the id to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) WithID(id int32) *PatchAlertRuleByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch alert rule by Id params
func (o *PatchAlertRuleByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAlertRuleByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
