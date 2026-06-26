/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPositionsByFilterResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPositionsByFilterResponse{}

// QueryPositionsByFilterResponse struct for QueryPositionsByFilterResponse
type QueryPositionsByFilterResponse struct {
	Positions            []QueryPositionsByFilterResponsePositionsInner `json:"positions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPositionsByFilterResponse QueryPositionsByFilterResponse

// NewQueryPositionsByFilterResponse instantiates a new QueryPositionsByFilterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPositionsByFilterResponse() *QueryPositionsByFilterResponse {
	this := QueryPositionsByFilterResponse{}
	return &this
}

// NewQueryPositionsByFilterResponseWithDefaults instantiates a new QueryPositionsByFilterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPositionsByFilterResponseWithDefaults() *QueryPositionsByFilterResponse {
	this := QueryPositionsByFilterResponse{}
	return &this
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *QueryPositionsByFilterResponse) GetPositions() []QueryPositionsByFilterResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []QueryPositionsByFilterResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsByFilterResponse) GetPositionsOk() ([]QueryPositionsByFilterResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *QueryPositionsByFilterResponse) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []QueryPositionsByFilterResponsePositionsInner and assigns it to the Positions field.
func (o *QueryPositionsByFilterResponse) SetPositions(v []QueryPositionsByFilterResponsePositionsInner) {
	o.Positions = v
}

func (o QueryPositionsByFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPositionsByFilterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPositionsByFilterResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryPositionsByFilterResponse := _QueryPositionsByFilterResponse{}

	err = json.Unmarshal(data, &varQueryPositionsByFilterResponse)

	if err != nil {
		return err
	}

	*o = QueryPositionsByFilterResponse(varQueryPositionsByFilterResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPositionsByFilterResponse struct {
	value *QueryPositionsByFilterResponse
	isSet bool
}

func (v NullableQueryPositionsByFilterResponse) Get() *QueryPositionsByFilterResponse {
	return v.value
}

func (v *NullableQueryPositionsByFilterResponse) Set(val *QueryPositionsByFilterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPositionsByFilterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPositionsByFilterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPositionsByFilterResponse(val *QueryPositionsByFilterResponse) *NullableQueryPositionsByFilterResponse {
	return &NullableQueryPositionsByFilterResponse{value: val, isSet: true}
}

func (v NullableQueryPositionsByFilterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPositionsByFilterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
