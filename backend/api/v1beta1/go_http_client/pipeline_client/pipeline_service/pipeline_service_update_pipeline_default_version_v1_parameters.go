// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

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

// NewPipelineServiceUpdatePipelineDefaultVersionV1Params creates a new PipelineServiceUpdatePipelineDefaultVersionV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPipelineServiceUpdatePipelineDefaultVersionV1Params() *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	return &PipelineServiceUpdatePipelineDefaultVersionV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPipelineServiceUpdatePipelineDefaultVersionV1ParamsWithTimeout creates a new PipelineServiceUpdatePipelineDefaultVersionV1Params object
// with the ability to set a timeout on a request.
func NewPipelineServiceUpdatePipelineDefaultVersionV1ParamsWithTimeout(timeout time.Duration) *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	return &PipelineServiceUpdatePipelineDefaultVersionV1Params{
		timeout: timeout,
	}
}

// NewPipelineServiceUpdatePipelineDefaultVersionV1ParamsWithContext creates a new PipelineServiceUpdatePipelineDefaultVersionV1Params object
// with the ability to set a context for a request.
func NewPipelineServiceUpdatePipelineDefaultVersionV1ParamsWithContext(ctx context.Context) *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	return &PipelineServiceUpdatePipelineDefaultVersionV1Params{
		Context: ctx,
	}
}

// NewPipelineServiceUpdatePipelineDefaultVersionV1ParamsWithHTTPClient creates a new PipelineServiceUpdatePipelineDefaultVersionV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewPipelineServiceUpdatePipelineDefaultVersionV1ParamsWithHTTPClient(client *http.Client) *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	return &PipelineServiceUpdatePipelineDefaultVersionV1Params{
		HTTPClient: client,
	}
}

/*
PipelineServiceUpdatePipelineDefaultVersionV1Params contains all the parameters to send to the API endpoint

	for the pipeline service update pipeline default version v1 operation.

	Typically these are written to a http.Request.
*/
type PipelineServiceUpdatePipelineDefaultVersionV1Params struct {

	/* PipelineID.

	   The ID of the pipeline to be updated.
	*/
	PipelineID string

	/* VersionID.

	   The ID of the default version.
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pipeline service update pipeline default version v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) WithDefaults() *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pipeline service update pipeline default version v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) WithTimeout(timeout time.Duration) *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) WithContext(ctx context.Context) *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) WithHTTPClient(client *http.Client) *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPipelineID adds the pipelineID to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) WithPipelineID(pipelineID string) *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	o.SetPipelineID(pipelineID)
	return o
}

// SetPipelineID adds the pipelineId to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) SetPipelineID(pipelineID string) {
	o.PipelineID = pipelineID
}

// WithVersionID adds the versionID to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) WithVersionID(versionID string) *PipelineServiceUpdatePipelineDefaultVersionV1Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the pipeline service update pipeline default version v1 params
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *PipelineServiceUpdatePipelineDefaultVersionV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pipeline_id
	if err := r.SetPathParam("pipeline_id", o.PipelineID); err != nil {
		return err
	}

	// path param version_id
	if err := r.SetPathParam("version_id", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
