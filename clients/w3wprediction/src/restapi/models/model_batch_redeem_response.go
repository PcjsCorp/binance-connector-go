/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BatchRedeemResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BatchRedeemResponse{}

// BatchRedeemResponse struct for BatchRedeemResponse
type BatchRedeemResponse struct {
	BatchId              *string                           `json:"batchId,omitempty"`
	Results              []BatchRedeemResponseResultsInner `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchRedeemResponse BatchRedeemResponse

// NewBatchRedeemResponse instantiates a new BatchRedeemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRedeemResponse() *BatchRedeemResponse {
	this := BatchRedeemResponse{}
	return &this
}

// NewBatchRedeemResponseWithDefaults instantiates a new BatchRedeemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRedeemResponseWithDefaults() *BatchRedeemResponse {
	this := BatchRedeemResponse{}
	return &this
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *BatchRedeemResponse) GetBatchId() string {
	if o == nil || common.IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRedeemResponse) GetBatchIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *BatchRedeemResponse) HasBatchId() bool {
	if o != nil && !common.IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *BatchRedeemResponse) SetBatchId(v string) {
	o.BatchId = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BatchRedeemResponse) GetResults() []BatchRedeemResponseResultsInner {
	if o == nil || common.IsNil(o.Results) {
		var ret []BatchRedeemResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRedeemResponse) GetResultsOk() ([]BatchRedeemResponseResultsInner, bool) {
	if o == nil || common.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BatchRedeemResponse) HasResults() bool {
	if o != nil && !common.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BatchRedeemResponseResultsInner and assigns it to the Results field.
func (o *BatchRedeemResponse) SetResults(v []BatchRedeemResponseResultsInner) {
	o.Results = v
}

func (o BatchRedeemResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchRedeemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	if !common.IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchRedeemResponse) UnmarshalJSON(data []byte) (err error) {
	varBatchRedeemResponse := _BatchRedeemResponse{}

	err = json.Unmarshal(data, &varBatchRedeemResponse)

	if err != nil {
		return err
	}

	*o = BatchRedeemResponse(varBatchRedeemResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "batchId")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchRedeemResponse struct {
	value *BatchRedeemResponse
	isSet bool
}

func (v NullableBatchRedeemResponse) Get() *BatchRedeemResponse {
	return v.value
}

func (v *NullableBatchRedeemResponse) Set(val *BatchRedeemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRedeemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRedeemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRedeemResponse(val *BatchRedeemResponse) *NullableBatchRedeemResponse {
	return &NullableBatchRedeemResponse{value: val, isSet: true}
}

func (v NullableBatchRedeemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRedeemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
