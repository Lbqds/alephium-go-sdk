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

// UnsignedTx struct for UnsignedTx
type UnsignedTx struct {
	TxId string `json:"txId"`
	Version int32 `json:"version"`
	NetworkId int32 `json:"networkId"`
	ScriptOpt *string `json:"scriptOpt,omitempty"`
	GasAmount int32 `json:"gasAmount"`
	GasPrice string `json:"gasPrice"`
	Inputs []AssetInput `json:"inputs"`
	FixedOutputs []FixedAssetOutput `json:"fixedOutputs"`
}

// NewUnsignedTx instantiates a new UnsignedTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsignedTx(txId string, version int32, networkId int32, gasAmount int32, gasPrice string, inputs []AssetInput, fixedOutputs []FixedAssetOutput) *UnsignedTx {
	this := UnsignedTx{}
	this.TxId = txId
	this.Version = version
	this.NetworkId = networkId
	this.GasAmount = gasAmount
	this.GasPrice = gasPrice
	this.Inputs = inputs
	this.FixedOutputs = fixedOutputs
	return &this
}

// NewUnsignedTxWithDefaults instantiates a new UnsignedTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsignedTxWithDefaults() *UnsignedTx {
	this := UnsignedTx{}
	return &this
}

// GetTxId returns the TxId field value
func (o *UnsignedTx) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *UnsignedTx) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *UnsignedTx) SetTxId(v string) {
	o.TxId = v
}

// GetVersion returns the Version field value
func (o *UnsignedTx) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UnsignedTx) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UnsignedTx) SetVersion(v int32) {
	o.Version = v
}

// GetNetworkId returns the NetworkId field value
func (o *UnsignedTx) GetNetworkId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *UnsignedTx) GetNetworkIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *UnsignedTx) SetNetworkId(v int32) {
	o.NetworkId = v
}

// GetScriptOpt returns the ScriptOpt field value if set, zero value otherwise.
func (o *UnsignedTx) GetScriptOpt() string {
	if o == nil || o.ScriptOpt == nil {
		var ret string
		return ret
	}
	return *o.ScriptOpt
}

// GetScriptOptOk returns a tuple with the ScriptOpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnsignedTx) GetScriptOptOk() (*string, bool) {
	if o == nil || o.ScriptOpt == nil {
		return nil, false
	}
	return o.ScriptOpt, true
}

// HasScriptOpt returns a boolean if a field has been set.
func (o *UnsignedTx) HasScriptOpt() bool {
	if o != nil && o.ScriptOpt != nil {
		return true
	}

	return false
}

// SetScriptOpt gets a reference to the given string and assigns it to the ScriptOpt field.
func (o *UnsignedTx) SetScriptOpt(v string) {
	o.ScriptOpt = &v
}

// GetGasAmount returns the GasAmount field value
func (o *UnsignedTx) GetGasAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasAmount
}

// GetGasAmountOk returns a tuple with the GasAmount field value
// and a boolean to check if the value has been set.
func (o *UnsignedTx) GetGasAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasAmount, true
}

// SetGasAmount sets field value
func (o *UnsignedTx) SetGasAmount(v int32) {
	o.GasAmount = v
}

// GetGasPrice returns the GasPrice field value
func (o *UnsignedTx) GetGasPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *UnsignedTx) GetGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *UnsignedTx) SetGasPrice(v string) {
	o.GasPrice = v
}

// GetInputs returns the Inputs field value
func (o *UnsignedTx) GetInputs() []AssetInput {
	if o == nil {
		var ret []AssetInput
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *UnsignedTx) GetInputsOk() ([]AssetInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *UnsignedTx) SetInputs(v []AssetInput) {
	o.Inputs = v
}

// GetFixedOutputs returns the FixedOutputs field value
func (o *UnsignedTx) GetFixedOutputs() []FixedAssetOutput {
	if o == nil {
		var ret []FixedAssetOutput
		return ret
	}

	return o.FixedOutputs
}

// GetFixedOutputsOk returns a tuple with the FixedOutputs field value
// and a boolean to check if the value has been set.
func (o *UnsignedTx) GetFixedOutputsOk() ([]FixedAssetOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.FixedOutputs, true
}

// SetFixedOutputs sets field value
func (o *UnsignedTx) SetFixedOutputs(v []FixedAssetOutput) {
	o.FixedOutputs = v
}

func (o UnsignedTx) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["txId"] = o.TxId
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["networkId"] = o.NetworkId
	}
	if o.ScriptOpt != nil {
		toSerialize["scriptOpt"] = o.ScriptOpt
	}
	if true {
		toSerialize["gasAmount"] = o.GasAmount
	}
	if true {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	if true {
		toSerialize["fixedOutputs"] = o.FixedOutputs
	}
	return json.Marshal(toSerialize)
}

type NullableUnsignedTx struct {
	value *UnsignedTx
	isSet bool
}

func (v NullableUnsignedTx) Get() *UnsignedTx {
	return v.value
}

func (v *NullableUnsignedTx) Set(val *UnsignedTx) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsignedTx) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsignedTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsignedTx(val *UnsignedTx) *NullableUnsignedTx {
	return &NullableUnsignedTx{value: val, isSet: true}
}

func (v NullableUnsignedTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsignedTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


