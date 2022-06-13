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

// Unban struct for Unban
type Unban struct {
	Peers []string `json:"peers"`
	Type string `json:"type"`
}

// NewUnban instantiates a new Unban object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnban(peers []string, type_ string) *Unban {
	this := Unban{}
	this.Peers = peers
	this.Type = type_
	return &this
}

// NewUnbanWithDefaults instantiates a new Unban object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnbanWithDefaults() *Unban {
	this := Unban{}
	return &this
}

// GetPeers returns the Peers field value
func (o *Unban) GetPeers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *Unban) GetPeersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *Unban) SetPeers(v []string) {
	o.Peers = v
}

// GetType returns the Type field value
func (o *Unban) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Unban) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Unban) SetType(v string) {
	o.Type = v
}

func (o Unban) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["peers"] = o.Peers
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableUnban struct {
	value *Unban
	isSet bool
}

func (v NullableUnban) Get() *Unban {
	return v.value
}

func (v *NullableUnban) Set(val *Unban) {
	v.value = val
	v.isSet = true
}

func (v NullableUnban) IsSet() bool {
	return v.isSet
}

func (v *NullableUnban) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnban(val *Unban) *NullableUnban {
	return &NullableUnban{value: val, isSet: true}
}

func (v NullableUnban) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnban) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

