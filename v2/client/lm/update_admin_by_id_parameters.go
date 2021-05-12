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

// NewUpdateAdminByIDParams creates a new UpdateAdminByIDParams object
// with the default values initialized.
func NewUpdateAdminByIDParams() *UpdateAdminByIDParams {
	var (
		changePasswordDefault = bool(false)
	)
	return &UpdateAdminByIDParams{
		ChangePassword: &changePasswordDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAdminByIDParamsWithTimeout creates a new UpdateAdminByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAdminByIDParamsWithTimeout(timeout time.Duration) *UpdateAdminByIDParams {
	var (
		changePasswordDefault = bool(false)
	)
	return &UpdateAdminByIDParams{
		ChangePassword: &changePasswordDefault,

		timeout: timeout,
	}
}

// NewUpdateAdminByIDParamsWithContext creates a new UpdateAdminByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAdminByIDParamsWithContext(ctx context.Context) *UpdateAdminByIDParams {
	var (
		changePasswordDefault = bool(false)
	)
	return &UpdateAdminByIDParams{
		ChangePassword: &changePasswordDefault,

		Context: ctx,
	}
}

// NewUpdateAdminByIDParamsWithHTTPClient creates a new UpdateAdminByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAdminByIDParamsWithHTTPClient(client *http.Client) *UpdateAdminByIDParams {
	var (
		changePasswordDefault = bool(false)
	)
	return &UpdateAdminByIDParams{
		ChangePassword: &changePasswordDefault,
		HTTPClient:     client,
	}
}

/*UpdateAdminByIDParams contains all the parameters to send to the API endpoint
for the update admin by Id operation typically these are written to a http.Request
*/
type UpdateAdminByIDParams struct {

	/*Body*/
	Body *models.Admin
	/*ChangePassword*/
	ChangePassword *bool
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update admin by Id params
func (o *UpdateAdminByIDParams) WithTimeout(timeout time.Duration) *UpdateAdminByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update admin by Id params
func (o *UpdateAdminByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update admin by Id params
func (o *UpdateAdminByIDParams) WithContext(ctx context.Context) *UpdateAdminByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update admin by Id params
func (o *UpdateAdminByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update admin by Id params
func (o *UpdateAdminByIDParams) WithHTTPClient(client *http.Client) *UpdateAdminByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update admin by Id params
func (o *UpdateAdminByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update admin by Id params
func (o *UpdateAdminByIDParams) WithBody(body *models.Admin) *UpdateAdminByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update admin by Id params
func (o *UpdateAdminByIDParams) SetBody(body *models.Admin) {
	o.Body = body
}

// WithChangePassword adds the changePassword to the update admin by Id params
func (o *UpdateAdminByIDParams) WithChangePassword(changePassword *bool) *UpdateAdminByIDParams {
	o.SetChangePassword(changePassword)
	return o
}

// SetChangePassword adds the changePassword to the update admin by Id params
func (o *UpdateAdminByIDParams) SetChangePassword(changePassword *bool) {
	o.ChangePassword = changePassword
}

// WithID adds the id to the update admin by Id params
func (o *UpdateAdminByIDParams) WithID(id int32) *UpdateAdminByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update admin by Id params
func (o *UpdateAdminByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAdminByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ChangePassword != nil {

		// query param changePassword
		var qrChangePassword bool
		if o.ChangePassword != nil {
			qrChangePassword = *o.ChangePassword
		}
		qChangePassword := swag.FormatBool(qrChangePassword)
		if qChangePassword != "" {
			if err := r.SetQueryParam("changePassword", qChangePassword); err != nil {
				return err
			}
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
