// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type PduSessionCreatedData struct {
	PduSessionType                PduSessionType                `json:"pduSessionType"`
	SscMode                       string                        `json:"sscMode"`
	HcnTunnelInfo                 *TunnelInfo                   `json:"hcnTunnelInfo"`
	SessionAmbr                   *Ambr                         `json:"sessionAmbr"`
	QosFlowsSetupList             []QosFlowSetupItem            `json:"qosFlowsSetupList"`
	HSmfInstanceId                string                        `json:"hSmfInstanceId"`
	PduSessionId                  int32                         `json:"pduSessionId,omitempty"`
	SNssai                        *Snssai                       `json:"sNssai,omitempty"`
	EnablePauseCharging           bool                          `json:"enablePauseCharging,omitempty"`
	UeIpv4Address                 string                        `json:"ueIpv4Address,omitempty"`
	UeIpv6Prefix                  string                        `json:"ueIpv6Prefix,omitempty"`
	N1SmInfoToUe                  *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	EpsPdnCnxInfo                 *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	EpsBearerInfo                 *[]EpsBearerInfo              `json:"epsBearerInfo,omitempty"`
	SupportedFeatures             string                        `json:"supportedFeatures,omitempty"`
	MaxIntegrityProtectedDataRate MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRate,omitempty"`
	AlwaysOnGranted               bool                          `json:"alwaysOnGranted,omitempty"`
	UpSecurity                    *UpSecurity                   `json:"upSecurity,omitempty"`
	RoamingChargingProfile        *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	HSmfServiceInstanceId         string                        `json:"hSmfServiceInstanceId,omitempty"`
	RecoveryTime                  *time.Time                    `json:"recoveryTime,omitempty"`
}
