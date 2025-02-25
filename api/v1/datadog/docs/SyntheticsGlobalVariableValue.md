# SyntheticsGlobalVariableValue

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Secure** | Pointer to **bool** | Determines if the value of the variable is hidden. | [optional] 
**Value** | Pointer to **string** | Value of the global variable. When reading a global variable, the value will not be present if the variable is hidden with the &#x60;secure&#x60; property. | [optional] 

## Methods

### NewSyntheticsGlobalVariableValue

`func NewSyntheticsGlobalVariableValue() *SyntheticsGlobalVariableValue`

NewSyntheticsGlobalVariableValue instantiates a new SyntheticsGlobalVariableValue object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsGlobalVariableValueWithDefaults

`func NewSyntheticsGlobalVariableValueWithDefaults() *SyntheticsGlobalVariableValue`

NewSyntheticsGlobalVariableValueWithDefaults instantiates a new SyntheticsGlobalVariableValue object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetSecure

`func (o *SyntheticsGlobalVariableValue) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *SyntheticsGlobalVariableValue) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *SyntheticsGlobalVariableValue) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *SyntheticsGlobalVariableValue) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetValue

`func (o *SyntheticsGlobalVariableValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SyntheticsGlobalVariableValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SyntheticsGlobalVariableValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SyntheticsGlobalVariableValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


