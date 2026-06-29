/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetQuotaStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetQuotaStatusResponse{}

// GetQuotaStatusResponse struct for GetQuotaStatusResponse
type GetQuotaStatusResponse struct {
	DailyLimit           *string `json:"dailyLimit,omitempty"`
	RemainingDailyLimit  *string `json:"remainingDailyLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetQuotaStatusResponse GetQuotaStatusResponse

// NewGetQuotaStatusResponse instantiates a new GetQuotaStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuotaStatusResponse() *GetQuotaStatusResponse {
	this := GetQuotaStatusResponse{}
	return &this
}

// NewGetQuotaStatusResponseWithDefaults instantiates a new GetQuotaStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuotaStatusResponseWithDefaults() *GetQuotaStatusResponse {
	this := GetQuotaStatusResponse{}
	return &this
}

// GetDailyLimit returns the DailyLimit field value if set, zero value otherwise.
func (o *GetQuotaStatusResponse) GetDailyLimit() string {
	if o == nil || common.IsNil(o.DailyLimit) {
		var ret string
		return ret
	}
	return *o.DailyLimit
}

// GetDailyLimitOk returns a tuple with the DailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuotaStatusResponse) GetDailyLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.DailyLimit) {
		return nil, false
	}
	return o.DailyLimit, true
}

// HasDailyLimit returns a boolean if a field has been set.
func (o *GetQuotaStatusResponse) HasDailyLimit() bool {
	if o != nil && !common.IsNil(o.DailyLimit) {
		return true
	}

	return false
}

// SetDailyLimit gets a reference to the given string and assigns it to the DailyLimit field.
func (o *GetQuotaStatusResponse) SetDailyLimit(v string) {
	o.DailyLimit = &v
}

// GetRemainingDailyLimit returns the RemainingDailyLimit field value if set, zero value otherwise.
func (o *GetQuotaStatusResponse) GetRemainingDailyLimit() string {
	if o == nil || common.IsNil(o.RemainingDailyLimit) {
		var ret string
		return ret
	}
	return *o.RemainingDailyLimit
}

// GetRemainingDailyLimitOk returns a tuple with the RemainingDailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuotaStatusResponse) GetRemainingDailyLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.RemainingDailyLimit) {
		return nil, false
	}
	return o.RemainingDailyLimit, true
}

// HasRemainingDailyLimit returns a boolean if a field has been set.
func (o *GetQuotaStatusResponse) HasRemainingDailyLimit() bool {
	if o != nil && !common.IsNil(o.RemainingDailyLimit) {
		return true
	}

	return false
}

// SetRemainingDailyLimit gets a reference to the given string and assigns it to the RemainingDailyLimit field.
func (o *GetQuotaStatusResponse) SetRemainingDailyLimit(v string) {
	o.RemainingDailyLimit = &v
}

func (o GetQuotaStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuotaStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DailyLimit) {
		toSerialize["dailyLimit"] = o.DailyLimit
	}
	if !common.IsNil(o.RemainingDailyLimit) {
		toSerialize["remainingDailyLimit"] = o.RemainingDailyLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetQuotaStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varGetQuotaStatusResponse := _GetQuotaStatusResponse{}

	err = json.Unmarshal(data, &varGetQuotaStatusResponse)

	if err != nil {
		return err
	}

	*o = GetQuotaStatusResponse(varGetQuotaStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dailyLimit")
		delete(additionalProperties, "remainingDailyLimit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetQuotaStatusResponse struct {
	value *GetQuotaStatusResponse
	isSet bool
}

func (v NullableGetQuotaStatusResponse) Get() *GetQuotaStatusResponse {
	return v.value
}

func (v *NullableGetQuotaStatusResponse) Set(val *GetQuotaStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuotaStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuotaStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuotaStatusResponse(val *GetQuotaStatusResponse) *NullableGetQuotaStatusResponse {
	return &NullableGetQuotaStatusResponse{value: val, isSet: true}
}

func (v NullableGetQuotaStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuotaStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
