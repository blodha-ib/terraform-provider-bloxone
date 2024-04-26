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

// checks if the ListServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServerResponse{}

// ListServerResponse The response format to retrieve __Server__ objects.
type ListServerResponse struct {
	// The list of Server objects.
	Results              []Server `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListServerResponse ListServerResponse

// NewListServerResponse instantiates a new ListServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServerResponse() *ListServerResponse {
	this := ListServerResponse{}
	return &this
}

// NewListServerResponseWithDefaults instantiates a new ListServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServerResponseWithDefaults() *ListServerResponse {
	this := ListServerResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListServerResponse) GetResults() []Server {
	if o == nil || IsNil(o.Results) {
		var ret []Server
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServerResponse) GetResultsOk() ([]Server, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListServerResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Server and assigns it to the Results field.
func (o *ListServerResponse) SetResults(v []Server) {
	o.Results = v
}

func (o ListServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListServerResponse) UnmarshalJSON(data []byte) (err error) {
	varListServerResponse := _ListServerResponse{}

	err = json.Unmarshal(data, &varListServerResponse)

	if err != nil {
		return err
	}

	*o = ListServerResponse(varListServerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListServerResponse struct {
	value *ListServerResponse
	isSet bool
}

func (v NullableListServerResponse) Get() *ListServerResponse {
	return v.value
}

func (v *NullableListServerResponse) Set(val *ListServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServerResponse(val *ListServerResponse) *NullableListServerResponse {
	return &NullableListServerResponse{value: val, isSet: true}
}

func (v NullableListServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
