# QueryOrderHistoryResponseOrdersInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** |  | [optional] 
**VendorOrderId** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**MarketTopicId** | Pointer to **int64** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**MarketTopicTitle** | Pointer to **string** |  | [optional] 
**MarketId** | Pointer to **int64** |  | [optional] 
**MarketTitle** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **string** |  | [optional] 
**OutcomeIndex** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**ModifyTime** | Pointer to **int64** |  | [optional] 
**TerminalTime** | Pointer to **int64** |  | [optional] 
**MakerUsdtAmount** | Pointer to **string** |  | [optional] 
**MakerShareQty** | Pointer to **string** |  | [optional] 
**FilledUsdtAmount** | Pointer to **string** |  | [optional] 
**FilledShareQty** | Pointer to **string** |  | [optional] 
**FillPercentage** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**MarketProviderFee** | Pointer to **string** |  | [optional] 
**NetworkFee** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryOrderHistoryResponseOrdersInner

`func NewQueryOrderHistoryResponseOrdersInner() *QueryOrderHistoryResponseOrdersInner`

NewQueryOrderHistoryResponseOrdersInner instantiates a new QueryOrderHistoryResponseOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOrderHistoryResponseOrdersInnerWithDefaults

`func NewQueryOrderHistoryResponseOrdersInnerWithDefaults() *QueryOrderHistoryResponseOrdersInner`

NewQueryOrderHistoryResponseOrdersInnerWithDefaults instantiates a new QueryOrderHistoryResponseOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *QueryOrderHistoryResponseOrdersInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryOrderHistoryResponseOrdersInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryOrderHistoryResponseOrdersInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetVendorOrderId

`func (o *QueryOrderHistoryResponseOrdersInner) GetVendorOrderId() string`

GetVendorOrderId returns the VendorOrderId field if non-nil, zero value otherwise.

### GetVendorOrderIdOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetVendorOrderIdOk() (*string, bool)`

GetVendorOrderIdOk returns a tuple with the VendorOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorOrderId

`func (o *QueryOrderHistoryResponseOrdersInner) SetVendorOrderId(v string)`

SetVendorOrderId sets VendorOrderId field to given value.

### HasVendorOrderId

`func (o *QueryOrderHistoryResponseOrdersInner) HasVendorOrderId() bool`

HasVendorOrderId returns a boolean if a field has been set.

### GetVendor

`func (o *QueryOrderHistoryResponseOrdersInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QueryOrderHistoryResponseOrdersInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *QueryOrderHistoryResponseOrdersInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetMarketTopicId

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketTopicId() int64`

GetMarketTopicId returns the MarketTopicId field if non-nil, zero value otherwise.

### GetMarketTopicIdOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketTopicIdOk() (*int64, bool)`

GetMarketTopicIdOk returns a tuple with the MarketTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicId

`func (o *QueryOrderHistoryResponseOrdersInner) SetMarketTopicId(v int64)`

SetMarketTopicId sets MarketTopicId field to given value.

### HasMarketTopicId

`func (o *QueryOrderHistoryResponseOrdersInner) HasMarketTopicId() bool`

HasMarketTopicId returns a boolean if a field has been set.

### GetSlug

`func (o *QueryOrderHistoryResponseOrdersInner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *QueryOrderHistoryResponseOrdersInner) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *QueryOrderHistoryResponseOrdersInner) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMarketTopicTitle

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketTopicTitle() string`

GetMarketTopicTitle returns the MarketTopicTitle field if non-nil, zero value otherwise.

### GetMarketTopicTitleOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketTopicTitleOk() (*string, bool)`

GetMarketTopicTitleOk returns a tuple with the MarketTopicTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicTitle

`func (o *QueryOrderHistoryResponseOrdersInner) SetMarketTopicTitle(v string)`

SetMarketTopicTitle sets MarketTopicTitle field to given value.

### HasMarketTopicTitle

`func (o *QueryOrderHistoryResponseOrdersInner) HasMarketTopicTitle() bool`

HasMarketTopicTitle returns a boolean if a field has been set.

### GetMarketId

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketId() int64`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketIdOk() (*int64, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *QueryOrderHistoryResponseOrdersInner) SetMarketId(v int64)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *QueryOrderHistoryResponseOrdersInner) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetMarketTitle

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketTitle() string`

