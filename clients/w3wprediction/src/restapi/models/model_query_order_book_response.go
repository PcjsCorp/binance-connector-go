/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryOrderBookResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOrderBookResponse{}

// QueryOrderBookResponse struct for QueryOrderBookResponse
type QueryOrderBookResponse struct {
	Outcome              *string                           `json:"outcome,omitempty"`
	TokenId              *string                           `json:"tokenId,omitempty"`
	Timestamp            *int64                            `json:"timestamp,omitempty"`
	Bids                 []QueryOrderBookResponseBidsInner `json:"bids,omitempty"`
	Asks                 []QueryOrderBookResponseAsksInner `json:"asks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryOrderBookResponse QueryOrderBookResponse

// NewQueryOrderBookResponse instantiates a new QueryOrderBookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOrderBookResponse() *QueryOrderBookResponse {
	this := QueryOrderBookResponse{}
	return &this
}

// NewQueryOrderBookResponseWithDefaults instantiates a new QueryOrderBookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOrderBookResponseWithDefaults() *QueryOrderBookResponse {
	this := QueryOrderBookResponse{}
	return &this
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *QueryOrderBookResponse) GetOutcome() string {
	if o == nil || common.IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponse) GetOutcomeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *QueryOrderBookResponse) HasOutcome() bool {
	if o != nil && !common.IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *QueryOrderBookResponse) SetOutcome(v string) {
	o.Outcome = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *QueryOrderBookResponse) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponse) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *QueryOrderBookResponse) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *QueryOrderBookResponse) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QueryOrderBookResponse) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponse) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QueryOrderBookResponse) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QueryOrderBookResponse) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *QueryOrderBookResponse) GetBids() []QueryOrderBookResponseBidsInner {
	if o == nil || common.IsNil(o.Bids) {
		var ret []QueryOrderBookResponseBidsInner
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponse) GetBidsOk() ([]QueryOrderBookResponseBidsInner, bool) {
	if o == nil || common.IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *QueryOrderBookResponse) HasBids() bool {
	if o != nil && !common.IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given []QueryOrderBookResponseBidsInner and assigns it to the Bids field.
func (o *QueryOrderBookResponse) SetBids(v []QueryOrderBookResponseBidsInner) {
	o.Bids = v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *QueryOrderBookResponse) GetAsks() []QueryOrderBookResponseAsksInner {
	if o == nil || common.IsNil(o.Asks) {
		var ret []QueryOrderBookResponseAsksInner
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderBookResponse) GetAsksOk() ([]QueryOrderBookResponseAsksInner, bool) {
	if o == nil || common.IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *QueryOrderBookResponse) HasAsks() bool {
	if o != nil && !common.IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given []QueryOrderBookResponseAsksInner and assigns it to the Asks field.
func (o *QueryOrderBookResponse) SetAsks(v []QueryOrderBookResponseAsksInner) {
	o.Asks = v
}

func (o QueryOrderBookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderBookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.Bids) {
		toSerialize["bids"] = o.Bids
	}
	if !common.IsNil(o.Asks) {
		toSerialize["asks"] = o.Asks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryOrderBookResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryOrderBookResponse := _QueryOrderBookResponse{}

	err = json.Unmarshal(data, &varQueryOrderBookResponse)

	if err != nil {
		return err
	}

	*o = QueryOrderBookResponse(varQueryOrderBookResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "bids")
		delete(additionalProperties, "asks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryOrderBookResponse struct {
	value *QueryOrderBookResponse
	isSet bool
}

func (v NullableQueryOrderBookResponse) Get() *QueryOrderBookResponse {
	return v.value
}

func (v *NullableQueryOrderBookResponse) Set(val *QueryOrderBookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOrderBookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOrderBookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOrderBookResponse(val *QueryOrderBookResponse) *NullableQueryOrderBookResponse {
	return &NullableQueryOrderBookResponse{value: val, isSet: true}
}

func (v NullableQueryOrderBookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOrderBookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
