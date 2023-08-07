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

// checks if the PatchedWritableL2VPNTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableL2VPNTerminationRequest{}

// PatchedWritableL2VPNTerminationRequest Adds support for custom fields and tags.
type PatchedWritableL2VPNTerminationRequest struct {
	L2vpn              *int32                 `json:"l2vpn,omitempty"`
	AssignedObjectType *string                `json:"assigned_object_type,omitempty"`
	AssignedObjectId   *int64                 `json:"assigned_object_id,omitempty"`
	Tags               []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields       map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewPatchedWritableL2VPNTerminationRequest instantiates a new PatchedWritableL2VPNTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableL2VPNTerminationRequest() *PatchedWritableL2VPNTerminationRequest {
	this := PatchedWritableL2VPNTerminationRequest{}
	return &this
}

// NewPatchedWritableL2VPNTerminationRequestWithDefaults instantiates a new PatchedWritableL2VPNTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableL2VPNTerminationRequestWithDefaults() *PatchedWritableL2VPNTerminationRequest {
	this := PatchedWritableL2VPNTerminationRequest{}
	return &this
}

// GetL2vpn returns the L2vpn field value if set, zero value otherwise.
func (o *PatchedWritableL2VPNTerminationRequest) GetL2vpn() int32 {
	if o == nil || IsNil(o.L2vpn) {
		var ret int32
		return ret
	}
	return *o.L2vpn
}

// GetL2vpnOk returns a tuple with the L2vpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableL2VPNTerminationRequest) GetL2vpnOk() (*int32, bool) {
	if o == nil || IsNil(o.L2vpn) {
		return nil, false
	}
	return o.L2vpn, true
}

// HasL2vpn returns a boolean if a field has been set.
func (o *PatchedWritableL2VPNTerminationRequest) HasL2vpn() bool {
	if o != nil && !IsNil(o.L2vpn) {
		return true
	}

	return false
}

// SetL2vpn gets a reference to the given int32 and assigns it to the L2vpn field.
func (o *PatchedWritableL2VPNTerminationRequest) SetL2vpn(v int32) {
	o.L2vpn = &v
}

// GetAssignedObjectType returns the AssignedObjectType field value if set, zero value otherwise.
func (o *PatchedWritableL2VPNTerminationRequest) GetAssignedObjectType() string {
	if o == nil || IsNil(o.AssignedObjectType) {
		var ret string
		return ret
	}
	return *o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableL2VPNTerminationRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedObjectType) {
		return nil, false
	}
	return o.AssignedObjectType, true
}

// HasAssignedObjectType returns a boolean if a field has been set.
func (o *PatchedWritableL2VPNTerminationRequest) HasAssignedObjectType() bool {
	if o != nil && !IsNil(o.AssignedObjectType) {
		return true
	}

	return false
}

// SetAssignedObjectType gets a reference to the given string and assigns it to the AssignedObjectType field.
func (o *PatchedWritableL2VPNTerminationRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType = &v
}

// GetAssignedObjectId returns the AssignedObjectId field value if set, zero value otherwise.
func (o *PatchedWritableL2VPNTerminationRequest) GetAssignedObjectId() int64 {
	if o == nil || IsNil(o.AssignedObjectId) {
		var ret int64
		return ret
	}
	return *o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableL2VPNTerminationRequest) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AssignedObjectId) {
		return nil, false
	}
	return o.AssignedObjectId, true
}

// HasAssignedObjectId returns a boolean if a field has been set.
func (o *PatchedWritableL2VPNTerminationRequest) HasAssignedObjectId() bool {
	if o != nil && !IsNil(o.AssignedObjectId) {
		return true
	}

	return false
}

// SetAssignedObjectId gets a reference to the given int64 and assigns it to the AssignedObjectId field.
func (o *PatchedWritableL2VPNTerminationRequest) SetAssignedObjectId(v int64) {
	o.AssignedObjectId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableL2VPNTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableL2VPNTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableL2VPNTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableL2VPNTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableL2VPNTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableL2VPNTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableL2VPNTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableL2VPNTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableL2VPNTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableL2VPNTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.L2vpn) {
		toSerialize["l2vpn"] = o.L2vpn
	}
	if !IsNil(o.AssignedObjectType) {
		toSerialize["assigned_object_type"] = o.AssignedObjectType
	}
	if !IsNil(o.AssignedObjectId) {
		toSerialize["assigned_object_id"] = o.AssignedObjectId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullablePatchedWritableL2VPNTerminationRequest struct {
	value *PatchedWritableL2VPNTerminationRequest
	isSet bool
}

func (v NullablePatchedWritableL2VPNTerminationRequest) Get() *PatchedWritableL2VPNTerminationRequest {
	return v.value
}

func (v *NullablePatchedWritableL2VPNTerminationRequest) Set(val *PatchedWritableL2VPNTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableL2VPNTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableL2VPNTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableL2VPNTerminationRequest(val *PatchedWritableL2VPNTerminationRequest) *NullablePatchedWritableL2VPNTerminationRequest {
	return &NullablePatchedWritableL2VPNTerminationRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableL2VPNTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableL2VPNTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
