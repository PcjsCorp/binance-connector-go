/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryTransferListResponseTransfersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryTransferListResponseTransfersInner{}

// QueryTransferListResponseTransfersInner struct for QueryTransferListResponseTransfersInner
type QueryTransferListResponseTransfersInner struct {
	TransferId           *string `json:"transferId,omitempty"`
	Direction            *string `json:"direction,omitempty"`
	Status               *string `json:"status,omitempty"`
	WalletAddress        *string `json:"walletAddress,omitempty"`
	FromToken            *string `json:"fromToken,omitempty"`
	FromTokenAmount      *string `json:"fromTokenAmount,omitempty"`
	ToToken              *string `json:"toToken,omitempty"`
	ToTokenAmount        *string `json:"toTokenAmount,omitempty"`
	ErrorCode            *string `json:"errorCode,omitempty"`
	ErrorMessage         *string `json:"errorMessage,omitempty"`
	CreateTime           *string `json:"createTime,omitempty"`
	UpdateTime           *string `json:"updateTime,omitempty"`
	CompleteAt           *string `json:"completeAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryTransferListResponseTransfersInner QueryTransferListResponseTransfersInner

// NewQueryTransferListResponseTransfersInner instantiates a new QueryTransferListResponseTransfersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryTransferListResponseTransfersInner() *QueryTransferListResponseTransfersInner {
	this := QueryTransferListResponseTransfersInner{}
	return &this
}

// NewQueryTransferListResponseTransfersInnerWithDefaults instantiates a new QueryTransferListResponseTransfersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTransferListResponseTransfersInnerWithDefaults() *QueryTransferListResponseTransfersInner {
	this := QueryTransferListResponseTransfersInner{}
	return &this
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetTransferId() string {
	if o == nil || common.IsNil(o.TransferId) {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetTransferIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferId) {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasTransferId() bool {
	if o != nil && !common.IsNil(o.TransferId) {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *QueryTransferListResponseTransfersInner) SetTransferId(v string) {
	o.TransferId = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetDirection() string {
	if o == nil || common.IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetDirectionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasDirection() bool {
	if o != nil && !common.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *QueryTransferListResponseTransfersInner) SetDirection(v string) {
	o.Direction = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryTransferListResponseTransfersInner) SetStatus(v string) {
	o.Status = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetWalletAddress() string {
	if o == nil || common.IsNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetWalletAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletAddress) {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasWalletAddress() bool {
	if o != nil && !common.IsNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *QueryTransferListResponseTransfersInner) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetFromToken returns the FromToken field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetFromToken() string {
	if o == nil || common.IsNil(o.FromToken) {
		var ret string
		return ret
	}
	return *o.FromToken
}

// GetFromTokenOk returns a tuple with the FromToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetFromTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromToken) {
		return nil, false
	}
	return o.FromToken, true
}

// HasFromToken returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasFromToken() bool {
	if o != nil && !common.IsNil(o.FromToken) {
		return true
	}

	return false
}

// SetFromToken gets a reference to the given string and assigns it to the FromToken field.
func (o *QueryTransferListResponseTransfersInner) SetFromToken(v string) {
	o.FromToken = &v
}

// GetFromTokenAmount returns the FromTokenAmount field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetFromTokenAmount() string {
	if o == nil || common.IsNil(o.FromTokenAmount) {
		var ret string
		return ret
	}
	return *o.FromTokenAmount
}

// GetFromTokenAmountOk returns a tuple with the FromTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetFromTokenAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromTokenAmount) {
		return nil, false
	}
	return o.FromTokenAmount, true
}

// HasFromTokenAmount returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasFromTokenAmount() bool {
	if o != nil && !common.IsNil(o.FromTokenAmount) {
		return true
	}

	return false
}

// SetFromTokenAmount gets a reference to the given string and assigns it to the FromTokenAmount field.
func (o *QueryTransferListResponseTransfersInner) SetFromTokenAmount(v string) {
	o.FromTokenAmount = &v
}

// GetToToken returns the ToToken field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetToToken() string {
	if o == nil || common.IsNil(o.ToToken) {
		var ret string
		return ret
	}
	return *o.ToToken
}

// GetToTokenOk returns a tuple with the ToToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetToTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToToken) {
		return nil, false
	}
	return o.ToToken, true
}

// HasToToken returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasToToken() bool {
	if o != nil && !common.IsNil(o.ToToken) {
		return true
	}

	return false
}

// SetToToken gets a reference to the given string and assigns it to the ToToken field.
func (o *QueryTransferListResponseTransfersInner) SetToToken(v string) {
	o.ToToken = &v
}

// GetToTokenAmount returns the ToTokenAmount field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetToTokenAmount() string {
	if o == nil || common.IsNil(o.ToTokenAmount) {
		var ret string
		return ret
	}
	return *o.ToTokenAmount
}

// GetToTokenAmountOk returns a tuple with the ToTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetToTokenAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.ToTokenAmount) {
		return nil, false
	}
	return o.ToTokenAmount, true
}

// HasToTokenAmount returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasToTokenAmount() bool {
	if o != nil && !common.IsNil(o.ToTokenAmount) {
		return true
	}

	return false
}

// SetToTokenAmount gets a reference to the given string and assigns it to the ToTokenAmount field.
func (o *QueryTransferListResponseTransfersInner) SetToTokenAmount(v string) {
	o.ToTokenAmount = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryTransferListResponseTransfersInner) GetErrorCode() string {
	if o == nil || common.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryTransferListResponseTransfersInner) GetErrorCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasErrorCode() bool {
	if o != nil && !common.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *QueryTransferListResponseTransfersInner) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryTransferListResponseTransfersInner) GetErrorMessage() string {
	if o == nil || common.IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryTransferListResponseTransfersInner) GetErrorMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasErrorMessage() bool {
	if o != nil && !common.IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *QueryTransferListResponseTransfersInner) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetCreateTime() string {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetCreateTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *QueryTransferListResponseTransfersInner) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetUpdateTime() string {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetUpdateTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *QueryTransferListResponseTransfersInner) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetCompleteAt returns the CompleteAt field value if set, zero value otherwise.
func (o *QueryTransferListResponseTransfersInner) GetCompleteAt() string {
	if o == nil || common.IsNil(o.CompleteAt) {
		var ret string
		return ret
	}
	return *o.CompleteAt
}

// GetCompleteAtOk returns a tuple with the CompleteAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTransferListResponseTransfersInner) GetCompleteAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.CompleteAt) {
		return nil, false
	}
	return o.CompleteAt, true
}

// HasCompleteAt returns a boolean if a field has been set.
func (o *QueryTransferListResponseTransfersInner) HasCompleteAt() bool {
	if o != nil && !common.IsNil(o.CompleteAt) {
		return true
	}

	return false
}

// SetCompleteAt gets a reference to the given string and assigns it to the CompleteAt field.
func (o *QueryTransferListResponseTransfersInner) SetCompleteAt(v string) {
	o.CompleteAt = &v
}

func (o QueryTransferListResponseTransfersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryTransferListResponseTransfersInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.WalletAddress) {
		toSerialize["walletAddress"] = o.WalletAddress
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
	if !common.IsNil(o.CompleteAt) {
		toSerialize["completeAt"] = o.CompleteAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryTransferListResponseTransfersInner) UnmarshalJSON(data []byte) (err error) {
	varQueryTransferListResponseTransfersInner := _QueryTransferListResponseTransfersInner{}

	err = json.Unmarshal(data, &varQueryTransferListResponseTransfersInner)

	if err != nil {
		return err
	}

	*o = QueryTransferListResponseTransfersInner(varQueryTransferListResponseTransfersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transferId")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "status")
		delete(additionalProperties, "walletAddress")
		delete(additionalProperties, "fromToken")
		delete(additionalProperties, "fromTokenAmount")
		delete(additionalProperties, "toToken")
		delete(additionalProperties, "toTokenAmount")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "completeAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryTransferListResponseTransfersInner struct {
	value *QueryTransferListResponseTransfersInner
	isSet bool
}

func (v NullableQueryTransferListResponseTransfersInner) Get() *QueryTransferListResponseTransfersInner {
	return v.value
}

func (v *NullableQueryTransferListResponseTransfersInner) Set(val *QueryTransferListResponseTransfersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryTransferListResponseTransfersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryTransferListResponseTransfersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryTransferListResponseTransfersInner(val *QueryTransferListResponseTransfersInner) *NullableQueryTransferListResponseTransfersInner {
	return &NullableQueryTransferListResponseTransfersInner{value: val, isSet: true}
}

func (v NullableQueryTransferListResponseTransfersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryTransferListResponseTransfersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
