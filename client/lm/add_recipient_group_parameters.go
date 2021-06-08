// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/logicmonitor/lm-sdk-go/models"
	"golang.org/x/net/context"
)

// NewAddRecipientGroupParams creates a new AddRecipientGroupParams object
// with the default values initialized.
func NewAddRecipientGroupParams() *AddRecipientGroupParams {
	var ()
	return &AddRecipientGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddRecipientGroupParamsWithTimeout creates a new AddRecipientGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddRecipientGroupParamsWithTimeout(timeout time.Duration) *AddRecipientGroupParams {
	var ()
	return &AddRecipientGroupParams{

		timeout: timeout,
	}
}

// NewAddRecipientGroupParamsWithContext creates a new AddRecipientGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddRecipientGroupParamsWithContext(ctx context.Context) *AddRecipientGroupParams {
	var ()
	return &AddRecipientGroupParams{

		Context: ctx,
	}
}

// NewAddRecipientGroupParamsWithHTTPClient creates a new AddRecipientGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddRecipientGroupParamsWithHTTPClient(client *http.Client) *AddRecipientGroupParams {
	var ()
	return &AddRecipientGroupParams{
		HTTPClient: client,
	}
}

/*AddRecipientGroupParams contains all the parameters to send to the API endpoint
for the add recipient group operation typically these are written to a http.Request
*/
type AddRecipientGroupParams struct {

	/*Body*/
	Body *models.RecipientGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add recipient group params
func (o *AddRecipientGroupParams) WithTimeout(timeout time.Duration) *AddRecipientGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add recipient group params
func (o *AddRecipientGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add recipient group params
func (o *AddRecipientGroupParams) WithContext(ctx context.Context) *AddRecipientGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add recipient group params
func (o *AddRecipientGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add recipient group params
func (o *AddRecipientGroupParams) WithHTTPClient(client *http.Client) *AddRecipientGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add recipient group params
func (o *AddRecipientGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add recipient group params
func (o *AddRecipientGroupParams) WithBody(body *models.RecipientGroup) *AddRecipientGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add recipient group params
func (o *AddRecipientGroupParams) SetBody(body *models.RecipientGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRecipientGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
