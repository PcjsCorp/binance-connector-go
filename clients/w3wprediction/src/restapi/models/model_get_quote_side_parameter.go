/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetQuoteSideParameter the model 'GetQuoteSideParameter'
type GetQuoteSideParameter string

// List of getQuote_side_parameter
const (
	GetQuoteSideParameterBuy  GetQuoteSideParameter = "BUY"
	GetQuoteSideParameterSell GetQuoteSideParameter = "SELL"
)

// All allowed values of GetQuoteSideParameter enum
var AllowedGetQuoteSideParameterEnumValues = []GetQuoteSideParameter{
	"BUY",
	"SELL",
}

func (v *GetQuoteSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetQuoteSideParameter(value)
	for _, existing := range AllowedGetQuoteSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetQuoteSideParameter", value)
}

// NewGetQuoteSideParameterFromValue returns a pointer to a valid GetQuoteSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetQuoteSideParameterFromValue(v string) (*GetQuoteSideParameter, error) {
	ev := GetQuoteSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetQuoteSideParameter: valid values are %v", v, AllowedGetQuoteSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetQuoteSideParameter) IsValid() bool {
	for _, existing := range AllowedGetQuoteSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getQuote_side_parameter value
func (v GetQuoteSideParameter) Ptr() *GetQuoteSideParameter {
	return &v
}

type NullableGetQuoteSideParameter struct {
	value *GetQuoteSideParameter
	isSet bool
}

func (v NullableGetQuoteSideParameter) Get() *GetQuoteSideParameter {
	return v.value
}

func (v *NullableGetQuoteSideParameter) Set(val *GetQuoteSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuoteSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuoteSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuoteSideParameter(val *GetQuoteSideParameter) *NullableGetQuoteSideParameter {
	return &NullableGetQuoteSideParameter{value: val, isSet: true}
}

func (v NullableGetQuoteSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuoteSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
