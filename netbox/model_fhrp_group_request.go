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

// checks if the FHRPGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FHRPGroupRequest{}

// FHRPGroupRequest Adds support for custom fields and tags.
type FHRPGroupRequest struct {
	Name *string `json:"name,omitempty"`
	// * `vrrp2` - VRRPv2 * `vrrp3` - VRRPv3 * `carp` - CARP * `clusterxl` - ClusterXL * `hsrp` - HSRP * `glbp` - GLBP * `other` - Other
	Protocol string `json:"protocol"`
	GroupId  int32  `json:"group_id"`
	// * `plaintext` - Plaintext * `md5` - MD5
	AuthType     *string                `json:"auth_type,omitempty"`
	AuthKey      *string                `json:"auth_key,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Comments     *string                `json:"comments,omitempty"`
	Tags         []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewFHRPGroupRequest instantiates a new FHRPGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFHRPGroupRequest(protocol string, groupId int32) *FHRPGroupRequest {
	this := FHRPGroupRequest{}
	this.Protocol = protocol
	this.GroupId = groupId
	return &this
}

// NewFHRPGroupRequestWithDefaults instantiates a new FHRPGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFHRPGroupRequestWithDefaults() *FHRPGroupRequest {
	this := FHRPGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FHRPGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FHRPGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FHRPGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value
func (o *FHRPGroupRequest) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *FHRPGroupRequest) SetProtocol(v string) {
	o.Protocol = v
}

// GetGroupId returns the GroupId field value
func (o *FHRPGroupRequest) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *FHRPGroupRequest) SetGroupId(v int32) {
	o.GroupId = v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *FHRPGroupRequest) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *FHRPGroupRequest) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *FHRPGroupRequest) SetAuthType(v string) {
	o.AuthType = &v
}

// GetAuthKey returns the AuthKey field value if set, zero value otherwise.
func (o *FHRPGroupRequest) GetAuthKey() string {
	if o == nil || IsNil(o.AuthKey) {
		var ret string
		return ret
	}
	return *o.AuthKey
}

// GetAuthKeyOk returns a tuple with the AuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetAuthKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AuthKey) {
		return nil, false
	}
	return o.AuthKey, true
}

// HasAuthKey returns a boolean if a field has been set.
func (o *FHRPGroupRequest) HasAuthKey() bool {
	if o != nil && !IsNil(o.AuthKey) {
		return true
	}

	return false
}

// SetAuthKey gets a reference to the given string and assigns it to the AuthKey field.
func (o *FHRPGroupRequest) SetAuthKey(v string) {
	o.AuthKey = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FHRPGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FHRPGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FHRPGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *FHRPGroupRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *FHRPGroupRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *FHRPGroupRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FHRPGroupRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FHRPGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *FHRPGroupRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *FHRPGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *FHRPGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *FHRPGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o FHRPGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FHRPGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["protocol"] = o.Protocol
	toSerialize["group_id"] = o.GroupId
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.AuthKey) {
		toSerialize["auth_key"] = o.AuthKey
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

type NullableFHRPGroupRequest struct {
	value *FHRPGroupRequest
	isSet bool
}

func (v NullableFHRPGroupRequest) Get() *FHRPGroupRequest {
	return v.value
}

func (v *NullableFHRPGroupRequest) Set(val *FHRPGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFHRPGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFHRPGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFHRPGroupRequest(val *FHRPGroupRequest) *NullableFHRPGroupRequest {
	return &NullableFHRPGroupRequest{value: val, isSet: true}
}

func (v NullableFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFHRPGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
