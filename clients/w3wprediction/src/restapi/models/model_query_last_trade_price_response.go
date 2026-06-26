/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryLastTradePriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLastTradePriceResponse{}

// QueryLastTradePriceResponse struct for QueryLastTradePriceResponse
type QueryLastTradePriceResponse struct {
	MarketId             *int64  `json:"marketId,omitempty"`
	LastTradePrice       *string `json:"lastTradePrice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryLastTradePriceResponse QueryLastTradePriceResponse

// NewQueryLastTradePriceResponse instantiates a new QueryLastTradePriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLastTradePriceResponse() *QueryLastTradePriceResponse {
	this := QueryLastTradePriceResponse{}
	return &this
}

// NewQueryLastTradePriceResponseWithDefaults instantiates a new QueryLastTradePriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLastTradePriceResponseWithDefaults() *QueryLastTradePriceResponse {
	this := QueryLastTradePriceResponse{}
	return &this
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *QueryLastTradePriceResponse) GetMarketId() int64 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int64
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLastTradePriceResponse) GetMarketIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *QueryLastTradePriceResponse) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int64 and assigns it to the MarketId field.
func (o *QueryLastTradePriceResponse) SetMarketId(v int64) {
	o.MarketId = &v
}

// GetLastTradePrice returns the LastTradePrice field value if set, zero value otherwise.
func (o *QueryLastTradePriceResponse) GetLastTradePrice() string {
	if o == nil || common.IsNil(o.LastTradePrice) {
		var ret string
		return ret
	}
	return *o.LastTradePrice
}

// GetLastTradePriceOk returns a tuple with the LastTradePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLastTradePriceResponse) GetLastTradePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastTradePrice) {
		return nil, false
	}
	return o.LastTradePrice, true
}

// HasLastTradePrice returns a boolean if a field has been set.
func (o *QueryLastTradePriceResponse) HasLastTradePrice() bool {
	if o != nil && !common.IsNil(o.LastTradePrice) {
		return true
	}

	return false
}

// SetLastTradePrice gets a reference to the given string and assigns it to the LastTradePrice field.
func (o *QueryLastTradePriceResponse) SetLastTradePrice(v string) {
	o.LastTradePrice = &v
}

func (o QueryLastTradePriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLastTradePriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarketId) {
		toSerialize["marketId"] = o.MarketId
	}
	if !common.IsNil(o.LastTradePrice) {
		toSerialize["lastTradePrice"] = o.LastTradePrice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryLastTradePriceResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryLastTradePriceResponse := _QueryLastTradePriceResponse{}

	err = json.Unmarshal(data, &varQueryLastTradePriceResponse)

	if err != nil {
		return err
	}

	*o = QueryLastTradePriceResponse(varQueryLastTradePriceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "lastTradePrice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryLastTradePriceResponse struct {
	value *QueryLastTradePriceResponse
	isSet bool
}

func (v NullableQueryLastTradePriceResponse) Get() *QueryLastTradePriceResponse {
	return v.value
}

func (v *NullableQueryLastTradePriceResponse) Set(val *QueryLastTradePriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLastTradePriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLastTradePriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLastTradePriceResponse(val *QueryLastTradePriceResponse) *NullableQueryLastTradePriceResponse {
	return &NullableQueryLastTradePriceResponse{value: val, isSet: true}
}

func (v NullableQueryLastTradePriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLastTradePriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
