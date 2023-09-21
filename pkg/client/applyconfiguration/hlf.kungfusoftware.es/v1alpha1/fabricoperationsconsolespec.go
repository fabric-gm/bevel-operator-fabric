/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricOperationsConsoleSpecApplyConfiguration represents an declarative configuration of the FabricOperationsConsoleSpec type for use
// with apply.
type FabricOperationsConsoleSpecApplyConfiguration struct {
	Auth             *FabricOperationsConsoleAuthApplyConfiguration    `json:"auth,omitempty"`
	Resources        *v1.ResourceRequirements                          `json:"resources,omitempty"`
	Image            *string                                           `json:"image,omitempty"`
	Tag              *string                                           `json:"tag,omitempty"`
	ImagePullPolicy  *v1.PullPolicy                                    `json:"imagePullPolicy,omitempty"`
	Tolerations      []v1.Toleration                                   `json:"tolerations,omitempty"`
	Replicas         *int                                              `json:"replicas,omitempty"`
	CouchDB          *FabricOperationsConsoleCouchDBApplyConfiguration `json:"couchDB,omitempty"`
	Env              []v1.EnvVar                                       `json:"env,omitempty"`
	ImagePullSecrets []v1.LocalObjectReference                         `json:"imagePullSecrets,omitempty"`
	Affinity         *v1.Affinity                                      `json:"affinity,omitempty"`
	Port             *int                                              `json:"port,omitempty"`
	Config           *string                                           `json:"config,omitempty"`
	Ingress          *IngressApplyConfiguration                        `json:"ingress,omitempty"`
	HostURL          *string                                           `json:"hostUrl,omitempty"`
}

// FabricOperationsConsoleSpecApplyConfiguration constructs an declarative configuration of the FabricOperationsConsoleSpec type for use with
// apply.
func FabricOperationsConsoleSpec() *FabricOperationsConsoleSpecApplyConfiguration {
	return &FabricOperationsConsoleSpecApplyConfiguration{}
}

// WithAuth sets the Auth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Auth field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithAuth(value *FabricOperationsConsoleAuthApplyConfiguration) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Auth = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithResources(value v1.ResourceRequirements) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithImage(value string) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithTag(value string) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Tag = &value
	return b
}

// WithImagePullPolicy sets the ImagePullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullPolicy field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithImagePullPolicy(value v1.PullPolicy) *FabricOperationsConsoleSpecApplyConfiguration {
	b.ImagePullPolicy = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithTolerations(values ...v1.Toleration) *FabricOperationsConsoleSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithReplicas(value int) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithCouchDB sets the CouchDB field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CouchDB field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithCouchDB(value *FabricOperationsConsoleCouchDBApplyConfiguration) *FabricOperationsConsoleSpecApplyConfiguration {
	b.CouchDB = value
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithEnv(values ...v1.EnvVar) *FabricOperationsConsoleSpecApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithImagePullSecrets(values ...v1.LocalObjectReference) *FabricOperationsConsoleSpecApplyConfiguration {
	for i := range values {
		b.ImagePullSecrets = append(b.ImagePullSecrets, values[i])
	}
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithAffinity(value v1.Affinity) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithPort(value int) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Port = &value
	return b
}

// WithConfig sets the Config field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Config field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithConfig(value string) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Config = &value
	return b
}

// WithIngress sets the Ingress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ingress field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithIngress(value *IngressApplyConfiguration) *FabricOperationsConsoleSpecApplyConfiguration {
	b.Ingress = value
	return b
}

// WithHostURL sets the HostURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostURL field is set to the value of the last call.
func (b *FabricOperationsConsoleSpecApplyConfiguration) WithHostURL(value string) *FabricOperationsConsoleSpecApplyConfiguration {
	b.HostURL = &value
	return b
}