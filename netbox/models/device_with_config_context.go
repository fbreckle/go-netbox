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

// DeviceWithConfigContext device with config context
//
// swagger:model DeviceWithConfigContext
type DeviceWithConfigContext struct {

	// Asset tag
	//
	// A unique tag used to identify this device
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// cluster
	Cluster *NestedCluster `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext map[string]interface{} `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// device role
	// Required: true
	DeviceRole *NestedDeviceRole `json:"device_role"`

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// face
	Face *DeviceWithConfigContextFace `json:"face,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData map[string]interface{} `json:"local_context_data,omitempty"`

	// location
	Location *NestedLocation `json:"location,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	Name *string `json:"name"`

	// parent device
	ParentDevice *NestedDevice `json:"parent_device,omitempty"`

	// platform
	Platform *NestedPlatform `json:"platform,omitempty"`

	// Position (U)
	// Minimum: 1
	Position *int64 `json:"position,omitempty"`

	// primary ip
	PrimaryIP *NestedIPAddress `json:"primary_ip,omitempty"`

	// primary ip4
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty"`

	// primary ip6
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty"`

	// rack
	Rack *NestedRack `json:"rack,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// site
	// Required: true
	Site *NestedSite `json:"site"`

	// status
	Status *DeviceWithConfigContextStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Vc position
	// Maximum: 255
	// Minimum: 0
	VcPosition *int64 `json:"vc_position,omitempty"`

	// Vc priority
	// Maximum: 255
	// Minimum: 0
	VcPriority *int64 `json:"vc_priority,omitempty"`

	// virtual chassis
	VirtualChassis *NestedVirtualChassis `json:"virtual_chassis,omitempty"`
}

// Validate validates this device with config context
func (m *DeviceWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualChassis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceWithConfigContext) validateAssetTag(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", *m.AssetTag, 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateDeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("device_role", "body", m.DeviceRole); err != nil {
		return err
	}

	if m.DeviceRole != nil {
		if err := m.DeviceRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateFace(formats strfmt.Registry) error {
	if swag.IsZero(m.Face) { // not required
		return nil
	}

	if m.Face != nil {
		if err := m.Face.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateParentDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentDevice) { // not required
		return nil
	}

	if m.ParentDevice != nil {
		if err := m.ParentDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", *m.Position, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIP(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIP) { // not required
		return nil
	}

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIp4(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp4) { // not required
		return nil
	}

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validatePrimaryIp6(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp6) { // not required
		return nil
	}

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateRack(formats strfmt.Registry) error {
	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if m.Rack != nil {
		if err := m.Rack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateSerial(formats strfmt.Registry) error {
	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", m.Serial, 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateTags(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceWithConfigContext) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVcPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.VcPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_position", "body", *m.VcPosition, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_position", "body", *m.VcPosition, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVcPriority(formats strfmt.Registry) error {
	if swag.IsZero(m.VcPriority) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_priority", "body", *m.VcPriority, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_priority", "body", *m.VcPriority, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) validateVirtualChassis(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualChassis) { // not required
		return nil
	}

	if m.VirtualChassis != nil {
		if err := m.VirtualChassis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_chassis")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device with config context based on the context it is used
func (m *DeviceWithConfigContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIp4(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIp6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualChassis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceWithConfigContext) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateConfigContext(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *DeviceWithConfigContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateDeviceRole(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceRole != nil {
		if err := m.DeviceRole.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceType != nil {
		if err := m.DeviceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateFace(ctx context.Context, formats strfmt.Registry) error {

	if m.Face != nil {
		if err := m.Face.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateParentDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentDevice != nil {
		if err := m.ParentDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if m.Platform != nil {
		if err := m.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidatePrimaryIP(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidatePrimaryIp4(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidatePrimaryIp6(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateRack(ctx context.Context, formats strfmt.Registry) error {

	if m.Rack != nil {
		if err := m.Rack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {
		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {
		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *DeviceWithConfigContext) contextValidateVirtualChassis(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualChassis != nil {
		if err := m.VirtualChassis.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_chassis")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContext) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceWithConfigContextFace Face
//
// swagger:model DeviceWithConfigContextFace
type DeviceWithConfigContextFace struct {

	// label
	// Required: true
	// Enum: [Front Rear]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [front rear]
	Value *string `json:"value"`
}

// Validate validates this device with config context face
func (m *DeviceWithConfigContextFace) Validate(formats strfmt.Registry) error {
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

var deviceWithConfigContextFaceTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Front","Rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextFaceTypeLabelPropEnum = append(deviceWithConfigContextFaceTypeLabelPropEnum, v)
	}
}

const (

	// DeviceWithConfigContextFaceLabelFront captures enum value "Front"
	DeviceWithConfigContextFaceLabelFront string = "Front"

	// DeviceWithConfigContextFaceLabelRear captures enum value "Rear"
	DeviceWithConfigContextFaceLabelRear string = "Rear"
)

// prop value enum
func (m *DeviceWithConfigContextFace) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceWithConfigContextFaceTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContextFace) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("face"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var deviceWithConfigContextFaceTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front","rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextFaceTypeValuePropEnum = append(deviceWithConfigContextFaceTypeValuePropEnum, v)
	}
}

const (

	// DeviceWithConfigContextFaceValueFront captures enum value "front"
	DeviceWithConfigContextFaceValueFront string = "front"

	// DeviceWithConfigContextFaceValueRear captures enum value "rear"
	DeviceWithConfigContextFaceValueRear string = "rear"
)

// prop value enum
func (m *DeviceWithConfigContextFace) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceWithConfigContextFaceTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContextFace) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("face"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device with config context face based on context it is used
func (m *DeviceWithConfigContextFace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContextFace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContextFace) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContextFace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceWithConfigContextStatus Status
//
// swagger:model DeviceWithConfigContextStatus
type DeviceWithConfigContextStatus struct {

	// label
	// Required: true
	// Enum: [Offline Active Planned Staged Failed Inventory Decommissioning]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [offline active planned staged failed inventory decommissioning]
	Value *string `json:"value"`
}

// Validate validates this device with config context status
func (m *DeviceWithConfigContextStatus) Validate(formats strfmt.Registry) error {
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

var deviceWithConfigContextStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offline","Active","Planned","Staged","Failed","Inventory","Decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextStatusTypeLabelPropEnum = append(deviceWithConfigContextStatusTypeLabelPropEnum, v)
	}
}

const (

	// DeviceWithConfigContextStatusLabelOffline captures enum value "Offline"
	DeviceWithConfigContextStatusLabelOffline string = "Offline"

	// DeviceWithConfigContextStatusLabelActive captures enum value "Active"
	DeviceWithConfigContextStatusLabelActive string = "Active"

	// DeviceWithConfigContextStatusLabelPlanned captures enum value "Planned"
	DeviceWithConfigContextStatusLabelPlanned string = "Planned"

	// DeviceWithConfigContextStatusLabelStaged captures enum value "Staged"
	DeviceWithConfigContextStatusLabelStaged string = "Staged"

	// DeviceWithConfigContextStatusLabelFailed captures enum value "Failed"
	DeviceWithConfigContextStatusLabelFailed string = "Failed"

	// DeviceWithConfigContextStatusLabelInventory captures enum value "Inventory"
	DeviceWithConfigContextStatusLabelInventory string = "Inventory"

	// DeviceWithConfigContextStatusLabelDecommissioning captures enum value "Decommissioning"
	DeviceWithConfigContextStatusLabelDecommissioning string = "Decommissioning"
)

// prop value enum
func (m *DeviceWithConfigContextStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceWithConfigContextStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContextStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var deviceWithConfigContextStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","staged","failed","inventory","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextStatusTypeValuePropEnum = append(deviceWithConfigContextStatusTypeValuePropEnum, v)
	}
}

const (

	// DeviceWithConfigContextStatusValueOffline captures enum value "offline"
	DeviceWithConfigContextStatusValueOffline string = "offline"

	// DeviceWithConfigContextStatusValueActive captures enum value "active"
	DeviceWithConfigContextStatusValueActive string = "active"

	// DeviceWithConfigContextStatusValuePlanned captures enum value "planned"
	DeviceWithConfigContextStatusValuePlanned string = "planned"

	// DeviceWithConfigContextStatusValueStaged captures enum value "staged"
	DeviceWithConfigContextStatusValueStaged string = "staged"

	// DeviceWithConfigContextStatusValueFailed captures enum value "failed"
	DeviceWithConfigContextStatusValueFailed string = "failed"

	// DeviceWithConfigContextStatusValueInventory captures enum value "inventory"
	DeviceWithConfigContextStatusValueInventory string = "inventory"

	// DeviceWithConfigContextStatusValueDecommissioning captures enum value "decommissioning"
	DeviceWithConfigContextStatusValueDecommissioning string = "decommissioning"
)

// prop value enum
func (m *DeviceWithConfigContextStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceWithConfigContextStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContextStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device with config context status based on context it is used
func (m *DeviceWithConfigContextStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContextStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContextStatus) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContextStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
