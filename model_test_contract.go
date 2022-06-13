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

// TestContract struct for TestContract
type TestContract struct {
	Group *int32 `json:"group,omitempty"`
	BlockHash *string `json:"blockHash,omitempty"`
	TxId *string `json:"txId,omitempty"`
	Address *string `json:"address,omitempty"`
	Bytecode string `json:"bytecode"`
	InitialFields []Val `json:"initialFields,omitempty"`
	InitialAsset *AssetState `json:"initialAsset,omitempty"`
	MethodIndex *int32 `json:"methodIndex,omitempty"`
	Args []Val `json:"args,omitempty"`
	ExistingContracts []ContractState `json:"existingContracts,omitempty"`
	InputAssets []TestInputAsset `json:"inputAssets,omitempty"`
}

// NewTestContract instantiates a new TestContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestContract(bytecode string) *TestContract {
	this := TestContract{}
	this.Bytecode = bytecode
	return &this
}

// NewTestContractWithDefaults instantiates a new TestContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestContractWithDefaults() *TestContract {
	this := TestContract{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *TestContract) GetGroup() int32 {
	if o == nil || o.Group == nil {
		var ret int32
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetGroupOk() (*int32, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *TestContract) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given int32 and assigns it to the Group field.
func (o *TestContract) SetGroup(v int32) {
	o.Group = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *TestContract) GetBlockHash() string {
	if o == nil || o.BlockHash == nil {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetBlockHashOk() (*string, bool) {
	if o == nil || o.BlockHash == nil {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *TestContract) HasBlockHash() bool {
	if o != nil && o.BlockHash != nil {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *TestContract) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *TestContract) GetTxId() string {
	if o == nil || o.TxId == nil {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetTxIdOk() (*string, bool) {
	if o == nil || o.TxId == nil {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *TestContract) HasTxId() bool {
	if o != nil && o.TxId != nil {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *TestContract) SetTxId(v string) {
	o.TxId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TestContract) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TestContract) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TestContract) SetAddress(v string) {
	o.Address = &v
}

// GetBytecode returns the Bytecode field value
func (o *TestContract) GetBytecode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bytecode
}

// GetBytecodeOk returns a tuple with the Bytecode field value
// and a boolean to check if the value has been set.
func (o *TestContract) GetBytecodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytecode, true
}

// SetBytecode sets field value
func (o *TestContract) SetBytecode(v string) {
	o.Bytecode = v
}

// GetInitialFields returns the InitialFields field value if set, zero value otherwise.
func (o *TestContract) GetInitialFields() []Val {
	if o == nil || o.InitialFields == nil {
		var ret []Val
		return ret
	}
	return o.InitialFields
}

// GetInitialFieldsOk returns a tuple with the InitialFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetInitialFieldsOk() ([]Val, bool) {
	if o == nil || o.InitialFields == nil {
		return nil, false
	}
	return o.InitialFields, true
}

// HasInitialFields returns a boolean if a field has been set.
func (o *TestContract) HasInitialFields() bool {
	if o != nil && o.InitialFields != nil {
		return true
	}

	return false
}

// SetInitialFields gets a reference to the given []Val and assigns it to the InitialFields field.
func (o *TestContract) SetInitialFields(v []Val) {
	o.InitialFields = v
}

// GetInitialAsset returns the InitialAsset field value if set, zero value otherwise.
func (o *TestContract) GetInitialAsset() AssetState {
	if o == nil || o.InitialAsset == nil {
		var ret AssetState
		return ret
	}
	return *o.InitialAsset
}

// GetInitialAssetOk returns a tuple with the InitialAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetInitialAssetOk() (*AssetState, bool) {
	if o == nil || o.InitialAsset == nil {
		return nil, false
	}
	return o.InitialAsset, true
}

// HasInitialAsset returns a boolean if a field has been set.
func (o *TestContract) HasInitialAsset() bool {
	if o != nil && o.InitialAsset != nil {
		return true
	}

	return false
}

// SetInitialAsset gets a reference to the given AssetState and assigns it to the InitialAsset field.
func (o *TestContract) SetInitialAsset(v AssetState) {
	o.InitialAsset = &v
}

// GetMethodIndex returns the MethodIndex field value if set, zero value otherwise.
func (o *TestContract) GetMethodIndex() int32 {
	if o == nil || o.MethodIndex == nil {
		var ret int32
		return ret
	}
	return *o.MethodIndex
}

// GetMethodIndexOk returns a tuple with the MethodIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetMethodIndexOk() (*int32, bool) {
	if o == nil || o.MethodIndex == nil {
		return nil, false
	}
	return o.MethodIndex, true
}

// HasMethodIndex returns a boolean if a field has been set.
func (o *TestContract) HasMethodIndex() bool {
	if o != nil && o.MethodIndex != nil {
		return true
	}

	return false
}

// SetMethodIndex gets a reference to the given int32 and assigns it to the MethodIndex field.
func (o *TestContract) SetMethodIndex(v int32) {
	o.MethodIndex = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *TestContract) GetArgs() []Val {
	if o == nil || o.Args == nil {
		var ret []Val
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetArgsOk() ([]Val, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *TestContract) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []Val and assigns it to the Args field.
func (o *TestContract) SetArgs(v []Val) {
	o.Args = v
}

// GetExistingContracts returns the ExistingContracts field value if set, zero value otherwise.
func (o *TestContract) GetExistingContracts() []ContractState {
	if o == nil || o.ExistingContracts == nil {
		var ret []ContractState
		return ret
	}
	return o.ExistingContracts
}

// GetExistingContractsOk returns a tuple with the ExistingContracts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetExistingContractsOk() ([]ContractState, bool) {
	if o == nil || o.ExistingContracts == nil {
		return nil, false
	}
	return o.ExistingContracts, true
}

// HasExistingContracts returns a boolean if a field has been set.
func (o *TestContract) HasExistingContracts() bool {
	if o != nil && o.ExistingContracts != nil {
		return true
	}

	return false
}

// SetExistingContracts gets a reference to the given []ContractState and assigns it to the ExistingContracts field.
func (o *TestContract) SetExistingContracts(v []ContractState) {
	o.ExistingContracts = v
}

// GetInputAssets returns the InputAssets field value if set, zero value otherwise.
func (o *TestContract) GetInputAssets() []TestInputAsset {
	if o == nil || o.InputAssets == nil {
		var ret []TestInputAsset
		return ret
	}
	return o.InputAssets
}

// GetInputAssetsOk returns a tuple with the InputAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestContract) GetInputAssetsOk() ([]TestInputAsset, bool) {
	if o == nil || o.InputAssets == nil {
		return nil, false
	}
	return o.InputAssets, true
}

// HasInputAssets returns a boolean if a field has been set.
func (o *TestContract) HasInputAssets() bool {
	if o != nil && o.InputAssets != nil {
		return true
	}

	return false
}

// SetInputAssets gets a reference to the given []TestInputAsset and assigns it to the InputAssets field.
func (o *TestContract) SetInputAssets(v []TestInputAsset) {
	o.InputAssets = v
}

func (o TestContract) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.BlockHash != nil {
		toSerialize["blockHash"] = o.BlockHash
	}
	if o.TxId != nil {
		toSerialize["txId"] = o.TxId
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["bytecode"] = o.Bytecode
	}
	if o.InitialFields != nil {
		toSerialize["initialFields"] = o.InitialFields
	}
	if o.InitialAsset != nil {
		toSerialize["initialAsset"] = o.InitialAsset
	}
	if o.MethodIndex != nil {
		toSerialize["methodIndex"] = o.MethodIndex
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.ExistingContracts != nil {
		toSerialize["existingContracts"] = o.ExistingContracts
	}
	if o.InputAssets != nil {
		toSerialize["inputAssets"] = o.InputAssets
	}
	return json.Marshal(toSerialize)
}

type NullableTestContract struct {
	value *TestContract
	isSet bool
}

func (v NullableTestContract) Get() *TestContract {
	return v.value
}

func (v *NullableTestContract) Set(val *TestContract) {
	v.value = val
	v.isSet = true
}

func (v NullableTestContract) IsSet() bool {
	return v.isSet
}

func (v *NullableTestContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestContract(val *TestContract) *NullableTestContract {
	return &NullableTestContract{value: val, isSet: true}
}

func (v NullableTestContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

