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

// NewPatchOpsNoteByIDParams creates a new PatchOpsNoteByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchOpsNoteByIDParams() *PatchOpsNoteByIDParams {
	return &PatchOpsNoteByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchOpsNoteByIDParamsWithTimeout creates a new PatchOpsNoteByIDParams object
// with the ability to set a timeout on a request.
func NewPatchOpsNoteByIDParamsWithTimeout(timeout time.Duration) *PatchOpsNoteByIDParams {
	return &PatchOpsNoteByIDParams{
		timeout: timeout,
	}
}

// NewPatchOpsNoteByIDParamsWithContext creates a new PatchOpsNoteByIDParams object
// with the ability to set a context for a request.
func NewPatchOpsNoteByIDParamsWithContext(ctx context.Context) *PatchOpsNoteByIDParams {
	return &PatchOpsNoteByIDParams{
		Context: ctx,
	}
}

// NewPatchOpsNoteByIDParamsWithHTTPClient creates a new PatchOpsNoteByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchOpsNoteByIDParamsWithHTTPClient(client *http.Client) *PatchOpsNoteByIDParams {
	return &PatchOpsNoteByIDParams{
		HTTPClient: client,
	}
}

/* PatchOpsNoteByIDParams contains all the parameters to send to the API endpoint
   for the patch ops note by Id operation.

   Typically these are written to a http.Request.
*/
type PatchOpsNoteByIDParams struct {

	// PatchFields.
	PatchFields *string

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1"
	UserAgent *string

	// Body.
	Body *models.OpsNote

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch ops note by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchOpsNoteByIDParams) WithDefaults() *PatchOpsNoteByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch ops note by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchOpsNoteByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1")
	)

	val := PatchOpsNoteByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) WithTimeout(timeout time.Duration) *PatchOpsNoteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) WithContext(ctx context.Context) *PatchOpsNoteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) WithHTTPClient(client *http.Client) *PatchOpsNoteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchFields adds the patchFields to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) WithPatchFields(patchFields *string) *PatchOpsNoteByIDParams {
	o.SetPatchFields(patchFields)
	return o
}

// SetPatchFields adds the patchFields to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) SetPatchFields(patchFields *string) {
	o.PatchFields = patchFields
}

// WithUserAgent adds the userAgent to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) WithUserAgent(userAgent *string) *PatchOpsNoteByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) WithBody(body *models.OpsNote) *PatchOpsNoteByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) SetBody(body *models.OpsNote) {
	o.Body = body
}

// WithID adds the id to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) WithID(id string) *PatchOpsNoteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch ops note by Id params
func (o *PatchOpsNoteByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchOpsNoteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PatchFields != nil {

		// query param PatchFields
		var qrPatchFields string

		if o.PatchFields != nil {
			qrPatchFields = *o.PatchFields
		}
		qPatchFields := qrPatchFields
		if qPatchFields != "" {

			if err := r.SetQueryParam("PatchFields", qPatchFields); err != nil {
				return err
			}
		}
	}

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
