/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the InternalDomainsDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalDomainsDeleteRequest{}

// InternalDomainsDeleteRequest The Internal Domains delete request.
type InternalDomainsDeleteRequest struct {
	// The list of InternalDomains object identifiers.
	Ids                  []int32 `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InternalDomainsDeleteRequest InternalDomainsDeleteRequest

// NewInternalDomainsDeleteRequest instantiates a new InternalDomainsDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalDomainsDeleteRequest() *InternalDomainsDeleteRequest {
	this := InternalDomainsDeleteRequest{}
	return &this
}

// NewInternalDomainsDeleteRequestWithDefaults instantiates a new InternalDomainsDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalDomainsDeleteRequestWithDefaults() *InternalDomainsDeleteRequest {
	this := InternalDomainsDeleteRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InternalDomainsDeleteRequest) GetIds() []int32 {
	if o == nil || IsNil(o.Ids) {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalDomainsDeleteRequest) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InternalDomainsDeleteRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *InternalDomainsDeleteRequest) SetIds(v []int32) {
	o.Ids = v
}

func (o InternalDomainsDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalDomainsDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InternalDomainsDeleteRequest) UnmarshalJSON(data []byte) (err error) {
	varInternalDomainsDeleteRequest := _InternalDomainsDeleteRequest{}

	err = json.Unmarshal(data, &varInternalDomainsDeleteRequest)

	if err != nil {
		return err
	}

	*o = InternalDomainsDeleteRequest(varInternalDomainsDeleteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternalDomainsDeleteRequest struct {
	value *InternalDomainsDeleteRequest
	isSet bool
}

func (v NullableInternalDomainsDeleteRequest) Get() *InternalDomainsDeleteRequest {
	return v.value
}

func (v *NullableInternalDomainsDeleteRequest) Set(val *InternalDomainsDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalDomainsDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalDomainsDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalDomainsDeleteRequest(val *InternalDomainsDeleteRequest) *NullableInternalDomainsDeleteRequest {
	return &NullableInternalDomainsDeleteRequest{value: val, isSet: true}
}

func (v NullableInternalDomainsDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalDomainsDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
