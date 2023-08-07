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

// checks if the NestedSiteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedSiteRequest{}

// NestedSiteRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedSiteRequest struct {
	// Full name of the site
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewNestedSiteRequest instantiates a new NestedSiteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedSiteRequest(name string, slug string) *NestedSiteRequest {
	this := NestedSiteRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedSiteRequestWithDefaults instantiates a new NestedSiteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedSiteRequestWithDefaults() *NestedSiteRequest {
	this := NestedSiteRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedSiteRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedSiteRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedSiteRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedSiteRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedSiteRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedSiteRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedSiteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedSiteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

type NullableNestedSiteRequest struct {
	value *NestedSiteRequest
	isSet bool
}

func (v NullableNestedSiteRequest) Get() *NestedSiteRequest {
	return v.value
}

func (v *NullableNestedSiteRequest) Set(val *NestedSiteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedSiteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedSiteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedSiteRequest(val *NestedSiteRequest) *NullableNestedSiteRequest {
	return &NullableNestedSiteRequest{value: val, isSet: true}
}

func (v NullableNestedSiteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedSiteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
