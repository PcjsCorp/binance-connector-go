# ListPredictionMarketsResponseMarketTopicsInner

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
**ImageUrl** | Pointer to **string** |  | [optional] 
**TopicType** | Pointer to **string** |  | [optional] 
**ChartType** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**ParticipantCount** | Pointer to **int32** |  | [optional] 
**Collateral** | Pointer to **string** |  | [optional] 
**FeeRateBps** | Pointer to **int32** |  | [optional] 
**SlippageBps** | Pointer to **int32** |  | [optional] 
**IsYieldBearing** | Pointer to **bool** |  | [optional] 
**TradeVolume** | Pointer to **string** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**PublishedAt** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Markets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListPredictionMarketsResponseMarketTopicsInner

`func NewListPredictionMarketsResponseMarketTopicsInner() *ListPredictionMarketsResponseMarketTopicsInner`

NewListPredictionMarketsResponseMarketTopicsInner instantiates a new ListPredictionMarketsResponseMarketTopicsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPredictionMarketsResponseMarketTopicsInnerWithDefaults

`func NewListPredictionMarketsResponseMarketTopicsInnerWithDefaults() *ListPredictionMarketsResponseMarketTopicsInner`

NewListPredictionMarketsResponseMarketTopicsInnerWithDefaults instantiates a new ListPredictionMarketsResponseMarketTopicsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketTopicId

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetMarketTopicId() int64`

GetMarketTopicId returns the MarketTopicId field if non-nil, zero value otherwise.

### GetMarketTopicIdOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetMarketTopicIdOk() (*int64, bool)`

GetMarketTopicIdOk returns a tuple with the MarketTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicId

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetMarketTopicId(v int64)`

SetMarketTopicId sets MarketTopicId field to given value.

### HasMarketTopicId

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasMarketTopicId() bool`

HasMarketTopicId returns a boolean if a field has been set.

### GetVendor

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetChainId

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetSlug

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetQuestion

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetDescription

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetTopicType

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetTopicType() string`

GetTopicType returns the TopicType field if non-nil, zero value otherwise.

### GetTopicTypeOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetTopicTypeOk() (*string, bool)`

GetTopicTypeOk returns a tuple with the TopicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicType

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetTopicType(v string)`

SetTopicType sets TopicType field to given value.

### HasTopicType

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasTopicType() bool`

HasTopicType returns a boolean if a field has been set.

### GetChartType

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetChartType() string`

GetChartType returns the ChartType field if non-nil, zero value otherwise.

### GetChartTypeOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetChartTypeOk() (*string, bool)`

GetChartTypeOk returns a tuple with the ChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartType

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetChartType(v string)`

SetChartType sets ChartType field to given value.

### HasChartType

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasChartType() bool`

HasChartType returns a boolean if a field has been set.

### GetSymbol

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetParticipantCount

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetCollateral

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetCollateral() string`

GetCollateral returns the Collateral field if non-nil, zero value otherwise.

### GetCollateralOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetCollateralOk() (*string, bool)`

GetCollateralOk returns a tuple with the Collateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateral

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetCollateral(v string)`

SetCollateral sets Collateral field to given value.

### HasCollateral

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasCollateral() bool`

HasCollateral returns a boolean if a field has been set.

### GetFeeRateBps

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetFeeRateBps() int32`

GetFeeRateBps returns the FeeRateBps field if non-nil, zero value otherwise.

### GetFeeRateBpsOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetFeeRateBpsOk() (*int32, bool)`

GetFeeRateBpsOk returns a tuple with the FeeRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRateBps

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetFeeRateBps(v int32)`

SetFeeRateBps sets FeeRateBps field to given value.

### HasFeeRateBps

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasFeeRateBps() bool`

HasFeeRateBps returns a boolean if a field has been set.

### GetSlippageBps

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetSlippageBps() int32`

GetSlippageBps returns the SlippageBps field if non-nil, zero value otherwise.

### GetSlippageBpsOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetSlippageBpsOk() (*int32, bool)`

GetSlippageBpsOk returns a tuple with the SlippageBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageBps

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetSlippageBps(v int32)`

SetSlippageBps sets SlippageBps field to given value.

### HasSlippageBps

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasSlippageBps() bool`

HasSlippageBps returns a boolean if a field has been set.

### GetIsYieldBearing

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetIsYieldBearing() bool`

GetIsYieldBearing returns the IsYieldBearing field if non-nil, zero value otherwise.

### GetIsYieldBearingOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetIsYieldBearingOk() (*bool, bool)`

GetIsYieldBearingOk returns a tuple with the IsYieldBearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYieldBearing

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetIsYieldBearing(v bool)`

SetIsYieldBearing sets IsYieldBearing field to given value.

### HasIsYieldBearing

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasIsYieldBearing() bool`

HasIsYieldBearing returns a boolean if a field has been set.

### GetTradeVolume

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetTradeVolume() string`

GetTradeVolume returns the TradeVolume field if non-nil, zero value otherwise.

### GetTradeVolumeOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetTradeVolumeOk() (*string, bool)`

GetTradeVolumeOk returns a tuple with the TradeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeVolume

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetTradeVolume(v string)`

SetTradeVolume sets TradeVolume field to given value.

### HasTradeVolume

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasTradeVolume() bool`

HasTradeVolume returns a boolean if a field has been set.

### GetLiquidity

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetPublishedAt

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetPublishedAt() int64`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetPublishedAtOk() (*int64, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetPublishedAt(v int64)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetStartDate

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMarkets

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetMarkets() []map[string]interface{}`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *ListPredictionMarketsResponseMarketTopicsInner) GetMarketsOk() (*[]map[string]interface{}, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *ListPredictionMarketsResponseMarketTopicsInner) SetMarkets(v []map[string]interface{})`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *ListPredictionMarketsResponseMarketTopicsInner) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.


[[Back to README]](../README.md)


