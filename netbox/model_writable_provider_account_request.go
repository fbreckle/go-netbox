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

// checks if the WritableProviderAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableProviderAccountRequest{}

// WritableProviderAccountRequest Adds support for custom fields and tags.
type WritableProviderAccountRequest struct {
	Provider     int32                  `json:"provider"`
	Name         *string                `json:"name,omitempty"`
	Account      string                 `json:"account"`
	Description  *string                `json:"description,omitempty"`
	Comments     *string                `json:"comments,omitempty"`
	Tags         []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewWritableProviderAccountRequest instantiates a new WritableProviderAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableProviderAccountRequest(provider int32, account string) *WritableProviderAccountRequest {
	this := WritableProviderAccountRequest{}
	this.Provider = provider
	this.Account = account
	return &this
}

// NewWritableProviderAccountRequestWithDefaults instantiates a new WritableProviderAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableProviderAccountRequestWithDefaults() *WritableProviderAccountRequest {
	this := WritableProviderAccountRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *WritableProviderAccountRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *WritableProviderAccountRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *WritableProviderAccountRequest) SetProvider(v int32) {
	o.Provider = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WritableProviderAccountRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableProviderAccountRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WritableProviderAccountRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WritableProviderAccountRequest) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value
func (o *WritableProviderAccountRequest) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *WritableProviderAccountRequest) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *WritableProviderAccountRequest) SetAccount(v string) {
	o.Account = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableProviderAccountRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableProviderAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableProviderAccountRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableProviderAccountRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableProviderAccountRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableProviderAccountRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableProviderAccountRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableProviderAccountRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableProviderAccountRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableProviderAccountRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableProviderAccountRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableProviderAccountRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableProviderAccountRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableProviderAccountRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableProviderAccountRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableProviderAccountRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableProviderAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableProviderAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["account"] = o.Account
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

type NullableWritableProviderAccountRequest struct {
	value *WritableProviderAccountRequest
	isSet bool
}

func (v NullableWritableProviderAccountRequest) Get() *WritableProviderAccountRequest {
	return v.value
}

func (v *NullableWritableProviderAccountRequest) Set(val *WritableProviderAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableProviderAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableProviderAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableProviderAccountRequest(val *WritableProviderAccountRequest) *NullableWritableProviderAccountRequest {
	return &NullableWritableProviderAccountRequest{value: val, isSet: true}
}

func (v NullableWritableProviderAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableProviderAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
