# WritableCableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | * &#x60;cat3&#x60; - CAT3 * &#x60;cat5&#x60; - CAT5 * &#x60;cat5e&#x60; - CAT5e * &#x60;cat6&#x60; - CAT6 * &#x60;cat6a&#x60; - CAT6a * &#x60;cat7&#x60; - CAT7 * &#x60;cat7a&#x60; - CAT7a * &#x60;cat8&#x60; - CAT8 * &#x60;dac-active&#x60; - Direct Attach Copper (Active) * &#x60;dac-passive&#x60; - Direct Attach Copper (Passive) * &#x60;mrj21-trunk&#x60; - MRJ21 Trunk * &#x60;coaxial&#x60; - Coaxial * &#x60;mmf&#x60; - Multimode Fiber * &#x60;mmf-om1&#x60; - Multimode Fiber (OM1) * &#x60;mmf-om2&#x60; - Multimode Fiber (OM2) * &#x60;mmf-om3&#x60; - Multimode Fiber (OM3) * &#x60;mmf-om4&#x60; - Multimode Fiber (OM4) * &#x60;mmf-om5&#x60; - Multimode Fiber (OM5) * &#x60;smf&#x60; - Singlemode Fiber * &#x60;smf-os1&#x60; - Singlemode Fiber (OS1) * &#x60;smf-os2&#x60; - Singlemode Fiber (OS2) * &#x60;aoc&#x60; - Active Optical Cabling (AOC) * &#x60;power&#x60; - Power | [optional] 
**ATerminations** | Pointer to [**[]GenericObjectRequest**](GenericObjectRequest.md) |  | [optional] 
**BTerminations** | Pointer to [**[]GenericObjectRequest**](GenericObjectRequest.md) |  | [optional] 
**Status** | Pointer to **string** | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | [optional] 
**Tenant** | Pointer to **NullableInt32** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **NullableFloat64** |  | [optional] 
**LengthUnit** | Pointer to **string** | * &#x60;km&#x60; - Kilometers * &#x60;m&#x60; - Meters * &#x60;cm&#x60; - Centimeters * &#x60;mi&#x60; - Miles * &#x60;ft&#x60; - Feet * &#x60;in&#x60; - Inches | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableCableRequest

`func NewWritableCableRequest() *WritableCableRequest`

NewWritableCableRequest instantiates a new WritableCableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCableRequestWithDefaults

`func NewWritableCableRequestWithDefaults() *WritableCableRequest`

NewWritableCableRequestWithDefaults instantiates a new WritableCableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WritableCableRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCableRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCableRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WritableCableRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetATerminations

`func (o *WritableCableRequest) GetATerminations() []GenericObjectRequest`

GetATerminations returns the ATerminations field if non-nil, zero value otherwise.

### GetATerminationsOk

`func (o *WritableCableRequest) GetATerminationsOk() (*[]GenericObjectRequest, bool)`

GetATerminationsOk returns a tuple with the ATerminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetATerminations

`func (o *WritableCableRequest) SetATerminations(v []GenericObjectRequest)`

SetATerminations sets ATerminations field to given value.

### HasATerminations

`func (o *WritableCableRequest) HasATerminations() bool`

HasATerminations returns a boolean if a field has been set.

### GetBTerminations

`func (o *WritableCableRequest) GetBTerminations() []GenericObjectRequest`

GetBTerminations returns the BTerminations field if non-nil, zero value otherwise.

### GetBTerminationsOk

`func (o *WritableCableRequest) GetBTerminationsOk() (*[]GenericObjectRequest, bool)`

GetBTerminationsOk returns a tuple with the BTerminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTerminations

`func (o *WritableCableRequest) SetBTerminations(v []GenericObjectRequest)`

SetBTerminations sets BTerminations field to given value.

### HasBTerminations

`func (o *WritableCableRequest) HasBTerminations() bool`

HasBTerminations returns a boolean if a field has been set.

### GetStatus

`func (o *WritableCableRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableCableRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableCableRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableCableRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *WritableCableRequest) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableCableRequest) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableCableRequest) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableCableRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableCableRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableCableRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLabel

`func (o *WritableCableRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableCableRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableCableRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableCableRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *WritableCableRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WritableCableRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WritableCableRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WritableCableRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *WritableCableRequest) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *WritableCableRequest) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *WritableCableRequest) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *WritableCableRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *WritableCableRequest) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *WritableCableRequest) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *WritableCableRequest) GetLengthUnit() string`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *WritableCableRequest) GetLengthUnitOk() (*string, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *WritableCableRequest) SetLengthUnit(v string)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *WritableCableRequest) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetDescription

`func (o *WritableCableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableCableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableCableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableCableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableCableRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableCableRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableCableRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableCableRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableCableRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCableRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCableRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCableRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCableRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCableRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCableRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCableRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


