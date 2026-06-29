/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPositionsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPositionsResponse{}

// QueryPositionsResponse struct for QueryPositionsResponse
type QueryPositionsResponse struct {
	Summary              *QueryPositionsResponseSummary         `json:"summary,omitempty"`
	Counts               *QueryPositionsResponseCounts          `json:"counts,omitempty"`
	Positions            []QueryPositionsResponsePositionsInner `json:"positions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPositionsResponse QueryPositionsResponse

// NewQueryPositionsResponse instantiates a new QueryPositionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPositionsResponse() *QueryPositionsResponse {
	this := QueryPositionsResponse{}
	return &this
}

// NewQueryPositionsResponseWithDefaults instantiates a new QueryPositionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPositionsResponseWithDefaults() *QueryPositionsResponse {
	this := QueryPositionsResponse{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *QueryPositionsResponse) GetSummary() QueryPositionsResponseSummary {
	if o == nil || common.IsNil(o.Summary) {
		var ret QueryPositionsResponseSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponse) GetSummaryOk() (*QueryPositionsResponseSummary, bool) {
	if o == nil || common.IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *QueryPositionsResponse) HasSummary() bool {
	if o != nil && !common.IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given QueryPositionsResponseSummary and assigns it to the Summary field.
func (o *QueryPositionsResponse) SetSummary(v QueryPositionsResponseSummary) {
	o.Summary = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *QueryPositionsResponse) GetCounts() QueryPositionsResponseCounts {
	if o == nil || common.IsNil(o.Counts) {
		var ret QueryPositionsResponseCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponse) GetCountsOk() (*QueryPositionsResponseCounts, bool) {
	if o == nil || common.IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *QueryPositionsResponse) HasCounts() bool {
	if o != nil && !common.IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given QueryPositionsResponseCounts and assigns it to the Counts field.
func (o *QueryPositionsResponse) SetCounts(v QueryPositionsResponseCounts) {
	o.Counts = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *QueryPositionsResponse) GetPositions() []QueryPositionsResponsePositionsInner {
	if o == nil || common.IsNil(o.Positions) {
		var ret []QueryPositionsResponsePositionsInner
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponse) GetPositionsOk() ([]QueryPositionsResponsePositionsInner, bool) {
	if o == nil || common.IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *QueryPositionsResponse) HasPositions() bool {
	if o != nil && !common.IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []QueryPositionsResponsePositionsInner and assigns it to the Positions field.
func (o *QueryPositionsResponse) SetPositions(v []QueryPositionsResponsePositionsInner) {
	o.Positions = v
}

func (o QueryPositionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPositionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !common.IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !common.IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPositionsResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryPositionsResponse := _QueryPositionsResponse{}

	err = json.Unmarshal(data, &varQueryPositionsResponse)

	if err != nil {
		return err
	}

	*o = QueryPositionsResponse(varQueryPositionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "summary")
		delete(additionalProperties, "counts")
		delete(additionalProperties, "positions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPositionsResponse struct {
	value *QueryPositionsResponse
	isSet bool
}

func (v NullableQueryPositionsResponse) Get() *QueryPositionsResponse {
	return v.value
}

func (v *NullableQueryPositionsResponse) Set(val *QueryPositionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPositionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPositionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPositionsResponse(val *QueryPositionsResponse) *NullableQueryPositionsResponse {
	return &NullableQueryPositionsResponse{value: val, isSet: true}
}

func (v NullableQueryPositionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPositionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
