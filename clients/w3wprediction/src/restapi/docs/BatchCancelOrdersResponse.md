# BatchCancelOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Canceled** | Pointer to **[]string** |  | [optional] 
**Failed** | Pointer to [**[]BatchCancelOrdersResponseFailedInner**](BatchCancelOrdersResponseFailedInner.md) |  | [optional] 

## Methods

### NewBatchCancelOrdersResponse

`func NewBatchCancelOrdersResponse() *BatchCancelOrdersResponse`

NewBatchCancelOrdersResponse instantiates a new BatchCancelOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCancelOrdersResponseWithDefaults

`func NewBatchCancelOrdersResponseWithDefaults() *BatchCancelOrdersResponse`

NewBatchCancelOrdersResponseWithDefaults instantiates a new BatchCancelOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanceled

`func (o *BatchCancelOrdersResponse) GetCanceled() []string`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *BatchCancelOrdersResponse) GetCanceledOk() (*[]string, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *BatchCancelOrdersResponse) SetCanceled(v []string)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *BatchCancelOrdersResponse) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### GetFailed

`func (o *BatchCancelOrdersResponse) GetFailed() []BatchCancelOrdersResponseFailedInner`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BatchCancelOrdersResponse) GetFailedOk() (*[]BatchCancelOrdersResponseFailedInner, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BatchCancelOrdersResponse) SetFailed(v []BatchCancelOrdersResponseFailedInner)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *BatchCancelOrdersResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to README]](../README.md)


