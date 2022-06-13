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

// SignResult struct for SignResult
type SignResult struct {
	Signature string `json:"signature"`
}

// NewSignResult instantiates a new SignResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignResult(signature string) *SignResult {
	this := SignResult{}
	this.Signature = signature
	return &this
}

// NewSignResultWithDefaults instantiates a new SignResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignResultWithDefaults() *SignResult {
	this := SignResult{}
	return &this
}

// GetSignature returns the Signature field value
func (o *SignResult) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *SignResult) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *SignResult) SetSignature(v string) {
	o.Signature = v
}

func (o SignResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableSignResult struct {
	value *SignResult
	isSet bool
}

func (v NullableSignResult) Get() *SignResult {
	return v.value
}

func (v *NullableSignResult) Set(val *SignResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSignResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSignResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignResult(val *SignResult) *NullableSignResult {
	return &NullableSignResult{value: val, isSet: true}
}

func (v NullableSignResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

