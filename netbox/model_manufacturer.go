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

// checks if the Manufacturer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Manufacturer{}

// Manufacturer Adds support for custom fields and tags.
type Manufacturer struct {
	Id                 int32                  `json:"id"`
	Url                string                 `json:"url"`
	Display            string                 `json:"display"`
	Name               string                 `json:"name"`
	Slug               string                 `json:"slug"`
	Description        *string                `json:"description,omitempty"`
	Tags               []NestedTag            `json:"tags,omitempty"`
	CustomFields       map[string]interface{} `json:"custom_fields,omitempty"`
	Created            NullableTime           `json:"created"`
	LastUpdated        NullableTime           `json:"last_updated"`
	DevicetypeCount    int32                  `json:"devicetype_count"`
	InventoryitemCount int32                  `json:"inventoryitem_count"`
	PlatformCount      int32                  `json:"platform_count"`
}

// NewManufacturer instantiates a new Manufacturer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManufacturer(id int32, url string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, devicetypeCount int32, inventoryitemCount int32, platformCount int32) *Manufacturer {
	this := Manufacturer{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Created = created
	this.LastUpdated = lastUpdated
	this.DevicetypeCount = devicetypeCount
	this.InventoryitemCount = inventoryitemCount
	this.PlatformCount = platformCount
	return &this
}

// NewManufacturerWithDefaults instantiates a new Manufacturer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManufacturerWithDefaults() *Manufacturer {
	this := Manufacturer{}
	return &this
}

// GetId returns the Id field value
func (o *Manufacturer) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Manufacturer) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Manufacturer) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Manufacturer) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *Manufacturer) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Manufacturer) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *Manufacturer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Manufacturer) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Manufacturer) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Manufacturer) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Manufacturer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Manufacturer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Manufacturer) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Manufacturer) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Manufacturer) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Manufacturer) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Manufacturer) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Manufacturer) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Manufacturer) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Manufacturer) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Manufacturer) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Manufacturer) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Manufacturer) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Manufacturer) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Manufacturer) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetDevicetypeCount returns the DevicetypeCount field value
func (o *Manufacturer) GetDevicetypeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DevicetypeCount
}

// GetDevicetypeCountOk returns a tuple with the DevicetypeCount field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetDevicetypeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevicetypeCount, true
}

// SetDevicetypeCount sets field value
func (o *Manufacturer) SetDevicetypeCount(v int32) {
	o.DevicetypeCount = v
}

// GetInventoryitemCount returns the InventoryitemCount field value
func (o *Manufacturer) GetInventoryitemCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InventoryitemCount
}

// GetInventoryitemCountOk returns a tuple with the InventoryitemCount field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetInventoryitemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryitemCount, true
}

// SetInventoryitemCount sets field value
func (o *Manufacturer) SetInventoryitemCount(v int32) {
	o.InventoryitemCount = v
}

// GetPlatformCount returns the PlatformCount field value
func (o *Manufacturer) GetPlatformCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlatformCount
}

// GetPlatformCountOk returns a tuple with the PlatformCount field value
// and a boolean to check if the value has been set.
func (o *Manufacturer) GetPlatformCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformCount, true
}

// SetPlatformCount sets field value
func (o *Manufacturer) SetPlatformCount(v int32) {
	o.PlatformCount = v
}

func (o Manufacturer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Manufacturer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	// skip: devicetype_count is readOnly
	// skip: inventoryitem_count is readOnly
	// skip: platform_count is readOnly
	return toSerialize, nil
}

type NullableManufacturer struct {
	value *Manufacturer
	isSet bool
}

func (v NullableManufacturer) Get() *Manufacturer {
	return v.value
}

func (v *NullableManufacturer) Set(val *Manufacturer) {
	v.value = val
	v.isSet = true
}

func (v NullableManufacturer) IsSet() bool {
	return v.isSet
}

func (v *NullableManufacturer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManufacturer(val *Manufacturer) *NullableManufacturer {
	return &NullableManufacturer{value: val, isSet: true}
}

func (v NullableManufacturer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManufacturer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
