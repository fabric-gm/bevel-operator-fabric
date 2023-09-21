/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// CsrApplyConfiguration represents an declarative configuration of the Csr type for use
// with apply.
type CsrApplyConfiguration struct {
	Hosts []string `json:"hosts,omitempty"`
	CN    *string  `json:"cn,omitempty"`
}

// CsrApplyConfiguration constructs an declarative configuration of the Csr type for use with
// apply.
func Csr() *CsrApplyConfiguration {
	return &CsrApplyConfiguration{}
}

// WithHosts adds the given value to the Hosts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Hosts field.
func (b *CsrApplyConfiguration) WithHosts(values ...string) *CsrApplyConfiguration {
	for i := range values {
		b.Hosts = append(b.Hosts, values[i])
	}
	return b
}

// WithCN sets the CN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CN field is set to the value of the last call.
func (b *CsrApplyConfiguration) WithCN(value string) *CsrApplyConfiguration {
	b.CN = &value
	return b
}