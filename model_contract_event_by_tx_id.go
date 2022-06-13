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

// ContractEventByTxId struct for ContractEventByTxId
type ContractEventByTxId struct {
	BlockHash string `json:"blockHash"`
	ContractAddress string `json:"contractAddress"`
	EventIndex int32 `json:"eventIndex"`
	Fields []Val `json:"fields"`
}

// NewContractEventByTxId instantiates a new ContractEventByTxId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractEventByTxId(blockHash string, contractAddress string, eventIndex int32, fields []Val) *ContractEventByTxId {
	this := ContractEventByTxId{}
	this.BlockHash = blockHash
	this.ContractAddress = contractAddress
	this.EventIndex = eventIndex
	this.Fields = fields
	return &this
}

// NewContractEventByTxIdWithDefaults instantiates a new ContractEventByTxId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractEventByTxIdWithDefaults() *ContractEventByTxId {
	this := ContractEventByTxId{}
	return &this
}

// GetBlockHash returns the BlockHash field value
func (o *ContractEventByTxId) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *ContractEventByTxId) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *ContractEventByTxId) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetContractAddress returns the ContractAddress field value
func (o *ContractEventByTxId) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ContractEventByTxId) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ContractEventByTxId) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetEventIndex returns the EventIndex field value
func (o *ContractEventByTxId) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *ContractEventByTxId) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *ContractEventByTxId) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetFields returns the Fields field value
func (o *ContractEventByTxId) GetFields() []Val {
	if o == nil {
		var ret []Val
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ContractEventByTxId) GetFieldsOk() ([]Val, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *ContractEventByTxId) SetFields(v []Val) {
	o.Fields = v
}

func (o ContractEventByTxId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockHash"] = o.BlockHash
	}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if true {
		toSerialize["eventIndex"] = o.EventIndex
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableContractEventByTxId struct {
	value *ContractEventByTxId
	isSet bool
}

func (v NullableContractEventByTxId) Get() *ContractEventByTxId {
	return v.value
}

func (v *NullableContractEventByTxId) Set(val *ContractEventByTxId) {
	v.value = val
	v.isSet = true
}

func (v NullableContractEventByTxId) IsSet() bool {
	return v.isSet
}

func (v *NullableContractEventByTxId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractEventByTxId(val *ContractEventByTxId) *NullableContractEventByTxId {
	return &NullableContractEventByTxId{value: val, isSet: true}
}

func (v NullableContractEventByTxId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractEventByTxId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

