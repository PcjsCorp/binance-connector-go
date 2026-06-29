/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListPredictionCategoriesResponseCategoriesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListPredictionCategoriesResponseCategoriesInner{}

// ListPredictionCategoriesResponseCategoriesInner struct for ListPredictionCategoriesResponseCategoriesInner
type ListPredictionCategoriesResponseCategoriesInner struct {
	Id                   *string                                                             `json:"id,omitempty"`
	Name                 *string                                                             `json:"name,omitempty"`
	Icon                 *string                                                             `json:"icon,omitempty"`
	Order                *int32                                                              `json:"order,omitempty"`
	Subcategories        []ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner `json:"subcategories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPredictionCategoriesResponseCategoriesInner ListPredictionCategoriesResponseCategoriesInner

// NewListPredictionCategoriesResponseCategoriesInner instantiates a new ListPredictionCategoriesResponseCategoriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPredictionCategoriesResponseCategoriesInner() *ListPredictionCategoriesResponseCategoriesInner {
	this := ListPredictionCategoriesResponseCategoriesInner{}
	return &this
}

// NewListPredictionCategoriesResponseCategoriesInnerWithDefaults instantiates a new ListPredictionCategoriesResponseCategoriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPredictionCategoriesResponseCategoriesInnerWithDefaults() *ListPredictionCategoriesResponseCategoriesInner {
	this := ListPredictionCategoriesResponseCategoriesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListPredictionCategoriesResponseCategoriesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListPredictionCategoriesResponseCategoriesInner) SetName(v string) {
	o.Name = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetIcon() string {
	if o == nil || common.IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetIconOk() (*string, bool) {
	if o == nil || common.IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) HasIcon() bool {
	if o != nil && !common.IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *ListPredictionCategoriesResponseCategoriesInner) SetIcon(v string) {
	o.Icon = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetOrder() int32 {
	if o == nil || common.IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetOrderOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) HasOrder() bool {
	if o != nil && !common.IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ListPredictionCategoriesResponseCategoriesInner) SetOrder(v int32) {
	o.Order = &v
}

// GetSubcategories returns the Subcategories field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetSubcategories() []ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner {
	if o == nil || common.IsNil(o.Subcategories) {
		var ret []ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner
		return ret
	}
	return o.Subcategories
}

// GetSubcategoriesOk returns a tuple with the Subcategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) GetSubcategoriesOk() ([]ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner, bool) {
	if o == nil || common.IsNil(o.Subcategories) {
		return nil, false
	}
	return o.Subcategories, true
}

// HasSubcategories returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponseCategoriesInner) HasSubcategories() bool {
	if o != nil && !common.IsNil(o.Subcategories) {
		return true
	}

	return false
}

// SetSubcategories gets a reference to the given []ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner and assigns it to the Subcategories field.
func (o *ListPredictionCategoriesResponseCategoriesInner) SetSubcategories(v []ListPredictionCategoriesResponseCategoriesInnerSubcategoriesInner) {
	o.Subcategories = v
}

func (o ListPredictionCategoriesResponseCategoriesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPredictionCategoriesResponseCategoriesInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Subcategories) {
		toSerialize["subcategories"] = o.Subcategories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPredictionCategoriesResponseCategoriesInner) UnmarshalJSON(data []byte) (err error) {
	varListPredictionCategoriesResponseCategoriesInner := _ListPredictionCategoriesResponseCategoriesInner{}

	err = json.Unmarshal(data, &varListPredictionCategoriesResponseCategoriesInner)

	if err != nil {
		return err
	}

	*o = ListPredictionCategoriesResponseCategoriesInner(varListPredictionCategoriesResponseCategoriesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "order")
		delete(additionalProperties, "subcategories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPredictionCategoriesResponseCategoriesInner struct {
	value *ListPredictionCategoriesResponseCategoriesInner
	isSet bool
}

func (v NullableListPredictionCategoriesResponseCategoriesInner) Get() *ListPredictionCategoriesResponseCategoriesInner {
	return v.value
}

func (v *NullableListPredictionCategoriesResponseCategoriesInner) Set(val *ListPredictionCategoriesResponseCategoriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListPredictionCategoriesResponseCategoriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListPredictionCategoriesResponseCategoriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPredictionCategoriesResponseCategoriesInner(val *ListPredictionCategoriesResponseCategoriesInner) *NullableListPredictionCategoriesResponseCategoriesInner {
	return &NullableListPredictionCategoriesResponseCategoriesInner{value: val, isSet: true}
}

func (v NullableListPredictionCategoriesResponseCategoriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPredictionCategoriesResponseCategoriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
