/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BatchCancelOrdersResponseFailedInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BatchCancelOrdersResponseFailedInner{}

// BatchCancelOrdersResponseFailedInner struct for BatchCancelOrdersResponseFailedInner
type BatchCancelOrdersResponseFailedInner struct {
	OrderId              *string `json:"orderId,omitempty"`
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchCancelOrdersResponseFailedInner BatchCancelOrdersResponseFailedInner

// NewBatchCancelOrdersResponseFailedInner instantiates a new BatchCancelOrdersResponseFailedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchCancelOrdersResponseFailedInner() *BatchCancelOrdersResponseFailedInner {
	this := BatchCancelOrdersResponseFailedInner{}
	return &this
}

// NewBatchCancelOrdersResponseFailedInnerWithDefaults instantiates a new BatchCancelOrdersResponseFailedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchCancelOrdersResponseFailedInnerWithDefaults() *BatchCancelOrdersResponseFailedInner {
	this := BatchCancelOrdersResponseFailedInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *BatchCancelOrdersResponseFailedInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchCancelOrdersResponseFailedInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *BatchCancelOrdersResponseFailedInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *BatchCancelOrdersResponseFailedInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BatchCancelOrdersResponseFailedInner) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchCancelOrdersResponseFailedInner) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BatchCancelOrdersResponseFailedInner) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BatchCancelOrdersResponseFailedInner) SetReason(v string) {
	o.Reason = &v
}

func (o BatchCancelOrdersResponseFailedInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchCancelOrdersResponseFailedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchCancelOrdersResponseFailedInner) UnmarshalJSON(data []byte) (err error) {
	varBatchCancelOrdersResponseFailedInner := _BatchCancelOrdersResponseFailedInner{}

	err = json.Unmarshal(data, &varBatchCancelOrdersResponseFailedInner)

	if err != nil {
		return err
	}

	*o = BatchCancelOrdersResponseFailedInner(varBatchCancelOrdersResponseFailedInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchCancelOrdersResponseFailedInner struct {
	value *BatchCancelOrdersResponseFailedInner
	isSet bool
}

func (v NullableBatchCancelOrdersResponseFailedInner) Get() *BatchCancelOrdersResponseFailedInner {
	return v.value
}

func (v *NullableBatchCancelOrdersResponseFailedInner) Set(val *BatchCancelOrdersResponseFailedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchCancelOrdersResponseFailedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchCancelOrdersResponseFailedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchCancelOrdersResponseFailedInner(val *BatchCancelOrdersResponseFailedInner) *NullableBatchCancelOrdersResponseFailedInner {
	return &NullableBatchCancelOrdersResponseFailedInner{value: val, isSet: true}
}

func (v NullableBatchCancelOrdersResponseFailedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchCancelOrdersResponseFailedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
