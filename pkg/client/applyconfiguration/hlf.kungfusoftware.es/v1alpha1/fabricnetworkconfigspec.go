// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricNetworkConfigSpecApplyConfiguration represents an declarative configuration of the FabricNetworkConfigSpec type for use
// with apply.
type FabricNetworkConfigSpecApplyConfiguration struct {
	Organization  *string                                         `json:"organization,omitempty"`
	Internal      *bool                                           `json:"internal,omitempty"`
	Organizations []string                                        `json:"organizations,omitempty"`
	Namespaces    []string                                        `json:"namespaces,omitempty"`
	Channels      []string                                        `json:"channels,omitempty"`
	Identities    []FabricNetworkConfigIdentityApplyConfiguration `json:"identities,omitempty"`
	SecretName    *string                                         `json:"secretName,omitempty"`
}

// FabricNetworkConfigSpecApplyConfiguration constructs an declarative configuration of the FabricNetworkConfigSpec type for use with
// apply.
func FabricNetworkConfigSpec() *FabricNetworkConfigSpecApplyConfiguration {
	return &FabricNetworkConfigSpecApplyConfiguration{}
}

// WithOrganization sets the Organization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Organization field is set to the value of the last call.
func (b *FabricNetworkConfigSpecApplyConfiguration) WithOrganization(value string) *FabricNetworkConfigSpecApplyConfiguration {
	b.Organization = &value
	return b
}

// WithInternal sets the Internal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Internal field is set to the value of the last call.
func (b *FabricNetworkConfigSpecApplyConfiguration) WithInternal(value bool) *FabricNetworkConfigSpecApplyConfiguration {
	b.Internal = &value
	return b
}

// WithOrganizations adds the given value to the Organizations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Organizations field.
func (b *FabricNetworkConfigSpecApplyConfiguration) WithOrganizations(values ...string) *FabricNetworkConfigSpecApplyConfiguration {
	for i := range values {
		b.Organizations = append(b.Organizations, values[i])
	}
	return b
}

// WithNamespaces adds the given value to the Namespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Namespaces field.
func (b *FabricNetworkConfigSpecApplyConfiguration) WithNamespaces(values ...string) *FabricNetworkConfigSpecApplyConfiguration {
	for i := range values {
		b.Namespaces = append(b.Namespaces, values[i])
	}
	return b
}

// WithChannels adds the given value to the Channels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Channels field.
func (b *FabricNetworkConfigSpecApplyConfiguration) WithChannels(values ...string) *FabricNetworkConfigSpecApplyConfiguration {
	for i := range values {
		b.Channels = append(b.Channels, values[i])
	}
	return b
}

// WithIdentities adds the given value to the Identities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Identities field.
func (b *FabricNetworkConfigSpecApplyConfiguration) WithIdentities(values ...*FabricNetworkConfigIdentityApplyConfiguration) *FabricNetworkConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithIdentities")
		}
		b.Identities = append(b.Identities, *values[i])
	}
	return b
}

// WithSecretName sets the SecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretName field is set to the value of the last call.
func (b *FabricNetworkConfigSpecApplyConfiguration) WithSecretName(value string) *FabricNetworkConfigSpecApplyConfiguration {
	b.SecretName = &value
	return b
}
