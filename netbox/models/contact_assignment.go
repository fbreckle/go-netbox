// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContactAssignment contact assignment
//
// swagger:model ContactAssignment
type ContactAssignment struct {

	// contact
	// Required: true
	Contact *NestedContact `json:"contact"`

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Object
	// Read Only: true
	Object interface{} `json:"object,omitempty"`

	// Object id
	// Required: true
	// Minimum: 0
	ObjectID *int64 `json:"object_id"`

	// priority
	Priority *ContactAssignmentPriority `json:"priority,omitempty"`

	// role
	Role *NestedContactRole `json:"role,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this contact assignment
func (m *ContactAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactAssignment) validateContact(formats strfmt.Registry) error {

	if err := validate.Required("contact", "body", m.Contact); err != nil {
		return err
	}

	if m.Contact != nil {
		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *ContactAssignment) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *ContactAssignment) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactAssignment) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactAssignment) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("object_id", "body", m.ObjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("object_id", "body", *m.ObjectID, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ContactAssignment) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if m.Priority != nil {
		if err := m.Priority.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priority")
			}
			return err
		}
	}

	return nil
}

func (m *ContactAssignment) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *ContactAssignment) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this contact assignment based on the context it is used
func (m *ContactAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactAssignment) contextValidateContact(ctx context.Context, formats strfmt.Registry) error {

	if m.Contact != nil {
		if err := m.Contact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *ContactAssignment) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *ContactAssignment) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ContactAssignment) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ContactAssignment) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *ContactAssignment) contextValidatePriority(ctx context.Context, formats strfmt.Registry) error {

	if m.Priority != nil {
		if err := m.Priority.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priority")
			}
			return err
		}
	}

	return nil
}

func (m *ContactAssignment) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {
		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *ContactAssignment) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAssignment) UnmarshalBinary(b []byte) error {
	var res ContactAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContactAssignmentPriority Priority
//
// swagger:model ContactAssignmentPriority
type ContactAssignmentPriority struct {

	// label
	// Required: true
	// Enum: [Primary Secondary Tertiary Inactive]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [primary secondary tertiary inactive]
	Value *string `json:"value"`
}

// Validate validates this contact assignment priority
func (m *ContactAssignmentPriority) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var contactAssignmentPriorityTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Primary","Secondary","Tertiary","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactAssignmentPriorityTypeLabelPropEnum = append(contactAssignmentPriorityTypeLabelPropEnum, v)
	}
}

const (

	// ContactAssignmentPriorityLabelPrimary captures enum value "Primary"
	ContactAssignmentPriorityLabelPrimary string = "Primary"

	// ContactAssignmentPriorityLabelSecondary captures enum value "Secondary"
	ContactAssignmentPriorityLabelSecondary string = "Secondary"

	// ContactAssignmentPriorityLabelTertiary captures enum value "Tertiary"
	ContactAssignmentPriorityLabelTertiary string = "Tertiary"

	// ContactAssignmentPriorityLabelInactive captures enum value "Inactive"
	ContactAssignmentPriorityLabelInactive string = "Inactive"
)

// prop value enum
func (m *ContactAssignmentPriority) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contactAssignmentPriorityTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContactAssignmentPriority) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("priority"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("priority"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var contactAssignmentPriorityTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","secondary","tertiary","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactAssignmentPriorityTypeValuePropEnum = append(contactAssignmentPriorityTypeValuePropEnum, v)
	}
}

const (

	// ContactAssignmentPriorityValuePrimary captures enum value "primary"
	ContactAssignmentPriorityValuePrimary string = "primary"

	// ContactAssignmentPriorityValueSecondary captures enum value "secondary"
	ContactAssignmentPriorityValueSecondary string = "secondary"

	// ContactAssignmentPriorityValueTertiary captures enum value "tertiary"
	ContactAssignmentPriorityValueTertiary string = "tertiary"

	// ContactAssignmentPriorityValueInactive captures enum value "inactive"
	ContactAssignmentPriorityValueInactive string = "inactive"
)

// prop value enum
func (m *ContactAssignmentPriority) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contactAssignmentPriorityTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContactAssignmentPriority) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("priority"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("priority"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this contact assignment priority based on context it is used
func (m *ContactAssignmentPriority) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactAssignmentPriority) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAssignmentPriority) UnmarshalBinary(b []byte) error {
	var res ContactAssignmentPriority
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
