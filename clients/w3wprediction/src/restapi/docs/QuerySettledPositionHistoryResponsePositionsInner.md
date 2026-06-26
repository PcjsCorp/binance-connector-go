# QuerySettledPositionHistoryResponsePositionsInner

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
**IsWinner** | Pointer to **bool** |  | [optional] 
**RedeemStatus** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **int64** |  | [optional] 
**FinalOutcome** | Pointer to **string** |  | [optional] 
**RealizedPnl** | Pointer to **string** |  | [optional] 
**Pnl** | Pointer to **string** |  | [optional] 
**ClaimAmount** | Pointer to **string** |  | [optional] 
**SettledDate** | Pointer to **int64** |  | [optional] 
**CreatedTime** | Pointer to **int64** |  | [optional] 
**UpdatedTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewQuerySettledPositionHistoryResponsePositionsInner

`func NewQuerySettledPositionHistoryResponsePositionsInner() *QuerySettledPositionHistoryResponsePositionsInner`

NewQuerySettledPositionHistoryResponsePositionsInner instantiates a new QuerySettledPositionHistoryResponsePositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySettledPositionHistoryResponsePositionsInnerWithDefaults

`func NewQuerySettledPositionHistoryResponsePositionsInnerWithDefaults() *QuerySettledPositionHistoryResponsePositionsInner`

NewQuerySettledPositionHistoryResponsePositionsInnerWithDefaults instantiates a new QuerySettledPositionHistoryResponsePositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPositionId() int64`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPositionIdOk() (*int64, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetPositionId(v int64)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetVendor

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetChainId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetTokenId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetCollateralSymbol

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCollateralSymbol() string`

GetCollateralSymbol returns the CollateralSymbol field if non-nil, zero value otherwise.

### GetCollateralSymbolOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCollateralSymbolOk() (*string, bool)`

GetCollateralSymbolOk returns a tuple with the CollateralSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralSymbol

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetCollateralSymbol(v string)`

SetCollateralSymbol sets CollateralSymbol field to given value.

### HasCollateralSymbol

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasCollateralSymbol() bool`

HasCollateralSymbol returns a boolean if a field has been set.

### GetTopicType

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTopicType() string`

GetTopicType returns the TopicType field if non-nil, zero value otherwise.

### GetTopicTypeOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTopicTypeOk() (*string, bool)`

GetTopicTypeOk returns a tuple with the TopicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicType

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetTopicType(v string)`

SetTopicType sets TopicType field to given value.

### HasTopicType

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasTopicType() bool`

HasTopicType returns a boolean if a field has been set.

### GetMarketTopicId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTopicId() int64`

GetMarketTopicId returns the MarketTopicId field if non-nil, zero value otherwise.

### GetMarketTopicIdOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTopicIdOk() (*int64, bool)`

GetMarketTopicIdOk returns a tuple with the MarketTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetMarketTopicId(v int64)`

SetMarketTopicId sets MarketTopicId field to given value.

### HasMarketTopicId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasMarketTopicId() bool`

HasMarketTopicId returns a boolean if a field has been set.

### GetMarketId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketId() int64`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketIdOk() (*int64, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetMarketId(v int64)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetMarketTopicTitle

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTopicTitle() string`

GetMarketTopicTitle returns the MarketTopicTitle field if non-nil, zero value otherwise.

### GetMarketTopicTitleOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTopicTitleOk() (*string, bool)`

GetMarketTopicTitleOk returns a tuple with the MarketTopicTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTopicTitle

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetMarketTopicTitle(v string)`

SetMarketTopicTitle sets MarketTopicTitle field to given value.

### HasMarketTopicTitle

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasMarketTopicTitle() bool`

HasMarketTopicTitle returns a boolean if a field has been set.

### GetMarketTitle

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTitle() string`

GetMarketTitle returns the MarketTitle field if non-nil, zero value otherwise.

### GetMarketTitleOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTitleOk() (*string, bool)`

GetMarketTitleOk returns a tuple with the MarketTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTitle

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetMarketTitle(v string)`

SetMarketTitle sets MarketTitle field to given value.

### HasMarketTitle

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasMarketTitle() bool`

HasMarketTitle returns a boolean if a field has been set.

### GetOutcomeName

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetOutcomeName() string`

GetOutcomeName returns the OutcomeName field if non-nil, zero value otherwise.

### GetOutcomeNameOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetOutcomeNameOk() (*string, bool)`

GetOutcomeNameOk returns a tuple with the OutcomeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeName

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetOutcomeName(v string)`

SetOutcomeName sets OutcomeName field to given value.

### HasOutcomeName

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasOutcomeName() bool`

