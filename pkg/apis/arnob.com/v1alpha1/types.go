package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Messi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MessiSpec   `json:"spec,omitempty"`
	Status MessiStatus `json:"status,omitempty"`
}

type MessiStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

type MessiSpec struct {
	DeploymentName string `json:"deploymentName"`
	DeploymentImage string `json:"deploymentImage"`
	Replicas       *int32 `json:"replicas"`
	ServiceName 	string `json:"serviceName"`
	ServicePort 	string `json:"servicePort"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MessiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Messi `json:"items,omitempty"`
}
