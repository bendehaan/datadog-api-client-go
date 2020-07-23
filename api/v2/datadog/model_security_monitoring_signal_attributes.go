/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// SecurityMonitoringSignalAttributes JSON object containing all signal attributes and their associated values.
type SecurityMonitoringSignalAttributes struct {
	// JSON object of attributes in the security signal.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The message the security signal, defined by the rule which which generated the signal.
	Message *string `json:"message,omitempty"`
	// Array of tags associated with the security signal.
	Tags *[]interface{} `json:"tags,omitempty"`
	// Timestamp of the security signal.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewSecurityMonitoringSignalAttributes instantiates a new SecurityMonitoringSignalAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringSignalAttributes() *SecurityMonitoringSignalAttributes {
	this := SecurityMonitoringSignalAttributes{}
	return &this
}

// NewSecurityMonitoringSignalAttributesWithDefaults instantiates a new SecurityMonitoringSignalAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringSignalAttributesWithDefaults() *SecurityMonitoringSignalAttributes {
	this := SecurityMonitoringSignalAttributes{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalAttributes) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalAttributes) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalAttributes) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *SecurityMonitoringSignalAttributes) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalAttributes) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SecurityMonitoringSignalAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalAttributes) GetTags() []interface{} {
	if o == nil || o.Tags == nil {
		var ret []interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalAttributes) GetTagsOk() (*[]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalAttributes) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []interface{} and assigns it to the Tags field.
func (o *SecurityMonitoringSignalAttributes) SetTags(v []interface{}) {
	o.Tags = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalAttributes) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalAttributes) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalAttributes) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SecurityMonitoringSignalAttributes) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o SecurityMonitoringSignalAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityMonitoringSignalAttributes struct {
	value *SecurityMonitoringSignalAttributes
	isSet bool
}

func (v NullableSecurityMonitoringSignalAttributes) Get() *SecurityMonitoringSignalAttributes {
	return v.value
}

func (v *NullableSecurityMonitoringSignalAttributes) Set(val *SecurityMonitoringSignalAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringSignalAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringSignalAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringSignalAttributes(val *SecurityMonitoringSignalAttributes) *NullableSecurityMonitoringSignalAttributes {
	return &NullableSecurityMonitoringSignalAttributes{value: val, isSet: true}
}

func (v NullableSecurityMonitoringSignalAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringSignalAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
