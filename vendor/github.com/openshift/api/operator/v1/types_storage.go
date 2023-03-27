package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Storage provides a means to configure an operator to manage the cluster storage operator. `cluster` is the canonical name.
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type Storage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +kubebuilder:validation:Required
	// +required
	Spec StorageSpec `json:"spec"`

	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status StorageStatus `json:"status"`
}

// CSIMigrationToggle indicates whether CSI migration should be enabled for drivers where it is optional.
// +kubebuilder:validation:Enum=Enabled
type CSIMigrationToggle string

const (
	CSIMigrationEnabled CSIMigrationToggle = "Enabled"
)

// StorageSpec is the specification of the desired behavior of the cluster storage operator.
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.csiMigration) || has(self.csiMigration)", message="CSIMigration is required once set"
type StorageSpec struct {
	OperatorSpec `json:",inline"`

	// CSIMigration enables CSI migration for drivers where it is optional.
	// This field is immutable once it is set and can not be undone.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="CSIMigration is immutable"
	// +optional
	CSIMigration CSIMigrationToggle `json:"csiMigration,omitempty"`
}

// StorageStatus defines the observed status of the cluster storage operator.
type StorageStatus struct {
	OperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageList contains a list of Storages.
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type StorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Storage `json:"items"`
}
