/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryActiveOrdersResponseOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryActiveOrdersResponseOrdersInner{}

// QueryActiveOrdersResponseOrdersInner struct for QueryActiveOrdersResponseOrdersInner
type QueryActiveOrdersResponseOrdersInner struct {
	OrderId              *string `json:"orderId,omitempty"`
	VendorOrderId        *string `json:"vendorOrderId,omitempty"`
	Vendor               *string `json:"vendor,omitempty"`
	MarketTopicId        *int64  `json:"marketTopicId,omitempty"`
	Slug                 *string `json:"slug,omitempty"`
	MarketTopicTitle     *string `json:"marketTopicTitle,omitempty"`
	MarketId             *int64  `json:"marketId,omitempty"`
	MarketTitle          *string `json:"marketTitle,omitempty"`
	Outcome              *string `json:"outcome,omitempty"`
	OutcomeIndex         *int32  `json:"outcomeIndex,omitempty"`
	Status               *string `json:"status,omitempty"`
	Side                 *string `json:"side,omitempty"`
	OrderType            *string `json:"orderType,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	ModifyTime           *int64  `json:"modifyTime,omitempty"`
	MakerUsdtAmount      *string `json:"makerUsdtAmount,omitempty"`
	MakerShareQty        *string `json:"makerShareQty,omitempty"`
	FilledUsdtAmount     *string `json:"filledUsdtAmount,omitempty"`
	FilledShareQty       *string `json:"filledShareQty,omitempty"`
	FillPercentage       *string `json:"fillPercentage,omitempty"`
	Price                *string `json:"price,omitempty"`
	MarketProviderFee    *string `json:"marketProviderFee,omitempty"`
	NetworkFee           *string `json:"networkFee,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryActiveOrdersResponseOrdersInner QueryActiveOrdersResponseOrdersInner

// NewQueryActiveOrdersResponseOrdersInner instantiates a new QueryActiveOrdersResponseOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryActiveOrdersResponseOrdersInner() *QueryActiveOrdersResponseOrdersInner {
	this := QueryActiveOrdersResponseOrdersInner{}
	return &this
}

// NewQueryActiveOrdersResponseOrdersInnerWithDefaults instantiates a new QueryActiveOrdersResponseOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryActiveOrdersResponseOrdersInnerWithDefaults() *QueryActiveOrdersResponseOrdersInner {
	this := QueryActiveOrdersResponseOrdersInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *QueryActiveOrdersResponseOrdersInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetVendorOrderId returns the VendorOrderId field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetVendorOrderId() string {
	if o == nil || common.IsNil(o.VendorOrderId) {
		var ret string
		return ret
	}
	return *o.VendorOrderId
}

// GetVendorOrderIdOk returns a tuple with the VendorOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetVendorOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.VendorOrderId) {
		return nil, false
	}
	return o.VendorOrderId, true
}

// HasVendorOrderId returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasVendorOrderId() bool {
	if o != nil && !common.IsNil(o.VendorOrderId) {
		return true
	}

	return false
}

// SetVendorOrderId gets a reference to the given string and assigns it to the VendorOrderId field.
func (o *QueryActiveOrdersResponseOrdersInner) SetVendorOrderId(v string) {
	o.VendorOrderId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetVendor() string {
	if o == nil || common.IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetVendorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasVendor() bool {
	if o != nil && !common.IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *QueryActiveOrdersResponseOrdersInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetMarketTopicId returns the MarketTopicId field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketTopicId() int64 {
	if o == nil || common.IsNil(o.MarketTopicId) {
		var ret int64
		return ret
	}
	return *o.MarketTopicId
}

// GetMarketTopicIdOk returns a tuple with the MarketTopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketTopicIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketTopicId) {
		return nil, false
	}
	return o.MarketTopicId, true
}

// HasMarketTopicId returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasMarketTopicId() bool {
	if o != nil && !common.IsNil(o.MarketTopicId) {
		return true
	}

	return false
}

// SetMarketTopicId gets a reference to the given int64 and assigns it to the MarketTopicId field.
func (o *QueryActiveOrdersResponseOrdersInner) SetMarketTopicId(v int64) {
	o.MarketTopicId = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetSlug() string {
	if o == nil || common.IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetSlugOk() (*string, bool) {
	if o == nil || common.IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasSlug() bool {
	if o != nil && !common.IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *QueryActiveOrdersResponseOrdersInner) SetSlug(v string) {
	o.Slug = &v
}

// GetMarketTopicTitle returns the MarketTopicTitle field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketTopicTitle() string {
	if o == nil || common.IsNil(o.MarketTopicTitle) {
		var ret string
		return ret
	}
	return *o.MarketTopicTitle
}

// GetMarketTopicTitleOk returns a tuple with the MarketTopicTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketTopicTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketTopicTitle) {
		return nil, false
	}
	return o.MarketTopicTitle, true
}

