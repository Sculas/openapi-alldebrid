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

// checks if the V4LinkInfosGet200ResponseDataInfosInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V4LinkInfosGet200ResponseDataInfosInner{}

// V4LinkInfosGet200ResponseDataInfosInner struct for V4LinkInfosGet200ResponseDataInfosInner
type V4LinkInfosGet200ResponseDataInfosInner struct {
	// Requested link
	Link *string `json:"link,omitempty"`
	// Link's file filename.
	Filename *string `json:"filename,omitempty"`
	// Link's file size in bytes.
	Size *int64 `json:"size,omitempty"`
	// Link host.
	Host *string `json:"host,omitempty"`
	// Host main domain
	HostDomain *string `json:"hostDomain,omitempty"`
}

// NewV4LinkInfosGet200ResponseDataInfosInner instantiates a new V4LinkInfosGet200ResponseDataInfosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4LinkInfosGet200ResponseDataInfosInner() *V4LinkInfosGet200ResponseDataInfosInner {
	this := V4LinkInfosGet200ResponseDataInfosInner{}
	return &this
}

// NewV4LinkInfosGet200ResponseDataInfosInnerWithDefaults instantiates a new V4LinkInfosGet200ResponseDataInfosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4LinkInfosGet200ResponseDataInfosInnerWithDefaults() *V4LinkInfosGet200ResponseDataInfosInner {
	this := V4LinkInfosGet200ResponseDataInfosInner{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *V4LinkInfosGet200ResponseDataInfosInner) SetLink(v string) {
	o.Link = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *V4LinkInfosGet200ResponseDataInfosInner) SetFilename(v string) {
	o.Filename = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *V4LinkInfosGet200ResponseDataInfosInner) SetSize(v int64) {
	o.Size = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *V4LinkInfosGet200ResponseDataInfosInner) SetHost(v string) {
	o.Host = &v
}

// GetHostDomain returns the HostDomain field value if set, zero value otherwise.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetHostDomain() string {
	if o == nil || IsNil(o.HostDomain) {
		var ret string
		return ret
	}
	return *o.HostDomain
}

// GetHostDomainOk returns a tuple with the HostDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) GetHostDomainOk() (*string, bool) {
	if o == nil || IsNil(o.HostDomain) {
		return nil, false
	}
	return o.HostDomain, true
}

// HasHostDomain returns a boolean if a field has been set.
func (o *V4LinkInfosGet200ResponseDataInfosInner) HasHostDomain() bool {
	if o != nil && !IsNil(o.HostDomain) {
		return true
	}

	return false
}

// SetHostDomain gets a reference to the given string and assigns it to the HostDomain field.
func (o *V4LinkInfosGet200ResponseDataInfosInner) SetHostDomain(v string) {
	o.HostDomain = &v
}

func (o V4LinkInfosGet200ResponseDataInfosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V4LinkInfosGet200ResponseDataInfosInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.HostDomain) {
		toSerialize["hostDomain"] = o.HostDomain
	}
	return toSerialize, nil
}

type NullableV4LinkInfosGet200ResponseDataInfosInner struct {
	value *V4LinkInfosGet200ResponseDataInfosInner
	isSet bool
}

func (v NullableV4LinkInfosGet200ResponseDataInfosInner) Get() *V4LinkInfosGet200ResponseDataInfosInner {
	return v.value
}

func (v *NullableV4LinkInfosGet200ResponseDataInfosInner) Set(val *V4LinkInfosGet200ResponseDataInfosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV4LinkInfosGet200ResponseDataInfosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV4LinkInfosGet200ResponseDataInfosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4LinkInfosGet200ResponseDataInfosInner(val *V4LinkInfosGet200ResponseDataInfosInner) *NullableV4LinkInfosGet200ResponseDataInfosInner {
	return &NullableV4LinkInfosGet200ResponseDataInfosInner{value: val, isSet: true}
}

func (v NullableV4LinkInfosGet200ResponseDataInfosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4LinkInfosGet200ResponseDataInfosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
