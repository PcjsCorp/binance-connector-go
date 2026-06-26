/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BatchRedeemResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BatchRedeemResponseResultsInner{}

// BatchRedeemResponseResultsInner struct for BatchRedeemResponseResultsInner
type BatchRedeemResponseResultsInner struct {
	RequestId            *string `json:"requestId,omitempty"`
	TxHash               *string `json:"txHash,omitempty"`
	Status               *string `json:"status,omitempty"`
	Error                *string `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchRedeemResponseResultsInner BatchRedeemResponseResultsInner

// NewBatchRedeemResponseResultsInner instantiates a new BatchRedeemResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRedeemResponseResultsInner() *BatchRedeemResponseResultsInner {
	this := BatchRedeemResponseResultsInner{}
	return &this
}

// NewBatchRedeemResponseResultsInnerWithDefaults instantiates a new BatchRedeemResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRedeemResponseResultsInnerWithDefaults() *BatchRedeemResponseResultsInner {
	this := BatchRedeemResponseResultsInner{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *BatchRedeemResponseResultsInner) GetRequestId() string {
	if o == nil || common.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRedeemResponseResultsInner) GetRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *BatchRedeemResponseResultsInner) HasRequestId() bool {
	if o != nil && !common.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *BatchRedeemResponseResultsInner) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *BatchRedeemResponseResultsInner) GetTxHash() string {
	if o == nil || common.IsNil(o.TxHash) {
		var ret string
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRedeemResponseResultsInner) GetTxHashOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *BatchRedeemResponseResultsInner) HasTxHash() bool {
	if o != nil && !common.IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given string and assigns it to the TxHash field.
func (o *BatchRedeemResponseResultsInner) SetTxHash(v string) {
	o.TxHash = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BatchRedeemResponseResultsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRedeemResponseResultsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BatchRedeemResponseResultsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BatchRedeemResponseResultsInner) SetStatus(v string) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchRedeemResponseResultsInner) GetError() string {
	if o == nil || common.IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchRedeemResponseResultsInner) GetErrorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BatchRedeemResponseResultsInner) HasError() bool {
	if o != nil && !common.IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *BatchRedeemResponseResultsInner) SetError(v string) {
	o.Error = &v
}

func (o BatchRedeemResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchRedeemResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !common.IsNil(o.TxHash) {
		toSerialize["txHash"] = o.TxHash
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchRedeemResponseResultsInner) UnmarshalJSON(data []byte) (err error) {
	varBatchRedeemResponseResultsInner := _BatchRedeemResponseResultsInner{}

	err = json.Unmarshal(data, &varBatchRedeemResponseResultsInner)

	if err != nil {
		return err
	}

	*o = BatchRedeemResponseResultsInner(varBatchRedeemResponseResultsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "txHash")
		delete(additionalProperties, "status")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchRedeemResponseResultsInner struct {
	value *BatchRedeemResponseResultsInner
	isSet bool
}

func (v NullableBatchRedeemResponseResultsInner) Get() *BatchRedeemResponseResultsInner {
	return v.value
}

func (v *NullableBatchRedeemResponseResultsInner) Set(val *BatchRedeemResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRedeemResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRedeemResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRedeemResponseResultsInner(val *BatchRedeemResponseResultsInner) *NullableBatchRedeemResponseResultsInner {
	return &NullableBatchRedeemResponseResultsInner{value: val, isSet: true}
}

func (v NullableBatchRedeemResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRedeemResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