GetMarketTitle returns the MarketTitle field if non-nil, zero value otherwise.

### GetMarketTitleOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketTitleOk() (*string, bool)`

GetMarketTitleOk returns a tuple with the MarketTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTitle

`func (o *QueryOrderHistoryResponseOrdersInner) SetMarketTitle(v string)`

SetMarketTitle sets MarketTitle field to given value.

### HasMarketTitle

`func (o *QueryOrderHistoryResponseOrdersInner) HasMarketTitle() bool`

HasMarketTitle returns a boolean if a field has been set.

### GetOutcome

`func (o *QueryOrderHistoryResponseOrdersInner) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *QueryOrderHistoryResponseOrdersInner) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *QueryOrderHistoryResponseOrdersInner) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutcomeIndex

`func (o *QueryOrderHistoryResponseOrdersInner) GetOutcomeIndex() int32`

GetOutcomeIndex returns the OutcomeIndex field if non-nil, zero value otherwise.

### GetOutcomeIndexOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetOutcomeIndexOk() (*int32, bool)`

GetOutcomeIndexOk returns a tuple with the OutcomeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeIndex

`func (o *QueryOrderHistoryResponseOrdersInner) SetOutcomeIndex(v int32)`

SetOutcomeIndex sets OutcomeIndex field to given value.

### HasOutcomeIndex

`func (o *QueryOrderHistoryResponseOrdersInner) HasOutcomeIndex() bool`

HasOutcomeIndex returns a boolean if a field has been set.

### GetStatus

