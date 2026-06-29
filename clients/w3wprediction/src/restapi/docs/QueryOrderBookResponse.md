# QueryOrderBookResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Outcome** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Bids** | Pointer to [**[]QueryOrderBookResponseBidsInner**](QueryOrderBookResponseBidsInner.md) |  | [optional] 
**Asks** | Pointer to [**[]QueryOrderBookResponseAsksInner**](QueryOrderBookResponseAsksInner.md) |  | [optional] 

## Methods

### NewQueryOrderBookResponse

`func NewQueryOrderBookResponse() *QueryOrderBookResponse`

NewQueryOrderBookResponse instantiates a new QueryOrderBookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOrderBookResponseWithDefaults

`func NewQueryOrderBookResponseWithDefaults() *QueryOrderBookResponse`

NewQueryOrderBookResponseWithDefaults instantiates a new QueryOrderBookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcome

`func (o *QueryOrderBookResponse) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *QueryOrderBookResponse) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *QueryOrderBookResponse) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *QueryOrderBookResponse) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetTokenId

`func (o *QueryOrderBookResponse) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *QueryOrderBookResponse) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *QueryOrderBookResponse) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *QueryOrderBookResponse) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTimestamp

`func (o *QueryOrderBookResponse) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *QueryOrderBookResponse) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *QueryOrderBookResponse) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *QueryOrderBookResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBids

`func (o *QueryOrderBookResponse) GetBids() []QueryOrderBookResponseBidsInner`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *QueryOrderBookResponse) GetBidsOk() (*[]QueryOrderBookResponseBidsInner, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *QueryOrderBookResponse) SetBids(v []QueryOrderBookResponseBidsInner)`

SetBids sets Bids field to given value.

### HasBids

`func (o *QueryOrderBookResponse) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *QueryOrderBookResponse) GetAsks() []QueryOrderBookResponseAsksInner`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *QueryOrderBookResponse) GetAsksOk() (*[]QueryOrderBookResponseAsksInner, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *QueryOrderBookResponse) SetAsks(v []QueryOrderBookResponseAsksInner)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *QueryOrderBookResponse) HasAsks() bool`

HasAsks returns a boolean if a field has been set.


[[Back to README]](../README.md)


