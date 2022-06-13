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

// BuildExecuteScriptTxResult struct for BuildExecuteScriptTxResult
type BuildExecuteScriptTxResult struct {
	FromGroup int32 `json:"fromGroup"`
	ToGroup int32 `json:"toGroup"`
	UnsignedTx string `json:"unsignedTx"`
	GasAmount int32 `json:"gasAmount"`
	GasPrice string `json:"gasPrice"`
	TxId string `json:"txId"`
}

// NewBuildExecuteScriptTxResult instantiates a new BuildExecuteScriptTxResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildExecuteScriptTxResult(fromGroup int32, toGroup int32, unsignedTx string, gasAmount int32, gasPrice string, txId string) *BuildExecuteScriptTxResult {
	this := BuildExecuteScriptTxResult{}
	this.FromGroup = fromGroup
	this.ToGroup = toGroup
	this.UnsignedTx = unsignedTx
	this.GasAmount = gasAmount
	this.GasPrice = gasPrice
	this.TxId = txId
	return &this
}

// NewBuildExecuteScriptTxResultWithDefaults instantiates a new BuildExecuteScriptTxResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildExecuteScriptTxResultWithDefaults() *BuildExecuteScriptTxResult {
	this := BuildExecuteScriptTxResult{}
	return &this
}

// GetFromGroup returns the FromGroup field value
func (o *BuildExecuteScriptTxResult) GetFromGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromGroup
}

// GetFromGroupOk returns a tuple with the FromGroup field value
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTxResult) GetFromGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromGroup, true
}

// SetFromGroup sets field value
func (o *BuildExecuteScriptTxResult) SetFromGroup(v int32) {
	o.FromGroup = v
}

// GetToGroup returns the ToGroup field value
func (o *BuildExecuteScriptTxResult) GetToGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToGroup
}

// GetToGroupOk returns a tuple with the ToGroup field value
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTxResult) GetToGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToGroup, true
}

// SetToGroup sets field value
func (o *BuildExecuteScriptTxResult) SetToGroup(v int32) {
	o.ToGroup = v
}

// GetUnsignedTx returns the UnsignedTx field value
func (o *BuildExecuteScriptTxResult) GetUnsignedTx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnsignedTx
}

// GetUnsignedTxOk returns a tuple with the UnsignedTx field value
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTxResult) GetUnsignedTxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnsignedTx, true
}

// SetUnsignedTx sets field value
func (o *BuildExecuteScriptTxResult) SetUnsignedTx(v string) {
	o.UnsignedTx = v
}

// GetGasAmount returns the GasAmount field value
func (o *BuildExecuteScriptTxResult) GetGasAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasAmount
}

// GetGasAmountOk returns a tuple with the GasAmount field value
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTxResult) GetGasAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasAmount, true
}

// SetGasAmount sets field value
func (o *BuildExecuteScriptTxResult) SetGasAmount(v int32) {
	o.GasAmount = v
}

// GetGasPrice returns the GasPrice field value
func (o *BuildExecuteScriptTxResult) GetGasPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTxResult) GetGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *BuildExecuteScriptTxResult) SetGasPrice(v string) {
	o.GasPrice = v
}

// GetTxId returns the TxId field value
func (o *BuildExecuteScriptTxResult) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTxResult) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *BuildExecuteScriptTxResult) SetTxId(v string) {
	o.TxId = v
}

func (o BuildExecuteScriptTxResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fromGroup"] = o.FromGroup
	}
	if true {
		toSerialize["toGroup"] = o.ToGroup
	}
	if true {
		toSerialize["unsignedTx"] = o.UnsignedTx
	}
	if true {
		toSerialize["gasAmount"] = o.GasAmount
	}
	if true {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if true {
		toSerialize["txId"] = o.TxId
	}
	return json.Marshal(toSerialize)
}

type NullableBuildExecuteScriptTxResult struct {
	value *BuildExecuteScriptTxResult
	isSet bool
}

func (v NullableBuildExecuteScriptTxResult) Get() *BuildExecuteScriptTxResult {
	return v.value
}

func (v *NullableBuildExecuteScriptTxResult) Set(val *BuildExecuteScriptTxResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildExecuteScriptTxResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildExecuteScriptTxResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildExecuteScriptTxResult(val *BuildExecuteScriptTxResult) *NullableBuildExecuteScriptTxResult {
	return &NullableBuildExecuteScriptTxResult{value: val, isSet: true}
}

func (v NullableBuildExecuteScriptTxResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildExecuteScriptTxResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

