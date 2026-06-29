# QueryPositionsResponsePositionsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | Pointer to **int64** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**CollateralSymbol** | Pointer to **string** |  | [optional] 
**TopicType** | Pointer to **string** |  | [optional] 
**MarketTopicId** | Pointer to **int64** |  | [optional] 
**MarketId** | Pointer to **int64** |  | [optional] 
**MarketTopicTitle** | Pointer to **string** |  | [optional] 
**MarketTitle** | Pointer to **string** |  | [optional] 
**OutcomeName** | Pointer to **string** |  | [optional] 
**OutcomeIndex** | Pointer to **int32** |  | [optional] 
**Shares** | Pointer to **string** |  | [optional] 
**AvgPrice** | Pointer to **string** |  | [optional] 
**TotalCost** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**CurrentPrice** | Pointer to **string** |  | [optional] 
**ToWin** | Pointer to **string** |  | [optional] 
**PositionStatus** | Pointer to **string** |  | [optional] 
**CanClaim** | Pointer to **bool** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**UnrealizedPnl** | Pointer to **string** |  | [optional] 
**UnrealizedPnlPercent** | Pointer to **string** |  | [optional] 
**RealizedPnl** | Pointer to **string** |  | [optional] 
**Pnl** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **int64** |  | [optional] 
**UpdatedTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryPositionsResponsePositionsInner

`func NewQueryPositionsResponsePositionsInner() *QueryPositionsResponsePositionsInner`

NewQueryPositionsResponsePositionsInner instantiates a new QueryPositionsResponsePositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPositionsResponsePositionsInnerWithDefaults

`func NewQueryPositionsResponsePositionsInnerWithDefaults() *QueryPositionsResponsePositionsInner`

NewQueryPositionsResponsePositionsInnerWithDefaults instantiates a new QueryPositionsResponsePositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *QueryPositionsResponsePositionsInner) GetPositionId() int64`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *QueryPositionsResponsePositionsInner) GetPositionIdOk() (*int64, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *QueryPositionsResponsePositionsInner) SetPositionId(v int64)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *QueryPositionsResponsePositionsInner) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetVendor

`func (o *QueryPositionsResponsePositionsInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QueryPositionsResponsePositionsInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QueryPositionsResponsePositionsInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *QueryPositionsResponsePositionsInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetChainId

`func (o *QueryPositionsResponsePositionsInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *QueryPositionsResponsePositionsInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *QueryPositionsResponsePositionsInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *QueryPositionsResponsePositionsInner) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetTokenId

`func (o *QueryPositionsResponsePositionsInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *QueryPositionsResponsePositionsInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *QueryPositionsResponsePositionsInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *QueryPositionsResponsePositionsInner) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetCollateralSymbol

`func (o *QueryPositionsResponsePositionsInner) GetCollateralSymbol() string`

GetCollateralSymbol returns the CollateralSymbol field if non-nil, zero value otherwise.

### GetCollateralSymbolOk

`func (o *QueryPositionsResponsePositionsInner) GetCollateralSymbolOk() (*string, bool)`

GetCollateralSymbolOk returns a tuple with the CollateralSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralSymbol

`func (o *QueryPositionsResponsePositionsInner) SetCollateralSymbol(v string)`

SetCollateralSymbol sets CollateralSymbol field to given value.

### HasCollateralSymbol

`func (o *QueryPositionsResponsePositionsInner) HasCollateralSymbol() bool`

HasCollateralSymbol returns a boolean if a field has been set.

### GetTopicType

`func (o *QueryPositionsResponsePositionsInner) GetTopicType() string`

GetTopicType returns the TopicType field if non-nil, zero value otherwise.

### GetTopicTypeOk

`func (o *QueryPositionsResponsePositionsInner) GetTopicTypeOk() (*string, bool)`

GetTopicTypeOk returns a tuple with the TopicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicType

`func (o *QueryPositionsResponsePositionsInner) SetTopicType(v string)`

SetTopicType sets TopicType field to given value.

### HasTopicType

`func (o *QueryPositionsResponsePositionsInner) HasTopicType() bool`

HasTopicType returns a boolean if a field has been set.

### GetMarketTopicId

`func (o *QueryPositionsResponsePositionsInner) GetMarketTopicId() int64`

GetMarketTopicId returns the MarketTopicId field if non-nil, zero value otherwise.

### GetMarketTopicIdOk

`func (o *QueryPositionsResponsePositionsInner) GetMarketTopicIdOk() (*int64, bool)`

GetMarketTopicIdOk returns a tuple with the MarketTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicId

`func (o *QueryPositionsResponsePositionsInner) SetMarketTopicId(v int64)`

SetMarketTopicId sets MarketTopicId field to given value.

### HasMarketTopicId

`func (o *QueryPositionsResponsePositionsInner) HasMarketTopicId() bool`

HasMarketTopicId returns a boolean if a field has been set.

### GetMarketId

`func (o *QueryPositionsResponsePositionsInner) GetMarketId() int64`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *QueryPositionsResponsePositionsInner) GetMarketIdOk() (*int64, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *QueryPositionsResponsePositionsInner) SetMarketId(v int64)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *QueryPositionsResponsePositionsInner) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetMarketTopicTitle

