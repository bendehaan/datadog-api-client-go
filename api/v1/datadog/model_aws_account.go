/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// AWSAccount Returns the AWS account associated with this integration.
type AWSAccount struct {
	// Only return AWS accounts that matches this access_key_id.
	AccessKeyId *string `json:"access_key_id,omitempty"`
	// Only return AWS accounts that match this account_id.
	AccountId *string `json:"account_id,omitempty"`
	// An object (in the form {\"namespace1\":true/false, \"namespace2\":true/false}) that enables or disables metric collection for specific AWS namespaces for this AWS account only. A list of namespaces can be found at the /v1/integration/aws/available_namespace_rules endpoint.
	AccountSpecificNamespaceRules *map[string]bool `json:"account_specific_namespace_rules,omitempty"`
	// The array of EC2 tags (in the form key:value) defines a filter that Datadog uses when collecting metrics from EC2. Wildcards, such as ? (for single characters) and * (for multiple characters) can also be used. Only hosts that match one of the defined tags will be imported into Datadog. The rest will be ignored. Host matching a given tag can also be excluded by adding ! before the tag. For example, `env:production,instance-type:c1.*,!region:us-east-1`
	FilterTags *[]string `json:"filter_tags,omitempty"`
	// Array of tags (in the form key:value) to add to all hosts and metrics reporting through this integration.
	HostTags *[]string `json:"host_tags,omitempty"`
	// Only return AWS accounts that matches this role_name.
	RoleName *string `json:"role_name,omitempty"`
}

// NewAWSAccount instantiates a new AWSAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSAccount() *AWSAccount {
	this := AWSAccount{}
	return &this
}

// NewAWSAccountWithDefaults instantiates a new AWSAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSAccountWithDefaults() *AWSAccount {
	this := AWSAccount{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AWSAccount) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccount) GetAccessKeyIdOk() (string, bool) {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret, false
	}
	return *o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AWSAccount) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AWSAccount) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AWSAccount) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccount) GetAccountIdOk() (string, bool) {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret, false
	}
	return *o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AWSAccount) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AWSAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountSpecificNamespaceRules returns the AccountSpecificNamespaceRules field value if set, zero value otherwise.
func (o *AWSAccount) GetAccountSpecificNamespaceRules() map[string]bool {
	if o == nil || o.AccountSpecificNamespaceRules == nil {
		var ret map[string]bool
		return ret
	}
	return *o.AccountSpecificNamespaceRules
}

// GetAccountSpecificNamespaceRulesOk returns a tuple with the AccountSpecificNamespaceRules field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccount) GetAccountSpecificNamespaceRulesOk() (map[string]bool, bool) {
	if o == nil || o.AccountSpecificNamespaceRules == nil {
		var ret map[string]bool
		return ret, false
	}
	return *o.AccountSpecificNamespaceRules, true
}

// HasAccountSpecificNamespaceRules returns a boolean if a field has been set.
func (o *AWSAccount) HasAccountSpecificNamespaceRules() bool {
	if o != nil && o.AccountSpecificNamespaceRules != nil {
		return true
	}

	return false
}

// SetAccountSpecificNamespaceRules gets a reference to the given map[string]bool and assigns it to the AccountSpecificNamespaceRules field.
func (o *AWSAccount) SetAccountSpecificNamespaceRules(v map[string]bool) {
	o.AccountSpecificNamespaceRules = &v
}

// GetFilterTags returns the FilterTags field value if set, zero value otherwise.
func (o *AWSAccount) GetFilterTags() []string {
	if o == nil || o.FilterTags == nil {
		var ret []string
		return ret
	}
	return *o.FilterTags
}

// GetFilterTagsOk returns a tuple with the FilterTags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccount) GetFilterTagsOk() ([]string, bool) {
	if o == nil || o.FilterTags == nil {
		var ret []string
		return ret, false
	}
	return *o.FilterTags, true
}

// HasFilterTags returns a boolean if a field has been set.
func (o *AWSAccount) HasFilterTags() bool {
	if o != nil && o.FilterTags != nil {
		return true
	}

	return false
}

// SetFilterTags gets a reference to the given []string and assigns it to the FilterTags field.
func (o *AWSAccount) SetFilterTags(v []string) {
	o.FilterTags = &v
}

// GetHostTags returns the HostTags field value if set, zero value otherwise.
func (o *AWSAccount) GetHostTags() []string {
	if o == nil || o.HostTags == nil {
		var ret []string
		return ret
	}
	return *o.HostTags
}

// GetHostTagsOk returns a tuple with the HostTags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccount) GetHostTagsOk() ([]string, bool) {
	if o == nil || o.HostTags == nil {
		var ret []string
		return ret, false
	}
	return *o.HostTags, true
}

// HasHostTags returns a boolean if a field has been set.
func (o *AWSAccount) HasHostTags() bool {
	if o != nil && o.HostTags != nil {
		return true
	}

	return false
}

// SetHostTags gets a reference to the given []string and assigns it to the HostTags field.
func (o *AWSAccount) SetHostTags(v []string) {
	o.HostTags = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *AWSAccount) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccount) GetRoleNameOk() (string, bool) {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret, false
	}
	return *o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *AWSAccount) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *AWSAccount) SetRoleName(v string) {
	o.RoleName = &v
}

func (o AWSAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKeyId != nil {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.AccountSpecificNamespaceRules != nil {
		toSerialize["account_specific_namespace_rules"] = o.AccountSpecificNamespaceRules
	}
	if o.FilterTags != nil {
		toSerialize["filter_tags"] = o.FilterTags
	}
	if o.HostTags != nil {
		toSerialize["host_tags"] = o.HostTags
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableAWSAccount struct {
	value *AWSAccount
	isSet bool
}

func (v NullableAWSAccount) Get() *AWSAccount {
	return v.value
}

func (v NullableAWSAccount) Set(val *AWSAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSAccount) IsSet() bool {
	return v.isSet
}

func (v NullableAWSAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSAccount(val *AWSAccount) *NullableAWSAccount {
	return &NullableAWSAccount{value: val, isSet: true}
}

func (v NullableAWSAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
