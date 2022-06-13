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

// BadRequest struct for BadRequest
type BadRequest struct {
	Detail string `json:"detail"`
}

// NewBadRequest instantiates a new BadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequest(detail string) *BadRequest {
	this := BadRequest{}
	this.Detail = detail
	return &this
}

// NewBadRequestWithDefaults instantiates a new BadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestWithDefaults() *BadRequest {
	this := BadRequest{}
	return &this
}

// GetDetail returns the Detail field value
func (o *BadRequest) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *BadRequest) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *BadRequest) SetDetail(v string) {
	o.Detail = v
}

func (o BadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableBadRequest struct {
	value *BadRequest
	isSet bool
}

func (v NullableBadRequest) Get() *BadRequest {
	return v.value
}

func (v *NullableBadRequest) Set(val *BadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequest(val *BadRequest) *NullableBadRequest {
	return &NullableBadRequest{value: val, isSet: true}
}

func (v NullableBadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

