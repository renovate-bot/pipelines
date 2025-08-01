// Code generated by go-swagger; DO NOT EDIT.

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/experiment_model"
)

// ExperimentServiceGetExperimentV1Reader is a Reader for the ExperimentServiceGetExperimentV1 structure.
type ExperimentServiceGetExperimentV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExperimentServiceGetExperimentV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExperimentServiceGetExperimentV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExperimentServiceGetExperimentV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExperimentServiceGetExperimentV1OK creates a ExperimentServiceGetExperimentV1OK with default headers values
func NewExperimentServiceGetExperimentV1OK() *ExperimentServiceGetExperimentV1OK {
	return &ExperimentServiceGetExperimentV1OK{}
}

/*
ExperimentServiceGetExperimentV1OK describes a response with status code 200, with default header values.

A successful response.
*/
type ExperimentServiceGetExperimentV1OK struct {
	Payload *experiment_model.APIExperiment
}

// IsSuccess returns true when this experiment service get experiment v1 o k response has a 2xx status code
func (o *ExperimentServiceGetExperimentV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this experiment service get experiment v1 o k response has a 3xx status code
func (o *ExperimentServiceGetExperimentV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this experiment service get experiment v1 o k response has a 4xx status code
func (o *ExperimentServiceGetExperimentV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this experiment service get experiment v1 o k response has a 5xx status code
func (o *ExperimentServiceGetExperimentV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this experiment service get experiment v1 o k response a status code equal to that given
func (o *ExperimentServiceGetExperimentV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the experiment service get experiment v1 o k response
func (o *ExperimentServiceGetExperimentV1OK) Code() int {
	return 200
}

func (o *ExperimentServiceGetExperimentV1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /apis/v1beta1/experiments/{id}][%d] experimentServiceGetExperimentV1OK %s", 200, payload)
}

func (o *ExperimentServiceGetExperimentV1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /apis/v1beta1/experiments/{id}][%d] experimentServiceGetExperimentV1OK %s", 200, payload)
}

func (o *ExperimentServiceGetExperimentV1OK) GetPayload() *experiment_model.APIExperiment {
	return o.Payload
}

func (o *ExperimentServiceGetExperimentV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(experiment_model.APIExperiment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExperimentServiceGetExperimentV1Default creates a ExperimentServiceGetExperimentV1Default with default headers values
func NewExperimentServiceGetExperimentV1Default(code int) *ExperimentServiceGetExperimentV1Default {
	return &ExperimentServiceGetExperimentV1Default{
		_statusCode: code,
	}
}

/*
ExperimentServiceGetExperimentV1Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ExperimentServiceGetExperimentV1Default struct {
	_statusCode int

	Payload *experiment_model.GooglerpcStatus
}

// IsSuccess returns true when this experiment service get experiment v1 default response has a 2xx status code
func (o *ExperimentServiceGetExperimentV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this experiment service get experiment v1 default response has a 3xx status code
func (o *ExperimentServiceGetExperimentV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this experiment service get experiment v1 default response has a 4xx status code
func (o *ExperimentServiceGetExperimentV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this experiment service get experiment v1 default response has a 5xx status code
func (o *ExperimentServiceGetExperimentV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this experiment service get experiment v1 default response a status code equal to that given
func (o *ExperimentServiceGetExperimentV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the experiment service get experiment v1 default response
func (o *ExperimentServiceGetExperimentV1Default) Code() int {
	return o._statusCode
}

func (o *ExperimentServiceGetExperimentV1Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /apis/v1beta1/experiments/{id}][%d] ExperimentService_GetExperimentV1 default %s", o._statusCode, payload)
}

func (o *ExperimentServiceGetExperimentV1Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /apis/v1beta1/experiments/{id}][%d] ExperimentService_GetExperimentV1 default %s", o._statusCode, payload)
}

func (o *ExperimentServiceGetExperimentV1Default) GetPayload() *experiment_model.GooglerpcStatus {
	return o.Payload
}

func (o *ExperimentServiceGetExperimentV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(experiment_model.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
