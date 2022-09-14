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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// L2VPN l2 v p n
//
// swagger:model L2VPN
type L2VPN struct {

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// export targets
	// Unique: true
	ExportTargets []*NestedRouteTarget `json:"export_targets"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Identifier
	Identifier *int64 `json:"identifier,omitempty"`

	// import targets
	// Unique: true
	ImportTargets []*NestedRouteTarget `json:"import_targets"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Slug
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// type
	Type *L2VPNType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this l2 v p n
func (m *L2VPN) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *L2VPN) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) validateExportTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportTargets) { // not required
		return nil
	}

	if err := validate.UniqueItems("export_targets", "body", m.ExportTargets); err != nil {
		return err
	}

	for i := 0; i < len(m.ExportTargets); i++ {
		if swag.IsZero(m.ExportTargets[i]) { // not required
			continue
		}

		if m.ExportTargets[i] != nil {
			if err := m.ExportTargets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("export_targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("export_targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *L2VPN) validateImportTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportTargets) { // not required
		return nil
	}

	if err := validate.UniqueItems("import_targets", "body", m.ImportTargets); err != nil {
		return err
	}

	for i := 0; i < len(m.ImportTargets); i++ {
		if swag.IsZero(m.ImportTargets[i]) { // not required
			continue
		}

		if m.ImportTargets[i] != nil {
			if err := m.ImportTargets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("import_targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("import_targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *L2VPN) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", *m.Slug, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", *m.Slug, 100); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", *m.Slug, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *L2VPN) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *L2VPN) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *L2VPN) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this l2 v p n based on the context it is used
func (m *L2VPN) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImportTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
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

func (m *L2VPN) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) contextValidateExportTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportTargets); i++ {

		if m.ExportTargets[i] != nil {
			if err := m.ExportTargets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("export_targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("export_targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *L2VPN) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) contextValidateImportTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ImportTargets); i++ {

		if m.ImportTargets[i] != nil {
			if err := m.ImportTargets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("import_targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("import_targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *L2VPN) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *L2VPN) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *L2VPN) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {
		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *L2VPN) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *L2VPN) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *L2VPN) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *L2VPN) UnmarshalBinary(b []byte) error {
	var res L2VPN
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// L2VPNType Type
//
// swagger:model L2VPNType
type L2VPNType struct {

	// label
	// Required: true
	// Enum: [VPWS VPLS VXLAN VXLAN-EVPN MPLS EVPN PBB EVPN EPL EVPL Ethernet Private LAN Ethernet Virtual Private LAN Ethernet Private Tree Ethernet Virtual Private Tree]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [vpws vpls vxlan vxlan-evpn mpls-evpn pbb-evpn epl evpl ep-lan evp-lan ep-tree evp-tree]
	Value *string `json:"value"`
}

// Validate validates this l2 v p n type
func (m *L2VPNType) Validate(formats strfmt.Registry) error {
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

var l2VPNTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VPWS","VPLS","VXLAN","VXLAN-EVPN","MPLS EVPN","PBB EVPN","EPL","EVPL","Ethernet Private LAN","Ethernet Virtual Private LAN","Ethernet Private Tree","Ethernet Virtual Private Tree"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		l2VPNTypeTypeLabelPropEnum = append(l2VPNTypeTypeLabelPropEnum, v)
	}
}