`func (o *QueryOrderHistoryResponseOrdersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryOrderHistoryResponseOrdersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryOrderHistoryResponseOrdersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSide

`func (o *QueryOrderHistoryResponseOrdersInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryOrderHistoryResponseOrdersInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryOrderHistoryResponseOrdersInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderType

`func (o *QueryOrderHistoryResponseOrdersInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *QueryOrderHistoryResponseOrdersInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *QueryOrderHistoryResponseOrdersInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetCreateTime

`func (o *QueryOrderHistoryResponseOrdersInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryOrderHistoryResponseOrdersInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryOrderHistoryResponseOrdersInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetModifyTime

`func (o *QueryOrderHistoryResponseOrdersInner) GetModifyTime() int64`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetModifyTimeOk() (*int64, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *QueryOrderHistoryResponseOrdersInner) SetModifyTime(v int64)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *QueryOrderHistoryResponseOrdersInner) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetTerminalTime

`func (o *QueryOrderHistoryResponseOrdersInner) GetTerminalTime() int64`

GetTerminalTime returns the TerminalTime field if non-nil, zero value otherwise.

### GetTerminalTimeOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetTerminalTimeOk() (*int64, bool)`

GetTerminalTimeOk returns a tuple with the TerminalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTime

`func (o *QueryOrderHistoryResponseOrdersInner) SetTerminalTime(v int64)`

SetTerminalTime sets TerminalTime field to given value.

### HasTerminalTime

`func (o *QueryOrderHistoryResponseOrdersInner) HasTerminalTime() bool`

HasTerminalTime returns a boolean if a field has been set.

### GetMakerUsdtAmount

`func (o *QueryOrderHistoryResponseOrdersInner) GetMakerUsdtAmount() string`

GetMakerUsdtAmount returns the MakerUsdtAmount field if non-nil, zero value otherwise.

### GetMakerUsdtAmountOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetMakerUsdtAmountOk() (*string, bool)`

GetMakerUsdtAmountOk returns a tuple with the MakerUsdtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerUsdtAmount

`func (o *QueryOrderHistoryResponseOrdersInner) SetMakerUsdtAmount(v string)`

SetMakerUsdtAmount sets MakerUsdtAmount field to given value.

### HasMakerUsdtAmount

`func (o *QueryOrderHistoryResponseOrdersInner) HasMakerUsdtAmount() bool`

HasMakerUsdtAmount returns a boolean if a field has been set.

### GetMakerShareQty

`func (o *QueryOrderHistoryResponseOrdersInner) GetMakerShareQty() string`

GetMakerShareQty returns the MakerShareQty field if non-nil, zero value otherwise.

### GetMakerShareQtyOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetMakerShareQtyOk() (*string, bool)`

GetMakerShareQtyOk returns a tuple with the MakerShareQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerShareQty

`func (o *QueryOrderHistoryResponseOrdersInner) SetMakerShareQty(v string)`

SetMakerShareQty sets MakerShareQty field to given value.

### HasMakerShareQty

`func (o *QueryOrderHistoryResponseOrdersInner) HasMakerShareQty() bool`

HasMakerShareQty returns a boolean if a field has been set.

### GetFilledUsdtAmount

`func (o *QueryOrderHistoryResponseOrdersInner) GetFilledUsdtAmount() string`

GetFilledUsdtAmount returns the FilledUsdtAmount field if non-nil, zero value otherwise.

### GetFilledUsdtAmountOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetFilledUsdtAmountOk() (*string, bool)`

GetFilledUsdtAmountOk returns a tuple with the FilledUsdtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledUsdtAmount

`func (o *QueryOrderHistoryResponseOrdersInner) SetFilledUsdtAmount(v string)`

SetFilledUsdtAmount sets FilledUsdtAmount field to given value.

### HasFilledUsdtAmount

`func (o *QueryOrderHistoryResponseOrdersInner) HasFilledUsdtAmount() bool`

HasFilledUsdtAmount returns a boolean if a field has been set.

### GetFilledShareQty

`func (o *QueryOrderHistoryResponseOrdersInner) GetFilledShareQty() string`

GetFilledShareQty returns the FilledShareQty field if non-nil, zero value otherwise.

### GetFilledShareQtyOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetFilledShareQtyOk() (*string, bool)`

GetFilledShareQtyOk returns a tuple with the FilledShareQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledShareQty

`func (o *QueryOrderHistoryResponseOrdersInner) SetFilledShareQty(v string)`

SetFilledShareQty sets FilledShareQty field to given value.

### HasFilledShareQty

`func (o *QueryOrderHistoryResponseOrdersInner) HasFilledShareQty() bool`

HasFilledShareQty returns a boolean if a field has been set.

### GetFillPercentage

`func (o *QueryOrderHistoryResponseOrdersInner) GetFillPercentage() string`

GetFillPercentage returns the FillPercentage field if non-nil, zero value otherwise.

### GetFillPercentageOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetFillPercentageOk() (*string, bool)`

GetFillPercentageOk returns a tuple with the FillPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPercentage

`func (o *QueryOrderHistoryResponseOrdersInner) SetFillPercentage(v string)`

SetFillPercentage sets FillPercentage field to given value.

### HasFillPercentage

`func (o *QueryOrderHistoryResponseOrdersInner) HasFillPercentage() bool`

HasFillPercentage returns a boolean if a field has been set.

### GetPrice

`func (o *QueryOrderHistoryResponseOrdersInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryOrderHistoryResponseOrdersInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryOrderHistoryResponseOrdersInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMarketProviderFee

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketProviderFee() string`

GetMarketProviderFee returns the MarketProviderFee field if non-nil, zero value otherwise.

### GetMarketProviderFeeOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetMarketProviderFeeOk() (*string, bool)`

GetMarketProviderFeeOk returns a tuple with the MarketProviderFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketProviderFee

`func (o *QueryOrderHistoryResponseOrdersInner) SetMarketProviderFee(v string)`

SetMarketProviderFee sets MarketProviderFee field to given value.

### HasMarketProviderFee

`func (o *QueryOrderHistoryResponseOrdersInner) HasMarketProviderFee() bool`

HasMarketProviderFee returns a boolean if a field has been set.

### GetNetworkFee

`func (o *QueryOrderHistoryResponseOrdersInner) GetNetworkFee() string`

GetNetworkFee returns the NetworkFee field if non-nil, zero value otherwise.

### GetNetworkFeeOk

`func (o *QueryOrderHistoryResponseOrdersInner) GetNetworkFeeOk() (*string, bool)`

GetNetworkFeeOk returns a tuple with the NetworkFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFee

`func (o *QueryOrderHistoryResponseOrdersInner) SetNetworkFee(v string)`

SetNetworkFee sets NetworkFee field to given value.

### HasNetworkFee

`func (o *QueryOrderHistoryResponseOrdersInner) HasNetworkFee() bool`

HasNetworkFee returns a boolean if a field has been set.


[[Back to README]](../README.md)


