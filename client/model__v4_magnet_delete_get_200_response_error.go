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

// checks if the V4MagnetDeleteGet200ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V4MagnetDeleteGet200ResponseError{}

// V4MagnetDeleteGet200ResponseError struct for V4MagnetDeleteGet200ResponseError
type V4MagnetDeleteGet200ResponseError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewV4MagnetDeleteGet200ResponseError instantiates a new V4MagnetDeleteGet200ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4MagnetDeleteGet200ResponseError() *V4MagnetDeleteGet200ResponseError {
	this := V4MagnetDeleteGet200ResponseError{}
	return &this
}

// NewV4MagnetDeleteGet200ResponseErrorWithDefaults instantiates a new V4MagnetDeleteGet200ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4MagnetDeleteGet200ResponseErrorWithDefaults() *V4MagnetDeleteGet200ResponseError {
	this := V4MagnetDeleteGet200ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *V4MagnetDeleteGet200ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetDeleteGet200ResponseError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *V4MagnetDeleteGet200ResponseError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *V4MagnetDeleteGet200ResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V4MagnetDeleteGet200ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetDeleteGet200ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V4MagnetDeleteGet200ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V4MagnetDeleteGet200ResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o V4MagnetDeleteGet200ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V4MagnetDeleteGet200ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableV4MagnetDeleteGet200ResponseError struct {
	value *V4MagnetDeleteGet200ResponseError
	isSet bool
}

func (v NullableV4MagnetDeleteGet200ResponseError) Get() *V4MagnetDeleteGet200ResponseError {
	return v.value
}

func (v *NullableV4MagnetDeleteGet200ResponseError) Set(val *V4MagnetDeleteGet200ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableV4MagnetDeleteGet200ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableV4MagnetDeleteGet200ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4MagnetDeleteGet200ResponseError(val *V4MagnetDeleteGet200ResponseError) *NullableV4MagnetDeleteGet200ResponseError {
	return &NullableV4MagnetDeleteGet200ResponseError{value: val, isSet: true}
}

func (v NullableV4MagnetDeleteGet200ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4MagnetDeleteGet200ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
