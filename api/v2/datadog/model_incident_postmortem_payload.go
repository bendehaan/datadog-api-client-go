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

// IncidentPostmortemPayload Represents the JSON API Payload of an Incident Postmortem Item.
type IncidentPostmortemPayload struct {
	Data IncidentPostmortem `json:"data"`
	// The User relationships.
	Included *[]UserJSONAPIResponse `json:"included,omitempty"`
}

// NewIncidentPostmortemPayload instantiates a new IncidentPostmortemPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentPostmortemPayload(data IncidentPostmortem) *IncidentPostmortemPayload {
	this := IncidentPostmortemPayload{}
	this.Data = data
	return &this
}

// NewIncidentPostmortemPayloadWithDefaults instantiates a new IncidentPostmortemPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentPostmortemPayloadWithDefaults() *IncidentPostmortemPayload {
	this := IncidentPostmortemPayload{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentPostmortemPayload) GetData() IncidentPostmortem {
	if o == nil {
		var ret IncidentPostmortem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemPayload) GetDataOk() (*IncidentPostmortem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentPostmortemPayload) SetData(v IncidentPostmortem) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentPostmortemPayload) GetIncluded() []UserJSONAPIResponse {
	if o == nil || o.Included == nil {
		var ret []UserJSONAPIResponse
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentPostmortemPayload) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []UserJSONAPIResponse and assigns it to the Included field.
func (o *IncidentPostmortemPayload) SetIncluded(v []UserJSONAPIResponse) {
	o.Included = &v
}

func (o IncidentPostmortemPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentPostmortemPayload struct {
	value *IncidentPostmortemPayload
	isSet bool
}

func (v NullableIncidentPostmortemPayload) Get() *IncidentPostmortemPayload {
	return v.value
}

func (v *NullableIncidentPostmortemPayload) Set(val *IncidentPostmortemPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentPostmortemPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentPostmortemPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentPostmortemPayload(val *IncidentPostmortemPayload) *NullableIncidentPostmortemPayload {
	return &NullableIncidentPostmortemPayload{value: val, isSet: true}
}

func (v NullableIncidentPostmortemPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentPostmortemPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}