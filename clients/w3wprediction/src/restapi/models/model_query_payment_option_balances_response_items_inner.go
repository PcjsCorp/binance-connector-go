/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPaymentOptionBalancesResponseItemsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPaymentOptionBalancesResponseItemsInner{}

// QueryPaymentOptionBalancesResponseItemsInner struct for QueryPaymentOptionBalancesResponseItemsInner
type QueryPaymentOptionBalancesResponseItemsInner struct {
	AccountType             *string `json:"accountType,omitempty"`
	AvailableBalanceDisplay *string `json:"availableBalanceDisplay,omitempty"`
	Enabled                 *bool   `json:"enabled,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryPaymentOptionBalancesResponseItemsInner QueryPaymentOptionBalancesResponseItemsInner

// NewQueryPaymentOptionBalancesResponseItemsInner instantiates a new QueryPaymentOptionBalancesResponseItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPaymentOptionBalancesResponseItemsInner() *QueryPaymentOptionBalancesResponseItemsInner {
	this := QueryPaymentOptionBalancesResponseItemsInner{}
	return &this
}

// NewQueryPaymentOptionBalancesResponseItemsInnerWithDefaults instantiates a new QueryPaymentOptionBalancesResponseItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPaymentOptionBalancesResponseItemsInnerWithDefaults() *QueryPaymentOptionBalancesResponseItemsInner {
	this := QueryPaymentOptionBalancesResponseItemsInner{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *QueryPaymentOptionBalancesResponseItemsInner) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPaymentOptionBalancesResponseItemsInner) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *QueryPaymentOptionBalancesResponseItemsInner) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *QueryPaymentOptionBalancesResponseItemsInner) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAvailableBalanceDisplay returns the AvailableBalanceDisplay field value if set, zero value otherwise.
func (o *QueryPaymentOptionBalancesResponseItemsInner) GetAvailableBalanceDisplay() string {
	if o == nil || common.IsNil(o.AvailableBalanceDisplay) {
		var ret string
		return ret
	}
	return *o.AvailableBalanceDisplay
}

// GetAvailableBalanceDisplayOk returns a tuple with the AvailableBalanceDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPaymentOptionBalancesResponseItemsInner) GetAvailableBalanceDisplayOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvailableBalanceDisplay) {
		return nil, false
	}
	return o.AvailableBalanceDisplay, true
}

// HasAvailableBalanceDisplay returns a boolean if a field has been set.
func (o *QueryPaymentOptionBalancesResponseItemsInner) HasAvailableBalanceDisplay() bool {
	if o != nil && !common.IsNil(o.AvailableBalanceDisplay) {
		return true
	}

	return false
}

// SetAvailableBalanceDisplay gets a reference to the given string and assigns it to the AvailableBalanceDisplay field.
func (o *QueryPaymentOptionBalancesResponseItemsInner) SetAvailableBalanceDisplay(v string) {
	o.AvailableBalanceDisplay = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *QueryPaymentOptionBalancesResponseItemsInner) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPaymentOptionBalancesResponseItemsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *QueryPaymentOptionBalancesResponseItemsInner) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *QueryPaymentOptionBalancesResponseItemsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o QueryPaymentOptionBalancesResponseItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPaymentOptionBalancesResponseItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !common.IsNil(o.AvailableBalanceDisplay) {
		toSerialize["availableBalanceDisplay"] = o.AvailableBalanceDisplay
	}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPaymentOptionBalancesResponseItemsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryPaymentOptionBalancesResponseItemsInner := _QueryPaymentOptionBalancesResponseItemsInner{}

	err = json.Unmarshal(data, &varQueryPaymentOptionBalancesResponseItemsInner)

	if err != nil {
		return err
	}

	*o = QueryPaymentOptionBalancesResponseItemsInner(varQueryPaymentOptionBalancesResponseItemsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountType")
		delete(additionalProperties, "availableBalanceDisplay")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPaymentOptionBalancesResponseItemsInner struct {
	value *QueryPaymentOptionBalancesResponseItemsInner
	isSet bool
}

func (v NullableQueryPaymentOptionBalancesResponseItemsInner) Get() *QueryPaymentOptionBalancesResponseItemsInner {
	return v.value
}

func (v *NullableQueryPaymentOptionBalancesResponseItemsInner) Set(val *QueryPaymentOptionBalancesResponseItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPaymentOptionBalancesResponseItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPaymentOptionBalancesResponseItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPaymentOptionBalancesResponseItemsInner(val *QueryPaymentOptionBalancesResponseItemsInner) *NullableQueryPaymentOptionBalancesResponseItemsInner {
	return &NullableQueryPaymentOptionBalancesResponseItemsInner{value: val, isSet: true}
}

func (v NullableQueryPaymentOptionBalancesResponseItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPaymentOptionBalancesResponseItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
