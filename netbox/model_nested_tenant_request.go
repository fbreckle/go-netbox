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

// checks if the NestedTenantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedTenantRequest{}

// NestedTenantRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedTenantRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewNestedTenantRequest instantiates a new NestedTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedTenantRequest(name string, slug string) *NestedTenantRequest {
	this := NestedTenantRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedTenantRequestWithDefaults instantiates a new NestedTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedTenantRequestWithDefaults() *NestedTenantRequest {
	this := NestedTenantRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedTenantRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedTenantRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedTenantRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedTenantRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedTenantRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedTenantRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedTenantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

type NullableNestedTenantRequest struct {
	value *NestedTenantRequest
	isSet bool
}

func (v NullableNestedTenantRequest) Get() *NestedTenantRequest {
	return v.value
}

func (v *NullableNestedTenantRequest) Set(val *NestedTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedTenantRequest(val *NestedTenantRequest) *NullableNestedTenantRequest {
	return &NullableNestedTenantRequest{value: val, isSet: true}
}

func (v NullableNestedTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
