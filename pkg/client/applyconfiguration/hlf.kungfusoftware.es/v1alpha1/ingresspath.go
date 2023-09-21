/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// IngressPathApplyConfiguration represents an declarative configuration of the IngressPath type for use
// with apply.
type IngressPathApplyConfiguration struct {
	Path     *string `json:"path,omitempty"`
	PathType *string `json:"pathType,omitempty"`
}

// IngressPathApplyConfiguration constructs an declarative configuration of the IngressPath type for use with
// apply.
func IngressPath() *IngressPathApplyConfiguration {
	return &IngressPathApplyConfiguration{}
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *IngressPathApplyConfiguration) WithPath(value string) *IngressPathApplyConfiguration {
	b.Path = &value
	return b
}

// WithPathType sets the PathType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PathType field is set to the value of the last call.
func (b *IngressPathApplyConfiguration) WithPathType(value string) *IngressPathApplyConfiguration {
	b.PathType = &value
	return b
}