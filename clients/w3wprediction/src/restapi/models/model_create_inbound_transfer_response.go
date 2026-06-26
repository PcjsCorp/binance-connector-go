/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CreateInboundTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateInboundTransferResponse{}

// CreateInboundTransferResponse struct for CreateInboundTransferResponse
type CreateInboundTransferResponse struct {
	TransferId           *string `json:"transferId,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateInboundTransferResponse CreateInboundTransferResponse

// NewCreateInboundTransferResponse instantiates a new CreateInboundTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInboundTransferResponse() *CreateInboundTransferResponse {
	this := CreateInboundTransferResponse{}
	return &this
}

// NewCreateInboundTransferResponseWithDefaults instantiates a new CreateInboundTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInboundTransferResponseWithDefaults() *CreateInboundTransferResponse {
	this := CreateInboundTransferResponse{}
	return &this
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *CreateInboundTransferResponse) GetTransferId() string {
	if o == nil || common.IsNil(o.TransferId) {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundTransferResponse) GetTransferIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *CreateInboundTransferResponse) HasTransferId() bool {
	if o != nil && !common.IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *CreateInboundTransferResponse) SetTransferId(v string) {
	o.TransferId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateInboundTransferResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundTransferResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateInboundTransferResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateInboundTransferResponse) SetStatus(v string) {
	o.Status = &v
}

func (o CreateInboundTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInboundTransferResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CreateInboundTransferResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateInboundTransferResponse := _CreateInboundTransferResponse{}

	err = json.Unmarshal(data, &varCreateInboundTransferResponse)

	if err != nil {
		return err
	}

	*o = CreateInboundTransferResponse(varCreateInboundTransferResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transferId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateInboundTransferResponse struct {
	value *CreateInboundTransferResponse
	isSet bool
}

func (v NullableCreateInboundTransferResponse) Get() *CreateInboundTransferResponse {
	return v.value
}

func (v *NullableCreateInboundTransferResponse) Set(val *CreateInboundTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInboundTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInboundTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInboundTransferResponse(val *CreateInboundTransferResponse) *NullableCreateInboundTransferResponse {
	return &NullableCreateInboundTransferResponse{value: val, isSet: true}
}

func (v NullableCreateInboundTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInboundTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
