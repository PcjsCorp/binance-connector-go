/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPositionsByFilterResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPositionsByFilterResponsePositionsInner{}

// QueryPositionsByFilterResponsePositionsInner struct for QueryPositionsByFilterResponsePositionsInner
type QueryPositionsByFilterResponsePositionsInner struct {
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
	PositionStatus       *string `json:"positionStatus,omitempty"`
	EndDate              *int64  `json:"endDate,omitempty"`
	UnrealizedPnl        *string `json:"unrealizedPnl,omitempty"`
	RealizedPnl          *string `json:"realizedPnl,omitempty"`
	Pnl                  *string `json:"pnl,omitempty"`
	CreatedTime          *int64  `json:"createdTime,omitempty"`
	UpdatedTime          *int64  `json:"updatedTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPositionsByFilterResponsePositionsInner QueryPositionsByFilterResponsePositionsInner

// NewQueryPositionsByFilterResponsePositionsInner instantiates a new QueryPositionsByFilterResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPositionsByFilterResponsePositionsInner() *QueryPositionsByFilterResponsePositionsInner {
	this := QueryPositionsByFilterResponsePositionsInner{}
	return &this
}

// NewQueryPositionsByFilterResponsePositionsInnerWithDefaults instantiates a new QueryPositionsByFilterResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPositionsByFilterResponsePositionsInnerWithDefaults() *QueryPositionsByFilterResponsePositionsInner {
	this := QueryPositionsByFilterResponsePositionsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetPositionId() int64 {
	if o == nil || common.IsNil(o.PositionId) {
		var ret int64
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetPositionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given int64 and assigns it to the PositionId field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetPositionId(v int64) {
	o.PositionId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetVendor() string {
	if o == nil || common.IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetVendorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasVendor() bool {
	if o != nil && !common.IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetChainId() string {
	if o == nil || common.IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetChainIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasChainId() bool {
	if o != nil && !common.IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetChainId(v string) {
	o.ChainId = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetCollateralSymbol returns the CollateralSymbol field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetCollateralSymbol() string {
	if o == nil || common.IsNil(o.CollateralSymbol) {
		var ret string
		return ret
	}
	return *o.CollateralSymbol
}

// GetCollateralSymbolOk returns a tuple with the CollateralSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetCollateralSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralSymbol) {
		return nil, false
	}
	return o.CollateralSymbol, true
}

// HasCollateralSymbol returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasCollateralSymbol() bool {
	if o != nil && !common.IsNil(o.CollateralSymbol) {
		return true
	}

	return false
}

// SetCollateralSymbol gets a reference to the given string and assigns it to the CollateralSymbol field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetCollateralSymbol(v string) {
	o.CollateralSymbol = &v
}

// GetTopicType returns the TopicType field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetTopicType() string {
	if o == nil || common.IsNil(o.TopicType) {
		var ret string
		return ret
	}
	return *o.TopicType
}

// GetTopicTypeOk returns a tuple with the TopicType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetTopicTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TopicType) {
		return nil, false
	}
	return o.TopicType, true
}

// HasTopicType returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasTopicType() bool {
	if o != nil && !common.IsNil(o.TopicType) {
		return true
	}

	return false
}

// SetTopicType gets a reference to the given string and assigns it to the TopicType field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetTopicType(v string) {
	o.TopicType = &v
}

// GetMarketTopicId returns the MarketTopicId field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetMarketTopicId() int64 {
	if o == nil || common.IsNil(o.MarketTopicId) {
		var ret int64
		return ret
	}
	return *o.MarketTopicId
}

// GetMarketTopicIdOk returns a tuple with the MarketTopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetMarketTopicIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketTopicId) {
		return nil, false
	}
	return o.MarketTopicId, true
}

// HasMarketTopicId returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasMarketTopicId() bool {
	if o != nil && !common.IsNil(o.MarketTopicId) {
		return true
	}

	return false
}

// SetMarketTopicId gets a reference to the given int64 and assigns it to the MarketTopicId field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetMarketTopicId(v int64) {
	o.MarketTopicId = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetMarketId() int64 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int64
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetMarketIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int64 and assigns it to the MarketId field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetMarketId(v int64) {
	o.MarketId = &v
}

// GetMarketTopicTitle returns the MarketTopicTitle field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetMarketTopicTitle() string {
	if o == nil || common.IsNil(o.MarketTopicTitle) {
		var ret string
		return ret
	}
	return *o.MarketTopicTitle
}

// GetMarketTopicTitleOk returns a tuple with the MarketTopicTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetMarketTopicTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketTopicTitle) {
		return nil, false
	}
	return o.MarketTopicTitle, true
}

// HasMarketTopicTitle returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasMarketTopicTitle() bool {
	if o != nil && !common.IsNil(o.MarketTopicTitle) {
		return true
	}

	return false
}

// SetMarketTopicTitle gets a reference to the given string and assigns it to the MarketTopicTitle field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetMarketTopicTitle(v string) {
	o.MarketTopicTitle = &v
}

// GetMarketTitle returns the MarketTitle field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetMarketTitle() string {
	if o == nil || common.IsNil(o.MarketTitle) {
		var ret string
		return ret
	}
	return *o.MarketTitle
}

// GetMarketTitleOk returns a tuple with the MarketTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetMarketTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketTitle) {
		return nil, false
	}
	return o.MarketTitle, true
}

// HasMarketTitle returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasMarketTitle() bool {
	if o != nil && !common.IsNil(o.MarketTitle) {
		return true
	}

	return false
}

// SetMarketTitle gets a reference to the given string and assigns it to the MarketTitle field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetMarketTitle(v string) {
	o.MarketTitle = &v
}

// GetOutcomeName returns the OutcomeName field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetOutcomeName() string {
	if o == nil || common.IsNil(o.OutcomeName) {
		var ret string
		return ret
	}
	return *o.OutcomeName
}

// GetOutcomeNameOk returns a tuple with the OutcomeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetOutcomeNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OutcomeName) {
		return nil, false
	}
	return o.OutcomeName, true
}

// HasOutcomeName returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasOutcomeName() bool {
	if o != nil && !common.IsNil(o.OutcomeName) {
		return true
	}

	return false
}

// SetOutcomeName gets a reference to the given string and assigns it to the OutcomeName field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetOutcomeName(v string) {
	o.OutcomeName = &v
}

// GetOutcomeIndex returns the OutcomeIndex field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetOutcomeIndex() int32 {
	if o == nil || common.IsNil(o.OutcomeIndex) {
		var ret int32
		return ret
	}
	return *o.OutcomeIndex
}

// GetOutcomeIndexOk returns a tuple with the OutcomeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetOutcomeIndexOk() (*int32, bool) {
	if o == nil || common.IsNil(o.OutcomeIndex) {
		return nil, false
	}
	return o.OutcomeIndex, true
}

// HasOutcomeIndex returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasOutcomeIndex() bool {
	if o != nil && !common.IsNil(o.OutcomeIndex) {
		return true
	}

	return false
}

// SetOutcomeIndex gets a reference to the given int32 and assigns it to the OutcomeIndex field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetOutcomeIndex(v int32) {
	o.OutcomeIndex = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetShares() string {
	if o == nil || common.IsNil(o.Shares) {
		var ret string
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetSharesOk() (*string, bool) {
	if o == nil || common.IsNil(o.Shares) {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasShares() bool {
	if o != nil && !common.IsNil(o.Shares) {
		return true
	}

	return false
}

// SetShares gets a reference to the given string and assigns it to the Shares field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetShares(v string) {
	o.Shares = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetTotalCost() string {
	if o == nil || common.IsNil(o.TotalCost) {
		var ret string
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetTotalCostOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCost) {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasTotalCost() bool {
	if o != nil && !common.IsNil(o.TotalCost) {
		return true
	}

	return false
}

// SetTotalCost gets a reference to the given string and assigns it to the TotalCost field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetTotalCost(v string) {
	o.TotalCost = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetValue() string {
	if o == nil || common.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetValue(v string) {
	o.Value = &v
}

// GetCurrentPrice returns the CurrentPrice field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetCurrentPrice() string {
	if o == nil || common.IsNil(o.CurrentPrice) {
		var ret string
		return ret
	}
	return *o.CurrentPrice
}

// GetCurrentPriceOk returns a tuple with the CurrentPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetCurrentPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentPrice) {
		return nil, false
	}
	return o.CurrentPrice, true
}

// HasCurrentPrice returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasCurrentPrice() bool {
	if o != nil && !common.IsNil(o.CurrentPrice) {
		return true
	}

	return false
}

// SetCurrentPrice gets a reference to the given string and assigns it to the CurrentPrice field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetCurrentPrice(v string) {
	o.CurrentPrice = &v
}

// GetPositionStatus returns the PositionStatus field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetPositionStatus() string {
	if o == nil || common.IsNil(o.PositionStatus) {
		var ret string
		return ret
	}
	return *o.PositionStatus
}

// GetPositionStatusOk returns a tuple with the PositionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetPositionStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionStatus) {
		return nil, false
	}
	return o.PositionStatus, true
}

// HasPositionStatus returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasPositionStatus() bool {
	if o != nil && !common.IsNil(o.PositionStatus) {
		return true
	}

	return false
}

// SetPositionStatus gets a reference to the given string and assigns it to the PositionStatus field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetPositionStatus(v string) {
	o.PositionStatus = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetEndDate() int64 {
	if o == nil || common.IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetEndDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasEndDate() bool {
	if o != nil && !common.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetUnrealizedPnl returns the UnrealizedPnl field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetUnrealizedPnl() string {
	if o == nil || common.IsNil(o.UnrealizedPnl) {
		var ret string
		return ret
	}
	return *o.UnrealizedPnl
}

// GetUnrealizedPnlOk returns a tuple with the UnrealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetUnrealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedPnl) {
		return nil, false
	}
	return o.UnrealizedPnl, true
}

// HasUnrealizedPnl returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasUnrealizedPnl() bool {
	if o != nil && !common.IsNil(o.UnrealizedPnl) {
		return true
	}

	return false
}

// SetUnrealizedPnl gets a reference to the given string and assigns it to the UnrealizedPnl field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetUnrealizedPnl(v string) {
	o.UnrealizedPnl = &v
}

// GetRealizedPnl returns the RealizedPnl field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetRealizedPnl() string {
	if o == nil || common.IsNil(o.RealizedPnl) {
		var ret string
		return ret
	}
	return *o.RealizedPnl
}

// GetRealizedPnlOk returns a tuple with the RealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetRealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.RealizedPnl) {
		return nil, false
	}
	return o.RealizedPnl, true
}

// HasRealizedPnl returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasRealizedPnl() bool {
	if o != nil && !common.IsNil(o.RealizedPnl) {
		return true
	}

	return false
}

// SetRealizedPnl gets a reference to the given string and assigns it to the RealizedPnl field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetRealizedPnl(v string) {
	o.RealizedPnl = &v
}

// GetPnl returns the Pnl field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetPnl() string {
	if o == nil || common.IsNil(o.Pnl) {
		var ret string
		return ret
	}
	return *o.Pnl
}

// GetPnlOk returns a tuple with the Pnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pnl) {
		return nil, false
	}
	return o.Pnl, true
}

// HasPnl returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasPnl() bool {
	if o != nil && !common.IsNil(o.Pnl) {
		return true
	}

	return false
}

// SetPnl gets a reference to the given string and assigns it to the Pnl field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetPnl(v string) {
	o.Pnl = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetCreatedTime() int64 {
	if o == nil || common.IsNil(o.CreatedTime) {
		var ret int64
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetCreatedTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasCreatedTime() bool {
	if o != nil && !common.IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given int64 and assigns it to the CreatedTime field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetCreatedTime(v int64) {
	o.CreatedTime = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponsePositionsInner) GetUpdatedTime() int64 {
	if o == nil || common.IsNil(o.UpdatedTime) {
		var ret int64
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) GetUpdatedTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponsePositionsInner) HasUpdatedTime() bool {
	if o != nil && !common.IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given int64 and assigns it to the UpdatedTime field.
func (o *QueryPositionsByFilterResponsePositionsInner) SetUpdatedTime(v int64) {
	o.UpdatedTime = &v
}

func (o QueryPositionsByFilterResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPositionsByFilterResponsePositionsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.PositionStatus) {
		toSerialize["positionStatus"] = o.PositionStatus
	}
	if !common.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !common.IsNil(o.UnrealizedPnl) {
		toSerialize["unrealizedPnl"] = o.UnrealizedPnl
	}
	if !common.IsNil(o.RealizedPnl) {
		toSerialize["realizedPnl"] = o.RealizedPnl
	}
	if !common.IsNil(o.Pnl) {
		toSerialize["pnl"] = o.Pnl
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

func (o *QueryPositionsByFilterResponsePositionsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryPositionsByFilterResponsePositionsInner := _QueryPositionsByFilterResponsePositionsInner{}

	err = json.Unmarshal(data, &varQueryPositionsByFilterResponsePositionsInner)

	if err != nil {
		return err
	}

	*o = QueryPositionsByFilterResponsePositionsInner(varQueryPositionsByFilterResponsePositionsInner)

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
		delete(additionalProperties, "positionStatus")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "unrealizedPnl")
		delete(additionalProperties, "realizedPnl")
		delete(additionalProperties, "pnl")
		delete(additionalProperties, "createdTime")
		delete(additionalProperties, "updatedTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPositionsByFilterResponsePositionsInner struct {
	value *QueryPositionsByFilterResponsePositionsInner
	isSet bool
}

func (v NullableQueryPositionsByFilterResponsePositionsInner) Get() *QueryPositionsByFilterResponsePositionsInner {
	return v.value
}

func (v *NullableQueryPositionsByFilterResponsePositionsInner) Set(val *QueryPositionsByFilterResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPositionsByFilterResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPositionsByFilterResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPositionsByFilterResponsePositionsInner(val *QueryPositionsByFilterResponsePositionsInner) *NullableQueryPositionsByFilterResponsePositionsInner {
	return &NullableQueryPositionsByFilterResponsePositionsInner{value: val, isSet: true}
}

func (v NullableQueryPositionsByFilterResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPositionsByFilterResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
