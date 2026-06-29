# GetMarketDetailResponse

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
**VariantData** | Pointer to [**GetMarketDetailResponseVariantData**](GetMarketDetailResponseVariantData.md) |  | [optional] 
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
**Timeline** | Pointer to [**[]GetMarketDetailResponseTimelineInner**](GetMarketDetailResponseTimelineInner.md) |  | [optional] 
**Markets** | Pointer to [**[]GetMarketDetailResponseMarketsInner**](GetMarketDetailResponseMarketsInner.md) |  | [optional] 

## Methods

### NewGetMarketDetailResponse

`func NewGetMarketDetailResponse() *GetMarketDetailResponse`

NewGetMarketDetailResponse instantiates a new GetMarketDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketDetailResponseWithDefaults

`func NewGetMarketDetailResponseWithDefaults() *GetMarketDetailResponse`

NewGetMarketDetailResponseWithDefaults instantiates a new GetMarketDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketTopicId

`func (o *GetMarketDetailResponse) GetMarketTopicId() int64`

GetMarketTopicId returns the MarketTopicId field if non-nil, zero value otherwise.

### GetMarketTopicIdOk

`func (o *GetMarketDetailResponse) GetMarketTopicIdOk() (*int64, bool)`

GetMarketTopicIdOk returns a tuple with the MarketTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicId

`func (o *GetMarketDetailResponse) SetMarketTopicId(v int64)`

SetMarketTopicId sets MarketTopicId field to given value.

### HasMarketTopicId

`func (o *GetMarketDetailResponse) HasMarketTopicId() bool`

HasMarketTopicId returns a boolean if a field has been set.

### GetVendor

`func (o *GetMarketDetailResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GetMarketDetailResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GetMarketDetailResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GetMarketDetailResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetChainId

`func (o *GetMarketDetailResponse) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GetMarketDetailResponse) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GetMarketDetailResponse) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *GetMarketDetailResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetSlug

`func (o *GetMarketDetailResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GetMarketDetailResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GetMarketDetailResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GetMarketDetailResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *GetMarketDetailResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetMarketDetailResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetMarketDetailResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetMarketDetailResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetQuestion

`func (o *GetMarketDetailResponse) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *GetMarketDetailResponse) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *GetMarketDetailResponse) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *GetMarketDetailResponse) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetDescription

