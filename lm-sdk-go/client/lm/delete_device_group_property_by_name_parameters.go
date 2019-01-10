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

// NewDeleteDeviceGroupPropertyByNameParams creates a new DeleteDeviceGroupPropertyByNameParams object
// with the default values initialized.
func NewDeleteDeviceGroupPropertyByNameParams() *DeleteDeviceGroupPropertyByNameParams {
	var ()
	return &DeleteDeviceGroupPropertyByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceGroupPropertyByNameParamsWithTimeout creates a new DeleteDeviceGroupPropertyByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeviceGroupPropertyByNameParamsWithTimeout(timeout time.Duration) *DeleteDeviceGroupPropertyByNameParams {
	var ()
	return &DeleteDeviceGroupPropertyByNameParams{

		timeout: timeout,
	}
}

// NewDeleteDeviceGroupPropertyByNameParamsWithContext creates a new DeleteDeviceGroupPropertyByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeviceGroupPropertyByNameParamsWithContext(ctx context.Context) *DeleteDeviceGroupPropertyByNameParams {
	var ()
	return &DeleteDeviceGroupPropertyByNameParams{

		Context: ctx,
	}
}

// NewDeleteDeviceGroupPropertyByNameParamsWithHTTPClient creates a new DeleteDeviceGroupPropertyByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeviceGroupPropertyByNameParamsWithHTTPClient(client *http.Client) *DeleteDeviceGroupPropertyByNameParams {
	var ()
	return &DeleteDeviceGroupPropertyByNameParams{
		HTTPClient: client,
	}
}

/*DeleteDeviceGroupPropertyByNameParams contains all the parameters to send to the API endpoint
for the delete device group property by name operation typically these are written to a http.Request
*/
type DeleteDeviceGroupPropertyByNameParams struct {

	/*Gid
	  group ID

	*/
	Gid int32
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) WithTimeout(timeout time.Duration) *DeleteDeviceGroupPropertyByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) WithContext(ctx context.Context) *DeleteDeviceGroupPropertyByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) WithHTTPClient(client *http.Client) *DeleteDeviceGroupPropertyByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGid adds the gid to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) WithGid(gid int32) *DeleteDeviceGroupPropertyByNameParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) SetGid(gid int32) {
	o.Gid = gid
}

// WithName adds the name to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) WithName(name string) *DeleteDeviceGroupPropertyByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete device group property by name params
func (o *DeleteDeviceGroupPropertyByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceGroupPropertyByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gid
	if err := r.SetPathParam("gid", swag.FormatInt32(o.Gid)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}