HasOutcomeName returns a boolean if a field has been set.

### GetOutcomeIndex

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetOutcomeIndex() int32`

GetOutcomeIndex returns the OutcomeIndex field if non-nil, zero value otherwise.

### GetOutcomeIndexOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetOutcomeIndexOk() (*int32, bool)`

GetOutcomeIndexOk returns a tuple with the OutcomeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeIndex

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetOutcomeIndex(v int32)`

SetOutcomeIndex sets OutcomeIndex field to given value.

### HasOutcomeIndex

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasOutcomeIndex() bool`

HasOutcomeIndex returns a boolean if a field has been set.

### GetShares

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetShares() string`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetSharesOk() (*string, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetShares(v string)`

SetShares sets Shares field to given value.

### HasShares

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetAvgPrice

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetTotalCost

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetValue

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrentPrice

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCurrentPrice() string`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCurrentPriceOk() (*string, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetCurrentPrice(v string)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### GetToWin

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetToWin() string`

GetToWin returns the ToWin field if non-nil, zero value otherwise.

### GetToWinOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetToWinOk() (*string, bool)`

GetToWinOk returns a tuple with the ToWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToWin

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetToWin(v string)`

SetToWin sets ToWin field to given value.

### HasToWin

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasToWin() bool`

HasToWin returns a boolean if a field has been set.

### GetPositionStatus

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPositionStatus() string`

GetPositionStatus returns the PositionStatus field if non-nil, zero value otherwise.

### GetPositionStatusOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPositionStatusOk() (*string, bool)`

GetPositionStatusOk returns a tuple with the PositionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionStatus

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetPositionStatus(v string)`

SetPositionStatus sets PositionStatus field to given value.

### HasPositionStatus

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasPositionStatus() bool`

HasPositionStatus returns a boolean if a field has been set.

### GetIsWinner

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetIsWinner() bool`

GetIsWinner returns the IsWinner field if non-nil, zero value otherwise.

### GetIsWinnerOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetIsWinnerOk() (*bool, bool)`

GetIsWinnerOk returns a tuple with the IsWinner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWinner

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetIsWinner(v bool)`

SetIsWinner sets IsWinner field to given value.

### HasIsWinner

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasIsWinner() bool`

HasIsWinner returns a boolean if a field has been set.

### GetRedeemStatus

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetRedeemStatus() string`

GetRedeemStatus returns the RedeemStatus field if non-nil, zero value otherwise.

### GetRedeemStatusOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetRedeemStatusOk() (*string, bool)`

GetRedeemStatusOk returns a tuple with the RedeemStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemStatus

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetRedeemStatus(v string)`

SetRedeemStatus sets RedeemStatus field to given value.

### HasRedeemStatus

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasRedeemStatus() bool`

HasRedeemStatus returns a boolean if a field has been set.

### GetEndDate

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFinalOutcome

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetFinalOutcome() string`

GetFinalOutcome returns the FinalOutcome field if non-nil, zero value otherwise.

### GetFinalOutcomeOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetFinalOutcomeOk() (*string, bool)`

GetFinalOutcomeOk returns a tuple with the FinalOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalOutcome

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetFinalOutcome(v string)`

SetFinalOutcome sets FinalOutcome field to given value.

### HasFinalOutcome

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasFinalOutcome() bool`

HasFinalOutcome returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetPnl

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetClaimAmount

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetClaimAmount() string`

GetClaimAmount returns the ClaimAmount field if non-nil, zero value otherwise.

### GetClaimAmountOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetClaimAmountOk() (*string, bool)`

GetClaimAmountOk returns a tuple with the ClaimAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimAmount

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetClaimAmount(v string)`

SetClaimAmount sets ClaimAmount field to given value.

### HasClaimAmount

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasClaimAmount() bool`

HasClaimAmount returns a boolean if a field has been set.

### GetSettledDate

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetSettledDate() int64`

GetSettledDate returns the SettledDate field if non-nil, zero value otherwise.

### GetSettledDateOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetSettledDateOk() (*int64, bool)`

GetSettledDateOk returns a tuple with the SettledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledDate

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetSettledDate(v int64)`

SetSettledDate sets SettledDate field to given value.

### HasSettledDate

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasSettledDate() bool`

HasSettledDate returns a boolean if a field has been set.

### GetCreatedTime

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *QuerySettledPositionHistoryResponsePositionsInner) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *QuerySettledPositionHistoryResponsePositionsInner) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *QuerySettledPositionHistoryResponsePositionsInner) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


