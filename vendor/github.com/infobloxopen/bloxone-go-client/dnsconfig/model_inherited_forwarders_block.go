/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the InheritedForwardersBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InheritedForwardersBlock{}

// InheritedForwardersBlock Inheritance block for fields: _forwarders_, _forwarders_only_, _use_root_forwarders_for_local_resolution_with_b1td_.
type InheritedForwardersBlock struct {
	// Defaults to _inherit_.
	Action *string `json:"action,omitempty"`
	// Human-readable display name for the object referred to by _source_.
	DisplayName *string `json:"display_name,omitempty"`
	// The resource identifier.
	Source               *string          `json:"source,omitempty"`
	Value                *ForwardersBlock `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InheritedForwardersBlock InheritedForwardersBlock

// NewInheritedForwardersBlock instantiates a new InheritedForwardersBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInheritedForwardersBlock() *InheritedForwardersBlock {
	this := InheritedForwardersBlock{}
	return &this
}

// NewInheritedForwardersBlockWithDefaults instantiates a new InheritedForwardersBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInheritedForwardersBlockWithDefaults() *InheritedForwardersBlock {
	this := InheritedForwardersBlock{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *InheritedForwardersBlock) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InheritedForwardersBlock) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *InheritedForwardersBlock) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *InheritedForwardersBlock) SetAction(v string) {
	o.Action = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InheritedForwardersBlock) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InheritedForwardersBlock) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InheritedForwardersBlock) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InheritedForwardersBlock) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *InheritedForwardersBlock) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InheritedForwardersBlock) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *InheritedForwardersBlock) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *InheritedForwardersBlock) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InheritedForwardersBlock) GetValue() ForwardersBlock {
	if o == nil || IsNil(o.Value) {
		var ret ForwardersBlock
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InheritedForwardersBlock) GetValueOk() (*ForwardersBlock, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InheritedForwardersBlock) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ForwardersBlock and assigns it to the Value field.
func (o *InheritedForwardersBlock) SetValue(v ForwardersBlock) {
	o.Value = &v
}

func (o InheritedForwardersBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InheritedForwardersBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InheritedForwardersBlock) UnmarshalJSON(data []byte) (err error) {
	varInheritedForwardersBlock := _InheritedForwardersBlock{}

	err = json.Unmarshal(data, &varInheritedForwardersBlock)

	if err != nil {
		return err
	}

	*o = InheritedForwardersBlock(varInheritedForwardersBlock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "source")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInheritedForwardersBlock struct {
	value *InheritedForwardersBlock
	isSet bool
}

func (v NullableInheritedForwardersBlock) Get() *InheritedForwardersBlock {
	return v.value
}

func (v *NullableInheritedForwardersBlock) Set(val *InheritedForwardersBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableInheritedForwardersBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableInheritedForwardersBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInheritedForwardersBlock(val *InheritedForwardersBlock) *NullableInheritedForwardersBlock {
	return &NullableInheritedForwardersBlock{value: val, isSet: true}
}

func (v NullableInheritedForwardersBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInheritedForwardersBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
