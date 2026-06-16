# AssetIndexResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Index** | Pointer to **string** |  | [optional] 
**BidBuffer** | Pointer to **string** |  | [optional] 
**AskBuffer** | Pointer to **string** |  | [optional] 
**BidRate** | Pointer to **string** |  | [optional] 
**AskRate** | Pointer to **string** |  | [optional] 
**AutoExchangeBidBuffer** | Pointer to **string** |  | [optional] 
**AutoExchangeAskBuffer** | Pointer to **string** |  | [optional] 
**AutoExchangeBidRate** | Pointer to **string** |  | [optional] 
**AutoExchangeAskRate** | Pointer to **string** |  | [optional] 

## Methods

### NewAssetIndexResponse

`func NewAssetIndexResponse() *AssetIndexResponse`

NewAssetIndexResponse instantiates a new AssetIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetIndexResponseWithDefaults

`func NewAssetIndexResponseWithDefaults() *AssetIndexResponse`

NewAssetIndexResponseWithDefaults instantiates a new AssetIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *AssetIndexResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AssetIndexResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AssetIndexResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AssetIndexResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *AssetIndexResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AssetIndexResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AssetIndexResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *AssetIndexResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetIndex

`func (o *AssetIndexResponse) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AssetIndexResponse) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AssetIndexResponse) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AssetIndexResponse) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetBidBuffer

`func (o *AssetIndexResponse) GetBidBuffer() string`

GetBidBuffer returns the BidBuffer field if non-nil, zero value otherwise.

### GetBidBufferOk

`func (o *AssetIndexResponse) GetBidBufferOk() (*string, bool)`

GetBidBufferOk returns a tuple with the BidBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidBuffer

`func (o *AssetIndexResponse) SetBidBuffer(v string)`

SetBidBuffer sets BidBuffer field to given value.

### HasBidBuffer

`func (o *AssetIndexResponse) HasBidBuffer() bool`

HasBidBuffer returns a boolean if a field has been set.

### GetAskBuffer

`func (o *AssetIndexResponse) GetAskBuffer() string`

GetAskBuffer returns the AskBuffer field if non-nil, zero value otherwise.

### GetAskBufferOk

`func (o *AssetIndexResponse) GetAskBufferOk() (*string, bool)`

GetAskBufferOk returns a tuple with the AskBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskBuffer

`func (o *AssetIndexResponse) SetAskBuffer(v string)`

SetAskBuffer sets AskBuffer field to given value.

### HasAskBuffer

`func (o *AssetIndexResponse) HasAskBuffer() bool`

HasAskBuffer returns a boolean if a field has been set.

### GetBidRate

`func (o *AssetIndexResponse) GetBidRate() string`

GetBidRate returns the BidRate field if non-nil, zero value otherwise.

### GetBidRateOk

`func (o *AssetIndexResponse) GetBidRateOk() (*string, bool)`

GetBidRateOk returns a tuple with the BidRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidRate

`func (o *AssetIndexResponse) SetBidRate(v string)`

SetBidRate sets BidRate field to given value.

### HasBidRate

`func (o *AssetIndexResponse) HasBidRate() bool`

HasBidRate returns a boolean if a field has been set.

### GetAskRate

`func (o *AssetIndexResponse) GetAskRate() string`

GetAskRate returns the AskRate field if non-nil, zero value otherwise.

### GetAskRateOk

`func (o *AssetIndexResponse) GetAskRateOk() (*string, bool)`

GetAskRateOk returns a tuple with the AskRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskRate

`func (o *AssetIndexResponse) SetAskRate(v string)`

SetAskRate sets AskRate field to given value.

### HasAskRate

`func (o *AssetIndexResponse) HasAskRate() bool`

HasAskRate returns a boolean if a field has been set.

### GetAutoExchangeBidBuffer

`func (o *AssetIndexResponse) GetAutoExchangeBidBuffer() string`

GetAutoExchangeBidBuffer returns the AutoExchangeBidBuffer field if non-nil, zero value otherwise.

### GetAutoExchangeBidBufferOk

`func (o *AssetIndexResponse) GetAutoExchangeBidBufferOk() (*string, bool)`

GetAutoExchangeBidBufferOk returns a tuple with the AutoExchangeBidBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExchangeBidBuffer

`func (o *AssetIndexResponse) SetAutoExchangeBidBuffer(v string)`

SetAutoExchangeBidBuffer sets AutoExchangeBidBuffer field to given value.

### HasAutoExchangeBidBuffer

`func (o *AssetIndexResponse) HasAutoExchangeBidBuffer() bool`

HasAutoExchangeBidBuffer returns a boolean if a field has been set.

### GetAutoExchangeAskBuffer

`func (o *AssetIndexResponse) GetAutoExchangeAskBuffer() string`

GetAutoExchangeAskBuffer returns the AutoExchangeAskBuffer field if non-nil, zero value otherwise.

### GetAutoExchangeAskBufferOk

`func (o *AssetIndexResponse) GetAutoExchangeAskBufferOk() (*string, bool)`

GetAutoExchangeAskBufferOk returns a tuple with the AutoExchangeAskBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExchangeAskBuffer

`func (o *AssetIndexResponse) SetAutoExchangeAskBuffer(v string)`

SetAutoExchangeAskBuffer sets AutoExchangeAskBuffer field to given value.

### HasAutoExchangeAskBuffer

`func (o *AssetIndexResponse) HasAutoExchangeAskBuffer() bool`

HasAutoExchangeAskBuffer returns a boolean if a field has been set.

### GetAutoExchangeBidRate

`func (o *AssetIndexResponse) GetAutoExchangeBidRate() string`

GetAutoExchangeBidRate returns the AutoExchangeBidRate field if non-nil, zero value otherwise.

### GetAutoExchangeBidRateOk

`func (o *AssetIndexResponse) GetAutoExchangeBidRateOk() (*string, bool)`

GetAutoExchangeBidRateOk returns a tuple with the AutoExchangeBidRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExchangeBidRate

`func (o *AssetIndexResponse) SetAutoExchangeBidRate(v string)`

SetAutoExchangeBidRate sets AutoExchangeBidRate field to given value.

### HasAutoExchangeBidRate

`func (o *AssetIndexResponse) HasAutoExchangeBidRate() bool`

HasAutoExchangeBidRate returns a boolean if a field has been set.

### GetAutoExchangeAskRate

`func (o *AssetIndexResponse) GetAutoExchangeAskRate() string`

GetAutoExchangeAskRate returns the AutoExchangeAskRate field if non-nil, zero value otherwise.

### GetAutoExchangeAskRateOk

`func (o *AssetIndexResponse) GetAutoExchangeAskRateOk() (*string, bool)`

GetAutoExchangeAskRateOk returns a tuple with the AutoExchangeAskRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExchangeAskRate

`func (o *AssetIndexResponse) SetAutoExchangeAskRate(v string)`

SetAutoExchangeAskRate sets AutoExchangeAskRate field to given value.

### HasAutoExchangeAskRate

`func (o *AssetIndexResponse) HasAutoExchangeAskRate() bool`

HasAutoExchangeAskRate returns a boolean if a field has been set.


[[Back to README]](../README.md)


