# QueryPositionsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**QueryPositionsResponseSummary**](QueryPositionsResponseSummary.md) |  | [optional] 
**Counts** | Pointer to [**QueryPositionsResponseCounts**](QueryPositionsResponseCounts.md) |  | [optional] 
**Positions** | Pointer to [**[]QueryPositionsResponsePositionsInner**](QueryPositionsResponsePositionsInner.md) |  | [optional] 

## Methods

### NewQueryPositionsResponse

`func NewQueryPositionsResponse() *QueryPositionsResponse`

NewQueryPositionsResponse instantiates a new QueryPositionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPositionsResponseWithDefaults

`func NewQueryPositionsResponseWithDefaults() *QueryPositionsResponse`

NewQueryPositionsResponseWithDefaults instantiates a new QueryPositionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *QueryPositionsResponse) GetSummary() QueryPositionsResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *QueryPositionsResponse) GetSummaryOk() (*QueryPositionsResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *QueryPositionsResponse) SetSummary(v QueryPositionsResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *QueryPositionsResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCounts

`func (o *QueryPositionsResponse) GetCounts() QueryPositionsResponseCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *QueryPositionsResponse) GetCountsOk() (*QueryPositionsResponseCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *QueryPositionsResponse) SetCounts(v QueryPositionsResponseCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *QueryPositionsResponse) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetPositions

`func (o *QueryPositionsResponse) GetPositions() []QueryPositionsResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *QueryPositionsResponse) GetPositionsOk() (*[]QueryPositionsResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *QueryPositionsResponse) SetPositions(v []QueryPositionsResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *QueryPositionsResponse) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