// HasMarketTopicTitle returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasMarketTopicTitle() bool {
	if o != nil && !common.IsNil(o.MarketTopicTitle) {
		return true
	}

	return false
}

// SetMarketTopicTitle gets a reference to the given string and assigns it to the MarketTopicTitle field.
func (o *QueryActiveOrdersResponseOrdersInner) SetMarketTopicTitle(v string) {
	o.MarketTopicTitle = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketId() int64 {
	if o == nil || common.IsNil(o.MarketId) {
		var ret int64
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MarketId) {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasMarketId() bool {
	if o != nil && !common.IsNil(o.MarketId) {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given int64 and assigns it to the MarketId field.
func (o *QueryActiveOrdersResponseOrdersInner) SetMarketId(v int64) {
	o.MarketId = &v
}

// GetMarketTitle returns the MarketTitle field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketTitle() string {
	if o == nil || common.IsNil(o.MarketTitle) {
		var ret string
		return ret
	}
	return *o.MarketTitle
}

// GetMarketTitleOk returns a tuple with the MarketTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketTitle) {
		return nil, false
	}
	return o.MarketTitle, true
}

// HasMarketTitle returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasMarketTitle() bool {
	if o != nil && !common.IsNil(o.MarketTitle) {
		return true
	}

	return false
}

// SetMarketTitle gets a reference to the given string and assigns it to the MarketTitle field.
func (o *QueryActiveOrdersResponseOrdersInner) SetMarketTitle(v string) {
	o.MarketTitle = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetOutcome() string {
	if o == nil || common.IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetOutcomeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasOutcome() bool {
	if o != nil && !common.IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *QueryActiveOrdersResponseOrdersInner) SetOutcome(v string) {
	o.Outcome = &v
}

// GetOutcomeIndex returns the OutcomeIndex field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetOutcomeIndex() int32 {
	if o == nil || common.IsNil(o.OutcomeIndex) {
		var ret int32
		return ret
	}
	return *o.OutcomeIndex
}

// GetOutcomeIndexOk returns a tuple with the OutcomeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetOutcomeIndexOk() (*int32, bool) {
	if o == nil || common.IsNil(o.OutcomeIndex) {
		return nil, false
	}
	return o.OutcomeIndex, true
}

// HasOutcomeIndex returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasOutcomeIndex() bool {
	if o != nil && !common.IsNil(o.OutcomeIndex) {
		return true
	}

	return false
}

// SetOutcomeIndex gets a reference to the given int32 and assigns it to the OutcomeIndex field.
func (o *QueryActiveOrdersResponseOrdersInner) SetOutcomeIndex(v int32) {
	o.OutcomeIndex = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryActiveOrdersResponseOrdersInner) SetStatus(v string) {
	o.Status = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryActiveOrdersResponseOrdersInner) SetSide(v string) {
	o.Side = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *QueryActiveOrdersResponseOrdersInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *QueryActiveOrdersResponseOrdersInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetModifyTime() int64 {
	if o == nil || common.IsNil(o.ModifyTime) {
		var ret int64
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetModifyTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ModifyTime) {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasModifyTime() bool {
	if o != nil && !common.IsNil(o.ModifyTime) {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given int64 and assigns it to the ModifyTime field.
func (o *QueryActiveOrdersResponseOrdersInner) SetModifyTime(v int64) {
	o.ModifyTime = &v
}

// GetMakerUsdtAmount returns the MakerUsdtAmount field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetMakerUsdtAmount() string {
	if o == nil || common.IsNil(o.MakerUsdtAmount) {
		var ret string
		return ret
	}
	return *o.MakerUsdtAmount
}

// GetMakerUsdtAmountOk returns a tuple with the MakerUsdtAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetMakerUsdtAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerUsdtAmount) {
		return nil, false
	}
	return o.MakerUsdtAmount, true
}

// HasMakerUsdtAmount returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasMakerUsdtAmount() bool {
	if o != nil && !common.IsNil(o.MakerUsdtAmount) {
		return true
	}

	return false
}

// SetMakerUsdtAmount gets a reference to the given string and assigns it to the MakerUsdtAmount field.
func (o *QueryActiveOrdersResponseOrdersInner) SetMakerUsdtAmount(v string) {
	o.MakerUsdtAmount = &v
}

// GetMakerShareQty returns the MakerShareQty field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetMakerShareQty() string {
	if o == nil || common.IsNil(o.MakerShareQty) {
		var ret string
		return ret
	}
	return *o.MakerShareQty
}

