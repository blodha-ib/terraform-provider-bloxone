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

// checks if the CopyIPSpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyIPSpaceResponse{}

// CopyIPSpaceResponse The response format to copy the __IPSpace__ object.
type CopyIPSpaceResponse struct {
	Result               *CopyResponse `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CopyIPSpaceResponse CopyIPSpaceResponse

// NewCopyIPSpaceResponse instantiates a new CopyIPSpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyIPSpaceResponse() *CopyIPSpaceResponse {
	this := CopyIPSpaceResponse{}
	return &this
}

// NewCopyIPSpaceResponseWithDefaults instantiates a new CopyIPSpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyIPSpaceResponseWithDefaults() *CopyIPSpaceResponse {
	this := CopyIPSpaceResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CopyIPSpaceResponse) GetResult() CopyResponse {
	if o == nil || IsNil(o.Result) {
		var ret CopyResponse
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyIPSpaceResponse) GetResultOk() (*CopyResponse, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CopyIPSpaceResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CopyResponse and assigns it to the Result field.
func (o *CopyIPSpaceResponse) SetResult(v CopyResponse) {
	o.Result = &v
}

func (o CopyIPSpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyIPSpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CopyIPSpaceResponse) UnmarshalJSON(data []byte) (err error) {
	varCopyIPSpaceResponse := _CopyIPSpaceResponse{}

	err = json.Unmarshal(data, &varCopyIPSpaceResponse)

	if err != nil {
		return err
	}

	*o = CopyIPSpaceResponse(varCopyIPSpaceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCopyIPSpaceResponse struct {
	value *CopyIPSpaceResponse
	isSet bool
}

func (v NullableCopyIPSpaceResponse) Get() *CopyIPSpaceResponse {
	return v.value
}

func (v *NullableCopyIPSpaceResponse) Set(val *CopyIPSpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyIPSpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyIPSpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyIPSpaceResponse(val *CopyIPSpaceResponse) *NullableCopyIPSpaceResponse {
	return &NullableCopyIPSpaceResponse{value: val, isSet: true}
}

func (v NullableCopyIPSpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyIPSpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
