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

// TxNotFound struct for TxNotFound
type TxNotFound struct {
	Type string `json:"type"`
}

// NewTxNotFound instantiates a new TxNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxNotFound(type_ string) *TxNotFound {
	this := TxNotFound{}
	this.Type = type_
	return &this
}

// NewTxNotFoundWithDefaults instantiates a new TxNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxNotFoundWithDefaults() *TxNotFound {
	this := TxNotFound{}
	return &this
}

// GetType returns the Type field value
func (o *TxNotFound) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TxNotFound) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TxNotFound) SetType(v string) {
	o.Type = v
}

func (o TxNotFound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTxNotFound struct {
	value *TxNotFound
	isSet bool
}

func (v NullableTxNotFound) Get() *TxNotFound {
	return v.value
}

func (v *NullableTxNotFound) Set(val *TxNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableTxNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableTxNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxNotFound(val *TxNotFound) *NullableTxNotFound {
	return &NullableTxNotFound{value: val, isSet: true}
}

func (v NullableTxNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

