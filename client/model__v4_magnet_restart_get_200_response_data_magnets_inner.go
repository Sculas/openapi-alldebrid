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

// checks if the V4MagnetRestartGet200ResponseDataMagnetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V4MagnetRestartGet200ResponseDataMagnetsInner{}

// V4MagnetRestartGet200ResponseDataMagnetsInner struct for V4MagnetRestartGet200ResponseDataMagnetsInner
type V4MagnetRestartGet200ResponseDataMagnetsInner struct {
	Magnet  *string                                             `json:"magnet,omitempty"`
	Message *string                                             `json:"message,omitempty"`
	Error   *V4MagnetRestartGet200ResponseDataMagnetsInnerError `json:"error,omitempty"`
}

// NewV4MagnetRestartGet200ResponseDataMagnetsInner instantiates a new V4MagnetRestartGet200ResponseDataMagnetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4MagnetRestartGet200ResponseDataMagnetsInner() *V4MagnetRestartGet200ResponseDataMagnetsInner {
	this := V4MagnetRestartGet200ResponseDataMagnetsInner{}
	return &this
}

// NewV4MagnetRestartGet200ResponseDataMagnetsInnerWithDefaults instantiates a new V4MagnetRestartGet200ResponseDataMagnetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4MagnetRestartGet200ResponseDataMagnetsInnerWithDefaults() *V4MagnetRestartGet200ResponseDataMagnetsInner {
	this := V4MagnetRestartGet200ResponseDataMagnetsInner{}
	return &this
}

// GetMagnet returns the Magnet field value if set, zero value otherwise.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) GetMagnet() string {
	if o == nil || IsNil(o.Magnet) {
		var ret string
		return ret
	}
	return *o.Magnet
}

// GetMagnetOk returns a tuple with the Magnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) GetMagnetOk() (*string, bool) {
	if o == nil || IsNil(o.Magnet) {
		return nil, false
	}
	return o.Magnet, true
}

// HasMagnet returns a boolean if a field has been set.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) HasMagnet() bool {
	if o != nil && !IsNil(o.Magnet) {
		return true
	}

	return false
}

// SetMagnet gets a reference to the given string and assigns it to the Magnet field.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) SetMagnet(v string) {
	o.Magnet = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) SetMessage(v string) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) GetError() V4MagnetRestartGet200ResponseDataMagnetsInnerError {
	if o == nil || IsNil(o.Error) {
		var ret V4MagnetRestartGet200ResponseDataMagnetsInnerError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) GetErrorOk() (*V4MagnetRestartGet200ResponseDataMagnetsInnerError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given V4MagnetRestartGet200ResponseDataMagnetsInnerError and assigns it to the Error field.
func (o *V4MagnetRestartGet200ResponseDataMagnetsInner) SetError(v V4MagnetRestartGet200ResponseDataMagnetsInnerError) {
	o.Error = &v
}

func (o V4MagnetRestartGet200ResponseDataMagnetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V4MagnetRestartGet200ResponseDataMagnetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Magnet) {
		toSerialize["magnet"] = o.Magnet
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableV4MagnetRestartGet200ResponseDataMagnetsInner struct {
	value *V4MagnetRestartGet200ResponseDataMagnetsInner
	isSet bool
}

func (v NullableV4MagnetRestartGet200ResponseDataMagnetsInner) Get() *V4MagnetRestartGet200ResponseDataMagnetsInner {
	return v.value
}

func (v *NullableV4MagnetRestartGet200ResponseDataMagnetsInner) Set(val *V4MagnetRestartGet200ResponseDataMagnetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV4MagnetRestartGet200ResponseDataMagnetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV4MagnetRestartGet200ResponseDataMagnetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4MagnetRestartGet200ResponseDataMagnetsInner(val *V4MagnetRestartGet200ResponseDataMagnetsInner) *NullableV4MagnetRestartGet200ResponseDataMagnetsInner {
	return &NullableV4MagnetRestartGet200ResponseDataMagnetsInner{value: val, isSet: true}
}

func (v NullableV4MagnetRestartGet200ResponseDataMagnetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4MagnetRestartGet200ResponseDataMagnetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
