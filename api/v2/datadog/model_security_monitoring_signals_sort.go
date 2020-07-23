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

// SecurityMonitoringSignalsSort Sort parameters when querying security signals.
type SecurityMonitoringSignalsSort string

// List of SecurityMonitoringSignalsSort
const (
	SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING  SecurityMonitoringSignalsSort = "timestamp"
	SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_DESCENDING SecurityMonitoringSignalsSort = "-timestamp"
)

func (v *SecurityMonitoringSignalsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityMonitoringSignalsSort(value)
	for _, existing := range []SecurityMonitoringSignalsSort{"timestamp", "-timestamp"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityMonitoringSignalsSort", *v)
}

// Ptr returns reference to SecurityMonitoringSignalsSort value
func (v SecurityMonitoringSignalsSort) Ptr() *SecurityMonitoringSignalsSort {
	return &v
}

type NullableSecurityMonitoringSignalsSort struct {
	value *SecurityMonitoringSignalsSort
	isSet bool
}

func (v NullableSecurityMonitoringSignalsSort) Get() *SecurityMonitoringSignalsSort {
	return v.value
}

func (v *NullableSecurityMonitoringSignalsSort) Set(val *SecurityMonitoringSignalsSort) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringSignalsSort) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringSignalsSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringSignalsSort(val *SecurityMonitoringSignalsSort) *NullableSecurityMonitoringSignalsSort {
	return &NullableSecurityMonitoringSignalsSort{value: val, isSet: true}
}

func (v NullableSecurityMonitoringSignalsSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringSignalsSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
