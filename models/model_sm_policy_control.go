// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SmPolicyControl struct {
	Context *SmPolicyContextData `json:"context" yaml:"context" bson:"context" mapstructure:"Context"`
	Policy  *SmPolicyDecision    `json:"policy" yaml:"policy" bson:"policy" mapstructure:"Policy"`
}
