/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ListPredictionMarketsSortByParameter the model 'ListPredictionMarketsSortByParameter'
type ListPredictionMarketsSortByParameter string

// List of listPredictionMarkets_sortBy_parameter
const (
	ListPredictionMarketsSortByParameterRecommended  ListPredictionMarketsSortByParameter = "RECOMMENDED"
	ListPredictionMarketsSortByParameterVolume       ListPredictionMarketsSortByParameter = "VOLUME"
	ListPredictionMarketsSortByParameterParticipants ListPredictionMarketsSortByParameter = "PARTICIPANTS"
	ListPredictionMarketsSortByParameterCreatedTime  ListPredictionMarketsSortByParameter = "CREATED_TIME"
	ListPredictionMarketsSortByParameterEndDate      ListPredictionMarketsSortByParameter = "END_DATE"
)

// All allowed values of ListPredictionMarketsSortByParameter enum
var AllowedListPredictionMarketsSortByParameterEnumValues = []ListPredictionMarketsSortByParameter{
	"RECOMMENDED",
	"VOLUME",
	"PARTICIPANTS",
	"CREATED_TIME",
	"END_DATE",
}

func (v *ListPredictionMarketsSortByParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListPredictionMarketsSortByParameter(value)
	for _, existing := range AllowedListPredictionMarketsSortByParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListPredictionMarketsSortByParameter", value)
}

// NewListPredictionMarketsSortByParameterFromValue returns a pointer to a valid ListPredictionMarketsSortByParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListPredictionMarketsSortByParameterFromValue(v string) (*ListPredictionMarketsSortByParameter, error) {
	ev := ListPredictionMarketsSortByParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListPredictionMarketsSortByParameter: valid values are %v", v, AllowedListPredictionMarketsSortByParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListPredictionMarketsSortByParameter) IsValid() bool {
	for _, existing := range AllowedListPredictionMarketsSortByParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to listPredictionMarkets_sortBy_parameter value
func (v ListPredictionMarketsSortByParameter) Ptr() *ListPredictionMarketsSortByParameter {
	return &v
}

type NullableListPredictionMarketsSortByParameter struct {
	value *ListPredictionMarketsSortByParameter
	isSet bool
}

func (v NullableListPredictionMarketsSortByParameter) Get() *ListPredictionMarketsSortByParameter {
	return v.value
}

func (v *NullableListPredictionMarketsSortByParameter) Set(val *ListPredictionMarketsSortByParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableListPredictionMarketsSortByParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableListPredictionMarketsSortByParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPredictionMarketsSortByParameter(val *ListPredictionMarketsSortByParameter) *NullableListPredictionMarketsSortByParameter {
	return &NullableListPredictionMarketsSortByParameter{value: val, isSet: true}
}

func (v NullableListPredictionMarketsSortByParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPredictionMarketsSortByParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
