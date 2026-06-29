/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PlaceOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceOrderResponse{}

// PlaceOrderResponse struct for PlaceOrderResponse
type PlaceOrderResponse struct {
	OrderId              *string `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlaceOrderResponse PlaceOrderResponse

// NewPlaceOrderResponse instantiates a new PlaceOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceOrderResponse() *PlaceOrderResponse {
	this := PlaceOrderResponse{}
	return &this
}

// NewPlaceOrderResponseWithDefaults instantiates a new PlaceOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceOrderResponseWithDefaults() *PlaceOrderResponse {
	this := PlaceOrderResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *PlaceOrderResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceOrderResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PlaceOrderResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *PlaceOrderResponse) SetOrderId(v string) {
	o.OrderId = &v
}

func (o PlaceOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlaceOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varPlaceOrderResponse := _PlaceOrderResponse{}

	err = json.Unmarshal(data, &varPlaceOrderResponse)

	if err != nil {
		return err
	}

	*o = PlaceOrderResponse(varPlaceOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlaceOrderResponse struct {
	value *PlaceOrderResponse
	isSet bool
}

func (v NullablePlaceOrderResponse) Get() *PlaceOrderResponse {
	return v.value
}

func (v *NullablePlaceOrderResponse) Set(val *PlaceOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceOrderResponse(val *PlaceOrderResponse) *NullablePlaceOrderResponse {
	return &NullablePlaceOrderResponse{value: val, isSet: true}
}

func (v NullablePlaceOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
