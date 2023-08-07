# WritablePowerPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **NullableInt32** |  | [optional] 
**ModuleType** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |          {module} is accepted as a substitution for the module bay position when attached to a module type.          | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to **string** | * &#x60;iec-60320-c6&#x60; - C6 * &#x60;iec-60320-c8&#x60; - C8 * &#x60;iec-60320-c14&#x60; - C14 * &#x60;iec-60320-c16&#x60; - C16 * &#x60;iec-60320-c20&#x60; - C20 * &#x60;iec-60320-c22&#x60; - C22 * &#x60;iec-60309-p-n-e-4h&#x60; - P+N+E 4H * &#x60;iec-60309-p-n-e-6h&#x60; - P+N+E 6H * &#x60;iec-60309-p-n-e-9h&#x60; - P+N+E 9H * &#x60;iec-60309-2p-e-4h&#x60; - 2P+E 4H * &#x60;iec-60309-2p-e-6h&#x60; - 2P+E 6H * &#x60;iec-60309-2p-e-9h&#x60; - 2P+E 9H * &#x60;iec-60309-3p-e-4h&#x60; - 3P+E 4H * &#x60;iec-60309-3p-e-6h&#x60; - 3P+E 6H * &#x60;iec-60309-3p-e-9h&#x60; - 3P+E 9H * &#x60;iec-60309-3p-n-e-4h&#x60; - 3P+N+E 4H * &#x60;iec-60309-3p-n-e-6h&#x60; - 3P+N+E 6H * &#x60;iec-60309-3p-n-e-9h&#x60; - 3P+N+E 9H * &#x60;nema-1-15p&#x60; - NEMA 1-15P * &#x60;nema-5-15p&#x60; - NEMA 5-15P * &#x60;nema-5-20p&#x60; - NEMA 5-20P * &#x60;nema-5-30p&#x60; - NEMA 5-30P * &#x60;nema-5-50p&#x60; - NEMA 5-50P * &#x60;nema-6-15p&#x60; - NEMA 6-15P * &#x60;nema-6-20p&#x60; - NEMA 6-20P * &#x60;nema-6-30p&#x60; - NEMA 6-30P * &#x60;nema-6-50p&#x60; - NEMA 6-50P * &#x60;nema-10-30p&#x60; - NEMA 10-30P * &#x60;nema-10-50p&#x60; - NEMA 10-50P * &#x60;nema-14-20p&#x60; - NEMA 14-20P * &#x60;nema-14-30p&#x60; - NEMA 14-30P * &#x60;nema-14-50p&#x60; - NEMA 14-50P * &#x60;nema-14-60p&#x60; - NEMA 14-60P * &#x60;nema-15-15p&#x60; - NEMA 15-15P * &#x60;nema-15-20p&#x60; - NEMA 15-20P * &#x60;nema-15-30p&#x60; - NEMA 15-30P * &#x60;nema-15-50p&#x60; - NEMA 15-50P * &#x60;nema-15-60p&#x60; - NEMA 15-60P * &#x60;nema-l1-15p&#x60; - NEMA L1-15P * &#x60;nema-l5-15p&#x60; - NEMA L5-15P * &#x60;nema-l5-20p&#x60; - NEMA L5-20P * &#x60;nema-l5-30p&#x60; - NEMA L5-30P * &#x60;nema-l5-50p&#x60; - NEMA L5-50P * &#x60;nema-l6-15p&#x60; - NEMA L6-15P * &#x60;nema-l6-20p&#x60; - NEMA L6-20P * &#x60;nema-l6-30p&#x60; - NEMA L6-30P * &#x60;nema-l6-50p&#x60; - NEMA L6-50P * &#x60;nema-l10-30p&#x60; - NEMA L10-30P * &#x60;nema-l14-20p&#x60; - NEMA L14-20P * &#x60;nema-l14-30p&#x60; - NEMA L14-30P * &#x60;nema-l14-50p&#x60; - NEMA L14-50P * &#x60;nema-l14-60p&#x60; - NEMA L14-60P * &#x60;nema-l15-20p&#x60; - NEMA L15-20P * &#x60;nema-l15-30p&#x60; - NEMA L15-30P * &#x60;nema-l15-50p&#x60; - NEMA L15-50P * &#x60;nema-l15-60p&#x60; - NEMA L15-60P * &#x60;nema-l21-20p&#x60; - NEMA L21-20P * &#x60;nema-l21-30p&#x60; - NEMA L21-30P * &#x60;nema-l22-30p&#x60; - NEMA L22-30P * &#x60;cs6361c&#x60; - CS6361C * &#x60;cs6365c&#x60; - CS6365C * &#x60;cs8165c&#x60; - CS8165C * &#x60;cs8265c&#x60; - CS8265C * &#x60;cs8365c&#x60; - CS8365C * &#x60;cs8465c&#x60; - CS8465C * &#x60;ita-c&#x60; - ITA Type C (CEE 7/16) * &#x60;ita-e&#x60; - ITA Type E (CEE 7/6) * &#x60;ita-f&#x60; - ITA Type F (CEE 7/4) * &#x60;ita-ef&#x60; - ITA Type E/F (CEE 7/7) * &#x60;ita-g&#x60; - ITA Type G (BS 1363) * &#x60;ita-h&#x60; - ITA Type H * &#x60;ita-i&#x60; - ITA Type I * &#x60;ita-j&#x60; - ITA Type J * &#x60;ita-k&#x60; - ITA Type K * &#x60;ita-l&#x60; - ITA Type L (CEI 23-50) * &#x60;ita-m&#x60; - ITA Type M (BS 546) * &#x60;ita-n&#x60; - ITA Type N * &#x60;ita-o&#x60; - ITA Type O * &#x60;usb-a&#x60; - USB Type A * &#x60;usb-b&#x60; - USB Type B * &#x60;usb-c&#x60; - USB Type C * &#x60;usb-mini-a&#x60; - USB Mini A * &#x60;usb-mini-b&#x60; - USB Mini B * &#x60;usb-micro-a&#x60; - USB Micro A * &#x60;usb-micro-b&#x60; - USB Micro B * &#x60;usb-micro-ab&#x60; - USB Micro AB * &#x60;usb-3-b&#x60; - USB 3.0 Type B * &#x60;usb-3-micro-b&#x60; - USB 3.0 Micro B * &#x60;dc-terminal&#x60; - DC Terminal * &#x60;saf-d-grid&#x60; - Saf-D-Grid * &#x60;neutrik-powercon-20&#x60; - Neutrik powerCON (20A) * &#x60;neutrik-powercon-32&#x60; - Neutrik powerCON (32A) * &#x60;neutrik-powercon-true1&#x60; - Neutrik powerCON TRUE1 * &#x60;neutrik-powercon-true1-top&#x60; - Neutrik powerCON TRUE1 TOP * &#x60;ubiquiti-smartpower&#x60; - Ubiquiti SmartPower * &#x60;hardwired&#x60; - Hardwired * &#x60;other&#x60; - Other | [optional] 
**MaximumDraw** | Pointer to **NullableInt32** | Maximum power draw (watts) | [optional] 
**AllocatedDraw** | Pointer to **NullableInt32** | Allocated power draw (watts) | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewWritablePowerPortTemplateRequest

