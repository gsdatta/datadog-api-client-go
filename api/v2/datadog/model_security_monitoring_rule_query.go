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

// SecurityMonitoringRuleQuery Query for matching rule.
type SecurityMonitoringRuleQuery struct {
	AgentRule   *SecurityMonitoringRuntimeAgentRule     `json:"agentRule,omitempty"`
	Aggregation *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
	// Field for which the cardinality is measured. Sent as an array.
	DistinctFields *[]string `json:"distinctFields,omitempty"`
	// Fields to group by.
	GroupByFields *[]string `json:"groupByFields,omitempty"`
	// The target field to aggregate over when using the sum or max aggregations.
	Metric *string `json:"metric,omitempty"`
	// Name of the query.
	Name *string `json:"name,omitempty"`
	// Query to run on logs.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSecurityMonitoringRuleQuery instantiates a new SecurityMonitoringRuleQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringRuleQuery() *SecurityMonitoringRuleQuery {
	this := SecurityMonitoringRuleQuery{}
	return &this
}

// NewSecurityMonitoringRuleQueryWithDefaults instantiates a new SecurityMonitoringRuleQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringRuleQueryWithDefaults() *SecurityMonitoringRuleQuery {
	this := SecurityMonitoringRuleQuery{}
	return &this
}

// GetAgentRule returns the AgentRule field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetAgentRule() SecurityMonitoringRuntimeAgentRule {
	if o == nil || o.AgentRule == nil {
		var ret SecurityMonitoringRuntimeAgentRule
		return ret
	}
	return *o.AgentRule
}

// GetAgentRuleOk returns a tuple with the AgentRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetAgentRuleOk() (*SecurityMonitoringRuntimeAgentRule, bool) {
	if o == nil || o.AgentRule == nil {
		return nil, false
	}
	return o.AgentRule, true
}

// HasAgentRule returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasAgentRule() bool {
	if o != nil && o.AgentRule != nil {
		return true
	}

	return false
}

// SetAgentRule gets a reference to the given SecurityMonitoringRuntimeAgentRule and assigns it to the AgentRule field.
func (o *SecurityMonitoringRuleQuery) SetAgentRule(v SecurityMonitoringRuntimeAgentRule) {
	o.AgentRule = &v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetAggregation() SecurityMonitoringRuleQueryAggregation {
	if o == nil || o.Aggregation == nil {
		var ret SecurityMonitoringRuleQueryAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetAggregationOk() (*SecurityMonitoringRuleQueryAggregation, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given SecurityMonitoringRuleQueryAggregation and assigns it to the Aggregation field.
func (o *SecurityMonitoringRuleQuery) SetAggregation(v SecurityMonitoringRuleQueryAggregation) {
	o.Aggregation = &v
}

// GetDistinctFields returns the DistinctFields field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetDistinctFields() []string {
	if o == nil || o.DistinctFields == nil {
		var ret []string
		return ret
	}
	return *o.DistinctFields
}

// GetDistinctFieldsOk returns a tuple with the DistinctFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetDistinctFieldsOk() (*[]string, bool) {
	if o == nil || o.DistinctFields == nil {
		return nil, false
	}
	return o.DistinctFields, true
}

// HasDistinctFields returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasDistinctFields() bool {
	if o != nil && o.DistinctFields != nil {
		return true
	}

	return false
}

// SetDistinctFields gets a reference to the given []string and assigns it to the DistinctFields field.
func (o *SecurityMonitoringRuleQuery) SetDistinctFields(v []string) {
	o.DistinctFields = &v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetGroupByFields() []string {
	if o == nil || o.GroupByFields == nil {
		var ret []string
		return ret
	}
	return *o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetGroupByFieldsOk() (*[]string, bool) {
	if o == nil || o.GroupByFields == nil {
		return nil, false
	}
	return o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasGroupByFields() bool {
	if o != nil && o.GroupByFields != nil {
		return true
	}

	return false
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *SecurityMonitoringRuleQuery) SetGroupByFields(v []string) {
	o.GroupByFields = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *SecurityMonitoringRuleQuery) SetMetric(v string) {
	o.Metric = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringRuleQuery) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQuery) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SecurityMonitoringRuleQuery) SetQuery(v string) {
	o.Query = &v
}

func (o SecurityMonitoringRuleQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AgentRule != nil {
		toSerialize["agentRule"] = o.AgentRule
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.DistinctFields != nil {
		toSerialize["distinctFields"] = o.DistinctFields
	}
	if o.GroupByFields != nil {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityMonitoringRuleQuery) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AgentRule      *SecurityMonitoringRuntimeAgentRule     `json:"agentRule,omitempty"`
		Aggregation    *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
		DistinctFields *[]string                               `json:"distinctFields,omitempty"`
		GroupByFields  *[]string                               `json:"groupByFields,omitempty"`
		Metric         *string                                 `json:"metric,omitempty"`
		Name           *string                                 `json:"name,omitempty"`
		Query          *string                                 `json:"query,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Aggregation; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.AgentRule = all.AgentRule
	o.Aggregation = all.Aggregation
	o.DistinctFields = all.DistinctFields
	o.GroupByFields = all.GroupByFields
	o.Metric = all.Metric
	o.Name = all.Name
	o.Query = all.Query
	return nil
}

type NullableSecurityMonitoringRuleQuery struct {
	value *SecurityMonitoringRuleQuery
	isSet bool
}

func (v NullableSecurityMonitoringRuleQuery) Get() *SecurityMonitoringRuleQuery {
	return v.value
}

func (v *NullableSecurityMonitoringRuleQuery) Set(val *SecurityMonitoringRuleQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleQuery(val *SecurityMonitoringRuleQuery) *NullableSecurityMonitoringRuleQuery {
	return &NullableSecurityMonitoringRuleQuery{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
