# AccountInfo

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**Agency** | Pointer to **string** |  | [optional] 
**BankCodeForPix** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountInfo

`func NewAccountInfo() *AccountInfo`

NewAccountInfo instantiates a new AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoWithDefaults

`func NewAccountInfoWithDefaults() *AccountInfo`

NewAccountInfoWithDefaults instantiates a new AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AccountInfo) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountInfo) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountInfo) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountInfo) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAgency

`func (o *AccountInfo) GetAgency() string`

GetAgency returns the Agency field if non-nil, zero value otherwise.

### GetAgencyOk

`func (o *AccountInfo) GetAgencyOk() (*string, bool)`

GetAgencyOk returns a tuple with the Agency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgency

`func (o *AccountInfo) SetAgency(v string)`

SetAgency sets Agency field to given value.

### HasAgency

`func (o *AccountInfo) HasAgency() bool`

HasAgency returns a boolean if a field has been set.

### GetBankCodeForPix

`func (o *AccountInfo) GetBankCodeForPix() string`

GetBankCodeForPix returns the BankCodeForPix field if non-nil, zero value otherwise.

### GetBankCodeForPixOk

`func (o *AccountInfo) GetBankCodeForPixOk() (*string, bool)`

GetBankCodeForPixOk returns a tuple with the BankCodeForPix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCodeForPix

`func (o *AccountInfo) SetBankCodeForPix(v string)`

SetBankCodeForPix sets BankCodeForPix field to given value.

### HasBankCodeForPix

`func (o *AccountInfo) HasBankCodeForPix() bool`

HasBankCodeForPix returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountInfo) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountInfo) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountInfo) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountInfo) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.


[[Back to README]](../README.md)


