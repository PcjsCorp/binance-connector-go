# QueryActiveOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Orders** | Pointer to [**[]QueryActiveOrdersResponseOrdersInner**](QueryActiveOrdersResponseOrdersInner.md) |  | [optional] 

## Methods

### NewQueryActiveOrdersResponse

`func NewQueryActiveOrdersResponse() *QueryActiveOrdersResponse`

NewQueryActiveOrdersResponse instantiates a new QueryActiveOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryActiveOrdersResponseWithDefaults

`func NewQueryActiveOrdersResponseWithDefaults() *QueryActiveOrdersResponse`

NewQueryActiveOrdersResponseWithDefaults instantiates a new QueryActiveOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryActiveOrdersResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryActiveOrdersResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryActiveOrdersResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryActiveOrdersResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOffset

`func (o *QueryActiveOrdersResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryActiveOrdersResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryActiveOrdersResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryActiveOrdersResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *QueryActiveOrdersResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryActiveOrdersResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryActiveOrdersResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryActiveOrdersResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrders

`func (o *QueryActiveOrdersResponse) GetOrders() []QueryActiveOrdersResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *QueryActiveOrdersResponse) GetOrdersOk() (*[]QueryActiveOrdersResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *QueryActiveOrdersResponse) SetOrders(v []QueryActiveOrdersResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *QueryActiveOrdersResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


