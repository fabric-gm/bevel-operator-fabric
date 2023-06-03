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

// FabricIstioApplyConfiguration represents an declarative configuration of the FabricIstio type for use
// with apply.
type FabricIstioApplyConfiguration struct {
	Port           *int     `json:"port,omitempty"`
	Hosts          []string `json:"hosts,omitempty"`
	IngressGateway *string  `json:"ingressGateway,omitempty"`
}

// FabricIstioApplyConfiguration constructs an declarative configuration of the FabricIstio type for use with
// apply.
func FabricIstio() *FabricIstioApplyConfiguration {
	return &FabricIstioApplyConfiguration{}
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *FabricIstioApplyConfiguration) WithPort(value int) *FabricIstioApplyConfiguration {
	b.Port = &value
	return b
}

// WithHosts adds the given value to the Hosts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Hosts field.
func (b *FabricIstioApplyConfiguration) WithHosts(values ...string) *FabricIstioApplyConfiguration {
	for i := range values {
		b.Hosts = append(b.Hosts, values[i])
	}
	return b
}

// WithIngressGateway sets the IngressGateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IngressGateway field is set to the value of the last call.
func (b *FabricIstioApplyConfiguration) WithIngressGateway(value string) *FabricIstioApplyConfiguration {
	b.IngressGateway = &value
	return b
}
