/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySettledPositionHistoryResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySettledPositionHistoryResponsePositionsInner{}

// QuerySettledPositionHistoryResponsePositionsInner struct for QuerySettledPositionHistoryResponsePositionsInner
type QuerySettledPositionHistoryResponsePositionsInner struct {
	PositionId           *int64  `json:"positionId,omitempty"`
	Vendor               *string `json:"vendor,omitempty"`
	ChainId              *string `json:"chainId,omitempty"`
	TokenId              *string `json:"tokenId,omitempty"`
	CollateralSymbol     *string `json:"collateralSymbol,omitempty"`
	TopicType            *string `json:"topicType,omitempty"`
	MarketTopicId        *int64  `json:"marketTopicId,omitempty"`
	MarketId             *int64  `json:"marketId,omitempty"`
	MarketTopicTitle     *string `json:"marketTopicTitle,omitempty"`
	MarketTitle          *string `json:"marketTitle,omitempty"`
	OutcomeName          *string `json:"outcomeName,omitempty"`
	OutcomeIndex         *int32  `json:"outcomeIndex,omitempty"`
	Shares               *string `json:"shares,omitempty"`
	AvgPrice             *string `json:"avgPrice,omitempty"`
	TotalCost            *string `json:"totalCost,omitempty"`
	Value                *string `json:"value,omitempty"`
	CurrentPrice         *string `json:"currentPrice,omitempty"`
	ToWin                *string `json:"toWin,omitempty"`
	PositionStatus       *string `json:"positionStatus,omitempty"`
	IsWinner             *bool   `json:"isWinner,omitempty"`
	RedeemStatus         *string `json:"redeemStatus,omitempty"`
	EndDate              *int64  `json:"endDate,omitempty"`
	FinalOutcome         *string `json:"finalOutcome,omitempty"`
	RealizedPnl          *string `json:"realizedPnl,omitempty"`
	Pnl                  *string `json:"pnl,omitempty"`
	ClaimAmount          *string `json:"claimAmount,omitempty"`
	SettledDate          *int64  `json:"settledDate,omitempty"`
	CreatedTime          *int64  `json:"createdTime,omitempty"`
	UpdatedTime          *int64  `json:"updatedTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySettledPositionHistoryResponsePositionsInner QuerySettledPositionHistoryResponsePositionsInner

// NewQuerySettledPositionHistoryResponsePositionsInner instantiates a new QuerySettledPositionHistoryResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySettledPositionHistoryResponsePositionsInner() *QuerySettledPositionHistoryResponsePositionsInner {
	this := QuerySettledPositionHistoryResponsePositionsInner{}
	return &this
}

// NewQuerySettledPositionHistoryResponsePositionsInnerWithDefaults instantiates a new QuerySettledPositionHistoryResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySettledPositionHistoryResponsePositionsInnerWithDefaults() *QuerySettledPositionHistoryResponsePositionsInner {
	this := QuerySettledPositionHistoryResponsePositionsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPositionId() int64 {
	if o == nil || common.IsNil(o.PositionId) {
		var ret int64
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPositionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given int64 and assigns it to the PositionId field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetPositionId(v int64) {
	o.PositionId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetVendor() string {
	if o == nil || common.IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetVendorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasVendor() bool {
	if o != nil && !common.IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetChainId() string {
	if o == nil || common.IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetChainIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasChainId() bool {
	if o != nil && !common.IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetChainId(v string) {
	o.ChainId = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetCollateralSymbol returns the CollateralSymbol field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCollateralSymbol() string {
	if o == nil || common.IsNil(o.CollateralSymbol) {
		var ret string
		return ret
	}
	return *o.CollateralSymbol
}

// GetCollateralSymbolOk returns a tuple with the CollateralSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCollateralSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralSymbol) {
		return nil, false
	}
	return o.CollateralSymbol, true
}

// HasCollateralSymbol returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasCollateralSymbol() bool {
	if o != nil && !common.IsNil(o.CollateralSymbol) {
		return true
	}

	return false
}

// SetCollateralSymbol gets a reference to the given string and assigns it to the CollateralSymbol field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetCollateralSymbol(v string) {
	o.CollateralSymbol = &v
}

// GetTopicType returns the TopicType field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTopicType() string {
	if o == nil || common.IsNil(o.TopicType) {
		var ret string
		return ret
	}
	return *o.TopicType
}

// GetTopicTypeOk returns a tuple with the TopicType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTopicTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TopicType) {
		return nil, false
	}
	return o.TopicType, true
}

// HasTopicType returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasTopicType() bool {
	if o != nil && !common.IsNil(o.TopicType) {
		return true
	}

	return false
}

// SetTopicType gets a reference to the given string and assigns it to the TopicType field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetTopicType(v string) {
	o.TopicType = &v
}

// GetMarketTopicId returns the MarketTopicId field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTopicId() int64 {
	if o == nil || common.IsNil(o.MarketTopicId) {
		var ret int64
		return ret
	}
	return *o.MarketTopicId
}

// GetMarketTopicIdOk returns a tuple with the MarketTopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTopicIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketTopicId) {
		return nil, false
	}
	return o.MarketTopicId, true
}

// HasMarketTopicId returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasMarketTopicId() bool {
	if o != nil && !common.IsNil(o.MarketTopicId) {
		return true
	}

	return false
}

// SetMarketTopicId gets a reference to the given int64 and assigns it to the MarketTopicId field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetMarketTopicId(v int64) {
	o.MarketTopicId = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketId() int64 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int64
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int64 and assigns it to the MarketId field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetMarketId(v int64) {
	o.MarketId = &v
}

// GetMarketTopicTitle returns the MarketTopicTitle field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTopicTitle() string {
	if o == nil || common.IsNil(o.MarketTopicTitle) {
		var ret string
		return ret
	}
	return *o.MarketTopicTitle
}

// GetMarketTopicTitleOk returns a tuple with the MarketTopicTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTopicTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketTopicTitle) {
		return nil, false
	}
	return o.MarketTopicTitle, true
}

// HasMarketTopicTitle returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasMarketTopicTitle() bool {
	if o != nil && !common.IsNil(o.MarketTopicTitle) {
		return true
	}

	return false
}

// SetMarketTopicTitle gets a reference to the given string and assigns it to the MarketTopicTitle field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetMarketTopicTitle(v string) {
	o.MarketTopicTitle = &v
}

// GetMarketTitle returns the MarketTitle field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTitle() string {
	if o == nil || common.IsNil(o.MarketTitle) {
		var ret string
		return ret
	}
	return *o.MarketTitle
}

// GetMarketTitleOk returns a tuple with the MarketTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetMarketTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketTitle) {
		return nil, false
	}
	return o.MarketTitle, true
}

// HasMarketTitle returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasMarketTitle() bool {
	if o != nil && !common.IsNil(o.MarketTitle) {
		return true
	}

	return false
}

// SetMarketTitle gets a reference to the given string and assigns it to the MarketTitle field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetMarketTitle(v string) {
	o.MarketTitle = &v
}

// GetOutcomeName returns the OutcomeName field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetOutcomeName() string {
	if o == nil || common.IsNil(o.OutcomeName) {
		var ret string
		return ret
	}
	return *o.OutcomeName
}

// GetOutcomeNameOk returns a tuple with the OutcomeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetOutcomeNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OutcomeName) {
		return nil, false
	}
	return o.OutcomeName, true
}

// HasOutcomeName returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasOutcomeName() bool {
	if o != nil && !common.IsNil(o.OutcomeName) {
		return true
	}

	return false
}

// SetOutcomeName gets a reference to the given string and assigns it to the OutcomeName field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetOutcomeName(v string) {
	o.OutcomeName = &v
}

// GetOutcomeIndex returns the OutcomeIndex field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetOutcomeIndex() int32 {
	if o == nil || common.IsNil(o.OutcomeIndex) {
		var ret int32
		return ret
	}
	return *o.OutcomeIndex
}

// GetOutcomeIndexOk returns a tuple with the OutcomeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetOutcomeIndexOk() (*int32, bool) {
	if o == nil || common.IsNil(o.OutcomeIndex) {
		return nil, false
	}
	return o.OutcomeIndex, true
}

// HasOutcomeIndex returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasOutcomeIndex() bool {
	if o != nil && !common.IsNil(o.OutcomeIndex) {
		return true
	}

	return false
}

// SetOutcomeIndex gets a reference to the given int32 and assigns it to the OutcomeIndex field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetOutcomeIndex(v int32) {
	o.OutcomeIndex = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetShares() string {
	if o == nil || common.IsNil(o.Shares) {
		var ret string
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetSharesOk() (*string, bool) {
	if o == nil || common.IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasShares() bool {
	if o != nil && !common.IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given string and assigns it to the Shares field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetShares(v string) {
	o.Shares = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTotalCost() string {
	if o == nil || common.IsNil(o.TotalCost) {
		var ret string
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetTotalCostOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCost) {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasTotalCost() bool {
	if o != nil && !common.IsNil(o.TotalCost) {
		return true
	}

	return false
}

// SetTotalCost gets a reference to the given string and assigns it to the TotalCost field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetTotalCost(v string) {
	o.TotalCost = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetValue() string {
	if o == nil || common.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetValue(v string) {
	o.Value = &v
}

// GetCurrentPrice returns the CurrentPrice field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCurrentPrice() string {
	if o == nil || common.IsNil(o.CurrentPrice) {
		var ret string
		return ret
	}
	return *o.CurrentPrice
}

// GetCurrentPriceOk returns a tuple with the CurrentPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCurrentPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentPrice) {
		return nil, false
	}
	return o.CurrentPrice, true
}

// HasCurrentPrice returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasCurrentPrice() bool {
	if o != nil && !common.IsNil(o.CurrentPrice) {
		return true
	}

	return false
}

// SetCurrentPrice gets a reference to the given string and assigns it to the CurrentPrice field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetCurrentPrice(v string) {
	o.CurrentPrice = &v
}

// GetToWin returns the ToWin field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetToWin() string {
	if o == nil || common.IsNil(o.ToWin) {
		var ret string
		return ret
	}
	return *o.ToWin
}

// GetToWinOk returns a tuple with the ToWin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetToWinOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToWin) {
		return nil, false
	}
	return o.ToWin, true
}

// HasToWin returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasToWin() bool {
	if o != nil && !common.IsNil(o.ToWin) {
		return true
	}

	return false
}

// SetToWin gets a reference to the given string and assigns it to the ToWin field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetToWin(v string) {
	o.ToWin = &v
}

// GetPositionStatus returns the PositionStatus field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPositionStatus() string {
	if o == nil || common.IsNil(o.PositionStatus) {
		var ret string
		return ret
	}
	return *o.PositionStatus
}

// GetPositionStatusOk returns a tuple with the PositionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPositionStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionStatus) {
		return nil, false
	}
	return o.PositionStatus, true
}

// HasPositionStatus returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasPositionStatus() bool {
	if o != nil && !common.IsNil(o.PositionStatus) {
		return true
	}

	return false
}

// SetPositionStatus gets a reference to the given string and assigns it to the PositionStatus field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetPositionStatus(v string) {
	o.PositionStatus = &v
}

// GetIsWinner returns the IsWinner field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetIsWinner() bool {
	if o == nil || common.IsNil(o.IsWinner) {
		var ret bool
		return ret
	}
	return *o.IsWinner
}

// GetIsWinnerOk returns a tuple with the IsWinner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetIsWinnerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsWinner) {
		return nil, false
	}
	return o.IsWinner, true
}

// HasIsWinner returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasIsWinner() bool {
	if o != nil && !common.IsNil(o.IsWinner) {
		return true
	}

	return false
}

// SetIsWinner gets a reference to the given bool and assigns it to the IsWinner field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetIsWinner(v bool) {
	o.IsWinner = &v
}

// GetRedeemStatus returns the RedeemStatus field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetRedeemStatus() string {
	if o == nil || common.IsNil(o.RedeemStatus) {
		var ret string
		return ret
	}
	return *o.RedeemStatus
}

// GetRedeemStatusOk returns a tuple with the RedeemStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetRedeemStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedeemStatus) {
		return nil, false
	}
	return o.RedeemStatus, true
}

// HasRedeemStatus returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasRedeemStatus() bool {
	if o != nil && !common.IsNil(o.RedeemStatus) {
		return true
	}

	return false
}

// SetRedeemStatus gets a reference to the given string and assigns it to the RedeemStatus field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetRedeemStatus(v string) {
	o.RedeemStatus = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetEndDate() int64 {
	if o == nil || common.IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetEndDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasEndDate() bool {
	if o != nil && !common.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetFinalOutcome returns the FinalOutcome field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetFinalOutcome() string {
	if o == nil || common.IsNil(o.FinalOutcome) {
		var ret string
		return ret
	}
	return *o.FinalOutcome
}

// GetFinalOutcomeOk returns a tuple with the FinalOutcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetFinalOutcomeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FinalOutcome) {
		return nil, false
	}
	return o.FinalOutcome, true
}

// HasFinalOutcome returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasFinalOutcome() bool {
	if o != nil && !common.IsNil(o.FinalOutcome) {
		return true
	}

	return false
}

// SetFinalOutcome gets a reference to the given string and assigns it to the FinalOutcome field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetFinalOutcome(v string) {
	o.FinalOutcome = &v
}

// GetRealizedPnl returns the RealizedPnl field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetRealizedPnl() string {
	if o == nil || common.IsNil(o.RealizedPnl) {
		var ret string
		return ret
	}
	return *o.RealizedPnl
}

// GetRealizedPnlOk returns a tuple with the RealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetRealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.RealizedPnl) {
		return nil, false
	}
	return o.RealizedPnl, true
}

// HasRealizedPnl returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasRealizedPnl() bool {
	if o != nil && !common.IsNil(o.RealizedPnl) {
		return true
	}

	return false
}

// SetRealizedPnl gets a reference to the given string and assigns it to the RealizedPnl field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetRealizedPnl(v string) {
	o.RealizedPnl = &v
}

// GetPnl returns the Pnl field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPnl() string {
	if o == nil || common.IsNil(o.Pnl) {
		var ret string
		return ret
	}
	return *o.Pnl
}

// GetPnlOk returns a tuple with the Pnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pnl) {
		return nil, false
	}
	return o.Pnl, true
}

// HasPnl returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasPnl() bool {
	if o != nil && !common.IsNil(o.Pnl) {
		return true
	}

	return false
}

// SetPnl gets a reference to the given string and assigns it to the Pnl field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetPnl(v string) {
	o.Pnl = &v
}

// GetClaimAmount returns the ClaimAmount field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetClaimAmount() string {
	if o == nil || common.IsNil(o.ClaimAmount) {
		var ret string
		return ret
	}
	return *o.ClaimAmount
}

// GetClaimAmountOk returns a tuple with the ClaimAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetClaimAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClaimAmount) {
		return nil, false
	}
	return o.ClaimAmount, true
}

// HasClaimAmount returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasClaimAmount() bool {
	if o != nil && !common.IsNil(o.ClaimAmount) {
		return true
	}

	return false
}

// SetClaimAmount gets a reference to the given string and assigns it to the ClaimAmount field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetClaimAmount(v string) {
	o.ClaimAmount = &v
}

// GetSettledDate returns the SettledDate field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetSettledDate() int64 {
	if o == nil || common.IsNil(o.SettledDate) {
		var ret int64
		return ret
	}
	return *o.SettledDate
}

// GetSettledDateOk returns a tuple with the SettledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetSettledDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SettledDate) {
		return nil, false
	}
	return o.SettledDate, true
}

// HasSettledDate returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasSettledDate() bool {
	if o != nil && !common.IsNil(o.SettledDate) {
		return true
	}

	return false
}

// SetSettledDate gets a reference to the given int64 and assigns it to the SettledDate field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetSettledDate(v int64) {
	o.SettledDate = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCreatedTime() int64 {
	if o == nil || common.IsNil(o.CreatedTime) {
		var ret int64
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetCreatedTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasCreatedTime() bool {
	if o != nil && !common.IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given int64 and assigns it to the CreatedTime field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetCreatedTime(v int64) {
	o.CreatedTime = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetUpdatedTime() int64 {
	if o == nil || common.IsNil(o.UpdatedTime) {
		var ret int64
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) GetUpdatedTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponsePositionsInner) HasUpdatedTime() bool {
	if o != nil && !common.IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given int64 and assigns it to the UpdatedTime field.
func (o *QuerySettledPositionHistoryResponsePositionsInner) SetUpdatedTime(v int64) {
	o.UpdatedTime = &v
}

func (o QuerySettledPositionHistoryResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySettledPositionHistoryResponsePositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !common.IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.CollateralSymbol) {
		toSerialize["collateralSymbol"] = o.CollateralSymbol
	}
	if !common.IsNil(o.TopicType) {
		toSerialize["topicType"] = o.TopicType
	}
	if !common.IsNil(o.MarketTopicId) {
		toSerialize["marketTopicId"] = o.MarketTopicId
	}
	if !common.IsNil(o.MarketId) {
		toSerialize["marketId"] = o.MarketId
	}
	if !common.IsNil(o.MarketTopicTitle) {
		toSerialize["marketTopicTitle"] = o.MarketTopicTitle
	}
	if !common.IsNil(o.MarketTitle) {
		toSerialize["marketTitle"] = o.MarketTitle
	}
	if !common.IsNil(o.OutcomeName) {
		toSerialize["outcomeName"] = o.OutcomeName
	}
	if !common.IsNil(o.OutcomeIndex) {
		toSerialize["outcomeIndex"] = o.OutcomeIndex
	}
	if !common.IsNil(o.Shares) {
		toSerialize["shares"] = o.Shares
	}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.TotalCost) {
		toSerialize["totalCost"] = o.TotalCost
	}
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !common.IsNil(o.CurrentPrice) {
		toSerialize["currentPrice"] = o.CurrentPrice
	}
	if !common.IsNil(o.ToWin) {
		toSerialize["toWin"] = o.ToWin
	}
	if !common.IsNil(o.PositionStatus) {
		toSerialize["positionStatus"] = o.PositionStatus
	}
	if !common.IsNil(o.IsWinner) {
		toSerialize["isWinner"] = o.IsWinner
	}
	if !common.IsNil(o.RedeemStatus) {
		toSerialize["redeemStatus"] = o.RedeemStatus
	}
	if !common.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !common.IsNil(o.FinalOutcome) {
		toSerialize["finalOutcome"] = o.FinalOutcome
	}
	if !common.IsNil(o.RealizedPnl) {
		toSerialize["realizedPnl"] = o.RealizedPnl
	}
	if !common.IsNil(o.Pnl) {
		toSerialize["pnl"] = o.Pnl
	}
	if !common.IsNil(o.ClaimAmount) {
		toSerialize["claimAmount"] = o.ClaimAmount
	}
	if !common.IsNil(o.SettledDate) {
		toSerialize["settledDate"] = o.SettledDate
	}
	if !common.IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if !common.IsNil(o.UpdatedTime) {
		toSerialize["updatedTime"] = o.UpdatedTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySettledPositionHistoryResponsePositionsInner) UnmarshalJSON(data []byte) (err error) {
	varQuerySettledPositionHistoryResponsePositionsInner := _QuerySettledPositionHistoryResponsePositionsInner{}

	err = json.Unmarshal(data, &varQuerySettledPositionHistoryResponsePositionsInner)

	if err != nil {
		return err
	}

	*o = QuerySettledPositionHistoryResponsePositionsInner(varQuerySettledPositionHistoryResponsePositionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "chainId")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "collateralSymbol")
		delete(additionalProperties, "topicType")
		delete(additionalProperties, "marketTopicId")
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "marketTopicTitle")
		delete(additionalProperties, "marketTitle")
		delete(additionalProperties, "outcomeName")
		delete(additionalProperties, "outcomeIndex")
		delete(additionalProperties, "shares")
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "totalCost")
		delete(additionalProperties, "value")
		delete(additionalProperties, "currentPrice")
		delete(additionalProperties, "toWin")
		delete(additionalProperties, "positionStatus")
		delete(additionalProperties, "isWinner")
		delete(additionalProperties, "redeemStatus")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "finalOutcome")
		delete(additionalProperties, "realizedPnl")
		delete(additionalProperties, "pnl")
		delete(additionalProperties, "claimAmount")
		delete(additionalProperties, "settledDate")
		delete(additionalProperties, "createdTime")
		delete(additionalProperties, "updatedTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySettledPositionHistoryResponsePositionsInner struct {
	value *QuerySettledPositionHistoryResponsePositionsInner
	isSet bool
}

func (v NullableQuerySettledPositionHistoryResponsePositionsInner) Get() *QuerySettledPositionHistoryResponsePositionsInner {
	return v.value
}

func (v *NullableQuerySettledPositionHistoryResponsePositionsInner) Set(val *QuerySettledPositionHistoryResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySettledPositionHistoryResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySettledPositionHistoryResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySettledPositionHistoryResponsePositionsInner(val *QuerySettledPositionHistoryResponsePositionsInner) *NullableQuerySettledPositionHistoryResponsePositionsInner {
	return &NullableQuerySettledPositionHistoryResponsePositionsInner{value: val, isSet: true}
}

func (v NullableQuerySettledPositionHistoryResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySettledPositionHistoryResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
