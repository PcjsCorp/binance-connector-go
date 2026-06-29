/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ListPredictionMarketsOrderByParameter the model 'ListPredictionMarketsOrderByParameter'
type ListPredictionMarketsOrderByParameter string

// List of listPredictionMarkets_orderBy_parameter
const (
	ListPredictionMarketsOrderByParameterAsc  ListPredictionMarketsOrderByParameter = "ASC"
	ListPredictionMarketsOrderByParameterDesc ListPredictionMarketsOrderByParameter = "DESC"
)

// All allowed values of ListPredictionMarketsOrderByParameter enum
var AllowedListPredictionMarketsOrderByParameterEnumValues = []ListPredictionMarketsOrderByParameter{
	"ASC",
	"DESC",
}

func (v *ListPredictionMarketsOrderByParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListPredictionMarketsOrderByParameter(value)
	for _, existing := range AllowedListPredictionMarketsOrderByParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListPredictionMarketsOrderByParameter", value)
}

// NewListPredictionMarketsOrderByParameterFromValue returns a pointer to a valid ListPredictionMarketsOrderByParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListPredictionMarketsOrderByParameterFromValue(v string) (*ListPredictionMarketsOrderByParameter, error) {
	ev := ListPredictionMarketsOrderByParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListPredictionMarketsOrderByParameter: valid values are %v", v, AllowedListPredictionMarketsOrderByParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListPredictionMarketsOrderByParameter) IsValid() bool {
	for _, existing := range AllowedListPredictionMarketsOrderByParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to listPredictionMarkets_orderBy_parameter value
func (v ListPredictionMarketsOrderByParameter) Ptr() *ListPredictionMarketsOrderByParameter {
	return &v
}

type NullableListPredictionMarketsOrderByParameter struct {
	value *ListPredictionMarketsOrderByParameter
	isSet bool
}

func (v NullableListPredictionMarketsOrderByParameter) Get() *ListPredictionMarketsOrderByParameter {
	return v.value
}

func (v *NullableListPredictionMarketsOrderByParameter) Set(val *ListPredictionMarketsOrderByParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableListPredictionMarketsOrderByParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableListPredictionMarketsOrderByParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPredictionMarketsOrderByParameter(val *ListPredictionMarketsOrderByParameter) *NullableListPredictionMarketsOrderByParameter {
	return &NullableListPredictionMarketsOrderByParameter{value: val, isSet: true}
}

func (v NullableListPredictionMarketsOrderByParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPredictionMarketsOrderByParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
