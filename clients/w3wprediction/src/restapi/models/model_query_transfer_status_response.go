/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryTransferStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryTransferStatusResponse{}

// QueryTransferStatusResponse struct for QueryTransferStatusResponse
type QueryTransferStatusResponse struct {
	TransferId           *string `json:"transferId,omitempty"`
	Direction            *string `json:"direction,omitempty"`
	Status               *string `json:"status,omitempty"`
	FromToken            *string `json:"fromToken,omitempty"`
	FromTokenAmount      *string `json:"fromTokenAmount,omitempty"`
	ToToken              *string `json:"toToken,omitempty"`
	ToTokenAmount        *string `json:"toTokenAmount,omitempty"`
	ErrorCode            *string `json:"errorCode,omitempty"`
	ErrorMessage         *string `json:"errorMessage,omitempty"`
	CreateTime           *string `json:"createTime,omitempty"`
	UpdateTime           *string `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryTransferStatusResponse QueryTransferStatusResponse

// NewQueryTransferStatusResponse instantiates a new QueryTransferStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryTransferStatusResponse() *QueryTransferStatusResponse {
	this := QueryTransferStatusResponse{}
	return &this
}

// NewQueryTransferStatusResponseWithDefaults instantiates a new QueryTransferStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTransferStatusResponseWithDefaults() *QueryTransferStatusResponse {
	this := QueryTransferStatusResponse{}
	return &this
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetTransferId() string {
	if o == nil || common.IsNil(o.TransferId) {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetTransferIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasTransferId() bool {
	if o != nil && !common.IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *QueryTransferStatusResponse) SetTransferId(v string) {
	o.TransferId = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetDirection() string {
	if o == nil || common.IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetDirectionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasDirection() bool {
	if o != nil && !common.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *QueryTransferStatusResponse) SetDirection(v string) {
	o.Direction = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryTransferStatusResponse) SetStatus(v string) {
	o.Status = &v
}

// GetFromToken returns the FromToken field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetFromToken() string {
	if o == nil || common.IsNil(o.FromToken) {
		var ret string
		return ret
	}
	return *o.FromToken
}

// GetFromTokenOk returns a tuple with the FromToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetFromTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromToken) {
		return nil, false
	}
	return o.FromToken, true
}

// HasFromToken returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasFromToken() bool {
	if o != nil && !common.IsNil(o.FromToken) {
		return true
	}

	return false
}

// SetFromToken gets a reference to the given string and assigns it to the FromToken field.
func (o *QueryTransferStatusResponse) SetFromToken(v string) {
	o.FromToken = &v
}

// GetFromTokenAmount returns the FromTokenAmount field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetFromTokenAmount() string {
	if o == nil || common.IsNil(o.FromTokenAmount) {
		var ret string
		return ret
	}
	return *o.FromTokenAmount
}

// GetFromTokenAmountOk returns a tuple with the FromTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetFromTokenAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromTokenAmount) {
		return nil, false
	}
	return o.FromTokenAmount, true
}

// HasFromTokenAmount returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasFromTokenAmount() bool {
	if o != nil && !common.IsNil(o.FromTokenAmount) {
		return true
	}

	return false
}

// SetFromTokenAmount gets a reference to the given string and assigns it to the FromTokenAmount field.
func (o *QueryTransferStatusResponse) SetFromTokenAmount(v string) {
	o.FromTokenAmount = &v
}

// GetToToken returns the ToToken field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetToToken() string {
	if o == nil || common.IsNil(o.ToToken) {
		var ret string
		return ret
	}
	return *o.ToToken
}

// GetToTokenOk returns a tuple with the ToToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetToTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToToken) {
		return nil, false
	}
	return o.ToToken, true
}

// HasToToken returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasToToken() bool {
	if o != nil && !common.IsNil(o.ToToken) {
		return true
	}

	return false
}

// SetToToken gets a reference to the given string and assigns it to the ToToken field.
func (o *QueryTransferStatusResponse) SetToToken(v string) {
	o.ToToken = &v
}

// GetToTokenAmount returns the ToTokenAmount field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetToTokenAmount() string {
	if o == nil || common.IsNil(o.ToTokenAmount) {
		var ret string
		return ret
	}
	return *o.ToTokenAmount
}

// GetToTokenAmountOk returns a tuple with the ToTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetToTokenAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToTokenAmount) {
		return nil, false
	}
	return o.ToTokenAmount, true
}

// HasToTokenAmount returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasToTokenAmount() bool {
	if o != nil && !common.IsNil(o.ToTokenAmount) {
		return true
	}

	return false
}

// SetToTokenAmount gets a reference to the given string and assigns it to the ToTokenAmount field.
func (o *QueryTransferStatusResponse) SetToTokenAmount(v string) {
	o.ToTokenAmount = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryTransferStatusResponse) GetErrorCode() string {
	if o == nil || common.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryTransferStatusResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasErrorCode() bool {
	if o != nil && !common.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *QueryTransferStatusResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryTransferStatusResponse) GetErrorMessage() string {
	if o == nil || common.IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryTransferStatusResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasErrorMessage() bool {
	if o != nil && !common.IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *QueryTransferStatusResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetCreateTime() string {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetCreateTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *QueryTransferStatusResponse) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryTransferStatusResponse) GetUpdateTime() string {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferStatusResponse) GetUpdateTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryTransferStatusResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *QueryTransferStatusResponse) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

func (o QueryTransferStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryTransferStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransferId) {
		toSerialize["transferId"] = o.TransferId
	}
	if !common.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.FromToken) {
		toSerialize["fromToken"] = o.FromToken
	}
	if !common.IsNil(o.FromTokenAmount) {
		toSerialize["fromTokenAmount"] = o.FromTokenAmount
	}
	if !common.IsNil(o.ToToken) {
		toSerialize["toToken"] = o.ToToken
	}
	if !common.IsNil(o.ToTokenAmount) {
		toSerialize["toTokenAmount"] = o.ToTokenAmount
	}
	if !common.IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !common.IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryTransferStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryTransferStatusResponse := _QueryTransferStatusResponse{}

	err = json.Unmarshal(data, &varQueryTransferStatusResponse)

	if err != nil {
		return err
	}

	*o = QueryTransferStatusResponse(varQueryTransferStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transferId")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "status")
		delete(additionalProperties, "fromToken")
		delete(additionalProperties, "fromTokenAmount")
		delete(additionalProperties, "toToken")
		delete(additionalProperties, "toTokenAmount")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryTransferStatusResponse struct {
	value *QueryTransferStatusResponse
	isSet bool
}

func (v NullableQueryTransferStatusResponse) Get() *QueryTransferStatusResponse {
	return v.value
}

func (v *NullableQueryTransferStatusResponse) Set(val *QueryTransferStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryTransferStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryTransferStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryTransferStatusResponse(val *QueryTransferStatusResponse) *NullableQueryTransferStatusResponse {
	return &NullableQueryTransferStatusResponse{value: val, isSet: true}
}

func (v NullableQueryTransferStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryTransferStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
