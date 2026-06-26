# GetPortfolioResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **string** |  | [optional] 
**WalletAddress** | Pointer to **string** |  | [optional] 
**ActivePositionsCount** | Pointer to **int32** |  | [optional] 
**TotalRealizedPnl** | Pointer to **string** |  | [optional] 
**TotalUnrealizedPnl** | Pointer to **string** |  | [optional] 
**TotalPnl** | Pointer to **string** |  | [optional] 
**TotalCostBasis** | Pointer to **string** |  | [optional] 
**TotalCurrentValue** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to [**[]GetPortfolioResponsePositionsInner**](GetPortfolioResponsePositionsInner.md) |  | [optional] 

## Methods

### NewGetPortfolioResponse

`func NewGetPortfolioResponse() *GetPortfolioResponse`

NewGetPortfolioResponse instantiates a new GetPortfolioResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPortfolioResponseWithDefaults

`func NewGetPortfolioResponseWithDefaults() *GetPortfolioResponse`

NewGetPortfolioResponseWithDefaults instantiates a new GetPortfolioResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *GetPortfolioResponse) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GetPortfolioResponse) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GetPortfolioResponse) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *GetPortfolioResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetWalletAddress

`func (o *GetPortfolioResponse) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *GetPortfolioResponse) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *GetPortfolioResponse) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *GetPortfolioResponse) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetActivePositionsCount

`func (o *GetPortfolioResponse) GetActivePositionsCount() int32`

GetActivePositionsCount returns the ActivePositionsCount field if non-nil, zero value otherwise.

### GetActivePositionsCountOk

`func (o *GetPortfolioResponse) GetActivePositionsCountOk() (*int32, bool)`

GetActivePositionsCountOk returns a tuple with the ActivePositionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePositionsCount

`func (o *GetPortfolioResponse) SetActivePositionsCount(v int32)`

SetActivePositionsCount sets ActivePositionsCount field to given value.

### HasActivePositionsCount

`func (o *GetPortfolioResponse) HasActivePositionsCount() bool`

HasActivePositionsCount returns a boolean if a field has been set.

### GetTotalRealizedPnl

`func (o *GetPortfolioResponse) GetTotalRealizedPnl() string`

GetTotalRealizedPnl returns the TotalRealizedPnl field if non-nil, zero value otherwise.

### GetTotalRealizedPnlOk

`func (o *GetPortfolioResponse) GetTotalRealizedPnlOk() (*string, bool)`

GetTotalRealizedPnlOk returns a tuple with the TotalRealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRealizedPnl

`func (o *GetPortfolioResponse) SetTotalRealizedPnl(v string)`

SetTotalRealizedPnl sets TotalRealizedPnl field to given value.

### HasTotalRealizedPnl

`func (o *GetPortfolioResponse) HasTotalRealizedPnl() bool`

HasTotalRealizedPnl returns a boolean if a field has been set.

### GetTotalUnrealizedPnl

`func (o *GetPortfolioResponse) GetTotalUnrealizedPnl() string`

GetTotalUnrealizedPnl returns the TotalUnrealizedPnl field if non-nil, zero value otherwise.

### GetTotalUnrealizedPnlOk

`func (o *GetPortfolioResponse) GetTotalUnrealizedPnlOk() (*string, bool)`

GetTotalUnrealizedPnlOk returns a tuple with the TotalUnrealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedPnl

`func (o *GetPortfolioResponse) SetTotalUnrealizedPnl(v string)`

SetTotalUnrealizedPnl sets TotalUnrealizedPnl field to given value.

### HasTotalUnrealizedPnl

`func (o *GetPortfolioResponse) HasTotalUnrealizedPnl() bool`

HasTotalUnrealizedPnl returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetPortfolioResponse) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetPortfolioResponse) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetPortfolioResponse) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetPortfolioResponse) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTotalCostBasis

`func (o *GetPortfolioResponse) GetTotalCostBasis() string`

GetTotalCostBasis returns the TotalCostBasis field if non-nil, zero value otherwise.

### GetTotalCostBasisOk

`func (o *GetPortfolioResponse) GetTotalCostBasisOk() (*string, bool)`

GetTotalCostBasisOk returns a tuple with the TotalCostBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostBasis

`func (o *GetPortfolioResponse) SetTotalCostBasis(v string)`

SetTotalCostBasis sets TotalCostBasis field to given value.

### HasTotalCostBasis

`func (o *GetPortfolioResponse) HasTotalCostBasis() bool`

HasTotalCostBasis returns a boolean if a field has been set.

### GetTotalCurrentValue

`func (o *GetPortfolioResponse) GetTotalCurrentValue() string`

GetTotalCurrentValue returns the TotalCurrentValue field if non-nil, zero value otherwise.

### GetTotalCurrentValueOk

`func (o *GetPortfolioResponse) GetTotalCurrentValueOk() (*string, bool)`

GetTotalCurrentValueOk returns a tuple with the TotalCurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrentValue

`func (o *GetPortfolioResponse) SetTotalCurrentValue(v string)`

SetTotalCurrentValue sets TotalCurrentValue field to given value.

### HasTotalCurrentValue

`func (o *GetPortfolioResponse) HasTotalCurrentValue() bool`

HasTotalCurrentValue returns a boolean if a field has been set.

### GetPositions

`func (o *GetPortfolioResponse) GetPositions() []GetPortfolioResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetPortfolioResponse) GetPositionsOk() (*[]GetPortfolioResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetPortfolioResponse) SetPositions(v []GetPortfolioResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetPortfolioResponse) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