// GetMakerShareQtyOk returns a tuple with the MakerShareQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetMakerShareQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerShareQty) {
		return nil, false
	}
	return o.MakerShareQty, true
}

// HasMakerShareQty returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasMakerShareQty() bool {
	if o != nil && !common.IsNil(o.MakerShareQty) {
		return true
	}

	return false
}

// SetMakerShareQty gets a reference to the given string and assigns it to the MakerShareQty field.
func (o *QueryActiveOrdersResponseOrdersInner) SetMakerShareQty(v string) {
	o.MakerShareQty = &v
}

// GetFilledUsdtAmount returns the FilledUsdtAmount field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetFilledUsdtAmount() string {
	if o == nil || common.IsNil(o.FilledUsdtAmount) {
		var ret string
		return ret
	}
	return *o.FilledUsdtAmount
}

// GetFilledUsdtAmountOk returns a tuple with the FilledUsdtAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetFilledUsdtAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledUsdtAmount) {
		return nil, false
	}
	return o.FilledUsdtAmount, true
}

// HasFilledUsdtAmount returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasFilledUsdtAmount() bool {
	if o != nil && !common.IsNil(o.FilledUsdtAmount) {
		return true
	}

	return false
}

// SetFilledUsdtAmount gets a reference to the given string and assigns it to the FilledUsdtAmount field.
func (o *QueryActiveOrdersResponseOrdersInner) SetFilledUsdtAmount(v string) {
	o.FilledUsdtAmount = &v
}

// GetFilledShareQty returns the FilledShareQty field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetFilledShareQty() string {
	if o == nil || common.IsNil(o.FilledShareQty) {
		var ret string
		return ret
	}
	return *o.FilledShareQty
}

// GetFilledShareQtyOk returns a tuple with the FilledShareQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetFilledShareQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledShareQty) {
		return nil, false
	}
	return o.FilledShareQty, true
}

// HasFilledShareQty returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasFilledShareQty() bool {
	if o != nil && !common.IsNil(o.FilledShareQty) {
		return true
	}

	return false
}

// SetFilledShareQty gets a reference to the given string and assigns it to the FilledShareQty field.
func (o *QueryActiveOrdersResponseOrdersInner) SetFilledShareQty(v string) {
	o.FilledShareQty = &v
}

// GetFillPercentage returns the FillPercentage field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetFillPercentage() string {
	if o == nil || common.IsNil(o.FillPercentage) {
		var ret string
		return ret
	}
	return *o.FillPercentage
}

// GetFillPercentageOk returns a tuple with the FillPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetFillPercentageOk() (*string, bool) {
	if o == nil || common.IsNil(o.FillPercentage) {
		return nil, false
	}
	return o.FillPercentage, true
}

// HasFillPercentage returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasFillPercentage() bool {
	if o != nil && !common.IsNil(o.FillPercentage) {
		return true
	}

	return false
}

// SetFillPercentage gets a reference to the given string and assigns it to the FillPercentage field.
func (o *QueryActiveOrdersResponseOrdersInner) SetFillPercentage(v string) {
	o.FillPercentage = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryActiveOrdersResponseOrdersInner) SetPrice(v string) {
	o.Price = &v
}

// GetMarketProviderFee returns the MarketProviderFee field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketProviderFee() string {
	if o == nil || common.IsNil(o.MarketProviderFee) {
		var ret string
		return ret
	}
	return *o.MarketProviderFee
}

// GetMarketProviderFeeOk returns a tuple with the MarketProviderFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetMarketProviderFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketProviderFee) {
		return nil, false
	}
	return o.MarketProviderFee, true
}

// HasMarketProviderFee returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasMarketProviderFee() bool {
	if o != nil && !common.IsNil(o.MarketProviderFee) {
		return true
	}

	return false
}

// SetMarketProviderFee gets a reference to the given string and assigns it to the MarketProviderFee field.
func (o *QueryActiveOrdersResponseOrdersInner) SetMarketProviderFee(v string) {
	o.MarketProviderFee = &v
}

// GetNetworkFee returns the NetworkFee field value if set, zero value otherwise.
func (o *QueryActiveOrdersResponseOrdersInner) GetNetworkFee() string {
	if o == nil || common.IsNil(o.NetworkFee) {
		var ret string
		return ret
	}
	return *o.NetworkFee
}

