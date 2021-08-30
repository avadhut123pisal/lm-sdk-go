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

// NewGetAPITokenListByAdminIDParams creates a new GetAPITokenListByAdminIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPITokenListByAdminIDParams() *GetAPITokenListByAdminIDParams {
	return &GetAPITokenListByAdminIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPITokenListByAdminIDParamsWithTimeout creates a new GetAPITokenListByAdminIDParams object
// with the ability to set a timeout on a request.
func NewGetAPITokenListByAdminIDParamsWithTimeout(timeout time.Duration) *GetAPITokenListByAdminIDParams {
	return &GetAPITokenListByAdminIDParams{
		timeout: timeout,
	}
}

// NewGetAPITokenListByAdminIDParamsWithContext creates a new GetAPITokenListByAdminIDParams object
// with the ability to set a context for a request.
func NewGetAPITokenListByAdminIDParamsWithContext(ctx context.Context) *GetAPITokenListByAdminIDParams {
	return &GetAPITokenListByAdminIDParams{
		Context: ctx,
	}
}

// NewGetAPITokenListByAdminIDParamsWithHTTPClient creates a new GetAPITokenListByAdminIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPITokenListByAdminIDParamsWithHTTPClient(client *http.Client) *GetAPITokenListByAdminIDParams {
	return &GetAPITokenListByAdminIDParams{
		HTTPClient: client,
	}
}

/* GetAPITokenListByAdminIDParams contains all the parameters to send to the API endpoint
   for the get Api token list by admin Id operation.

   Typically these are written to a http.Request.
*/
type GetAPITokenListByAdminIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// AdminID.
	//
	// Format: int32
	AdminID int32

	// Fields.
	Fields *string

	// Filter.
	Filter *string

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

// WithDefaults hydrates default values in the get Api token list by admin Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITokenListByAdminIDParams) WithDefaults() *GetAPITokenListByAdminIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Api token list by admin Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITokenListByAdminIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetAPITokenListByAdminIDParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithTimeout(timeout time.Duration) *GetAPITokenListByAdminIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithContext(ctx context.Context) *GetAPITokenListByAdminIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithHTTPClient(client *http.Client) *GetAPITokenListByAdminIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithUserAgent(userAgent *string) *GetAPITokenListByAdminIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithAdminID adds the adminID to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithAdminID(adminID int32) *GetAPITokenListByAdminIDParams {
	o.SetAdminID(adminID)
	return o
}

// SetAdminID adds the adminId to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetAdminID(adminID int32) {
	o.AdminID = adminID
}

// WithFields adds the fields to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithFields(fields *string) *GetAPITokenListByAdminIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithFilter(filter *string) *GetAPITokenListByAdminIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithOffset(offset *int32) *GetAPITokenListByAdminIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) WithSize(size *int32) *GetAPITokenListByAdminIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get Api token list by admin Id params
func (o *GetAPITokenListByAdminIDParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPITokenListByAdminIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param adminId
	if err := r.SetPathParam("adminId", swag.FormatInt32(o.AdminID)); err != nil {
		return err
	}

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
