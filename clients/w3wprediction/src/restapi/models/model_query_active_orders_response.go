/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryActiveOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryActiveOrdersResponse{}

// QueryActiveOrdersResponse struct for QueryActiveOrdersResponse
type QueryActiveOrdersResponse struct {
	Total                *int32                                 `json:"total,omitempty"`
	Offset               *int32                                 `json:"offset,omitempty"`
	Limit                *int32                                 `json:"limit,omitempty"`
	Orders               []QueryActiveOrdersResponseOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryActiveOrdersResponse QueryActiveOrdersResponse

// NewQueryActiveOrdersResponse instantiates a new QueryActiveOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryActiveOrdersResponse() *QueryActiveOrdersResponse {
	this := QueryActiveOrdersResponse{}
	return &this
}

// NewQueryActiveOrdersResponseWithDefaults instantiates a new QueryActiveOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryActiveOrdersResponseWithDefaults() *QueryActiveOrdersResponse {
	this := QueryActiveOrdersResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponse) GetTotal() int32 {
	if o == nil || common.IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponse) GetTotalOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *QueryActiveOrdersResponse) SetTotal(v int32) {
	o.Total = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponse) GetOffset() int32 {
	if o == nil || common.IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponse) HasOffset() bool {
	if o != nil && !common.IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *QueryActiveOrdersResponse) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponse) GetLimit() int32 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponse) GetLimitOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponse) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryActiveOrdersResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponse) GetOrders() []QueryActiveOrdersResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryActiveOrdersResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponse) GetOrdersOk() ([]QueryActiveOrdersResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryActiveOrdersResponseOrdersInner and assigns it to the Orders field.
func (o *QueryActiveOrdersResponse) SetOrders(v []QueryActiveOrdersResponseOrdersInner) {
	o.Orders = v
}

func (o QueryActiveOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryActiveOrdersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryActiveOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryActiveOrdersResponse := _QueryActiveOrdersResponse{}

	err = json.Unmarshal(data, &varQueryActiveOrdersResponse)

	if err != nil {
		return err
	}

	*o = QueryActiveOrdersResponse(varQueryActiveOrdersResponse)

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

type NullableQueryActiveOrdersResponse struct {
	value *QueryActiveOrdersResponse
	isSet bool
}

func (v NullableQueryActiveOrdersResponse) Get() *QueryActiveOrdersResponse {
	return v.value
}

func (v *NullableQueryActiveOrdersResponse) Set(val *QueryActiveOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryActiveOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryActiveOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryActiveOrdersResponse(val *QueryActiveOrdersResponse) *NullableQueryActiveOrdersResponse {
	return &NullableQueryActiveOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryActiveOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryActiveOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
