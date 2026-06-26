/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListPredictionWalletsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListPredictionWalletsResponse{}

// ListPredictionWalletsResponse struct for ListPredictionWalletsResponse
type ListPredictionWalletsResponse struct {
	Wallets              []ListPredictionWalletsResponseWalletsInner `json:"wallets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPredictionWalletsResponse ListPredictionWalletsResponse

// NewListPredictionWalletsResponse instantiates a new ListPredictionWalletsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPredictionWalletsResponse() *ListPredictionWalletsResponse {
	this := ListPredictionWalletsResponse{}
	return &this
}

// NewListPredictionWalletsResponseWithDefaults instantiates a new ListPredictionWalletsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPredictionWalletsResponseWithDefaults() *ListPredictionWalletsResponse {
	this := ListPredictionWalletsResponse{}
	return &this
}

// GetWallets returns the Wallets field value if set, zero value otherwise.
func (o *ListPredictionWalletsResponse) GetWallets() []ListPredictionWalletsResponseWalletsInner {
	if o == nil || common.IsNil(o.Wallets) {
		var ret []ListPredictionWalletsResponseWalletsInner
		return ret
	}
	return o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionWalletsResponse) GetWalletsOk() ([]ListPredictionWalletsResponseWalletsInner, bool) {
	if o == nil || common.IsNil(o.Wallets) {
		return nil, false
	}
	return o.Wallets, true
}

// HasWallets returns a boolean if a field has been set.
func (o *ListPredictionWalletsResponse) HasWallets() bool {
	if o != nil && !common.IsNil(o.Wallets) {
		return true
	}

	return false
}

// SetWallets gets a reference to the given []ListPredictionWalletsResponseWalletsInner and assigns it to the Wallets field.
func (o *ListPredictionWalletsResponse) SetWallets(v []ListPredictionWalletsResponseWalletsInner) {
	o.Wallets = v
}

func (o ListPredictionWalletsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPredictionWalletsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Wallets) {
		toSerialize["wallets"] = o.Wallets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPredictionWalletsResponse) UnmarshalJSON(data []byte) (err error) {
	varListPredictionWalletsResponse := _ListPredictionWalletsResponse{}

	err = json.Unmarshal(data, &varListPredictionWalletsResponse)

	if err != nil {
		return err
	}

	*o = ListPredictionWalletsResponse(varListPredictionWalletsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "wallets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPredictionWalletsResponse struct {
	value *ListPredictionWalletsResponse
	isSet bool
}

func (v NullableListPredictionWalletsResponse) Get() *ListPredictionWalletsResponse {
	return v.value
}

func (v *NullableListPredictionWalletsResponse) Set(val *ListPredictionWalletsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPredictionWalletsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPredictionWalletsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPredictionWalletsResponse(val *ListPredictionWalletsResponse) *NullableListPredictionWalletsResponse {
	return &NullableListPredictionWalletsResponse{value: val, isSet: true}
}

func (v NullableListPredictionWalletsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPredictionWalletsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
