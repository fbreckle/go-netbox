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

// checks if the ASNRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ASNRequest{}

// ASNRequest Adds support for custom fields and tags.
type ASNRequest struct {
	// 16- or 32-bit autonomous system number
	Asn          int64                       `json:"asn"`
	Rir          NullableNestedRIRRequest    `json:"rir,omitempty"`
	Tenant       NullableNestedTenantRequest `json:"tenant,omitempty"`
	Description  *string                     `json:"description,omitempty"`
	Comments     *string                     `json:"comments,omitempty"`
	Tags         []NestedTagRequest          `json:"tags,omitempty"`
	CustomFields map[string]interface{}      `json:"custom_fields,omitempty"`
}

// NewASNRequest instantiates a new ASNRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewASNRequest(asn int64) *ASNRequest {
	this := ASNRequest{}
	this.Asn = asn
	return &this
}

// NewASNRequestWithDefaults instantiates a new ASNRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewASNRequestWithDefaults() *ASNRequest {
	this := ASNRequest{}
	return &this
}

// GetAsn returns the Asn field value
func (o *ASNRequest) GetAsn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *ASNRequest) GetAsnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *ASNRequest) SetAsn(v int64) {
	o.Asn = v
}

// GetRir returns the Rir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ASNRequest) GetRir() NestedRIRRequest {
	if o == nil || IsNil(o.Rir.Get()) {
		var ret NestedRIRRequest
		return ret
	}
	return *o.Rir.Get()
}

// GetRirOk returns a tuple with the Rir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASNRequest) GetRirOk() (*NestedRIRRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rir.Get(), o.Rir.IsSet()
}

// HasRir returns a boolean if a field has been set.
func (o *ASNRequest) HasRir() bool {
	if o != nil && o.Rir.IsSet() {
		return true
	}

	return false
}

// SetRir gets a reference to the given NullableNestedRIRRequest and assigns it to the Rir field.
func (o *ASNRequest) SetRir(v NestedRIRRequest) {
	o.Rir.Set(&v)
}

// SetRirNil sets the value for Rir to be an explicit nil
func (o *ASNRequest) SetRirNil() {
	o.Rir.Set(nil)
}

// UnsetRir ensures that no value is present for Rir, not even an explicit nil
func (o *ASNRequest) UnsetRir() {
	o.Rir.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ASNRequest) GetTenant() NestedTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret NestedTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASNRequest) GetTenantOk() (*NestedTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *ASNRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenantRequest and assigns it to the Tenant field.
func (o *ASNRequest) SetTenant(v NestedTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *ASNRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *ASNRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ASNRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASNRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ASNRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ASNRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ASNRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASNRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ASNRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ASNRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ASNRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASNRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ASNRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ASNRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ASNRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASNRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ASNRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ASNRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ASNRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ASNRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asn"] = o.Asn
	if o.Rir.IsSet() {
		toSerialize["rir"] = o.Rir.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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

type NullableASNRequest struct {
	value *ASNRequest
	isSet bool
}

func (v NullableASNRequest) Get() *ASNRequest {
	return v.value
}

func (v *NullableASNRequest) Set(val *ASNRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableASNRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableASNRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableASNRequest(val *ASNRequest) *NullableASNRequest {
	return &NullableASNRequest{value: val, isSet: true}
}

func (v NullableASNRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableASNRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
