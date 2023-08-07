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

// checks if the DeviceWithConfigContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceWithConfigContext{}

// DeviceWithConfigContext Adds support for custom fields and tags.
type DeviceWithConfigContext struct {
	Id         int32                  `json:"id"`
	Url        string                 `json:"url"`
	Display    string                 `json:"display"`
	Name       NullableString         `json:"name,omitempty"`
	DeviceType NestedDeviceType       `json:"device_type"`
	DeviceRole NestedDeviceRole       `json:"device_role"`
	Tenant     NullableNestedTenant   `json:"tenant,omitempty"`
	Platform   NullableNestedPlatform `json:"platform,omitempty"`
	// Chassis serial number, assigned by the manufacturer
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag       NullableString               `json:"asset_tag,omitempty"`
	Site           NestedSite                   `json:"site"`
	Location       NullableNestedLocation       `json:"location,omitempty"`
	Rack           NullableNestedRack           `json:"rack,omitempty"`
	Position       NullableFloat64              `json:"position,omitempty"`
	Face           *CableStatus                 `json:"face,omitempty"`
	ParentDevice   NestedDevice                 `json:"parent_device"`
	Status         *CableStatus                 `json:"status,omitempty"`
	Airflow        *CableStatus                 `json:"airflow,omitempty"`
	PrimaryIp      NestedIPAddress              `json:"primary_ip"`
	PrimaryIp4     NullableNestedIPAddress      `json:"primary_ip4,omitempty"`
	PrimaryIp6     NullableNestedIPAddress      `json:"primary_ip6,omitempty"`
	Cluster        NullableNestedCluster        `json:"cluster,omitempty"`
	VirtualChassis NullableNestedVirtualChassis `json:"virtual_chassis,omitempty"`
	VcPosition     NullableInt32                `json:"vc_position,omitempty"`
	// Virtual chassis master election priority
	VcPriority  NullableInt32 `json:"vc_priority,omitempty"`
	Description *string       `json:"description,omitempty"`
	Comments    *string       `json:"comments,omitempty"`
	// Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData map[string]interface{} `json:"local_context_data,omitempty"`
	Tags             []NestedTag            `json:"tags,omitempty"`
	CustomFields     map[string]interface{} `json:"custom_fields,omitempty"`
	ConfigContext    map[string]interface{} `json:"config_context"`
	Created          NullableTime           `json:"created"`
	LastUpdated      NullableTime           `json:"last_updated"`
}

// NewDeviceWithConfigContext instantiates a new DeviceWithConfigContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceWithConfigContext(id int32, url string, display string, deviceType NestedDeviceType, deviceRole NestedDeviceRole, site NestedSite, parentDevice NestedDevice, primaryIp NestedIPAddress, configContext map[string]interface{}, created NullableTime, lastUpdated NullableTime) *DeviceWithConfigContext {
	this := DeviceWithConfigContext{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.DeviceType = deviceType
	this.DeviceRole = deviceRole
	this.Site = site
	this.ParentDevice = parentDevice
	this.PrimaryIp = primaryIp
	this.ConfigContext = configContext
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewDeviceWithConfigContextWithDefaults instantiates a new DeviceWithConfigContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithConfigContextWithDefaults() *DeviceWithConfigContext {
	this := DeviceWithConfigContext{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceWithConfigContext) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceWithConfigContext) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DeviceWithConfigContext) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeviceWithConfigContext) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *DeviceWithConfigContext) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DeviceWithConfigContext) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DeviceWithConfigContext) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *DeviceWithConfigContext) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetName() {
	o.Name.Unset()
}

