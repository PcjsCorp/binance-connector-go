# QueryPaymentOptionBalancesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]QueryPaymentOptionBalancesResponseItemsInner**](QueryPaymentOptionBalancesResponseItemsInner.md) |  | [optional] 

## Methods

### NewQueryPaymentOptionBalancesResponse

`func NewQueryPaymentOptionBalancesResponse() *QueryPaymentOptionBalancesResponse`

NewQueryPaymentOptionBalancesResponse instantiates a new QueryPaymentOptionBalancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPaymentOptionBalancesResponseWithDefaults

`func NewQueryPaymentOptionBalancesResponseWithDefaults() *QueryPaymentOptionBalancesResponse`

NewQueryPaymentOptionBalancesResponseWithDefaults instantiates a new QueryPaymentOptionBalancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *QueryPaymentOptionBalancesResponse) GetItems() []QueryPaymentOptionBalancesResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QueryPaymentOptionBalancesResponse) GetItemsOk() (*[]QueryPaymentOptionBalancesResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QueryPaymentOptionBalancesResponse) SetItems(v []QueryPaymentOptionBalancesResponseItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *QueryPaymentOptionBalancesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to README]](../README.md)


