/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// CreateOutboundTransferSourceBizParameter the model 'CreateOutboundTransferSourceBizParameter'
type CreateOutboundTransferSourceBizParameter string

// List of createOutboundTransfer_sourceBiz_parameter
const (
	CreateOutboundTransferSourceBizParameterUserTransfer  CreateOutboundTransferSourceBizParameter = "USER_TRANSFER"
	CreateOutboundTransferSourceBizParameterPredictionBuy CreateOutboundTransferSourceBizParameter = "PREDICTION_BUY"
)

// All allowed values of CreateOutboundTransferSourceBizParameter enum
var AllowedCreateOutboundTransferSourceBizParameterEnumValues = []CreateOutboundTransferSourceBizParameter{
	"USER_TRANSFER",
	"PREDICTION_BUY",
}

func (v *CreateOutboundTransferSourceBizParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateOutboundTransferSourceBizParameter(value)
	for _, existing := range AllowedCreateOutboundTransferSourceBizParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateOutboundTransferSourceBizParameter", value)
}

// NewCreateOutboundTransferSourceBizParameterFromValue returns a pointer to a valid CreateOutboundTransferSourceBizParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateOutboundTransferSourceBizParameterFromValue(v string) (*CreateOutboundTransferSourceBizParameter, error) {
	ev := CreateOutboundTransferSourceBizParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateOutboundTransferSourceBizParameter: valid values are %v", v, AllowedCreateOutboundTransferSourceBizParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateOutboundTransferSourceBizParameter) IsValid() bool {
	for _, existing := range AllowedCreateOutboundTransferSourceBizParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to createOutboundTransfer_sourceBiz_parameter value
func (v CreateOutboundTransferSourceBizParameter) Ptr() *CreateOutboundTransferSourceBizParameter {
	return &v
}

type NullableCreateOutboundTransferSourceBizParameter struct {
	value *CreateOutboundTransferSourceBizParameter
	isSet bool
}

func (v NullableCreateOutboundTransferSourceBizParameter) Get() *CreateOutboundTransferSourceBizParameter {
	return v.value
}

func (v *NullableCreateOutboundTransferSourceBizParameter) Set(val *CreateOutboundTransferSourceBizParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOutboundTransferSourceBizParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOutboundTransferSourceBizParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOutboundTransferSourceBizParameter(val *CreateOutboundTransferSourceBizParameter) *NullableCreateOutboundTransferSourceBizParameter {
	return &NullableCreateOutboundTransferSourceBizParameter{value: val, isSet: true}
}

func (v NullableCreateOutboundTransferSourceBizParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOutboundTransferSourceBizParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
