/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPortfolioResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPortfolioResponse{}

// GetPortfolioResponse struct for GetPortfolioResponse
type GetPortfolioResponse struct {
	ChainId              *string                              `json:"chainId,omitempty"`
	WalletAddress        *string                              `json:"walletAddress,omitempty"`
	ActivePositionsCount *int32                               `json:"activePositionsCount,omitempty"`
	TotalRealizedPnl     *string                              `json:"totalRealizedPnl,omitempty"`
	TotalUnrealizedPnl   *string                              `json:"totalUnrealizedPnl,omitempty"`
	TotalPnl             *string                              `json:"totalPnl,omitempty"`
	TotalCostBasis       *string                              `json:"totalCostBasis,omitempty"`
	TotalCurrentValue    *string                              `json:"totalCurrentValue,omitempty"`
	Positions            []GetPortfolioResponsePositionsInner `json:"positions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPortfolioResponse GetPortfolioResponse

// NewGetPortfolioResponse instantiates a new GetPortfolioResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortfolioResponse() *GetPortfolioResponse {
	this := GetPortfolioResponse{}
	return &this
}

// NewGetPortfolioResponseWithDefaults instantiates a new GetPortfolioResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortfolioResponseWithDefaults() *GetPortfolioResponse {
	this := GetPortfolioResponse{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetChainId() string {
	if o == nil || common.IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetChainIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasChainId() bool {
	if o != nil && !common.IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *GetPortfolioResponse) SetChainId(v string) {
	o.ChainId = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetWalletAddress() string {
	if o == nil || common.IsNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetWalletAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletAddress) {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasWalletAddress() bool {
	if o != nil && !common.IsNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *GetPortfolioResponse) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetActivePositionsCount returns the ActivePositionsCount field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetActivePositionsCount() int32 {
	if o == nil || common.IsNil(o.ActivePositionsCount) {
		var ret int32
		return ret
	}
	return *o.ActivePositionsCount
}

// GetActivePositionsCountOk returns a tuple with the ActivePositionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetActivePositionsCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.ActivePositionsCount) {
		return nil, false
	}
	return o.ActivePositionsCount, true
}

// HasActivePositionsCount returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasActivePositionsCount() bool {
	if o != nil && !common.IsNil(o.ActivePositionsCount) {
		return true
	}

	return false
}

// SetActivePositionsCount gets a reference to the given int32 and assigns it to the ActivePositionsCount field.
func (o *GetPortfolioResponse) SetActivePositionsCount(v int32) {
	o.ActivePositionsCount = &v
}

// GetTotalRealizedPnl returns the TotalRealizedPnl field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetTotalRealizedPnl() string {
	if o == nil || common.IsNil(o.TotalRealizedPnl) {
		var ret string
		return ret
	}
	return *o.TotalRealizedPnl
}

// GetTotalRealizedPnlOk returns a tuple with the TotalRealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetTotalRealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalRealizedPnl) {
		return nil, false
	}
	return o.TotalRealizedPnl, true
}

// HasTotalRealizedPnl returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasTotalRealizedPnl() bool {
	if o != nil && !common.IsNil(o.TotalRealizedPnl) {
		return true
	}

	return false
}

// SetTotalRealizedPnl gets a reference to the given string and assigns it to the TotalRealizedPnl field.
func (o *GetPortfolioResponse) SetTotalRealizedPnl(v string) {
	o.TotalRealizedPnl = &v
}

// GetTotalUnrealizedPnl returns the TotalUnrealizedPnl field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetTotalUnrealizedPnl() string {
	if o == nil || common.IsNil(o.TotalUnrealizedPnl) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedPnl
}

// GetTotalUnrealizedPnlOk returns a tuple with the TotalUnrealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetTotalUnrealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedPnl) {
		return nil, false
	}
	return o.TotalUnrealizedPnl, true
}

// HasTotalUnrealizedPnl returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasTotalUnrealizedPnl() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedPnl) {
		return true
	}

	return false
}

// SetTotalUnrealizedPnl gets a reference to the given string and assigns it to the TotalUnrealizedPnl field.
func (o *GetPortfolioResponse) SetTotalUnrealizedPnl(v string) {
	o.TotalUnrealizedPnl = &v
}

// GetTotalPnl returns the TotalPnl field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetTotalPnl() string {
	if o == nil || common.IsNil(o.TotalPnl) {
		var ret string
		return ret
	}
	return *o.TotalPnl
}

// GetTotalPnlOk returns a tuple with the TotalPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetTotalPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPnl) {
		return nil, false
	}
	return o.TotalPnl, true
}

// HasTotalPnl returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasTotalPnl() bool {
	if o != nil && !common.IsNil(o.TotalPnl) {
		return true
	}

	return false
}

// SetTotalPnl gets a reference to the given string and assigns it to the TotalPnl field.
func (o *GetPortfolioResponse) SetTotalPnl(v string) {
	o.TotalPnl = &v
}

// GetTotalCostBasis returns the TotalCostBasis field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetTotalCostBasis() string {
	if o == nil || common.IsNil(o.TotalCostBasis) {
		var ret string
		return ret
	}
	return *o.TotalCostBasis
}

// GetTotalCostBasisOk returns a tuple with the TotalCostBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetTotalCostBasisOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCostBasis) {
		return nil, false
	}
	return o.TotalCostBasis, true
}

// HasTotalCostBasis returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasTotalCostBasis() bool {
	if o != nil && !common.IsNil(o.TotalCostBasis) {
		return true
	}

	return false
}

// SetTotalCostBasis gets a reference to the given string and assigns it to the TotalCostBasis field.
func (o *GetPortfolioResponse) SetTotalCostBasis(v string) {
	o.TotalCostBasis = &v
}

// GetTotalCurrentValue returns the TotalCurrentValue field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetTotalCurrentValue() string {
	if o == nil || common.IsNil(o.TotalCurrentValue) {
		var ret string
		return ret
	}
	return *o.TotalCurrentValue
}

// GetTotalCurrentValueOk returns a tuple with the TotalCurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetTotalCurrentValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalCurrentValue) {
		return nil, false
	}
	return o.TotalCurrentValue, true
}

// HasTotalCurrentValue returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasTotalCurrentValue() bool {
	if o != nil && !common.IsNil(o.TotalCurrentValue) {
		return true
	}

	return false
}

// SetTotalCurrentValue gets a reference to the given string and assigns it to the TotalCurrentValue field.
func (o *GetPortfolioResponse) SetTotalCurrentValue(v string) {
	o.TotalCurrentValue = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *GetPortfolioResponse) GetPositions() []GetPortfolioResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []GetPortfolioResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortfolioResponse) GetPositionsOk() ([]GetPortfolioResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *GetPortfolioResponse) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []GetPortfolioResponsePositionsInner and assigns it to the Positions field.
func (o *GetPortfolioResponse) SetPositions(v []GetPortfolioResponsePositionsInner) {
	o.Positions = v
}

func (o GetPortfolioResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortfolioResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !common.IsNil(o.WalletAddress) {
		toSerialize["walletAddress"] = o.WalletAddress
	}
	if !common.IsNil(o.ActivePositionsCount) {
		toSerialize["activePositionsCount"] = o.ActivePositionsCount
	}
	if !common.IsNil(o.TotalRealizedPnl) {
		toSerialize["totalRealizedPnl"] = o.TotalRealizedPnl
	}
	if !common.IsNil(o.TotalUnrealizedPnl) {
		toSerialize["totalUnrealizedPnl"] = o.TotalUnrealizedPnl
	}
	if !common.IsNil(o.TotalPnl) {
		toSerialize["totalPnl"] = o.TotalPnl
	}
	if !common.IsNil(o.TotalCostBasis) {
		toSerialize["totalCostBasis"] = o.TotalCostBasis
	}
	if !common.IsNil(o.TotalCurrentValue) {
		toSerialize["totalCurrentValue"] = o.TotalCurrentValue
	}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPortfolioResponse) UnmarshalJSON(data []byte) (err error) {
	varGetPortfolioResponse := _GetPortfolioResponse{}

	err = json.Unmarshal(data, &varGetPortfolioResponse)

	if err != nil {
		return err
	}

	*o = GetPortfolioResponse(varGetPortfolioResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chainId")
		delete(additionalProperties, "walletAddress")
		delete(additionalProperties, "activePositionsCount")
		delete(additionalProperties, "totalRealizedPnl")
		delete(additionalProperties, "totalUnrealizedPnl")
		delete(additionalProperties, "totalPnl")
		delete(additionalProperties, "totalCostBasis")
		delete(additionalProperties, "totalCurrentValue")
		delete(additionalProperties, "positions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPortfolioResponse struct {
	value *GetPortfolioResponse
	isSet bool
}

func (v NullableGetPortfolioResponse) Get() *GetPortfolioResponse {
	return v.value
}

func (v *NullableGetPortfolioResponse) Set(val *GetPortfolioResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortfolioResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortfolioResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortfolioResponse(val *GetPortfolioResponse) *NullableGetPortfolioResponse {
	return &NullableGetPortfolioResponse{value: val, isSet: true}
}

func (v NullableGetPortfolioResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortfolioResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
