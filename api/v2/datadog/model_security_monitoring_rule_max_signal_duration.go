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

import (
	"fmt"
)

// SecurityMonitoringRuleMaxSignalDuration A signal will “close” regardless of the query being matched once the time exceeds the maximum duration. This time is calculated from the first seen timestamp.
type SecurityMonitoringRuleMaxSignalDuration int32

// List of SecurityMonitoringRuleMaxSignalDuration
const (
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ZERO_MINUTES    SecurityMonitoringRuleMaxSignalDuration = 0
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_MINUTE      SecurityMonitoringRuleMaxSignalDuration = 60
	SECURITYMONITORINGRULEMAXSIGNALDURATION_FIVE_MINUTES    SecurityMonitoringRuleMaxSignalDuration = 300
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TEN_MINUTES     SecurityMonitoringRuleMaxSignalDuration = 600
	SECURITYMONITORINGRULEMAXSIGNALDURATION_FIFTEEN_MINUTES SecurityMonitoringRuleMaxSignalDuration = 900
	SECURITYMONITORINGRULEMAXSIGNALDURATION_THIRTY_MINUTES  SecurityMonitoringRuleMaxSignalDuration = 1800
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_HOUR        SecurityMonitoringRuleMaxSignalDuration = 3600
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TWO_HOURS       SecurityMonitoringRuleMaxSignalDuration = 7200
	SECURITYMONITORINGRULEMAXSIGNALDURATION_THREE_HOURS     SecurityMonitoringRuleMaxSignalDuration = 10800
	SECURITYMONITORINGRULEMAXSIGNALDURATION_SIX_HOURS       SecurityMonitoringRuleMaxSignalDuration = 21600
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TWELVE_HOURS    SecurityMonitoringRuleMaxSignalDuration = 43200
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY         SecurityMonitoringRuleMaxSignalDuration = 86400
)

func (v *SecurityMonitoringRuleMaxSignalDuration) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityMonitoringRuleMaxSignalDuration(value)
	for _, existing := range []SecurityMonitoringRuleMaxSignalDuration{0, 60, 300, 600, 900, 1800, 3600, 7200, 10800, 21600, 43200, 86400} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityMonitoringRuleMaxSignalDuration", *v)
}

// Ptr returns reference to SecurityMonitoringRuleMaxSignalDuration value
func (v SecurityMonitoringRuleMaxSignalDuration) Ptr() *SecurityMonitoringRuleMaxSignalDuration {
	return &v
}

type NullableSecurityMonitoringRuleMaxSignalDuration struct {
	value *SecurityMonitoringRuleMaxSignalDuration
	isSet bool
}

func (v NullableSecurityMonitoringRuleMaxSignalDuration) Get() *SecurityMonitoringRuleMaxSignalDuration {
	return v.value
}

func (v *NullableSecurityMonitoringRuleMaxSignalDuration) Set(val *SecurityMonitoringRuleMaxSignalDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleMaxSignalDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleMaxSignalDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleMaxSignalDuration(val *SecurityMonitoringRuleMaxSignalDuration) *NullableSecurityMonitoringRuleMaxSignalDuration {
	return &NullableSecurityMonitoringRuleMaxSignalDuration{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleMaxSignalDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleMaxSignalDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
