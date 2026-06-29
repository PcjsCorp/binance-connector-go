/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPositionsResponseSummary type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPositionsResponseSummary{}

// QueryPositionsResponseSummary struct for QueryPositionsResponseSummary
type QueryPositionsResponseSummary struct {
	TotalValue              *string `json:"totalValue,omitempty"`
	PositionValue           *string `json:"positionValue,omitempty"`
	WalletBalance           *string `json:"walletBalance,omitempty"`
	TotalClaimableAmount    *string `json:"totalClaimableAmount,omitempty"`
	TodayRealizedPnl        *string `json:"todayRealizedPnl,omitempty"`
	TodayRealizedPnlPercent *string `json:"todayRealizedPnlPercent,omitempty"`
	TodayTotalCost          *string `json:"todayTotalCost,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryPositionsResponseSummary QueryPositionsResponseSummary

// NewQueryPositionsResponseSummary instantiates a new QueryPositionsResponseSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPositionsResponseSummary() *QueryPositionsResponseSummary {
	this := QueryPositionsResponseSummary{}
	return &this
}

// NewQueryPositionsResponseSummaryWithDefaults instantiates a new QueryPositionsResponseSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPositionsResponseSummaryWithDefaults() *QueryPositionsResponseSummary {
	this := QueryPositionsResponseSummary{}
	return &this
}

// GetTotalValue returns the TotalValue field value if set, zero value otherwise.
func (o *QueryPositionsResponseSummary) GetTotalValue() string {
	if o == nil || common.IsNil(o.TotalValue) {
		var ret string
		return ret
	}
	return *o.TotalValue
}

// GetTotalValueOk returns a tuple with the TotalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseSummary) GetTotalValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalValue) {
		return nil, false
	}
	return o.TotalValue, true
}

// HasTotalValue returns a boolean if a field has been set.
func (o *QueryPositionsResponseSummary) HasTotalValue() bool {
	if o != nil && !common.IsNil(o.TotalValue) {
		return true
	}

	return false
}

// SetTotalValue gets a reference to the given string and assigns it to the TotalValue field.
func (o *QueryPositionsResponseSummary) SetTotalValue(v string) {
	o.TotalValue = &v
}

// GetPositionValue returns the PositionValue field value if set, zero value otherwise.
func (o *QueryPositionsResponseSummary) GetPositionValue() string {
	if o == nil || common.IsNil(o.PositionValue) {
		var ret string
		return ret
	}
	return *o.PositionValue
}

// GetPositionValueOk returns a tuple with the PositionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseSummary) GetPositionValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionValue) {
		return nil, false
	}
	return o.PositionValue, true
}

// HasPositionValue returns a boolean if a field has been set.
func (o *QueryPositionsResponseSummary) HasPositionValue() bool {
	if o != nil && !common.IsNil(o.PositionValue) {
		return true
	}

	return false
}

// SetPositionValue gets a reference to the given string and assigns it to the PositionValue field.
func (o *QueryPositionsResponseSummary) SetPositionValue(v string) {
	o.PositionValue = &v
}

// GetWalletBalance returns the WalletBalance field value if set, zero value otherwise.
func (o *QueryPositionsResponseSummary) GetWalletBalance() string {
	if o == nil || common.IsNil(o.WalletBalance) {
		var ret string
		return ret
	}
	return *o.WalletBalance
}

// GetWalletBalanceOk returns a tuple with the WalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseSummary) GetWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletBalance) {
		return nil, false
	}
	return o.WalletBalance, true
}

// HasWalletBalance returns a boolean if a field has been set.
func (o *QueryPositionsResponseSummary) HasWalletBalance() bool {
	if o != nil && !common.IsNil(o.WalletBalance) {
		return true
	}

	return false
}

// SetWalletBalance gets a reference to the given string and assigns it to the WalletBalance field.
func (o *QueryPositionsResponseSummary) SetWalletBalance(v string) {
	o.WalletBalance = &v
}

// GetTotalClaimableAmount returns the TotalClaimableAmount field value if set, zero value otherwise.
func (o *QueryPositionsResponseSummary) GetTotalClaimableAmount() string {
	if o == nil || common.IsNil(o.TotalClaimableAmount) {
		var ret string
		return ret
	}
	return *o.TotalClaimableAmount
}

// GetTotalClaimableAmountOk returns a tuple with the TotalClaimableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseSummary) GetTotalClaimableAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalClaimableAmount) {
		return nil, false
	}
	return o.TotalClaimableAmount, true
}

// HasTotalClaimableAmount returns a boolean if a field has been set.
func (o *QueryPositionsResponseSummary) HasTotalClaimableAmount() bool {
	if o != nil && !common.IsNil(o.TotalClaimableAmount) {
		return true
	}

	return false
}

// SetTotalClaimableAmount gets a reference to the given string and assigns it to the TotalClaimableAmount field.
func (o *QueryPositionsResponseSummary) SetTotalClaimableAmount(v string) {
	o.TotalClaimableAmount = &v
}

// GetTodayRealizedPnl returns the TodayRealizedPnl field value if set, zero value otherwise.
func (o *QueryPositionsResponseSummary) GetTodayRealizedPnl() string {
	if o == nil || common.IsNil(o.TodayRealizedPnl) {
		var ret string
		return ret
	}
	return *o.TodayRealizedPnl
}

// GetTodayRealizedPnlOk returns a tuple with the TodayRealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseSummary) GetTodayRealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.TodayRealizedPnl) {
		return nil, false
	}
	return o.TodayRealizedPnl, true
}

// HasTodayRealizedPnl returns a boolean if a field has been set.
func (o *QueryPositionsResponseSummary) HasTodayRealizedPnl() bool {
	if o != nil && !common.IsNil(o.TodayRealizedPnl) {
		return true
	}

	return false
}

// SetTodayRealizedPnl gets a reference to the given string and assigns it to the TodayRealizedPnl field.
func (o *QueryPositionsResponseSummary) SetTodayRealizedPnl(v string) {
	o.TodayRealizedPnl = &v
}

// GetTodayRealizedPnlPercent returns the TodayRealizedPnlPercent field value if set, zero value otherwise.
func (o *QueryPositionsResponseSummary) GetTodayRealizedPnlPercent() string {
	if o == nil || common.IsNil(o.TodayRealizedPnlPercent) {
		var ret string
		return ret
	}
	return *o.TodayRealizedPnlPercent
}

// GetTodayRealizedPnlPercentOk returns a tuple with the TodayRealizedPnlPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseSummary) GetTodayRealizedPnlPercentOk() (*string, bool) {
	if o == nil || common.IsNil(o.TodayRealizedPnlPercent) {
		return nil, false
	}
	return o.TodayRealizedPnlPercent, true
}

// HasTodayRealizedPnlPercent returns a boolean if a field has been set.
func (o *QueryPositionsResponseSummary) HasTodayRealizedPnlPercent() bool {
	if o != nil && !common.IsNil(o.TodayRealizedPnlPercent) {
		return true
	}

	return false
}

// SetTodayRealizedPnlPercent gets a reference to the given string and assigns it to the TodayRealizedPnlPercent field.
func (o *QueryPositionsResponseSummary) SetTodayRealizedPnlPercent(v string) {
	o.TodayRealizedPnlPercent = &v
}

// GetTodayTotalCost returns the TodayTotalCost field value if set, zero value otherwise.
func (o *QueryPositionsResponseSummary) GetTodayTotalCost() string {
	if o == nil || common.IsNil(o.TodayTotalCost) {
		var ret string
		return ret
	}
	return *o.TodayTotalCost
}

// GetTodayTotalCostOk returns a tuple with the TodayTotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseSummary) GetTodayTotalCostOk() (*string, bool) {
	if o == nil || common.IsNil(o.TodayTotalCost) {
		return nil, false
	}
	return o.TodayTotalCost, true
}

// HasTodayTotalCost returns a boolean if a field has been set.
func (o *QueryPositionsResponseSummary) HasTodayTotalCost() bool {
	if o != nil && !common.IsNil(o.TodayTotalCost) {
		return true
	}

	return false
}

// SetTodayTotalCost gets a reference to the given string and assigns it to the TodayTotalCost field.
func (o *QueryPositionsResponseSummary) SetTodayTotalCost(v string) {
	o.TodayTotalCost = &v
}

func (o QueryPositionsResponseSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPositionsResponseSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalValue) {
		toSerialize["totalValue"] = o.TotalValue
	}
	if !common.IsNil(o.PositionValue) {
		toSerialize["positionValue"] = o.PositionValue
	}
	if !common.IsNil(o.WalletBalance) {
		toSerialize["walletBalance"] = o.WalletBalance
	}
	if !common.IsNil(o.TotalClaimableAmount) {
		toSerialize["totalClaimableAmount"] = o.TotalClaimableAmount
	}
	if !common.IsNil(o.TodayRealizedPnl) {
		toSerialize["todayRealizedPnl"] = o.TodayRealizedPnl
	}
	if !common.IsNil(o.TodayRealizedPnlPercent) {
		toSerialize["todayRealizedPnlPercent"] = o.TodayRealizedPnlPercent
	}
	if !common.IsNil(o.TodayTotalCost) {
		toSerialize["todayTotalCost"] = o.TodayTotalCost
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPositionsResponseSummary) UnmarshalJSON(data []byte) (err error) {
	varQueryPositionsResponseSummary := _QueryPositionsResponseSummary{}

	err = json.Unmarshal(data, &varQueryPositionsResponseSummary)

	if err != nil {
		return err
	}

	*o = QueryPositionsResponseSummary(varQueryPositionsResponseSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalValue")
		delete(additionalProperties, "positionValue")
		delete(additionalProperties, "walletBalance")
		delete(additionalProperties, "totalClaimableAmount")
		delete(additionalProperties, "todayRealizedPnl")
		delete(additionalProperties, "todayRealizedPnlPercent")
		delete(additionalProperties, "todayTotalCost")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPositionsResponseSummary struct {
	value *QueryPositionsResponseSummary
	isSet bool
}

func (v NullableQueryPositionsResponseSummary) Get() *QueryPositionsResponseSummary {
	return v.value
}

func (v *NullableQueryPositionsResponseSummary) Set(val *QueryPositionsResponseSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPositionsResponseSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPositionsResponseSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPositionsResponseSummary(val *QueryPositionsResponseSummary) *NullableQueryPositionsResponseSummary {
	return &NullableQueryPositionsResponseSummary{value: val, isSet: true}
}

func (v NullableQueryPositionsResponseSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPositionsResponseSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
