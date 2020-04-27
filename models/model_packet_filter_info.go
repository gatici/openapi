/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PacketFilterInfo struct {
	// An identifier of packet filter.
	PackFiltId string `json:"packFiltId,omitempty" yaml:"packFiltId" bson:"packFiltId" mapstructure:"PackFiltId"`
	// Defines a packet filter for an IP flow.Refer to subclause 5.3.54 of 3GPP TS 29.212 [23] for encoding.
	PackFiltCont string `json:"packFiltCont,omitempty" yaml:"packFiltCont" bson:"packFiltCont" mapstructure:"PackFiltCont"`
	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass string `json:"tosTrafficClass,omitempty" yaml:"tosTrafficClass" bson:"tosTrafficClass" mapstructure:"TosTrafficClass"`
	// The security parameter index of the IPSec packet.
	Spi string `json:"spi,omitempty" yaml:"spi" bson:"spi" mapstructure:"Spi"`
	// The Ipv6 flow label header field.
	FlowLabel     string        `json:"flowLabel,omitempty" yaml:"flowLabel" bson:"flowLabel" mapstructure:"FlowLabel"`
	FlowDirection FlowDirection `json:"flowDirection,omitempty" yaml:"flowDirection" bson:"flowDirection" mapstructure:"FlowDirection"`
}
