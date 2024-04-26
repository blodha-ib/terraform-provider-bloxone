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

// checks if the ListACLResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListACLResponse{}

// ListACLResponse The ACL object list response format.
type ListACLResponse struct {
	// List of ACL objects.
	Results              []ACL `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListACLResponse ListACLResponse

// NewListACLResponse instantiates a new ListACLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListACLResponse() *ListACLResponse {
	this := ListACLResponse{}
	return &this
}

// NewListACLResponseWithDefaults instantiates a new ListACLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListACLResponseWithDefaults() *ListACLResponse {
	this := ListACLResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListACLResponse) GetResults() []ACL {
	if o == nil || IsNil(o.Results) {
		var ret []ACL
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListACLResponse) GetResultsOk() ([]ACL, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListACLResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ACL and assigns it to the Results field.
func (o *ListACLResponse) SetResults(v []ACL) {
	o.Results = v
}

func (o ListACLResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListACLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListACLResponse) UnmarshalJSON(data []byte) (err error) {
	varListACLResponse := _ListACLResponse{}

	err = json.Unmarshal(data, &varListACLResponse)

	if err != nil {
		return err
	}

	*o = ListACLResponse(varListACLResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListACLResponse struct {
	value *ListACLResponse
	isSet bool
}

func (v NullableListACLResponse) Get() *ListACLResponse {
	return v.value
}

func (v *NullableListACLResponse) Set(val *ListACLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListACLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListACLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListACLResponse(val *ListACLResponse) *NullableListACLResponse {
	return &NullableListACLResponse{value: val, isSet: true}
}

func (v NullableListACLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListACLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
