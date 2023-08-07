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

// checks if the WritableServiceTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableServiceTemplateRequest{}

// WritableServiceTemplateRequest Adds support for custom fields and tags.
type WritableServiceTemplateRequest struct {
	Name  string  `json:"name"`
	Ports []int32 `json:"ports"`
	// * `tcp` - TCP * `udp` - UDP * `sctp` - SCTP
	Protocol     string                 `json:"protocol"`
	Description  *string                `json:"description,omitempty"`
	Comments     *string                `json:"comments,omitempty"`
	Tags         []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewWritableServiceTemplateRequest instantiates a new WritableServiceTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableServiceTemplateRequest(name string, ports []int32, protocol string) *WritableServiceTemplateRequest {
	this := WritableServiceTemplateRequest{}
	this.Name = name
	this.Ports = ports
	this.Protocol = protocol
	return &this
}

// NewWritableServiceTemplateRequestWithDefaults instantiates a new WritableServiceTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableServiceTemplateRequestWithDefaults() *WritableServiceTemplateRequest {
	this := WritableServiceTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableServiceTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableServiceTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableServiceTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetPorts returns the Ports field value
func (o *WritableServiceTemplateRequest) GetPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *WritableServiceTemplateRequest) GetPortsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *WritableServiceTemplateRequest) SetPorts(v []int32) {
	o.Ports = v
}

// GetProtocol returns the Protocol field value
func (o *WritableServiceTemplateRequest) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *WritableServiceTemplateRequest) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *WritableServiceTemplateRequest) SetProtocol(v string) {
	o.Protocol = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableServiceTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableServiceTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableServiceTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableServiceTemplateRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceTemplateRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableServiceTemplateRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableServiceTemplateRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableServiceTemplateRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceTemplateRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableServiceTemplateRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableServiceTemplateRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableServiceTemplateRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableServiceTemplateRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableServiceTemplateRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableServiceTemplateRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableServiceTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableServiceTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["ports"] = o.Ports
	toSerialize["protocol"] = o.Protocol
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

type NullableWritableServiceTemplateRequest struct {
	value *WritableServiceTemplateRequest
	isSet bool
}

func (v NullableWritableServiceTemplateRequest) Get() *WritableServiceTemplateRequest {
	return v.value
}

func (v *NullableWritableServiceTemplateRequest) Set(val *WritableServiceTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableServiceTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableServiceTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableServiceTemplateRequest(val *WritableServiceTemplateRequest) *NullableWritableServiceTemplateRequest {
	return &NullableWritableServiceTemplateRequest{value: val, isSet: true}
}

func (v NullableWritableServiceTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableServiceTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
