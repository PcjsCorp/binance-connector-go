# MarketSearchResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**MarketTopicId** | Pointer to **int64** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TopicType** | Pointer to **string** |  | [optional] 
**ChartType** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**ParticipantCount** | Pointer to **int32** |  | [optional] 
**Collateral** | Pointer to **string** |  | [optional] 
**TradeVolume** | Pointer to **string** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Markets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewMarketSearchResponseInner

`func NewMarketSearchResponseInner() *MarketSearchResponseInner`

NewMarketSearchResponseInner instantiates a new MarketSearchResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketSearchResponseInnerWithDefaults

`func NewMarketSearchResponseInnerWithDefaults() *MarketSearchResponseInner`

NewMarketSearchResponseInnerWithDefaults instantiates a new MarketSearchResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketTopicId

`func (o *MarketSearchResponseInner) GetMarketTopicId() int64`

GetMarketTopicId returns the MarketTopicId field if non-nil, zero value otherwise.

### GetMarketTopicIdOk

`func (o *MarketSearchResponseInner) GetMarketTopicIdOk() (*int64, bool)`

GetMarketTopicIdOk returns a tuple with the MarketTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicId

`func (o *MarketSearchResponseInner) SetMarketTopicId(v int64)`

SetMarketTopicId sets MarketTopicId field to given value.

### HasMarketTopicId

`func (o *MarketSearchResponseInner) HasMarketTopicId() bool`

HasMarketTopicId returns a boolean if a field has been set.

### GetVendor

`func (o *MarketSearchResponseInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *MarketSearchResponseInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *MarketSearchResponseInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *MarketSearchResponseInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetChainId

`func (o *MarketSearchResponseInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *MarketSearchResponseInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *MarketSearchResponseInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *MarketSearchResponseInner) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetSlug

`func (o *MarketSearchResponseInner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MarketSearchResponseInner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MarketSearchResponseInner) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *MarketSearchResponseInner) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *MarketSearchResponseInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketSearchResponseInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketSearchResponseInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketSearchResponseInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetQuestion

`func (o *MarketSearchResponseInner) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *MarketSearchResponseInner) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *MarketSearchResponseInner) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *MarketSearchResponseInner) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetDescription

`func (o *MarketSearchResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketSearchResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketSearchResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketSearchResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTopicType

`func (o *MarketSearchResponseInner) GetTopicType() string`

GetTopicType returns the TopicType field if non-nil, zero value otherwise.

### GetTopicTypeOk

`func (o *MarketSearchResponseInner) GetTopicTypeOk() (*string, bool)`

GetTopicTypeOk returns a tuple with the TopicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicType

`func (o *MarketSearchResponseInner) SetTopicType(v string)`

SetTopicType sets TopicType field to given value.

### HasTopicType

`func (o *MarketSearchResponseInner) HasTopicType() bool`

HasTopicType returns a boolean if a field has been set.

### GetChartType

`func (o *MarketSearchResponseInner) GetChartType() string`

GetChartType returns the ChartType field if non-nil, zero value otherwise.

### GetChartTypeOk

`func (o *MarketSearchResponseInner) GetChartTypeOk() (*string, bool)`

GetChartTypeOk returns a tuple with the ChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartType

`func (o *MarketSearchResponseInner) SetChartType(v string)`

SetChartType sets ChartType field to given value.

### HasChartType

`func (o *MarketSearchResponseInner) HasChartType() bool`

HasChartType returns a boolean if a field has been set.

### GetSymbol

`func (o *MarketSearchResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarketSearchResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarketSearchResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarketSearchResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetParticipantCount

`func (o *MarketSearchResponseInner) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *MarketSearchResponseInner) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *MarketSearchResponseInner) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *MarketSearchResponseInner) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetCollateral

`func (o *MarketSearchResponseInner) GetCollateral() string`

GetCollateral returns the Collateral field if non-nil, zero value otherwise.

### GetCollateralOk

`func (o *MarketSearchResponseInner) GetCollateralOk() (*string, bool)`

GetCollateralOk returns a tuple with the Collateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateral

`func (o *MarketSearchResponseInner) SetCollateral(v string)`

SetCollateral sets Collateral field to given value.

### HasCollateral

`func (o *MarketSearchResponseInner) HasCollateral() bool`

HasCollateral returns a boolean if a field has been set.

### GetTradeVolume

`func (o *MarketSearchResponseInner) GetTradeVolume() string`

GetTradeVolume returns the TradeVolume field if non-nil, zero value otherwise.

### GetTradeVolumeOk

`func (o *MarketSearchResponseInner) GetTradeVolumeOk() (*string, bool)`

GetTradeVolumeOk returns a tuple with the TradeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeVolume

`func (o *MarketSearchResponseInner) SetTradeVolume(v string)`

SetTradeVolume sets TradeVolume field to given value.

### HasTradeVolume

`func (o *MarketSearchResponseInner) HasTradeVolume() bool`

HasTradeVolume returns a boolean if a field has been set.

### GetLiquidity

`func (o *MarketSearchResponseInner) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *MarketSearchResponseInner) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *MarketSearchResponseInner) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *MarketSearchResponseInner) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStartDate

`func (o *MarketSearchResponseInner) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MarketSearchResponseInner) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MarketSearchResponseInner) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *MarketSearchResponseInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *MarketSearchResponseInner) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MarketSearchResponseInner) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MarketSearchResponseInner) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *MarketSearchResponseInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *MarketSearchResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketSearchResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketSearchResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketSearchResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMarkets

`func (o *MarketSearchResponseInner) GetMarkets() []map[string]interface{}`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *MarketSearchResponseInner) GetMarketsOk() (*[]map[string]interface{}, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *MarketSearchResponseInner) SetMarkets(v []map[string]interface{})`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *MarketSearchResponseInner) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.


[[Back to README]](../README.md)


