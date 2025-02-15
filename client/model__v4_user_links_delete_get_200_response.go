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

// checks if the V4UserLinksDeleteGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V4UserLinksDeleteGet200Response{}

// V4UserLinksDeleteGet200Response struct for V4UserLinksDeleteGet200Response
type V4UserLinksDeleteGet200Response struct {
	Status *string                              `json:"status,omitempty"`
	Data   *V4UserLinksDeleteGet200ResponseData `json:"data,omitempty"`
	Error  *V4MagnetUploadGet200ResponseError   `json:"error,omitempty"`
}

// NewV4UserLinksDeleteGet200Response instantiates a new V4UserLinksDeleteGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4UserLinksDeleteGet200Response() *V4UserLinksDeleteGet200Response {
	this := V4UserLinksDeleteGet200Response{}
	return &this
}

// NewV4UserLinksDeleteGet200ResponseWithDefaults instantiates a new V4UserLinksDeleteGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4UserLinksDeleteGet200ResponseWithDefaults() *V4UserLinksDeleteGet200Response {
	this := V4UserLinksDeleteGet200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V4UserLinksDeleteGet200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4UserLinksDeleteGet200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V4UserLinksDeleteGet200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V4UserLinksDeleteGet200Response) SetStatus(v string) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V4UserLinksDeleteGet200Response) GetData() V4UserLinksDeleteGet200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret V4UserLinksDeleteGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4UserLinksDeleteGet200Response) GetDataOk() (*V4UserLinksDeleteGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V4UserLinksDeleteGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given V4UserLinksDeleteGet200ResponseData and assigns it to the Data field.
func (o *V4UserLinksDeleteGet200Response) SetData(v V4UserLinksDeleteGet200ResponseData) {
	o.Data = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V4UserLinksDeleteGet200Response) GetError() V4MagnetUploadGet200ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret V4MagnetUploadGet200ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4UserLinksDeleteGet200Response) GetErrorOk() (*V4MagnetUploadGet200ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V4UserLinksDeleteGet200Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given V4MagnetUploadGet200ResponseError and assigns it to the Error field.
func (o *V4UserLinksDeleteGet200Response) SetError(v V4MagnetUploadGet200ResponseError) {
	o.Error = &v
}

func (o V4UserLinksDeleteGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V4UserLinksDeleteGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableV4UserLinksDeleteGet200Response struct {
	value *V4UserLinksDeleteGet200Response
	isSet bool
}

func (v NullableV4UserLinksDeleteGet200Response) Get() *V4UserLinksDeleteGet200Response {
	return v.value
}

func (v *NullableV4UserLinksDeleteGet200Response) Set(val *V4UserLinksDeleteGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV4UserLinksDeleteGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV4UserLinksDeleteGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4UserLinksDeleteGet200Response(val *V4UserLinksDeleteGet200Response) *NullableV4UserLinksDeleteGet200Response {
	return &NullableV4UserLinksDeleteGet200Response{value: val, isSet: true}
}

func (v NullableV4UserLinksDeleteGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4UserLinksDeleteGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
