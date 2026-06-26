/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryOrderBookResponseAsksInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOrderBookResponseAsksInner{}

// QueryOrderBookResponseAsksInner struct for QueryOrderBookResponseAsksInner
type QueryOrderBookResponseAsksInner struct {
	Price                *string `json:"price,omitempty"`
	Size                 *string `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryOrderBookResponseAsksInner QueryOrderBookResponseAsksInner

// NewQueryOrderBookResponseAsksInner instantiates a new QueryOrderBookResponseAsksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOrderBookResponseAsksInner() *QueryOrderBookResponseAsksInner {
	this := QueryOrderBookResponseAsksInner{}
	return &this
}

// NewQueryOrderBookResponseAsksInnerWithDefaults instantiates a new QueryOrderBookResponseAsksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOrderBookResponseAsksInnerWithDefaults() *QueryOrderBookResponseAsksInner {
	this := QueryOrderBookResponseAsksInner{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryOrderBookResponseAsksInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponseAsksInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryOrderBookResponseAsksInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryOrderBookResponseAsksInner) SetPrice(v string) {
	o.Price = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *QueryOrderBookResponseAsksInner) GetSize() string {
	if o == nil || common.IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponseAsksInner) GetSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *QueryOrderBookResponseAsksInner) HasSize() bool {
	if o != nil && !common.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *QueryOrderBookResponseAsksInner) SetSize(v string) {
	o.Size = &v
}

func (o QueryOrderBookResponseAsksInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderBookResponseAsksInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryOrderBookResponseAsksInner) UnmarshalJSON(data []byte) (err error) {
	varQueryOrderBookResponseAsksInner := _QueryOrderBookResponseAsksInner{}

	err = json.Unmarshal(data, &varQueryOrderBookResponseAsksInner)

	if err != nil {
		return err
	}

	*o = QueryOrderBookResponseAsksInner(varQueryOrderBookResponseAsksInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryOrderBookResponseAsksInner struct {
	value *QueryOrderBookResponseAsksInner
	isSet bool
}

func (v NullableQueryOrderBookResponseAsksInner) Get() *QueryOrderBookResponseAsksInner {
	return v.value
}

func (v *NullableQueryOrderBookResponseAsksInner) Set(val *QueryOrderBookResponseAsksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOrderBookResponseAsksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOrderBookResponseAsksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOrderBookResponseAsksInner(val *QueryOrderBookResponseAsksInner) *NullableQueryOrderBookResponseAsksInner {
	return &NullableQueryOrderBookResponseAsksInner{value: val, isSet: true}
}

func (v NullableQueryOrderBookResponseAsksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOrderBookResponseAsksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
