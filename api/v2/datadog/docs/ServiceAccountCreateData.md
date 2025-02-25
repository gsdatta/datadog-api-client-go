# ServiceAccountCreateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**ServiceAccountCreateAttributes**](ServiceAccountCreateAttributes.md) |  | 
**Relationships** | Pointer to [**UserRelationships**](UserRelationships.md) |  | [optional] 
**Type** | [**UsersType**](UsersType.md) |  | [default to USERSTYPE_USERS]

## Methods

### NewServiceAccountCreateData

`func NewServiceAccountCreateData(attributes ServiceAccountCreateAttributes, type_ UsersType) *ServiceAccountCreateData`

NewServiceAccountCreateData instantiates a new ServiceAccountCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServiceAccountCreateDataWithDefaults

`func NewServiceAccountCreateDataWithDefaults() *ServiceAccountCreateData`

NewServiceAccountCreateDataWithDefaults instantiates a new ServiceAccountCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *ServiceAccountCreateData) GetAttributes() ServiceAccountCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceAccountCreateData) GetAttributesOk() (*ServiceAccountCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceAccountCreateData) SetAttributes(v ServiceAccountCreateAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ServiceAccountCreateData) GetRelationships() UserRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceAccountCreateData) GetRelationshipsOk() (*UserRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceAccountCreateData) SetRelationships(v UserRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceAccountCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *ServiceAccountCreateData) GetType() UsersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceAccountCreateData) GetTypeOk() (*UsersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceAccountCreateData) SetType(v UsersType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