// GetDeviceType returns the DeviceType field value
func (o *DeviceWithConfigContext) GetDeviceType() NestedDeviceType {
	if o == nil {
		var ret NestedDeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetDeviceTypeOk() (*NestedDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *DeviceWithConfigContext) SetDeviceType(v NestedDeviceType) {
	o.DeviceType = v
}

// GetDeviceRole returns the DeviceRole field value
func (o *DeviceWithConfigContext) GetDeviceRole() NestedDeviceRole {
	if o == nil {
		var ret NestedDeviceRole
		return ret
	}

	return o.DeviceRole
}

// GetDeviceRoleOk returns a tuple with the DeviceRole field value
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetDeviceRoleOk() (*NestedDeviceRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceRole, true
}

// SetDeviceRole sets field value
func (o *DeviceWithConfigContext) SetDeviceRole(v NestedDeviceRole) {
	o.DeviceRole = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetTenant() NestedTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret NestedTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetTenantOk() (*NestedTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenant and assigns it to the Tenant field.
func (o *DeviceWithConfigContext) SetTenant(v NestedTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *DeviceWithConfigContext) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetPlatform() NestedPlatform {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret NestedPlatform
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetPlatformOk() (*NestedPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableNestedPlatform and assigns it to the Platform field.
func (o *DeviceWithConfigContext) SetPlatform(v NestedPlatform) {
	o.Platform.Set(&v)
}

// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *DeviceWithConfigContext) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetPlatform() {
	o.Platform.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *DeviceWithConfigContext) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *DeviceWithConfigContext) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *DeviceWithConfigContext) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *DeviceWithConfigContext) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetSite returns the Site field value
func (o *DeviceWithConfigContext) GetSite() NestedSite {
	if o == nil {
		var ret NestedSite
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetSiteOk() (*NestedSite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *DeviceWithConfigContext) SetSite(v NestedSite) {
	o.Site = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetLocation() NestedLocation {
	if o == nil || IsNil(o.Location.Get()) {
		var ret NestedLocation
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetLocationOk() (*NestedLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableNestedLocation and assigns it to the Location field.
func (o *DeviceWithConfigContext) SetLocation(v NestedLocation) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *DeviceWithConfigContext) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetLocation() {
	o.Location.Unset()
}

// GetRack returns the Rack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetRack() NestedRack {
	if o == nil || IsNil(o.Rack.Get()) {
		var ret NestedRack
		return ret
	}
	return *o.Rack.Get()
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetRackOk() (*NestedRack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rack.Get(), o.Rack.IsSet()
}

// HasRack returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasRack() bool {
	if o != nil && o.Rack.IsSet() {
		return true
	}

	return false
}

// SetRack gets a reference to the given NullableNestedRack and assigns it to the Rack field.
func (o *DeviceWithConfigContext) SetRack(v NestedRack) {
	o.Rack.Set(&v)
}

// SetRackNil sets the value for Rack to be an explicit nil
func (o *DeviceWithConfigContext) SetRackNil() {
	o.Rack.Set(nil)
}

// UnsetRack ensures that no value is present for Rack, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetRack() {
	o.Rack.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetPosition() float64 {
	if o == nil || IsNil(o.Position.Get()) {
		var ret float64
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetPositionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableFloat64 and assigns it to the Position field.
func (o *DeviceWithConfigContext) SetPosition(v float64) {
	o.Position.Set(&v)
}

// SetPositionNil sets the value for Position to be an explicit nil
func (o *DeviceWithConfigContext) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetPosition() {
	o.Position.Unset()
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *DeviceWithConfigContext) GetFace() CableStatus {
	if o == nil || IsNil(o.Face) {
		var ret CableStatus
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetFaceOk() (*CableStatus, bool) {
	if o == nil || IsNil(o.Face) {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasFace() bool {
	if o != nil && !IsNil(o.Face) {
		return true
	}

	return false
}

// SetFace gets a reference to the given CableStatus and assigns it to the Face field.
func (o *DeviceWithConfigContext) SetFace(v CableStatus) {
	o.Face = &v
}

// GetParentDevice returns the ParentDevice field value
func (o *DeviceWithConfigContext) GetParentDevice() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.ParentDevice
}

// GetParentDeviceOk returns a tuple with the ParentDevice field value
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetParentDeviceOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentDevice, true
}

// SetParentDevice sets field value
func (o *DeviceWithConfigContext) SetParentDevice(v NestedDevice) {
	o.ParentDevice = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceWithConfigContext) GetStatus() CableStatus {
	if o == nil || IsNil(o.Status) {
		var ret CableStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetStatusOk() (*CableStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CableStatus and assigns it to the Status field.
func (o *DeviceWithConfigContext) SetStatus(v CableStatus) {
	o.Status = &v
}

// GetAirflow returns the Airflow field value if set, zero value otherwise.
func (o *DeviceWithConfigContext) GetAirflow() CableStatus {
	if o == nil || IsNil(o.Airflow) {
		var ret CableStatus
		return ret
	}
	return *o.Airflow
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetAirflowOk() (*CableStatus, bool) {
	if o == nil || IsNil(o.Airflow) {
		return nil, false
	}
	return o.Airflow, true
}

// HasAirflow returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasAirflow() bool {
	if o != nil && !IsNil(o.Airflow) {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given CableStatus and assigns it to the Airflow field.
func (o *DeviceWithConfigContext) SetAirflow(v CableStatus) {
	o.Airflow = &v
}

// GetPrimaryIp returns the PrimaryIp field value
func (o *DeviceWithConfigContext) GetPrimaryIp() NestedIPAddress {
	if o == nil {
		var ret NestedIPAddress
		return ret
	}

	return o.PrimaryIp
}

// GetPrimaryIpOk returns a tuple with the PrimaryIp field value
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetPrimaryIpOk() (*NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryIp, true
}

// SetPrimaryIp sets field value
func (o *DeviceWithConfigContext) SetPrimaryIp(v NestedIPAddress) {
	o.PrimaryIp = v
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetPrimaryIp4() NestedIPAddress {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret NestedIPAddress
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetPrimaryIp4Ok() (*NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableNestedIPAddress and assigns it to the PrimaryIp4 field.
func (o *DeviceWithConfigContext) SetPrimaryIp4(v NestedIPAddress) {
	o.PrimaryIp4.Set(&v)
}

// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *DeviceWithConfigContext) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetPrimaryIp6() NestedIPAddress {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret NestedIPAddress
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetPrimaryIp6Ok() (*NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableNestedIPAddress and assigns it to the PrimaryIp6 field.
func (o *DeviceWithConfigContext) SetPrimaryIp6(v NestedIPAddress) {
	o.PrimaryIp6.Set(&v)
}

// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *DeviceWithConfigContext) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetCluster() NestedCluster {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret NestedCluster
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetClusterOk() (*NestedCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableNestedCluster and assigns it to the Cluster field.
func (o *DeviceWithConfigContext) SetCluster(v NestedCluster) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *DeviceWithConfigContext) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetCluster() {
	o.Cluster.Unset()
}

// GetVirtualChassis returns the VirtualChassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetVirtualChassis() NestedVirtualChassis {
	if o == nil || IsNil(o.VirtualChassis.Get()) {
		var ret NestedVirtualChassis
		return ret
	}
	return *o.VirtualChassis.Get()
}

// GetVirtualChassisOk returns a tuple with the VirtualChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetVirtualChassisOk() (*NestedVirtualChassis, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualChassis.Get(), o.VirtualChassis.IsSet()
}

// HasVirtualChassis returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasVirtualChassis() bool {
	if o != nil && o.VirtualChassis.IsSet() {
		return true
	}

	return false
}

// SetVirtualChassis gets a reference to the given NullableNestedVirtualChassis and assigns it to the VirtualChassis field.
func (o *DeviceWithConfigContext) SetVirtualChassis(v NestedVirtualChassis) {
	o.VirtualChassis.Set(&v)
}

// SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil
func (o *DeviceWithConfigContext) SetVirtualChassisNil() {
	o.VirtualChassis.Set(nil)
}

// UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetVirtualChassis() {
	o.VirtualChassis.Unset()
}

// GetVcPosition returns the VcPosition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetVcPosition() int32 {
	if o == nil || IsNil(o.VcPosition.Get()) {
		var ret int32
		return ret
	}
	return *o.VcPosition.Get()
}

// GetVcPositionOk returns a tuple with the VcPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetVcPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcPosition.Get(), o.VcPosition.IsSet()
}

// HasVcPosition returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasVcPosition() bool {
	if o != nil && o.VcPosition.IsSet() {
		return true
	}

	return false
}

// SetVcPosition gets a reference to the given NullableInt32 and assigns it to the VcPosition field.
func (o *DeviceWithConfigContext) SetVcPosition(v int32) {
	o.VcPosition.Set(&v)
}

// SetVcPositionNil sets the value for VcPosition to be an explicit nil
func (o *DeviceWithConfigContext) SetVcPositionNil() {
	o.VcPosition.Set(nil)
}

// UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetVcPosition() {
	o.VcPosition.Unset()
}

// GetVcPriority returns the VcPriority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetVcPriority() int32 {
	if o == nil || IsNil(o.VcPriority.Get()) {
		var ret int32
		return ret
	}
	return *o.VcPriority.Get()
}

// GetVcPriorityOk returns a tuple with the VcPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetVcPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcPriority.Get(), o.VcPriority.IsSet()
}

// HasVcPriority returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasVcPriority() bool {
	if o != nil && o.VcPriority.IsSet() {
		return true
	}

	return false
}

// SetVcPriority gets a reference to the given NullableInt32 and assigns it to the VcPriority field.
func (o *DeviceWithConfigContext) SetVcPriority(v int32) {
	o.VcPriority.Set(&v)
}

// SetVcPriorityNil sets the value for VcPriority to be an explicit nil
func (o *DeviceWithConfigContext) SetVcPriorityNil() {
	o.VcPriority.Set(nil)
}

// UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
func (o *DeviceWithConfigContext) UnsetVcPriority() {
	o.VcPriority.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceWithConfigContext) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceWithConfigContext) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *DeviceWithConfigContext) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *DeviceWithConfigContext) SetComments(v string) {
	o.Comments = &v
}

// GetLocalContextData returns the LocalContextData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceWithConfigContext) GetLocalContextData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.LocalContextData
}

// GetLocalContextDataOk returns a tuple with the LocalContextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetLocalContextDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LocalContextData) {
		return map[string]interface{}{}, false
	}
	return o.LocalContextData, true
}

// HasLocalContextData returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasLocalContextData() bool {
	if o != nil && IsNil(o.LocalContextData) {
		return true
	}

	return false
}

// SetLocalContextData gets a reference to the given map[string]interface{} and assigns it to the LocalContextData field.
func (o *DeviceWithConfigContext) SetLocalContextData(v map[string]interface{}) {
	o.LocalContextData = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceWithConfigContext) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *DeviceWithConfigContext) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceWithConfigContext) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithConfigContext) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceWithConfigContext) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceWithConfigContext) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetConfigContext returns the ConfigContext field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *DeviceWithConfigContext) GetConfigContext() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ConfigContext
}