`func NewWritablePowerPortTemplateRequest(name string, ) *WritablePowerPortTemplateRequest`

NewWritablePowerPortTemplateRequest instantiates a new WritablePowerPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerPortTemplateRequestWithDefaults

`func NewWritablePowerPortTemplateRequestWithDefaults() *WritablePowerPortTemplateRequest`

NewWritablePowerPortTemplateRequestWithDefaults instantiates a new WritablePowerPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *WritablePowerPortTemplateRequest) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WritablePowerPortTemplateRequest) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WritablePowerPortTemplateRequest) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WritablePowerPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *WritablePowerPortTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *WritablePowerPortTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *WritablePowerPortTemplateRequest) GetModuleType() int32`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *WritablePowerPortTemplateRequest) GetModuleTypeOk() (*int32, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *WritablePowerPortTemplateRequest) SetModuleType(v int32)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *WritablePowerPortTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *WritablePowerPortTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *WritablePowerPortTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *WritablePowerPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritablePowerPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritablePowerPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritablePowerPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritablePowerPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritablePowerPortTemplateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePowerPortTemplateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePowerPortTemplateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePowerPortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMaximumDraw

`func (o *WritablePowerPortTemplateRequest) GetMaximumDraw() int32`

GetMaximumDraw returns the MaximumDraw field if non-nil, zero value otherwise.

### GetMaximumDrawOk

`func (o *WritablePowerPortTemplateRequest) GetMaximumDrawOk() (*int32, bool)`

GetMaximumDrawOk returns a tuple with the MaximumDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDraw

`func (o *WritablePowerPortTemplateRequest) SetMaximumDraw(v int32)`

SetMaximumDraw sets MaximumDraw field to given value.

### HasMaximumDraw

`func (o *WritablePowerPortTemplateRequest) HasMaximumDraw() bool`

HasMaximumDraw returns a boolean if a field has been set.

### SetMaximumDrawNil

`func (o *WritablePowerPortTemplateRequest) SetMaximumDrawNil(b bool)`

 SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil

### UnsetMaximumDraw
`func (o *WritablePowerPortTemplateRequest) UnsetMaximumDraw()`

UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
### GetAllocatedDraw

`func (o *WritablePowerPortTemplateRequest) GetAllocatedDraw() int32`

GetAllocatedDraw returns the AllocatedDraw field if non-nil, zero value otherwise.

### GetAllocatedDrawOk

`func (o *WritablePowerPortTemplateRequest) GetAllocatedDrawOk() (*int32, bool)`

GetAllocatedDrawOk returns a tuple with the AllocatedDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedDraw

`func (o *WritablePowerPortTemplateRequest) SetAllocatedDraw(v int32)`

SetAllocatedDraw sets AllocatedDraw field to given value.

### HasAllocatedDraw

`func (o *WritablePowerPortTemplateRequest) HasAllocatedDraw() bool`

HasAllocatedDraw returns a boolean if a field has been set.

### SetAllocatedDrawNil

`func (o *WritablePowerPortTemplateRequest) SetAllocatedDrawNil(b bool)`

 SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil

### UnsetAllocatedDraw
`func (o *WritablePowerPortTemplateRequest) UnsetAllocatedDraw()`

UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
### GetDescription

`func (o *WritablePowerPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePowerPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePowerPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePowerPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


