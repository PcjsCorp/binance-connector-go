# QueryTransferListResponseTransfersInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TransferId** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**WalletAddress** | Pointer to **string** |  | [optional] 
**FromToken** | Pointer to **string** |  | [optional] 
**FromTokenAmount** | Pointer to **string** |  | [optional] 
**ToToken** | Pointer to **string** |  | [optional] 
**ToTokenAmount** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**CompleteAt** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryTransferListResponseTransfersInner

`func NewQueryTransferListResponseTransfersInner() *QueryTransferListResponseTransfersInner`

NewQueryTransferListResponseTransfersInner instantiates a new QueryTransferListResponseTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTransferListResponseTransfersInnerWithDefaults

`func NewQueryTransferListResponseTransfersInnerWithDefaults() *QueryTransferListResponseTransfersInner`

NewQueryTransferListResponseTransfersInnerWithDefaults instantiates a new QueryTransferListResponseTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferId

`func (o *QueryTransferListResponseTransfersInner) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *QueryTransferListResponseTransfersInner) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *QueryTransferListResponseTransfersInner) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *QueryTransferListResponseTransfersInner) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### GetDirection

`func (o *QueryTransferListResponseTransfersInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QueryTransferListResponseTransfersInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QueryTransferListResponseTransfersInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *QueryTransferListResponseTransfersInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetStatus

`func (o *QueryTransferListResponseTransfersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryTransferListResponseTransfersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryTransferListResponseTransfersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryTransferListResponseTransfersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWalletAddress

`func (o *QueryTransferListResponseTransfersInner) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *QueryTransferListResponseTransfersInner) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *QueryTransferListResponseTransfersInner) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *QueryTransferListResponseTransfersInner) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetFromToken

`func (o *QueryTransferListResponseTransfersInner) GetFromToken() string`

GetFromToken returns the FromToken field if non-nil, zero value otherwise.

### GetFromTokenOk

`func (o *QueryTransferListResponseTransfersInner) GetFromTokenOk() (*string, bool)`

GetFromTokenOk returns a tuple with the FromToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromToken

`func (o *QueryTransferListResponseTransfersInner) SetFromToken(v string)`

SetFromToken sets FromToken field to given value.

### HasFromToken

`func (o *QueryTransferListResponseTransfersInner) HasFromToken() bool`

HasFromToken returns a boolean if a field has been set.

### GetFromTokenAmount

`func (o *QueryTransferListResponseTransfersInner) GetFromTokenAmount() string`

GetFromTokenAmount returns the FromTokenAmount field if non-nil, zero value otherwise.

### GetFromTokenAmountOk

`func (o *QueryTransferListResponseTransfersInner) GetFromTokenAmountOk() (*string, bool)`

GetFromTokenAmountOk returns a tuple with the FromTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTokenAmount

`func (o *QueryTransferListResponseTransfersInner) SetFromTokenAmount(v string)`

SetFromTokenAmount sets FromTokenAmount field to given value.

### HasFromTokenAmount

`func (o *QueryTransferListResponseTransfersInner) HasFromTokenAmount() bool`

HasFromTokenAmount returns a boolean if a field has been set.

### GetToToken

`func (o *QueryTransferListResponseTransfersInner) GetToToken() string`

GetToToken returns the ToToken field if non-nil, zero value otherwise.

### GetToTokenOk

`func (o *QueryTransferListResponseTransfersInner) GetToTokenOk() (*string, bool)`

GetToTokenOk returns a tuple with the ToToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToToken

`func (o *QueryTransferListResponseTransfersInner) SetToToken(v string)`

SetToToken sets ToToken field to given value.

### HasToToken

`func (o *QueryTransferListResponseTransfersInner) HasToToken() bool`

HasToToken returns a boolean if a field has been set.

### GetToTokenAmount

`func (o *QueryTransferListResponseTransfersInner) GetToTokenAmount() string`

GetToTokenAmount returns the ToTokenAmount field if non-nil, zero value otherwise.

### GetToTokenAmountOk

`func (o *QueryTransferListResponseTransfersInner) GetToTokenAmountOk() (*string, bool)`

GetToTokenAmountOk returns a tuple with the ToTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTokenAmount

`func (o *QueryTransferListResponseTransfersInner) SetToTokenAmount(v string)`

SetToTokenAmount sets ToTokenAmount field to given value.

### HasToTokenAmount

`func (o *QueryTransferListResponseTransfersInner) HasToTokenAmount() bool`

HasToTokenAmount returns a boolean if a field has been set.

### GetErrorCode

`func (o *QueryTransferListResponseTransfersInner) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *QueryTransferListResponseTransfersInner) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *QueryTransferListResponseTransfersInner) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *QueryTransferListResponseTransfersInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *QueryTransferListResponseTransfersInner) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *QueryTransferListResponseTransfersInner) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *QueryTransferListResponseTransfersInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *QueryTransferListResponseTransfersInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *QueryTransferListResponseTransfersInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *QueryTransferListResponseTransfersInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *QueryTransferListResponseTransfersInner) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *QueryTransferListResponseTransfersInner) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCreateTime

`func (o *QueryTransferListResponseTransfersInner) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryTransferListResponseTransfersInner) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryTransferListResponseTransfersInner) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryTransferListResponseTransfersInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryTransferListResponseTransfersInner) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryTransferListResponseTransfersInner) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryTransferListResponseTransfersInner) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryTransferListResponseTransfersInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetCompleteAt

`func (o *QueryTransferListResponseTransfersInner) GetCompleteAt() string`

GetCompleteAt returns the CompleteAt field if non-nil, zero value otherwise.

### GetCompleteAtOk

`func (o *QueryTransferListResponseTransfersInner) GetCompleteAtOk() (*string, bool)`

GetCompleteAtOk returns a tuple with the CompleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteAt

`func (o *QueryTransferListResponseTransfersInner) SetCompleteAt(v string)`

SetCompleteAt sets CompleteAt field to given value.

### HasCompleteAt

`func (o *QueryTransferListResponseTransfersInner) HasCompleteAt() bool`

HasCompleteAt returns a boolean if a field has been set.


[[Back to README]](../README.md)


