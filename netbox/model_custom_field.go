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

// checks if the CustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomField{}

// CustomField Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CustomField struct {
	Id           int32       `json:"id"`
	Url          string      `json:"url"`
	Display      string      `json:"display"`
	ContentTypes []string    `json:"content_types"`
	Type         CableStatus `json:"type"`
	ObjectType   *string     `json:"object_type,omitempty"`
	DataType     string      `json:"data_type"`
	// Internal field name
	Name string `json:"name"`
	// Name of the field as displayed to users (if not provided, the field's name will be used)
	Label *string `json:"label,omitempty"`
	// Custom fields within the same group will be displayed together
	GroupName   *string `json:"group_name,omitempty"`
	Description *string `json:"description,omitempty"`
	// If true, this field is required when creating new objects or editing an existing object.
	Required *bool `json:"required,omitempty"`
	// Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored.
	SearchWeight *int32       `json:"search_weight,omitempty"`
	FilterLogic  *CableStatus `json:"filter_logic,omitempty"`
	UiVisibility *CableStatus `json:"ui_visibility,omitempty"`
	// Replicate this value when cloning objects
	IsCloneable *bool `json:"is_cloneable,omitempty"`
	// Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \"Foo\").
	Default map[string]interface{} `json:"default,omitempty"`
	// Fields with higher weights appear lower in a form.
	Weight *int32 `json:"weight,omitempty"`
	// Minimum allowed value (for numeric fields)
	ValidationMinimum NullableInt32 `json:"validation_minimum,omitempty"`
	// Maximum allowed value (for numeric fields)
	ValidationMaximum NullableInt32 `json:"validation_maximum,omitempty"`
	// Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, <code>^[A-Z]{3}$</code> will limit values to exactly three uppercase letters.
	ValidationRegex *string `json:"validation_regex,omitempty"`
	// Comma-separated list of available choices (for selection fields)
	Choices     []string     `json:"choices,omitempty"`
	Created     NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
}

// NewCustomField instantiates a new CustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomField(id int32, url string, display string, contentTypes []string, type_ CableStatus, dataType string, name string, created NullableTime, lastUpdated NullableTime) *CustomField {
	this := CustomField{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.ContentTypes = contentTypes
	this.Type = type_
	this.DataType = dataType
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewCustomFieldWithDefaults instantiates a new CustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldWithDefaults() *CustomField {
	this := CustomField{}
	return &this
}

// GetId returns the Id field value
func (o *CustomField) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomField) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CustomField) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CustomField) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *CustomField) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CustomField) SetDisplay(v string) {
	o.Display = v
}

// GetContentTypes returns the ContentTypes field value
func (o *CustomField) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *CustomField) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetType returns the Type field value
func (o *CustomField) GetType() CableStatus {
	if o == nil {
		var ret CableStatus
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetTypeOk() (*CableStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomField) SetType(v CableStatus) {
	o.Type = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *CustomField) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *CustomField) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *CustomField) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetDataType returns the DataType field value
func (o *CustomField) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *CustomField) SetDataType(v string) {
	o.DataType = v
}

// GetName returns the Name field value
func (o *CustomField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomField) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CustomField) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomField) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CustomField) SetLabel(v string) {
	o.Label = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *CustomField) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *CustomField) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *CustomField) SetGroupName(v string) {
	o.GroupName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomField) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomField) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomField) SetDescription(v string) {
	o.Description = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *CustomField) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *CustomField) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *CustomField) SetRequired(v bool) {
	o.Required = &v
}

// GetSearchWeight returns the SearchWeight field value if set, zero value otherwise.
func (o *CustomField) GetSearchWeight() int32 {
	if o == nil || IsNil(o.SearchWeight) {
		var ret int32
		return ret
	}
	return *o.SearchWeight
}

// GetSearchWeightOk returns a tuple with the SearchWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetSearchWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.SearchWeight) {
		return nil, false
	}
	return o.SearchWeight, true
}

