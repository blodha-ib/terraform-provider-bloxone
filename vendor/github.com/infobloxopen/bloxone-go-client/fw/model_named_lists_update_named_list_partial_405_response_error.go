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

// checks if the NamedListsUpdateNamedListPartial405ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamedListsUpdateNamedListPartial405ResponseError{}

// NamedListsUpdateNamedListPartial405ResponseError struct for NamedListsUpdateNamedListPartial405ResponseError
type NamedListsUpdateNamedListPartial405ResponseError struct {
	Code                 *string `json:"code,omitempty"`
	Message              *string `json:"message,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NamedListsUpdateNamedListPartial405ResponseError NamedListsUpdateNamedListPartial405ResponseError

// NewNamedListsUpdateNamedListPartial405ResponseError instantiates a new NamedListsUpdateNamedListPartial405ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedListsUpdateNamedListPartial405ResponseError() *NamedListsUpdateNamedListPartial405ResponseError {
	this := NamedListsUpdateNamedListPartial405ResponseError{}
	return &this
}

// NewNamedListsUpdateNamedListPartial405ResponseErrorWithDefaults instantiates a new NamedListsUpdateNamedListPartial405ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedListsUpdateNamedListPartial405ResponseErrorWithDefaults() *NamedListsUpdateNamedListPartial405ResponseError {
	this := NamedListsUpdateNamedListPartial405ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *NamedListsUpdateNamedListPartial405ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedListsUpdateNamedListPartial405ResponseError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *NamedListsUpdateNamedListPartial405ResponseError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *NamedListsUpdateNamedListPartial405ResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NamedListsUpdateNamedListPartial405ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedListsUpdateNamedListPartial405ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NamedListsUpdateNamedListPartial405ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NamedListsUpdateNamedListPartial405ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NamedListsUpdateNamedListPartial405ResponseError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedListsUpdateNamedListPartial405ResponseError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NamedListsUpdateNamedListPartial405ResponseError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NamedListsUpdateNamedListPartial405ResponseError) SetStatus(v string) {
	o.Status = &v
}

func (o NamedListsUpdateNamedListPartial405ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamedListsUpdateNamedListPartial405ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NamedListsUpdateNamedListPartial405ResponseError) UnmarshalJSON(data []byte) (err error) {
	varNamedListsUpdateNamedListPartial405ResponseError := _NamedListsUpdateNamedListPartial405ResponseError{}

	err = json.Unmarshal(data, &varNamedListsUpdateNamedListPartial405ResponseError)

	if err != nil {
		return err
	}

	*o = NamedListsUpdateNamedListPartial405ResponseError(varNamedListsUpdateNamedListPartial405ResponseError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNamedListsUpdateNamedListPartial405ResponseError struct {
	value *NamedListsUpdateNamedListPartial405ResponseError
	isSet bool
}

func (v NullableNamedListsUpdateNamedListPartial405ResponseError) Get() *NamedListsUpdateNamedListPartial405ResponseError {
	return v.value
}

func (v *NullableNamedListsUpdateNamedListPartial405ResponseError) Set(val *NamedListsUpdateNamedListPartial405ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedListsUpdateNamedListPartial405ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedListsUpdateNamedListPartial405ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedListsUpdateNamedListPartial405ResponseError(val *NamedListsUpdateNamedListPartial405ResponseError) *NullableNamedListsUpdateNamedListPartial405ResponseError {
	return &NullableNamedListsUpdateNamedListPartial405ResponseError{value: val, isSet: true}
}

func (v NullableNamedListsUpdateNamedListPartial405ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedListsUpdateNamedListPartial405ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
