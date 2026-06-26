/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BatchCancelOrdersCancelInfoListParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BatchCancelOrdersCancelInfoListParameterInner{}

// BatchCancelOrdersCancelInfoListParameterInner struct for BatchCancelOrdersCancelInfoListParameterInner
type BatchCancelOrdersCancelInfoListParameterInner struct {
	// Internal order ID to cancel
	OrderId              string `json:"orderId"`
	AdditionalProperties map[string]interface{}
}

type _BatchCancelOrdersCancelInfoListParameterInner BatchCancelOrdersCancelInfoListParameterInner

// NewBatchCancelOrdersCancelInfoListParameterInner instantiates a new BatchCancelOrdersCancelInfoListParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchCancelOrdersCancelInfoListParameterInner(orderId string) *BatchCancelOrdersCancelInfoListParameterInner {
	this := BatchCancelOrdersCancelInfoListParameterInner{}
	this.OrderId = orderId
	return &this
}

// NewBatchCancelOrdersCancelInfoListParameterInnerWithDefaults instantiates a new BatchCancelOrdersCancelInfoListParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchCancelOrdersCancelInfoListParameterInnerWithDefaults() *BatchCancelOrdersCancelInfoListParameterInner {
	this := BatchCancelOrdersCancelInfoListParameterInner{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *BatchCancelOrdersCancelInfoListParameterInner) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *BatchCancelOrdersCancelInfoListParameterInner) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *BatchCancelOrdersCancelInfoListParameterInner) SetOrderId(v string) {
	o.OrderId = v
}

func (o BatchCancelOrdersCancelInfoListParameterInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchCancelOrdersCancelInfoListParameterInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchCancelOrdersCancelInfoListParameterInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBatchCancelOrdersCancelInfoListParameterInner := _BatchCancelOrdersCancelInfoListParameterInner{}

	err = json.Unmarshal(data, &varBatchCancelOrdersCancelInfoListParameterInner)

	if err != nil {
		return err
	}

	*o = BatchCancelOrdersCancelInfoListParameterInner(varBatchCancelOrdersCancelInfoListParameterInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchCancelOrdersCancelInfoListParameterInner struct {
	value *BatchCancelOrdersCancelInfoListParameterInner
	isSet bool
}

func (v NullableBatchCancelOrdersCancelInfoListParameterInner) Get() *BatchCancelOrdersCancelInfoListParameterInner {
	return v.value
}

func (v *NullableBatchCancelOrdersCancelInfoListParameterInner) Set(val *BatchCancelOrdersCancelInfoListParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchCancelOrdersCancelInfoListParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchCancelOrdersCancelInfoListParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchCancelOrdersCancelInfoListParameterInner(val *BatchCancelOrdersCancelInfoListParameterInner) *NullableBatchCancelOrdersCancelInfoListParameterInner {
	return &NullableBatchCancelOrdersCancelInfoListParameterInner{value: val, isSet: true}
}

func (v NullableBatchCancelOrdersCancelInfoListParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchCancelOrdersCancelInfoListParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
