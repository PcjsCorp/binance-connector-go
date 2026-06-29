/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetRedeemStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetRedeemStatusResponse{}

// GetRedeemStatusResponse struct for GetRedeemStatusResponse
type GetRedeemStatusResponse struct {
	TxHash               *string `json:"txHash,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRedeemStatusResponse GetRedeemStatusResponse

// NewGetRedeemStatusResponse instantiates a new GetRedeemStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRedeemStatusResponse() *GetRedeemStatusResponse {
	this := GetRedeemStatusResponse{}
	return &this
}

// NewGetRedeemStatusResponseWithDefaults instantiates a new GetRedeemStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRedeemStatusResponseWithDefaults() *GetRedeemStatusResponse {
	this := GetRedeemStatusResponse{}
	return &this
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *GetRedeemStatusResponse) GetTxHash() string {
	if o == nil || common.IsNil(o.TxHash) {
		var ret string
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRedeemStatusResponse) GetTxHashOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *GetRedeemStatusResponse) HasTxHash() bool {
	if o != nil && !common.IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given string and assigns it to the TxHash field.
func (o *GetRedeemStatusResponse) SetTxHash(v string) {
	o.TxHash = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetRedeemStatusResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRedeemStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetRedeemStatusResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetRedeemStatusResponse) SetStatus(v string) {
	o.Status = &v
}

func (o GetRedeemStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRedeemStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TxHash) {
		toSerialize["txHash"] = o.TxHash
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRedeemStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varGetRedeemStatusResponse := _GetRedeemStatusResponse{}

	err = json.Unmarshal(data, &varGetRedeemStatusResponse)

	if err != nil {
		return err
	}

	*o = GetRedeemStatusResponse(varGetRedeemStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "txHash")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRedeemStatusResponse struct {
	value *GetRedeemStatusResponse
	isSet bool
}

func (v NullableGetRedeemStatusResponse) Get() *GetRedeemStatusResponse {
	return v.value
}

func (v *NullableGetRedeemStatusResponse) Set(val *GetRedeemStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRedeemStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRedeemStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRedeemStatusResponse(val *GetRedeemStatusResponse) *NullableGetRedeemStatusResponse {
	return &NullableGetRedeemStatusResponse{value: val, isSet: true}
}

func (v NullableGetRedeemStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRedeemStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
