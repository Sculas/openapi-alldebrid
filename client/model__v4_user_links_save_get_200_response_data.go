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

// checks if the V4UserLinksSaveGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V4UserLinksSaveGet200ResponseData{}

// V4UserLinksSaveGet200ResponseData struct for V4UserLinksSaveGet200ResponseData
type V4UserLinksSaveGet200ResponseData struct {
	Message *string `json:"message,omitempty"`
}

// NewV4UserLinksSaveGet200ResponseData instantiates a new V4UserLinksSaveGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4UserLinksSaveGet200ResponseData() *V4UserLinksSaveGet200ResponseData {
	this := V4UserLinksSaveGet200ResponseData{}
	return &this
}

// NewV4UserLinksSaveGet200ResponseDataWithDefaults instantiates a new V4UserLinksSaveGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4UserLinksSaveGet200ResponseDataWithDefaults() *V4UserLinksSaveGet200ResponseData {
	this := V4UserLinksSaveGet200ResponseData{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V4UserLinksSaveGet200ResponseData) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4UserLinksSaveGet200ResponseData) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V4UserLinksSaveGet200ResponseData) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V4UserLinksSaveGet200ResponseData) SetMessage(v string) {
	o.Message = &v
}

func (o V4UserLinksSaveGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V4UserLinksSaveGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableV4UserLinksSaveGet200ResponseData struct {
	value *V4UserLinksSaveGet200ResponseData
	isSet bool
}

func (v NullableV4UserLinksSaveGet200ResponseData) Get() *V4UserLinksSaveGet200ResponseData {
	return v.value
}

func (v *NullableV4UserLinksSaveGet200ResponseData) Set(val *V4UserLinksSaveGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableV4UserLinksSaveGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableV4UserLinksSaveGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4UserLinksSaveGet200ResponseData(val *V4UserLinksSaveGet200ResponseData) *NullableV4UserLinksSaveGet200ResponseData {
	return &NullableV4UserLinksSaveGet200ResponseData{value: val, isSet: true}
}

func (v NullableV4UserLinksSaveGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4UserLinksSaveGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
