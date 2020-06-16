// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterStatus cluster status
//
// swagger:model ClusterStatus
type ClusterStatus struct {

	// Current version of the cluster. This can refer to a commit hash or
	// any valid version string in the context.
	//
	CurrentVersion string `json:"current_version,omitempty"`

	// Last working version of the cluster. This can refer to a commit
	// hash or any valid version string in the context. In case any
	// problems are defined for the current_version then it should be
	// safe to roll back to this last version.
	//
	LastVersion string `json:"last_version,omitempty"`

	// Next version of the cluster. This field indicates that the cluster is
	// being updated to a new version. This can refer to a commit hash or any
	// valid version string in the context.
	//
	NextVersion string `json:"next_version,omitempty"`

	// problems
	Problems []*ClusterStatusProblemsItems0 `json:"problems"`
}

// Validate validates this cluster status
func (m *ClusterStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProblems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatus) validateProblems(formats strfmt.Registry) error {

	if swag.IsZero(m.Problems) { // not required
		return nil
	}

	for i := 0; i < len(m.Problems); i++ {
		if swag.IsZero(m.Problems[i]) { // not required
			continue
		}

		if m.Problems[i] != nil {
			if err := m.Problems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("problems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatus) UnmarshalBinary(b []byte) error {
	var res ClusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterStatusProblemsItems0 cluster status problems items0
//
// swagger:model ClusterStatusProblemsItems0
type ClusterStatusProblemsItems0 struct {

	// A human-readable explanation specific to this occurrence of
	// the problem.
	//
	Detail string `json:"detail,omitempty"`

	// A URI reference that identifies the specific occurrence of
	// the problem.
	//
	Instance string `json:"instance,omitempty"`

	// The HTTP status code generated by the origin server for this
	// occurence of the problem.
	//
	Status int32 `json:"status,omitempty"`

	// A short, human-readable summary of the problem type.
	//
	// Required: true
	Title *string `json:"title"`

	// A URI reference the indentifies the problem type.
	// Required: true
	Type *string `json:"type"`

	// cluster status problems items0
	ClusterStatusProblemsItems0 map[string]string `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *ClusterStatusProblemsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// A human-readable explanation specific to this occurrence of
		// the problem.
		//
		Detail string `json:"detail,omitempty"`

		// A URI reference that identifies the specific occurrence of
		// the problem.
		//
		Instance string `json:"instance,omitempty"`

		// The HTTP status code generated by the origin server for this
		// occurence of the problem.
		//
		Status int32 `json:"status,omitempty"`

		// A short, human-readable summary of the problem type.
		//
		// Required: true
		Title *string `json:"title"`

		// A URI reference the indentifies the problem type.
		// Required: true
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ClusterStatusProblemsItems0

	rcv.Detail = stage1.Detail
	rcv.Instance = stage1.Instance
	rcv.Status = stage1.Status
	rcv.Title = stage1.Title
	rcv.Type = stage1.Type
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "detail")
	delete(stage2, "instance")
	delete(stage2, "status")
	delete(stage2, "title")
	delete(stage2, "type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]string)
		for k, v := range stage2 {
			var toadd string
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.ClusterStatusProblemsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m ClusterStatusProblemsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// A human-readable explanation specific to this occurrence of
		// the problem.
		//
		Detail string `json:"detail,omitempty"`

		// A URI reference that identifies the specific occurrence of
		// the problem.
		//
		Instance string `json:"instance,omitempty"`

		// The HTTP status code generated by the origin server for this
		// occurence of the problem.
		//
		Status int32 `json:"status,omitempty"`

		// A short, human-readable summary of the problem type.
		//
		// Required: true
		Title *string `json:"title"`

		// A URI reference the indentifies the problem type.
		// Required: true
		Type *string `json:"type"`
	}

	stage1.Detail = m.Detail
	stage1.Instance = m.Instance
	stage1.Status = m.Status
	stage1.Title = m.Title
	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.ClusterStatusProblemsItems0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.ClusterStatusProblemsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this cluster status problems items0
func (m *ClusterStatusProblemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatusProblemsItems0) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *ClusterStatusProblemsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatusProblemsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatusProblemsItems0) UnmarshalBinary(b []byte) error {
	var res ClusterStatusProblemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
