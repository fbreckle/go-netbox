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

// NestedFHRPGroup nested f h r p group
//
// swagger:model NestedFHRPGroup
type NestedFHRPGroup struct {

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Group ID
	// Required: true
	// Maximum: 32767
	// Minimum: 0
	GroupID *int64 `json:"group_id"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Protocol
	// Required: true
	// Enum: [vrrp2 vrrp3 hsrp glbp carp other]
	Protocol *string `json:"protocol"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this nested f h r p group
func (m *NestedFHRPGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
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

func (m *NestedFHRPGroup) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	if err := validate.MinimumInt("group_id", "body", *m.GroupID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("group_id", "body", *m.GroupID, 32767, false); err != nil {
		return err
	}

	return nil
}

var nestedFHRPGroupTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vrrp2","vrrp3","hsrp","glbp","carp","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nestedFHRPGroupTypeProtocolPropEnum = append(nestedFHRPGroupTypeProtocolPropEnum, v)
	}
}

const (

	// NestedFHRPGroupProtocolVrrp2 captures enum value "vrrp2"
	NestedFHRPGroupProtocolVrrp2 string = "vrrp2"

	// NestedFHRPGroupProtocolVrrp3 captures enum value "vrrp3"
	NestedFHRPGroupProtocolVrrp3 string = "vrrp3"

	// NestedFHRPGroupProtocolHsrp captures enum value "hsrp"
	NestedFHRPGroupProtocolHsrp string = "hsrp"

	// NestedFHRPGroupProtocolGlbp captures enum value "glbp"
	NestedFHRPGroupProtocolGlbp string = "glbp"

	// NestedFHRPGroupProtocolCarp captures enum value "carp"
	NestedFHRPGroupProtocolCarp string = "carp"

	// NestedFHRPGroupProtocolOther captures enum value "other"
	NestedFHRPGroupProtocolOther string = "other"
)

// prop value enum
func (m *NestedFHRPGroup) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nestedFHRPGroupTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NestedFHRPGroup) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *NestedFHRPGroup) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested f h r p group based on the context it is used
func (m *NestedFHRPGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
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

func (m *NestedFHRPGroup) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *NestedFHRPGroup) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NestedFHRPGroup) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedFHRPGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedFHRPGroup) UnmarshalBinary(b []byte) error {
	var res NestedFHRPGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
