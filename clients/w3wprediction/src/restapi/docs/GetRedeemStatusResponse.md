# GetRedeemStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetRedeemStatusResponse

`func NewGetRedeemStatusResponse() *GetRedeemStatusResponse`

NewGetRedeemStatusResponse instantiates a new GetRedeemStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedeemStatusResponseWithDefaults

`func NewGetRedeemStatusResponseWithDefaults() *GetRedeemStatusResponse`

NewGetRedeemStatusResponseWithDefaults instantiates a new GetRedeemStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *GetRedeemStatusResponse) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *GetRedeemStatusResponse) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *GetRedeemStatusResponse) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *GetRedeemStatusResponse) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetStatus

`func (o *GetRedeemStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRedeemStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRedeemStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRedeemStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


