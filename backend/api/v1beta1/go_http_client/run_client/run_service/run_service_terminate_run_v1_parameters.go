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

// NewRunServiceTerminateRunV1Params creates a new RunServiceTerminateRunV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRunServiceTerminateRunV1Params() *RunServiceTerminateRunV1Params {
	return &RunServiceTerminateRunV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRunServiceTerminateRunV1ParamsWithTimeout creates a new RunServiceTerminateRunV1Params object
// with the ability to set a timeout on a request.
func NewRunServiceTerminateRunV1ParamsWithTimeout(timeout time.Duration) *RunServiceTerminateRunV1Params {
	return &RunServiceTerminateRunV1Params{
		timeout: timeout,
	}
}

// NewRunServiceTerminateRunV1ParamsWithContext creates a new RunServiceTerminateRunV1Params object
// with the ability to set a context for a request.
func NewRunServiceTerminateRunV1ParamsWithContext(ctx context.Context) *RunServiceTerminateRunV1Params {
	return &RunServiceTerminateRunV1Params{
		Context: ctx,
	}
}

// NewRunServiceTerminateRunV1ParamsWithHTTPClient creates a new RunServiceTerminateRunV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewRunServiceTerminateRunV1ParamsWithHTTPClient(client *http.Client) *RunServiceTerminateRunV1Params {
	return &RunServiceTerminateRunV1Params{
		HTTPClient: client,
	}
}

/*
RunServiceTerminateRunV1Params contains all the parameters to send to the API endpoint

	for the run service terminate run v1 operation.

	Typically these are written to a http.Request.
*/
type RunServiceTerminateRunV1Params struct {

	/* RunID.

	   The ID of the run to be terminated.
	*/
	RunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the run service terminate run v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunServiceTerminateRunV1Params) WithDefaults() *RunServiceTerminateRunV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the run service terminate run v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunServiceTerminateRunV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the run service terminate run v1 params
func (o *RunServiceTerminateRunV1Params) WithTimeout(timeout time.Duration) *RunServiceTerminateRunV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run service terminate run v1 params
func (o *RunServiceTerminateRunV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run service terminate run v1 params
func (o *RunServiceTerminateRunV1Params) WithContext(ctx context.Context) *RunServiceTerminateRunV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run service terminate run v1 params
func (o *RunServiceTerminateRunV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run service terminate run v1 params
func (o *RunServiceTerminateRunV1Params) WithHTTPClient(client *http.Client) *RunServiceTerminateRunV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run service terminate run v1 params
func (o *RunServiceTerminateRunV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRunID adds the runID to the run service terminate run v1 params
func (o *RunServiceTerminateRunV1Params) WithRunID(runID string) *RunServiceTerminateRunV1Params {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the run service terminate run v1 params
func (o *RunServiceTerminateRunV1Params) SetRunID(runID string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *RunServiceTerminateRunV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param run_id
	if err := r.SetPathParam("run_id", o.RunID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
