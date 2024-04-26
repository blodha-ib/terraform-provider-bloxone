/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
	"fmt"
)

// checks if the ForwardNSG type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForwardNSG{}

// ForwardNSG Forward DNS Server Group for forward zones.
type ForwardNSG struct {
	// Optional. Comment for the object.
	Comment *string `json:"comment,omitempty"`
	// Optional. External DNS servers to forward to. Order is not significant.
	ExternalForwarders []Forwarder `json:"external_forwarders,omitempty"`
	// Optional. _true_ to only forward.
	ForwardersOnly *bool `json:"forwarders_only,omitempty"`
	// The resource identifier.
	Hosts []string `json:"hosts,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The resource identifier.
	InternalForwarders []string `json:"internal_forwarders,omitempty"`
	// Name of the object.
	Name string `json:"name"`
	// The resource identifier.
	Nsgs []string `json:"nsgs,omitempty"`
	// Tagging specifics.
	Tags                 map[string]interface{} `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForwardNSG ForwardNSG

// NewForwardNSG instantiates a new ForwardNSG object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwardNSG(name string) *ForwardNSG {
	this := ForwardNSG{}
	this.Name = name
	return &this
}

// NewForwardNSGWithDefaults instantiates a new ForwardNSG object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwardNSGWithDefaults() *ForwardNSG {
	this := ForwardNSG{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ForwardNSG) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ForwardNSG) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ForwardNSG) SetComment(v string) {
	o.Comment = &v
}

// GetExternalForwarders returns the ExternalForwarders field value if set, zero value otherwise.
func (o *ForwardNSG) GetExternalForwarders() []Forwarder {
	if o == nil || IsNil(o.ExternalForwarders) {
		var ret []Forwarder
		return ret
	}
	return o.ExternalForwarders
}

// GetExternalForwardersOk returns a tuple with the ExternalForwarders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetExternalForwardersOk() ([]Forwarder, bool) {
	if o == nil || IsNil(o.ExternalForwarders) {
		return nil, false
	}
	return o.ExternalForwarders, true
}

// HasExternalForwarders returns a boolean if a field has been set.
func (o *ForwardNSG) HasExternalForwarders() bool {
	if o != nil && !IsNil(o.ExternalForwarders) {
		return true
	}

	return false
}

// SetExternalForwarders gets a reference to the given []Forwarder and assigns it to the ExternalForwarders field.
func (o *ForwardNSG) SetExternalForwarders(v []Forwarder) {
	o.ExternalForwarders = v
}

// GetForwardersOnly returns the ForwardersOnly field value if set, zero value otherwise.
func (o *ForwardNSG) GetForwardersOnly() bool {
	if o == nil || IsNil(o.ForwardersOnly) {
		var ret bool
		return ret
	}
	return *o.ForwardersOnly
}

// GetForwardersOnlyOk returns a tuple with the ForwardersOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetForwardersOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ForwardersOnly) {
		return nil, false
	}
	return o.ForwardersOnly, true
}

// HasForwardersOnly returns a boolean if a field has been set.
func (o *ForwardNSG) HasForwardersOnly() bool {
	if o != nil && !IsNil(o.ForwardersOnly) {
		return true
	}

	return false
}

// SetForwardersOnly gets a reference to the given bool and assigns it to the ForwardersOnly field.
func (o *ForwardNSG) SetForwardersOnly(v bool) {
	o.ForwardersOnly = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *ForwardNSG) GetHosts() []string {
	if o == nil || IsNil(o.Hosts) {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *ForwardNSG) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *ForwardNSG) SetHosts(v []string) {
	o.Hosts = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ForwardNSG) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ForwardNSG) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ForwardNSG) SetId(v string) {
	o.Id = &v
}

// GetInternalForwarders returns the InternalForwarders field value if set, zero value otherwise.
func (o *ForwardNSG) GetInternalForwarders() []string {
	if o == nil || IsNil(o.InternalForwarders) {
		var ret []string
		return ret
	}
	return o.InternalForwarders
}

// GetInternalForwardersOk returns a tuple with the InternalForwarders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetInternalForwardersOk() ([]string, bool) {
	if o == nil || IsNil(o.InternalForwarders) {
		return nil, false
	}
	return o.InternalForwarders, true
}

// HasInternalForwarders returns a boolean if a field has been set.
func (o *ForwardNSG) HasInternalForwarders() bool {
	if o != nil && !IsNil(o.InternalForwarders) {
		return true
	}

	return false
}

// SetInternalForwarders gets a reference to the given []string and assigns it to the InternalForwarders field.
func (o *ForwardNSG) SetInternalForwarders(v []string) {
	o.InternalForwarders = v
}

// GetName returns the Name field value
func (o *ForwardNSG) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ForwardNSG) SetName(v string) {
	o.Name = v
}

// GetNsgs returns the Nsgs field value if set, zero value otherwise.
func (o *ForwardNSG) GetNsgs() []string {
	if o == nil || IsNil(o.Nsgs) {
		var ret []string
		return ret
	}
	return o.Nsgs
}

// GetNsgsOk returns a tuple with the Nsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetNsgsOk() ([]string, bool) {
	if o == nil || IsNil(o.Nsgs) {
		return nil, false
	}
	return o.Nsgs, true
}

// HasNsgs returns a boolean if a field has been set.
func (o *ForwardNSG) HasNsgs() bool {
	if o != nil && !IsNil(o.Nsgs) {
		return true
	}

	return false
}

// SetNsgs gets a reference to the given []string and assigns it to the Nsgs field.
func (o *ForwardNSG) SetNsgs(v []string) {
	o.Nsgs = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ForwardNSG) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardNSG) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ForwardNSG) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ForwardNSG) SetTags(v map[string]interface{}) {
	o.Tags = v
}

func (o ForwardNSG) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForwardNSG) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ExternalForwarders) {
		toSerialize["external_forwarders"] = o.ExternalForwarders
	}
	if !IsNil(o.ForwardersOnly) {
		toSerialize["forwarders_only"] = o.ForwardersOnly
	}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InternalForwarders) {
		toSerialize["internal_forwarders"] = o.InternalForwarders
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Nsgs) {
		toSerialize["nsgs"] = o.Nsgs
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ForwardNSG) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varForwardNSG := _ForwardNSG{}

	err = json.Unmarshal(data, &varForwardNSG)

	if err != nil {
		return err
	}

	*o = ForwardNSG(varForwardNSG)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "external_forwarders")
		delete(additionalProperties, "forwarders_only")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "id")
		delete(additionalProperties, "internal_forwarders")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nsgs")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForwardNSG struct {
	value *ForwardNSG
	isSet bool
}

func (v NullableForwardNSG) Get() *ForwardNSG {
	return v.value
}

func (v *NullableForwardNSG) Set(val *ForwardNSG) {
	v.value = val
	v.isSet = true
}

func (v NullableForwardNSG) IsSet() bool {
	return v.isSet
}

func (v *NullableForwardNSG) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwardNSG(val *ForwardNSG) *NullableForwardNSG {
	return &NullableForwardNSG{value: val, isSet: true}
}

func (v NullableForwardNSG) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwardNSG) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
