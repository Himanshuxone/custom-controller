package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// nvme struct will handle the nvme storage details
type Nvme struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NvmeSpec   `json:"spec,omitempty"`
	Status            NvmeStatus `json:"status,omitempty"`
}

// NvmeSpec will contains subsystem, controller and namespace details of memory pages in blocks
type NvmeSpec struct {
	Subsystem  string `json:"subsys,omitempty"`
	Controller string `json:"controller,omitempty"`
	NameSpace  string `json:"namespace,omitempty"`
}

type NvmeStatus struct {
	Attached bool `json:"attached,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NvmeList will contains list of nvme storage
type NvmeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Nvme `json:"items"`
}
