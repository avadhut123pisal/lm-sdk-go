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

// NewScheduleAutoDiscoveryByDeviceIDParams creates a new ScheduleAutoDiscoveryByDeviceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScheduleAutoDiscoveryByDeviceIDParams() *ScheduleAutoDiscoveryByDeviceIDParams {
	return &ScheduleAutoDiscoveryByDeviceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleAutoDiscoveryByDeviceIDParamsWithTimeout creates a new ScheduleAutoDiscoveryByDeviceIDParams object
// with the ability to set a timeout on a request.
func NewScheduleAutoDiscoveryByDeviceIDParamsWithTimeout(timeout time.Duration) *ScheduleAutoDiscoveryByDeviceIDParams {
	return &ScheduleAutoDiscoveryByDeviceIDParams{
		timeout: timeout,
	}
}

// NewScheduleAutoDiscoveryByDeviceIDParamsWithContext creates a new ScheduleAutoDiscoveryByDeviceIDParams object
// with the ability to set a context for a request.
func NewScheduleAutoDiscoveryByDeviceIDParamsWithContext(ctx context.Context) *ScheduleAutoDiscoveryByDeviceIDParams {
	return &ScheduleAutoDiscoveryByDeviceIDParams{
		Context: ctx,
	}
}

// NewScheduleAutoDiscoveryByDeviceIDParamsWithHTTPClient creates a new ScheduleAutoDiscoveryByDeviceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewScheduleAutoDiscoveryByDeviceIDParamsWithHTTPClient(client *http.Client) *ScheduleAutoDiscoveryByDeviceIDParams {
	return &ScheduleAutoDiscoveryByDeviceIDParams{
		HTTPClient: client,
	}
}

/* ScheduleAutoDiscoveryByDeviceIDParams contains all the parameters to send to the API endpoint
   for the schedule auto discovery by device Id operation.

   Typically these are written to a http.Request.
*/
type ScheduleAutoDiscoveryByDeviceIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// End.
	//
	// Format: int64
	End *int64

	// ID.
	//
	// Format: int32
	ID int32

	// NetflowFilter.
	NetflowFilter *string

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schedule auto discovery by device Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithDefaults() *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedule auto discovery by device Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")
	)

	val := ScheduleAutoDiscoveryByDeviceIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithTimeout(timeout time.Duration) *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithContext(ctx context.Context) *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithHTTPClient(client *http.Client) *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithUserAgent(userAgent *string) *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithEnd adds the end to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithEnd(end *int64) *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetEnd(end *int64) {
	o.End = end
}

// WithID adds the id to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithID(id int32) *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetID(id int32) {
	o.ID = id
}

// WithNetflowFilter adds the netflowFilter to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithNetflowFilter(netflowFilter *string) *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithStart adds the start to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WithStart(start *int64) *ScheduleAutoDiscoveryByDeviceIDParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the schedule auto discovery by device Id params
func (o *ScheduleAutoDiscoveryByDeviceIDParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleAutoDiscoveryByDeviceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
