/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1beta1 "k8s.io/api/networking/v1beta1"
)

// IngressApplyConfiguration represents an declarative configuration of the Ingress type for use
// with apply.
type IngressApplyConfiguration struct {
	Enabled     *bool                           `json:"enabled,omitempty"`
	ClassName   *string                         `json:"className,omitempty"`
	Annotations map[string]string               `json:"annotations,omitempty"`
	TLS         []v1beta1.IngressTLS            `json:"tls,omitempty"`
	Hosts       []IngressHostApplyConfiguration `json:"hosts,omitempty"`
}

// IngressApplyConfiguration constructs an declarative configuration of the Ingress type for use with
// apply.
func Ingress() *IngressApplyConfiguration {
	return &IngressApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *IngressApplyConfiguration) WithEnabled(value bool) *IngressApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithClassName sets the ClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClassName field is set to the value of the last call.
func (b *IngressApplyConfiguration) WithClassName(value string) *IngressApplyConfiguration {
	b.ClassName = &value
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *IngressApplyConfiguration) WithAnnotations(entries map[string]string) *IngressApplyConfiguration {
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithTLS adds the given value to the TLS field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TLS field.
func (b *IngressApplyConfiguration) WithTLS(values ...v1beta1.IngressTLS) *IngressApplyConfiguration {
	for i := range values {
		b.TLS = append(b.TLS, values[i])
	}
	return b
}

// WithHosts adds the given value to the Hosts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Hosts field.
func (b *IngressApplyConfiguration) WithHosts(values ...*IngressHostApplyConfiguration) *IngressApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHosts")
		}
		b.Hosts = append(b.Hosts, *values[i])
	}
	return b
}