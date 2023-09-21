/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	status "github.com/kfsoftware/hlf-operator/pkg/status"
)

// FabricOrderingServiceStatusApplyConfiguration represents an declarative configuration of the FabricOrderingServiceStatus type for use
// with apply.
type FabricOrderingServiceStatusApplyConfiguration struct {
	Conditions *status.Conditions         `json:"conditions,omitempty"`
	Status     *v1alpha1.DeploymentStatus `json:"status,omitempty"`
}

// FabricOrderingServiceStatusApplyConfiguration constructs an declarative configuration of the FabricOrderingServiceStatus type for use with
// apply.
func FabricOrderingServiceStatus() *FabricOrderingServiceStatusApplyConfiguration {
	return &FabricOrderingServiceStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *FabricOrderingServiceStatusApplyConfiguration) WithConditions(value status.Conditions) *FabricOrderingServiceStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FabricOrderingServiceStatusApplyConfiguration) WithStatus(value v1alpha1.DeploymentStatus) *FabricOrderingServiceStatusApplyConfiguration {
	b.Status = &value
	return b
}