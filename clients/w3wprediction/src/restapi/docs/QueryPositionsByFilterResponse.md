# QueryPositionsByFilterResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Positions** | Pointer to [**[]QueryPositionsByFilterResponsePositionsInner**](QueryPositionsByFilterResponsePositionsInner.md) |  | [optional] 

## Methods

### NewQueryPositionsByFilterResponse

`func NewQueryPositionsByFilterResponse() *QueryPositionsByFilterResponse`

NewQueryPositionsByFilterResponse instantiates a new QueryPositionsByFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPositionsByFilterResponseWithDefaults

`func NewQueryPositionsByFilterResponseWithDefaults() *QueryPositionsByFilterResponse`

NewQueryPositionsByFilterResponseWithDefaults instantiates a new QueryPositionsByFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositions

`func (o *QueryPositionsByFilterResponse) GetPositions() []QueryPositionsByFilterResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *QueryPositionsByFilterResponse) GetPositionsOk() (*[]QueryPositionsByFilterResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *QueryPositionsByFilterResponse) SetPositions(v []QueryPositionsByFilterResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *QueryPositionsByFilterResponse) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


