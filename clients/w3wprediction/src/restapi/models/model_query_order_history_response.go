/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryOrderHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOrderHistoryResponse{}

// QueryOrderHistoryResponse struct for QueryOrderHistoryResponse
type QueryOrderHistoryResponse struct {
	Total                *int32                                 `json:"total,omitempty"`
	Offset               *int32                                 `json:"offset,omitempty"`
	Limit                *int32                                 `json:"limit,omitempty"`
	Orders               []QueryOrderHistoryResponseOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryOrderHistoryResponse QueryOrderHistoryResponse

// NewQueryOrderHistoryResponse instantiates a new QueryOrderHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOrderHistoryResponse() *QueryOrderHistoryResponse {
	this := QueryOrderHistoryResponse{}
	return &this
}

// NewQueryOrderHistoryResponseWithDefaults instantiates a new QueryOrderHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOrderHistoryResponseWithDefaults() *QueryOrderHistoryResponse {
	this := QueryOrderHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryOrderHistoryResponse) GetTotal() int32 {
	if o == nil || common.IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderHistoryResponse) GetTotalOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryOrderHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *QueryOrderHistoryResponse) SetTotal(v int32) {
	o.Total = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryOrderHistoryResponse) GetOffset() int32 {
	if o == nil || common.IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderHistoryResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryOrderHistoryResponse) HasOffset() bool {
	if o != nil && !common.IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *QueryOrderHistoryResponse) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryOrderHistoryResponse) GetLimit() int32 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderHistoryResponse) GetLimitOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryOrderHistoryResponse) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryOrderHistoryResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryOrderHistoryResponse) GetOrders() []QueryOrderHistoryResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryOrderHistoryResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderHistoryResponse) GetOrdersOk() ([]QueryOrderHistoryResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryOrderHistoryResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryOrderHistoryResponseOrdersInner and assigns it to the Orders field.
func (o *QueryOrderHistoryResponse) SetOrders(v []QueryOrderHistoryResponseOrdersInner) {
	o.Orders = v
}

func (o QueryOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryOrderHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryOrderHistoryResponse := _QueryOrderHistoryResponse{}

	err = json.Unmarshal(data, &varQueryOrderHistoryResponse)

	if err != nil {
		return err
	}

	*o = QueryOrderHistoryResponse(varQueryOrderHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "orders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryOrderHistoryResponse struct {
	value *QueryOrderHistoryResponse
	isSet bool
}

func (v NullableQueryOrderHistoryResponse) Get() *QueryOrderHistoryResponse {
	return v.value
}

func (v *NullableQueryOrderHistoryResponse) Set(val *QueryOrderHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOrderHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOrderHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOrderHistoryResponse(val *QueryOrderHistoryResponse) *NullableQueryOrderHistoryResponse {
	return &NullableQueryOrderHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOrderHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
