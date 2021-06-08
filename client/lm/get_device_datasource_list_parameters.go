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
	"github.com/go-openapi/swag"
	"golang.org/x/net/context"
)

// NewGetDeviceDatasourceListParams creates a new GetDeviceDatasourceListParams object
// with the default values initialized.
func NewGetDeviceDatasourceListParams() *GetDeviceDatasourceListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceDatasourceListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceListParamsWithTimeout creates a new GetDeviceDatasourceListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceDatasourceListParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceDatasourceListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetDeviceDatasourceListParamsWithContext creates a new GetDeviceDatasourceListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceDatasourceListParamsWithContext(ctx context.Context) *GetDeviceDatasourceListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceDatasourceListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetDeviceDatasourceListParamsWithHTTPClient creates a new GetDeviceDatasourceListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceDatasourceListParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceDatasourceListParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetDeviceDatasourceListParams contains all the parameters to send to the API endpoint
for the get device datasource list operation typically these are written to a http.Request
*/
type GetDeviceDatasourceListParams struct {

	/*DeviceID*/
	DeviceID int32
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

// WithTimeout adds the timeout to the get device datasource list params
func (o *GetDeviceDatasourceListParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource list params
func (o *GetDeviceDatasourceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource list params
func (o *GetDeviceDatasourceListParams) WithContext(ctx context.Context) *GetDeviceDatasourceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource list params
func (o *GetDeviceDatasourceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource list params
func (o *GetDeviceDatasourceListParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource list params
func (o *GetDeviceDatasourceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device datasource list params
func (o *GetDeviceDatasourceListParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource list params
func (o *GetDeviceDatasourceListParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device datasource list params
func (o *GetDeviceDatasourceListParams) WithFields(fields *string) *GetDeviceDatasourceListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device datasource list params
func (o *GetDeviceDatasourceListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device datasource list params
func (o *GetDeviceDatasourceListParams) WithFilter(filter *string) *GetDeviceDatasourceListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device datasource list params
func (o *GetDeviceDatasourceListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get device datasource list params
func (o *GetDeviceDatasourceListParams) WithOffset(offset *int32) *GetDeviceDatasourceListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device datasource list params
func (o *GetDeviceDatasourceListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device datasource list params
func (o *GetDeviceDatasourceListParams) WithSize(size *int32) *GetDeviceDatasourceListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device datasource list params
func (o *GetDeviceDatasourceListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
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
