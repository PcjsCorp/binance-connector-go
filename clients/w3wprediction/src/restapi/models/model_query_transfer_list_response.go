/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryTransferListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryTransferListResponse{}

// QueryTransferListResponse struct for QueryTransferListResponse
type QueryTransferListResponse struct {
	Transfers            []QueryTransferListResponseTransfersInner `json:"transfers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryTransferListResponse QueryTransferListResponse

// NewQueryTransferListResponse instantiates a new QueryTransferListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryTransferListResponse() *QueryTransferListResponse {
	this := QueryTransferListResponse{}
	return &this
}

// NewQueryTransferListResponseWithDefaults instantiates a new QueryTransferListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTransferListResponseWithDefaults() *QueryTransferListResponse {
	this := QueryTransferListResponse{}
	return &this
}

// GetTransfers returns the Transfers field value if set, zero value otherwise.
func (o *QueryTransferListResponse) GetTransfers() []QueryTransferListResponseTransfersInner {
	if o == nil || common.IsNil(o.Transfers) {
		var ret []QueryTransferListResponseTransfersInner
		return ret
	}
	return o.Transfers
}

// GetTransfersOk returns a tuple with the Transfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponse) GetTransfersOk() ([]QueryTransferListResponseTransfersInner, bool) {
	if o == nil || common.IsNil(o.Transfers) {
		return nil, false
	}
	return o.Transfers, true
}

// HasTransfers returns a boolean if a field has been set.
func (o *QueryTransferListResponse) HasTransfers() bool {
	if o != nil && !common.IsNil(o.Transfers) {
		return true
	}

	return false
}

// SetTransfers gets a reference to the given []QueryTransferListResponseTransfersInner and assigns it to the Transfers field.
func (o *QueryTransferListResponse) SetTransfers(v []QueryTransferListResponseTransfersInner) {
	o.Transfers = v
}

func (o QueryTransferListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryTransferListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Transfers) {
		toSerialize["transfers"] = o.Transfers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryTransferListResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryTransferListResponse := _QueryTransferListResponse{}

	err = json.Unmarshal(data, &varQueryTransferListResponse)

	if err != nil {
		return err
	}

	*o = QueryTransferListResponse(varQueryTransferListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transfers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryTransferListResponse struct {
	value *QueryTransferListResponse
	isSet bool
}

func (v NullableQueryTransferListResponse) Get() *QueryTransferListResponse {
	return v.value
}

func (v *NullableQueryTransferListResponse) Set(val *QueryTransferListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryTransferListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryTransferListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryTransferListResponse(val *QueryTransferListResponse) *NullableQueryTransferListResponse {
	return &NullableQueryTransferListResponse{value: val, isSet: true}
}

func (v NullableQueryTransferListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryTransferListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
