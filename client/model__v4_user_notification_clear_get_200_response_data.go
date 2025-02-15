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

// checks if the V4UserNotificationClearGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V4UserNotificationClearGet200ResponseData{}

// V4UserNotificationClearGet200ResponseData struct for V4UserNotificationClearGet200ResponseData
type V4UserNotificationClearGet200ResponseData struct {
	Message *string `json:"message,omitempty"`
}

// NewV4UserNotificationClearGet200ResponseData instantiates a new V4UserNotificationClearGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4UserNotificationClearGet200ResponseData() *V4UserNotificationClearGet200ResponseData {
	this := V4UserNotificationClearGet200ResponseData{}
	return &this
}

// NewV4UserNotificationClearGet200ResponseDataWithDefaults instantiates a new V4UserNotificationClearGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4UserNotificationClearGet200ResponseDataWithDefaults() *V4UserNotificationClearGet200ResponseData {
	this := V4UserNotificationClearGet200ResponseData{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V4UserNotificationClearGet200ResponseData) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4UserNotificationClearGet200ResponseData) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V4UserNotificationClearGet200ResponseData) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V4UserNotificationClearGet200ResponseData) SetMessage(v string) {
	o.Message = &v
}

func (o V4UserNotificationClearGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V4UserNotificationClearGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableV4UserNotificationClearGet200ResponseData struct {
	value *V4UserNotificationClearGet200ResponseData
	isSet bool
}

func (v NullableV4UserNotificationClearGet200ResponseData) Get() *V4UserNotificationClearGet200ResponseData {
	return v.value
}

func (v *NullableV4UserNotificationClearGet200ResponseData) Set(val *V4UserNotificationClearGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableV4UserNotificationClearGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableV4UserNotificationClearGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4UserNotificationClearGet200ResponseData(val *V4UserNotificationClearGet200ResponseData) *NullableV4UserNotificationClearGet200ResponseData {
	return &NullableV4UserNotificationClearGet200ResponseData{value: val, isSet: true}
}

func (v NullableV4UserNotificationClearGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4UserNotificationClearGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
