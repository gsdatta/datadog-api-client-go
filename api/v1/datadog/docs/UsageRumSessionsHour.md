# UsageRumSessionsHour

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**SessionCount** | Pointer to **int64** | Contains the number of RUM Sessions. | [optional] 
**SessionCountAndroid** | Pointer to **int64** | Contains the number of mobile RUM Sessions on Android (data available beginning December 1, 2020). | [optional] 
**SessionCountIos** | Pointer to **int64** | Contains the number of mobile RUM Sessions on iOS (data available beginning December 1, 2020). | [optional] 

## Methods

### NewUsageRumSessionsHour

`func NewUsageRumSessionsHour() *UsageRumSessionsHour`

NewUsageRumSessionsHour instantiates a new UsageRumSessionsHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageRumSessionsHourWithDefaults

`func NewUsageRumSessionsHourWithDefaults() *UsageRumSessionsHour`

NewUsageRumSessionsHourWithDefaults instantiates a new UsageRumSessionsHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHour

`func (o *UsageRumSessionsHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageRumSessionsHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageRumSessionsHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageRumSessionsHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetSessionCount

`func (o *UsageRumSessionsHour) GetSessionCount() int64`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *UsageRumSessionsHour) GetSessionCountOk() (*int64, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *UsageRumSessionsHour) SetSessionCount(v int64)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *UsageRumSessionsHour) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSessionCountAndroid

`func (o *UsageRumSessionsHour) GetSessionCountAndroid() int64`

GetSessionCountAndroid returns the SessionCountAndroid field if non-nil, zero value otherwise.

### GetSessionCountAndroidOk

`func (o *UsageRumSessionsHour) GetSessionCountAndroidOk() (*int64, bool)`

GetSessionCountAndroidOk returns a tuple with the SessionCountAndroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCountAndroid

`func (o *UsageRumSessionsHour) SetSessionCountAndroid(v int64)`

SetSessionCountAndroid sets SessionCountAndroid field to given value.

### HasSessionCountAndroid

`func (o *UsageRumSessionsHour) HasSessionCountAndroid() bool`

HasSessionCountAndroid returns a boolean if a field has been set.

### GetSessionCountIos

`func (o *UsageRumSessionsHour) GetSessionCountIos() int64`

GetSessionCountIos returns the SessionCountIos field if non-nil, zero value otherwise.

### GetSessionCountIosOk

`func (o *UsageRumSessionsHour) GetSessionCountIosOk() (*int64, bool)`

GetSessionCountIosOk returns a tuple with the SessionCountIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCountIos

`func (o *UsageRumSessionsHour) SetSessionCountIos(v int64)`

SetSessionCountIos sets SessionCountIos field to given value.

### HasSessionCountIos

`func (o *UsageRumSessionsHour) HasSessionCountIos() bool`

HasSessionCountIos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


