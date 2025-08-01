// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunServiceReportRunMetricsV1Body run service report run metrics v1 body
//
// swagger:model RunServiceReportRunMetricsV1Body
type RunServiceReportRunMetricsV1Body struct {

	// List of metrics to report.
	Metrics []*APIRunMetric `json:"metrics"`
}

// Validate validates this run service report run metrics v1 body
func (m *RunServiceReportRunMetricsV1Body) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunServiceReportRunMetricsV1Body) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {
		if swag.IsZero(m.Metrics[i]) { // not required
			continue
		}

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this run service report run metrics v1 body based on the context it is used
func (m *RunServiceReportRunMetricsV1Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunServiceReportRunMetricsV1Body) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metrics); i++ {

		if m.Metrics[i] != nil {

			if swag.IsZero(m.Metrics[i]) { // not required
				return nil
			}

			if err := m.Metrics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunServiceReportRunMetricsV1Body) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunServiceReportRunMetricsV1Body) UnmarshalBinary(b []byte) error {
	var res RunServiceReportRunMetricsV1Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
