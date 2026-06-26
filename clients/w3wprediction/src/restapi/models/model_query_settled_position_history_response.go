/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySettledPositionHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySettledPositionHistoryResponse{}

// QuerySettledPositionHistoryResponse struct for QuerySettledPositionHistoryResponse
type QuerySettledPositionHistoryResponse struct {
	TotalCount           *int32                                              `json:"totalCount,omitempty"`
	Positions            []QuerySettledPositionHistoryResponsePositionsInner `json:"positions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySettledPositionHistoryResponse QuerySettledPositionHistoryResponse

// NewQuerySettledPositionHistoryResponse instantiates a new QuerySettledPositionHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySettledPositionHistoryResponse() *QuerySettledPositionHistoryResponse {
	this := QuerySettledPositionHistoryResponse{}
	return &this
}

// NewQuerySettledPositionHistoryResponseWithDefaults instantiates a new QuerySettledPositionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySettledPositionHistoryResponseWithDefaults() *QuerySettledPositionHistoryResponse {
	this := QuerySettledPositionHistoryResponse{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponse) GetTotalCount() int32 {
	if o == nil || common.IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponse) HasTotalCount() bool {
	if o != nil && !common.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *QuerySettledPositionHistoryResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *QuerySettledPositionHistoryResponse) GetPositions() []QuerySettledPositionHistoryResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []QuerySettledPositionHistoryResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySettledPositionHistoryResponse) GetPositionsOk() ([]QuerySettledPositionHistoryResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *QuerySettledPositionHistoryResponse) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []QuerySettledPositionHistoryResponsePositionsInner and assigns it to the Positions field.
func (o *QuerySettledPositionHistoryResponse) SetPositions(v []QuerySettledPositionHistoryResponsePositionsInner) {
	o.Positions = v
}

func (o QuerySettledPositionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySettledPositionHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySettledPositionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySettledPositionHistoryResponse := _QuerySettledPositionHistoryResponse{}

	err = json.Unmarshal(data, &varQuerySettledPositionHistoryResponse)

	if err != nil {
		return err
	}

	*o = QuerySettledPositionHistoryResponse(varQuerySettledPositionHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalCount")
		delete(additionalProperties, "positions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySettledPositionHistoryResponse struct {
	value *QuerySettledPositionHistoryResponse
	isSet bool
}

func (v NullableQuerySettledPositionHistoryResponse) Get() *QuerySettledPositionHistoryResponse {
	return v.value
}

func (v *NullableQuerySettledPositionHistoryResponse) Set(val *QuerySettledPositionHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySettledPositionHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySettledPositionHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySettledPositionHistoryResponse(val *QuerySettledPositionHistoryResponse) *NullableQuerySettledPositionHistoryResponse {
	return &NullableQuerySettledPositionHistoryResponse{value: val, isSet: true}
}

func (v NullableQuerySettledPositionHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySettledPositionHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
