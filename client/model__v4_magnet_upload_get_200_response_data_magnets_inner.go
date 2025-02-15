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

// checks if the V4MagnetUploadGet200ResponseDataMagnetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V4MagnetUploadGet200ResponseDataMagnetsInner{}

// V4MagnetUploadGet200ResponseDataMagnetsInner struct for V4MagnetUploadGet200ResponseDataMagnetsInner
type V4MagnetUploadGet200ResponseDataMagnetsInner struct {
	Magnet *string                                            `json:"magnet,omitempty"`
	Hash   *string                                            `json:"hash,omitempty"`
	Name   *string                                            `json:"name,omitempty"`
	Size   *int64                                             `json:"size,omitempty"`
	Ready  *bool                                              `json:"ready,omitempty"`
	Id     *int64                                             `json:"id,omitempty"`
	Error  *V4MagnetUploadGet200ResponseDataMagnetsInnerError `json:"error,omitempty"`
}

// NewV4MagnetUploadGet200ResponseDataMagnetsInner instantiates a new V4MagnetUploadGet200ResponseDataMagnetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4MagnetUploadGet200ResponseDataMagnetsInner() *V4MagnetUploadGet200ResponseDataMagnetsInner {
	this := V4MagnetUploadGet200ResponseDataMagnetsInner{}
	return &this
}

// NewV4MagnetUploadGet200ResponseDataMagnetsInnerWithDefaults instantiates a new V4MagnetUploadGet200ResponseDataMagnetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4MagnetUploadGet200ResponseDataMagnetsInnerWithDefaults() *V4MagnetUploadGet200ResponseDataMagnetsInner {
	this := V4MagnetUploadGet200ResponseDataMagnetsInner{}
	return &this
}

// GetMagnet returns the Magnet field value if set, zero value otherwise.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetMagnet() string {
	if o == nil || IsNil(o.Magnet) {
		var ret string
		return ret
	}
	return *o.Magnet
}

// GetMagnetOk returns a tuple with the Magnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetMagnetOk() (*string, bool) {
	if o == nil || IsNil(o.Magnet) {
		return nil, false
	}
	return o.Magnet, true
}

// HasMagnet returns a boolean if a field has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) HasMagnet() bool {
	if o != nil && !IsNil(o.Magnet) {
		return true
	}

	return false
}

// SetMagnet gets a reference to the given string and assigns it to the Magnet field.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) SetMagnet(v string) {
	o.Magnet = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) SetHash(v string) {
	o.Hash = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) SetSize(v int64) {
	o.Size = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetReady() bool {
	if o == nil || IsNil(o.Ready) {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.Ready) {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) HasReady() bool {
	if o != nil && !IsNil(o.Ready) {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) SetReady(v bool) {
	o.Ready = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) SetId(v int64) {
	o.Id = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetError() V4MagnetUploadGet200ResponseDataMagnetsInnerError {
	if o == nil || IsNil(o.Error) {
		var ret V4MagnetUploadGet200ResponseDataMagnetsInnerError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) GetErrorOk() (*V4MagnetUploadGet200ResponseDataMagnetsInnerError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given V4MagnetUploadGet200ResponseDataMagnetsInnerError and assigns it to the Error field.
func (o *V4MagnetUploadGet200ResponseDataMagnetsInner) SetError(v V4MagnetUploadGet200ResponseDataMagnetsInnerError) {
	o.Error = &v
}

func (o V4MagnetUploadGet200ResponseDataMagnetsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V4MagnetUploadGet200ResponseDataMagnetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Magnet) {
		toSerialize["magnet"] = o.Magnet
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Ready) {
		toSerialize["ready"] = o.Ready
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableV4MagnetUploadGet200ResponseDataMagnetsInner struct {
	value *V4MagnetUploadGet200ResponseDataMagnetsInner
	isSet bool
}

func (v NullableV4MagnetUploadGet200ResponseDataMagnetsInner) Get() *V4MagnetUploadGet200ResponseDataMagnetsInner {
	return v.value
}

func (v *NullableV4MagnetUploadGet200ResponseDataMagnetsInner) Set(val *V4MagnetUploadGet200ResponseDataMagnetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV4MagnetUploadGet200ResponseDataMagnetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV4MagnetUploadGet200ResponseDataMagnetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4MagnetUploadGet200ResponseDataMagnetsInner(val *V4MagnetUploadGet200ResponseDataMagnetsInner) *NullableV4MagnetUploadGet200ResponseDataMagnetsInner {
	return &NullableV4MagnetUploadGet200ResponseDataMagnetsInner{value: val, isSet: true}
}

func (v NullableV4MagnetUploadGet200ResponseDataMagnetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4MagnetUploadGet200ResponseDataMagnetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
