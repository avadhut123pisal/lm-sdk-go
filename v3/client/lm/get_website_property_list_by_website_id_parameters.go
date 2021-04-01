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

// NewGetWebsitePropertyListByWebsiteIDParams creates a new GetWebsitePropertyListByWebsiteIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsitePropertyListByWebsiteIDParams() *GetWebsitePropertyListByWebsiteIDParams {
	return &GetWebsitePropertyListByWebsiteIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsitePropertyListByWebsiteIDParamsWithTimeout creates a new GetWebsitePropertyListByWebsiteIDParams object
// with the ability to set a timeout on a request.
func NewGetWebsitePropertyListByWebsiteIDParamsWithTimeout(timeout time.Duration) *GetWebsitePropertyListByWebsiteIDParams {
	return &GetWebsitePropertyListByWebsiteIDParams{
		timeout: timeout,
	}
}

// NewGetWebsitePropertyListByWebsiteIDParamsWithContext creates a new GetWebsitePropertyListByWebsiteIDParams object
// with the ability to set a context for a request.
func NewGetWebsitePropertyListByWebsiteIDParamsWithContext(ctx context.Context) *GetWebsitePropertyListByWebsiteIDParams {
	return &GetWebsitePropertyListByWebsiteIDParams{
		Context: ctx,
	}
}

// NewGetWebsitePropertyListByWebsiteIDParamsWithHTTPClient creates a new GetWebsitePropertyListByWebsiteIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsitePropertyListByWebsiteIDParamsWithHTTPClient(client *http.Client) *GetWebsitePropertyListByWebsiteIDParams {
	return &GetWebsitePropertyListByWebsiteIDParams{
		HTTPClient: client,
	}
}

/* GetWebsitePropertyListByWebsiteIDParams contains all the parameters to send to the API endpoint
   for the get website property list by website Id operation.

   Typically these are written to a http.Request.
*/
type GetWebsitePropertyListByWebsiteIDParams struct {

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get website property list by website Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsitePropertyListByWebsiteIDParams) WithDefaults() *GetWebsitePropertyListByWebsiteIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get website property list by website Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsitePropertyListByWebsiteIDParams) SetDefaults() {
	var (
		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetWebsitePropertyListByWebsiteIDParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) WithTimeout(timeout time.Duration) *GetWebsitePropertyListByWebsiteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) WithContext(ctx context.Context) *GetWebsitePropertyListByWebsiteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) WithHTTPClient(client *http.Client) *GetWebsitePropertyListByWebsiteIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) WithFields(fields *string) *GetWebsitePropertyListByWebsiteIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) WithFilter(filter *string) *GetWebsitePropertyListByWebsiteIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) WithID(id int32) *GetWebsitePropertyListByWebsiteIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) SetID(id int32) {
	o.ID = id
}

// WithOffset adds the offset to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) WithOffset(offset *int32) *GetWebsitePropertyListByWebsiteIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) WithSize(size *int32) *GetWebsitePropertyListByWebsiteIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get website property list by website Id params
func (o *GetWebsitePropertyListByWebsiteIDParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsitePropertyListByWebsiteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
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