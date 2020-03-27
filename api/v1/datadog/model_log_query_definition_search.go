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

// LogQueryDefinitionSearch struct for LogQueryDefinitionSearch
type LogQueryDefinitionSearch struct {
	Query string `json:"query"`
}

// NewLogQueryDefinitionSearch instantiates a new LogQueryDefinitionSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogQueryDefinitionSearch(query string) *LogQueryDefinitionSearch {
	this := LogQueryDefinitionSearch{}
	this.Query = query
	return &this
}

// NewLogQueryDefinitionSearchWithDefaults instantiates a new LogQueryDefinitionSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogQueryDefinitionSearchWithDefaults() *LogQueryDefinitionSearch {
	this := LogQueryDefinitionSearch{}
	return &this
}

// GetQuery returns the Query field value
func (o *LogQueryDefinitionSearch) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LogQueryDefinitionSearch) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *LogQueryDefinitionSearch) SetQuery(v string) {
	o.Query = v
}

func (o LogQueryDefinitionSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableLogQueryDefinitionSearch struct {
	value *LogQueryDefinitionSearch
	isSet bool
}

func (v NullableLogQueryDefinitionSearch) Get() *LogQueryDefinitionSearch {
	return v.value
}

func (v *NullableLogQueryDefinitionSearch) Set(val *LogQueryDefinitionSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableLogQueryDefinitionSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableLogQueryDefinitionSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogQueryDefinitionSearch(val *LogQueryDefinitionSearch) *NullableLogQueryDefinitionSearch {
	return &NullableLogQueryDefinitionSearch{value: val, isSet: true}
}

func (v NullableLogQueryDefinitionSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogQueryDefinitionSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
