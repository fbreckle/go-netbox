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

// checks if the AvailableASN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableASN{}

// AvailableASN Representation of an ASN which does not exist in the database.
type AvailableASN struct {
	Asn int32 `json:"asn"`
}

// NewAvailableASN instantiates a new AvailableASN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableASN(asn int32) *AvailableASN {
	this := AvailableASN{}
	this.Asn = asn
	return &this
}

// NewAvailableASNWithDefaults instantiates a new AvailableASN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableASNWithDefaults() *AvailableASN {
	this := AvailableASN{}
	return &this
}

// GetAsn returns the Asn field value
func (o *AvailableASN) GetAsn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *AvailableASN) GetAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *AvailableASN) SetAsn(v int32) {
	o.Asn = v
}

func (o AvailableASN) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableASN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: asn is readOnly
	return toSerialize, nil
}

type NullableAvailableASN struct {
	value *AvailableASN
	isSet bool
}

func (v NullableAvailableASN) Get() *AvailableASN {
	return v.value
}

func (v *NullableAvailableASN) Set(val *AvailableASN) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableASN) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableASN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableASN(val *AvailableASN) *NullableAvailableASN {
	return &NullableAvailableASN{value: val, isSet: true}
}

func (v NullableAvailableASN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableASN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
