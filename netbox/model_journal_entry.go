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

// checks if the JournalEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JournalEntry{}

// JournalEntry Adds support for custom fields and tags.
type JournalEntry struct {
	Id                 int32                  `json:"id"`
	Url                string                 `json:"url"`
	Display            string                 `json:"display"`
	AssignedObjectType string                 `json:"assigned_object_type"`
	AssignedObjectId   int64                  `json:"assigned_object_id"`
	AssignedObject     map[string]interface{} `json:"assigned_object"`
	Created            NullableTime           `json:"created"`
	CreatedBy          NullableInt32          `json:"created_by,omitempty"`
	Kind               *CableStatus           `json:"kind,omitempty"`
	Comments           string                 `json:"comments"`
	Tags               []NestedTag            `json:"tags,omitempty"`
	CustomFields       map[string]interface{} `json:"custom_fields,omitempty"`
	LastUpdated        NullableTime           `json:"last_updated"`
}

// NewJournalEntry instantiates a new JournalEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournalEntry(id int32, url string, display string, assignedObjectType string, assignedObjectId int64, assignedObject map[string]interface{}, created NullableTime, comments string, lastUpdated NullableTime) *JournalEntry {
	this := JournalEntry{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.AssignedObjectType = assignedObjectType
	this.AssignedObjectId = assignedObjectId
	this.AssignedObject = assignedObject
	this.Created = created
	this.Comments = comments
	this.LastUpdated = lastUpdated
	return &this
}

// NewJournalEntryWithDefaults instantiates a new JournalEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalEntryWithDefaults() *JournalEntry {
	this := JournalEntry{}
	return &this
}

// GetId returns the Id field value
func (o *JournalEntry) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JournalEntry) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *JournalEntry) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *JournalEntry) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *JournalEntry) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *JournalEntry) SetDisplay(v string) {
	o.Display = v
}

// GetAssignedObjectType returns the AssignedObjectType field value
func (o *JournalEntry) GetAssignedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectType, true
}

// SetAssignedObjectType sets field value
func (o *JournalEntry) SetAssignedObjectType(v string) {
	o.AssignedObjectType = v
}

// GetAssignedObjectId returns the AssignedObjectId field value
func (o *JournalEntry) GetAssignedObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectId, true
}

// SetAssignedObjectId sets field value
func (o *JournalEntry) SetAssignedObjectId(v int64) {
	o.AssignedObjectId = v
}

// GetAssignedObject returns the AssignedObject field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *JournalEntry) GetAssignedObject() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.AssignedObject
}

// GetAssignedObjectOk returns a tuple with the AssignedObject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalEntry) GetAssignedObjectOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AssignedObject) {
		return map[string]interface{}{}, false
	}
	return o.AssignedObject, true
}

// SetAssignedObject sets field value
func (o *JournalEntry) SetAssignedObject(v map[string]interface{}) {
	o.AssignedObject = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *JournalEntry) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalEntry) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *JournalEntry) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JournalEntry) GetCreatedBy() int32 {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret int32
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalEntry) GetCreatedByOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *JournalEntry) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableInt32 and assigns it to the CreatedBy field.
func (o *JournalEntry) SetCreatedBy(v int32) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *JournalEntry) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *JournalEntry) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *JournalEntry) GetKind() CableStatus {
	if o == nil || IsNil(o.Kind) {
		var ret CableStatus
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetKindOk() (*CableStatus, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *JournalEntry) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given CableStatus and assigns it to the Kind field.
func (o *JournalEntry) SetKind(v CableStatus) {
	o.Kind = &v
}

// GetComments returns the Comments field value
func (o *JournalEntry) GetComments() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *JournalEntry) SetComments(v string) {
	o.Comments = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *JournalEntry) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *JournalEntry) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *JournalEntry) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *JournalEntry) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalEntry) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *JournalEntry) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *JournalEntry) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *JournalEntry) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalEntry) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *JournalEntry) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o JournalEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JournalEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["assigned_object_type"] = o.AssignedObjectType
	toSerialize["assigned_object_id"] = o.AssignedObjectId
	if o.AssignedObject != nil {
		toSerialize["assigned_object"] = o.AssignedObject
	}
	toSerialize["created"] = o.Created.Get()
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["comments"] = o.Comments
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["last_updated"] = o.LastUpdated.Get()
	return toSerialize, nil
}

type NullableJournalEntry struct {
	value *JournalEntry
	isSet bool
}

func (v NullableJournalEntry) Get() *JournalEntry {
	return v.value
}

func (v *NullableJournalEntry) Set(val *JournalEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableJournalEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableJournalEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournalEntry(val *JournalEntry) *NullableJournalEntry {
	return &NullableJournalEntry{value: val, isSet: true}
}

func (v NullableJournalEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournalEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
