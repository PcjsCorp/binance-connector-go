/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// PlaceOrderAccountTypeParameter the model 'PlaceOrderAccountTypeParameter'
type PlaceOrderAccountTypeParameter string

// List of placeOrder_accountType_parameter
const (
	PlaceOrderAccountTypeParameterSpot    PlaceOrderAccountTypeParameter = "SPOT"
	PlaceOrderAccountTypeParameterFunding PlaceOrderAccountTypeParameter = "FUNDING"
)

// All allowed values of PlaceOrderAccountTypeParameter enum
var AllowedPlaceOrderAccountTypeParameterEnumValues = []PlaceOrderAccountTypeParameter{
	"SPOT",
	"FUNDING",
}

func (v *PlaceOrderAccountTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaceOrderAccountTypeParameter(value)
	for _, existing := range AllowedPlaceOrderAccountTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaceOrderAccountTypeParameter", value)
}

// NewPlaceOrderAccountTypeParameterFromValue returns a pointer to a valid PlaceOrderAccountTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaceOrderAccountTypeParameterFromValue(v string) (*PlaceOrderAccountTypeParameter, error) {
	ev := PlaceOrderAccountTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaceOrderAccountTypeParameter: valid values are %v", v, AllowedPlaceOrderAccountTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaceOrderAccountTypeParameter) IsValid() bool {
	for _, existing := range AllowedPlaceOrderAccountTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to placeOrder_accountType_parameter value
func (v PlaceOrderAccountTypeParameter) Ptr() *PlaceOrderAccountTypeParameter {
	return &v
}

type NullablePlaceOrderAccountTypeParameter struct {
	value *PlaceOrderAccountTypeParameter
	isSet bool
}

func (v NullablePlaceOrderAccountTypeParameter) Get() *PlaceOrderAccountTypeParameter {
	return v.value
}

func (v *NullablePlaceOrderAccountTypeParameter) Set(val *PlaceOrderAccountTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceOrderAccountTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceOrderAccountTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceOrderAccountTypeParameter(val *PlaceOrderAccountTypeParameter) *NullablePlaceOrderAccountTypeParameter {
	return &NullablePlaceOrderAccountTypeParameter{value: val, isSet: true}
}

func (v NullablePlaceOrderAccountTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceOrderAccountTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
