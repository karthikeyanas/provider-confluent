/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BasicObservation struct {
}

type BasicParameters struct {
}

type ClusterObservation struct {

	// API Version defines the schema version of this representation of a Kafka cluster.
	APIVersion *string `json:"apiVersion,omitempty" tf:"api_version,omitempty"`

	// The bootstrap endpoint used by Kafka clients to connect to the Kafka cluster.
	BootstrapEndpoint *string `json:"bootstrapEndpoint,omitempty" tf:"bootstrap_endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Kind defines the object Kafka cluster represents.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The Confluent Resource Name of the Kafka cluster suitable for confluent_role_binding's crn_pattern.
	RbacCrn *string `json:"rbacCrn,omitempty" tf:"rbac_crn,omitempty"`

	// The REST endpoint of the Kafka cluster.
	RestEndpoint *string `json:"restEndpoint,omitempty" tf:"rest_endpoint,omitempty"`
}

type ClusterParameters struct {

	// The availability zone configuration of the Kafka cluster.
	// +kubebuilder:validation:Required
	Availability *string `json:"availability" tf:"availability,omitempty"`

	// +kubebuilder:validation:Optional
	Basic []BasicParameters `json:"basic,omitempty" tf:"basic,omitempty"`

	// The cloud service provider that runs the Kafka cluster.
	// +kubebuilder:validation:Required
	Cloud *string `json:"cloud" tf:"cloud,omitempty"`

	// +kubebuilder:validation:Optional
	Dedicated []DedicatedParameters `json:"dedicated,omitempty" tf:"dedicated,omitempty"`

	// The name of the Kafka cluster.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	// +kubebuilder:validation:Required
	Environment []EnvironmentParameters `json:"environment" tf:"environment,omitempty"`

	// Network represents a network (VPC) in Confluent Cloud. All Networks exist within Confluent-managed cloud provider accounts.
	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// The cloud service provider region where the Kafka cluster is running.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Standard []StandardParameters `json:"standard,omitempty" tf:"standard,omitempty"`
}

type DedicatedObservation struct {
}

type DedicatedParameters struct {

	// The number of Confluent Kafka Units (CKUs) for Dedicated cluster types. MULTI_ZONE dedicated clusters must have at least two CKUs.
	// +kubebuilder:validation:Required
	Cku *float64 `json:"cku" tf:"cku,omitempty"`

	// The ID of the encryption key that is used to encrypt the data in the Kafka cluster.
	// +kubebuilder:validation:Optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`
}

type EnvironmentObservation struct {
}

type EnvironmentParameters struct {

	// The unique identifier for the environment.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type NetworkObservation struct {
}

type NetworkParameters struct {

	// The unique identifier for the network.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type StandardObservation struct {
}

type StandardParameters struct {
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
