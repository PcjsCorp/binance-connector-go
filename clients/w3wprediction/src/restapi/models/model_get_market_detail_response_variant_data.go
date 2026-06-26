/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarketDetailResponseVariantData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarketDetailResponseVariantData{}

// GetMarketDetailResponseVariantData struct for GetMarketDetailResponseVariantData
type GetMarketDetailResponseVariantData struct {
	Type                 *string `json:"type,omitempty"`
	StartPrice           *string `json:"startPrice,omitempty"`
	EndPrice             *string `json:"endPrice,omitempty"`
	PriceFeedId          *string `json:"priceFeedId,omitempty"`
	PriceFeedProvider    *string `json:"priceFeedProvider,omitempty"`
	PriceFeedSymbol      *string `json:"priceFeedSymbol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMarketDetailResponseVariantData GetMarketDetailResponseVariantData

// NewGetMarketDetailResponseVariantData instantiates a new GetMarketDetailResponseVariantData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketDetailResponseVariantData() *GetMarketDetailResponseVariantData {
	this := GetMarketDetailResponseVariantData{}
	return &this
}

// NewGetMarketDetailResponseVariantDataWithDefaults instantiates a new GetMarketDetailResponseVariantData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketDetailResponseVariantDataWithDefaults() *GetMarketDetailResponseVariantData {
	this := GetMarketDetailResponseVariantData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetMarketDetailResponseVariantData) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseVariantData) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetMarketDetailResponseVariantData) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetMarketDetailResponseVariantData) SetType(v string) {
	o.Type = &v
}

// GetStartPrice returns the StartPrice field value if set, zero value otherwise.
func (o *GetMarketDetailResponseVariantData) GetStartPrice() string {
	if o == nil || common.IsNil(o.StartPrice) {
		var ret string
		return ret
	}
	return *o.StartPrice
}

// GetStartPriceOk returns a tuple with the StartPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseVariantData) GetStartPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StartPrice) {
		return nil, false
	}
	return o.StartPrice, true
}

// HasStartPrice returns a boolean if a field has been set.
func (o *GetMarketDetailResponseVariantData) HasStartPrice() bool {
	if o != nil && !common.IsNil(o.StartPrice) {
		return true
	}

	return false
}

// SetStartPrice gets a reference to the given string and assigns it to the StartPrice field.
func (o *GetMarketDetailResponseVariantData) SetStartPrice(v string) {
	o.StartPrice = &v
}

// GetEndPrice returns the EndPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetMarketDetailResponseVariantData) GetEndPrice() string {
	if o == nil || common.IsNil(o.EndPrice) {
		var ret string
		return ret
	}
	return *o.EndPrice
}

// GetEndPriceOk returns a tuple with the EndPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetMarketDetailResponseVariantData) GetEndPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EndPrice) {
		return nil, false
	}
	return o.EndPrice, true
}

// HasEndPrice returns a boolean if a field has been set.
func (o *GetMarketDetailResponseVariantData) HasEndPrice() bool {
	if o != nil && !common.IsNil(o.EndPrice) {
		return true
	}

	return false
}

// SetEndPrice gets a reference to the given NullableString and assigns it to the EndPrice field.
func (o *GetMarketDetailResponseVariantData) SetEndPrice(v string) {
	o.EndPrice = &v
}

// GetPriceFeedId returns the PriceFeedId field value if set, zero value otherwise.
func (o *GetMarketDetailResponseVariantData) GetPriceFeedId() string {
	if o == nil || common.IsNil(o.PriceFeedId) {
		var ret string
		return ret
	}
	return *o.PriceFeedId
}

// GetPriceFeedIdOk returns a tuple with the PriceFeedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseVariantData) GetPriceFeedIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceFeedId) {
		return nil, false
	}
	return o.PriceFeedId, true
}

// HasPriceFeedId returns a boolean if a field has been set.
func (o *GetMarketDetailResponseVariantData) HasPriceFeedId() bool {
	if o != nil && !common.IsNil(o.PriceFeedId) {
		return true
	}

	return false
}

// SetPriceFeedId gets a reference to the given string and assigns it to the PriceFeedId field.
func (o *GetMarketDetailResponseVariantData) SetPriceFeedId(v string) {
	o.PriceFeedId = &v
}

// GetPriceFeedProvider returns the PriceFeedProvider field value if set, zero value otherwise.
func (o *GetMarketDetailResponseVariantData) GetPriceFeedProvider() string {
	if o == nil || common.IsNil(o.PriceFeedProvider) {
		var ret string
		return ret
	}
	return *o.PriceFeedProvider
}

// GetPriceFeedProviderOk returns a tuple with the PriceFeedProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseVariantData) GetPriceFeedProviderOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceFeedProvider) {
		return nil, false
	}
	return o.PriceFeedProvider, true
}

// HasPriceFeedProvider returns a boolean if a field has been set.
func (o *GetMarketDetailResponseVariantData) HasPriceFeedProvider() bool {
	if o != nil && !common.IsNil(o.PriceFeedProvider) {
		return true
	}

	return false
}

// SetPriceFeedProvider gets a reference to the given string and assigns it to the PriceFeedProvider field.
func (o *GetMarketDetailResponseVariantData) SetPriceFeedProvider(v string) {
	o.PriceFeedProvider = &v
}

// GetPriceFeedSymbol returns the PriceFeedSymbol field value if set, zero value otherwise.
func (o *GetMarketDetailResponseVariantData) GetPriceFeedSymbol() string {
	if o == nil || common.IsNil(o.PriceFeedSymbol) {
		var ret string
		return ret
	}
	return *o.PriceFeedSymbol
}

// GetPriceFeedSymbolOk returns a tuple with the PriceFeedSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponseVariantData) GetPriceFeedSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceFeedSymbol) {
		return nil, false
	}
	return o.PriceFeedSymbol, true
}

// HasPriceFeedSymbol returns a boolean if a field has been set.
func (o *GetMarketDetailResponseVariantData) HasPriceFeedSymbol() bool {
	if o != nil && !common.IsNil(o.PriceFeedSymbol) {
		return true
	}

	return false
}

// SetPriceFeedSymbol gets a reference to the given string and assigns it to the PriceFeedSymbol field.
func (o *GetMarketDetailResponseVariantData) SetPriceFeedSymbol(v string) {
	o.PriceFeedSymbol = &v
}

func (o GetMarketDetailResponseVariantData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketDetailResponseVariantData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.StartPrice) {
		toSerialize["startPrice"] = o.StartPrice
	}
	if !common.IsNil(o.EndPrice) {
		toSerialize["endPrice"] = o.EndPrice
	}
	if !common.IsNil(o.PriceFeedId) {
		toSerialize["priceFeedId"] = o.PriceFeedId
	}
	if !common.IsNil(o.PriceFeedProvider) {
		toSerialize["priceFeedProvider"] = o.PriceFeedProvider
	}
	if !common.IsNil(o.PriceFeedSymbol) {
		toSerialize["priceFeedSymbol"] = o.PriceFeedSymbol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarketDetailResponseVariantData) UnmarshalJSON(data []byte) (err error) {
	varGetMarketDetailResponseVariantData := _GetMarketDetailResponseVariantData{}

	err = json.Unmarshal(data, &varGetMarketDetailResponseVariantData)

	if err != nil {
		return err
	}

	*o = GetMarketDetailResponseVariantData(varGetMarketDetailResponseVariantData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "startPrice")
		delete(additionalProperties, "endPrice")
		delete(additionalProperties, "priceFeedId")
		delete(additionalProperties, "priceFeedProvider")
		delete(additionalProperties, "priceFeedSymbol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarketDetailResponseVariantData struct {
	value *GetMarketDetailResponseVariantData
	isSet bool
}

func (v NullableGetMarketDetailResponseVariantData) Get() *GetMarketDetailResponseVariantData {
	return v.value
}

func (v *NullableGetMarketDetailResponseVariantData) Set(val *GetMarketDetailResponseVariantData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketDetailResponseVariantData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketDetailResponseVariantData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketDetailResponseVariantData(val *GetMarketDetailResponseVariantData) *NullableGetMarketDetailResponseVariantData {
	return &NullableGetMarketDetailResponseVariantData{value: val, isSet: true}
}

func (v NullableGetMarketDetailResponseVariantData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketDetailResponseVariantData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