`func (o *QueryPositionsResponsePositionsInner) GetMarketTopicTitle() string`

GetMarketTopicTitle returns the MarketTopicTitle field if non-nil, zero value otherwise.

### GetMarketTopicTitleOk

`func (o *QueryPositionsResponsePositionsInner) GetMarketTopicTitleOk() (*string, bool)`

GetMarketTopicTitleOk returns a tuple with the MarketTopicTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicTitle

`func (o *QueryPositionsResponsePositionsInner) SetMarketTopicTitle(v string)`

SetMarketTopicTitle sets MarketTopicTitle field to given value.

### HasMarketTopicTitle

`func (o *QueryPositionsResponsePositionsInner) HasMarketTopicTitle() bool`

HasMarketTopicTitle returns a boolean if a field has been set.

### GetMarketTitle

`func (o *QueryPositionsResponsePositionsInner) GetMarketTitle() string`

GetMarketTitle returns the MarketTitle field if non-nil, zero value otherwise.

### GetMarketTitleOk

`func (o *QueryPositionsResponsePositionsInner) GetMarketTitleOk() (*string, bool)`

GetMarketTitleOk returns a tuple with the MarketTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTitle

`func (o *QueryPositionsResponsePositionsInner) SetMarketTitle(v string)`

SetMarketTitle sets MarketTitle field to given value.

### HasMarketTitle

`func (o *QueryPositionsResponsePositionsInner) HasMarketTitle() bool`

HasMarketTitle returns a boolean if a field has been set.

### GetOutcomeName

`func (o *QueryPositionsResponsePositionsInner) GetOutcomeName() string`

GetOutcomeName returns the OutcomeName field if non-nil, zero value otherwise.

### GetOutcomeNameOk

`func (o *QueryPositionsResponsePositionsInner) GetOutcomeNameOk() (*string, bool)`

GetOutcomeNameOk returns a tuple with the OutcomeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeName

`func (o *QueryPositionsResponsePositionsInner) SetOutcomeName(v string)`

SetOutcomeName sets OutcomeName field to given value.

### HasOutcomeName

`func (o *QueryPositionsResponsePositionsInner) HasOutcomeName() bool`

HasOutcomeName returns a boolean if a field has been set.

### GetOutcomeIndex

`func (o *QueryPositionsResponsePositionsInner) GetOutcomeIndex() int32`

GetOutcomeIndex returns the OutcomeIndex field if non-nil, zero value otherwise.

### GetOutcomeIndexOk

`func (o *QueryPositionsResponsePositionsInner) GetOutcomeIndexOk() (*int32, bool)`

GetOutcomeIndexOk returns a tuple with the OutcomeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeIndex

`func (o *QueryPositionsResponsePositionsInner) SetOutcomeIndex(v int32)`

SetOutcomeIndex sets OutcomeIndex field to given value.

### HasOutcomeIndex

`func (o *QueryPositionsResponsePositionsInner) HasOutcomeIndex() bool`

HasOutcomeIndex returns a boolean if a field has been set.

### GetShares

`func (o *QueryPositionsResponsePositionsInner) GetShares() string`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *QueryPositionsResponsePositionsInner) GetSharesOk() (*string, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *QueryPositionsResponsePositionsInner) SetShares(v string)`

SetShares sets Shares field to given value.

### HasShares

`func (o *QueryPositionsResponsePositionsInner) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetAvgPrice

`func (o *QueryPositionsResponsePositionsInner) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *QueryPositionsResponsePositionsInner) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *QueryPositionsResponsePositionsInner) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *QueryPositionsResponsePositionsInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetTotalCost

`func (o *QueryPositionsResponsePositionsInner) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *QueryPositionsResponsePositionsInner) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *QueryPositionsResponsePositionsInner) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *QueryPositionsResponsePositionsInner) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetValue

`func (o *QueryPositionsResponsePositionsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryPositionsResponsePositionsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryPositionsResponsePositionsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *QueryPositionsResponsePositionsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrentPrice

`func (o *QueryPositionsResponsePositionsInner) GetCurrentPrice() string`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *QueryPositionsResponsePositionsInner) GetCurrentPriceOk() (*string, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *QueryPositionsResponsePositionsInner) SetCurrentPrice(v string)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *QueryPositionsResponsePositionsInner) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### GetToWin

`func (o *QueryPositionsResponsePositionsInner) GetToWin() string`

GetToWin returns the ToWin field if non-nil, zero value otherwise.

### GetToWinOk

`func (o *QueryPositionsResponsePositionsInner) GetToWinOk() (*string, bool)`

GetToWinOk returns a tuple with the ToWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToWin

`func (o *QueryPositionsResponsePositionsInner) SetToWin(v string)`

SetToWin sets ToWin field to given value.

### HasToWin

`func (o *QueryPositionsResponsePositionsInner) HasToWin() bool`

HasToWin returns a boolean if a field has been set.

### GetPositionStatus

`func (o *QueryPositionsResponsePositionsInner) GetPositionStatus() string`

GetPositionStatus returns the PositionStatus field if non-nil, zero value otherwise.

### GetPositionStatusOk

`func (o *QueryPositionsResponsePositionsInner) GetPositionStatusOk() (*string, bool)`

GetPositionStatusOk returns a tuple with the PositionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionStatus

`func (o *QueryPositionsResponsePositionsInner) SetPositionStatus(v string)`

SetPositionStatus sets PositionStatus field to given value.

### HasPositionStatus

`func (o *QueryPositionsResponsePositionsInner) HasPositionStatus() bool`

HasPositionStatus returns a boolean if a field has been set.

### GetCanClaim

`func (o *QueryPositionsResponsePositionsInner) GetCanClaim() bool`

GetCanClaim returns the CanClaim field if non-nil, zero value otherwise.

### GetCanClaimOk

`func (o *QueryPositionsResponsePositionsInner) GetCanClaimOk() (*bool, bool)`

GetCanClaimOk returns a tuple with the CanClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanClaim

`func (o *QueryPositionsResponsePositionsInner) SetCanClaim(v bool)`

SetCanClaim sets CanClaim field to given value.

### HasCanClaim

`func (o *QueryPositionsResponsePositionsInner) HasCanClaim() bool`

HasCanClaim returns a boolean if a field has been set.

### GetEndDate

`func (o *QueryPositionsResponsePositionsInner) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *QueryPositionsResponsePositionsInner) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *QueryPositionsResponsePositionsInner) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *QueryPositionsResponsePositionsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetUnrealizedPnl

`func (o *QueryPositionsResponsePositionsInner) GetUnrealizedPnl() string`

GetUnrealizedPnl returns the UnrealizedPnl field if non-nil, zero value otherwise.

### GetUnrealizedPnlOk

`func (o *QueryPositionsResponsePositionsInner) GetUnrealizedPnlOk() (*string, bool)`

GetUnrealizedPnlOk returns a tuple with the UnrealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPnl

`func (o *QueryPositionsResponsePositionsInner) SetUnrealizedPnl(v string)`

SetUnrealizedPnl sets UnrealizedPnl field to given value.

### HasUnrealizedPnl

`func (o *QueryPositionsResponsePositionsInner) HasUnrealizedPnl() bool`

HasUnrealizedPnl returns a boolean if a field has been set.

### GetUnrealizedPnlPercent

`func (o *QueryPositionsResponsePositionsInner) GetUnrealizedPnlPercent() string`

GetUnrealizedPnlPercent returns the UnrealizedPnlPercent field if non-nil, zero value otherwise.

### GetUnrealizedPnlPercentOk

`func (o *QueryPositionsResponsePositionsInner) GetUnrealizedPnlPercentOk() (*string, bool)`

GetUnrealizedPnlPercentOk returns a tuple with the UnrealizedPnlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPnlPercent

`func (o *QueryPositionsResponsePositionsInner) SetUnrealizedPnlPercent(v string)`

SetUnrealizedPnlPercent sets UnrealizedPnlPercent field to given value.

### HasUnrealizedPnlPercent

`func (o *QueryPositionsResponsePositionsInner) HasUnrealizedPnlPercent() bool`

HasUnrealizedPnlPercent returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *QueryPositionsResponsePositionsInner) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *QueryPositionsResponsePositionsInner) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *QueryPositionsResponsePositionsInner) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *QueryPositionsResponsePositionsInner) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetPnl

`func (o *QueryPositionsResponsePositionsInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *QueryPositionsResponsePositionsInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *QueryPositionsResponsePositionsInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *QueryPositionsResponsePositionsInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetCreatedTime

`func (o *QueryPositionsResponsePositionsInner) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *QueryPositionsResponsePositionsInner) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *QueryPositionsResponsePositionsInner) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *QueryPositionsResponsePositionsInner) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *QueryPositionsResponsePositionsInner) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *QueryPositionsResponsePositionsInner) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *QueryPositionsResponsePositionsInner) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *QueryPositionsResponsePositionsInner) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


