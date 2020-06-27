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
type ApisixServiceRef struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// ApisixUpstreamsSpec defines the desired state of ApisixUpstreams
type ApisixUpstreamsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ApisixUpstreams. Edit ApisixUpstreams_types.go to remove/update
	ServiceRef 		ApisixServiceRef	`json:"serviceRef"`
	Type 			string 				`json:"type"`
}

type ApisixUpstreamNode struct {
	Host 			string 				`json:"host"`
	Port 			int64        		`json:"port"`
	Weight 			int64     			`json:"weight"`
}

type ApisixUpstreamTimeout struct {
	Connect 		int64           	`json:"connect"`
	Send 			int64 				`json:"send"`
	Read 			int64				`json:"read"`
}

// ApisixUpstreamsStatus defines the observed state of ApisixUpstreams
type ApisixUpstreamsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Nodes 			[]ApisixUpstreamNode 	`json:"nodes"`
	Retries 		int64 					`json:"retries"`
	Type 			string 					`json:"type"`
	Timeout			ApisixUpstreamTimeout	`json:"timeout"`
}

// +kubebuilder:object:root=true

// ApisixUpstreams is the Schema for the apisixupstreams API
type ApisixUpstreams struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApisixUpstreamsSpec   `json:"spec,omitempty"`
	Status ApisixUpstreamsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApisixUpstreamsList contains a list of ApisixUpstreams
type ApisixUpstreamsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApisixUpstreams `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApisixUpstreams{}, &ApisixUpstreamsList{})
}
