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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchDeviceParams creates a new PatchDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDeviceParams() *PatchDeviceParams {
	return &PatchDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDeviceParamsWithTimeout creates a new PatchDeviceParams object
// with the ability to set a timeout on a request.
func NewPatchDeviceParamsWithTimeout(timeout time.Duration) *PatchDeviceParams {
	return &PatchDeviceParams{
		timeout: timeout,
	}
}

// NewPatchDeviceParamsWithContext creates a new PatchDeviceParams object
// with the ability to set a context for a request.
func NewPatchDeviceParamsWithContext(ctx context.Context) *PatchDeviceParams {
	return &PatchDeviceParams{
		Context: ctx,
	}
}

// NewPatchDeviceParamsWithHTTPClient creates a new PatchDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDeviceParamsWithHTTPClient(client *http.Client) *PatchDeviceParams {
	return &PatchDeviceParams{
		HTTPClient: client,
	}
}

/* PatchDeviceParams contains all the parameters to send to the API endpoint
   for the patch device operation.

   Typically these are written to a http.Request.
*/
type PatchDeviceParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty"
	UserAgent *string

	// Body.
	Body *models.Device

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

	// OpType.
	//
	// Default: "refresh"
	OpType *string

	// PatchFields.
	PatchFields *string

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceParams) WithDefaults() *PatchDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-7-gdde4eda-dirty")

		opTypeDefault = string("refresh")
	)

	val := PatchDeviceParams{
		UserAgent: &userAgentDefault,
		OpType:    &opTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch device params
func (o *PatchDeviceParams) WithTimeout(timeout time.Duration) *PatchDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch device params
func (o *PatchDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch device params
func (o *PatchDeviceParams) WithContext(ctx context.Context) *PatchDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch device params
func (o *PatchDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch device params
func (o *PatchDeviceParams) WithHTTPClient(client *http.Client) *PatchDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch device params
func (o *PatchDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the patch device params
func (o *PatchDeviceParams) WithUserAgent(userAgent *string) *PatchDeviceParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch device params
func (o *PatchDeviceParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch device params
func (o *PatchDeviceParams) WithBody(body *models.Device) *PatchDeviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch device params
func (o *PatchDeviceParams) SetBody(body *models.Device) {
	o.Body = body
}

// WithEnd adds the end to the patch device params
func (o *PatchDeviceParams) WithEnd(end *int64) *PatchDeviceParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the patch device params
func (o *PatchDeviceParams) SetEnd(end *int64) {
	o.End = end
}

// WithID adds the id to the patch device params
func (o *PatchDeviceParams) WithID(id int32) *PatchDeviceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch device params
func (o *PatchDeviceParams) SetID(id int32) {
	o.ID = id
}

// WithNetflowFilter adds the netflowFilter to the patch device params
func (o *PatchDeviceParams) WithNetflowFilter(netflowFilter *string) *PatchDeviceParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the patch device params
func (o *PatchDeviceParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithOpType adds the opType to the patch device params
func (o *PatchDeviceParams) WithOpType(opType *string) *PatchDeviceParams {
	o.SetOpType(opType)
	return o
}

// SetOpType adds the opType to the patch device params
func (o *PatchDeviceParams) SetOpType(opType *string) {
	o.OpType = opType
}

// WithPatchFields adds the patchFields to the patch device params
func (o *PatchDeviceParams) WithPatchFields(patchFields *string) *PatchDeviceParams {
	o.SetPatchFields(patchFields)
	return o
}

// SetPatchFields adds the patchFields to the patch device params
func (o *PatchDeviceParams) SetPatchFields(patchFields *string) {
	o.PatchFields = patchFields
}

// WithStart adds the start to the patch device params
func (o *PatchDeviceParams) WithStart(start *int64) *PatchDeviceParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the patch device params
func (o *PatchDeviceParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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

	if o.OpType != nil {

		// query param opType
		var qrOpType string

		if o.OpType != nil {
			qrOpType = *o.OpType
		}
		qOpType := qrOpType
		if qOpType != "" {

			if err := r.SetQueryParam("opType", qOpType); err != nil {
				return err
			}
		}
	}

	if o.PatchFields != nil {

		// query param patchFields
		var qrPatchFields string

		if o.PatchFields != nil {
			qrPatchFields = *o.PatchFields
		}
		qPatchFields := qrPatchFields
		if qPatchFields != "" {

			if err := r.SetQueryParam("patchFields", qPatchFields); err != nil {
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
