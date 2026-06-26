/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarketSearchResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarketSearchResponse{}

// MarketSearchResponse struct for MarketSearchResponse
type MarketSearchResponse struct {
	Items []MarketSearchResponseInner
}

// NewMarketSearchResponse instantiates a new MarketSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketSearchResponse() *MarketSearchResponse {
	this := MarketSearchResponse{}
	return &this
}

// NewMarketSearchResponseWithDefaults instantiates a new MarketSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketSearchResponseWithDefaults() *MarketSearchResponse {
	this := MarketSearchResponse{}
	return &this
}

func (o MarketSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarketSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarketSearchResponse struct {
	value MarketSearchResponse
	isSet bool
}

func (v NullableMarketSearchResponse) Get() MarketSearchResponse {
	return v.value
}

func (v *NullableMarketSearchResponse) Set(val MarketSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketSearchResponse) Unset() {
	v.value = MarketSearchResponse{}
	v.isSet = false
}

func NewNullableMarketSearchResponse(val MarketSearchResponse) *NullableMarketSearchResponse {
	return &NullableMarketSearchResponse{value: val, isSet: true}
}

func (v NullableMarketSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
