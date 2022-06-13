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

// ChainParams struct for ChainParams
type ChainParams struct {
	NetworkId int32 `json:"networkId"`
	NumZerosAtLeastInHash int32 `json:"numZerosAtLeastInHash"`
	GroupNumPerBroker int32 `json:"groupNumPerBroker"`
	Groups int32 `json:"groups"`
}

// NewChainParams instantiates a new ChainParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChainParams(networkId int32, numZerosAtLeastInHash int32, groupNumPerBroker int32, groups int32) *ChainParams {
	this := ChainParams{}
	this.NetworkId = networkId
	this.NumZerosAtLeastInHash = numZerosAtLeastInHash
	this.GroupNumPerBroker = groupNumPerBroker
	this.Groups = groups
	return &this
}

// NewChainParamsWithDefaults instantiates a new ChainParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChainParamsWithDefaults() *ChainParams {
	this := ChainParams{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *ChainParams) GetNetworkId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *ChainParams) GetNetworkIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *ChainParams) SetNetworkId(v int32) {
	o.NetworkId = v
}

// GetNumZerosAtLeastInHash returns the NumZerosAtLeastInHash field value
func (o *ChainParams) GetNumZerosAtLeastInHash() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumZerosAtLeastInHash
}

// GetNumZerosAtLeastInHashOk returns a tuple with the NumZerosAtLeastInHash field value
// and a boolean to check if the value has been set.
func (o *ChainParams) GetNumZerosAtLeastInHashOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumZerosAtLeastInHash, true
}

// SetNumZerosAtLeastInHash sets field value
func (o *ChainParams) SetNumZerosAtLeastInHash(v int32) {
	o.NumZerosAtLeastInHash = v
}

// GetGroupNumPerBroker returns the GroupNumPerBroker field value
func (o *ChainParams) GetGroupNumPerBroker() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupNumPerBroker
}

// GetGroupNumPerBrokerOk returns a tuple with the GroupNumPerBroker field value
// and a boolean to check if the value has been set.
func (o *ChainParams) GetGroupNumPerBrokerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupNumPerBroker, true
}

// SetGroupNumPerBroker sets field value
func (o *ChainParams) SetGroupNumPerBroker(v int32) {
	o.GroupNumPerBroker = v
}

// GetGroups returns the Groups field value
func (o *ChainParams) GetGroups() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *ChainParams) GetGroupsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value
func (o *ChainParams) SetGroups(v int32) {
	o.Groups = v
}

func (o ChainParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["networkId"] = o.NetworkId
	}
	if true {
		toSerialize["numZerosAtLeastInHash"] = o.NumZerosAtLeastInHash
	}
	if true {
		toSerialize["groupNumPerBroker"] = o.GroupNumPerBroker
	}
	if true {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableChainParams struct {
	value *ChainParams
	isSet bool
}

func (v NullableChainParams) Get() *ChainParams {
	return v.value
}

func (v *NullableChainParams) Set(val *ChainParams) {
	v.value = val
	v.isSet = true
}

func (v NullableChainParams) IsSet() bool {
	return v.isSet
}

func (v *NullableChainParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChainParams(val *ChainParams) *NullableChainParams {
	return &NullableChainParams{value: val, isSet: true}
}

func (v NullableChainParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChainParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


