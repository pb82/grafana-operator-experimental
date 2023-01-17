//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"encoding/json"
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentV1) DeepCopyInto(out *DeploymentV1) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentV1.
func (in *DeploymentV1) DeepCopy() *DeploymentV1 {
	if in == nil {
		return nil
	}
	out := new(DeploymentV1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentV1PodSpec) DeepCopyInto(out *DeploymentV1PodSpec) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EphemeralContainers != nil {
		in, out := &in.EphemeralContainers, &out.EphemeralContainers
		*out = make([]v1.EphemeralContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AutomountServiceAccountToken != nil {
		in, out := &in.AutomountServiceAccountToken, &out.AutomountServiceAccountToken
		*out = new(bool)
		**out = **in
	}
	if in.ShareProcessNamespace != nil {
		in, out := &in.ShareProcessNamespace, &out.ShareProcessNamespace
		*out = new(bool)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]v1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int32)
		**out = **in
	}
	if in.DNSConfig != nil {
		in, out := &in.DNSConfig, &out.DNSConfig
		*out = new(v1.PodDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessGates != nil {
		in, out := &in.ReadinessGates, &out.ReadinessGates
		*out = make([]v1.PodReadinessGate, len(*in))
		copy(*out, *in)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.EnableServiceLinks != nil {
		in, out := &in.EnableServiceLinks, &out.EnableServiceLinks
		*out = new(bool)
		**out = **in
	}
	if in.PreemptionPolicy != nil {
		in, out := &in.PreemptionPolicy, &out.PreemptionPolicy
		*out = new(v1.PreemptionPolicy)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SetHostnameAsFQDN != nil {
		in, out := &in.SetHostnameAsFQDN, &out.SetHostnameAsFQDN
		*out = new(bool)
		**out = **in
	}
	if in.OS != nil {
		in, out := &in.OS, &out.OS
		*out = new(v1.PodOS)
		**out = **in
	}
	if in.HostUsers != nil {
		in, out := &in.HostUsers, &out.HostUsers
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentV1PodSpec.
func (in *DeploymentV1PodSpec) DeepCopy() *DeploymentV1PodSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentV1PodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentV1PodTemplateSpec) DeepCopyInto(out *DeploymentV1PodTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(DeploymentV1PodSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentV1PodTemplateSpec.
func (in *DeploymentV1PodTemplateSpec) DeepCopy() *DeploymentV1PodTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentV1PodTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentV1Spec) DeepCopyInto(out *DeploymentV1Spec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(DeploymentV1PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(appsv1.DeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.ProgressDeadlineSeconds != nil {
		in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentV1Spec.
func (in *DeploymentV1Spec) DeepCopy() *DeploymentV1Spec {
	if in == nil {
		return nil
	}
	out := new(DeploymentV1Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *External) DeepCopyInto(out *External) {
	*out = *in
	if in.ApiKey != nil {
		in, out := &in.ApiKey, &out.ApiKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AdminUser != nil {
		in, out := &in.AdminUser, &out.AdminUser
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AdminPassword != nil {
		in, out := &in.AdminPassword, &out.AdminPassword
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new External.
func (in *External) DeepCopy() *External {
	if in == nil {
		return nil
	}
	out := new(External)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grafana) DeepCopyInto(out *Grafana) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grafana.
func (in *Grafana) DeepCopy() *Grafana {
	if in == nil {
		return nil
	}
	out := new(Grafana)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Grafana) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaClient) DeepCopyInto(out *GrafanaClient) {
	*out = *in
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int)
		**out = **in
	}
	if in.PreferIngress != nil {
		in, out := &in.PreferIngress, &out.PreferIngress
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaClient.
func (in *GrafanaClient) DeepCopy() *GrafanaClient {
	if in == nil {
		return nil
	}
	out := new(GrafanaClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaContainer) DeepCopyInto(out *GrafanaContainer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaContainer.
func (in *GrafanaContainer) DeepCopy() *GrafanaContainer {
	if in == nil {
		return nil
	}
	out := new(GrafanaContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboard) DeepCopyInto(out *GrafanaDashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboard.
func (in *GrafanaDashboard) DeepCopy() *GrafanaDashboard {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardDatasource) DeepCopyInto(out *GrafanaDashboardDatasource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardDatasource.
func (in *GrafanaDashboardDatasource) DeepCopy() *GrafanaDashboardDatasource {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardList) DeepCopyInto(out *GrafanaDashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GrafanaDashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardList.
func (in *GrafanaDashboardList) DeepCopy() *GrafanaDashboardList {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardSpec) DeepCopyInto(out *GrafanaDashboardSpec) {
	*out = *in
	if in.InstanceSelector != nil {
		in, out := &in.InstanceSelector, &out.InstanceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make(PluginList, len(*in))
		copy(*out, *in)
	}
	out.ContentCacheDuration = in.ContentCacheDuration
	if in.Datasources != nil {
		in, out := &in.Datasources, &out.Datasources
		*out = make([]GrafanaDashboardDatasource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardSpec.
func (in *GrafanaDashboardSpec) DeepCopy() *GrafanaDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardStatus) DeepCopyInto(out *GrafanaDashboardStatus) {
	*out = *in
	if in.ContentCache != nil {
		in, out := &in.ContentCache, &out.ContentCache
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	in.ContentTimestamp.DeepCopyInto(&out.ContentTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardStatus.
func (in *GrafanaDashboardStatus) DeepCopy() *GrafanaDashboardStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDatasource) DeepCopyInto(out *GrafanaDatasource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDatasource.
func (in *GrafanaDatasource) DeepCopy() *GrafanaDatasource {
	if in == nil {
		return nil
	}
	out := new(GrafanaDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDatasource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDatasourceInternal) DeepCopyInto(out *GrafanaDatasourceInternal) {
	*out = *in
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(int64)
		**out = **in
	}
	if in.IsDefault != nil {
		in, out := &in.IsDefault, &out.IsDefault
		*out = new(bool)
		**out = **in
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(bool)
		**out = **in
	}
	if in.Editable != nil {
		in, out := &in.Editable, &out.Editable
		*out = new(bool)
		**out = **in
	}
	if in.JSONData != nil {
		in, out := &in.JSONData, &out.JSONData
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.SecureJSONData != nil {
		in, out := &in.SecureJSONData, &out.SecureJSONData
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDatasourceInternal.
func (in *GrafanaDatasourceInternal) DeepCopy() *GrafanaDatasourceInternal {
	if in == nil {
		return nil
	}
	out := new(GrafanaDatasourceInternal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDatasourceList) DeepCopyInto(out *GrafanaDatasourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GrafanaDatasource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDatasourceList.
func (in *GrafanaDatasourceList) DeepCopy() *GrafanaDatasourceList {
	if in == nil {
		return nil
	}
	out := new(GrafanaDatasourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDatasourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDatasourceSpec) DeepCopyInto(out *GrafanaDatasourceSpec) {
	*out = *in
	if in.Datasource != nil {
		in, out := &in.Datasource, &out.Datasource
		*out = new(GrafanaDatasourceInternal)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceSelector != nil {
		in, out := &in.InstanceSelector, &out.InstanceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make(PluginList, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDatasourceSpec.
func (in *GrafanaDatasourceSpec) DeepCopy() *GrafanaDatasourceSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaDatasourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDatasourceStatus) DeepCopyInto(out *GrafanaDatasourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDatasourceStatus.
func (in *GrafanaDatasourceStatus) DeepCopy() *GrafanaDatasourceStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaDatasourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaFolder) DeepCopyInto(out *GrafanaFolder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaFolder.
func (in *GrafanaFolder) DeepCopy() *GrafanaFolder {
	if in == nil {
		return nil
	}
	out := new(GrafanaFolder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaFolder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaFolderList) DeepCopyInto(out *GrafanaFolderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GrafanaFolder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaFolderList.
func (in *GrafanaFolderList) DeepCopy() *GrafanaFolderList {
	if in == nil {
		return nil
	}
	out := new(GrafanaFolderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaFolderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaFolderSpec) DeepCopyInto(out *GrafanaFolderSpec) {
	*out = *in
	if in.InstanceSelector != nil {
		in, out := &in.InstanceSelector, &out.InstanceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaFolderSpec.
func (in *GrafanaFolderSpec) DeepCopy() *GrafanaFolderSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaFolderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaFolderStatus) DeepCopyInto(out *GrafanaFolderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaFolderStatus.
func (in *GrafanaFolderStatus) DeepCopy() *GrafanaFolderStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaFolderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaList) DeepCopyInto(out *GrafanaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Grafana, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaList.
func (in *GrafanaList) DeepCopy() *GrafanaList {
	if in == nil {
		return nil
	}
	out := new(GrafanaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaPlugin) DeepCopyInto(out *GrafanaPlugin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaPlugin.
func (in *GrafanaPlugin) DeepCopy() *GrafanaPlugin {
	if in == nil {
		return nil
	}
	out := new(GrafanaPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaSpec) DeepCopyInto(out *GrafanaSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(IngressNetworkingV1)
		(*in).DeepCopyInto(*out)
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(RouteOpenshiftV1)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceV1)
		(*in).DeepCopyInto(*out)
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(DeploymentV1)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(PersistentVolumeClaimV1)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccountV1)
		(*in).DeepCopyInto(*out)
	}
	if in.Client != nil {
		in, out := &in.Client, &out.Client
		*out = new(GrafanaClient)
		(*in).DeepCopyInto(*out)
	}
	if in.Jsonnet != nil {
		in, out := &in.Jsonnet, &out.Jsonnet
		*out = new(JsonnetConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.GrafanaContainer != nil {
		in, out := &in.GrafanaContainer, &out.GrafanaContainer
		*out = new(GrafanaContainer)
		**out = **in
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(External)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaSpec.
func (in *GrafanaSpec) DeepCopy() *GrafanaSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaStatus) DeepCopyInto(out *GrafanaStatus) {
	*out = *in
	if in.Dashboards != nil {
		in, out := &in.Dashboards, &out.Dashboards
		*out = make(NamespacedResourceList, len(*in))
		copy(*out, *in)
	}
	if in.Datasources != nil {
		in, out := &in.Datasources, &out.Datasources
		*out = make(NamespacedResourceList, len(*in))
		copy(*out, *in)
	}
	if in.Folders != nil {
		in, out := &in.Folders, &out.Folders
		*out = make(NamespacedResourceList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaStatus.
func (in *GrafanaStatus) DeepCopy() *GrafanaStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNetworkingV1) DeepCopyInto(out *IngressNetworkingV1) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(networkingv1.IngressSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNetworkingV1.
func (in *IngressNetworkingV1) DeepCopy() *IngressNetworkingV1 {
	if in == nil {
		return nil
	}
	out := new(IngressNetworkingV1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonnetConfig) DeepCopyInto(out *JsonnetConfig) {
	*out = *in
	if in.LibraryLabelSelector != nil {
		in, out := &in.LibraryLabelSelector, &out.LibraryLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonnetConfig.
func (in *JsonnetConfig) DeepCopy() *JsonnetConfig {
	if in == nil {
		return nil
	}
	out := new(JsonnetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in NamespacedResourceList) DeepCopyInto(out *NamespacedResourceList) {
	{
		in := &in
		*out = make(NamespacedResourceList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedResourceList.
func (in NamespacedResourceList) DeepCopy() NamespacedResourceList {
	if in == nil {
		return nil
	}
	out := new(NamespacedResourceList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorReconcileVars) DeepCopyInto(out *OperatorReconcileVars) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorReconcileVars.
func (in *OperatorReconcileVars) DeepCopy() *OperatorReconcileVars {
	if in == nil {
		return nil
	}
	out := new(OperatorReconcileVars)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimV1) DeepCopyInto(out *PersistentVolumeClaimV1) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(PersistentVolumeClaimV1Spec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimV1.
func (in *PersistentVolumeClaimV1) DeepCopy() *PersistentVolumeClaimV1 {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimV1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimV1Spec) DeepCopyInto(out *PersistentVolumeClaimV1Spec) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.VolumeMode != nil {
		in, out := &in.VolumeMode, &out.VolumeMode
		*out = new(v1.PersistentVolumeMode)
		**out = **in
	}
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = new(v1.TypedLocalObjectReference)
		(*in).DeepCopyInto(*out)
	}
	if in.DataSourceRef != nil {
		in, out := &in.DataSourceRef, &out.DataSourceRef
		*out = new(v1.TypedLocalObjectReference)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimV1Spec.
func (in *PersistentVolumeClaimV1Spec) DeepCopy() *PersistentVolumeClaimV1Spec {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimV1Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PluginList) DeepCopyInto(out *PluginList) {
	{
		in := &in
		*out = make(PluginList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginList.
func (in PluginList) DeepCopy() PluginList {
	if in == nil {
		return nil
	}
	out := new(PluginList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PluginMap) DeepCopyInto(out *PluginMap) {
	{
		in := &in
		*out = make(PluginMap, len(*in))
		for key, val := range *in {
			var outVal []GrafanaPlugin
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PluginList, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginMap.
func (in PluginMap) DeepCopy() PluginMap {
	if in == nil {
		return nil
	}
	out := new(PluginMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteOpenShiftV1Spec) DeepCopyInto(out *RouteOpenShiftV1Spec) {
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(routev1.RouteTargetReference)
		(*in).DeepCopyInto(*out)
	}
	if in.AlternateBackends != nil {
		in, out := &in.AlternateBackends, &out.AlternateBackends
		*out = make([]routev1.RouteTargetReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(routev1.RoutePort)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(routev1.TLSConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteOpenShiftV1Spec.
func (in *RouteOpenShiftV1Spec) DeepCopy() *RouteOpenShiftV1Spec {
	if in == nil {
		return nil
	}
	out := new(RouteOpenShiftV1Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteOpenshiftV1) DeepCopyInto(out *RouteOpenshiftV1) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(RouteOpenShiftV1Spec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteOpenshiftV1.
func (in *RouteOpenshiftV1) DeepCopy() *RouteOpenshiftV1 {
	if in == nil {
		return nil
	}
	out := new(RouteOpenshiftV1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountV1) DeepCopyInto(out *ServiceAccountV1) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.AutomountServiceAccountToken != nil {
		in, out := &in.AutomountServiceAccountToken, &out.AutomountServiceAccountToken
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountV1.
func (in *ServiceAccountV1) DeepCopy() *ServiceAccountV1 {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountV1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceV1) DeepCopyInto(out *ServiceV1) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceV1.
func (in *ServiceV1) DeepCopy() *ServiceV1 {
	if in == nil {
		return nil
	}
	out := new(ServiceV1)
	in.DeepCopyInto(out)
	return out
}
