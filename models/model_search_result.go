// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SearchResult struct {
	ValidityPeriod       int32       `json:"validityPeriod,omitempty" yaml:"validityPeriod" bson:"validityPeriod" mapstructure:"ValidityPeriod"`
	NfInstances          []NfProfile `json:"nfInstances" yaml:"nfInstances" bson:"nfInstances" mapstructure:"NfInstances"`
	NrfSupportedFeatures string      `json:"nrfSupportedFeatures,omitempty" yaml:"nrfSupportedFeatures" bson:"nrfSupportedFeatures" mapstructure:"NrfSupportedFeatures"`
}
