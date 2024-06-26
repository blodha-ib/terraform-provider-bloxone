/*
Host Activation Service

Host activation service provides a RESTful interface to manage cert and join token object. Join tokens are essentially a password that allows on-prem hosts to auto-associate themselves to a customer's account and receive a signed cert.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infraprovision

import (
	"encoding/json"
	"time"
)

// checks if the CSR type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CSR{}

// CSR Represents a certificate signing request from an on-prem host.
type CSR struct {
	ActivationCode *string         `json:"activation_code,omitempty"`
	ClientIp       *TypesInetValue `json:"client_ip,omitempty"`
	CreatedAt      *time.Time      `json:"created_at,omitempty"`
	Csr            *string         `json:"csr,omitempty"`
	HostSerial     *string         `json:"host_serial,omitempty"`
	// The resource identifier.
	Id                   *string         `json:"id,omitempty"`
	JoinToken            *JoinToken      `json:"join_token,omitempty"`
	Ophid                *string         `json:"ophid,omitempty"`
	Signature            *string         `json:"signature,omitempty"`
	SrcIp                *TypesInetValue `json:"src_ip,omitempty"`
	State                *CSRState       `json:"state,omitempty"`
	UpdatedAt            *time.Time      `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CSR CSR

// NewCSR instantiates a new CSR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSR() *CSR {
	this := CSR{}
	var state CSRState = CSRSTATE_UNKNOWN
	this.State = &state
	return &this
}

// NewCSRWithDefaults instantiates a new CSR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSRWithDefaults() *CSR {
	this := CSR{}
	var state CSRState = CSRSTATE_UNKNOWN
	this.State = &state
	return &this
}

// GetActivationCode returns the ActivationCode field value if set, zero value otherwise.
func (o *CSR) GetActivationCode() string {
	if o == nil || IsNil(o.ActivationCode) {
		var ret string
		return ret
	}
	return *o.ActivationCode
}

// GetActivationCodeOk returns a tuple with the ActivationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetActivationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ActivationCode) {
		return nil, false
	}
	return o.ActivationCode, true
}

// HasActivationCode returns a boolean if a field has been set.
func (o *CSR) HasActivationCode() bool {
	if o != nil && !IsNil(o.ActivationCode) {
		return true
	}

	return false
}

// SetActivationCode gets a reference to the given string and assigns it to the ActivationCode field.
func (o *CSR) SetActivationCode(v string) {
	o.ActivationCode = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *CSR) GetClientIp() TypesInetValue {
	if o == nil || IsNil(o.ClientIp) {
		var ret TypesInetValue
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetClientIpOk() (*TypesInetValue, bool) {
	if o == nil || IsNil(o.ClientIp) {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *CSR) HasClientIp() bool {
	if o != nil && !IsNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given TypesInetValue and assigns it to the ClientIp field.
func (o *CSR) SetClientIp(v TypesInetValue) {
	o.ClientIp = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CSR) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CSR) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CSR) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCsr returns the Csr field value if set, zero value otherwise.
func (o *CSR) GetCsr() string {
	if o == nil || IsNil(o.Csr) {
		var ret string
		return ret
	}
	return *o.Csr
}

// GetCsrOk returns a tuple with the Csr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetCsrOk() (*string, bool) {
	if o == nil || IsNil(o.Csr) {
		return nil, false
	}
	return o.Csr, true
}

// HasCsr returns a boolean if a field has been set.
func (o *CSR) HasCsr() bool {
	if o != nil && !IsNil(o.Csr) {
		return true
	}

	return false
}

// SetCsr gets a reference to the given string and assigns it to the Csr field.
func (o *CSR) SetCsr(v string) {
	o.Csr = &v
}

// GetHostSerial returns the HostSerial field value if set, zero value otherwise.
func (o *CSR) GetHostSerial() string {
	if o == nil || IsNil(o.HostSerial) {
		var ret string
		return ret
	}
	return *o.HostSerial
}

// GetHostSerialOk returns a tuple with the HostSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetHostSerialOk() (*string, bool) {
	if o == nil || IsNil(o.HostSerial) {
		return nil, false
	}
	return o.HostSerial, true
}

// HasHostSerial returns a boolean if a field has been set.
func (o *CSR) HasHostSerial() bool {
	if o != nil && !IsNil(o.HostSerial) {
		return true
	}

	return false
}

// SetHostSerial gets a reference to the given string and assigns it to the HostSerial field.
func (o *CSR) SetHostSerial(v string) {
	o.HostSerial = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CSR) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CSR) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CSR) SetId(v string) {
	o.Id = &v
}

// GetJoinToken returns the JoinToken field value if set, zero value otherwise.
func (o *CSR) GetJoinToken() JoinToken {
	if o == nil || IsNil(o.JoinToken) {
		var ret JoinToken
		return ret
	}
	return *o.JoinToken
}

// GetJoinTokenOk returns a tuple with the JoinToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetJoinTokenOk() (*JoinToken, bool) {
	if o == nil || IsNil(o.JoinToken) {
		return nil, false
	}
	return o.JoinToken, true
}

// HasJoinToken returns a boolean if a field has been set.
func (o *CSR) HasJoinToken() bool {
	if o != nil && !IsNil(o.JoinToken) {
		return true
	}

	return false
}

// SetJoinToken gets a reference to the given JoinToken and assigns it to the JoinToken field.
func (o *CSR) SetJoinToken(v JoinToken) {
	o.JoinToken = &v
}

// GetOphid returns the Ophid field value if set, zero value otherwise.
func (o *CSR) GetOphid() string {
	if o == nil || IsNil(o.Ophid) {
		var ret string
		return ret
	}
	return *o.Ophid
}

// GetOphidOk returns a tuple with the Ophid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetOphidOk() (*string, bool) {
	if o == nil || IsNil(o.Ophid) {
		return nil, false
	}
	return o.Ophid, true
}

// HasOphid returns a boolean if a field has been set.
func (o *CSR) HasOphid() bool {
	if o != nil && !IsNil(o.Ophid) {
		return true
	}

	return false
}

// SetOphid gets a reference to the given string and assigns it to the Ophid field.
func (o *CSR) SetOphid(v string) {
	o.Ophid = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *CSR) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *CSR) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *CSR) SetSignature(v string) {
	o.Signature = &v
}

// GetSrcIp returns the SrcIp field value if set, zero value otherwise.
func (o *CSR) GetSrcIp() TypesInetValue {
	if o == nil || IsNil(o.SrcIp) {
		var ret TypesInetValue
		return ret
	}
	return *o.SrcIp
}

// GetSrcIpOk returns a tuple with the SrcIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetSrcIpOk() (*TypesInetValue, bool) {
	if o == nil || IsNil(o.SrcIp) {
		return nil, false
	}
	return o.SrcIp, true
}

// HasSrcIp returns a boolean if a field has been set.
func (o *CSR) HasSrcIp() bool {
	if o != nil && !IsNil(o.SrcIp) {
		return true
	}

	return false
}

// SetSrcIp gets a reference to the given TypesInetValue and assigns it to the SrcIp field.
func (o *CSR) SetSrcIp(v TypesInetValue) {
	o.SrcIp = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CSR) GetState() CSRState {
	if o == nil || IsNil(o.State) {
		var ret CSRState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetStateOk() (*CSRState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CSR) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CSRState and assigns it to the State field.
func (o *CSR) SetState(v CSRState) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CSR) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSR) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CSR) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CSR) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CSR) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CSR) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivationCode) {
		toSerialize["activation_code"] = o.ActivationCode
	}
	if !IsNil(o.ClientIp) {
		toSerialize["client_ip"] = o.ClientIp
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Csr) {
		toSerialize["csr"] = o.Csr
	}
	if !IsNil(o.HostSerial) {
		toSerialize["host_serial"] = o.HostSerial
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.JoinToken) {
		toSerialize["join_token"] = o.JoinToken
	}
	if !IsNil(o.Ophid) {
		toSerialize["ophid"] = o.Ophid
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.SrcIp) {
		toSerialize["src_ip"] = o.SrcIp
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CSR) UnmarshalJSON(data []byte) (err error) {
	varCSR := _CSR{}

	err = json.Unmarshal(data, &varCSR)

	if err != nil {
		return err
	}

	*o = CSR(varCSR)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activation_code")
		delete(additionalProperties, "client_ip")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "csr")
		delete(additionalProperties, "host_serial")
		delete(additionalProperties, "id")
		delete(additionalProperties, "join_token")
		delete(additionalProperties, "ophid")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "src_ip")
		delete(additionalProperties, "state")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCSR struct {
	value *CSR
	isSet bool
}

func (v NullableCSR) Get() *CSR {
	return v.value
}

func (v *NullableCSR) Set(val *CSR) {
	v.value = val
	v.isSet = true
}

func (v NullableCSR) IsSet() bool {
	return v.isSet
}

func (v *NullableCSR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSR(val *CSR) *NullableCSR {
	return &NullableCSR{value: val, isSet: true}
}

func (v NullableCSR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
