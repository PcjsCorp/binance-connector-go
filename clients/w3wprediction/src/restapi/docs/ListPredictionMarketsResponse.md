# ListPredictionMarketsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**MarketTopics** | Pointer to [**[]ListPredictionMarketsResponseMarketTopicsInner**](ListPredictionMarketsResponseMarketTopicsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewListPredictionMarketsResponse

`func NewListPredictionMarketsResponse() *ListPredictionMarketsResponse`

NewListPredictionMarketsResponse instantiates a new ListPredictionMarketsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPredictionMarketsResponseWithDefaults

`func NewListPredictionMarketsResponseWithDefaults() *ListPredictionMarketsResponse`

NewListPredictionMarketsResponseWithDefaults instantiates a new ListPredictionMarketsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketTopics

`func (o *ListPredictionMarketsResponse) GetMarketTopics() []ListPredictionMarketsResponseMarketTopicsInner`

GetMarketTopics returns the MarketTopics field if non-nil, zero value otherwise.

### GetMarketTopicsOk

`func (o *ListPredictionMarketsResponse) GetMarketTopicsOk() (*[]ListPredictionMarketsResponseMarketTopicsInner, bool)`

GetMarketTopicsOk returns a tuple with the MarketTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopics

`func (o *ListPredictionMarketsResponse) SetMarketTopics(v []ListPredictionMarketsResponseMarketTopicsInner)`

SetMarketTopics sets MarketTopics field to given value.

### HasMarketTopics

`func (o *ListPredictionMarketsResponse) HasMarketTopics() bool`

HasMarketTopics returns a boolean if a field has been set.

### GetTotal

`func (o *ListPredictionMarketsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListPredictionMarketsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListPredictionMarketsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListPredictionMarketsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOffset

`func (o *ListPredictionMarketsResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListPredictionMarketsResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListPredictionMarketsResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListPredictionMarketsResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *ListPredictionMarketsResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListPredictionMarketsResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListPredictionMarketsResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListPredictionMarketsResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetHasMore

`func (o *ListPredictionMarketsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListPredictionMarketsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListPredictionMarketsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListPredictionMarketsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to README]](../README.md)


