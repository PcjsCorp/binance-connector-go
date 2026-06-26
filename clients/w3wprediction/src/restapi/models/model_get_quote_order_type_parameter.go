/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetQuoteOrderTypeParameter the model 'GetQuoteOrderTypeParameter'
type GetQuoteOrderTypeParameter string

// List of getQuote_orderType_parameter
const (
	GetQuoteOrderTypeParameterMarket GetQuoteOrderTypeParameter = "MARKET"
	GetQuoteOrderTypeParameterLimit  GetQuoteOrderTypeParameter = "LIMIT"
)

// All allowed values of GetQuoteOrderTypeParameter enum
var AllowedGetQuoteOrderTypeParameterEnumValues = []GetQuoteOrderTypeParameter{
	"MARKET",
	"LIMIT",
}

func (v *GetQuoteOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetQuoteOrderTypeParameter(value)
	for _, existing := range AllowedGetQuoteOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetQuoteOrderTypeParameter", value)
}

// NewGetQuoteOrderTypeParameterFromValue returns a pointer to a valid GetQuoteOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetQuoteOrderTypeParameterFromValue(v string) (*GetQuoteOrderTypeParameter, error) {
	ev := GetQuoteOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetQuoteOrderTypeParameter: valid values are %v", v, AllowedGetQuoteOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetQuoteOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetQuoteOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getQuote_orderType_parameter value
func (v GetQuoteOrderTypeParameter) Ptr() *GetQuoteOrderTypeParameter {
	return &v
}

type NullableGetQuoteOrderTypeParameter struct {
	value *GetQuoteOrderTypeParameter
	isSet bool
}

func (v NullableGetQuoteOrderTypeParameter) Get() *GetQuoteOrderTypeParameter {
	return v.value
}

func (v *NullableGetQuoteOrderTypeParameter) Set(val *GetQuoteOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuoteOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuoteOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuoteOrderTypeParameter(val *GetQuoteOrderTypeParameter) *NullableGetQuoteOrderTypeParameter {
	return &NullableGetQuoteOrderTypeParameter{value: val, isSet: true}
}

func (v NullableGetQuoteOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuoteOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
