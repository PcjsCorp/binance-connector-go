/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListPredictionMarketsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListPredictionMarketsResponse{}

// ListPredictionMarketsResponse struct for ListPredictionMarketsResponse
type ListPredictionMarketsResponse struct {
	MarketTopics         []ListPredictionMarketsResponseMarketTopicsInner `json:"marketTopics,omitempty"`
	Total                *int32                                           `json:"total,omitempty"`
	Offset               *int32                                           `json:"offset,omitempty"`
	Limit                *int32                                           `json:"limit,omitempty"`
	HasMore              *bool                                            `json:"hasMore,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPredictionMarketsResponse ListPredictionMarketsResponse

// NewListPredictionMarketsResponse instantiates a new ListPredictionMarketsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPredictionMarketsResponse() *ListPredictionMarketsResponse {
	this := ListPredictionMarketsResponse{}
	return &this
}

// NewListPredictionMarketsResponseWithDefaults instantiates a new ListPredictionMarketsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPredictionMarketsResponseWithDefaults() *ListPredictionMarketsResponse {
	this := ListPredictionMarketsResponse{}
	return &this
}

// GetMarketTopics returns the MarketTopics field value if set, zero value otherwise.
func (o *ListPredictionMarketsResponse) GetMarketTopics() []ListPredictionMarketsResponseMarketTopicsInner {
	if o == nil || common.IsNil(o.MarketTopics) {
		var ret []ListPredictionMarketsResponseMarketTopicsInner
		return ret
	}
	return o.MarketTopics
}

// GetMarketTopicsOk returns a tuple with the MarketTopics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionMarketsResponse) GetMarketTopicsOk() ([]ListPredictionMarketsResponseMarketTopicsInner, bool) {
	if o == nil || common.IsNil(o.MarketTopics) {
		return nil, false
	}
	return o.MarketTopics, true
}

// HasMarketTopics returns a boolean if a field has been set.
func (o *ListPredictionMarketsResponse) HasMarketTopics() bool {
	if o != nil && !common.IsNil(o.MarketTopics) {
		return true
	}

	return false
}

// SetMarketTopics gets a reference to the given []ListPredictionMarketsResponseMarketTopicsInner and assigns it to the MarketTopics field.
func (o *ListPredictionMarketsResponse) SetMarketTopics(v []ListPredictionMarketsResponseMarketTopicsInner) {
	o.MarketTopics = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ListPredictionMarketsResponse) GetTotal() int32 {
	if o == nil || common.IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionMarketsResponse) GetTotalOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ListPredictionMarketsResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ListPredictionMarketsResponse) SetTotal(v int32) {
	o.Total = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListPredictionMarketsResponse) GetOffset() int32 {
	if o == nil || common.IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionMarketsResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListPredictionMarketsResponse) HasOffset() bool {
	if o != nil && !common.IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListPredictionMarketsResponse) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListPredictionMarketsResponse) GetLimit() int32 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionMarketsResponse) GetLimitOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListPredictionMarketsResponse) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListPredictionMarketsResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ListPredictionMarketsResponse) GetHasMore() bool {
	if o == nil || common.IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionMarketsResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ListPredictionMarketsResponse) HasHasMore() bool {
	if o != nil && !common.IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ListPredictionMarketsResponse) SetHasMore(v bool) {
	o.HasMore = &v
}

func (o ListPredictionMarketsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPredictionMarketsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarketTopics) {
		toSerialize["marketTopics"] = o.MarketTopics
	}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPredictionMarketsResponse) UnmarshalJSON(data []byte) (err error) {
	varListPredictionMarketsResponse := _ListPredictionMarketsResponse{}

	err = json.Unmarshal(data, &varListPredictionMarketsResponse)

	if err != nil {
		return err
	}

	*o = ListPredictionMarketsResponse(varListPredictionMarketsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marketTopics")
		delete(additionalProperties, "total")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "hasMore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPredictionMarketsResponse struct {
	value *ListPredictionMarketsResponse
	isSet bool
}

func (v NullableListPredictionMarketsResponse) Get() *ListPredictionMarketsResponse {
	return v.value
}

func (v *NullableListPredictionMarketsResponse) Set(val *ListPredictionMarketsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPredictionMarketsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPredictionMarketsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPredictionMarketsResponse(val *ListPredictionMarketsResponse) *NullableListPredictionMarketsResponse {
	return &NullableListPredictionMarketsResponse{value: val, isSet: true}
}

func (v NullableListPredictionMarketsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPredictionMarketsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
