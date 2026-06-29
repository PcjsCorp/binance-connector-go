# QuerySettledPositionHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Positions** | Pointer to [**[]QuerySettledPositionHistoryResponsePositionsInner**](QuerySettledPositionHistoryResponsePositionsInner.md) |  | [optional] 

## Methods

### NewQuerySettledPositionHistoryResponse

`func NewQuerySettledPositionHistoryResponse() *QuerySettledPositionHistoryResponse`

NewQuerySettledPositionHistoryResponse instantiates a new QuerySettledPositionHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySettledPositionHistoryResponseWithDefaults

`func NewQuerySettledPositionHistoryResponseWithDefaults() *QuerySettledPositionHistoryResponse`

NewQuerySettledPositionHistoryResponseWithDefaults instantiates a new QuerySettledPositionHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *QuerySettledPositionHistoryResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *QuerySettledPositionHistoryResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *QuerySettledPositionHistoryResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *QuerySettledPositionHistoryResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetPositions

`func (o *QuerySettledPositionHistoryResponse) GetPositions() []QuerySettledPositionHistoryResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *QuerySettledPositionHistoryResponse) GetPositionsOk() (*[]QuerySettledPositionHistoryResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *QuerySettledPositionHistoryResponse) SetPositions(v []QuerySettledPositionHistoryResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *QuerySettledPositionHistoryResponse) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


