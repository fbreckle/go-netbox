/*
NetBox API

API to access NetBox

API version: 3.5.0 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableCableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableCableRequest{}

// PatchedWritableCableRequest Adds support for custom fields and tags.
type PatchedWritableCableRequest struct {
	// * `cat3` - CAT3 * `cat5` - CAT5 * `cat5e` - CAT5e * `cat6` - CAT6 * `cat6a` - CAT6a * `cat7` - CAT7 * `cat7a` - CAT7a * `cat8` - CAT8 * `dac-active` - Direct Attach Copper (Active) * `dac-passive` - Direct Attach Copper (Passive) * `mrj21-trunk` - MRJ21 Trunk * `coaxial` - Coaxial * `mmf` - Multimode Fiber * `mmf-om1` - Multimode Fiber (OM1) * `mmf-om2` - Multimode Fiber (OM2) * `mmf-om3` - Multimode Fiber (OM3) * `mmf-om4` - Multimode Fiber (OM4) * `mmf-om5` - Multimode Fiber (OM5) * `smf` - Singlemode Fiber * `smf-os1` - Singlemode Fiber (OS1) * `smf-os2` - Singlemode Fiber (OS2) * `aoc` - Active Optical Cabling (AOC) * `power` - Power
	Type          *string                `json:"type,omitempty"`
	ATerminations []GenericObjectRequest `json:"a_terminations,omitempty"`
	BTerminations []GenericObjectRequest `json:"b_terminations,omitempty"`
	// * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning
	Status *string         `json:"status,omitempty"`
	Tenant NullableInt32   `json:"tenant,omitempty"`
	Label  *string         `json:"label,omitempty"`
	Color  *string         `json:"color,omitempty"`
	Length NullableFloat64 `json:"length,omitempty"`
	// * `km` - Kilometers * `m` - Meters * `cm` - Centimeters * `mi` - Miles * `ft` - Feet * `in` - Inches
	LengthUnit   *string                `json:"length_unit,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Comments     *string                `json:"comments,omitempty"`
	Tags         []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewPatchedWritableCableRequest instantiates a new PatchedWritableCableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCableRequest() *PatchedWritableCableRequest {
	this := PatchedWritableCableRequest{}
	return &this
}

// NewPatchedWritableCableRequestWithDefaults instantiates a new PatchedWritableCableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCableRequestWithDefaults() *PatchedWritableCableRequest {
	this := PatchedWritableCableRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedWritableCableRequest) SetType(v string) {
	o.Type = &v
}

// GetATerminations returns the ATerminations field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetATerminations() []GenericObjectRequest {
	if o == nil || IsNil(o.ATerminations) {
		var ret []GenericObjectRequest
		return ret
	}
	return o.ATerminations
}

// GetATerminationsOk returns a tuple with the ATerminations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetATerminationsOk() ([]GenericObjectRequest, bool) {
	if o == nil || IsNil(o.ATerminations) {
		return nil, false
	}
	return o.ATerminations, true
}

// HasATerminations returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasATerminations() bool {
	if o != nil && !IsNil(o.ATerminations) {
		return true
	}

	return false
}

// SetATerminations gets a reference to the given []GenericObjectRequest and assigns it to the ATerminations field.
func (o *PatchedWritableCableRequest) SetATerminations(v []GenericObjectRequest) {
	o.ATerminations = v
}

// GetBTerminations returns the BTerminations field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetBTerminations() []GenericObjectRequest {
	if o == nil || IsNil(o.BTerminations) {
		var ret []GenericObjectRequest
		return ret
	}
	return o.BTerminations
}

// GetBTerminationsOk returns a tuple with the BTerminations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetBTerminationsOk() ([]GenericObjectRequest, bool) {
	if o == nil || IsNil(o.BTerminations) {
		return nil, false
	}
	return o.BTerminations, true
}

// HasBTerminations returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasBTerminations() bool {
	if o != nil && !IsNil(o.BTerminations) {
		return true
	}

	return false
}

// SetBTerminations gets a reference to the given []GenericObjectRequest and assigns it to the BTerminations field.
func (o *PatchedWritableCableRequest) SetBTerminations(v []GenericObjectRequest) {
	o.BTerminations = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchedWritableCableRequest) SetStatus(v string) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCableRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCableRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *PatchedWritableCableRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableCableRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableCableRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableCableRequest) SetLabel(v string) {
	o.Label = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *PatchedWritableCableRequest) SetColor(v string) {
	o.Color = &v
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCableRequest) GetLength() float64 {
	if o == nil || IsNil(o.Length.Get()) {
		var ret float64
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCableRequest) GetLengthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasLength() bool {
	if o != nil && o.Length.IsSet() {
		return true
	}

	return false
}

// SetLength gets a reference to the given NullableFloat64 and assigns it to the Length field.
func (o *PatchedWritableCableRequest) SetLength(v float64) {
	o.Length.Set(&v)
}

// SetLengthNil sets the value for Length to be an explicit nil
func (o *PatchedWritableCableRequest) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil
func (o *PatchedWritableCableRequest) UnsetLength() {
	o.Length.Unset()
}

// GetLengthUnit returns the LengthUnit field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetLengthUnit() string {
	if o == nil || IsNil(o.LengthUnit) {
		var ret string
		return ret
	}
	return *o.LengthUnit
}

// GetLengthUnitOk returns a tuple with the LengthUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetLengthUnitOk() (*string, bool) {
	if o == nil || IsNil(o.LengthUnit) {
		return nil, false
	}
	return o.LengthUnit, true
}

// HasLengthUnit returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasLengthUnit() bool {
	if o != nil && !IsNil(o.LengthUnit) {
		return true
	}

	return false
}

// SetLengthUnit gets a reference to the given string and assigns it to the LengthUnit field.
func (o *PatchedWritableCableRequest) SetLengthUnit(v string) {
	o.LengthUnit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableCableRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableCableRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableCableRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableCableRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCableRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableCableRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableCableRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableCableRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableCableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ATerminations) {
		toSerialize["a_terminations"] = o.ATerminations
	}
	if !IsNil(o.BTerminations) {
		toSerialize["b_terminations"] = o.BTerminations
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if o.Length.IsSet() {
		toSerialize["length"] = o.Length.Get()
	}
	if !IsNil(o.LengthUnit) {
		toSerialize["length_unit"] = o.LengthUnit
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullablePatchedWritableCableRequest struct {
	value *PatchedWritableCableRequest
	isSet bool
}

func (v NullablePatchedWritableCableRequest) Get() *PatchedWritableCableRequest {
	return v.value
}

func (v *NullablePatchedWritableCableRequest) Set(val *PatchedWritableCableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCableRequest(val *PatchedWritableCableRequest) *NullablePatchedWritableCableRequest {
	return &NullablePatchedWritableCableRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableCableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
