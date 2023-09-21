/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricOperatorAPIAuthApplyConfiguration represents an declarative configuration of the FabricOperatorAPIAuth type for use
// with apply.
type FabricOperatorAPIAuthApplyConfiguration struct {
	OIDCJWKS      *string `json:"oidcJWKS,omitempty"`
	OIDCIssuer    *string `json:"oidcIssuer,omitempty"`
	OIDCAuthority *string `json:"oidcAuthority,omitempty"`
	OIDCClientId  *string `json:"oidcClientId,omitempty"`
	OIDCScope     *string `json:"oidcScope,omitempty"`
}

// FabricOperatorAPIAuthApplyConfiguration constructs an declarative configuration of the FabricOperatorAPIAuth type for use with
// apply.
func FabricOperatorAPIAuth() *FabricOperatorAPIAuthApplyConfiguration {
	return &FabricOperatorAPIAuthApplyConfiguration{}
}

// WithOIDCJWKS sets the OIDCJWKS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCJWKS field is set to the value of the last call.
func (b *FabricOperatorAPIAuthApplyConfiguration) WithOIDCJWKS(value string) *FabricOperatorAPIAuthApplyConfiguration {
	b.OIDCJWKS = &value
	return b
}

// WithOIDCIssuer sets the OIDCIssuer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCIssuer field is set to the value of the last call.
func (b *FabricOperatorAPIAuthApplyConfiguration) WithOIDCIssuer(value string) *FabricOperatorAPIAuthApplyConfiguration {
	b.OIDCIssuer = &value
	return b
}

// WithOIDCAuthority sets the OIDCAuthority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCAuthority field is set to the value of the last call.
func (b *FabricOperatorAPIAuthApplyConfiguration) WithOIDCAuthority(value string) *FabricOperatorAPIAuthApplyConfiguration {
	b.OIDCAuthority = &value
	return b
}

// WithOIDCClientId sets the OIDCClientId field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCClientId field is set to the value of the last call.
func (b *FabricOperatorAPIAuthApplyConfiguration) WithOIDCClientId(value string) *FabricOperatorAPIAuthApplyConfiguration {
	b.OIDCClientId = &value
	return b
}

// WithOIDCScope sets the OIDCScope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCScope field is set to the value of the last call.
func (b *FabricOperatorAPIAuthApplyConfiguration) WithOIDCScope(value string) *FabricOperatorAPIAuthApplyConfiguration {
	b.OIDCScope = &value
	return b
}