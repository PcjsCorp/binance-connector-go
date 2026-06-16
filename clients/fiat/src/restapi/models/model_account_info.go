/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInfo{}

// AccountInfo struct for AccountInfo
type AccountInfo struct {
	AccountNumber        *string `json:"accountNumber,omitempty"`
	Agency               *string `json:"agency,omitempty"`
	BankCodeForPix       *string `json:"bankCodeForPix,omitempty"`
	AccountType          *string `json:"accountType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountInfo AccountInfo

// NewAccountInfo instantiates a new AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfo() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// NewAccountInfoWithDefaults instantiates a new AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoWithDefaults() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountInfo) GetAccountNumber() string {
	if o == nil || common.IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountInfo) HasAccountNumber() bool {
	if o != nil && !common.IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountInfo) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAgency returns the Agency field value if set, zero value otherwise.
func (o *AccountInfo) GetAgency() string {
	if o == nil || common.IsNil(o.Agency) {
		var ret string
		return ret
	}
	return *o.Agency
}

// GetAgencyOk returns a tuple with the Agency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAgencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Agency) {
		return nil, false
	}
	return o.Agency, true
}

// HasAgency returns a boolean if a field has been set.
func (o *AccountInfo) HasAgency() bool {
	if o != nil && !common.IsNil(o.Agency) {
		return true
	}

	return false
}

// SetAgency gets a reference to the given string and assigns it to the Agency field.
func (o *AccountInfo) SetAgency(v string) {
	o.Agency = &v
}

// GetBankCodeForPix returns the BankCodeForPix field value if set, zero value otherwise.
func (o *AccountInfo) GetBankCodeForPix() string {
	if o == nil || common.IsNil(o.BankCodeForPix) {
		var ret string
		return ret
	}
	return *o.BankCodeForPix
}

// GetBankCodeForPixOk returns a tuple with the BankCodeForPix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetBankCodeForPixOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankCodeForPix) {
		return nil, false
	}
	return o.BankCodeForPix, true
}

// HasBankCodeForPix returns a boolean if a field has been set.
func (o *AccountInfo) HasBankCodeForPix() bool {
	if o != nil && !common.IsNil(o.BankCodeForPix) {
		return true
	}

	return false
}

// SetBankCodeForPix gets a reference to the given string and assigns it to the BankCodeForPix field.
func (o *AccountInfo) SetBankCodeForPix(v string) {
	o.BankCodeForPix = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *AccountInfo) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountInfo) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *AccountInfo) SetAccountType(v string) {
	o.AccountType = &v
}

func (o AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !common.IsNil(o.Agency) {
		toSerialize["agency"] = o.Agency
	}
	if !common.IsNil(o.BankCodeForPix) {
		toSerialize["bankCodeForPix"] = o.BankCodeForPix
	}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInfo) UnmarshalJSON(data []byte) (err error) {
	varAccountInfo := _AccountInfo{}

	err = json.Unmarshal(data, &varAccountInfo)

	if err != nil {
		return err
	}

	*o = AccountInfo(varAccountInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountNumber")
		delete(additionalProperties, "agency")
		delete(additionalProperties, "bankCodeForPix")
		delete(additionalProperties, "accountType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInfo struct {
	value *AccountInfo
	isSet bool
}

func (v NullableAccountInfo) Get() *AccountInfo {
	return v.value
}

func (v *NullableAccountInfo) Set(val *AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfo(val *AccountInfo) *NullableAccountInfo {
	return &NullableAccountInfo{value: val, isSet: true}
}

func (v NullableAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
