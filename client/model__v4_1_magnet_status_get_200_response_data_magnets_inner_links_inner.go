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

// checks if the V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner{}

// V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner struct for V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner
type V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner struct {
	Link     *string `json:"link,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Size     *int64  `json:"size,omitempty"`
}

// NewV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner instantiates a new V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner() *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner {
	this := V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner{}
	return &this
}

// NewV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInnerWithDefaults instantiates a new V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInnerWithDefaults() *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner {
	this := V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) SetLink(v string) {
	o.Link = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) SetFilename(v string) {
	o.Filename = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) SetSize(v int64) {
	o.Size = &v
}

func (o V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner struct {
	value *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner
	isSet bool
}

func (v NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) Get() *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner {
	return v.value
}

func (v *NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) Set(val *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner(val *V41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) *NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner {
	return &NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner{value: val, isSet: true}
}

func (v NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV41MagnetStatusGet200ResponseDataMagnetsInnerLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
