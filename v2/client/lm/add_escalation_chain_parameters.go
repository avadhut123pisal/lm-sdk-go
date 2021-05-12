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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/v2/models"
)

// NewAddEscalationChainParams creates a new AddEscalationChainParams object
// with the default values initialized.
func NewAddEscalationChainParams() *AddEscalationChainParams {
	var ()
	return &AddEscalationChainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddEscalationChainParamsWithTimeout creates a new AddEscalationChainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddEscalationChainParamsWithTimeout(timeout time.Duration) *AddEscalationChainParams {
	var ()
	return &AddEscalationChainParams{

		timeout: timeout,
	}
}

// NewAddEscalationChainParamsWithContext creates a new AddEscalationChainParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddEscalationChainParamsWithContext(ctx context.Context) *AddEscalationChainParams {
	var ()
	return &AddEscalationChainParams{

		Context: ctx,
	}
}

// NewAddEscalationChainParamsWithHTTPClient creates a new AddEscalationChainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddEscalationChainParamsWithHTTPClient(client *http.Client) *AddEscalationChainParams {
	var ()
	return &AddEscalationChainParams{
		HTTPClient: client,
	}
}

/*AddEscalationChainParams contains all the parameters to send to the API endpoint
for the add escalation chain operation typically these are written to a http.Request
*/
type AddEscalationChainParams struct {

	/*Body*/
	Body *models.EscalatingChain

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add escalation chain params
func (o *AddEscalationChainParams) WithTimeout(timeout time.Duration) *AddEscalationChainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add escalation chain params
func (o *AddEscalationChainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add escalation chain params
func (o *AddEscalationChainParams) WithContext(ctx context.Context) *AddEscalationChainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add escalation chain params
func (o *AddEscalationChainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add escalation chain params
func (o *AddEscalationChainParams) WithHTTPClient(client *http.Client) *AddEscalationChainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add escalation chain params
func (o *AddEscalationChainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add escalation chain params
func (o *AddEscalationChainParams) WithBody(body *models.EscalatingChain) *AddEscalationChainParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add escalation chain params
func (o *AddEscalationChainParams) SetBody(body *models.EscalatingChain) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddEscalationChainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
