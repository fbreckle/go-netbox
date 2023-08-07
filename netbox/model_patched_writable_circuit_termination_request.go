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

// checks if the PatchedWritableCircuitTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableCircuitTerminationRequest{}

// PatchedWritableCircuitTerminationRequest Adds support for custom fields and tags.
type PatchedWritableCircuitTerminationRequest struct {
	Circuit *int32 `json:"circuit,omitempty"`
	// * `A` - A * `Z` - Z
	TermSide        *string       `json:"term_side,omitempty"`
	Site            NullableInt32 `json:"site,omitempty"`
	ProviderNetwork NullableInt32 `json:"provider_network,omitempty"`
	// Physical circuit speed
	PortSpeed NullableInt32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt32 `json:"upstream_speed,omitempty"`
	// ID of the local cross-connect
	XconnectId *string `json:"xconnect_id,omitempty"`
	// Patch panel ID and port number(s)
	PpInfo      *string `json:"pp_info,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool                  `json:"mark_connected,omitempty"`
	Tags          []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields  map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewPatchedWritableCircuitTerminationRequest instantiates a new PatchedWritableCircuitTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCircuitTerminationRequest() *PatchedWritableCircuitTerminationRequest {
	this := PatchedWritableCircuitTerminationRequest{}
	return &this
}

// NewPatchedWritableCircuitTerminationRequestWithDefaults instantiates a new PatchedWritableCircuitTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCircuitTerminationRequestWithDefaults() *PatchedWritableCircuitTerminationRequest {
	this := PatchedWritableCircuitTerminationRequest{}
	return &this
}

// GetCircuit returns the Circuit field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTerminationRequest) GetCircuit() int32 {
	if o == nil || IsNil(o.Circuit) {
		var ret int32
		return ret
	}
	return *o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTerminationRequest) GetCircuitOk() (*int32, bool) {
	if o == nil || IsNil(o.Circuit) {
		return nil, false
	}
	return o.Circuit, true
}

// HasCircuit returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasCircuit() bool {
	if o != nil && !IsNil(o.Circuit) {
		return true
	}

	return false
}

// SetCircuit gets a reference to the given int32 and assigns it to the Circuit field.
func (o *PatchedWritableCircuitTerminationRequest) SetCircuit(v int32) {
	o.Circuit = &v
}

// GetTermSide returns the TermSide field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTerminationRequest) GetTermSide() string {
	if o == nil || IsNil(o.TermSide) {
		var ret string
		return ret
	}
	return *o.TermSide
}

// GetTermSideOk returns a tuple with the TermSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTerminationRequest) GetTermSideOk() (*string, bool) {
	if o == nil || IsNil(o.TermSide) {
		return nil, false
	}
	return o.TermSide, true
}

// HasTermSide returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasTermSide() bool {
	if o != nil && !IsNil(o.TermSide) {
		return true
	}

	return false
}

// SetTermSide gets a reference to the given string and assigns it to the TermSide field.
func (o *PatchedWritableCircuitTerminationRequest) SetTermSide(v string) {
	o.TermSide = &v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTerminationRequest) GetSite() int32 {
	if o == nil || IsNil(o.Site.Get()) {
		var ret int32
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTerminationRequest) GetSiteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableInt32 and assigns it to the Site field.
func (o *PatchedWritableCircuitTerminationRequest) SetSite(v int32) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *PatchedWritableCircuitTerminationRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *PatchedWritableCircuitTerminationRequest) UnsetSite() {
	o.Site.Unset()
}

// GetProviderNetwork returns the ProviderNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTerminationRequest) GetProviderNetwork() int32 {
	if o == nil || IsNil(o.ProviderNetwork.Get()) {
		var ret int32
		return ret
	}
	return *o.ProviderNetwork.Get()
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTerminationRequest) GetProviderNetworkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderNetwork.Get(), o.ProviderNetwork.IsSet()
}

// HasProviderNetwork returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasProviderNetwork() bool {
	if o != nil && o.ProviderNetwork.IsSet() {
		return true
	}

	return false
}

// SetProviderNetwork gets a reference to the given NullableInt32 and assigns it to the ProviderNetwork field.
func (o *PatchedWritableCircuitTerminationRequest) SetProviderNetwork(v int32) {
	o.ProviderNetwork.Set(&v)
}

// SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil
func (o *PatchedWritableCircuitTerminationRequest) SetProviderNetworkNil() {
	o.ProviderNetwork.Set(nil)
}

// UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
func (o *PatchedWritableCircuitTerminationRequest) UnsetProviderNetwork() {
	o.ProviderNetwork.Unset()
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTerminationRequest) GetPortSpeed() int32 {
	if o == nil || IsNil(o.PortSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTerminationRequest) GetPortSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt32 and assigns it to the PortSpeed field.
func (o *PatchedWritableCircuitTerminationRequest) SetPortSpeed(v int32) {
	o.PortSpeed.Set(&v)
}

// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *PatchedWritableCircuitTerminationRequest) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *PatchedWritableCircuitTerminationRequest) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTerminationRequest) GetUpstreamSpeed() int32 {
	if o == nil || IsNil(o.UpstreamSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt32 and assigns it to the UpstreamSpeed field.
func (o *PatchedWritableCircuitTerminationRequest) SetUpstreamSpeed(v int32) {
	o.UpstreamSpeed.Set(&v)
}

// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *PatchedWritableCircuitTerminationRequest) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *PatchedWritableCircuitTerminationRequest) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTerminationRequest) GetXconnectId() string {
	if o == nil || IsNil(o.XconnectId) {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTerminationRequest) GetXconnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.XconnectId) {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasXconnectId() bool {
	if o != nil && !IsNil(o.XconnectId) {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *PatchedWritableCircuitTerminationRequest) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetPpInfo returns the PpInfo field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTerminationRequest) GetPpInfo() string {
	if o == nil || IsNil(o.PpInfo) {
		var ret string
		return ret
	}
	return *o.PpInfo
}

// GetPpInfoOk returns a tuple with the PpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTerminationRequest) GetPpInfoOk() (*string, bool) {
	if o == nil || IsNil(o.PpInfo) {
		return nil, false
	}
	return o.PpInfo, true
}

// HasPpInfo returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasPpInfo() bool {
	if o != nil && !IsNil(o.PpInfo) {
		return true
	}

	return false
}

// SetPpInfo gets a reference to the given string and assigns it to the PpInfo field.
func (o *PatchedWritableCircuitTerminationRequest) SetPpInfo(v string) {
	o.PpInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTerminationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTerminationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableCircuitTerminationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTerminationRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTerminationRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PatchedWritableCircuitTerminationRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableCircuitTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableCircuitTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableCircuitTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Circuit) {
		toSerialize["circuit"] = o.Circuit
	}
	if !IsNil(o.TermSide) {
		toSerialize["term_side"] = o.TermSide
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.ProviderNetwork.IsSet() {
		toSerialize["provider_network"] = o.ProviderNetwork.Get()
	}
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if !IsNil(o.XconnectId) {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if !IsNil(o.PpInfo) {
		toSerialize["pp_info"] = o.PpInfo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullablePatchedWritableCircuitTerminationRequest struct {
	value *PatchedWritableCircuitTerminationRequest
	isSet bool
}

func (v NullablePatchedWritableCircuitTerminationRequest) Get() *PatchedWritableCircuitTerminationRequest {
	return v.value
}

func (v *NullablePatchedWritableCircuitTerminationRequest) Set(val *PatchedWritableCircuitTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCircuitTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCircuitTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCircuitTerminationRequest(val *PatchedWritableCircuitTerminationRequest) *NullablePatchedWritableCircuitTerminationRequest {
	return &NullablePatchedWritableCircuitTerminationRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCircuitTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
