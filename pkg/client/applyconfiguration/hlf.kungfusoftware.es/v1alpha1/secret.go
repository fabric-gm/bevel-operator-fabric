/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// SecretApplyConfiguration represents an declarative configuration of the Secret type for use
// with apply.
type SecretApplyConfiguration struct {
	Enrollment *EnrollmentApplyConfiguration `json:"enrollment,omitempty"`
}

// SecretApplyConfiguration constructs an declarative configuration of the Secret type for use with
// apply.
func Secret() *SecretApplyConfiguration {
	return &SecretApplyConfiguration{}
}

// WithEnrollment sets the Enrollment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollment field is set to the value of the last call.
func (b *SecretApplyConfiguration) WithEnrollment(value *EnrollmentApplyConfiguration) *SecretApplyConfiguration {
	b.Enrollment = value
	return b
}