// HasSearchWeight returns a boolean if a field has been set.
func (o *CustomField) HasSearchWeight() bool {
	if o != nil && !IsNil(o.SearchWeight) {
		return true
	}

	return false
}

// SetSearchWeight gets a reference to the given int32 and assigns it to the SearchWeight field.
func (o *CustomField) SetSearchWeight(v int32) {
	o.SearchWeight = &v
}

// GetFilterLogic returns the FilterLogic field value if set, zero value otherwise.
func (o *CustomField) GetFilterLogic() CableStatus {
	if o == nil || IsNil(o.FilterLogic) {
		var ret CableStatus
		return ret
	}
	return *o.FilterLogic
}

// GetFilterLogicOk returns a tuple with the FilterLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetFilterLogicOk() (*CableStatus, bool) {
	if o == nil || IsNil(o.FilterLogic) {
		return nil, false
	}
	return o.FilterLogic, true
}

// HasFilterLogic returns a boolean if a field has been set.
func (o *CustomField) HasFilterLogic() bool {
	if o != nil && !IsNil(o.FilterLogic) {
		return true
	}

	return false
}

// SetFilterLogic gets a reference to the given CableStatus and assigns it to the FilterLogic field.
func (o *CustomField) SetFilterLogic(v CableStatus) {
	o.FilterLogic = &v
}

// GetUiVisibility returns the UiVisibility field value if set, zero value otherwise.
func (o *CustomField) GetUiVisibility() CableStatus {
	if o == nil || IsNil(o.UiVisibility) {
		var ret CableStatus
		return ret
	}
	return *o.UiVisibility
}

// GetUiVisibilityOk returns a tuple with the UiVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetUiVisibilityOk() (*CableStatus, bool) {
	if o == nil || IsNil(o.UiVisibility) {
		return nil, false
	}
	return o.UiVisibility, true
}

// HasUiVisibility returns a boolean if a field has been set.
func (o *CustomField) HasUiVisibility() bool {
	if o != nil && !IsNil(o.UiVisibility) {
		return true
	}

	return false
}

// SetUiVisibility gets a reference to the given CableStatus and assigns it to the UiVisibility field.
func (o *CustomField) SetUiVisibility(v CableStatus) {
	o.UiVisibility = &v
}

// GetIsCloneable returns the IsCloneable field value if set, zero value otherwise.
func (o *CustomField) GetIsCloneable() bool {
	if o == nil || IsNil(o.IsCloneable) {
		var ret bool
		return ret
	}
	return *o.IsCloneable
}

// GetIsCloneableOk returns a tuple with the IsCloneable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetIsCloneableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloneable) {
		return nil, false
	}
	return o.IsCloneable, true
}

// HasIsCloneable returns a boolean if a field has been set.
func (o *CustomField) HasIsCloneable() bool {
	if o != nil && !IsNil(o.IsCloneable) {
		return true
	}

	return false
}

// SetIsCloneable gets a reference to the given bool and assigns it to the IsCloneable field.
func (o *CustomField) SetIsCloneable(v bool) {
	o.IsCloneable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomField) GetDefault() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetDefaultOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Default) {
		return map[string]interface{}{}, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CustomField) HasDefault() bool {
	if o != nil && IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given map[string]interface{} and assigns it to the Default field.
func (o *CustomField) SetDefault(v map[string]interface{}) {
	o.Default = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *CustomField) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *CustomField) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *CustomField) SetWeight(v int32) {
	o.Weight = &v
}

// GetValidationMinimum returns the ValidationMinimum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomField) GetValidationMinimum() int32 {
	if o == nil || IsNil(o.ValidationMinimum.Get()) {
		var ret int32
		return ret
	}
	return *o.ValidationMinimum.Get()
}

// GetValidationMinimumOk returns a tuple with the ValidationMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetValidationMinimumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationMinimum.Get(), o.ValidationMinimum.IsSet()
}

// HasValidationMinimum returns a boolean if a field has been set.
func (o *CustomField) HasValidationMinimum() bool {
	if o != nil && o.ValidationMinimum.IsSet() {
		return true
	}

	return false
}

