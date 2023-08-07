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

// checks if the VirtualChassisRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualChassisRequest{}

// VirtualChassisRequest Adds support for custom fields and tags.
type VirtualChassisRequest struct {
	Name         string                      `json:"name"`
	Domain       *string                     `json:"domain,omitempty"`
	Master       NullableNestedDeviceRequest `json:"master,omitempty"`
	Description  *string                     `json:"description,omitempty"`
	Comments     *string                     `json:"comments,omitempty"`
	Tags         []NestedTagRequest          `json:"tags,omitempty"`
	CustomFields map[string]interface{}      `json:"custom_fields,omitempty"`
}

// NewVirtualChassisRequest instantiates a new VirtualChassisRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualChassisRequest(name string) *VirtualChassisRequest {
	this := VirtualChassisRequest{}
	this.Name = name
	return &this
}

// NewVirtualChassisRequestWithDefaults instantiates a new VirtualChassisRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualChassisRequestWithDefaults() *VirtualChassisRequest {
	this := VirtualChassisRequest{}
	return &this
}

// GetName returns the Name field value
func (o *VirtualChassisRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualChassisRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualChassisRequest) SetName(v string) {
	o.Name = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *VirtualChassisRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *VirtualChassisRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *VirtualChassisRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetMaster returns the Master field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualChassisRequest) GetMaster() NestedDeviceRequest {
	if o == nil || IsNil(o.Master.Get()) {
		var ret NestedDeviceRequest
		return ret
	}
	return *o.Master.Get()
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualChassisRequest) GetMasterOk() (*NestedDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Master.Get(), o.Master.IsSet()
}

// HasMaster returns a boolean if a field has been set.
func (o *VirtualChassisRequest) HasMaster() bool {
	if o != nil && o.Master.IsSet() {
		return true
	}

	return false
}

// SetMaster gets a reference to the given NullableNestedDeviceRequest and assigns it to the Master field.
func (o *VirtualChassisRequest) SetMaster(v NestedDeviceRequest) {
	o.Master.Set(&v)
}

// SetMasterNil sets the value for Master to be an explicit nil
func (o *VirtualChassisRequest) SetMasterNil() {
	o.Master.Set(nil)
}

// UnsetMaster ensures that no value is present for Master, not even an explicit nil
func (o *VirtualChassisRequest) UnsetMaster() {
	o.Master.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualChassisRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualChassisRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualChassisRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VirtualChassisRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VirtualChassisRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VirtualChassisRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualChassisRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualChassisRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *VirtualChassisRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VirtualChassisRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VirtualChassisRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VirtualChassisRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o VirtualChassisRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualChassisRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if o.Master.IsSet() {
		toSerialize["master"] = o.Master.Get()
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

type NullableVirtualChassisRequest struct {
	value *VirtualChassisRequest
	isSet bool
}

func (v NullableVirtualChassisRequest) Get() *VirtualChassisRequest {
	return v.value
}

func (v *NullableVirtualChassisRequest) Set(val *VirtualChassisRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualChassisRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualChassisRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualChassisRequest(val *VirtualChassisRequest) *NullableVirtualChassisRequest {
	return &NullableVirtualChassisRequest{value: val, isSet: true}
}

func (v NullableVirtualChassisRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualChassisRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
