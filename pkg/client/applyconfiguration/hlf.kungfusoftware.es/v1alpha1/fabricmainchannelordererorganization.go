/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelOrdererOrganizationApplyConfiguration represents an declarative configuration of the FabricMainChannelOrdererOrganization type for use
// with apply.
type FabricMainChannelOrdererOrganizationApplyConfiguration struct {
	MSPID                  *string                                                  `json:"mspID,omitempty"`
	CAName                 *string                                                  `json:"caName,omitempty"`
	CANamespace            *string                                                  `json:"caNamespace,omitempty"`
	TLSCACert              *string                                                  `json:"tlsCACert,omitempty"`
	SignCACert             *string                                                  `json:"signCACert,omitempty"`
	OrdererEndpoints       []string                                                 `json:"ordererEndpoints,omitempty"`
	OrderersToJoin         []FabricMainChannelOrdererNodeApplyConfiguration         `json:"orderersToJoin,omitempty"`
	ExternalOrderersToJoin []FabricMainChannelExternalOrdererNodeApplyConfiguration `json:"externalOrderersToJoin,omitempty"`
}

// FabricMainChannelOrdererOrganizationApplyConfiguration constructs an declarative configuration of the FabricMainChannelOrdererOrganization type for use with
// apply.
func FabricMainChannelOrdererOrganization() *FabricMainChannelOrdererOrganizationApplyConfiguration {
	return &FabricMainChannelOrdererOrganizationApplyConfiguration{}
}

// WithMSPID sets the MSPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MSPID field is set to the value of the last call.
func (b *FabricMainChannelOrdererOrganizationApplyConfiguration) WithMSPID(value string) *FabricMainChannelOrdererOrganizationApplyConfiguration {
	b.MSPID = &value
	return b
}

// WithCAName sets the CAName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CAName field is set to the value of the last call.
func (b *FabricMainChannelOrdererOrganizationApplyConfiguration) WithCAName(value string) *FabricMainChannelOrdererOrganizationApplyConfiguration {
	b.CAName = &value
	return b
}

// WithCANamespace sets the CANamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CANamespace field is set to the value of the last call.
func (b *FabricMainChannelOrdererOrganizationApplyConfiguration) WithCANamespace(value string) *FabricMainChannelOrdererOrganizationApplyConfiguration {
	b.CANamespace = &value
	return b
}

// WithTLSCACert sets the TLSCACert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSCACert field is set to the value of the last call.
func (b *FabricMainChannelOrdererOrganizationApplyConfiguration) WithTLSCACert(value string) *FabricMainChannelOrdererOrganizationApplyConfiguration {
	b.TLSCACert = &value
	return b
}

// WithSignCACert sets the SignCACert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignCACert field is set to the value of the last call.
func (b *FabricMainChannelOrdererOrganizationApplyConfiguration) WithSignCACert(value string) *FabricMainChannelOrdererOrganizationApplyConfiguration {
	b.SignCACert = &value
	return b
}

// WithOrdererEndpoints adds the given value to the OrdererEndpoints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OrdererEndpoints field.
func (b *FabricMainChannelOrdererOrganizationApplyConfiguration) WithOrdererEndpoints(values ...string) *FabricMainChannelOrdererOrganizationApplyConfiguration {
	for i := range values {
		b.OrdererEndpoints = append(b.OrdererEndpoints, values[i])
	}
	return b
}

// WithOrderersToJoin adds the given value to the OrderersToJoin field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OrderersToJoin field.
func (b *FabricMainChannelOrdererOrganizationApplyConfiguration) WithOrderersToJoin(values ...*FabricMainChannelOrdererNodeApplyConfiguration) *FabricMainChannelOrdererOrganizationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOrderersToJoin")
		}
		b.OrderersToJoin = append(b.OrderersToJoin, *values[i])
	}
	return b
}

// WithExternalOrderersToJoin adds the given value to the ExternalOrderersToJoin field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalOrderersToJoin field.
func (b *FabricMainChannelOrdererOrganizationApplyConfiguration) WithExternalOrderersToJoin(values ...*FabricMainChannelExternalOrdererNodeApplyConfiguration) *FabricMainChannelOrdererOrganizationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExternalOrderersToJoin")
		}
		b.ExternalOrderersToJoin = append(b.ExternalOrderersToJoin, *values[i])
	}
	return b
}
