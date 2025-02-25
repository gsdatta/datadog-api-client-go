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

// SLOCorrectionResponseAttributes The attribute object associated with the SLO correction.
type SLOCorrectionResponseAttributes struct {
	Category *SLOCorrectionCategory `json:"category,omitempty"`
	Creator  *Creator               `json:"creator,omitempty"`
	// Description of the correction being made.
	Description *string `json:"description,omitempty"`
	// Ending time of the correction in epoch seconds.
	End *int64 `json:"end,omitempty"`
	// ID of the SLO that this correction will be applied to.
	SloId *string `json:"slo_id,omitempty"`
	// Starting time of the correction in epoch seconds.
	Start *int64 `json:"start,omitempty"`
	// The timezone to display in the UI for the correction times (defaults to \"UTC\").
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSLOCorrectionResponseAttributes instantiates a new SLOCorrectionResponseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOCorrectionResponseAttributes() *SLOCorrectionResponseAttributes {
	this := SLOCorrectionResponseAttributes{}
	return &this
}

// NewSLOCorrectionResponseAttributesWithDefaults instantiates a new SLOCorrectionResponseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOCorrectionResponseAttributesWithDefaults() *SLOCorrectionResponseAttributes {
	this := SLOCorrectionResponseAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SLOCorrectionResponseAttributes) GetCategory() SLOCorrectionCategory {
	if o == nil || o.Category == nil {
		var ret SLOCorrectionCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionResponseAttributes) GetCategoryOk() (*SLOCorrectionCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SLOCorrectionResponseAttributes) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given SLOCorrectionCategory and assigns it to the Category field.
func (o *SLOCorrectionResponseAttributes) SetCategory(v SLOCorrectionCategory) {
	o.Category = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SLOCorrectionResponseAttributes) GetCreator() Creator {
	if o == nil || o.Creator == nil {
		var ret Creator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionResponseAttributes) GetCreatorOk() (*Creator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SLOCorrectionResponseAttributes) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given Creator and assigns it to the Creator field.
func (o *SLOCorrectionResponseAttributes) SetCreator(v Creator) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SLOCorrectionResponseAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SLOCorrectionResponseAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SLOCorrectionResponseAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SLOCorrectionResponseAttributes) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionResponseAttributes) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SLOCorrectionResponseAttributes) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *SLOCorrectionResponseAttributes) SetEnd(v int64) {
	o.End = &v
}

// GetSloId returns the SloId field value if set, zero value otherwise.
func (o *SLOCorrectionResponseAttributes) GetSloId() string {
	if o == nil || o.SloId == nil {
		var ret string
		return ret
	}
	return *o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionResponseAttributes) GetSloIdOk() (*string, bool) {
	if o == nil || o.SloId == nil {
		return nil, false
	}
	return o.SloId, true
}

// HasSloId returns a boolean if a field has been set.
func (o *SLOCorrectionResponseAttributes) HasSloId() bool {
	if o != nil && o.SloId != nil {
		return true
	}

	return false
}

// SetSloId gets a reference to the given string and assigns it to the SloId field.
func (o *SLOCorrectionResponseAttributes) SetSloId(v string) {
	o.SloId = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SLOCorrectionResponseAttributes) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionResponseAttributes) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SLOCorrectionResponseAttributes) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *SLOCorrectionResponseAttributes) SetStart(v int64) {
	o.Start = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *SLOCorrectionResponseAttributes) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionResponseAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *SLOCorrectionResponseAttributes) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *SLOCorrectionResponseAttributes) SetTimezone(v string) {
	o.Timezone = &v
}

func (o SLOCorrectionResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.SloId != nil {
		toSerialize["slo_id"] = o.SloId
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

func (o *SLOCorrectionResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Category    *SLOCorrectionCategory `json:"category,omitempty"`
		Creator     *Creator               `json:"creator,omitempty"`
		Description *string                `json:"description,omitempty"`
		End         *int64                 `json:"end,omitempty"`
		SloId       *string                `json:"slo_id,omitempty"`
		Start       *int64                 `json:"start,omitempty"`
		Timezone    *string                `json:"timezone,omitempty"`
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
	if v := all.Category; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Category = all.Category
	o.Creator = all.Creator
	o.Description = all.Description
	o.End = all.End
	o.SloId = all.SloId
	o.Start = all.Start
	o.Timezone = all.Timezone
	return nil
}

type NullableSLOCorrectionResponseAttributes struct {
	value *SLOCorrectionResponseAttributes
	isSet bool
}

func (v NullableSLOCorrectionResponseAttributes) Get() *SLOCorrectionResponseAttributes {
	return v.value
}

func (v *NullableSLOCorrectionResponseAttributes) Set(val *SLOCorrectionResponseAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOCorrectionResponseAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOCorrectionResponseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOCorrectionResponseAttributes(val *SLOCorrectionResponseAttributes) *NullableSLOCorrectionResponseAttributes {
	return &NullableSLOCorrectionResponseAttributes{value: val, isSet: true}
}

func (v NullableSLOCorrectionResponseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOCorrectionResponseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
