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

// checks if the WritableVirtualDeviceContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableVirtualDeviceContextRequest{}

// WritableVirtualDeviceContextRequest Adds support for custom fields and tags.
type WritableVirtualDeviceContextRequest struct {
	Name   string        `json:"name"`
	Device NullableInt32 `json:"device,omitempty"`
	// Numeric identifier unique to the parent device
	Identifier NullableInt32 `json:"identifier,omitempty"`
	Tenant     NullableInt32 `json:"tenant,omitempty"`
	PrimaryIp4 NullableInt32 `json:"primary_ip4,omitempty"`
	PrimaryIp6 NullableInt32 `json:"primary_ip6,omitempty"`
	// * `active` - Active * `planned` - Planned * `offline` - Offline
	Status       string                 `json:"status"`
	Description  *string                `json:"description,omitempty"`
	Comments     *string                `json:"comments,omitempty"`
	Tags         []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewWritableVirtualDeviceContextRequest instantiates a new WritableVirtualDeviceContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVirtualDeviceContextRequest(name string, status string) *WritableVirtualDeviceContextRequest {
	this := WritableVirtualDeviceContextRequest{}
	this.Name = name
	this.Status = status
	return &this
}

// NewWritableVirtualDeviceContextRequestWithDefaults instantiates a new WritableVirtualDeviceContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVirtualDeviceContextRequestWithDefaults() *WritableVirtualDeviceContextRequest {
	this := WritableVirtualDeviceContextRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableVirtualDeviceContextRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualDeviceContextRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableVirtualDeviceContextRequest) SetName(v string) {
	o.Name = v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualDeviceContextRequest) GetDevice() int32 {
	if o == nil || IsNil(o.Device.Get()) {
		var ret int32
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualDeviceContextRequest) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableInt32 and assigns it to the Device field.
func (o *WritableVirtualDeviceContextRequest) SetDevice(v int32) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *WritableVirtualDeviceContextRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *WritableVirtualDeviceContextRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualDeviceContextRequest) GetIdentifier() int32 {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret int32
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualDeviceContextRequest) GetIdentifierOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableInt32 and assigns it to the Identifier field.
func (o *WritableVirtualDeviceContextRequest) SetIdentifier(v int32) {
	o.Identifier.Set(&v)
}

// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *WritableVirtualDeviceContextRequest) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *WritableVirtualDeviceContextRequest) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualDeviceContextRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualDeviceContextRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *WritableVirtualDeviceContextRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableVirtualDeviceContextRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableVirtualDeviceContextRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualDeviceContextRequest) GetPrimaryIp4() int32 {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret int32
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualDeviceContextRequest) GetPrimaryIp4Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableInt32 and assigns it to the PrimaryIp4 field.
func (o *WritableVirtualDeviceContextRequest) SetPrimaryIp4(v int32) {
	o.PrimaryIp4.Set(&v)
}

// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *WritableVirtualDeviceContextRequest) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *WritableVirtualDeviceContextRequest) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualDeviceContextRequest) GetPrimaryIp6() int32 {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret int32
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualDeviceContextRequest) GetPrimaryIp6Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableInt32 and assigns it to the PrimaryIp6 field.
func (o *WritableVirtualDeviceContextRequest) SetPrimaryIp6(v int32) {
	o.PrimaryIp6.Set(&v)
}

// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *WritableVirtualDeviceContextRequest) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *WritableVirtualDeviceContextRequest) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetStatus returns the Status field value
func (o *WritableVirtualDeviceContextRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualDeviceContextRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WritableVirtualDeviceContextRequest) SetStatus(v string) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableVirtualDeviceContextRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualDeviceContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableVirtualDeviceContextRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableVirtualDeviceContextRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualDeviceContextRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableVirtualDeviceContextRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVirtualDeviceContextRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualDeviceContextRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableVirtualDeviceContextRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVirtualDeviceContextRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualDeviceContextRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVirtualDeviceContextRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableVirtualDeviceContextRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableVirtualDeviceContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableVirtualDeviceContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	toSerialize["status"] = o.Status
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

type NullableWritableVirtualDeviceContextRequest struct {
	value *WritableVirtualDeviceContextRequest
	isSet bool
}

func (v NullableWritableVirtualDeviceContextRequest) Get() *WritableVirtualDeviceContextRequest {
	return v.value
}

func (v *NullableWritableVirtualDeviceContextRequest) Set(val *WritableVirtualDeviceContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVirtualDeviceContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVirtualDeviceContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVirtualDeviceContextRequest(val *WritableVirtualDeviceContextRequest) *NullableWritableVirtualDeviceContextRequest {
	return &NullableWritableVirtualDeviceContextRequest{value: val, isSet: true}
}

func (v NullableWritableVirtualDeviceContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVirtualDeviceContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
