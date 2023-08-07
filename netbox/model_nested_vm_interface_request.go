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

// checks if the NestedVMInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedVMInterfaceRequest{}

// NestedVMInterfaceRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedVMInterfaceRequest struct {
	Name string `json:"name"`
}

// NewNestedVMInterfaceRequest instantiates a new NestedVMInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedVMInterfaceRequest(name string) *NestedVMInterfaceRequest {
	this := NestedVMInterfaceRequest{}
	this.Name = name
	return &this
}

// NewNestedVMInterfaceRequestWithDefaults instantiates a new NestedVMInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedVMInterfaceRequestWithDefaults() *NestedVMInterfaceRequest {
	this := NestedVMInterfaceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedVMInterfaceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedVMInterfaceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedVMInterfaceRequest) SetName(v string) {
	o.Name = v
}

func (o NestedVMInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedVMInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableNestedVMInterfaceRequest struct {
	value *NestedVMInterfaceRequest
	isSet bool
}

func (v NullableNestedVMInterfaceRequest) Get() *NestedVMInterfaceRequest {
	return v.value
}

func (v *NullableNestedVMInterfaceRequest) Set(val *NestedVMInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedVMInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedVMInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedVMInterfaceRequest(val *NestedVMInterfaceRequest) *NullableNestedVMInterfaceRequest {
	return &NullableNestedVMInterfaceRequest{value: val, isSet: true}
}

func (v NullableNestedVMInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedVMInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