`func (o *GetMarketDetailResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMarketDetailResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMarketDetailResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMarketDetailResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *GetMarketDetailResponse) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GetMarketDetailResponse) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GetMarketDetailResponse) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GetMarketDetailResponse) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetTopicType

`func (o *GetMarketDetailResponse) GetTopicType() string`

GetTopicType returns the TopicType field if non-nil, zero value otherwise.

### GetTopicTypeOk

`func (o *GetMarketDetailResponse) GetTopicTypeOk() (*string, bool)`

GetTopicTypeOk returns a tuple with the TopicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicType

`func (o *GetMarketDetailResponse) SetTopicType(v string)`

SetTopicType sets TopicType field to given value.

### HasTopicType

`func (o *GetMarketDetailResponse) HasTopicType() bool`

HasTopicType returns a boolean if a field has been set.

### GetChartType

`func (o *GetMarketDetailResponse) GetChartType() string`

GetChartType returns the ChartType field if non-nil, zero value otherwise.

### GetChartTypeOk

`func (o *GetMarketDetailResponse) GetChartTypeOk() (*string, bool)`

GetChartTypeOk returns a tuple with the ChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartType

`func (o *GetMarketDetailResponse) SetChartType(v string)`

SetChartType sets ChartType field to given value.

### HasChartType

`func (o *GetMarketDetailResponse) HasChartType() bool`

HasChartType returns a boolean if a field has been set.

### GetSymbol

`func (o *GetMarketDetailResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetMarketDetailResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetMarketDetailResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetMarketDetailResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetVariantData

`func (o *GetMarketDetailResponse) GetVariantData() GetMarketDetailResponseVariantData`

GetVariantData returns the VariantData field if non-nil, zero value otherwise.

### GetVariantDataOk

`func (o *GetMarketDetailResponse) GetVariantDataOk() (*GetMarketDetailResponseVariantData, bool)`

GetVariantDataOk returns a tuple with the VariantData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantData

`func (o *GetMarketDetailResponse) SetVariantData(v GetMarketDetailResponseVariantData)`

SetVariantData sets VariantData field to given value.

### HasVariantData

`func (o *GetMarketDetailResponse) HasVariantData() bool`

HasVariantData returns a boolean if a field has been set.

### GetParticipantCount

`func (o *GetMarketDetailResponse) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *GetMarketDetailResponse) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *GetMarketDetailResponse) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *GetMarketDetailResponse) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetCollateral

`func (o *GetMarketDetailResponse) GetCollateral() string`

GetCollateral returns the Collateral field if non-nil, zero value otherwise.

### GetCollateralOk

`func (o *GetMarketDetailResponse) GetCollateralOk() (*string, bool)`

GetCollateralOk returns a tuple with the Collateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateral

`func (o *GetMarketDetailResponse) SetCollateral(v string)`

SetCollateral sets Collateral field to given value.

### HasCollateral

`func (o *GetMarketDetailResponse) HasCollateral() bool`

HasCollateral returns a boolean if a field has been set.

### GetFeeRateBps

`func (o *GetMarketDetailResponse) GetFeeRateBps() int32`

GetFeeRateBps returns the FeeRateBps field if non-nil, zero value otherwise.

### GetFeeRateBpsOk

`func (o *GetMarketDetailResponse) GetFeeRateBpsOk() (*int32, bool)`

GetFeeRateBpsOk returns a tuple with the FeeRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRateBps

`func (o *GetMarketDetailResponse) SetFeeRateBps(v int32)`

SetFeeRateBps sets FeeRateBps field to given value.

### HasFeeRateBps

`func (o *GetMarketDetailResponse) HasFeeRateBps() bool`

HasFeeRateBps returns a boolean if a field has been set.

### GetSlippageBps

`func (o *GetMarketDetailResponse) GetSlippageBps() int32`

GetSlippageBps returns the SlippageBps field if non-nil, zero value otherwise.

### GetSlippageBpsOk

`func (o *GetMarketDetailResponse) GetSlippageBpsOk() (*int32, bool)`

GetSlippageBpsOk returns a tuple with the SlippageBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageBps

`func (o *GetMarketDetailResponse) SetSlippageBps(v int32)`

SetSlippageBps sets SlippageBps field to given value.

### HasSlippageBps

`func (o *GetMarketDetailResponse) HasSlippageBps() bool`

HasSlippageBps returns a boolean if a field has been set.

### GetIsYieldBearing

`func (o *GetMarketDetailResponse) GetIsYieldBearing() bool`

GetIsYieldBearing returns the IsYieldBearing field if non-nil, zero value otherwise.

### GetIsYieldBearingOk

`func (o *GetMarketDetailResponse) GetIsYieldBearingOk() (*bool, bool)`

GetIsYieldBearingOk returns a tuple with the IsYieldBearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYieldBearing

`func (o *GetMarketDetailResponse) SetIsYieldBearing(v bool)`

SetIsYieldBearing sets IsYieldBearing field to given value.

### HasIsYieldBearing

`func (o *GetMarketDetailResponse) HasIsYieldBearing() bool`

HasIsYieldBearing returns a boolean if a field has been set.

### GetTradeVolume

`func (o *GetMarketDetailResponse) GetTradeVolume() string`

GetTradeVolume returns the TradeVolume field if non-nil, zero value otherwise.

### GetTradeVolumeOk

`func (o *GetMarketDetailResponse) GetTradeVolumeOk() (*string, bool)`

GetTradeVolumeOk returns a tuple with the TradeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeVolume

`func (o *GetMarketDetailResponse) SetTradeVolume(v string)`

SetTradeVolume sets TradeVolume field to given value.

### HasTradeVolume

`func (o *GetMarketDetailResponse) HasTradeVolume() bool`

HasTradeVolume returns a boolean if a field has been set.

### GetLiquidity

`func (o *GetMarketDetailResponse) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *GetMarketDetailResponse) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *GetMarketDetailResponse) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *GetMarketDetailResponse) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetPublishedAt

`func (o *GetMarketDetailResponse) GetPublishedAt() int64`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *GetMarketDetailResponse) GetPublishedAtOk() (*int64, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *GetMarketDetailResponse) SetPublishedAt(v int64)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *GetMarketDetailResponse) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetStartDate

`func (o *GetMarketDetailResponse) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetMarketDetailResponse) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetMarketDetailResponse) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetMarketDetailResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetMarketDetailResponse) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetMarketDetailResponse) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetMarketDetailResponse) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetMarketDetailResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *GetMarketDetailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMarketDetailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMarketDetailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMarketDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeline

`func (o *GetMarketDetailResponse) GetTimeline() []GetMarketDetailResponseTimelineInner`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *GetMarketDetailResponse) GetTimelineOk() (*[]GetMarketDetailResponseTimelineInner, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *GetMarketDetailResponse) SetTimeline(v []GetMarketDetailResponseTimelineInner)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *GetMarketDetailResponse) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetMarkets

`func (o *GetMarketDetailResponse) GetMarkets() []GetMarketDetailResponseMarketsInner`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *GetMarketDetailResponse) GetMarketsOk() (*[]GetMarketDetailResponseMarketsInner, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *GetMarketDetailResponse) SetMarkets(v []GetMarketDetailResponseMarketsInner)`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *GetMarketDetailResponse) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.


[[Back to README]](../README.md)


