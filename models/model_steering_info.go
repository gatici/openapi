// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SteeringInfo struct {
	PlmnId         *PlmnId      `json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
	AccessTechList []AccessTech `json:"accessTechList,omitempty" yaml:"accessTechList" bson:"accessTechList" mapstructure:"AccessTechList"`
}