const (

	// L2VPNTypeLabelVPWS captures enum value "VPWS"
	L2VPNTypeLabelVPWS string = "VPWS"

	// L2VPNTypeLabelVPLS captures enum value "VPLS"
	L2VPNTypeLabelVPLS string = "VPLS"

	// L2VPNTypeLabelVXLAN captures enum value "VXLAN"
	L2VPNTypeLabelVXLAN string = "VXLAN"

	// L2VPNTypeLabelVXLANDashEVPN captures enum value "VXLAN-EVPN"
	L2VPNTypeLabelVXLANDashEVPN string = "VXLAN-EVPN"

	// L2VPNTypeLabelMPLSEVPN captures enum value "MPLS EVPN"
	L2VPNTypeLabelMPLSEVPN string = "MPLS EVPN"

	// L2VPNTypeLabelPBBEVPN captures enum value "PBB EVPN"
	L2VPNTypeLabelPBBEVPN string = "PBB EVPN"

	// L2VPNTypeLabelEPL captures enum value "EPL"
	L2VPNTypeLabelEPL string = "EPL"

	// L2VPNTypeLabelEVPL captures enum value "EVPL"
	L2VPNTypeLabelEVPL string = "EVPL"

	// L2VPNTypeLabelEthernetPrivateLAN captures enum value "Ethernet Private LAN"
	L2VPNTypeLabelEthernetPrivateLAN string = "Ethernet Private LAN"

	// L2VPNTypeLabelEthernetVirtualPrivateLAN captures enum value "Ethernet Virtual Private LAN"
	L2VPNTypeLabelEthernetVirtualPrivateLAN string = "Ethernet Virtual Private LAN"

	// L2VPNTypeLabelEthernetPrivateTree captures enum value "Ethernet Private Tree"
	L2VPNTypeLabelEthernetPrivateTree string = "Ethernet Private Tree"

	// L2VPNTypeLabelEthernetVirtualPrivateTree captures enum value "Ethernet Virtual Private Tree"
	L2VPNTypeLabelEthernetVirtualPrivateTree string = "Ethernet Virtual Private Tree"
)

// prop value enum
func (m *L2VPNType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, l2VPNTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *L2VPNType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var l2VPNTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vpws","vpls","vxlan","vxlan-evpn","mpls-evpn","pbb-evpn","epl","evpl","ep-lan","evp-lan","ep-tree","evp-tree"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		l2VPNTypeTypeValuePropEnum = append(l2VPNTypeTypeValuePropEnum, v)
	}
}

const (

	// L2VPNTypeValueVpws captures enum value "vpws"
	L2VPNTypeValueVpws string = "vpws"

	// L2VPNTypeValueVpls captures enum value "vpls"
	L2VPNTypeValueVpls string = "vpls"

	// L2VPNTypeValueVxlan captures enum value "vxlan"
	L2VPNTypeValueVxlan string = "vxlan"

	// L2VPNTypeValueVxlanDashEvpn captures enum value "vxlan-evpn"
	L2VPNTypeValueVxlanDashEvpn string = "vxlan-evpn"

	// L2VPNTypeValueMplsDashEvpn captures enum value "mpls-evpn"
	L2VPNTypeValueMplsDashEvpn string = "mpls-evpn"

	// L2VPNTypeValuePbbDashEvpn captures enum value "pbb-evpn"
	L2VPNTypeValuePbbDashEvpn string = "pbb-evpn"

	// L2VPNTypeValueEpl captures enum value "epl"
	L2VPNTypeValueEpl string = "epl"

	// L2VPNTypeValueEvpl captures enum value "evpl"
	L2VPNTypeValueEvpl string = "evpl"

	// L2VPNTypeValueEpDashLan captures enum value "ep-lan"
	L2VPNTypeValueEpDashLan string = "ep-lan"

	// L2VPNTypeValueEvpDashLan captures enum value "evp-lan"
	L2VPNTypeValueEvpDashLan string = "evp-lan"

	// L2VPNTypeValueEpDashTree captures enum value "ep-tree"
	L2VPNTypeValueEpDashTree string = "ep-tree"

	// L2VPNTypeValueEvpDashTree captures enum value "evp-tree"
	L2VPNTypeValueEvpDashTree string = "evp-tree"
)

// prop value enum
func (m *L2VPNType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, l2VPNTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *L2VPNType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this l2 v p n type based on context it is used
func (m *L2VPNType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *L2VPNType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *L2VPNType) UnmarshalBinary(b []byte) error {
	var res L2VPNType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
