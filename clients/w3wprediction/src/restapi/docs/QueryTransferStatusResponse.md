# QueryTransferStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TransferId** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**FromToken** | Pointer to **string** |  | [optional] 
**FromTokenAmount** | Pointer to **string** |  | [optional] 
**ToToken** | Pointer to **string** |  | [optional] 
**ToTokenAmount** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryTransferStatusResponse

`func NewQueryTransferStatusResponse() *QueryTransferStatusResponse`

NewQueryTransferStatusResponse instantiates a new QueryTransferStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTransferStatusResponseWithDefaults

`func NewQueryTransferStatusResponseWithDefaults() *QueryTransferStatusResponse`

NewQueryTransferStatusResponseWithDefaults instantiates a new QueryTransferStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferId

`func (o *QueryTransferStatusResponse) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *QueryTransferStatusResponse) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *QueryTransferStatusResponse) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *QueryTransferStatusResponse) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### GetDirection

`func (o *QueryTransferStatusResponse) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QueryTransferStatusResponse) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QueryTransferStatusResponse) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *QueryTransferStatusResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetStatus

`func (o *QueryTransferStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryTransferStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryTransferStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryTransferStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFromToken

`func (o *QueryTransferStatusResponse) GetFromToken() string`

GetFromToken returns the FromToken field if non-nil, zero value otherwise.

### GetFromTokenOk

`func (o *QueryTransferStatusResponse) GetFromTokenOk() (*string, bool)`

GetFromTokenOk returns a tuple with the FromToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromToken

`func (o *QueryTransferStatusResponse) SetFromToken(v string)`

SetFromToken sets FromToken field to given value.

### HasFromToken

`func (o *QueryTransferStatusResponse) HasFromToken() bool`

HasFromToken returns a boolean if a field has been set.

### GetFromTokenAmount

`func (o *QueryTransferStatusResponse) GetFromTokenAmount() string`

GetFromTokenAmount returns the FromTokenAmount field if non-nil, zero value otherwise.

### GetFromTokenAmountOk

`func (o *QueryTransferStatusResponse) GetFromTokenAmountOk() (*string, bool)`

GetFromTokenAmountOk returns a tuple with the FromTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTokenAmount

`func (o *QueryTransferStatusResponse) SetFromTokenAmount(v string)`

SetFromTokenAmount sets FromTokenAmount field to given value.

### HasFromTokenAmount

`func (o *QueryTransferStatusResponse) HasFromTokenAmount() bool`

HasFromTokenAmount returns a boolean if a field has been set.

### GetToToken

`func (o *QueryTransferStatusResponse) GetToToken() string`

GetToToken returns the ToToken field if non-nil, zero value otherwise.

### GetToTokenOk

`func (o *QueryTransferStatusResponse) GetToTokenOk() (*string, bool)`

GetToTokenOk returns a tuple with the ToToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToToken

`func (o *QueryTransferStatusResponse) SetToToken(v string)`

SetToToken sets ToToken field to given value.

### HasToToken

`func (o *QueryTransferStatusResponse) HasToToken() bool`

HasToToken returns a boolean if a field has been set.

### GetToTokenAmount

`func (o *QueryTransferStatusResponse) GetToTokenAmount() string`

GetToTokenAmount returns the ToTokenAmount field if non-nil, zero value otherwise.

### GetToTokenAmountOk

`func (o *QueryTransferStatusResponse) GetToTokenAmountOk() (*string, bool)`

GetToTokenAmountOk returns a tuple with the ToTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTokenAmount

`func (o *QueryTransferStatusResponse) SetToTokenAmount(v string)`

SetToTokenAmount sets ToTokenAmount field to given value.

### HasToTokenAmount

`func (o *QueryTransferStatusResponse) HasToTokenAmount() bool`

HasToTokenAmount returns a boolean if a field has been set.

### GetErrorCode

`func (o *QueryTransferStatusResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *QueryTransferStatusResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *QueryTransferStatusResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *QueryTransferStatusResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *QueryTransferStatusResponse) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *QueryTransferStatusResponse) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *QueryTransferStatusResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *QueryTransferStatusResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *QueryTransferStatusResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *QueryTransferStatusResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *QueryTransferStatusResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *QueryTransferStatusResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCreateTime

`func (o *QueryTransferStatusResponse) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryTransferStatusResponse) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryTransferStatusResponse) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryTransferStatusResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryTransferStatusResponse) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryTransferStatusResponse) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryTransferStatusResponse) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryTransferStatusResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


