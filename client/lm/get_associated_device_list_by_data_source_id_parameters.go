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

// NewGetAssociatedDeviceListByDataSourceIDParams creates a new GetAssociatedDeviceListByDataSourceIDParams object
// with the default values initialized.
func NewGetAssociatedDeviceListByDataSourceIDParams() *GetAssociatedDeviceListByDataSourceIDParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetAssociatedDeviceListByDataSourceIDParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssociatedDeviceListByDataSourceIDParamsWithTimeout creates a new GetAssociatedDeviceListByDataSourceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssociatedDeviceListByDataSourceIDParamsWithTimeout(timeout time.Duration) *GetAssociatedDeviceListByDataSourceIDParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetAssociatedDeviceListByDataSourceIDParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetAssociatedDeviceListByDataSourceIDParamsWithContext creates a new GetAssociatedDeviceListByDataSourceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssociatedDeviceListByDataSourceIDParamsWithContext(ctx context.Context) *GetAssociatedDeviceListByDataSourceIDParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetAssociatedDeviceListByDataSourceIDParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetAssociatedDeviceListByDataSourceIDParamsWithHTTPClient creates a new GetAssociatedDeviceListByDataSourceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssociatedDeviceListByDataSourceIDParamsWithHTTPClient(client *http.Client) *GetAssociatedDeviceListByDataSourceIDParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetAssociatedDeviceListByDataSourceIDParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetAssociatedDeviceListByDataSourceIDParams contains all the parameters to send to the API endpoint
for the get associated device list by data source Id operation typically these are written to a http.Request
*/
type GetAssociatedDeviceListByDataSourceIDParams struct {

	/*Fields*/
	Fields *string
	/*Filter*/
	Filter *string
	/*ID*/
	ID int32
	/*Offset*/
	Offset *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) WithTimeout(timeout time.Duration) *GetAssociatedDeviceListByDataSourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) WithContext(ctx context.Context) *GetAssociatedDeviceListByDataSourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) WithHTTPClient(client *http.Client) *GetAssociatedDeviceListByDataSourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) WithFields(fields *string) *GetAssociatedDeviceListByDataSourceIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) WithFilter(filter *string) *GetAssociatedDeviceListByDataSourceIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) WithID(id int32) *GetAssociatedDeviceListByDataSourceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) SetID(id int32) {
	o.ID = id
}

// WithOffset adds the offset to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) WithOffset(offset *int32) *GetAssociatedDeviceListByDataSourceIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) WithSize(size *int32) *GetAssociatedDeviceListByDataSourceIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get associated device list by data source Id params
func (o *GetAssociatedDeviceListByDataSourceIDParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssociatedDeviceListByDataSourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
