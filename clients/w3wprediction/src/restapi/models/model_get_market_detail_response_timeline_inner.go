/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarketDetailResponseTimelineInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarketDetailResponseTimelineInner{}

// GetMarketDetailResponseTimelineInner struct for GetMarketDetailResponseTimelineInner
type GetMarketDetailResponseTimelineInner struct {
	MarketTopicId        *int64 `json:"marketTopicId,omitempty"`
	StartDate            *int64 `json:"startDate,omitempty"`
	EndDate              *int64 `json:"endDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMarketDetailResponseTimelineInner GetMarketDetailResponseTimelineInner

// NewGetMarketDetailResponseTimelineInner instantiates a new GetMarketDetailResponseTimelineInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketDetailResponseTimelineInner() *GetMarketDetailResponseTimelineInner {
	this := GetMarketDetailResponseTimelineInner{}
	return &this
}

// NewGetMarketDetailResponseTimelineInnerWithDefaults instantiates a new GetMarketDetailResponseTimelineInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketDetailResponseTimelineInnerWithDefaults() *GetMarketDetailResponseTimelineInner {
	this := GetMarketDetailResponseTimelineInner{}
	return &this
}

// GetMarketTopicId returns the MarketTopicId field value if set, zero value otherwise.
func (o *GetMarketDetailResponseTimelineInner) GetMarketTopicId() int64 {
	if o == nil || common.IsNil(o.MarketTopicId) {
		var ret int64
		return ret
	}
	return *o.MarketTopicId
}

// GetMarketTopicIdOk returns a tuple with the MarketTopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseTimelineInner) GetMarketTopicIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketTopicId) {
		return nil, false
	}
	return o.MarketTopicId, true
}

// HasMarketTopicId returns a boolean if a field has been set.
func (o *GetMarketDetailResponseTimelineInner) HasMarketTopicId() bool {
	if o != nil && !common.IsNil(o.MarketTopicId) {
		return true
	}

	return false
}

// SetMarketTopicId gets a reference to the given int64 and assigns it to the MarketTopicId field.
func (o *GetMarketDetailResponseTimelineInner) SetMarketTopicId(v int64) {
	o.MarketTopicId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetMarketDetailResponseTimelineInner) GetStartDate() int64 {
	if o == nil || common.IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseTimelineInner) GetStartDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetMarketDetailResponseTimelineInner) HasStartDate() bool {
	if o != nil && !common.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *GetMarketDetailResponseTimelineInner) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetMarketDetailResponseTimelineInner) GetEndDate() int64 {
	if o == nil || common.IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseTimelineInner) GetEndDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetMarketDetailResponseTimelineInner) HasEndDate() bool {
	if o != nil && !common.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *GetMarketDetailResponseTimelineInner) SetEndDate(v int64) {
	o.EndDate = &v
}

func (o GetMarketDetailResponseTimelineInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketDetailResponseTimelineInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarketTopicId) {
		toSerialize["marketTopicId"] = o.MarketTopicId
	}
	if !common.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !common.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarketDetailResponseTimelineInner) UnmarshalJSON(data []byte) (err error) {
	varGetMarketDetailResponseTimelineInner := _GetMarketDetailResponseTimelineInner{}

	err = json.Unmarshal(data, &varGetMarketDetailResponseTimelineInner)

	if err != nil {
		return err
	}

	*o = GetMarketDetailResponseTimelineInner(varGetMarketDetailResponseTimelineInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marketTopicId")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarketDetailResponseTimelineInner struct {
	value *GetMarketDetailResponseTimelineInner
	isSet bool
}

func (v NullableGetMarketDetailResponseTimelineInner) Get() *GetMarketDetailResponseTimelineInner {
	return v.value
}

func (v *NullableGetMarketDetailResponseTimelineInner) Set(val *GetMarketDetailResponseTimelineInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketDetailResponseTimelineInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketDetailResponseTimelineInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketDetailResponseTimelineInner(val *GetMarketDetailResponseTimelineInner) *NullableGetMarketDetailResponseTimelineInner {
	return &NullableGetMarketDetailResponseTimelineInner{value: val, isSet: true}
}

func (v NullableGetMarketDetailResponseTimelineInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketDetailResponseTimelineInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
