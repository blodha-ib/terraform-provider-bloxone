/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
	"time"
)

// checks if the AccessCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessCode{}

// AccessCode struct for AccessCode
type AccessCode struct {
	// Auto generated unique Bypass Code value
	AccessKey *string `json:"access_key,omitempty"`
	// The time when the Bypass Code object was activated.
	Activation *time.Time `json:"activation,omitempty"`
	// The time when the Bypass Code object was created.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	Description *string    `json:"description,omitempty"`
	// The time when the Bypass Code object was expired.
	Expiration *time.Time `json:"expiration,omitempty"`
	// The name of Bypass Code
	Name *string `json:"name,omitempty"`
	// The list of SecurityPolicy object identifiers.
	PolicyIds []int32 `json:"policy_ids,omitempty"`
	// The list of selected security rules
	Rules []AccessCodeRule `json:"rules,omitempty"`
	// The time when the Bypass Code object was last updated.
	UpdatedTime          *time.Time `json:"updated_time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessCode AccessCode

// NewAccessCode instantiates a new AccessCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessCode() *AccessCode {
	this := AccessCode{}
	return &this
}

// NewAccessCodeWithDefaults instantiates a new AccessCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessCodeWithDefaults() *AccessCode {
	this := AccessCode{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *AccessCode) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *AccessCode) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *AccessCode) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetActivation returns the Activation field value if set, zero value otherwise.
func (o *AccessCode) GetActivation() time.Time {
	if o == nil || IsNil(o.Activation) {
		var ret time.Time
		return ret
	}
	return *o.Activation
}

// GetActivationOk returns a tuple with the Activation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetActivationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Activation) {
		return nil, false
	}
	return o.Activation, true
}

// HasActivation returns a boolean if a field has been set.
func (o *AccessCode) HasActivation() bool {
	if o != nil && !IsNil(o.Activation) {
		return true
	}

	return false
}

// SetActivation gets a reference to the given time.Time and assigns it to the Activation field.
func (o *AccessCode) SetActivation(v time.Time) {
	o.Activation = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *AccessCode) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *AccessCode) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *AccessCode) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessCode) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessCode) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessCode) SetDescription(v string) {
	o.Description = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *AccessCode) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *AccessCode) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *AccessCode) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessCode) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessCode) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessCode) SetName(v string) {
	o.Name = &v
}

// GetPolicyIds returns the PolicyIds field value if set, zero value otherwise.
func (o *AccessCode) GetPolicyIds() []int32 {
	if o == nil || IsNil(o.PolicyIds) {
		var ret []int32
		return ret
	}
	return o.PolicyIds
}

// GetPolicyIdsOk returns a tuple with the PolicyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetPolicyIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.PolicyIds) {
		return nil, false
	}
	return o.PolicyIds, true
}

// HasPolicyIds returns a boolean if a field has been set.
func (o *AccessCode) HasPolicyIds() bool {
	if o != nil && !IsNil(o.PolicyIds) {
		return true
	}

	return false
}

// SetPolicyIds gets a reference to the given []int32 and assigns it to the PolicyIds field.
func (o *AccessCode) SetPolicyIds(v []int32) {
	o.PolicyIds = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *AccessCode) GetRules() []AccessCodeRule {
	if o == nil || IsNil(o.Rules) {
		var ret []AccessCodeRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetRulesOk() ([]AccessCodeRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *AccessCode) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []AccessCodeRule and assigns it to the Rules field.
func (o *AccessCode) SetRules(v []AccessCodeRule) {
	o.Rules = v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *AccessCode) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCode) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *AccessCode) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *AccessCode) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

func (o AccessCode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.Activation) {
		toSerialize["activation"] = o.Activation
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PolicyIds) {
		toSerialize["policy_ids"] = o.PolicyIds
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updated_time"] = o.UpdatedTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessCode) UnmarshalJSON(data []byte) (err error) {
	varAccessCode := _AccessCode{}

	err = json.Unmarshal(data, &varAccessCode)

	if err != nil {
		return err
	}

	*o = AccessCode(varAccessCode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "activation")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "description")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "name")
		delete(additionalProperties, "policy_ids")
		delete(additionalProperties, "rules")
		delete(additionalProperties, "updated_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessCode struct {
	value *AccessCode
	isSet bool
}

func (v NullableAccessCode) Get() *AccessCode {
	return v.value
}

func (v *NullableAccessCode) Set(val *AccessCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessCode(val *AccessCode) *NullableAccessCode {
	return &NullableAccessCode{value: val, isSet: true}
}

func (v NullableAccessCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
