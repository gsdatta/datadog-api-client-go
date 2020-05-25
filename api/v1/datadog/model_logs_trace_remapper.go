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

// LogsTraceRemapper There are two ways to improve correlation between application traces and logs.    1. Follow the documentation on [how to inject a trace ID in the application logs](https://docs.datadoghq.com/tracing/connect_logs_and_traces)   and by default log integrations take care of all the rest of the setup.    2. Use the Trace remapper processor to define a log attribute as its associated trace ID.
type LogsTraceRemapper struct {
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// Array of source attributes.
	Sources *[]string             `json:"sources,omitempty"`
	Type    LogsTraceRemapperType `json:"type"`
}

// NewLogsTraceRemapper instantiates a new LogsTraceRemapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsTraceRemapper(type_ LogsTraceRemapperType) *LogsTraceRemapper {
	this := LogsTraceRemapper{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	this.Type = type_
	return &this
}

// NewLogsTraceRemapperWithDefaults instantiates a new LogsTraceRemapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsTraceRemapperWithDefaults() *LogsTraceRemapper {
	this := LogsTraceRemapper{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var type_ LogsTraceRemapperType = "trace-id-remapper"
	this.Type = type_
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsTraceRemapper) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTraceRemapper) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsTraceRemapper) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsTraceRemapper) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsTraceRemapper) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTraceRemapper) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsTraceRemapper) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsTraceRemapper) SetName(v string) {
	o.Name = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *LogsTraceRemapper) GetSources() []string {
	if o == nil || o.Sources == nil {
		var ret []string
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsTraceRemapper) GetSourcesOk() (*[]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *LogsTraceRemapper) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *LogsTraceRemapper) SetSources(v []string) {
	o.Sources = &v
}

// GetType returns the Type field value
func (o *LogsTraceRemapper) GetType() LogsTraceRemapperType {
	if o == nil {
		var ret LogsTraceRemapperType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsTraceRemapper) GetTypeOk() (*LogsTraceRemapperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogsTraceRemapper) SetType(v LogsTraceRemapperType) {
	o.Type = v
}

func (o LogsTraceRemapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLogsTraceRemapper struct {
	value *LogsTraceRemapper
	isSet bool
}

func (v NullableLogsTraceRemapper) Get() *LogsTraceRemapper {
	return v.value
}

func (v *NullableLogsTraceRemapper) Set(val *LogsTraceRemapper) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsTraceRemapper) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsTraceRemapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsTraceRemapper(val *LogsTraceRemapper) *NullableLogsTraceRemapper {
	return &NullableLogsTraceRemapper{value: val, isSet: true}
}

func (v NullableLogsTraceRemapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsTraceRemapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
