// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new run service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new run service API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new run service API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for run service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RunServiceArchiveRunV1(params *RunServiceArchiveRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceArchiveRunV1OK, error)

	RunServiceCreateRunV1(params *RunServiceCreateRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceCreateRunV1OK, error)

	RunServiceDeleteRunV1(params *RunServiceDeleteRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceDeleteRunV1OK, error)

	RunServiceGetRunV1(params *RunServiceGetRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceGetRunV1OK, error)

	RunServiceListRunsV1(params *RunServiceListRunsV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceListRunsV1OK, error)

	RunServiceReadArtifactV1(params *RunServiceReadArtifactV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceReadArtifactV1OK, error)

	RunServiceReportRunMetricsV1(params *RunServiceReportRunMetricsV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceReportRunMetricsV1OK, error)

	RunServiceRetryRunV1(params *RunServiceRetryRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceRetryRunV1OK, error)

	RunServiceTerminateRunV1(params *RunServiceTerminateRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceTerminateRunV1OK, error)

	RunServiceUnarchiveRunV1(params *RunServiceUnarchiveRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceUnarchiveRunV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
RunServiceArchiveRunV1 archives a run
*/
func (a *Client) RunServiceArchiveRunV1(params *RunServiceArchiveRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceArchiveRunV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceArchiveRunV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_ArchiveRunV1",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/runs/{id}:archive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceArchiveRunV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceArchiveRunV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceArchiveRunV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceCreateRunV1 creates a new run
*/
func (a *Client) RunServiceCreateRunV1(params *RunServiceCreateRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceCreateRunV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceCreateRunV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_CreateRunV1",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceCreateRunV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceCreateRunV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceCreateRunV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceDeleteRunV1 deletes a run
*/
func (a *Client) RunServiceDeleteRunV1(params *RunServiceDeleteRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceDeleteRunV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceDeleteRunV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_DeleteRunV1",
		Method:             "DELETE",
		PathPattern:        "/apis/v1beta1/runs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceDeleteRunV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceDeleteRunV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceDeleteRunV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceGetRunV1 finds a specific run by ID
*/
func (a *Client) RunServiceGetRunV1(params *RunServiceGetRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceGetRunV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceGetRunV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_GetRunV1",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/runs/{run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceGetRunV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceGetRunV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceGetRunV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceListRunsV1 finds all runs
*/
func (a *Client) RunServiceListRunsV1(params *RunServiceListRunsV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceListRunsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceListRunsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_ListRunsV1",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceListRunsV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceListRunsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceListRunsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceReadArtifactV1 finds a run s artifact data
*/
func (a *Client) RunServiceReadArtifactV1(params *RunServiceReadArtifactV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceReadArtifactV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceReadArtifactV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_ReadArtifactV1",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/runs/{run_id}/nodes/{node_id}/artifacts/{artifact_name}:read",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceReadArtifactV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceReadArtifactV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceReadArtifactV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceReportRunMetricsV1 reports run metrics reports metrics of a run each metric is reported in its own transaction so this API accepts partial failures metric can be uniquely identified by run id node id name duplicate reporting will be ignored by the API first reporting wins
*/
func (a *Client) RunServiceReportRunMetricsV1(params *RunServiceReportRunMetricsV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceReportRunMetricsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceReportRunMetricsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_ReportRunMetricsV1",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/runs/{run_id}:reportMetrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceReportRunMetricsV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceReportRunMetricsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceReportRunMetricsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceRetryRunV1 res initiates a failed or terminated run
*/
func (a *Client) RunServiceRetryRunV1(params *RunServiceRetryRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceRetryRunV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceRetryRunV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_RetryRunV1",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/runs/{run_id}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceRetryRunV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceRetryRunV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceRetryRunV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceTerminateRunV1 terminates an active run
*/
func (a *Client) RunServiceTerminateRunV1(params *RunServiceTerminateRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceTerminateRunV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceTerminateRunV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_TerminateRunV1",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/runs/{run_id}/terminate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceTerminateRunV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceTerminateRunV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceTerminateRunV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunServiceUnarchiveRunV1 restores an archived run
*/
func (a *Client) RunServiceUnarchiveRunV1(params *RunServiceUnarchiveRunV1Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunServiceUnarchiveRunV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceUnarchiveRunV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunService_UnarchiveRunV1",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/runs/{id}:unarchive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceUnarchiveRunV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunServiceUnarchiveRunV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunServiceUnarchiveRunV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
