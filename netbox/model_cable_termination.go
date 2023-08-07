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

// checks if the CableTermination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CableTermination{}

// CableTermination Adds support for custom fields and tags.
type CableTermination struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Cable   int32  `json:"cable"`
	// * `A` - A * `B` - B
	CableEnd        string                 `json:"cable_end"`
	TerminationType string                 `json:"termination_type"`
	TerminationId   int64                  `json:"termination_id"`
	Termination     map[string]interface{} `json:"termination"`
	Created         NullableTime           `json:"created"`
	LastUpdated     NullableTime           `json:"last_updated"`
}

// NewCableTermination instantiates a new CableTermination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCableTermination(id int32, url string, display string, cable int32, cableEnd string, terminationType string, terminationId int64, termination map[string]interface{}, created NullableTime, lastUpdated NullableTime) *CableTermination {
	this := CableTermination{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Cable = cable
	this.CableEnd = cableEnd
	this.TerminationType = terminationType
	this.TerminationId = terminationId
	this.Termination = termination
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewCableTerminationWithDefaults instantiates a new CableTermination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCableTerminationWithDefaults() *CableTermination {
	this := CableTermination{}
	return &this
}

// GetId returns the Id field value
func (o *CableTermination) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CableTermination) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CableTermination) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CableTermination) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CableTermination) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CableTermination) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *CableTermination) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CableTermination) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CableTermination) SetDisplay(v string) {
	o.Display = v
}

// GetCable returns the Cable field value
func (o *CableTermination) GetCable() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cable
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
func (o *CableTermination) GetCableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cable, true
}

// SetCable sets field value
func (o *CableTermination) SetCable(v int32) {
	o.Cable = v
}

// GetCableEnd returns the CableEnd field value
func (o *CableTermination) GetCableEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CableEnd
}

// GetCableEndOk returns a tuple with the CableEnd field value
// and a boolean to check if the value has been set.
func (o *CableTermination) GetCableEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CableEnd, true
}

// SetCableEnd sets field value
func (o *CableTermination) SetCableEnd(v string) {
	o.CableEnd = v
}

// GetTerminationType returns the TerminationType field value
func (o *CableTermination) GetTerminationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminationType
}

// GetTerminationTypeOk returns a tuple with the TerminationType field value
// and a boolean to check if the value has been set.
func (o *CableTermination) GetTerminationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationType, true
}

// SetTerminationType sets field value
func (o *CableTermination) SetTerminationType(v string) {
	o.TerminationType = v
}

// GetTerminationId returns the TerminationId field value
func (o *CableTermination) GetTerminationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TerminationId
}

// GetTerminationIdOk returns a tuple with the TerminationId field value
// and a boolean to check if the value has been set.
func (o *CableTermination) GetTerminationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationId, true
}

// SetTerminationId sets field value
func (o *CableTermination) SetTerminationId(v int64) {
	o.TerminationId = v
}

// GetTermination returns the Termination field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *CableTermination) GetTermination() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Termination
}

// GetTerminationOk returns a tuple with the Termination field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CableTermination) GetTerminationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Termination) {
		return map[string]interface{}{}, false
	}
	return o.Termination, true
}

// SetTermination sets field value
func (o *CableTermination) SetTermination(v map[string]interface{}) {
	o.Termination = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CableTermination) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CableTermination) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *CableTermination) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CableTermination) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CableTermination) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *CableTermination) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o CableTermination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CableTermination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["cable"] = o.Cable
	toSerialize["cable_end"] = o.CableEnd
	toSerialize["termination_type"] = o.TerminationType
	toSerialize["termination_id"] = o.TerminationId
	if o.Termination != nil {
		toSerialize["termination"] = o.Termination
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	return toSerialize, nil
}

type NullableCableTermination struct {
	value *CableTermination
	isSet bool
}

func (v NullableCableTermination) Get() *CableTermination {
	return v.value
}

func (v *NullableCableTermination) Set(val *CableTermination) {
	v.value = val
	v.isSet = true
}

func (v NullableCableTermination) IsSet() bool {
	return v.isSet
}

func (v *NullableCableTermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableTermination(val *CableTermination) *NullableCableTermination {
	return &NullableCableTermination{value: val, isSet: true}
}

func (v NullableCableTermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableTermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
