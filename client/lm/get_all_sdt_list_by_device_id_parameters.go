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

// NewGetAllSDTListByDeviceIDParams creates a new GetAllSDTListByDeviceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllSDTListByDeviceIDParams() *GetAllSDTListByDeviceIDParams {
	return &GetAllSDTListByDeviceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllSDTListByDeviceIDParamsWithTimeout creates a new GetAllSDTListByDeviceIDParams object
// with the ability to set a timeout on a request.
func NewGetAllSDTListByDeviceIDParamsWithTimeout(timeout time.Duration) *GetAllSDTListByDeviceIDParams {
	return &GetAllSDTListByDeviceIDParams{
		timeout: timeout,
	}
}

// NewGetAllSDTListByDeviceIDParamsWithContext creates a new GetAllSDTListByDeviceIDParams object
// with the ability to set a context for a request.
func NewGetAllSDTListByDeviceIDParamsWithContext(ctx context.Context) *GetAllSDTListByDeviceIDParams {
	return &GetAllSDTListByDeviceIDParams{
		Context: ctx,
	}
}

// NewGetAllSDTListByDeviceIDParamsWithHTTPClient creates a new GetAllSDTListByDeviceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllSDTListByDeviceIDParamsWithHTTPClient(client *http.Client) *GetAllSDTListByDeviceIDParams {
	return &GetAllSDTListByDeviceIDParams{
		HTTPClient: client,
	}
}

/* GetAllSDTListByDeviceIDParams contains all the parameters to send to the API endpoint
   for the get all SDT list by device Id operation.

   Typically these are written to a http.Request.
*/
type GetAllSDTListByDeviceIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// End.
	//
	// Format: int64
	End *int64

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

	// NetflowFilter.
	NetflowFilter *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all SDT list by device Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSDTListByDeviceIDParams) WithDefaults() *GetAllSDTListByDeviceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all SDT list by device Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSDTListByDeviceIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetAllSDTListByDeviceIDParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithTimeout(timeout time.Duration) *GetAllSDTListByDeviceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithContext(ctx context.Context) *GetAllSDTListByDeviceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithHTTPClient(client *http.Client) *GetAllSDTListByDeviceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithUserAgent(userAgent *string) *GetAllSDTListByDeviceIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithEnd adds the end to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithEnd(end *int64) *GetAllSDTListByDeviceIDParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetEnd(end *int64) {
	o.End = end
}

// WithFields adds the fields to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithFields(fields *string) *GetAllSDTListByDeviceIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithFilter(filter *string) *GetAllSDTListByDeviceIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithID(id int32) *GetAllSDTListByDeviceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetID(id int32) {
	o.ID = id
}

// WithNetflowFilter adds the netflowFilter to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithNetflowFilter(netflowFilter *string) *GetAllSDTListByDeviceIDParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithOffset adds the offset to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithOffset(offset *int32) *GetAllSDTListByDeviceIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithSize(size *int32) *GetAllSDTListByDeviceIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetSize(size *int32) {
	o.Size = size
}

// WithStart adds the start to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) WithStart(start *int64) *GetAllSDTListByDeviceIDParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get all SDT list by device Id params
func (o *GetAllSDTListByDeviceIDParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllSDTListByDeviceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}
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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.NetflowFilter != nil {

		// query param netflowFilter
		var qrNetflowFilter string

		if o.NetflowFilter != nil {
			qrNetflowFilter = *o.NetflowFilter
		}
		qNetflowFilter := qrNetflowFilter
		if qNetflowFilter != "" {

			if err := r.SetQueryParam("netflowFilter", qNetflowFilter); err != nil {
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

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
