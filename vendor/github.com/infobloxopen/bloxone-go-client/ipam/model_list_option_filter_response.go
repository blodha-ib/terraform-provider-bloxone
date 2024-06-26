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

// checks if the ListOptionFilterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOptionFilterResponse{}

// ListOptionFilterResponse The response format to retrieve __OptionFilter__ objects.
type ListOptionFilterResponse struct {
	// The list of OptionFilter objects.
	Results              []OptionFilter `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListOptionFilterResponse ListOptionFilterResponse

// NewListOptionFilterResponse instantiates a new ListOptionFilterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOptionFilterResponse() *ListOptionFilterResponse {
	this := ListOptionFilterResponse{}
	return &this
}

// NewListOptionFilterResponseWithDefaults instantiates a new ListOptionFilterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOptionFilterResponseWithDefaults() *ListOptionFilterResponse {
	this := ListOptionFilterResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListOptionFilterResponse) GetResults() []OptionFilter {
	if o == nil || IsNil(o.Results) {
		var ret []OptionFilter
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOptionFilterResponse) GetResultsOk() ([]OptionFilter, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListOptionFilterResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OptionFilter and assigns it to the Results field.
func (o *ListOptionFilterResponse) SetResults(v []OptionFilter) {
	o.Results = v
}

func (o ListOptionFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOptionFilterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListOptionFilterResponse) UnmarshalJSON(data []byte) (err error) {
	varListOptionFilterResponse := _ListOptionFilterResponse{}

	err = json.Unmarshal(data, &varListOptionFilterResponse)

	if err != nil {
		return err
	}

	*o = ListOptionFilterResponse(varListOptionFilterResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListOptionFilterResponse struct {
	value *ListOptionFilterResponse
	isSet bool
}

func (v NullableListOptionFilterResponse) Get() *ListOptionFilterResponse {
	return v.value
}

func (v *NullableListOptionFilterResponse) Set(val *ListOptionFilterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOptionFilterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOptionFilterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOptionFilterResponse(val *ListOptionFilterResponse) *NullableListOptionFilterResponse {
	return &NullableListOptionFilterResponse{value: val, isSet: true}
}

func (v NullableListOptionFilterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOptionFilterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
