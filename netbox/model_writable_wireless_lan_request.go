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

// checks if the WritableWirelessLANRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableWirelessLANRequest{}

// WritableWirelessLANRequest Adds support for custom fields and tags.
type WritableWirelessLANRequest struct {
	Ssid        string        `json:"ssid"`
	Description *string       `json:"description,omitempty"`
	Group       NullableInt32 `json:"group,omitempty"`
	// * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated
	Status *string       `json:"status,omitempty"`
	Vlan   NullableInt32 `json:"vlan,omitempty"`
	Tenant NullableInt32 `json:"tenant,omitempty"`
	// * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise
	AuthType *string `json:"auth_type,omitempty"`
	// * `auto` - Auto * `tkip` - TKIP * `aes` - AES
	AuthCipher   *string                `json:"auth_cipher,omitempty"`
	AuthPsk      *string                `json:"auth_psk,omitempty"`
	Comments     *string                `json:"comments,omitempty"`
	Tags         []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewWritableWirelessLANRequest instantiates a new WritableWirelessLANRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableWirelessLANRequest(ssid string) *WritableWirelessLANRequest {
	this := WritableWirelessLANRequest{}
	this.Ssid = ssid
	return &this
}

// NewWritableWirelessLANRequestWithDefaults instantiates a new WritableWirelessLANRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableWirelessLANRequestWithDefaults() *WritableWirelessLANRequest {
	this := WritableWirelessLANRequest{}
	return &this
}

// GetSsid returns the Ssid field value
func (o *WritableWirelessLANRequest) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *WritableWirelessLANRequest) SetSsid(v string) {
	o.Ssid = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableWirelessLANRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableWirelessLANRequest) GetGroup() int32 {
	if o == nil || IsNil(o.Group.Get()) {
		var ret int32
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableWirelessLANRequest) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableInt32 and assigns it to the Group field.
func (o *WritableWirelessLANRequest) SetGroup(v int32) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *WritableWirelessLANRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *WritableWirelessLANRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WritableWirelessLANRequest) SetStatus(v string) {
	o.Status = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableWirelessLANRequest) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan.Get()) {
		var ret int32
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableWirelessLANRequest) GetVlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableInt32 and assigns it to the Vlan field.
func (o *WritableWirelessLANRequest) SetVlan(v int32) {
	o.Vlan.Set(&v)
}

// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *WritableWirelessLANRequest) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *WritableWirelessLANRequest) UnsetVlan() {
	o.Vlan.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableWirelessLANRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableWirelessLANRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *WritableWirelessLANRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableWirelessLANRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableWirelessLANRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *WritableWirelessLANRequest) SetAuthType(v string) {
	o.AuthType = &v
}

// GetAuthCipher returns the AuthCipher field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetAuthCipher() string {
	if o == nil || IsNil(o.AuthCipher) {
		var ret string
		return ret
	}
	return *o.AuthCipher
}

// GetAuthCipherOk returns a tuple with the AuthCipher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetAuthCipherOk() (*string, bool) {
	if o == nil || IsNil(o.AuthCipher) {
		return nil, false
	}
	return o.AuthCipher, true
}

// HasAuthCipher returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasAuthCipher() bool {
	if o != nil && !IsNil(o.AuthCipher) {
		return true
	}

	return false
}

// SetAuthCipher gets a reference to the given string and assigns it to the AuthCipher field.
func (o *WritableWirelessLANRequest) SetAuthCipher(v string) {
	o.AuthCipher = &v
}

// GetAuthPsk returns the AuthPsk field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetAuthPsk() string {
	if o == nil || IsNil(o.AuthPsk) {
		var ret string
		return ret
	}
	return *o.AuthPsk
}

// GetAuthPskOk returns a tuple with the AuthPsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetAuthPskOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPsk) {
		return nil, false
	}
	return o.AuthPsk, true
}

// HasAuthPsk returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasAuthPsk() bool {
	if o != nil && !IsNil(o.AuthPsk) {
		return true
	}

	return false
}

// SetAuthPsk gets a reference to the given string and assigns it to the AuthPsk field.
func (o *WritableWirelessLANRequest) SetAuthPsk(v string) {
	o.AuthPsk = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableWirelessLANRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableWirelessLANRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableWirelessLANRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableWirelessLANRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableWirelessLANRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableWirelessLANRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableWirelessLANRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableWirelessLANRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ssid"] = o.Ssid
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.AuthCipher) {
		toSerialize["auth_cipher"] = o.AuthCipher
	}
	if !IsNil(o.AuthPsk) {
		toSerialize["auth_psk"] = o.AuthPsk
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

type NullableWritableWirelessLANRequest struct {
	value *WritableWirelessLANRequest
	isSet bool
}

func (v NullableWritableWirelessLANRequest) Get() *WritableWirelessLANRequest {
	return v.value
}

func (v *NullableWritableWirelessLANRequest) Set(val *WritableWirelessLANRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableWirelessLANRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableWirelessLANRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableWirelessLANRequest(val *WritableWirelessLANRequest) *NullableWritableWirelessLANRequest {
	return &NullableWritableWirelessLANRequest{value: val, isSet: true}
}

func (v NullableWritableWirelessLANRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableWirelessLANRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
