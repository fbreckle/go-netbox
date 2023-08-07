/*
NetBox API

API to access NetBox

API version: 3.5.0 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// checks if the FHRPGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FHRPGroup{}

// FHRPGroup Adds support for custom fields and tags.
type FHRPGroup struct {
	Id      int32   `json:"id"`
	Name    *string `json:"name,omitempty"`
	Url     string  `json:"url"`
	Display string  `json:"display"`
	// * `vrrp2` - VRRPv2 * `vrrp3` - VRRPv3 * `carp` - CARP * `clusterxl` - ClusterXL * `hsrp` - HSRP * `glbp` - GLBP * `other` - Other
	Protocol string `json:"protocol"`
	GroupId  int32  `json:"group_id"`
	// * `plaintext` - Plaintext * `md5` - MD5
	AuthType     *string                `json:"auth_type,omitempty"`
	AuthKey      *string                `json:"auth_key,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Comments     *string                `json:"comments,omitempty"`
	Tags         []NestedTag            `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created      NullableTime           `json:"created"`
	LastUpdated  NullableTime           `json:"last_updated"`
	IpAddresses  []NestedIPAddress      `json:"ip_addresses"`
}

// NewFHRPGroup instantiates a new FHRPGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFHRPGroup(id int32, url string, display string, protocol string, groupId int32, created NullableTime, lastUpdated NullableTime, ipAddresses []NestedIPAddress) *FHRPGroup {
	this := FHRPGroup{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Protocol = protocol
	this.GroupId = groupId
	this.Created = created
	this.LastUpdated = lastUpdated
	this.IpAddresses = ipAddresses
	return &this
}

// NewFHRPGroupWithDefaults instantiates a new FHRPGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFHRPGroupWithDefaults() *FHRPGroup {
	this := FHRPGroup{}
	return &this
}

// GetId returns the Id field value
func (o *FHRPGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FHRPGroup) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FHRPGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FHRPGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FHRPGroup) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value
func (o *FHRPGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FHRPGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *FHRPGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *FHRPGroup) SetDisplay(v string) {
	o.Display = v
}

// GetProtocol returns the Protocol field value
func (o *FHRPGroup) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *FHRPGroup) SetProtocol(v string) {
	o.Protocol = v
}

// GetGroupId returns the GroupId field value
func (o *FHRPGroup) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *FHRPGroup) SetGroupId(v int32) {
	o.GroupId = v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *FHRPGroup) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *FHRPGroup) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *FHRPGroup) SetAuthType(v string) {
	o.AuthType = &v
}

// GetAuthKey returns the AuthKey field value if set, zero value otherwise.
func (o *FHRPGroup) GetAuthKey() string {
	if o == nil || IsNil(o.AuthKey) {
		var ret string
		return ret
	}
	return *o.AuthKey
}

// GetAuthKeyOk returns a tuple with the AuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetAuthKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AuthKey) {
		return nil, false
	}
	return o.AuthKey, true
}

// HasAuthKey returns a boolean if a field has been set.
func (o *FHRPGroup) HasAuthKey() bool {
	if o != nil && !IsNil(o.AuthKey) {
		return true
	}

	return false
}

// SetAuthKey gets a reference to the given string and assigns it to the AuthKey field.
func (o *FHRPGroup) SetAuthKey(v string) {
	o.AuthKey = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FHRPGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FHRPGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FHRPGroup) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *FHRPGroup) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *FHRPGroup) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *FHRPGroup) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FHRPGroup) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FHRPGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *FHRPGroup) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *FHRPGroup) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *FHRPGroup) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *FHRPGroup) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FHRPGroup) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FHRPGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *FHRPGroup) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FHRPGroup) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FHRPGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *FHRPGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetIpAddresses returns the IpAddresses field value
func (o *FHRPGroup) GetIpAddresses() []NestedIPAddress {
	if o == nil {
		var ret []NestedIPAddress
		return ret
	}

	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value
// and a boolean to check if the value has been set.
func (o *FHRPGroup) GetIpAddressesOk() ([]NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// SetIpAddresses sets field value
func (o *FHRPGroup) SetIpAddresses(v []NestedIPAddress) {
	o.IpAddresses = v
}

func (o FHRPGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FHRPGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	// skip: url is readOnly
	// skip: display is readOnly
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	// skip: ip_addresses is readOnly
	return toSerialize, nil
}

type NullableFHRPGroup struct {
	value *FHRPGroup
	isSet bool
}

func (v NullableFHRPGroup) Get() *FHRPGroup {
	return v.value
}

func (v *NullableFHRPGroup) Set(val *FHRPGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFHRPGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFHRPGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFHRPGroup(val *FHRPGroup) *NullableFHRPGroup {
	return &NullableFHRPGroup{value: val, isSet: true}
}

func (v NullableFHRPGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFHRPGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