// GetConfigContextOk returns a tuple with the ConfigContext field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetConfigContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConfigContext) {
		return map[string]interface{}{}, false
	}
	return o.ConfigContext, true
}

// SetConfigContext sets field value
func (o *DeviceWithConfigContext) SetConfigContext(v map[string]interface{}) {
	o.ConfigContext = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceWithConfigContext) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *DeviceWithConfigContext) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceWithConfigContext) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceWithConfigContext) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *DeviceWithConfigContext) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o DeviceWithConfigContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceWithConfigContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["device_type"] = o.DeviceType
	toSerialize["device_role"] = o.DeviceRole
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	toSerialize["site"] = o.Site
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Rack.IsSet() {
		toSerialize["rack"] = o.Rack.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if !IsNil(o.Face) {
		toSerialize["face"] = o.Face
	}
	// skip: parent_device is readOnly
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Airflow) {
		toSerialize["airflow"] = o.Airflow
	}
	// skip: primary_ip is readOnly
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	if o.Cluster.IsSet() {
		toSerialize["cluster"] = o.Cluster.Get()
	}
	if o.VirtualChassis.IsSet() {
		toSerialize["virtual_chassis"] = o.VirtualChassis.Get()
	}
	if o.VcPosition.IsSet() {
		toSerialize["vc_position"] = o.VcPosition.Get()
	}
	if o.VcPriority.IsSet() {
		toSerialize["vc_priority"] = o.VcPriority.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.LocalContextData != nil {
		toSerialize["local_context_data"] = o.LocalContextData
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.ConfigContext != nil {
		toSerialize["config_context"] = o.ConfigContext
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	return toSerialize, nil
}

type NullableDeviceWithConfigContext struct {
	value *DeviceWithConfigContext
	isSet bool
}

func (v NullableDeviceWithConfigContext) Get() *DeviceWithConfigContext {
	return v.value
}

func (v *NullableDeviceWithConfigContext) Set(val *DeviceWithConfigContext) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceWithConfigContext) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceWithConfigContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceWithConfigContext(val *DeviceWithConfigContext) *NullableDeviceWithConfigContext {
	return &NullableDeviceWithConfigContext{value: val, isSet: true}
}

func (v NullableDeviceWithConfigContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceWithConfigContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
