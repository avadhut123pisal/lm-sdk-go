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

// NewGetWebsiteCheckpointDataByIDParams creates a new GetWebsiteCheckpointDataByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsiteCheckpointDataByIDParams() *GetWebsiteCheckpointDataByIDParams {
	return &GetWebsiteCheckpointDataByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsiteCheckpointDataByIDParamsWithTimeout creates a new GetWebsiteCheckpointDataByIDParams object
// with the ability to set a timeout on a request.
func NewGetWebsiteCheckpointDataByIDParamsWithTimeout(timeout time.Duration) *GetWebsiteCheckpointDataByIDParams {
	return &GetWebsiteCheckpointDataByIDParams{
		timeout: timeout,
	}
}

// NewGetWebsiteCheckpointDataByIDParamsWithContext creates a new GetWebsiteCheckpointDataByIDParams object
// with the ability to set a context for a request.
func NewGetWebsiteCheckpointDataByIDParamsWithContext(ctx context.Context) *GetWebsiteCheckpointDataByIDParams {
	return &GetWebsiteCheckpointDataByIDParams{
		Context: ctx,
	}
}

// NewGetWebsiteCheckpointDataByIDParamsWithHTTPClient creates a new GetWebsiteCheckpointDataByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsiteCheckpointDataByIDParamsWithHTTPClient(client *http.Client) *GetWebsiteCheckpointDataByIDParams {
	return &GetWebsiteCheckpointDataByIDParams{
		HTTPClient: client,
	}
}

/* GetWebsiteCheckpointDataByIDParams contains all the parameters to send to the API endpoint
   for the get website checkpoint data by Id operation.

   Typically these are written to a http.Request.
*/
type GetWebsiteCheckpointDataByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1"
	UserAgent *string

	/* Aggregate.

	   the aggregate option

	   Default: "none"
	*/
	Aggregate *string

	// CheckID.
	//
	// Format: int32
	CheckID int32

	// Datapoints.
	Datapoints *string

	// End.
	//
	// Format: int64
	End *int64

	// Format.
	//
	// Default: "json"
	Format *string

	// Period.
	//
	// Format: double
	// Default: 1
	Period *float64

	// SrvID.
	//
	// Format: int32
	SrvID int32

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get website checkpoint data by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteCheckpointDataByIDParams) WithDefaults() *GetWebsiteCheckpointDataByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get website checkpoint data by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteCheckpointDataByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1")

		aggregateDefault = string("none")

		endDefault = int64(0)

		formatDefault = string("json")

		periodDefault = float64(1)

		startDefault = int64(0)
	)

	val := GetWebsiteCheckpointDataByIDParams{
		UserAgent: &userAgentDefault,
		Aggregate: &aggregateDefault,
		End:       &endDefault,
		Format:    &formatDefault,
		Period:    &periodDefault,
		Start:     &startDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithTimeout(timeout time.Duration) *GetWebsiteCheckpointDataByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithContext(ctx context.Context) *GetWebsiteCheckpointDataByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithHTTPClient(client *http.Client) *GetWebsiteCheckpointDataByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithUserAgent(userAgent *string) *GetWebsiteCheckpointDataByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithAggregate adds the aggregate to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithAggregate(aggregate *string) *GetWebsiteCheckpointDataByIDParams {
	o.SetAggregate(aggregate)
	return o
}

// SetAggregate adds the aggregate to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetAggregate(aggregate *string) {
	o.Aggregate = aggregate
}

// WithCheckID adds the checkID to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithCheckID(checkID int32) *GetWebsiteCheckpointDataByIDParams {
	o.SetCheckID(checkID)
	return o
}

// SetCheckID adds the checkId to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetCheckID(checkID int32) {
	o.CheckID = checkID
}

// WithDatapoints adds the datapoints to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithDatapoints(datapoints *string) *GetWebsiteCheckpointDataByIDParams {
	o.SetDatapoints(datapoints)
	return o
}

// SetDatapoints adds the datapoints to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetDatapoints(datapoints *string) {
	o.Datapoints = datapoints
}

// WithEnd adds the end to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithEnd(end *int64) *GetWebsiteCheckpointDataByIDParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetEnd(end *int64) {
	o.End = end
}

// WithFormat adds the format to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithFormat(format *string) *GetWebsiteCheckpointDataByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithPeriod adds the period to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithPeriod(period *float64) *GetWebsiteCheckpointDataByIDParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetPeriod(period *float64) {
	o.Period = period
}

// WithSrvID adds the srvID to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithSrvID(srvID int32) *GetWebsiteCheckpointDataByIDParams {
	o.SetSrvID(srvID)
	return o
}

// SetSrvID adds the srvId to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetSrvID(srvID int32) {
	o.SrvID = srvID
}

// WithStart adds the start to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) WithStart(start *int64) *GetWebsiteCheckpointDataByIDParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get website checkpoint data by Id params
func (o *GetWebsiteCheckpointDataByIDParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteCheckpointDataByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Aggregate != nil {

		// query param aggregate
		var qrAggregate string

		if o.Aggregate != nil {
			qrAggregate = *o.Aggregate
		}
		qAggregate := qrAggregate
		if qAggregate != "" {

			if err := r.SetQueryParam("aggregate", qAggregate); err != nil {
				return err
			}
		}
	}

	// path param checkId
	if err := r.SetPathParam("checkId", swag.FormatInt32(o.CheckID)); err != nil {
		return err
	}

	if o.Datapoints != nil {

		// query param datapoints
		var qrDatapoints string

		if o.Datapoints != nil {
			qrDatapoints = *o.Datapoints
		}
		qDatapoints := qrDatapoints
		if qDatapoints != "" {

			if err := r.SetQueryParam("datapoints", qDatapoints); err != nil {
				return err
			}
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

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	if o.Period != nil {

		// query param period
		var qrPeriod float64

		if o.Period != nil {
			qrPeriod = *o.Period
		}
		qPeriod := swag.FormatFloat64(qrPeriod)
		if qPeriod != "" {

			if err := r.SetQueryParam("period", qPeriod); err != nil {
				return err
			}
		}
	}

	// path param srvId
	if err := r.SetPathParam("srvId", swag.FormatInt32(o.SrvID)); err != nil {
		return err
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
