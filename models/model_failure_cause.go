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

type FailureCause string

// List of FailureCause
const (
	FailureCause_RULE_EVENT     FailureCause = "PCC_RULE_EVENT"
	FailureCause_QOS_FLOW_EVENT FailureCause = "PCC_QOS_FLOW_EVENT"
)
