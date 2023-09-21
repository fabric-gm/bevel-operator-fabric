/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricPeerExternalCouchDBApplyConfiguration represents an declarative configuration of the FabricPeerExternalCouchDB type for use
// with apply.
type FabricPeerExternalCouchDBApplyConfiguration struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Host    *string `json:"host,omitempty"`
	Port    *int    `json:"port,omitempty"`
}

// FabricPeerExternalCouchDBApplyConfiguration constructs an declarative configuration of the FabricPeerExternalCouchDB type for use with
// apply.
func FabricPeerExternalCouchDB() *FabricPeerExternalCouchDBApplyConfiguration {
	return &FabricPeerExternalCouchDBApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *FabricPeerExternalCouchDBApplyConfiguration) WithEnabled(value bool) *FabricPeerExternalCouchDBApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *FabricPeerExternalCouchDBApplyConfiguration) WithHost(value string) *FabricPeerExternalCouchDBApplyConfiguration {
	b.Host = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *FabricPeerExternalCouchDBApplyConfiguration) WithPort(value int) *FabricPeerExternalCouchDBApplyConfiguration {
	b.Port = &value
	return b
}