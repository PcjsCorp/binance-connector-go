# BatchRedeemResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]BatchRedeemResponseResultsInner**](BatchRedeemResponseResultsInner.md) |  | [optional] 

## Methods

### NewBatchRedeemResponse

`func NewBatchRedeemResponse() *BatchRedeemResponse`

NewBatchRedeemResponse instantiates a new BatchRedeemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRedeemResponseWithDefaults

`func NewBatchRedeemResponseWithDefaults() *BatchRedeemResponse`

NewBatchRedeemResponseWithDefaults instantiates a new BatchRedeemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *BatchRedeemResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *BatchRedeemResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *BatchRedeemResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *BatchRedeemResponse) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetResults

`func (o *BatchRedeemResponse) GetResults() []BatchRedeemResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchRedeemResponse) GetResultsOk() (*[]BatchRedeemResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchRedeemResponse) SetResults(v []BatchRedeemResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *BatchRedeemResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to README]](../README.md)


