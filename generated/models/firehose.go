// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Firehose firehose
//
// swagger:model Firehose
type Firehose struct {

	// cluster
	// Example: data_engineering
	Cluster string `json:"cluster,omitempty"`

	// configs
	Configs *FirehoseConfig `json:"configs,omitempty"`

	// created at
	// Example: 2022-06-23T16:49:15.885541Z
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// description
	// Example: This firehose consumes from booking events and ingests to redis
	Description string `json:"description,omitempty"`

	// name
	// Example: booking-events-ingester
	Name string `json:"name,omitempty"`

	// state
	State *FirehoseState `json:"state,omitempty"`

	// team
	// Example: pricing
	// Read Only: true
	Team string `json:"team,omitempty"`

	// title
	// Example: Booking Events Ingester
	Title string `json:"title,omitempty"`

	// updated at
	// Example: 2022-06-23T16:49:15.885541Z
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// urn
	// Example: orn:foo:firehose:fh1
	// Read Only: true
	Urn string `json:"urn,omitempty"`
}

// Validate validates this firehose
func (m *Firehose) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Firehose) validateConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.Configs) { // not required
		return nil
	}

	if m.Configs != nil {
		if err := m.Configs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configs")
			}
			return err
		}
	}

	return nil
}

func (m *Firehose) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Firehose) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *Firehose) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this firehose based on the context it is used
func (m *Firehose) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUrn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Firehose) contextValidateConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.Configs != nil {
		if err := m.Configs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configs")
			}
			return err
		}
	}

	return nil
}

func (m *Firehose) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Firehose) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *Firehose) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "team", "body", string(m.Team)); err != nil {
		return err
	}

	return nil
}

func (m *Firehose) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Firehose) contextValidateUrn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "urn", "body", string(m.Urn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Firehose) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Firehose) UnmarshalBinary(b []byte) error {
	var res Firehose
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
