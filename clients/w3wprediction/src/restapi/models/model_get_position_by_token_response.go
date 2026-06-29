/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPositionByTokenResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPositionByTokenResponse{}

// GetPositionByTokenResponse struct for GetPositionByTokenResponse
type GetPositionByTokenResponse struct {
	Position             *GetPositionByTokenResponsePosition `json:"position,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPositionByTokenResponse GetPositionByTokenResponse

// NewGetPositionByTokenResponse instantiates a new GetPositionByTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPositionByTokenResponse() *GetPositionByTokenResponse {
	this := GetPositionByTokenResponse{}
	return &this
}

// NewGetPositionByTokenResponseWithDefaults instantiates a new GetPositionByTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPositionByTokenResponseWithDefaults() *GetPositionByTokenResponse {
	this := GetPositionByTokenResponse{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *GetPositionByTokenResponse) GetPosition() GetPositionByTokenResponsePosition {
	if o == nil || common.IsNil(o.Position) {
		var ret GetPositionByTokenResponsePosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionByTokenResponse) GetPositionOk() (*GetPositionByTokenResponsePosition, bool) {
	if o == nil || common.IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *GetPositionByTokenResponse) HasPosition() bool {
	if o != nil && !common.IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given GetPositionByTokenResponsePosition and assigns it to the Position field.
func (o *GetPositionByTokenResponse) SetPosition(v GetPositionByTokenResponsePosition) {
	o.Position = &v
}

func (o GetPositionByTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPositionByTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPositionByTokenResponse) UnmarshalJSON(data []byte) (err error) {
	varGetPositionByTokenResponse := _GetPositionByTokenResponse{}

	err = json.Unmarshal(data, &varGetPositionByTokenResponse)

	if err != nil {
		return err
	}

	*o = GetPositionByTokenResponse(varGetPositionByTokenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "position")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPositionByTokenResponse struct {
	value *GetPositionByTokenResponse
	isSet bool
}

func (v NullableGetPositionByTokenResponse) Get() *GetPositionByTokenResponse {
	return v.value
}

func (v *NullableGetPositionByTokenResponse) Set(val *GetPositionByTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPositionByTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPositionByTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPositionByTokenResponse(val *GetPositionByTokenResponse) *NullableGetPositionByTokenResponse {
	return &NullableGetPositionByTokenResponse{value: val, isSet: true}
}

func (v NullableGetPositionByTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPositionByTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
