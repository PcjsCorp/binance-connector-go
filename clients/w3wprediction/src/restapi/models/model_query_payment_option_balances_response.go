/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPaymentOptionBalancesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPaymentOptionBalancesResponse{}

// QueryPaymentOptionBalancesResponse struct for QueryPaymentOptionBalancesResponse
type QueryPaymentOptionBalancesResponse struct {
	Items                []QueryPaymentOptionBalancesResponseItemsInner `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryPaymentOptionBalancesResponse QueryPaymentOptionBalancesResponse

// NewQueryPaymentOptionBalancesResponse instantiates a new QueryPaymentOptionBalancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPaymentOptionBalancesResponse() *QueryPaymentOptionBalancesResponse {
	this := QueryPaymentOptionBalancesResponse{}
	return &this
}

// NewQueryPaymentOptionBalancesResponseWithDefaults instantiates a new QueryPaymentOptionBalancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPaymentOptionBalancesResponseWithDefaults() *QueryPaymentOptionBalancesResponse {
	this := QueryPaymentOptionBalancesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *QueryPaymentOptionBalancesResponse) GetItems() []QueryPaymentOptionBalancesResponseItemsInner {
	if o == nil || common.IsNil(o.Items) {
		var ret []QueryPaymentOptionBalancesResponseItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPaymentOptionBalancesResponse) GetItemsOk() ([]QueryPaymentOptionBalancesResponseItemsInner, bool) {
	if o == nil || common.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QueryPaymentOptionBalancesResponse) HasItems() bool {
	if o != nil && !common.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []QueryPaymentOptionBalancesResponseItemsInner and assigns it to the Items field.
func (o *QueryPaymentOptionBalancesResponse) SetItems(v []QueryPaymentOptionBalancesResponseItemsInner) {
	o.Items = v
}

func (o QueryPaymentOptionBalancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPaymentOptionBalancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryPaymentOptionBalancesResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryPaymentOptionBalancesResponse := _QueryPaymentOptionBalancesResponse{}

	err = json.Unmarshal(data, &varQueryPaymentOptionBalancesResponse)

	if err != nil {
		return err
	}

	*o = QueryPaymentOptionBalancesResponse(varQueryPaymentOptionBalancesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryPaymentOptionBalancesResponse struct {
	value *QueryPaymentOptionBalancesResponse
	isSet bool
}

func (v NullableQueryPaymentOptionBalancesResponse) Get() *QueryPaymentOptionBalancesResponse {
	return v.value
}

func (v *NullableQueryPaymentOptionBalancesResponse) Set(val *QueryPaymentOptionBalancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPaymentOptionBalancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPaymentOptionBalancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPaymentOptionBalancesResponse(val *QueryPaymentOptionBalancesResponse) *NullableQueryPaymentOptionBalancesResponse {
	return &NullableQueryPaymentOptionBalancesResponse{value: val, isSet: true}
}

func (v NullableQueryPaymentOptionBalancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPaymentOptionBalancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
