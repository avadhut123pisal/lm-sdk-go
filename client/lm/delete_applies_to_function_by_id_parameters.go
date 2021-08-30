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

// NewDeleteAppliesToFunctionByIDParams creates a new DeleteAppliesToFunctionByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAppliesToFunctionByIDParams() *DeleteAppliesToFunctionByIDParams {
	return &DeleteAppliesToFunctionByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppliesToFunctionByIDParamsWithTimeout creates a new DeleteAppliesToFunctionByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAppliesToFunctionByIDParamsWithTimeout(timeout time.Duration) *DeleteAppliesToFunctionByIDParams {
	return &DeleteAppliesToFunctionByIDParams{
		timeout: timeout,
	}
}

// NewDeleteAppliesToFunctionByIDParamsWithContext creates a new DeleteAppliesToFunctionByIDParams object
// with the ability to set a context for a request.
func NewDeleteAppliesToFunctionByIDParamsWithContext(ctx context.Context) *DeleteAppliesToFunctionByIDParams {
	return &DeleteAppliesToFunctionByIDParams{
		Context: ctx,
	}
}

// NewDeleteAppliesToFunctionByIDParamsWithHTTPClient creates a new DeleteAppliesToFunctionByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAppliesToFunctionByIDParamsWithHTTPClient(client *http.Client) *DeleteAppliesToFunctionByIDParams {
	return &DeleteAppliesToFunctionByIDParams{
		HTTPClient: client,
	}
}

/* DeleteAppliesToFunctionByIDParams contains all the parameters to send to the API endpoint
   for the delete applies to function by Id operation.

   Typically these are written to a http.Request.
*/
type DeleteAppliesToFunctionByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// ID.
	//
	// Format: int32
	ID int32

	// IgnoreReference.
	IgnoreReference *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete applies to function by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAppliesToFunctionByIDParams) WithDefaults() *DeleteAppliesToFunctionByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete applies to function by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAppliesToFunctionByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")

		ignoreReferenceDefault = bool(false)
	)

	val := DeleteAppliesToFunctionByIDParams{
		UserAgent:       &userAgentDefault,
		IgnoreReference: &ignoreReferenceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) WithTimeout(timeout time.Duration) *DeleteAppliesToFunctionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) WithContext(ctx context.Context) *DeleteAppliesToFunctionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) WithHTTPClient(client *http.Client) *DeleteAppliesToFunctionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) WithUserAgent(userAgent *string) *DeleteAppliesToFunctionByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) WithID(id int32) *DeleteAppliesToFunctionByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) SetID(id int32) {
	o.ID = id
}

// WithIgnoreReference adds the ignoreReference to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) WithIgnoreReference(ignoreReference *bool) *DeleteAppliesToFunctionByIDParams {
	o.SetIgnoreReference(ignoreReference)
	return o
}

// SetIgnoreReference adds the ignoreReference to the delete applies to function by Id params
func (o *DeleteAppliesToFunctionByIDParams) SetIgnoreReference(ignoreReference *bool) {
	o.IgnoreReference = ignoreReference
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppliesToFunctionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.IgnoreReference != nil {

		// query param ignoreReference
		var qrIgnoreReference bool

		if o.IgnoreReference != nil {
			qrIgnoreReference = *o.IgnoreReference
		}
		qIgnoreReference := swag.FormatBool(qrIgnoreReference)
		if qIgnoreReference != "" {

			if err := r.SetQueryParam("ignoreReference", qIgnoreReference); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
