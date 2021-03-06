/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// WalletRestoreResult struct for WalletRestoreResult
type WalletRestoreResult struct {
	WalletName string `json:"walletName"`
}

// NewWalletRestoreResult instantiates a new WalletRestoreResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletRestoreResult(walletName string) *WalletRestoreResult {
	this := WalletRestoreResult{}
	this.WalletName = walletName
	return &this
}

// NewWalletRestoreResultWithDefaults instantiates a new WalletRestoreResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletRestoreResultWithDefaults() *WalletRestoreResult {
	this := WalletRestoreResult{}
	return &this
}

// GetWalletName returns the WalletName field value
func (o *WalletRestoreResult) GetWalletName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletName
}

// GetWalletNameOk returns a tuple with the WalletName field value
// and a boolean to check if the value has been set.
func (o *WalletRestoreResult) GetWalletNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletName, true
}

// SetWalletName sets field value
func (o *WalletRestoreResult) SetWalletName(v string) {
	o.WalletName = v
}

func (o WalletRestoreResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["walletName"] = o.WalletName
	}
	return json.Marshal(toSerialize)
}

type NullableWalletRestoreResult struct {
	value *WalletRestoreResult
	isSet bool
}

func (v NullableWalletRestoreResult) Get() *WalletRestoreResult {
	return v.value
}

func (v *NullableWalletRestoreResult) Set(val *WalletRestoreResult) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletRestoreResult) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletRestoreResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletRestoreResult(val *WalletRestoreResult) *NullableWalletRestoreResult {
	return &NullableWalletRestoreResult{value: val, isSet: true}
}

func (v NullableWalletRestoreResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletRestoreResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


