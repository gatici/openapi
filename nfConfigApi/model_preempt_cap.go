// SPDX-FileCopyrightText: 2025 Canonical Ltd
//
// SPDX-License-Identifier: Apache-2.0
//

/*
WebConsole NFConfig API

API for managing access, mobility, policy, session and PLMN information.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nfConfigApi

import (
	"encoding/json"
	"fmt"
)

// PreemptCap the model 'PreemptCap'
type PreemptCap string

// List of PreemptCap
const (
	PREEMPTCAP_NOT_PREEMPT PreemptCap = "NOT_PREEMPT"
	PREEMPTCAP_MAY_PREEMPT PreemptCap = "MAY_PREEMPT"
)

// All allowed values of PreemptCap enum
var AllowedPreemptCapEnumValues = []PreemptCap{
	"NOT_PREEMPT",
	"MAY_PREEMPT",
}

func (v *PreemptCap) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PreemptCap(value)
	for _, existing := range AllowedPreemptCapEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PreemptCap", value)
}

// NewPreemptCapFromValue returns a pointer to a valid PreemptCap
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPreemptCapFromValue(v string) (*PreemptCap, error) {
	ev := PreemptCap(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PreemptCap: valid values are %v", v, AllowedPreemptCapEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PreemptCap) IsValid() bool {
	for _, existing := range AllowedPreemptCapEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PreemptCap value
func (v PreemptCap) Ptr() *PreemptCap {
	return &v
}

type NullablePreemptCap struct {
	value *PreemptCap
	isSet bool
}

func (v NullablePreemptCap) Get() *PreemptCap {
	return v.value
}

func (v *NullablePreemptCap) Set(val *PreemptCap) {
	v.value = val
	v.isSet = true
}

func (v NullablePreemptCap) IsSet() bool {
	return v.isSet
}

func (v *NullablePreemptCap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreemptCap(val *PreemptCap) *NullablePreemptCap {
	return &NullablePreemptCap{value: val, isSet: true}
}

func (v NullablePreemptCap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreemptCap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
