/*
Copyright 2021.

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

package v1alpha1

import (
	"github.com/k8ssandra/k8ssandra-operator/pkg/images"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	ReaperLabel     = "k8ssandra.io/reaper"
	DefaultKeyspace = "reaper_db"
)

type ReaperDatacenterTemplate struct {

	// The image to use for the Reaper pod main container.
	// The default is "thelastpickle/cassandra-reaper:3.1.0".
	// +optional
	// +kubebuilder:default={repository:"thelastpickle",name:"cassandra-reaper",tag:"3.1.0"}
	ContainerImage *images.Image `json:"containerImage,omitempty"`

	// The image to use for tje Reaper pod init container (that performs schema migrations).
	// The default is "thelastpickle/cassandra-reaper:3.1.0".
	// +optional
	// +kubebuilder:default={repository:"thelastpickle",name:"cassandra-reaper",tag:"3.1.0"}
	InitContainerImage *images.Image `json:"initContainerImage,omitempty"`

	// +kubebuilder:default="default"
	// +optional
	ServiceAccountName string `json:"ServiceAccountName,omitempty"`

	// Auto scheduling properties. When you enable the auto-schedule feature, Reaper dynamically schedules repairs for
	// all non-system keyspaces in a cluster. A cluster's keyspaces are monitored and any modifications (additions or
	// removals) are detected. When a new keyspace is created, a new repair schedule is created automatically for that
	// keyspace. Conversely, when a keyspace is removed, the corresponding repair schedule is deleted.
	// +optional
	AutoScheduling AutoScheduling `json:"autoScheduling,omitempty"`

	// LivenessProbe sets the Reaper liveness probe. Leave nil to use defaults.
	// +optional
	LivenessProbe *corev1.Probe `json:"livenessProbe,omitempty"`

	// ReadinessProbe sets the Reaper readiness probe. Leave nil to use defaults.
	// +optional
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty"`

	// Affinity applied to the Reaper pods.
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// Tolerations applied to the Reaper pods.
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// PodSecurityContext contains a pod-level SecurityContext to apply to Reaper pods.
	// +optional
	PodSecurityContext *corev1.PodSecurityContext `json:"podSecurityContext,omitempty"`

	// SecurityContext applied to the Reaper main container.
	// +optional
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`

	// InitContainerSecurityContext is the SecurityContext applied to the Reaper init container, used to perform schema
	// migrations.
	// +optional
	InitContainerSecurityContext *corev1.SecurityContext `json:"initContainerSecurityContext,omitempty"`
}

// AutoScheduling includes options to configure the auto scheduling of repairs for new clusters.
type AutoScheduling struct {

	// +optional
	// +kubebuilder:default=false
	Enabled bool `json:"enabled,omitempty"`

	// RepairType is the type of repair to create:
	// - REGULAR creates a regular repair (non-adaptive and non-incremental);
	// - ADAPTIVE creates an adaptive repair; adaptive repairs are most suited for Cassandra 3.
	// - INCREMENTAL creates an incremental repair; incremental repairs should only be used with Cassandra 4+.
	// - AUTO chooses between ADAPTIVE and INCREMENTAL depending on the Cassandra server version; ADAPTIVE for Cassandra
	// 3 and INCREMENTAL for Cassandra 4+.
	// +optional
	// +kubebuilder:default="AUTO"
	// +kubebuilder:validation:Enum:=REGULAR;ADAPTIVE;INCREMENTAL;AUTO
	RepairType string `json:"repairType,omitempty"`

	// PercentUnrepairedThreshold is the percentage of unrepaired data over which an incremental repair should be
	// started. Only relevant when using repair type INCREMENTAL.
	// +optional
	// +kubebuilder:default=10
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	PercentUnrepairedThreshold int `json:"percentUnrepairedThreshold,omitempty"`

	// InitialDelay is the amount of delay time before the schedule period starts. Must be a valid ISO-8601 duration
	// string. The default is "PT15S" (15 seconds).
	// +optional
	// +kubebuilder:default="PT15S"
	// +kubebuilder:validation:Pattern:="([-+]?)P(?:([-+]?[0-9]+)D)?(T(?:([-+]?[0-9]+)H)?(?:([-+]?[0-9]+)M)?(?:([-+]?[0-9]+)(?:[.,]([0-9]{0,9}))?S)?)?"
	InitialDelay string `json:"initialDelayPeriod,omitempty"`

	// PeriodBetweenPolls is the interval time to wait before checking whether to start a repair task. Must be a valid
	// ISO-8601 duration string. The default is "PT10M" (10 minutes).
	// +optional
	// +kubebuilder:default="PT10M"
	// +kubebuilder:validation:Pattern:="([-+]?)P(?:([-+]?[0-9]+)D)?(T(?:([-+]?[0-9]+)H)?(?:([-+]?[0-9]+)M)?(?:([-+]?[0-9]+)(?:[.,]([0-9]{0,9}))?S)?)?"
	PeriodBetweenPolls string `json:"periodBetweenPolls,omitempty"`

	// TimeBeforeFirstSchedule is the grace period before the first repair in the schedule is started. Must be a valid
	// ISO-8601 duration string. The default is "PT5M" (5 minutes).
	// +optional
	// +kubebuilder:default="PT5M"
	// +kubebuilder:validation:Pattern:="([-+]?)P(?:([-+]?[0-9]+)D)?(T(?:([-+]?[0-9]+)H)?(?:([-+]?[0-9]+)M)?(?:([-+]?[0-9]+)(?:[.,]([0-9]{0,9}))?S)?)?"
	TimeBeforeFirstSchedule string `json:"timeBeforeFirstSchedule,omitempty"`

	// ScheduleSpreadPeriod is the time spacing between each of the repair schedules that is to be carried out. Must be
	// a valid ISO-8601 duration string. The default is "PT6H" (6 hours).
	// +optional
	// +kubebuilder:default="PT6H"
	// +kubebuilder:validation:Pattern:="([-+]?)P(?:([-+]?[0-9]+)D)?(T(?:([-+]?[0-9]+)H)?(?:([-+]?[0-9]+)M)?(?:([-+]?[0-9]+)(?:[.,]([0-9]{0,9}))?S)?)?"
	ScheduleSpreadPeriod string `json:"scheduleSpreadPeriod,omitempty"`

	// ExcludedClusters are the clusters that are to be excluded from the repair schedule.
	// +optional
	ExcludedClusters []string `json:"excludedClusters,omitempty"`

	// ExcludedKeyspaces are the keyspaces that are to be excluded from the repair schedule.
	// +optional
	ExcludedKeyspaces []string `json:"excludedKeyspaces,omitempty"`
}

type ReaperClusterTemplate struct {
	ReaperDatacenterTemplate `json:",inline"`

	// The keyspace to use to store Reaper's state. Will default to "reaper_db" if unspecified. Will be created if it
	// does not exist, and if this Reaper resource is managed by K8ssandra.
	// +kubebuilder:default="reaper_db"
	// +optional
	Keyspace string `json:"keyspace,omitempty"`

	// Defines the username and password that Reaper will use to authenticate CQL connections to Cassandra clusters.
	// These credentials will be automatically turned into CQL roles by cass-operator when bootstrapping the datacenter,
	// then passed to the Reaper instance, so that it can authenticate against nodes in the datacenter using CQL. If CQL
	// authentication is not required, leave this field empty. The secret must be in the same namespace as Reaper itself
	// and must contain two keys: "username" and "password".
	// +optional
	CassandraUserSecretRef string `json:"cassandraUserSecretRef,omitempty"`

	// Defines the username and password that Reaper will use to authenticate JMX connections to Cassandra clusters.
	// These credentials will be automatically passed to each Cassandra node in the datacenter, as well as to the Reaper
	// instance, so that the latter can authenticate against the former. If JMX authentication is not required, leave
	// this field empty. The secret must be in the same namespace as Reaper itself and must contain two keys: "username"
	// and "password".
	// +optional
	JmxUserSecretRef string `json:"jmxUserSecretRef,omitempty"`
}

// CassandraDatacenterRef references the target Cassandra DC that Reaper should manage.
// TODO this object could be used by Stargate too; which currently cannot locate DCs outside of its own namespace.
type CassandraDatacenterRef struct {

	// The datacenter name.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// The datacenter namespace. If empty, the datacenter will be assumed to reside in the same namespace as the Reaper
	// instance.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// ReaperSpec defines the desired state of Reaper
type ReaperSpec struct {
	ReaperClusterTemplate `json:",inline"`

	// DatacenterRef is the reference of a CassandraDatacenter resource that this Reaper instance should manage. It will
	// also be used as the backend for persisting Reaper's state. Reaper must be able to access the JMX port (7199 by
	// default) and the CQL port (9042 by default) on this DC.
	// +kubebuilder:validation:Required
	DatacenterRef CassandraDatacenterRef `json:"datacenterRef"`

	// DatacenterAvailability indicates to Reaper its deployment in relation to the target datacenter's network.
	// For single-DC clusters, the default (LOCAL) is fine. For multi-DC clusters, it is recommended to use EACH,
	// provided that there is one Reaper instance managing each DC in the cluster; otherwise, if one single Reaper
	// instance is going to manage more than one DC in the cluster, use LOCAL and remote DCs will be handled internally
	// by Cassandra itself.
	// See https://cassandra-reaper.io/docs/usage/multi_dc/.
	// +optional
	// +kubebuilder:default="LOCAL"
	// +kubebuilder:validation:Enum:=LOCAL;ALL;EACH
	DatacenterAvailability string `json:"datacenterAvailability,omitempty"`
}

// ReaperProgress is a word summarizing the state of a Reaper resource.
type ReaperProgress string

const (
	// ReaperProgressPending is Reaper's status when it's waiting for the datacenter to become ready.
	ReaperProgressPending = ReaperProgress("Pending")
	// ReaperProgressDeploying is Reaper's status when it's waiting for the Reaper instance and its associated service
	// to become ready.
	ReaperProgressDeploying = ReaperProgress("Deploying")
	// ReaperProgressConfiguring is Reaper's status when the Reaper instance is ready for work and is being connected
	// its target datacenter.
	ReaperProgressConfiguring = ReaperProgress("Configuring")
	// ReaperProgressRunning is Reaper's status when Reaper is up and running.
	ReaperProgressRunning = ReaperProgress("Running")
)

type ReaperConditionType string

const (
	ReaperReady ReaperConditionType = "Ready"
)

type ReaperCondition struct {
	Type   ReaperConditionType    `json:"type"`
	Status corev1.ConditionStatus `json:"status"`

	// LastTransitionTime is the last time the condition transited from one status to another.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
}

// ReaperStatus defines the observed state of Reaper
type ReaperStatus struct {

	// Progress is the progress of this Reaper object.
	// +kubebuilder:validation:Enum=Pending;Deploying;Configuring;Running
	Progress ReaperProgress `json:"progress"`

	// +optional
	Conditions []ReaperCondition `json:"conditions,omitempty"`
}

func (in *ReaperStatus) GetConditionStatus(conditionType ReaperConditionType) corev1.ConditionStatus {
	if in != nil {
		for _, condition := range in.Conditions {
			if condition.Type == conditionType {
				return condition.Status
			}
		}
	}
	return corev1.ConditionUnknown
}

func (in *ReaperStatus) SetCondition(condition ReaperCondition) {
	for i, c := range in.Conditions {
		if c.Type == condition.Type {
			in.Conditions[i] = condition
			return
		}
	}
	in.Conditions = append(in.Conditions, condition)
}

func (in *ReaperStatus) IsReady() bool {
	return in != nil && in.GetConditionStatus(ReaperReady) == corev1.ConditionTrue
}

func (in *ReaperStatus) SetReady() {
	now := metav1.Now()
	in.SetCondition(ReaperCondition{
		Type:               ReaperReady,
		Status:             corev1.ConditionTrue,
		LastTransitionTime: &now,
	})
}

func (in *ReaperStatus) SetNotReady() {
	now := metav1.Now()
	in.SetCondition(ReaperCondition{
		Type:               ReaperReady,
		Status:             corev1.ConditionFalse,
		LastTransitionTime: &now,
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="DC",type=string,JSONPath=`.spec.datacenterRef.name`
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.progress`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Reaper is the Schema for the reapers API
type Reaper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReaperSpec   `json:"spec,omitempty"`
	Status ReaperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReaperList contains a list of Reaper
type ReaperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reaper `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Reaper{}, &ReaperList{})
}
