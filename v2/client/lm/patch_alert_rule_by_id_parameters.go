// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/v2/models"
)

// NewPatchAlertRuleByIDParams creates a new PatchAlertRuleByIDParams object
// with the default values initialized.
func NewPatchAlertRuleByIDParams() *PatchAlertRuleByIDParams {
	var ()
	return &PatchAlertRuleByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAlertRuleByIDParamsWithTimeout creates a new PatchAlertRuleByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAlertRuleByIDParamsWithTimeout(timeout time.Duration) *PatchAlertRuleByIDParams {
	var ()
	return &PatchAlertRuleByIDParams{

		timeout: timeout,
	}
}

// NewPatchAlertRuleByIDParamsWithContext creates a new PatchAlertRuleByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAlertRuleByIDParamsWithContext(ctx context.Context) *PatchAlertRuleByIDParams {
	var ()
	return &PatchAlertRuleByIDParams{

		Context: ctx,
	}
}

// NewPatchAlertRuleByIDParamsWithHTTPClient creates a new PatchAlertRuleByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAlertRuleByIDParamsWithHTTPClient(client *http.Client) *PatchAlertRuleByIDParams {
	var ()
	return &PatchAlertRuleByIDParams{
		HTTPClient: client,
	}
}

/*PatchAlertRuleByIDParams contains all the parameters to send to the API endpoint
for the patch alert rule by Id operation typically these are written to a http.Request
*/
type PatchAlertRuleByIDParams struct {

	/*Body*/
	Body *models.AlertRule
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
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
