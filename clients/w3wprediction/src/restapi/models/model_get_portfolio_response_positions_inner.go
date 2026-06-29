/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioResponsePositionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioResponsePositionsInner{}

// GetPortfolioResponsePositionsInner struct for GetPortfolioResponsePositionsInner
type GetPortfolioResponsePositionsInner struct {
	Id                   *int64  `json:"id,omitempty"`
	WalletAddress        *string `json:"walletAddress,omitempty"`
	MarketTopicId        *int64  `json:"marketTopicId,omitempty"`
	MarketId             *int64  `json:"marketId,omitempty"`
	TokenId              *string `json:"tokenId,omitempty"`
	Vendor               *string `json:"vendor,omitempty"`
	CurrentShares        *string `json:"currentShares,omitempty"`
	AvgPrice             *string `json:"avgPrice,omitempty"`
	CurrentPrice         *string `json:"currentPrice,omitempty"`
	RealizedPnl          *string `json:"realizedPnl,omitempty"`
	UnrealizedPnl        *string `json:"unrealizedPnl,omitempty"`
	TotalPnl             *string `json:"totalPnl,omitempty"`
	PnlPercentage        *string `json:"pnlPercentage,omitempty"`
	IsResolved           *bool   `json:"isResolved,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPortfolioResponsePositionsInner GetPortfolioResponsePositionsInner

// NewGetPortfolioResponsePositionsInner instantiates a new GetPortfolioResponsePositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioResponsePositionsInner() *GetPortfolioResponsePositionsInner {
	this := GetPortfolioResponsePositionsInner{}
	return &this
}

// NewGetPortfolioResponsePositionsInnerWithDefaults instantiates a new GetPortfolioResponsePositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioResponsePositionsInnerWithDefaults() *GetPortfolioResponsePositionsInner {
	this := GetPortfolioResponsePositionsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetPortfolioResponsePositionsInner) SetId(v int64) {
	o.Id = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetWalletAddress() string {
	if o == nil || common.IsNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetWalletAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletAddress) {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasWalletAddress() bool {
	if o != nil && !common.IsNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *GetPortfolioResponsePositionsInner) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetMarketTopicId returns the MarketTopicId field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetMarketTopicId() int64 {
	if o == nil || common.IsNil(o.MarketTopicId) {
		var ret int64
		return ret
	}
	return *o.MarketTopicId
}

// GetMarketTopicIdOk returns a tuple with the MarketTopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetMarketTopicIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketTopicId) {
		return nil, false
	}
	return o.MarketTopicId, true
}

// HasMarketTopicId returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasMarketTopicId() bool {
	if o != nil && !common.IsNil(o.MarketTopicId) {
		return true
	}

	return false
}

// SetMarketTopicId gets a reference to the given int64 and assigns it to the MarketTopicId field.
func (o *GetPortfolioResponsePositionsInner) SetMarketTopicId(v int64) {
	o.MarketTopicId = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetMarketId() int64 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int64
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetMarketIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int64 and assigns it to the MarketId field.
func (o *GetPortfolioResponsePositionsInner) SetMarketId(v int64) {
	o.MarketId = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetPortfolioResponsePositionsInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetVendor() string {
	if o == nil || common.IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetVendorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasVendor() bool {
	if o != nil && !common.IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *GetPortfolioResponsePositionsInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetCurrentShares returns the CurrentShares field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetCurrentShares() string {
	if o == nil || common.IsNil(o.CurrentShares) {
		var ret string
		return ret
	}
	return *o.CurrentShares
}

// GetCurrentSharesOk returns a tuple with the CurrentShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetCurrentSharesOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentShares) {
		return nil, false
	}
	return o.CurrentShares, true
}

// HasCurrentShares returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasCurrentShares() bool {
	if o != nil && !common.IsNil(o.CurrentShares) {
		return true
	}

	return false
}

// SetCurrentShares gets a reference to the given string and assigns it to the CurrentShares field.
func (o *GetPortfolioResponsePositionsInner) SetCurrentShares(v string) {
	o.CurrentShares = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *GetPortfolioResponsePositionsInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetCurrentPrice returns the CurrentPrice field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetCurrentPrice() string {
	if o == nil || common.IsNil(o.CurrentPrice) {
		var ret string
		return ret
	}
	return *o.CurrentPrice
}

// GetCurrentPriceOk returns a tuple with the CurrentPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetCurrentPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrentPrice) {
		return nil, false
	}
	return o.CurrentPrice, true
}

// HasCurrentPrice returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasCurrentPrice() bool {
	if o != nil && !common.IsNil(o.CurrentPrice) {
		return true
	}

	return false
}

// SetCurrentPrice gets a reference to the given string and assigns it to the CurrentPrice field.
func (o *GetPortfolioResponsePositionsInner) SetCurrentPrice(v string) {
	o.CurrentPrice = &v
}

// GetRealizedPnl returns the RealizedPnl field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetRealizedPnl() string {
	if o == nil || common.IsNil(o.RealizedPnl) {
		var ret string
		return ret
	}
	return *o.RealizedPnl
}

// GetRealizedPnlOk returns a tuple with the RealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetRealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.RealizedPnl) {
		return nil, false
	}
	return o.RealizedPnl, true
}

// HasRealizedPnl returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasRealizedPnl() bool {
	if o != nil && !common.IsNil(o.RealizedPnl) {
		return true
	}

	return false
}

// SetRealizedPnl gets a reference to the given string and assigns it to the RealizedPnl field.
func (o *GetPortfolioResponsePositionsInner) SetRealizedPnl(v string) {
	o.RealizedPnl = &v
}

// GetUnrealizedPnl returns the UnrealizedPnl field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetUnrealizedPnl() string {
	if o == nil || common.IsNil(o.UnrealizedPnl) {
		var ret string
		return ret
	}
	return *o.UnrealizedPnl
}

// GetUnrealizedPnlOk returns a tuple with the UnrealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetUnrealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnrealizedPnl) {
		return nil, false
	}
	return o.UnrealizedPnl, true
}

// HasUnrealizedPnl returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasUnrealizedPnl() bool {
	if o != nil && !common.IsNil(o.UnrealizedPnl) {
		return true
	}

	return false
}

// SetUnrealizedPnl gets a reference to the given string and assigns it to the UnrealizedPnl field.
func (o *GetPortfolioResponsePositionsInner) SetUnrealizedPnl(v string) {
	o.UnrealizedPnl = &v
}

// GetTotalPnl returns the TotalPnl field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetTotalPnl() string {
	if o == nil || common.IsNil(o.TotalPnl) {
		var ret string
		return ret
	}
	return *o.TotalPnl
}

// GetTotalPnlOk returns a tuple with the TotalPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetTotalPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPnl) {
		return nil, false
	}
	return o.TotalPnl, true
}

// HasTotalPnl returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasTotalPnl() bool {
	if o != nil && !common.IsNil(o.TotalPnl) {
		return true
	}

	return false
}

// SetTotalPnl gets a reference to the given string and assigns it to the TotalPnl field.
func (o *GetPortfolioResponsePositionsInner) SetTotalPnl(v string) {
	o.TotalPnl = &v
}

// GetPnlPercentage returns the PnlPercentage field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetPnlPercentage() string {
	if o == nil || common.IsNil(o.PnlPercentage) {
		var ret string
		return ret
	}
	return *o.PnlPercentage
}

// GetPnlPercentageOk returns a tuple with the PnlPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetPnlPercentageOk() (*string, bool) {
	if o == nil || common.IsNil(o.PnlPercentage) {
		return nil, false
	}
	return o.PnlPercentage, true
}

// HasPnlPercentage returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasPnlPercentage() bool {
	if o != nil && !common.IsNil(o.PnlPercentage) {
		return true
	}

	return false
}

// SetPnlPercentage gets a reference to the given string and assigns it to the PnlPercentage field.
func (o *GetPortfolioResponsePositionsInner) SetPnlPercentage(v string) {
	o.PnlPercentage = &v
}

// GetIsResolved returns the IsResolved field value if set, zero value otherwise.
func (o *GetPortfolioResponsePositionsInner) GetIsResolved() bool {
	if o == nil || common.IsNil(o.IsResolved) {
		var ret bool
		return ret
	}
	return *o.IsResolved
}

// GetIsResolvedOk returns a tuple with the IsResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponsePositionsInner) GetIsResolvedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsResolved) {
		return nil, false
	}
	return o.IsResolved, true
}

// HasIsResolved returns a boolean if a field has been set.
func (o *GetPortfolioResponsePositionsInner) HasIsResolved() bool {
	if o != nil && !common.IsNil(o.IsResolved) {
		return true
	}

	return false
}

// SetIsResolved gets a reference to the given bool and assigns it to the IsResolved field.
func (o *GetPortfolioResponsePositionsInner) SetIsResolved(v bool) {
	o.IsResolved = &v
}

func (o GetPortfolioResponsePositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioResponsePositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.WalletAddress) {
		toSerialize["walletAddress"] = o.WalletAddress
	}
	if !common.IsNil(o.MarketTopicId) {
		toSerialize["marketTopicId"] = o.MarketTopicId
	}
	if !common.IsNil(o.MarketId) {
		toSerialize["marketId"] = o.MarketId
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !common.IsNil(o.CurrentShares) {
		toSerialize["currentShares"] = o.CurrentShares
	}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.CurrentPrice) {
		toSerialize["currentPrice"] = o.CurrentPrice
	}
	if !common.IsNil(o.RealizedPnl) {
		toSerialize["realizedPnl"] = o.RealizedPnl
	}
	if !common.IsNil(o.UnrealizedPnl) {
		toSerialize["unrealizedPnl"] = o.UnrealizedPnl
	}
	if !common.IsNil(o.TotalPnl) {
		toSerialize["totalPnl"] = o.TotalPnl
	}
	if !common.IsNil(o.PnlPercentage) {
		toSerialize["pnlPercentage"] = o.PnlPercentage
	}
	if !common.IsNil(o.IsResolved) {
		toSerialize["isResolved"] = o.IsResolved
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPortfolioResponsePositionsInner) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioResponsePositionsInner := _GetPortfolioResponsePositionsInner{}

	err = json.Unmarshal(data, &varGetPortfolioResponsePositionsInner)

	if err != nil {
		return err
	}

	*o = GetPortfolioResponsePositionsInner(varGetPortfolioResponsePositionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "walletAddress")
		delete(additionalProperties, "marketTopicId")
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "currentShares")
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "currentPrice")
		delete(additionalProperties, "realizedPnl")
		delete(additionalProperties, "unrealizedPnl")
		delete(additionalProperties, "totalPnl")
		delete(additionalProperties, "pnlPercentage")
		delete(additionalProperties, "isResolved")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioResponsePositionsInner struct {
	value *GetPortfolioResponsePositionsInner
	isSet bool
}

func (v NullableGetPortfolioResponsePositionsInner) Get() *GetPortfolioResponsePositionsInner {
	return v.value
}

func (v *NullableGetPortfolioResponsePositionsInner) Set(val *GetPortfolioResponsePositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioResponsePositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioResponsePositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioResponsePositionsInner(val *GetPortfolioResponsePositionsInner) *NullableGetPortfolioResponsePositionsInner {
	return &NullableGetPortfolioResponsePositionsInner{value: val, isSet: true}
}

func (v NullableGetPortfolioResponsePositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioResponsePositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
