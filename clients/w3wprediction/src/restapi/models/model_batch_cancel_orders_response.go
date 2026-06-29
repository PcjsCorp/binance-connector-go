/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BatchCancelOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BatchCancelOrdersResponse{}

// BatchCancelOrdersResponse struct for BatchCancelOrdersResponse
type BatchCancelOrdersResponse struct {
	Canceled             []string                               `json:"canceled,omitempty"`
	Failed               []BatchCancelOrdersResponseFailedInner `json:"failed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchCancelOrdersResponse BatchCancelOrdersResponse

// NewBatchCancelOrdersResponse instantiates a new BatchCancelOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchCancelOrdersResponse() *BatchCancelOrdersResponse {
	this := BatchCancelOrdersResponse{}
	return &this
}

// NewBatchCancelOrdersResponseWithDefaults instantiates a new BatchCancelOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchCancelOrdersResponseWithDefaults() *BatchCancelOrdersResponse {
	this := BatchCancelOrdersResponse{}
	return &this
}

// GetCanceled returns the Canceled field value if set, zero value otherwise.
func (o *BatchCancelOrdersResponse) GetCanceled() []string {
	if o == nil || common.IsNil(o.Canceled) {
		var ret []string
		return ret
	}
	return o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchCancelOrdersResponse) GetCanceledOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Canceled) {
		return nil, false
	}
	return o.Canceled, true
}

// HasCanceled returns a boolean if a field has been set.
func (o *BatchCancelOrdersResponse) HasCanceled() bool {
	if o != nil && !common.IsNil(o.Canceled) {
		return true
	}

	return false
}

// SetCanceled gets a reference to the given []string and assigns it to the Canceled field.
func (o *BatchCancelOrdersResponse) SetCanceled(v []string) {
	o.Canceled = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *BatchCancelOrdersResponse) GetFailed() []BatchCancelOrdersResponseFailedInner {
	if o == nil || common.IsNil(o.Failed) {
		var ret []BatchCancelOrdersResponseFailedInner
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchCancelOrdersResponse) GetFailedOk() ([]BatchCancelOrdersResponseFailedInner, bool) {
	if o == nil || common.IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *BatchCancelOrdersResponse) HasFailed() bool {
	if o != nil && !common.IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given []BatchCancelOrdersResponseFailedInner and assigns it to the Failed field.
func (o *BatchCancelOrdersResponse) SetFailed(v []BatchCancelOrdersResponseFailedInner) {
	o.Failed = v
}

func (o BatchCancelOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchCancelOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Canceled) {
		toSerialize["canceled"] = o.Canceled
	}
	if !common.IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchCancelOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varBatchCancelOrdersResponse := _BatchCancelOrdersResponse{}

	err = json.Unmarshal(data, &varBatchCancelOrdersResponse)

	if err != nil {
		return err
	}

	*o = BatchCancelOrdersResponse(varBatchCancelOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "canceled")
		delete(additionalProperties, "failed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchCancelOrdersResponse struct {
	value *BatchCancelOrdersResponse
	isSet bool
}

func (v NullableBatchCancelOrdersResponse) Get() *BatchCancelOrdersResponse {
	return v.value
}

func (v *NullableBatchCancelOrdersResponse) Set(val *BatchCancelOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchCancelOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchCancelOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchCancelOrdersResponse(val *BatchCancelOrdersResponse) *NullableBatchCancelOrdersResponse {
	return &NullableBatchCancelOrdersResponse{value: val, isSet: true}
}

func (v NullableBatchCancelOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchCancelOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
