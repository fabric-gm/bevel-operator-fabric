/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelExternalOrdererOrganizationApplyConfiguration represents an declarative configuration of the FabricMainChannelExternalOrdererOrganization type for use
// with apply.
type FabricMainChannelExternalOrdererOrganizationApplyConfiguration struct {
	MSPID            *string  `json:"mspID,omitempty"`
	TLSRootCert      *string  `json:"tlsRootCert,omitempty"`
	SignRootCert     *string  `json:"signRootCert,omitempty"`
	OrdererEndpoints []string `json:"ordererEndpoints,omitempty"`
}

// FabricMainChannelExternalOrdererOrganizationApplyConfiguration constructs an declarative configuration of the FabricMainChannelExternalOrdererOrganization type for use with
// apply.
func FabricMainChannelExternalOrdererOrganization() *FabricMainChannelExternalOrdererOrganizationApplyConfiguration {
	return &FabricMainChannelExternalOrdererOrganizationApplyConfiguration{}
}

// WithMSPID sets the MSPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MSPID field is set to the value of the last call.
func (b *FabricMainChannelExternalOrdererOrganizationApplyConfiguration) WithMSPID(value string) *FabricMainChannelExternalOrdererOrganizationApplyConfiguration {
	b.MSPID = &value
	return b
}

// WithTLSRootCert sets the TLSRootCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSRootCert field is set to the value of the last call.
func (b *FabricMainChannelExternalOrdererOrganizationApplyConfiguration) WithTLSRootCert(value string) *FabricMainChannelExternalOrdererOrganizationApplyConfiguration {
	b.TLSRootCert = &value
	return b
}

// WithSignRootCert sets the SignRootCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignRootCert field is set to the value of the last call.
func (b *FabricMainChannelExternalOrdererOrganizationApplyConfiguration) WithSignRootCert(value string) *FabricMainChannelExternalOrdererOrganizationApplyConfiguration {
	b.SignRootCert = &value
	return b
}

// WithOrdererEndpoints adds the given value to the OrdererEndpoints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OrdererEndpoints field.
func (b *FabricMainChannelExternalOrdererOrganizationApplyConfiguration) WithOrdererEndpoints(values ...string) *FabricMainChannelExternalOrdererOrganizationApplyConfiguration {
	for i := range values {
		b.OrdererEndpoints = append(b.OrdererEndpoints, values[i])
	}
	return b
}