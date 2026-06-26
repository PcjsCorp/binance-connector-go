# QueryTransferListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Transfers** | Pointer to [**[]QueryTransferListResponseTransfersInner**](QueryTransferListResponseTransfersInner.md) |  | [optional] 

## Methods

### NewQueryTransferListResponse

`func NewQueryTransferListResponse() *QueryTransferListResponse`

NewQueryTransferListResponse instantiates a new QueryTransferListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTransferListResponseWithDefaults

`func NewQueryTransferListResponseWithDefaults() *QueryTransferListResponse`

NewQueryTransferListResponseWithDefaults instantiates a new QueryTransferListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransfers

`func (o *QueryTransferListResponse) GetTransfers() []QueryTransferListResponseTransfersInner`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *QueryTransferListResponse) GetTransfersOk() (*[]QueryTransferListResponseTransfersInner, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *QueryTransferListResponse) SetTransfers(v []QueryTransferListResponseTransfersInner)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *QueryTransferListResponse) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.


[[Back to README]](../README.md)


