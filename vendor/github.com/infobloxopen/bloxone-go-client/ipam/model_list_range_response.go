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

// checks if the ListRangeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRangeResponse{}

// ListRangeResponse The response format to retrieve __Range__ objects.
type ListRangeResponse struct {
	// The list of Range objects.
	Results              []Range `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListRangeResponse ListRangeResponse

// NewListRangeResponse instantiates a new ListRangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRangeResponse() *ListRangeResponse {
	this := ListRangeResponse{}
	return &this
}

// NewListRangeResponseWithDefaults instantiates a new ListRangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRangeResponseWithDefaults() *ListRangeResponse {
	this := ListRangeResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListRangeResponse) GetResults() []Range {
	if o == nil || IsNil(o.Results) {
		var ret []Range
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRangeResponse) GetResultsOk() ([]Range, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListRangeResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Range and assigns it to the Results field.
func (o *ListRangeResponse) SetResults(v []Range) {
	o.Results = v
}

func (o ListRangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRangeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListRangeResponse) UnmarshalJSON(data []byte) (err error) {
	varListRangeResponse := _ListRangeResponse{}

	err = json.Unmarshal(data, &varListRangeResponse)

	if err != nil {
		return err
	}

	*o = ListRangeResponse(varListRangeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListRangeResponse struct {
	value *ListRangeResponse
	isSet bool
}

func (v NullableListRangeResponse) Get() *ListRangeResponse {
	return v.value
}

func (v *NullableListRangeResponse) Set(val *ListRangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRangeResponse(val *ListRangeResponse) *NullableListRangeResponse {
	return &NullableListRangeResponse{value: val, isSet: true}
}

func (v NullableListRangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
