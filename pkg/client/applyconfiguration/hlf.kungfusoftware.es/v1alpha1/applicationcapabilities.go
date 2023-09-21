/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// ApplicationCapabilitiesApplyConfiguration represents an declarative configuration of the ApplicationCapabilities type for use
// with apply.
type ApplicationCapabilitiesApplyConfiguration struct {
	V2_0 *bool `json:"V2_0,omitempty"`
}

// ApplicationCapabilitiesApplyConfiguration constructs an declarative configuration of the ApplicationCapabilities type for use with
// apply.
func ApplicationCapabilities() *ApplicationCapabilitiesApplyConfiguration {
	return &ApplicationCapabilitiesApplyConfiguration{}
}

// WithV2_0 sets the V2_0 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V2_0 field is set to the value of the last call.
func (b *ApplicationCapabilitiesApplyConfiguration) WithV2_0(value bool) *ApplicationCapabilitiesApplyConfiguration {
	b.V2_0 = &value
	return b
}