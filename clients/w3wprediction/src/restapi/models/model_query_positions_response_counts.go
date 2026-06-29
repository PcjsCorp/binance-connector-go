/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPositionsResponseCounts type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPositionsResponseCounts{}

// QueryPositionsResponseCounts struct for QueryPositionsResponseCounts
type QueryPositionsResponseCounts struct {
	OngoingCount         *int32 `json:"ongoingCount,omitempty"`
	EndedCount           *int32 `json:"endedCount,omitempty"`
	PendingClaimCount    *int32 `json:"pendingClaimCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPositionsResponseCounts QueryPositionsResponseCounts

// NewQueryPositionsResponseCounts instantiates a new QueryPositionsResponseCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPositionsResponseCounts() *QueryPositionsResponseCounts {
	this := QueryPositionsResponseCounts{}
	return &this
}

// NewQueryPositionsResponseCountsWithDefaults instantiates a new QueryPositionsResponseCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPositionsResponseCountsWithDefaults() *QueryPositionsResponseCounts {
	this := QueryPositionsResponseCounts{}
	return &this
}

// GetOngoingCount returns the OngoingCount field value if set, zero value otherwise.
func (o *QueryPositionsResponseCounts) GetOngoingCount() int32 {
	if o == nil || common.IsNil(o.OngoingCount) {
		var ret int32
		return ret
	}
	return *o.OngoingCount
}

// GetOngoingCountOk returns a tuple with the OngoingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseCounts) GetOngoingCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.OngoingCount) {
		return nil, false
	}
	return o.OngoingCount, true
}

// HasOngoingCount returns a boolean if a field has been set.
func (o *QueryPositionsResponseCounts) HasOngoingCount() bool {
	if o != nil && !common.IsNil(o.OngoingCount) {
		return true
	}

	return false
}

// SetOngoingCount gets a reference to the given int32 and assigns it to the OngoingCount field.
func (o *QueryPositionsResponseCounts) SetOngoingCount(v int32) {
	o.OngoingCount = &v
}

// GetEndedCount returns the EndedCount field value if set, zero value otherwise.
func (o *QueryPositionsResponseCounts) GetEndedCount() int32 {
	if o == nil || common.IsNil(o.EndedCount) {
		var ret int32
		return ret
	}
	return *o.EndedCount
}

// GetEndedCountOk returns a tuple with the EndedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseCounts) GetEndedCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.EndedCount) {
		return nil, false
	}
	return o.EndedCount, true
}

// HasEndedCount returns a boolean if a field has been set.
func (o *QueryPositionsResponseCounts) HasEndedCount() bool {
	if o != nil && !common.IsNil(o.EndedCount) {
		return true
	}

	return false
}

// SetEndedCount gets a reference to the given int32 and assigns it to the EndedCount field.
func (o *QueryPositionsResponseCounts) SetEndedCount(v int32) {
	o.EndedCount = &v
}

// GetPendingClaimCount returns the PendingClaimCount field value if set, zero value otherwise.
func (o *QueryPositionsResponseCounts) GetPendingClaimCount() int32 {
	if o == nil || common.IsNil(o.PendingClaimCount) {
		var ret int32
		return ret
	}
	return *o.PendingClaimCount
}

// GetPendingClaimCountOk returns a tuple with the PendingClaimCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPositionsResponseCounts) GetPendingClaimCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PendingClaimCount) {
		return nil, false
	}
	return o.PendingClaimCount, true
}

// HasPendingClaimCount returns a boolean if a field has been set.
func (o *QueryPositionsResponseCounts) HasPendingClaimCount() bool {
	if o != nil && !common.IsNil(o.PendingClaimCount) {
		return true
	}

	return false
}

// SetPendingClaimCount gets a reference to the given int32 and assigns it to the PendingClaimCount field.
func (o *QueryPositionsResponseCounts) SetPendingClaimCount(v int32) {
	o.PendingClaimCount = &v
}

func (o QueryPositionsResponseCounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPositionsResponseCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OngoingCount) {
		toSerialize["ongoingCount"] = o.OngoingCount
	}
	if !common.IsNil(o.EndedCount) {
		toSerialize["endedCount"] = o.EndedCount
	}
	if !common.IsNil(o.PendingClaimCount) {
		toSerialize["pendingClaimCount"] = o.PendingClaimCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPositionsResponseCounts) UnmarshalJSON(data []byte) (err error) {
	varQueryPositionsResponseCounts := _QueryPositionsResponseCounts{}

	err = json.Unmarshal(data, &varQueryPositionsResponseCounts)

	if err != nil {
		return err
	}

	*o = QueryPositionsResponseCounts(varQueryPositionsResponseCounts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ongoingCount")
		delete(additionalProperties, "endedCount")
		delete(additionalProperties, "pendingClaimCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPositionsResponseCounts struct {
	value *QueryPositionsResponseCounts
	isSet bool
}

func (v NullableQueryPositionsResponseCounts) Get() *QueryPositionsResponseCounts {
	return v.value
}

func (v *NullableQueryPositionsResponseCounts) Set(val *QueryPositionsResponseCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPositionsResponseCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPositionsResponseCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPositionsResponseCounts(val *QueryPositionsResponseCounts) *NullableQueryPositionsResponseCounts {
	return &NullableQueryPositionsResponseCounts{value: val, isSet: true}
}

func (v NullableQueryPositionsResponseCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPositionsResponseCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
