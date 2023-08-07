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

// checks if the NestedContactRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedContactRole{}

// NestedContactRole Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedContactRole struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
}

// NewNestedContactRole instantiates a new NestedContactRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedContactRole(id int32, url string, display string, name string, slug string) *NestedContactRole {
	this := NestedContactRole{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedContactRoleWithDefaults instantiates a new NestedContactRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedContactRoleWithDefaults() *NestedContactRole {
	this := NestedContactRole{}
	return &this
}

// GetId returns the Id field value
func (o *NestedContactRole) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedContactRole) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedContactRole) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedContactRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedContactRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedContactRole) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedContactRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedContactRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedContactRole) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedContactRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedContactRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedContactRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedContactRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedContactRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedContactRole) SetSlug(v string) {
	o.Slug = v
}

func (o NestedContactRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedContactRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

type NullableNestedContactRole struct {
	value *NestedContactRole
	isSet bool
}

func (v NullableNestedContactRole) Get() *NestedContactRole {
	return v.value
}

func (v *NullableNestedContactRole) Set(val *NestedContactRole) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedContactRole) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedContactRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedContactRole(val *NestedContactRole) *NullableNestedContactRole {
	return &NullableNestedContactRole{value: val, isSet: true}
}

func (v NullableNestedContactRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedContactRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
