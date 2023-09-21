/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration represents an declarative configuration of the FabricMainChannelAdminPeerOrganizationSpec type for use
// with apply.
type FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration struct {
	MSPID *string `json:"mspID,omitempty"`
}

// FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration constructs an declarative configuration of the FabricMainChannelAdminPeerOrganizationSpec type for use with
// apply.
func FabricMainChannelAdminPeerOrganizationSpec() *FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration {
	return &FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration{}
}

// WithMSPID sets the MSPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MSPID field is set to the value of the last call.
func (b *FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration) WithMSPID(value string) *FabricMainChannelAdminPeerOrganizationSpecApplyConfiguration {
	b.MSPID = &value
	return b
}