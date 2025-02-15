/*
Alldebrid API

Welcome to the Alldebrid API! You can use this API to access various Alldebrid services from custom applications or scripts.

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the V4MagnetRestartGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V4MagnetRestartGet200ResponseData{}

// V4MagnetRestartGet200ResponseData struct for V4MagnetRestartGet200ResponseData
type V4MagnetRestartGet200ResponseData struct {
	Magnets []V4MagnetRestartGet200ResponseDataMagnetsInner `json:"magnets,omitempty"`
}

// NewV4MagnetRestartGet200ResponseData instantiates a new V4MagnetRestartGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4MagnetRestartGet200ResponseData() *V4MagnetRestartGet200ResponseData {
	this := V4MagnetRestartGet200ResponseData{}
	return &this
}

// NewV4MagnetRestartGet200ResponseDataWithDefaults instantiates a new V4MagnetRestartGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4MagnetRestartGet200ResponseDataWithDefaults() *V4MagnetRestartGet200ResponseData {
	this := V4MagnetRestartGet200ResponseData{}
	return &this
}

// GetMagnets returns the Magnets field value if set, zero value otherwise.
func (o *V4MagnetRestartGet200ResponseData) GetMagnets() []V4MagnetRestartGet200ResponseDataMagnetsInner {
	if o == nil || IsNil(o.Magnets) {
		var ret []V4MagnetRestartGet200ResponseDataMagnetsInner
		return ret
	}
	return o.Magnets
}

// GetMagnetsOk returns a tuple with the Magnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetRestartGet200ResponseData) GetMagnetsOk() ([]V4MagnetRestartGet200ResponseDataMagnetsInner, bool) {
	if o == nil || IsNil(o.Magnets) {
		return nil, false
	}
	return o.Magnets, true
}

// HasMagnets returns a boolean if a field has been set.
func (o *V4MagnetRestartGet200ResponseData) HasMagnets() bool {
	if o != nil && !IsNil(o.Magnets) {
		return true
	}

	return false
}

// SetMagnets gets a reference to the given []V4MagnetRestartGet200ResponseDataMagnetsInner and assigns it to the Magnets field.
func (o *V4MagnetRestartGet200ResponseData) SetMagnets(v []V4MagnetRestartGet200ResponseDataMagnetsInner) {
	o.Magnets = v
}

func (o V4MagnetRestartGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V4MagnetRestartGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Magnets) {
		toSerialize["magnets"] = o.Magnets
	}
	return toSerialize, nil
}

type NullableV4MagnetRestartGet200ResponseData struct {
	value *V4MagnetRestartGet200ResponseData
	isSet bool
}

func (v NullableV4MagnetRestartGet200ResponseData) Get() *V4MagnetRestartGet200ResponseData {
	return v.value
}

func (v *NullableV4MagnetRestartGet200ResponseData) Set(val *V4MagnetRestartGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableV4MagnetRestartGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableV4MagnetRestartGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4MagnetRestartGet200ResponseData(val *V4MagnetRestartGet200ResponseData) *NullableV4MagnetRestartGet200ResponseData {
	return &NullableV4MagnetRestartGet200ResponseData{value: val, isSet: true}
}

func (v NullableV4MagnetRestartGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4MagnetRestartGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
