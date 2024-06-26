/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the HostnameRewriteBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostnameRewriteBlock{}

// HostnameRewriteBlock Hostname Rewrite grouping fields.
type HostnameRewriteBlock struct {
	// The inheritance configuration for _hostname_rewrite_char_ field.
	HostnameRewriteChar *string `json:"hostname_rewrite_char,omitempty"`
	// The inheritance configuration for _hostname_rewrite_enabled_ field.
	HostnameRewriteEnabled *bool `json:"hostname_rewrite_enabled,omitempty"`
	// The inheritance configuration for _hostname_rewrite_regex_ field.
	HostnameRewriteRegex *string `json:"hostname_rewrite_regex,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HostnameRewriteBlock HostnameRewriteBlock

// NewHostnameRewriteBlock instantiates a new HostnameRewriteBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostnameRewriteBlock() *HostnameRewriteBlock {
	this := HostnameRewriteBlock{}
	return &this
}

// NewHostnameRewriteBlockWithDefaults instantiates a new HostnameRewriteBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostnameRewriteBlockWithDefaults() *HostnameRewriteBlock {
	this := HostnameRewriteBlock{}
	return &this
}

// GetHostnameRewriteChar returns the HostnameRewriteChar field value if set, zero value otherwise.
func (o *HostnameRewriteBlock) GetHostnameRewriteChar() string {
	if o == nil || IsNil(o.HostnameRewriteChar) {
		var ret string
		return ret
	}
	return *o.HostnameRewriteChar
}

// GetHostnameRewriteCharOk returns a tuple with the HostnameRewriteChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostnameRewriteBlock) GetHostnameRewriteCharOk() (*string, bool) {
	if o == nil || IsNil(o.HostnameRewriteChar) {
		return nil, false
	}
	return o.HostnameRewriteChar, true
}

// HasHostnameRewriteChar returns a boolean if a field has been set.
func (o *HostnameRewriteBlock) HasHostnameRewriteChar() bool {
	if o != nil && !IsNil(o.HostnameRewriteChar) {
		return true
	}

	return false
}

// SetHostnameRewriteChar gets a reference to the given string and assigns it to the HostnameRewriteChar field.
func (o *HostnameRewriteBlock) SetHostnameRewriteChar(v string) {
	o.HostnameRewriteChar = &v
}

// GetHostnameRewriteEnabled returns the HostnameRewriteEnabled field value if set, zero value otherwise.
func (o *HostnameRewriteBlock) GetHostnameRewriteEnabled() bool {
	if o == nil || IsNil(o.HostnameRewriteEnabled) {
		var ret bool
		return ret
	}
	return *o.HostnameRewriteEnabled
}

// GetHostnameRewriteEnabledOk returns a tuple with the HostnameRewriteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostnameRewriteBlock) GetHostnameRewriteEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HostnameRewriteEnabled) {
		return nil, false
	}
	return o.HostnameRewriteEnabled, true
}

// HasHostnameRewriteEnabled returns a boolean if a field has been set.
func (o *HostnameRewriteBlock) HasHostnameRewriteEnabled() bool {
	if o != nil && !IsNil(o.HostnameRewriteEnabled) {
		return true
	}

	return false
}

// SetHostnameRewriteEnabled gets a reference to the given bool and assigns it to the HostnameRewriteEnabled field.
func (o *HostnameRewriteBlock) SetHostnameRewriteEnabled(v bool) {
	o.HostnameRewriteEnabled = &v
}

// GetHostnameRewriteRegex returns the HostnameRewriteRegex field value if set, zero value otherwise.
func (o *HostnameRewriteBlock) GetHostnameRewriteRegex() string {
	if o == nil || IsNil(o.HostnameRewriteRegex) {
		var ret string
		return ret
	}
	return *o.HostnameRewriteRegex
}

// GetHostnameRewriteRegexOk returns a tuple with the HostnameRewriteRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostnameRewriteBlock) GetHostnameRewriteRegexOk() (*string, bool) {
	if o == nil || IsNil(o.HostnameRewriteRegex) {
		return nil, false
	}
	return o.HostnameRewriteRegex, true
}

// HasHostnameRewriteRegex returns a boolean if a field has been set.
func (o *HostnameRewriteBlock) HasHostnameRewriteRegex() bool {
	if o != nil && !IsNil(o.HostnameRewriteRegex) {
		return true
	}

	return false
}

// SetHostnameRewriteRegex gets a reference to the given string and assigns it to the HostnameRewriteRegex field.
func (o *HostnameRewriteBlock) SetHostnameRewriteRegex(v string) {
	o.HostnameRewriteRegex = &v
}

func (o HostnameRewriteBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostnameRewriteBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostnameRewriteChar) {
		toSerialize["hostname_rewrite_char"] = o.HostnameRewriteChar
	}
	if !IsNil(o.HostnameRewriteEnabled) {
		toSerialize["hostname_rewrite_enabled"] = o.HostnameRewriteEnabled
	}
	if !IsNil(o.HostnameRewriteRegex) {
		toSerialize["hostname_rewrite_regex"] = o.HostnameRewriteRegex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HostnameRewriteBlock) UnmarshalJSON(data []byte) (err error) {
	varHostnameRewriteBlock := _HostnameRewriteBlock{}

	err = json.Unmarshal(data, &varHostnameRewriteBlock)

	if err != nil {
		return err
	}

	*o = HostnameRewriteBlock(varHostnameRewriteBlock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hostname_rewrite_char")
		delete(additionalProperties, "hostname_rewrite_enabled")
		delete(additionalProperties, "hostname_rewrite_regex")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHostnameRewriteBlock struct {
	value *HostnameRewriteBlock
	isSet bool
}

func (v NullableHostnameRewriteBlock) Get() *HostnameRewriteBlock {
	return v.value
}

func (v *NullableHostnameRewriteBlock) Set(val *HostnameRewriteBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableHostnameRewriteBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableHostnameRewriteBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostnameRewriteBlock(val *HostnameRewriteBlock) *NullableHostnameRewriteBlock {
	return &NullableHostnameRewriteBlock{value: val, isSet: true}
}

func (v NullableHostnameRewriteBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostnameRewriteBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
