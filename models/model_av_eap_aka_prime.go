// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NudmUEAU
 *
 * UDM UE Authentication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AvEapAkaPrime struct {
	AvType  AvType `json:"avType" yaml:"avType" bson:"avType" mapstructure:"AvType"`
	Rand    string `json:"rand" yaml:"rand" bson:"rand" mapstructure:"Rand"`
	Xres    string `json:"xres" yaml:"xres" bson:"xres" mapstructure:"Xres"`
	Autn    string `json:"autn" yaml:"autn" bson:"autn" mapstructure:"Autn"`
	CkPrime string `json:"ckPrime" yaml:"ckPrime" bson:"ckPrime" mapstructure:"CkPrime"`
	IkPrime string `json:"ikPrime" yaml:"ikPrime" bson:"ikPrime" mapstructure:"IkPrime"`
}