// GetNetworkFeeOk returns a tuple with the NetworkFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryActiveOrdersResponseOrdersInner) GetNetworkFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkFee) {
		return nil, false
	}
	return o.NetworkFee, true
}

// HasNetworkFee returns a boolean if a field has been set.
func (o *QueryActiveOrdersResponseOrdersInner) HasNetworkFee() bool {
	if o != nil && !common.IsNil(o.NetworkFee) {
		return true
	}

	return false
}

// SetNetworkFee gets a reference to the given string and assigns it to the NetworkFee field.
func (o *QueryActiveOrdersResponseOrdersInner) SetNetworkFee(v string) {
	o.NetworkFee = &v
}

func (o QueryActiveOrdersResponseOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryActiveOrdersResponseOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.VendorOrderId) {
		toSerialize["vendorOrderId"] = o.VendorOrderId
	}
	if !common.IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !common.IsNil(o.MarketTopicId) {
		toSerialize["marketTopicId"] = o.MarketTopicId
	}
	if !common.IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !common.IsNil(o.MarketTopicTitle) {
		toSerialize["marketTopicTitle"] = o.MarketTopicTitle
	}
	if !common.IsNil(o.MarketId) {
		toSerialize["marketId"] = o.MarketId
	}
	if !common.IsNil(o.MarketTitle) {
		toSerialize["marketTitle"] = o.MarketTitle
	}
	if !common.IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !common.IsNil(o.OutcomeIndex) {
		toSerialize["outcomeIndex"] = o.OutcomeIndex
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.ModifyTime) {
		toSerialize["modifyTime"] = o.ModifyTime
	}
	if !common.IsNil(o.MakerUsdtAmount) {
		toSerialize["makerUsdtAmount"] = o.MakerUsdtAmount
	}
	if !common.IsNil(o.MakerShareQty) {
		toSerialize["makerShareQty"] = o.MakerShareQty
	}
	if !common.IsNil(o.FilledUsdtAmount) {
		toSerialize["filledUsdtAmount"] = o.FilledUsdtAmount
	}
	if !common.IsNil(o.FilledShareQty) {
		toSerialize["filledShareQty"] = o.FilledShareQty
	}
	if !common.IsNil(o.FillPercentage) {
		toSerialize["fillPercentage"] = o.FillPercentage
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.MarketProviderFee) {
		toSerialize["marketProviderFee"] = o.MarketProviderFee
	}
	if !common.IsNil(o.NetworkFee) {
		toSerialize["networkFee"] = o.NetworkFee
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryActiveOrdersResponseOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varQueryActiveOrdersResponseOrdersInner := _QueryActiveOrdersResponseOrdersInner{}

	err = json.Unmarshal(data, &varQueryActiveOrdersResponseOrdersInner)

	if err != nil {
		return err
	}

	*o = QueryActiveOrdersResponseOrdersInner(varQueryActiveOrdersResponseOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "vendorOrderId")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "marketTopicId")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "marketTopicTitle")
		delete(additionalProperties, "marketId")
		delete(additionalProperties, "marketTitle")
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "outcomeIndex")
		delete(additionalProperties, "status")
		delete(additionalProperties, "side")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "modifyTime")
		delete(additionalProperties, "makerUsdtAmount")
		delete(additionalProperties, "makerShareQty")
		delete(additionalProperties, "filledUsdtAmount")
		delete(additionalProperties, "filledShareQty")
		delete(additionalProperties, "fillPercentage")
		delete(additionalProperties, "price")
		delete(additionalProperties, "marketProviderFee")
		delete(additionalProperties, "networkFee")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryActiveOrdersResponseOrdersInner struct {
	value *QueryActiveOrdersResponseOrdersInner
	isSet bool
}

func (v NullableQueryActiveOrdersResponseOrdersInner) Get() *QueryActiveOrdersResponseOrdersInner {
	return v.value
}

func (v *NullableQueryActiveOrdersResponseOrdersInner) Set(val *QueryActiveOrdersResponseOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryActiveOrdersResponseOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryActiveOrdersResponseOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryActiveOrdersResponseOrdersInner(val *QueryActiveOrdersResponseOrdersInner) *NullableQueryActiveOrdersResponseOrdersInner {
	return &NullableQueryActiveOrdersResponseOrdersInner{value: val, isSet: true}
}

func (v NullableQueryActiveOrdersResponseOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryActiveOrdersResponseOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
