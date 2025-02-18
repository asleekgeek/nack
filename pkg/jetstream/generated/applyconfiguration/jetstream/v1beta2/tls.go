// Copyright 2025 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

// TLSApplyConfiguration represents a declarative configuration of the TLS type for use
// with apply.
type TLSApplyConfiguration struct {
	ClientCert *string  `json:"clientCert,omitempty"`
	ClientKey  *string  `json:"clientKey,omitempty"`
	RootCAs    []string `json:"rootCas,omitempty"`
}

// TLSApplyConfiguration constructs a declarative configuration of the TLS type for use with
// apply.
func TLS() *TLSApplyConfiguration {
	return &TLSApplyConfiguration{}
}

// WithClientCert sets the ClientCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientCert field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithClientCert(value string) *TLSApplyConfiguration {
	b.ClientCert = &value
	return b
}

// WithClientKey sets the ClientKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientKey field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithClientKey(value string) *TLSApplyConfiguration {
	b.ClientKey = &value
	return b
}

// WithRootCAs adds the given value to the RootCAs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RootCAs field.
func (b *TLSApplyConfiguration) WithRootCAs(values ...string) *TLSApplyConfiguration {
	for i := range values {
		b.RootCAs = append(b.RootCAs, values[i])
	}
	return b
}
