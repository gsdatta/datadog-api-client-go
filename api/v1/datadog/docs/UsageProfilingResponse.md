# UsageProfilingResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Usage** | Pointer to [**[]UsageProfilingHour**](UsageProfilingHour.md) | Get hourly usage for profiled hosts. | [optional] 

## Methods

### NewUsageProfilingResponse

`func NewUsageProfilingResponse() *UsageProfilingResponse`

NewUsageProfilingResponse instantiates a new UsageProfilingResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageProfilingResponseWithDefaults

`func NewUsageProfilingResponseWithDefaults() *UsageProfilingResponse`

NewUsageProfilingResponseWithDefaults instantiates a new UsageProfilingResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUsage

`func (o *UsageProfilingResponse) GetUsage() []UsageProfilingHour`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageProfilingResponse) GetUsageOk() (*[]UsageProfilingHour, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageProfilingResponse) SetUsage(v []UsageProfilingHour)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageProfilingResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


