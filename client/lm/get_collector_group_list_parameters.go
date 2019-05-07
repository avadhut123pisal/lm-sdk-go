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
)

// NewGetCollectorGroupListParams creates a new GetCollectorGroupListParams object
// with the default values initialized.
func NewGetCollectorGroupListParams() *GetCollectorGroupListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetCollectorGroupListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCollectorGroupListParamsWithTimeout creates a new GetCollectorGroupListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCollectorGroupListParamsWithTimeout(timeout time.Duration) *GetCollectorGroupListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetCollectorGroupListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetCollectorGroupListParamsWithContext creates a new GetCollectorGroupListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCollectorGroupListParamsWithContext(ctx context.Context) *GetCollectorGroupListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetCollectorGroupListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetCollectorGroupListParamsWithHTTPClient creates a new GetCollectorGroupListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCollectorGroupListParamsWithHTTPClient(client *http.Client) *GetCollectorGroupListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetCollectorGroupListParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetCollectorGroupListParams contains all the parameters to send to the API endpoint
for the get collector group list operation typically these are written to a http.Request
*/
type GetCollectorGroupListParams struct {

	/*Fields*/
	Fields *string
	/*Filter*/
	Filter *string
	/*Offset*/
	Offset *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get collector group list params
func (o *GetCollectorGroupListParams) WithTimeout(timeout time.Duration) *GetCollectorGroupListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get collector group list params
func (o *GetCollectorGroupListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get collector group list params
func (o *GetCollectorGroupListParams) WithContext(ctx context.Context) *GetCollectorGroupListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get collector group list params
func (o *GetCollectorGroupListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get collector group list params
func (o *GetCollectorGroupListParams) WithHTTPClient(client *http.Client) *GetCollectorGroupListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get collector group list params
func (o *GetCollectorGroupListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get collector group list params
func (o *GetCollectorGroupListParams) WithFields(fields *string) *GetCollectorGroupListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get collector group list params
func (o *GetCollectorGroupListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get collector group list params
func (o *GetCollectorGroupListParams) WithFilter(filter *string) *GetCollectorGroupListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get collector group list params
func (o *GetCollectorGroupListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get collector group list params
func (o *GetCollectorGroupListParams) WithOffset(offset *int32) *GetCollectorGroupListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get collector group list params
func (o *GetCollectorGroupListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get collector group list params
func (o *GetCollectorGroupListParams) WithSize(size *int32) *GetCollectorGroupListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get collector group list params
func (o *GetCollectorGroupListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetCollectorGroupListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}