// SetValidationMinimum gets a reference to the given NullableInt32 and assigns it to the ValidationMinimum field.
func (o *CustomField) SetValidationMinimum(v int32) {
	o.ValidationMinimum.Set(&v)
}

// SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil
func (o *CustomField) SetValidationMinimumNil() {
	o.ValidationMinimum.Set(nil)
}

// UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
func (o *CustomField) UnsetValidationMinimum() {
	o.ValidationMinimum.Unset()
}

// GetValidationMaximum returns the ValidationMaximum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomField) GetValidationMaximum() int32 {
	if o == nil || IsNil(o.ValidationMaximum.Get()) {
		var ret int32
		return ret
	}
	return *o.ValidationMaximum.Get()
}

// GetValidationMaximumOk returns a tuple with the ValidationMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetValidationMaximumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationMaximum.Get(), o.ValidationMaximum.IsSet()
}

// HasValidationMaximum returns a boolean if a field has been set.
func (o *CustomField) HasValidationMaximum() bool {
	if o != nil && o.ValidationMaximum.IsSet() {
		return true
	}

	return false
}

// SetValidationMaximum gets a reference to the given NullableInt32 and assigns it to the ValidationMaximum field.
func (o *CustomField) SetValidationMaximum(v int32) {
	o.ValidationMaximum.Set(&v)
}

// SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil
func (o *CustomField) SetValidationMaximumNil() {
	o.ValidationMaximum.Set(nil)
}

// UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
func (o *CustomField) UnsetValidationMaximum() {
	o.ValidationMaximum.Unset()
}

// GetValidationRegex returns the ValidationRegex field value if set, zero value otherwise.
func (o *CustomField) GetValidationRegex() string {
	if o == nil || IsNil(o.ValidationRegex) {
		var ret string
		return ret
	}
	return *o.ValidationRegex
}

// GetValidationRegexOk returns a tuple with the ValidationRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetValidationRegexOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationRegex) {
		return nil, false
	}
	return o.ValidationRegex, true
}

// HasValidationRegex returns a boolean if a field has been set.
func (o *CustomField) HasValidationRegex() bool {
	if o != nil && !IsNil(o.ValidationRegex) {
		return true
	}

	return false
}

// SetValidationRegex gets a reference to the given string and assigns it to the ValidationRegex field.
func (o *CustomField) SetValidationRegex(v string) {
	o.ValidationRegex = &v
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomField) GetChoices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetChoicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *CustomField) HasChoices() bool {
	if o != nil && IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []string and assigns it to the Choices field.
func (o *CustomField) SetChoices(v []string) {
	o.Choices = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CustomField) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *CustomField) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CustomField) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *CustomField) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o CustomField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["type"] = o.Type
	if !IsNil(o.ObjectType) {
		toSerialize["object_type"] = o.ObjectType
	}
	// skip: data_type is readOnly
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.SearchWeight) {
		toSerialize["search_weight"] = o.SearchWeight
	}
	if !IsNil(o.FilterLogic) {
		toSerialize["filter_logic"] = o.FilterLogic
	}
	if !IsNil(o.UiVisibility) {
		toSerialize["ui_visibility"] = o.UiVisibility
	}
	if !IsNil(o.IsCloneable) {
		toSerialize["is_cloneable"] = o.IsCloneable
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if o.ValidationMinimum.IsSet() {
		toSerialize["validation_minimum"] = o.ValidationMinimum.Get()
	}
	if o.ValidationMaximum.IsSet() {
		toSerialize["validation_maximum"] = o.ValidationMaximum.Get()
	}
	if !IsNil(o.ValidationRegex) {
		toSerialize["validation_regex"] = o.ValidationRegex
	}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	return toSerialize, nil
}

type NullableCustomField struct {
	value *CustomField
	isSet bool
}

func (v NullableCustomField) Get() *CustomField {
	return v.value
}

func (v *NullableCustomField) Set(val *CustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomField(val *CustomField) *NullableCustomField {
	return &NullableCustomField{value: val, isSet: true}
}

func (v NullableCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
