/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListPredictionWalletsResponseWalletsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListPredictionWalletsResponseWalletsInner{}

// ListPredictionWalletsResponseWalletsInner struct for ListPredictionWalletsResponseWalletsInner
type ListPredictionWalletsResponseWalletsInner struct {
	WalletAddress        *string `json:"walletAddress,omitempty"`
	WalletId             *string `json:"walletId,omitempty"`
	RegisteredTime       *int64  `json:"registeredTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPredictionWalletsResponseWalletsInner ListPredictionWalletsResponseWalletsInner

// NewListPredictionWalletsResponseWalletsInner instantiates a new ListPredictionWalletsResponseWalletsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPredictionWalletsResponseWalletsInner() *ListPredictionWalletsResponseWalletsInner {
	this := ListPredictionWalletsResponseWalletsInner{}
	return &this
}

// NewListPredictionWalletsResponseWalletsInnerWithDefaults instantiates a new ListPredictionWalletsResponseWalletsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPredictionWalletsResponseWalletsInnerWithDefaults() *ListPredictionWalletsResponseWalletsInner {
	this := ListPredictionWalletsResponseWalletsInner{}
	return &this
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *ListPredictionWalletsResponseWalletsInner) GetWalletAddress() string {
	if o == nil || common.IsNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionWalletsResponseWalletsInner) GetWalletAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletAddress) {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *ListPredictionWalletsResponseWalletsInner) HasWalletAddress() bool {
	if o != nil && !common.IsNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *ListPredictionWalletsResponseWalletsInner) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *ListPredictionWalletsResponseWalletsInner) GetWalletId() string {
	if o == nil || common.IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionWalletsResponseWalletsInner) GetWalletIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *ListPredictionWalletsResponseWalletsInner) HasWalletId() bool {
	if o != nil && !common.IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *ListPredictionWalletsResponseWalletsInner) SetWalletId(v string) {
	o.WalletId = &v
}

// GetRegisteredTime returns the RegisteredTime field value if set, zero value otherwise.
func (o *ListPredictionWalletsResponseWalletsInner) GetRegisteredTime() int64 {
	if o == nil || common.IsNil(o.RegisteredTime) {
		var ret int64
		return ret
	}
	return *o.RegisteredTime
}

// GetRegisteredTimeOk returns a tuple with the RegisteredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionWalletsResponseWalletsInner) GetRegisteredTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RegisteredTime) {
		return nil, false
	}
	return o.RegisteredTime, true
}

// HasRegisteredTime returns a boolean if a field has been set.
func (o *ListPredictionWalletsResponseWalletsInner) HasRegisteredTime() bool {
	if o != nil && !common.IsNil(o.RegisteredTime) {
		return true
	}

	return false
}

// SetRegisteredTime gets a reference to the given int64 and assigns it to the RegisteredTime field.
func (o *ListPredictionWalletsResponseWalletsInner) SetRegisteredTime(v int64) {
	o.RegisteredTime = &v
}

func (o ListPredictionWalletsResponseWalletsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPredictionWalletsResponseWalletsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.WalletAddress) {
		toSerialize["walletAddress"] = o.WalletAddress
	}
	if !common.IsNil(o.WalletId) {
		toSerialize["walletId"] = o.WalletId
	}
	if !common.IsNil(o.RegisteredTime) {
		toSerialize["registeredTime"] = o.RegisteredTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPredictionWalletsResponseWalletsInner) UnmarshalJSON(data []byte) (err error) {
	varListPredictionWalletsResponseWalletsInner := _ListPredictionWalletsResponseWalletsInner{}

	err = json.Unmarshal(data, &varListPredictionWalletsResponseWalletsInner)

	if err != nil {
		return err
	}

	*o = ListPredictionWalletsResponseWalletsInner(varListPredictionWalletsResponseWalletsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "walletAddress")
		delete(additionalProperties, "walletId")
		delete(additionalProperties, "registeredTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPredictionWalletsResponseWalletsInner struct {
	value *ListPredictionWalletsResponseWalletsInner
	isSet bool
}

func (v NullableListPredictionWalletsResponseWalletsInner) Get() *ListPredictionWalletsResponseWalletsInner {
	return v.value
}

func (v *NullableListPredictionWalletsResponseWalletsInner) Set(val *ListPredictionWalletsResponseWalletsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListPredictionWalletsResponseWalletsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListPredictionWalletsResponseWalletsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPredictionWalletsResponseWalletsInner(val *ListPredictionWalletsResponseWalletsInner) *NullableListPredictionWalletsResponseWalletsInner {
	return &NullableListPredictionWalletsResponseWalletsInner{value: val, isSet: true}
}

func (v NullableListPredictionWalletsResponseWalletsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPredictionWalletsResponseWalletsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
