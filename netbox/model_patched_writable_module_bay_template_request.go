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

// checks if the PatchedWritableModuleBayTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableModuleBayTemplateRequest{}

// PatchedWritableModuleBayTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableModuleBayTemplateRequest struct {
	DeviceType *int32 `json:"device_type,omitempty"`
	//          {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// Identifier to reference when renaming installed components
	Position    *string `json:"position,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewPatchedWritableModuleBayTemplateRequest instantiates a new PatchedWritableModuleBayTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableModuleBayTemplateRequest() *PatchedWritableModuleBayTemplateRequest {
	this := PatchedWritableModuleBayTemplateRequest{}
	return &this
}

// NewPatchedWritableModuleBayTemplateRequestWithDefaults instantiates a new PatchedWritableModuleBayTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableModuleBayTemplateRequestWithDefaults() *PatchedWritableModuleBayTemplateRequest {
	this := PatchedWritableModuleBayTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayTemplateRequest) GetDeviceType() int32 {
	if o == nil || IsNil(o.DeviceType) {
		var ret int32
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayTemplateRequest) GetDeviceTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayTemplateRequest) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given int32 and assigns it to the DeviceType field.
func (o *PatchedWritableModuleBayTemplateRequest) SetDeviceType(v int32) {
	o.DeviceType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableModuleBayTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableModuleBayTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayTemplateRequest) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayTemplateRequest) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayTemplateRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *PatchedWritableModuleBayTemplateRequest) SetPosition(v string) {
	o.Position = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableModuleBayTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleBayTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableModuleBayTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableModuleBayTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedWritableModuleBayTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableModuleBayTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullablePatchedWritableModuleBayTemplateRequest struct {
	value *PatchedWritableModuleBayTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableModuleBayTemplateRequest) Get() *PatchedWritableModuleBayTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableModuleBayTemplateRequest) Set(val *PatchedWritableModuleBayTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableModuleBayTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableModuleBayTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableModuleBayTemplateRequest(val *PatchedWritableModuleBayTemplateRequest) *NullablePatchedWritableModuleBayTemplateRequest {
	return &NullablePatchedWritableModuleBayTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableModuleBayTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableModuleBayTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
