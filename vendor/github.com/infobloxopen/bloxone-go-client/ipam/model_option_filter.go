/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the OptionFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionFilter{}

// OptionFilter An __OptionFilter__ object (_dhcp/option_filter_) applies options to DHCP clients with matching options. It must be configured in the _filters_ list for a scope to be effective.
type OptionFilter struct {
	// The description for the option filter. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The list of DHCP options for the option filter. May be either a specific option or a group of options.
	DhcpOptions []OptionItem `json:"dhcp_options,omitempty"`
	// The configuration for header option filename field.
	HeaderOptionFilename *string `json:"header_option_filename,omitempty"`
	// The configuration for header option server address field.
	HeaderOptionServerAddress *string `json:"header_option_server_address,omitempty"`
	// The configuration for header option server name field.
	HeaderOptionServerName *string `json:"header_option_server_name,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The lease lifetime duration in seconds.
	LeaseTime *int64 `json:"lease_time,omitempty"`
	// The name of the option filter. Must contain 1 to 256 characters. Can include UTF-8.
	Name string `json:"name"`
	// The type of protocol of option filter (_ip4_ or _ip6_).
	Protocol *string `json:"protocol,omitempty"`
	// The role of DHCP filter (_values_ or _selection_).  Defaults to _values_.
	Role  *string              `json:"role,omitempty"`
	Rules OptionFilterRuleList `json:"rules"`
	// The tags for the option filter in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The resource identifier.
	VendorSpecificOptionOptionSpace *string `json:"vendor_specific_option_option_space,omitempty"`
	AdditionalProperties            map[string]interface{}
}

type _OptionFilter OptionFilter

// NewOptionFilter instantiates a new OptionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionFilter(name string, rules OptionFilterRuleList) *OptionFilter {
	this := OptionFilter{}
	this.Name = name
	this.Rules = rules
	return &this
}

// NewOptionFilterWithDefaults instantiates a new OptionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionFilterWithDefaults() *OptionFilter {
	this := OptionFilter{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *OptionFilter) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *OptionFilter) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *OptionFilter) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OptionFilter) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OptionFilter) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OptionFilter) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *OptionFilter) GetDhcpOptions() []OptionItem {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret []OptionItem
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetDhcpOptionsOk() ([]OptionItem, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *OptionFilter) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []OptionItem and assigns it to the DhcpOptions field.
func (o *OptionFilter) SetDhcpOptions(v []OptionItem) {
	o.DhcpOptions = v
}

// GetHeaderOptionFilename returns the HeaderOptionFilename field value if set, zero value otherwise.
func (o *OptionFilter) GetHeaderOptionFilename() string {
	if o == nil || IsNil(o.HeaderOptionFilename) {
		var ret string
		return ret
	}
	return *o.HeaderOptionFilename
}

// GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetHeaderOptionFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderOptionFilename) {
		return nil, false
	}
	return o.HeaderOptionFilename, true
}

// HasHeaderOptionFilename returns a boolean if a field has been set.
func (o *OptionFilter) HasHeaderOptionFilename() bool {
	if o != nil && !IsNil(o.HeaderOptionFilename) {
		return true
	}

	return false
}

// SetHeaderOptionFilename gets a reference to the given string and assigns it to the HeaderOptionFilename field.
func (o *OptionFilter) SetHeaderOptionFilename(v string) {
	o.HeaderOptionFilename = &v
}

// GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field value if set, zero value otherwise.
func (o *OptionFilter) GetHeaderOptionServerAddress() string {
	if o == nil || IsNil(o.HeaderOptionServerAddress) {
		var ret string
		return ret
	}
	return *o.HeaderOptionServerAddress
}

// GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetHeaderOptionServerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderOptionServerAddress) {
		return nil, false
	}
	return o.HeaderOptionServerAddress, true
}

// HasHeaderOptionServerAddress returns a boolean if a field has been set.
func (o *OptionFilter) HasHeaderOptionServerAddress() bool {
	if o != nil && !IsNil(o.HeaderOptionServerAddress) {
		return true
	}

	return false
}

// SetHeaderOptionServerAddress gets a reference to the given string and assigns it to the HeaderOptionServerAddress field.
func (o *OptionFilter) SetHeaderOptionServerAddress(v string) {
	o.HeaderOptionServerAddress = &v
}

// GetHeaderOptionServerName returns the HeaderOptionServerName field value if set, zero value otherwise.
func (o *OptionFilter) GetHeaderOptionServerName() string {
	if o == nil || IsNil(o.HeaderOptionServerName) {
		var ret string
		return ret
	}
	return *o.HeaderOptionServerName
}

// GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetHeaderOptionServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderOptionServerName) {
		return nil, false
	}
	return o.HeaderOptionServerName, true
}

// HasHeaderOptionServerName returns a boolean if a field has been set.
func (o *OptionFilter) HasHeaderOptionServerName() bool {
	if o != nil && !IsNil(o.HeaderOptionServerName) {
		return true
	}

	return false
}

// SetHeaderOptionServerName gets a reference to the given string and assigns it to the HeaderOptionServerName field.
func (o *OptionFilter) SetHeaderOptionServerName(v string) {
	o.HeaderOptionServerName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionFilter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptionFilter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OptionFilter) SetId(v string) {
	o.Id = &v
}

// GetLeaseTime returns the LeaseTime field value if set, zero value otherwise.
func (o *OptionFilter) GetLeaseTime() int64 {
	if o == nil || IsNil(o.LeaseTime) {
		var ret int64
		return ret
	}
	return *o.LeaseTime
}

// GetLeaseTimeOk returns a tuple with the LeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetLeaseTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.LeaseTime) {
		return nil, false
	}
	return o.LeaseTime, true
}

// HasLeaseTime returns a boolean if a field has been set.
func (o *OptionFilter) HasLeaseTime() bool {
	if o != nil && !IsNil(o.LeaseTime) {
		return true
	}

	return false
}

// SetLeaseTime gets a reference to the given int64 and assigns it to the LeaseTime field.
func (o *OptionFilter) SetLeaseTime(v int64) {
	o.LeaseTime = &v
}

// GetName returns the Name field value
func (o *OptionFilter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OptionFilter) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *OptionFilter) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *OptionFilter) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *OptionFilter) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OptionFilter) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OptionFilter) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OptionFilter) SetRole(v string) {
	o.Role = &v
}

// GetRules returns the Rules field value
func (o *OptionFilter) GetRules() OptionFilterRuleList {
	if o == nil {
		var ret OptionFilterRuleList
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetRulesOk() (*OptionFilterRuleList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *OptionFilter) SetRules(v OptionFilterRuleList) {
	o.Rules = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OptionFilter) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OptionFilter) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *OptionFilter) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OptionFilter) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OptionFilter) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OptionFilter) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field value if set, zero value otherwise.
func (o *OptionFilter) GetVendorSpecificOptionOptionSpace() string {
	if o == nil || IsNil(o.VendorSpecificOptionOptionSpace) {
		var ret string
		return ret
	}
	return *o.VendorSpecificOptionOptionSpace
}

// GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionFilter) GetVendorSpecificOptionOptionSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.VendorSpecificOptionOptionSpace) {
		return nil, false
	}
	return o.VendorSpecificOptionOptionSpace, true
}

// HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.
func (o *OptionFilter) HasVendorSpecificOptionOptionSpace() bool {
	if o != nil && !IsNil(o.VendorSpecificOptionOptionSpace) {
		return true
	}

	return false
}

// SetVendorSpecificOptionOptionSpace gets a reference to the given string and assigns it to the VendorSpecificOptionOptionSpace field.
func (o *OptionFilter) SetVendorSpecificOptionOptionSpace(v string) {
	o.VendorSpecificOptionOptionSpace = &v
}

func (o OptionFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcp_options"] = o.DhcpOptions
	}
	if !IsNil(o.HeaderOptionFilename) {
		toSerialize["header_option_filename"] = o.HeaderOptionFilename
	}
	if !IsNil(o.HeaderOptionServerAddress) {
		toSerialize["header_option_server_address"] = o.HeaderOptionServerAddress
	}
	if !IsNil(o.HeaderOptionServerName) {
		toSerialize["header_option_server_name"] = o.HeaderOptionServerName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LeaseTime) {
		toSerialize["lease_time"] = o.LeaseTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	toSerialize["rules"] = o.Rules
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VendorSpecificOptionOptionSpace) {
		toSerialize["vendor_specific_option_option_space"] = o.VendorSpecificOptionOptionSpace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"rules",
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

	varOptionFilter := _OptionFilter{}

	err = json.Unmarshal(data, &varOptionFilter)

	if err != nil {
		return err
	}

	*o = OptionFilter(varOptionFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "dhcp_options")
		delete(additionalProperties, "header_option_filename")
		delete(additionalProperties, "header_option_server_address")
		delete(additionalProperties, "header_option_server_name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lease_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "role")
		delete(additionalProperties, "rules")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "vendor_specific_option_option_space")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOptionFilter struct {
	value *OptionFilter
	isSet bool
}

func (v NullableOptionFilter) Get() *OptionFilter {
	return v.value
}

func (v *NullableOptionFilter) Set(val *OptionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionFilter(val *OptionFilter) *NullableOptionFilter {
	return &NullableOptionFilter{value: val, isSet: true}
}

func (v NullableOptionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}