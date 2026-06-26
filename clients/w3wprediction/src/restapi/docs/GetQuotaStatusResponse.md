# GetQuotaStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**DailyLimit** | Pointer to **string** |  | [optional] 
**RemainingDailyLimit** | Pointer to **string** |  | [optional] 

## Methods

### NewGetQuotaStatusResponse

`func NewGetQuotaStatusResponse() *GetQuotaStatusResponse`

NewGetQuotaStatusResponse instantiates a new GetQuotaStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuotaStatusResponseWithDefaults

`func NewGetQuotaStatusResponseWithDefaults() *GetQuotaStatusResponse`

NewGetQuotaStatusResponseWithDefaults instantiates a new GetQuotaStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyLimit

`func (o *GetQuotaStatusResponse) GetDailyLimit() string`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *GetQuotaStatusResponse) GetDailyLimitOk() (*string, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *GetQuotaStatusResponse) SetDailyLimit(v string)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *GetQuotaStatusResponse) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetRemainingDailyLimit

`func (o *GetQuotaStatusResponse) GetRemainingDailyLimit() string`

GetRemainingDailyLimit returns the RemainingDailyLimit field if non-nil, zero value otherwise.

### GetRemainingDailyLimitOk

`func (o *GetQuotaStatusResponse) GetRemainingDailyLimitOk() (*string, bool)`

GetRemainingDailyLimitOk returns a tuple with the RemainingDailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingDailyLimit

`func (o *GetQuotaStatusResponse) SetRemainingDailyLimit(v string)`

SetRemainingDailyLimit sets RemainingDailyLimit field to given value.

### HasRemainingDailyLimit

`func (o *GetQuotaStatusResponse) HasRemainingDailyLimit() bool`

HasRemainingDailyLimit returns a boolean if a field has been set.


[[Back to README]](../README.md)


