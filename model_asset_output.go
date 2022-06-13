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

// AssetOutput struct for AssetOutput
type AssetOutput struct {
	Hint int32 `json:"hint"`
	Key string `json:"key"`
	AttoAlphAmount string `json:"attoAlphAmount"`
	Address string `json:"address"`
	Tokens []Token `json:"tokens"`
	LockTime int64 `json:"lockTime"`
	Message string `json:"message"`
	Type string `json:"type"`
}

// NewAssetOutput instantiates a new AssetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOutput(hint int32, key string, attoAlphAmount string, address string, tokens []Token, lockTime int64, message string, type_ string) *AssetOutput {
	this := AssetOutput{}
	this.Hint = hint
	this.Key = key
	this.AttoAlphAmount = attoAlphAmount
	this.Address = address
	this.Tokens = tokens
	this.LockTime = lockTime
	this.Message = message
	this.Type = type_
	return &this
}

// NewAssetOutputWithDefaults instantiates a new AssetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOutputWithDefaults() *AssetOutput {
	this := AssetOutput{}
	return &this
}

// GetHint returns the Hint field value
func (o *AssetOutput) GetHint() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hint
}

// GetHintOk returns a tuple with the Hint field value
// and a boolean to check if the value has been set.
func (o *AssetOutput) GetHintOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hint, true
}

// SetHint sets field value
func (o *AssetOutput) SetHint(v int32) {
	o.Hint = v
}

// GetKey returns the Key field value
func (o *AssetOutput) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AssetOutput) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AssetOutput) SetKey(v string) {
	o.Key = v
}

// GetAttoAlphAmount returns the AttoAlphAmount field value
func (o *AssetOutput) GetAttoAlphAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttoAlphAmount
}

// GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field value
// and a boolean to check if the value has been set.
func (o *AssetOutput) GetAttoAlphAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttoAlphAmount, true
}

// SetAttoAlphAmount sets field value
func (o *AssetOutput) SetAttoAlphAmount(v string) {
	o.AttoAlphAmount = v
}

// GetAddress returns the Address field value
func (o *AssetOutput) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AssetOutput) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AssetOutput) SetAddress(v string) {
	o.Address = v
}

// GetTokens returns the Tokens field value
func (o *AssetOutput) GetTokens() []Token {
	if o == nil {
		var ret []Token
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *AssetOutput) GetTokensOk() ([]Token, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *AssetOutput) SetTokens(v []Token) {
	o.Tokens = v
}

// GetLockTime returns the LockTime field value
func (o *AssetOutput) GetLockTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LockTime
}

// GetLockTimeOk returns a tuple with the LockTime field value
// and a boolean to check if the value has been set.
func (o *AssetOutput) GetLockTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockTime, true
}

// SetLockTime sets field value
func (o *AssetOutput) SetLockTime(v int64) {
	o.LockTime = v
}

// GetMessage returns the Message field value
func (o *AssetOutput) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AssetOutput) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AssetOutput) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value
func (o *AssetOutput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssetOutput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AssetOutput) SetType(v string) {
	o.Type = v
}

func (o AssetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hint"] = o.Hint
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["attoAlphAmount"] = o.AttoAlphAmount
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["tokens"] = o.Tokens
	}
	if true {
		toSerialize["lockTime"] = o.LockTime
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAssetOutput struct {
	value *AssetOutput
	isSet bool
}

func (v NullableAssetOutput) Get() *AssetOutput {
	return v.value
}

func (v *NullableAssetOutput) Set(val *AssetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOutput(val *AssetOutput) *NullableAssetOutput {
	return &NullableAssetOutput{value: val, isSet: true}
}

func (v NullableAssetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

