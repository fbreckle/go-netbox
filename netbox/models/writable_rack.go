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

// WritableRack writable rack
//
// swagger:model WritableRack
type WritableRack struct {

	// Asset tag
	//
	// A unique tag used to identify this rack
	// Max Length: 50
	AssetTag *string `json:"asset_tag"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Descending units
	//
	// Units are numbered top-to-bottom
	DescUnits bool `json:"desc_units"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Facility ID
	// Max Length: 50
	FacilityID *string `json:"facility_id"`

	// * `2-post-frame` - 2-post frame
	// * `4-post-frame` - 4-post frame
	// * `4-post-cabinet` - 4-post cabinet
	// * `wall-frame` - wall-mounted frame
	// * `wall-frame-vertical` - wall-mounted frame (vertical)
	// * `wall-cabinet` - wall-mounted cabinet
	// * `wall-cabinet-vertical` - wall-mounted cabinet (vertical)
	// Enum: ["2-post-frame","4-post-frame","4-post-cabinet","wall-frame","wall-frame-vertical","wall-cabinet","wall-cabinet-vertical",""]
	FormFactor string `json:"form_factor"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Location
	Location *int64 `json:"location"`

	// Max weight
	//
	// Maximum load capacity for the rack
	// Maximum: 2.147483647e+09
	// Minimum: 0
	MaxWeight *int64 `json:"max_weight"`

	// Mounting depth
	//
	// Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	// Maximum: 32767
	// Minimum: 0
	MountingDepth *int64 `json:"mounting_depth"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Outer depth
	//
	// Outer dimension of rack (depth)
	// Maximum: 32767
	// Minimum: 0
	OuterDepth *int64 `json:"outer_depth"`

	// Outer unit
	// Enum: ["mm","in"]
	OuterUnit string `json:"outer_unit"`

	// Outer width
	//
	// Outer dimension of rack (width)
	// Maximum: 32767
	// Minimum: 0
	OuterWidth *int64 `json:"outer_width"`

	// Powerfeed count
	// Read Only: true
	PowerfeedCount int64 `json:"powerfeed_count,omitempty"`

	// Role
	//
	// Functional role
	Role *int64 `json:"role"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// Site
	// Required: true
	Site *int64 `json:"site"`

	// Status
	// Enum: ["reserved","available","planned","active","deprecated"]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant"`

	// Height (U)
	//
	// Height in rack units
	// Maximum: 100
	// Minimum: 1
	UHeight int64 `json:"u_height,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Weight
	Weight *float64 `json:"weight"`

	// Weight unit
	// Enum: ["kg","g","lb","oz"]
	WeightUnit string `json:"weight_unit"`

	// Width
	//
	// Rail-to-rail width
	// Enum: [10,19,21,23]
	Width int64 `json:"width,omitempty"`
}

// Validate validates this writable rack
func (m *WritableRack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountingDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterWidth(formats); err != nil {
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

	if err := m.validateUHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRack) validateAssetTag(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", *m.AssetTag, 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateFacilityID(formats strfmt.Registry) error {
	if swag.IsZero(m.FacilityID) { // not required
		return nil
	}

	if err := validate.MaxLength("facility_id", "body", *m.FacilityID, 50); err != nil {
		return err
	}

	return nil
}

var writableRackTypeFormFactorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2-post-frame","4-post-frame","4-post-cabinet","wall-frame","wall-frame-vertical","wall-cabinet","wall-cabinet-vertical",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeFormFactorPropEnum = append(writableRackTypeFormFactorPropEnum, v)
	}
}

const (

	// WritableRackFormFactorNr2DashPostDashFrame captures enum value "2-post-frame"
	WritableRackFormFactorNr2DashPostDashFrame string = "2-post-frame"

	// WritableRackFormFactorNr4DashPostDashFrame captures enum value "4-post-frame"
	WritableRackFormFactorNr4DashPostDashFrame string = "4-post-frame"

	// WritableRackFormFactorNr4DashPostDashCabinet captures enum value "4-post-cabinet"
	WritableRackFormFactorNr4DashPostDashCabinet string = "4-post-cabinet"

	// WritableRackFormFactorWallDashFrame captures enum value "wall-frame"
	WritableRackFormFactorWallDashFrame string = "wall-frame"

	// WritableRackFormFactorWallDashFrameDashVertical captures enum value "wall-frame-vertical"
	WritableRackFormFactorWallDashFrameDashVertical string = "wall-frame-vertical"

	// WritableRackFormFactorWallDashCabinet captures enum value "wall-cabinet"
	WritableRackFormFactorWallDashCabinet string = "wall-cabinet"

	// WritableRackFormFactorWallDashCabinetDashVertical captures enum value "wall-cabinet-vertical"
	WritableRackFormFactorWallDashCabinetDashVertical string = "wall-cabinet-vertical"

	// WritableRackFormFactorEmpty captures enum value ""
	WritableRackFormFactorEmpty string = ""
)

// prop value enum
func (m *WritableRack) validateFormFactorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeFormFactorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateFormFactor(formats strfmt.Registry) error {
	if swag.IsZero(m.FormFactor) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormFactorEnum("form_factor", "body", m.FormFactor); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateMaxWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_weight", "body", *m.MaxWeight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_weight", "body", *m.MaxWeight, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateMountingDepth(formats strfmt.Registry) error {
	if swag.IsZero(m.MountingDepth) { // not required
		return nil
	}

	if err := validate.MinimumInt("mounting_depth", "body", *m.MountingDepth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mounting_depth", "body", *m.MountingDepth, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateName(formats strfmt.Registry) error {

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

func (m *WritableRack) validateOuterDepth(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterDepth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_depth", "body", *m.OuterDepth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_depth", "body", *m.OuterDepth, 32767, false); err != nil {
		return err
	}

	return nil
}

var writableRackTypeOuterUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mm","in"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeOuterUnitPropEnum = append(writableRackTypeOuterUnitPropEnum, v)
	}
}

const (

	// WritableRackOuterUnitMm captures enum value "mm"
	WritableRackOuterUnitMm string = "mm"

	// WritableRackOuterUnitIn captures enum value "in"
	WritableRackOuterUnitIn string = "in"
)

// prop value enum
func (m *WritableRack) validateOuterUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeOuterUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateOuterUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateOuterUnitEnum("outer_unit", "body", m.OuterUnit); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateOuterWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.OuterWidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_width", "body", *m.OuterWidth, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_width", "body", *m.OuterWidth, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateSerial(formats strfmt.Registry) error {
	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", m.Serial, 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	return nil
}

var writableRackTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["reserved","available","planned","active","deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeStatusPropEnum = append(writableRackTypeStatusPropEnum, v)
	}
}

const (

	// WritableRackStatusReserved captures enum value "reserved"
	WritableRackStatusReserved string = "reserved"

	// WritableRackStatusAvailable captures enum value "available"
	WritableRackStatusAvailable string = "available"

	// WritableRackStatusPlanned captures enum value "planned"
	WritableRackStatusPlanned string = "planned"

	// WritableRackStatusActive captures enum value "active"
	WritableRackStatusActive string = "active"

	// WritableRackStatusDeprecated captures enum value "deprecated"
	WritableRackStatusDeprecated string = "deprecated"
)

// prop value enum
func (m *WritableRack) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateTags(formats strfmt.Registry) error {
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

func (m *WritableRack) validateUHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", m.UHeight, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", m.UHeight, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

var writableRackTypeWeightUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kg","g","lb","oz"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeWeightUnitPropEnum = append(writableRackTypeWeightUnitPropEnum, v)
	}
}

const (

	// WritableRackWeightUnitKg captures enum value "kg"
	WritableRackWeightUnitKg string = "kg"

	// WritableRackWeightUnitG captures enum value "g"
	WritableRackWeightUnitG string = "g"

	// WritableRackWeightUnitLb captures enum value "lb"
	WritableRackWeightUnitLb string = "lb"

	// WritableRackWeightUnitOz captures enum value "oz"
	WritableRackWeightUnitOz string = "oz"
)

// prop value enum
func (m *WritableRack) validateWeightUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeWeightUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateWeightUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.WeightUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateWeightUnitEnum("weight_unit", "body", m.WeightUnit); err != nil {
		return err
	}

	return nil
}

var writableRackTypeWidthPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[10,19,21,23]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRackTypeWidthPropEnum = append(writableRackTypeWidthPropEnum, v)
	}
}

// prop value enum
func (m *WritableRack) validateWidthEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, writableRackTypeWidthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRack) validateWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.Width) { // not required
		return nil
	}

	// value enum
	if err := m.validateWidthEnum("width", "body", m.Width); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable rack based on the context it is used
func (m *WritableRack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceCount(ctx, formats); err != nil {
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

	if err := m.contextValidatePowerfeedCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *WritableRack) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) contextValidateDeviceCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "device_count", "body", int64(m.DeviceCount)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) contextValidatePowerfeedCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "powerfeed_count", "body", int64(m.PowerfeedCount)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRack) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

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

func (m *WritableRack) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRack) UnmarshalBinary(b []byte) error {
	var res WritableRack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
