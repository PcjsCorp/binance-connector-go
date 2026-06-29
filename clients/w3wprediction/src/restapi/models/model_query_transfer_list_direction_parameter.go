/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryTransferListDirectionParameter the model 'QueryTransferListDirectionParameter'
type QueryTransferListDirectionParameter string

// List of queryTransferList_direction_parameter
const (
	QueryTransferListDirectionParameterInbound  QueryTransferListDirectionParameter = "INBOUND"
	QueryTransferListDirectionParameterOutbound QueryTransferListDirectionParameter = "OUTBOUND"
)

// All allowed values of QueryTransferListDirectionParameter enum
var AllowedQueryTransferListDirectionParameterEnumValues = []QueryTransferListDirectionParameter{
	"INBOUND",
	"OUTBOUND",
}

func (v *QueryTransferListDirectionParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryTransferListDirectionParameter(value)
	for _, existing := range AllowedQueryTransferListDirectionParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryTransferListDirectionParameter", value)
}

// NewQueryTransferListDirectionParameterFromValue returns a pointer to a valid QueryTransferListDirectionParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryTransferListDirectionParameterFromValue(v string) (*QueryTransferListDirectionParameter, error) {
	ev := QueryTransferListDirectionParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryTransferListDirectionParameter: valid values are %v", v, AllowedQueryTransferListDirectionParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryTransferListDirectionParameter) IsValid() bool {
	for _, existing := range AllowedQueryTransferListDirectionParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryTransferList_direction_parameter value
func (v QueryTransferListDirectionParameter) Ptr() *QueryTransferListDirectionParameter {
	return &v
}

type NullableQueryTransferListDirectionParameter struct {
	value *QueryTransferListDirectionParameter
	isSet bool
}

func (v NullableQueryTransferListDirectionParameter) Get() *QueryTransferListDirectionParameter {
	return v.value
}

func (v *NullableQueryTransferListDirectionParameter) Set(val *QueryTransferListDirectionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryTransferListDirectionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryTransferListDirectionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryTransferListDirectionParameter(val *QueryTransferListDirectionParameter) *NullableQueryTransferListDirectionParameter {
	return &NullableQueryTransferListDirectionParameter{value: val, isSet: true}
}

func (v NullableQueryTransferListDirectionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryTransferListDirectionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
