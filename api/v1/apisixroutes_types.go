/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ApisixRoutesIngressRef struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// ApisixRoutesSpec defines the desired state of ApisixRoutes
type ApisixRoutesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ApisixRoutes. Edit ApisixRoutes_types.go to remove/update
	IngressRef ApisixRoutesIngressRef `json:"ingressRef"`
	Type       string                 `json:"type"`
}

type ApisixPluginsProxyRewrite struct {
	Uri     string            `json:"uri"`
	Host    string            `json:"host"`
	Scheme  string            `json:"scheme"`
	Headers map[string]string `json:"headers"`
}

type ApisixPluginsObject struct {
	ProxyRewrite ApisixPluginsProxyRewrite `json:"proxy-rewrite"`
}

// ApisixRoutesStatus defines the observed state of ApisixRoutes
type ApisixRoutesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Uri             string              `json:"uri"`
	Host            string              `json:"host"`
	Priority        int64               `json:"priority"`
	Methods         []string            `json:"methods"`
	Plugins         ApisixPluginsObject `json:"plugins"`
	ServiceProtocol string              `json:"service_protocol"`
	UpstreamId      string              `json:"upstream_id"`
}

// +kubebuilder:object:root=true

// ApisixRoutes is the Schema for the apisixroutes API
type ApisixRoutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApisixRoutesSpec   `json:"spec,omitempty"`
	Status ApisixRoutesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApisixRoutesList contains a list of ApisixRoutes
type ApisixRoutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApisixRoutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApisixRoutes{}, &ApisixRoutesList{})
}
