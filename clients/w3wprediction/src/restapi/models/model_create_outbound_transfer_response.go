/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CreateOutboundTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateOutboundTransferResponse{}

// CreateOutboundTransferResponse struct for CreateOutboundTransferResponse
type CreateOutboundTransferResponse struct {
	TransferId           *string `json:"transferId,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOutboundTransferResponse CreateOutboundTransferResponse

// NewCreateOutboundTransferResponse instantiates a new CreateOutboundTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOutboundTransferResponse() *CreateOutboundTransferResponse {
	this := CreateOutboundTransferResponse{}
	return &this
}

// NewCreateOutboundTransferResponseWithDefaults instantiates a new CreateOutboundTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOutboundTransferResponseWithDefaults() *CreateOutboundTransferResponse {
	this := CreateOutboundTransferResponse{}
	return &this
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *CreateOutboundTransferResponse) GetTransferId() string {
	if o == nil || common.IsNil(o.TransferId) {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutboundTransferResponse) GetTransferIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *CreateOutboundTransferResponse) HasTransferId() bool {
	if o != nil && !common.IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *CreateOutboundTransferResponse) SetTransferId(v string) {
	o.TransferId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateOutboundTransferResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOutboundTransferResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateOutboundTransferResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateOutboundTransferResponse) SetStatus(v string) {
	o.Status = &v
}

func (o CreateOutboundTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOutboundTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransferId) {
		toSerialize["transferId"] = o.TransferId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOutboundTransferResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateOutboundTransferResponse := _CreateOutboundTransferResponse{}

	err = json.Unmarshal(data, &varCreateOutboundTransferResponse)

	if err != nil {
		return err
	}

	*o = CreateOutboundTransferResponse(varCreateOutboundTransferResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transferId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOutboundTransferResponse struct {
	value *CreateOutboundTransferResponse
	isSet bool
}

func (v NullableCreateOutboundTransferResponse) Get() *CreateOutboundTransferResponse {
	return v.value
}

func (v *NullableCreateOutboundTransferResponse) Set(val *CreateOutboundTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOutboundTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOutboundTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOutboundTransferResponse(val *CreateOutboundTransferResponse) *NullableCreateOutboundTransferResponse {
	return &NullableCreateOutboundTransferResponse{value: val, isSet: true}
}

func (v NullableCreateOutboundTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOutboundTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
