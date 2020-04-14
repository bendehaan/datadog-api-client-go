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

// QueryMetadataPagination Pagination properties.
type QueryMetadataPagination struct {
	// maximum number of pages to return
	Limit *int64 `json:"limit,omitempty"`
	// the next offset to retreive the next set of data
	NextOffset *int64 `json:"next_offset,omitempty"`
	// the index of the first element in the results
	Offset *int64 `json:"offset,omitempty"`
}

// NewQueryMetadataPagination instantiates a new QueryMetadataPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMetadataPagination() *QueryMetadataPagination {
	this := QueryMetadataPagination{}
	return &this
}

// NewQueryMetadataPaginationWithDefaults instantiates a new QueryMetadataPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMetadataPaginationWithDefaults() *QueryMetadataPagination {
	this := QueryMetadataPagination{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryMetadataPagination) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMetadataPagination) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryMetadataPagination) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *QueryMetadataPagination) SetLimit(v int64) {
	o.Limit = &v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *QueryMetadataPagination) GetNextOffset() int64 {
	if o == nil || o.NextOffset == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMetadataPagination) GetNextOffsetOk() (*int64, bool) {
	if o == nil || o.NextOffset == nil {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *QueryMetadataPagination) HasNextOffset() bool {
	if o != nil && o.NextOffset != nil {
		return true
	}

	return false
}

// SetNextOffset gets a reference to the given int64 and assigns it to the NextOffset field.
func (o *QueryMetadataPagination) SetNextOffset(v int64) {
	o.NextOffset = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryMetadataPagination) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMetadataPagination) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryMetadataPagination) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *QueryMetadataPagination) SetOffset(v int64) {
	o.Offset = &v
}

func (o QueryMetadataPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextOffset != nil {
		toSerialize["next_offset"] = o.NextOffset
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	return json.Marshal(toSerialize)
}

type NullableQueryMetadataPagination struct {
	value *QueryMetadataPagination
	isSet bool
}

func (v NullableQueryMetadataPagination) Get() *QueryMetadataPagination {
	return v.value
}

func (v *NullableQueryMetadataPagination) Set(val *QueryMetadataPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMetadataPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMetadataPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMetadataPagination(val *QueryMetadataPagination) *NullableQueryMetadataPagination {
	return &NullableQueryMetadataPagination{value: val, isSet: true}
}

func (v NullableQueryMetadataPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMetadataPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}