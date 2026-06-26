/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryOrderBookResponseBidsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOrderBookResponseBidsInner{}

// QueryOrderBookResponseBidsInner struct for QueryOrderBookResponseBidsInner
type QueryOrderBookResponseBidsInner struct {
	Price                *string `json:"price,omitempty"`
	Size                 *string `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryOrderBookResponseBidsInner QueryOrderBookResponseBidsInner

// NewQueryOrderBookResponseBidsInner instantiates a new QueryOrderBookResponseBidsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOrderBookResponseBidsInner() *QueryOrderBookResponseBidsInner {
	this := QueryOrderBookResponseBidsInner{}
	return &this
}

// NewQueryOrderBookResponseBidsInnerWithDefaults instantiates a new QueryOrderBookResponseBidsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOrderBookResponseBidsInnerWithDefaults() *QueryOrderBookResponseBidsInner {
	this := QueryOrderBookResponseBidsInner{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryOrderBookResponseBidsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponseBidsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryOrderBookResponseBidsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryOrderBookResponseBidsInner) SetPrice(v string) {
	o.Price = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *QueryOrderBookResponseBidsInner) GetSize() string {
	if o == nil || common.IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponseBidsInner) GetSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *QueryOrderBookResponseBidsInner) HasSize() bool {
	if o != nil && !common.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *QueryOrderBookResponseBidsInner) SetSize(v string) {
	o.Size = &v
}

func (o QueryOrderBookResponseBidsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderBookResponseBidsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryOrderBookResponseBidsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryOrderBookResponseBidsInner := _QueryOrderBookResponseBidsInner{}

	err = json.Unmarshal(data, &varQueryOrderBookResponseBidsInner)

	if err != nil {
		return err
	}

	*o = QueryOrderBookResponseBidsInner(varQueryOrderBookResponseBidsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryOrderBookResponseBidsInner struct {
	value *QueryOrderBookResponseBidsInner
	isSet bool
}

func (v NullableQueryOrderBookResponseBidsInner) Get() *QueryOrderBookResponseBidsInner {
	return v.value
}

func (v *NullableQueryOrderBookResponseBidsInner) Set(val *QueryOrderBookResponseBidsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOrderBookResponseBidsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOrderBookResponseBidsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOrderBookResponseBidsInner(val *QueryOrderBookResponseBidsInner) *NullableQueryOrderBookResponseBidsInner {
	return &NullableQueryOrderBookResponseBidsInner{value: val, isSet: true}
}

func (v NullableQueryOrderBookResponseBidsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOrderBookResponseBidsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
