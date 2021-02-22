/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SecurityMonitoringRuleDetectionMethod The detection method.
type SecurityMonitoringRuleDetectionMethod string

// List of SecurityMonitoringRuleDetectionMethod
const (
	SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD SecurityMonitoringRuleDetectionMethod = "threshold"
	SECURITYMONITORINGRULEDETECTIONMETHOD_NEW_VALUE SecurityMonitoringRuleDetectionMethod = "new_value"
)

var allowedSecurityMonitoringRuleDetectionMethodEnumValues = []SecurityMonitoringRuleDetectionMethod{
	"threshold",
	"new_value",
}

func (v *SecurityMonitoringRuleDetectionMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityMonitoringRuleDetectionMethod(value)
	for _, existing := range allowedSecurityMonitoringRuleDetectionMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityMonitoringRuleDetectionMethod", value)
}

// NewSecurityMonitoringRuleDetectionMethodFromValue returns a pointer to a valid SecurityMonitoringRuleDetectionMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityMonitoringRuleDetectionMethodFromValue(v string) (*SecurityMonitoringRuleDetectionMethod, error) {
	ev := SecurityMonitoringRuleDetectionMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleDetectionMethod: valid values are %v", v, allowedSecurityMonitoringRuleDetectionMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityMonitoringRuleDetectionMethod) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleDetectionMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleDetectionMethod value
func (v SecurityMonitoringRuleDetectionMethod) Ptr() *SecurityMonitoringRuleDetectionMethod {
	return &v
}

type NullableSecurityMonitoringRuleDetectionMethod struct {
	value *SecurityMonitoringRuleDetectionMethod
	isSet bool
}

func (v NullableSecurityMonitoringRuleDetectionMethod) Get() *SecurityMonitoringRuleDetectionMethod {
	return v.value
}

func (v *NullableSecurityMonitoringRuleDetectionMethod) Set(val *SecurityMonitoringRuleDetectionMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleDetectionMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleDetectionMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleDetectionMethod(val *SecurityMonitoringRuleDetectionMethod) *NullableSecurityMonitoringRuleDetectionMethod {
	return &NullableSecurityMonitoringRuleDetectionMethod{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleDetectionMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleDetectionMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
