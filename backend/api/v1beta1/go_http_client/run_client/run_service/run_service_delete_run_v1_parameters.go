// Code generated by go-swagger; DO NOT EDIT.

package run_service

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
)

// NewRunServiceDeleteRunV1Params creates a new RunServiceDeleteRunV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRunServiceDeleteRunV1Params() *RunServiceDeleteRunV1Params {
	return &RunServiceDeleteRunV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRunServiceDeleteRunV1ParamsWithTimeout creates a new RunServiceDeleteRunV1Params object
// with the ability to set a timeout on a request.
func NewRunServiceDeleteRunV1ParamsWithTimeout(timeout time.Duration) *RunServiceDeleteRunV1Params {
	return &RunServiceDeleteRunV1Params{
		timeout: timeout,
	}
}

// NewRunServiceDeleteRunV1ParamsWithContext creates a new RunServiceDeleteRunV1Params object
// with the ability to set a context for a request.
func NewRunServiceDeleteRunV1ParamsWithContext(ctx context.Context) *RunServiceDeleteRunV1Params {
	return &RunServiceDeleteRunV1Params{
		Context: ctx,
	}
}

// NewRunServiceDeleteRunV1ParamsWithHTTPClient creates a new RunServiceDeleteRunV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewRunServiceDeleteRunV1ParamsWithHTTPClient(client *http.Client) *RunServiceDeleteRunV1Params {
	return &RunServiceDeleteRunV1Params{
		HTTPClient: client,
	}
}

/*
RunServiceDeleteRunV1Params contains all the parameters to send to the API endpoint

	for the run service delete run v1 operation.

	Typically these are written to a http.Request.
*/
type RunServiceDeleteRunV1Params struct {

	/* ID.

	   The ID of the run to be deleted.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the run service delete run v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunServiceDeleteRunV1Params) WithDefaults() *RunServiceDeleteRunV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the run service delete run v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunServiceDeleteRunV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the run service delete run v1 params
func (o *RunServiceDeleteRunV1Params) WithTimeout(timeout time.Duration) *RunServiceDeleteRunV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run service delete run v1 params
func (o *RunServiceDeleteRunV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run service delete run v1 params
func (o *RunServiceDeleteRunV1Params) WithContext(ctx context.Context) *RunServiceDeleteRunV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run service delete run v1 params
func (o *RunServiceDeleteRunV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run service delete run v1 params
func (o *RunServiceDeleteRunV1Params) WithHTTPClient(client *http.Client) *RunServiceDeleteRunV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run service delete run v1 params
func (o *RunServiceDeleteRunV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the run service delete run v1 params
func (o *RunServiceDeleteRunV1Params) WithID(id string) *RunServiceDeleteRunV1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the run service delete run v1 params
func (o *RunServiceDeleteRunV1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RunServiceDeleteRunV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
