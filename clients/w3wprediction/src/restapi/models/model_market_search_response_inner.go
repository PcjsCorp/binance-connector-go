/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarketSearchResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarketSearchResponseInner{}

// MarketSearchResponseInner struct for MarketSearchResponseInner
type MarketSearchResponseInner struct {
	MarketTopicId        *int64                   `json:"marketTopicId,omitempty"`
	Vendor               *string                  `json:"vendor,omitempty"`
	ChainId              *string                  `json:"chainId,omitempty"`
	Slug                 *string                  `json:"slug,omitempty"`
	Title                *string                  `json:"title,omitempty"`
	Question             *string                  `json:"question,omitempty"`
	Description          *string                  `json:"description,omitempty"`
	TopicType            *string                  `json:"topicType,omitempty"`
	ChartType            *string                  `json:"chartType,omitempty"`
	Symbol               *string                  `json:"symbol,omitempty"`
	ParticipantCount     *int32                   `json:"participantCount,omitempty"`
	Collateral           *string                  `json:"collateral,omitempty"`
	TradeVolume          *string                  `json:"tradeVolume,omitempty"`
	Liquidity            *string                  `json:"liquidity,omitempty"`
	StartDate            *int64                   `json:"startDate,omitempty"`
	EndDate              *int64                   `json:"endDate,omitempty"`
	Status               *string                  `json:"status,omitempty"`
	Markets              []map[string]interface{} `json:"markets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketSearchResponseInner MarketSearchResponseInner

// NewMarketSearchResponseInner instantiates a new MarketSearchResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketSearchResponseInner() *MarketSearchResponseInner {
	this := MarketSearchResponseInner{}
	return &this
}

// NewMarketSearchResponseInnerWithDefaults instantiates a new MarketSearchResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketSearchResponseInnerWithDefaults() *MarketSearchResponseInner {
	this := MarketSearchResponseInner{}
	return &this
}

// GetMarketTopicId returns the MarketTopicId field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetMarketTopicId() int64 {
	if o == nil || common.IsNil(o.MarketTopicId) {
		var ret int64
		return ret
	}
	return *o.MarketTopicId
}

// GetMarketTopicIdOk returns a tuple with the MarketTopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetMarketTopicIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketTopicId) {
		return nil, false
	}
	return o.MarketTopicId, true
}

// HasMarketTopicId returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasMarketTopicId() bool {
	if o != nil && !common.IsNil(o.MarketTopicId) {
		return true
	}

	return false
}

// SetMarketTopicId gets a reference to the given int64 and assigns it to the MarketTopicId field.
func (o *MarketSearchResponseInner) SetMarketTopicId(v int64) {
	o.MarketTopicId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetVendor() string {
	if o == nil || common.IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetVendorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasVendor() bool {
	if o != nil && !common.IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *MarketSearchResponseInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetChainId() string {
	if o == nil || common.IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetChainIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasChainId() bool {
	if o != nil && !common.IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *MarketSearchResponseInner) SetChainId(v string) {
	o.ChainId = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetSlug() string {
	if o == nil || common.IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetSlugOk() (*string, bool) {
	if o == nil || common.IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasSlug() bool {
	if o != nil && !common.IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *MarketSearchResponseInner) SetSlug(v string) {
	o.Slug = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetTitle() string {
	if o == nil || common.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasTitle() bool {
	if o != nil && !common.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MarketSearchResponseInner) SetTitle(v string) {
	o.Title = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetQuestion() string {
	if o == nil || common.IsNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetQuestionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasQuestion() bool {
	if o != nil && !common.IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *MarketSearchResponseInner) SetQuestion(v string) {
	o.Question = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MarketSearchResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetTopicType returns the TopicType field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetTopicType() string {
	if o == nil || common.IsNil(o.TopicType) {
		var ret string
		return ret
	}
	return *o.TopicType
}

// GetTopicTypeOk returns a tuple with the TopicType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetTopicTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TopicType) {
		return nil, false
	}
	return o.TopicType, true
}

// HasTopicType returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasTopicType() bool {
	if o != nil && !common.IsNil(o.TopicType) {
		return true
	}

	return false
}

// SetTopicType gets a reference to the given string and assigns it to the TopicType field.
func (o *MarketSearchResponseInner) SetTopicType(v string) {
	o.TopicType = &v
}

// GetChartType returns the ChartType field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetChartType() string {
	if o == nil || common.IsNil(o.ChartType) {
		var ret string
		return ret
	}
	return *o.ChartType
}

// GetChartTypeOk returns a tuple with the ChartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetChartTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChartType) {
		return nil, false
	}
	return o.ChartType, true
}

// HasChartType returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasChartType() bool {
	if o != nil && !common.IsNil(o.ChartType) {
		return true
	}

	return false
}

// SetChartType gets a reference to the given string and assigns it to the ChartType field.
func (o *MarketSearchResponseInner) SetChartType(v string) {
	o.ChartType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarketSearchResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetParticipantCount returns the ParticipantCount field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetParticipantCount() int32 {
	if o == nil || common.IsNil(o.ParticipantCount) {
		var ret int32
		return ret
	}
	return *o.ParticipantCount
}

// GetParticipantCountOk returns a tuple with the ParticipantCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetParticipantCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.ParticipantCount) {
		return nil, false
	}
	return o.ParticipantCount, true
}

// HasParticipantCount returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasParticipantCount() bool {
	if o != nil && !common.IsNil(o.ParticipantCount) {
		return true
	}

	return false
}

// SetParticipantCount gets a reference to the given int32 and assigns it to the ParticipantCount field.
func (o *MarketSearchResponseInner) SetParticipantCount(v int32) {
	o.ParticipantCount = &v
}

// GetCollateral returns the Collateral field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetCollateral() string {
	if o == nil || common.IsNil(o.Collateral) {
		var ret string
		return ret
	}
	return *o.Collateral
}

// GetCollateralOk returns a tuple with the Collateral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetCollateralOk() (*string, bool) {
	if o == nil || common.IsNil(o.Collateral) {
		return nil, false
	}
	return o.Collateral, true
}

// HasCollateral returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasCollateral() bool {
	if o != nil && !common.IsNil(o.Collateral) {
		return true
	}

	return false
}

// SetCollateral gets a reference to the given string and assigns it to the Collateral field.
func (o *MarketSearchResponseInner) SetCollateral(v string) {
	o.Collateral = &v
}

// GetTradeVolume returns the TradeVolume field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetTradeVolume() string {
	if o == nil || common.IsNil(o.TradeVolume) {
		var ret string
		return ret
	}
	return *o.TradeVolume
}

// GetTradeVolumeOk returns a tuple with the TradeVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetTradeVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeVolume) {
		return nil, false
	}
	return o.TradeVolume, true
}

// HasTradeVolume returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasTradeVolume() bool {
	if o != nil && !common.IsNil(o.TradeVolume) {
		return true
	}

	return false
}

// SetTradeVolume gets a reference to the given string and assigns it to the TradeVolume field.
func (o *MarketSearchResponseInner) SetTradeVolume(v string) {
	o.TradeVolume = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *MarketSearchResponseInner) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetStartDate() int64 {
	if o == nil || common.IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetStartDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasStartDate() bool {
	if o != nil && !common.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *MarketSearchResponseInner) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetEndDate() int64 {
	if o == nil || common.IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetEndDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasEndDate() bool {
	if o != nil && !common.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *MarketSearchResponseInner) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MarketSearchResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *MarketSearchResponseInner) GetMarkets() []map[string]interface{} {
	if o == nil || common.IsNil(o.Markets) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketSearchResponseInner) GetMarketsOk() ([]map[string]interface{}, bool) {
	if o == nil || common.IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *MarketSearchResponseInner) HasMarkets() bool {
	if o != nil && !common.IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given []map[string]interface{} and assigns it to the Markets field.
func (o *MarketSearchResponseInner) SetMarkets(v []map[string]interface{}) {
	o.Markets = v
}

func (o MarketSearchResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketSearchResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MarketTopicId) {
		toSerialize["marketTopicId"] = o.MarketTopicId
	}
	if !common.IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !common.IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !common.IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !common.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !common.IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.TopicType) {
		toSerialize["topicType"] = o.TopicType
	}
	if !common.IsNil(o.ChartType) {
		toSerialize["chartType"] = o.ChartType
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.ParticipantCount) {
		toSerialize["participantCount"] = o.ParticipantCount
	}
	if !common.IsNil(o.Collateral) {
		toSerialize["collateral"] = o.Collateral
	}
	if !common.IsNil(o.TradeVolume) {
		toSerialize["tradeVolume"] = o.TradeVolume
	}
	if !common.IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !common.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !common.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarketSearchResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMarketSearchResponseInner := _MarketSearchResponseInner{}

	err = json.Unmarshal(data, &varMarketSearchResponseInner)

	if err != nil {
		return err
	}

	*o = MarketSearchResponseInner(varMarketSearchResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marketTopicId")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "chainId")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "title")
		delete(additionalProperties, "question")
		delete(additionalProperties, "description")
		delete(additionalProperties, "topicType")
		delete(additionalProperties, "chartType")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "participantCount")
		delete(additionalProperties, "collateral")
		delete(additionalProperties, "tradeVolume")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "status")
		delete(additionalProperties, "markets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarketSearchResponseInner struct {
	value *MarketSearchResponseInner
	isSet bool
}

func (v NullableMarketSearchResponseInner) Get() *MarketSearchResponseInner {
	return v.value
}

func (v *NullableMarketSearchResponseInner) Set(val *MarketSearchResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketSearchResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketSearchResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketSearchResponseInner(val *MarketSearchResponseInner) *NullableMarketSearchResponseInner {
	return &NullableMarketSearchResponseInner{value: val, isSet: true}
}

func (v NullableMarketSearchResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketSearchResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
