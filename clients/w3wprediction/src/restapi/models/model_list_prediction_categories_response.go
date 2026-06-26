/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListPredictionCategoriesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListPredictionCategoriesResponse{}

// ListPredictionCategoriesResponse struct for ListPredictionCategoriesResponse
type ListPredictionCategoriesResponse struct {
	Categories           []ListPredictionCategoriesResponseCategoriesInner `json:"categories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPredictionCategoriesResponse ListPredictionCategoriesResponse

// NewListPredictionCategoriesResponse instantiates a new ListPredictionCategoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPredictionCategoriesResponse() *ListPredictionCategoriesResponse {
	this := ListPredictionCategoriesResponse{}
	return &this
}

// NewListPredictionCategoriesResponseWithDefaults instantiates a new ListPredictionCategoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPredictionCategoriesResponseWithDefaults() *ListPredictionCategoriesResponse {
	this := ListPredictionCategoriesResponse{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *ListPredictionCategoriesResponse) GetCategories() []ListPredictionCategoriesResponseCategoriesInner {
	if o == nil || common.IsNil(o.Categories) {
		var ret []ListPredictionCategoriesResponseCategoriesInner
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPredictionCategoriesResponse) GetCategoriesOk() ([]ListPredictionCategoriesResponseCategoriesInner, bool) {
	if o == nil || common.IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *ListPredictionCategoriesResponse) HasCategories() bool {
	if o != nil && !common.IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []ListPredictionCategoriesResponseCategoriesInner and assigns it to the Categories field.
func (o *ListPredictionCategoriesResponse) SetCategories(v []ListPredictionCategoriesResponseCategoriesInner) {
	o.Categories = v
}

func (o ListPredictionCategoriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPredictionCategoriesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPredictionCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	varListPredictionCategoriesResponse := _ListPredictionCategoriesResponse{}

	err = json.Unmarshal(data, &varListPredictionCategoriesResponse)

	if err != nil {
		return err
	}

	*o = ListPredictionCategoriesResponse(varListPredictionCategoriesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "categories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPredictionCategoriesResponse struct {
	value *ListPredictionCategoriesResponse
	isSet bool
}

func (v NullableListPredictionCategoriesResponse) Get() *ListPredictionCategoriesResponse {
	return v.value
}

func (v *NullableListPredictionCategoriesResponse) Set(val *ListPredictionCategoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPredictionCategoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPredictionCategoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPredictionCategoriesResponse(val *ListPredictionCategoriesResponse) *NullableListPredictionCategoriesResponse {
	return &NullableListPredictionCategoriesResponse{value: val, isSet: true}
}

func (v NullableListPredictionCategoriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPredictionCategoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
