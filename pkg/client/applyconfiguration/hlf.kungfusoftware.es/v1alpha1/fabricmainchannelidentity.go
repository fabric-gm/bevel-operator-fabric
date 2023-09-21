/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelIdentityApplyConfiguration represents an declarative configuration of the FabricMainChannelIdentity type for use
// with apply.
type FabricMainChannelIdentityApplyConfiguration struct {
	SecretNamespace *string `json:"secretNamespace,omitempty"`
	SecretName      *string `json:"secretName,omitempty"`
	SecretKey       *string `json:"secretKey,omitempty"`
}

// FabricMainChannelIdentityApplyConfiguration constructs an declarative configuration of the FabricMainChannelIdentity type for use with
// apply.
func FabricMainChannelIdentity() *FabricMainChannelIdentityApplyConfiguration {
	return &FabricMainChannelIdentityApplyConfiguration{}
}

// WithSecretNamespace sets the SecretNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretNamespace field is set to the value of the last call.
func (b *FabricMainChannelIdentityApplyConfiguration) WithSecretNamespace(value string) *FabricMainChannelIdentityApplyConfiguration {
	b.SecretNamespace = &value
	return b
}

// WithSecretName sets the SecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretName field is set to the value of the last call.
func (b *FabricMainChannelIdentityApplyConfiguration) WithSecretName(value string) *FabricMainChannelIdentityApplyConfiguration {
	b.SecretName = &value
	return b
}

// WithSecretKey sets the SecretKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretKey field is set to the value of the last call.
func (b *FabricMainChannelIdentityApplyConfiguration) WithSecretKey(value string) *FabricMainChannelIdentityApplyConfiguration {
	b.SecretKey = &value
	return b
}