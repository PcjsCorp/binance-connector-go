/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetQuoteFundingSourceParameter the model 'GetQuoteFundingSourceParameter'
type GetQuoteFundingSourceParameter string

// List of getQuote_fundingSource_parameter
const (
	GetQuoteFundingSourceParameterMpc GetQuoteFundingSourceParameter = "MPC"
	GetQuoteFundingSourceParameterCex GetQuoteFundingSourceParameter = "CEX"
)

// All allowed values of GetQuoteFundingSourceParameter enum
var AllowedGetQuoteFundingSourceParameterEnumValues = []GetQuoteFundingSourceParameter{
	"MPC",
	"CEX",
}

func (v *GetQuoteFundingSourceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetQuoteFundingSourceParameter(value)
	for _, existing := range AllowedGetQuoteFundingSourceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetQuoteFundingSourceParameter", value)
}

// NewGetQuoteFundingSourceParameterFromValue returns a pointer to a valid GetQuoteFundingSourceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetQuoteFundingSourceParameterFromValue(v string) (*GetQuoteFundingSourceParameter, error) {
	ev := GetQuoteFundingSourceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetQuoteFundingSourceParameter: valid values are %v", v, AllowedGetQuoteFundingSourceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetQuoteFundingSourceParameter) IsValid() bool {
	for _, existing := range AllowedGetQuoteFundingSourceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getQuote_fundingSource_parameter value
func (v GetQuoteFundingSourceParameter) Ptr() *GetQuoteFundingSourceParameter {
	return &v
}

type NullableGetQuoteFundingSourceParameter struct {
	value *GetQuoteFundingSourceParameter
	isSet bool
}

func (v NullableGetQuoteFundingSourceParameter) Get() *GetQuoteFundingSourceParameter {
	return v.value
}

func (v *NullableGetQuoteFundingSourceParameter) Set(val *GetQuoteFundingSourceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuoteFundingSourceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuoteFundingSourceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuoteFundingSourceParameter(val *GetQuoteFundingSourceParameter) *NullableGetQuoteFundingSourceParameter {
	return &NullableGetQuoteFundingSourceParameter{value: val, isSet: true}
}

func (v NullableGetQuoteFundingSourceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuoteFundingSourceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
