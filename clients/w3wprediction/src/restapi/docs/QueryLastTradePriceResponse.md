# QueryLastTradePriceResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**MarketId** | Pointer to **int64** |  | [optional] 
**LastTradePrice** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryLastTradePriceResponse

`func NewQueryLastTradePriceResponse() *QueryLastTradePriceResponse`

NewQueryLastTradePriceResponse instantiates a new QueryLastTradePriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLastTradePriceResponseWithDefaults

`func NewQueryLastTradePriceResponseWithDefaults() *QueryLastTradePriceResponse`

NewQueryLastTradePriceResponseWithDefaults instantiates a new QueryLastTradePriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketId

`func (o *QueryLastTradePriceResponse) GetMarketId() int64`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *QueryLastTradePriceResponse) GetMarketIdOk() (*int64, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *QueryLastTradePriceResponse) SetMarketId(v int64)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *QueryLastTradePriceResponse) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetLastTradePrice

`func (o *QueryLastTradePriceResponse) GetLastTradePrice() string`

GetLastTradePrice returns the LastTradePrice field if non-nil, zero value otherwise.

### GetLastTradePriceOk

`func (o *QueryLastTradePriceResponse) GetLastTradePriceOk() (*string, bool)`

GetLastTradePriceOk returns a tuple with the LastTradePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradePrice

`func (o *QueryLastTradePriceResponse) SetLastTradePrice(v string)`

SetLastTradePrice sets LastTradePrice field to given value.

### HasLastTradePrice

`func (o *QueryLastTradePriceResponse) HasLastTradePrice() bool`

HasLastTradePrice returns a boolean if a field has been set.


[[Back to README]](../README.md)


