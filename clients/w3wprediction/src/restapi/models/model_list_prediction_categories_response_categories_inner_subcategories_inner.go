/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner{}

// ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner struct for ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner
type ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner struct {
	Id                   *string `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Icon                 *string `json:"icon,omitempty"`
	Order                *int32  `json:"order,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner

// NewListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner instantiates a new ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner() *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner {
	this := ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner{}
	return &this
}

// NewListPredictionCategoriesResponseCategoriesInnerSubcategoriesInnerWithDefaults instantiates a new ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPredictionCategoriesResponseCategoriesInnerSubcategoriesInnerWithDefaults() *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner {
	this := ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) SetName(v string) {
	o.Name = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) GetIcon() string {
	if o == nil || common.IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) GetIconOk() (*string, bool) {
	if o == nil || common.IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) HasIcon() bool {
	if o != nil && !common.IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) SetIcon(v string) {
	o.Icon = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) GetOrder() int32 {
	if o == nil || common.IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) GetOrderOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) HasOrder() bool {
	if o != nil && !common.IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) SetOrder(v int32) {
	o.Order = &v
}

func (o ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !common.IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) UnmarshalJSON(data []byte) (err error) {
	varListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner := _ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner{}

	err = json.Unmarshal(data, &varListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner)

	if err != nil {
		return err
	}

	*o = ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner(varListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "order")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner struct {
	value *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner
	isSet bool
}

func (v NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) Get() *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner {
	return v.value
}

func (v *NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) Set(val *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner(val *ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) *NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner {
	return &NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner{value: val, isSet: true}
}

func (v NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
