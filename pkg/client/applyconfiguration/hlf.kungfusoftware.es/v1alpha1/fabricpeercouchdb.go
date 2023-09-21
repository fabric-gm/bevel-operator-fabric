/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricPeerCouchDBApplyConfiguration represents an declarative configuration of the FabricPeerCouchDB type for use
// with apply.
type FabricPeerCouchDBApplyConfiguration struct {
	User            *string                                      `json:"user,omitempty"`
	Password        *string                                      `json:"password,omitempty"`
	Image           *string                                      `json:"image,omitempty"`
	Tag             *string                                      `json:"tag,omitempty"`
	PullPolicy      *v1.PullPolicy                               `json:"pullPolicy,omitempty"`
	ExternalCouchDB *FabricPeerExternalCouchDBApplyConfiguration `json:"externalCouchDB,omitempty"`
}

// FabricPeerCouchDBApplyConfiguration constructs an declarative configuration of the FabricPeerCouchDB type for use with
// apply.
func FabricPeerCouchDB() *FabricPeerCouchDBApplyConfiguration {
	return &FabricPeerCouchDBApplyConfiguration{}
}

// WithUser sets the User field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the User field is set to the value of the last call.
func (b *FabricPeerCouchDBApplyConfiguration) WithUser(value string) *FabricPeerCouchDBApplyConfiguration {
	b.User = &value
	return b
}

// WithPassword sets the Password field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Password field is set to the value of the last call.
func (b *FabricPeerCouchDBApplyConfiguration) WithPassword(value string) *FabricPeerCouchDBApplyConfiguration {
	b.Password = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *FabricPeerCouchDBApplyConfiguration) WithImage(value string) *FabricPeerCouchDBApplyConfiguration {
	b.Image = &value
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *FabricPeerCouchDBApplyConfiguration) WithTag(value string) *FabricPeerCouchDBApplyConfiguration {
	b.Tag = &value
	return b
}

// WithPullPolicy sets the PullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PullPolicy field is set to the value of the last call.
func (b *FabricPeerCouchDBApplyConfiguration) WithPullPolicy(value v1.PullPolicy) *FabricPeerCouchDBApplyConfiguration {
	b.PullPolicy = &value
	return b
}

// WithExternalCouchDB sets the ExternalCouchDB field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalCouchDB field is set to the value of the last call.
func (b *FabricPeerCouchDBApplyConfiguration) WithExternalCouchDB(value *FabricPeerExternalCouchDBApplyConfiguration) *FabricPeerCouchDBApplyConfiguration {
	b.ExternalCouchDB = value
	return b
}