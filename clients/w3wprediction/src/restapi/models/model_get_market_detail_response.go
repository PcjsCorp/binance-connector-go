/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetMarketDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetMarketDetailResponse{}

// GetMarketDetailResponse struct for GetMarketDetailResponse
type GetMarketDetailResponse struct {
	MarketTopicId        *int64                                 `json:"marketTopicId,omitempty"`
	Vendor               *string                                `json:"vendor,omitempty"`
	ChainId              *string                                `json:"chainId,omitempty"`
	Slug                 *string                                `json:"slug,omitempty"`
	Title                *string                                `json:"title,omitempty"`
	Question             *string                                `json:"question,omitempty"`
	Description          *string                                `json:"description,omitempty"`
	ImageUrl             *string                                `json:"imageUrl,omitempty"`
	TopicType            *string                                `json:"topicType,omitempty"`
	ChartType            *string                                `json:"chartType,omitempty"`
	Symbol               *string                                `json:"symbol,omitempty"`
	VariantData          *GetMarketDetailResponseVariantData    `json:"variantData,omitempty"`
	ParticipantCount     *int32                                 `json:"participantCount,omitempty"`
	Collateral           *string                                `json:"collateral,omitempty"`
	FeeRateBps           *int32                                 `json:"feeRateBps,omitempty"`
	SlippageBps          *int32                                 `json:"slippageBps,omitempty"`
	IsYieldBearing       *bool                                  `json:"isYieldBearing,omitempty"`
	TradeVolume          *string                                `json:"tradeVolume,omitempty"`
	Liquidity            *string                                `json:"liquidity,omitempty"`
	PublishedAt          *int64                                 `json:"publishedAt,omitempty"`
	StartDate            *int64                                 `json:"startDate,omitempty"`
	EndDate              *int64                                 `json:"endDate,omitempty"`
	Status               *string                                `json:"status,omitempty"`
	Timeline             []GetMarketDetailResponseTimelineInner `json:"timeline,omitempty"`
	Markets              []GetMarketDetailResponseMarketsInner  `json:"markets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMarketDetailResponse GetMarketDetailResponse

// NewGetMarketDetailResponse instantiates a new GetMarketDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketDetailResponse() *GetMarketDetailResponse {
	this := GetMarketDetailResponse{}
	return &this
}

// NewGetMarketDetailResponseWithDefaults instantiates a new GetMarketDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketDetailResponseWithDefaults() *GetMarketDetailResponse {
	this := GetMarketDetailResponse{}
	return &this
}

// GetMarketTopicId returns the MarketTopicId field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetMarketTopicId() int64 {
	if o == nil || common.IsNil(o.MarketTopicId) {
		var ret int64
		return ret
	}
	return *o.MarketTopicId
}

// GetMarketTopicIdOk returns a tuple with the MarketTopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetMarketTopicIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketTopicId) {
		return nil, false
	}
	return o.MarketTopicId, true
}

// HasMarketTopicId returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasMarketTopicId() bool {
	if o != nil && !common.IsNil(o.MarketTopicId) {
		return true
	}

	return false
}

// SetMarketTopicId gets a reference to the given int64 and assigns it to the MarketTopicId field.
func (o *GetMarketDetailResponse) SetMarketTopicId(v int64) {
	o.MarketTopicId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetVendor() string {
	if o == nil || common.IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetVendorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasVendor() bool {
	if o != nil && !common.IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *GetMarketDetailResponse) SetVendor(v string) {
	o.Vendor = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetChainId() string {
	if o == nil || common.IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetChainIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasChainId() bool {
	if o != nil && !common.IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *GetMarketDetailResponse) SetChainId(v string) {
	o.ChainId = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetSlug() string {
	if o == nil || common.IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetSlugOk() (*string, bool) {
	if o == nil || common.IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasSlug() bool {
	if o != nil && !common.IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *GetMarketDetailResponse) SetSlug(v string) {
	o.Slug = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetTitle() string {
	if o == nil || common.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasTitle() bool {
	if o != nil && !common.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetMarketDetailResponse) SetTitle(v string) {
	o.Title = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetQuestion() string {
	if o == nil || common.IsNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetQuestionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasQuestion() bool {
	if o != nil && !common.IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *GetMarketDetailResponse) SetQuestion(v string) {
	o.Question = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetMarketDetailResponse) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetImageUrl() string {
	if o == nil || common.IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetImageUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasImageUrl() bool {
	if o != nil && !common.IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GetMarketDetailResponse) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetTopicType returns the TopicType field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetTopicType() string {
	if o == nil || common.IsNil(o.TopicType) {
		var ret string
		return ret
	}
	return *o.TopicType
}

// GetTopicTypeOk returns a tuple with the TopicType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetTopicTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TopicType) {
		return nil, false
	}
	return o.TopicType, true
}

// HasTopicType returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasTopicType() bool {
	if o != nil && !common.IsNil(o.TopicType) {
		return true
	}

	return false
}

// SetTopicType gets a reference to the given string and assigns it to the TopicType field.
func (o *GetMarketDetailResponse) SetTopicType(v string) {
	o.TopicType = &v
}

// GetChartType returns the ChartType field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetChartType() string {
	if o == nil || common.IsNil(o.ChartType) {
		var ret string
		return ret
	}
	return *o.ChartType
}

// GetChartTypeOk returns a tuple with the ChartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetChartTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChartType) {
		return nil, false
	}
	return o.ChartType, true
}

// HasChartType returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasChartType() bool {
	if o != nil && !common.IsNil(o.ChartType) {
		return true
	}

	return false
}

// SetChartType gets a reference to the given string and assigns it to the ChartType field.
func (o *GetMarketDetailResponse) SetChartType(v string) {
	o.ChartType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetMarketDetailResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetVariantData returns the VariantData field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetVariantData() GetMarketDetailResponseVariantData {
	if o == nil || common.IsNil(o.VariantData) {
		var ret GetMarketDetailResponseVariantData
		return ret
	}
	return *o.VariantData
}

// GetVariantDataOk returns a tuple with the VariantData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetVariantDataOk() (*GetMarketDetailResponseVariantData, bool) {
	if o == nil || common.IsNil(o.VariantData) {
		return nil, false
	}
	return o.VariantData, true
}

// HasVariantData returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasVariantData() bool {
	if o != nil && !common.IsNil(o.VariantData) {
		return true
	}

	return false
}

// SetVariantData gets a reference to the given GetMarketDetailResponseVariantData and assigns it to the VariantData field.
func (o *GetMarketDetailResponse) SetVariantData(v GetMarketDetailResponseVariantData) {
	o.VariantData = &v
}

// GetParticipantCount returns the ParticipantCount field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetParticipantCount() int32 {
	if o == nil || common.IsNil(o.ParticipantCount) {
		var ret int32
		return ret
	}
	return *o.ParticipantCount
}

// GetParticipantCountOk returns a tuple with the ParticipantCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetParticipantCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.ParticipantCount) {
		return nil, false
	}
	return o.ParticipantCount, true
}

// HasParticipantCount returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasParticipantCount() bool {
	if o != nil && !common.IsNil(o.ParticipantCount) {
		return true
	}

	return false
}

// SetParticipantCount gets a reference to the given int32 and assigns it to the ParticipantCount field.
func (o *GetMarketDetailResponse) SetParticipantCount(v int32) {
	o.ParticipantCount = &v
}

// GetCollateral returns the Collateral field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetCollateral() string {
	if o == nil || common.IsNil(o.Collateral) {
		var ret string
		return ret
	}
	return *o.Collateral
}

// GetCollateralOk returns a tuple with the Collateral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetCollateralOk() (*string, bool) {
	if o == nil || common.IsNil(o.Collateral) {
		return nil, false
	}
	return o.Collateral, true
}

// HasCollateral returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasCollateral() bool {
	if o != nil && !common.IsNil(o.Collateral) {
		return true
	}

	return false
}

// SetCollateral gets a reference to the given string and assigns it to the Collateral field.
func (o *GetMarketDetailResponse) SetCollateral(v string) {
	o.Collateral = &v
}

// GetFeeRateBps returns the FeeRateBps field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetFeeRateBps() int32 {
	if o == nil || common.IsNil(o.FeeRateBps) {
		var ret int32
		return ret
	}
	return *o.FeeRateBps
}

// GetFeeRateBpsOk returns a tuple with the FeeRateBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetFeeRateBpsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.FeeRateBps) {
		return nil, false
	}
	return o.FeeRateBps, true
}

// HasFeeRateBps returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasFeeRateBps() bool {
	if o != nil && !common.IsNil(o.FeeRateBps) {
		return true
	}

	return false
}

// SetFeeRateBps gets a reference to the given int32 and assigns it to the FeeRateBps field.
func (o *GetMarketDetailResponse) SetFeeRateBps(v int32) {
	o.FeeRateBps = &v
}

// GetSlippageBps returns the SlippageBps field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetSlippageBps() int32 {
	if o == nil || common.IsNil(o.SlippageBps) {
		var ret int32
		return ret
	}
	return *o.SlippageBps
}

// GetSlippageBpsOk returns a tuple with the SlippageBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetSlippageBpsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.SlippageBps) {
		return nil, false
	}
	return o.SlippageBps, true
}

// HasSlippageBps returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasSlippageBps() bool {
	if o != nil && !common.IsNil(o.SlippageBps) {
		return true
	}

	return false
}

// SetSlippageBps gets a reference to the given int32 and assigns it to the SlippageBps field.
func (o *GetMarketDetailResponse) SetSlippageBps(v int32) {
	o.SlippageBps = &v
}

// GetIsYieldBearing returns the IsYieldBearing field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetIsYieldBearing() bool {
	if o == nil || common.IsNil(o.IsYieldBearing) {
		var ret bool
		return ret
	}
	return *o.IsYieldBearing
}

// GetIsYieldBearingOk returns a tuple with the IsYieldBearing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetIsYieldBearingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsYieldBearing) {
		return nil, false
	}
	return o.IsYieldBearing, true
}

// HasIsYieldBearing returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasIsYieldBearing() bool {
	if o != nil && !common.IsNil(o.IsYieldBearing) {
		return true
	}

	return false
}

// SetIsYieldBearing gets a reference to the given bool and assigns it to the IsYieldBearing field.
func (o *GetMarketDetailResponse) SetIsYieldBearing(v bool) {
	o.IsYieldBearing = &v
}

// GetTradeVolume returns the TradeVolume field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetTradeVolume() string {
	if o == nil || common.IsNil(o.TradeVolume) {
		var ret string
		return ret
	}
	return *o.TradeVolume
}

// GetTradeVolumeOk returns a tuple with the TradeVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetTradeVolumeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeVolume) {
		return nil, false
	}
	return o.TradeVolume, true
}

// HasTradeVolume returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasTradeVolume() bool {
	if o != nil && !common.IsNil(o.TradeVolume) {
		return true
	}

	return false
}

// SetTradeVolume gets a reference to the given string and assigns it to the TradeVolume field.
func (o *GetMarketDetailResponse) SetTradeVolume(v string) {
	o.TradeVolume = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *GetMarketDetailResponse) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetPublishedAt() int64 {
	if o == nil || common.IsNil(o.PublishedAt) {
		var ret int64
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetPublishedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasPublishedAt() bool {
	if o != nil && !common.IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given int64 and assigns it to the PublishedAt field.
func (o *GetMarketDetailResponse) SetPublishedAt(v int64) {
	o.PublishedAt = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetStartDate() int64 {
	if o == nil || common.IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetStartDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasStartDate() bool {
	if o != nil && !common.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *GetMarketDetailResponse) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetEndDate() int64 {
	if o == nil || common.IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetEndDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasEndDate() bool {
	if o != nil && !common.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *GetMarketDetailResponse) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetMarketDetailResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTimeline returns the Timeline field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetTimeline() []GetMarketDetailResponseTimelineInner {
	if o == nil || common.IsNil(o.Timeline) {
		var ret []GetMarketDetailResponseTimelineInner
		return ret
	}
	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetTimelineOk() ([]GetMarketDetailResponseTimelineInner, bool) {
	if o == nil || common.IsNil(o.Timeline) {
		return nil, false
	}
	return o.Timeline, true
}

// HasTimeline returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasTimeline() bool {
	if o != nil && !common.IsNil(o.Timeline) {
		return true
	}

	return false
}

// SetTimeline gets a reference to the given []GetMarketDetailResponseTimelineInner and assigns it to the Timeline field.
func (o *GetMarketDetailResponse) SetTimeline(v []GetMarketDetailResponseTimelineInner) {
	o.Timeline = v
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *GetMarketDetailResponse) GetMarkets() []GetMarketDetailResponseMarketsInner {
	if o == nil || common.IsNil(o.Markets) {
		var ret []GetMarketDetailResponseMarketsInner
		return ret
	}
	return o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketDetailResponse) GetMarketsOk() ([]GetMarketDetailResponseMarketsInner, bool) {
	if o == nil || common.IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *GetMarketDetailResponse) HasMarkets() bool {
	if o != nil && !common.IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given []GetMarketDetailResponseMarketsInner and assigns it to the Markets field.
func (o *GetMarketDetailResponse) SetMarkets(v []GetMarketDetailResponseMarketsInner) {
	o.Markets = v
}

func (o GetMarketDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketDetailResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
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
	if !common.IsNil(o.VariantData) {
		toSerialize["variantData"] = o.VariantData
	}
	if !common.IsNil(o.ParticipantCount) {
		toSerialize["participantCount"] = o.ParticipantCount
	}
	if !common.IsNil(o.Collateral) {
		toSerialize["collateral"] = o.Collateral
	}
	if !common.IsNil(o.FeeRateBps) {
		toSerialize["feeRateBps"] = o.FeeRateBps
	}
	if !common.IsNil(o.SlippageBps) {
		toSerialize["slippageBps"] = o.SlippageBps
	}
	if !common.IsNil(o.IsYieldBearing) {
		toSerialize["isYieldBearing"] = o.IsYieldBearing
	}
	if !common.IsNil(o.TradeVolume) {
		toSerialize["tradeVolume"] = o.TradeVolume
	}
	if !common.IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !common.IsNil(o.PublishedAt) {
		toSerialize["publishedAt"] = o.PublishedAt
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
	if !common.IsNil(o.Timeline) {
		toSerialize["timeline"] = o.Timeline
	}
	if !common.IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMarketDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varGetMarketDetailResponse := _GetMarketDetailResponse{}

	err = json.Unmarshal(data, &varGetMarketDetailResponse)

	if err != nil {
		return err
	}

	*o = GetMarketDetailResponse(varGetMarketDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "marketTopicId")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "chainId")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "title")
		delete(additionalProperties, "question")
		delete(additionalProperties, "description")
		delete(additionalProperties, "imageUrl")
		delete(additionalProperties, "topicType")
		delete(additionalProperties, "chartType")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "variantData")
		delete(additionalProperties, "participantCount")
		delete(additionalProperties, "collateral")
		delete(additionalProperties, "feeRateBps")
		delete(additionalProperties, "slippageBps")
		delete(additionalProperties, "isYieldBearing")
		delete(additionalProperties, "tradeVolume")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "publishedAt")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeline")
		delete(additionalProperties, "markets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMarketDetailResponse struct {
	value *GetMarketDetailResponse
	isSet bool
}

func (v NullableGetMarketDetailResponse) Get() *GetMarketDetailResponse {
	return v.value
}

func (v *NullableGetMarketDetailResponse) Set(val *GetMarketDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketDetailResponse(val *GetMarketDetailResponse) *NullableGetMarketDetailResponse {
	return &NullableGetMarketDetailResponse{value: val, isSet: true}
}

func (v NullableGetMarketDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
