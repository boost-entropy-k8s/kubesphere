// +build !ignore_autogenerated

/*
Copyright 2020 The KubeSphere Authors.

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

package v1beta2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOverride) DeepCopyInto(out *ClusterOverride) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOverride.
func (in *ClusterOverride) DeepCopy() *ClusterOverride {
	if in == nil {
		return nil
	}
	out := new(ClusterOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationConfig) DeepCopyInto(out *FederatedNotificationConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(GenericFederatedStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationConfig.
func (in *FederatedNotificationConfig) DeepCopy() *FederatedNotificationConfig {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedNotificationConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationConfigList) DeepCopyInto(out *FederatedNotificationConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedNotificationConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationConfigList.
func (in *FederatedNotificationConfigList) DeepCopy() *FederatedNotificationConfigList {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedNotificationConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationConfigSpec) DeepCopyInto(out *FederatedNotificationConfigSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]GenericOverrideItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationConfigSpec.
func (in *FederatedNotificationConfigSpec) DeepCopy() *FederatedNotificationConfigSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationReceiver) DeepCopyInto(out *FederatedNotificationReceiver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(GenericFederatedStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationReceiver.
func (in *FederatedNotificationReceiver) DeepCopy() *FederatedNotificationReceiver {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedNotificationReceiver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationReceiverList) DeepCopyInto(out *FederatedNotificationReceiverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedNotificationReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationReceiverList.
func (in *FederatedNotificationReceiverList) DeepCopy() *FederatedNotificationReceiverList {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationReceiverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedNotificationReceiverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationReceiverSpec) DeepCopyInto(out *FederatedNotificationReceiverSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]GenericOverrideItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationReceiverSpec.
func (in *FederatedNotificationReceiverSpec) DeepCopy() *FederatedNotificationReceiverSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationReceiverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationRouter) DeepCopyInto(out *FederatedNotificationRouter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(GenericFederatedStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationRouter.
func (in *FederatedNotificationRouter) DeepCopy() *FederatedNotificationRouter {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedNotificationRouter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationRouterList) DeepCopyInto(out *FederatedNotificationRouterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedNotificationRouter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationRouterList.
func (in *FederatedNotificationRouterList) DeepCopy() *FederatedNotificationRouterList {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationRouterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedNotificationRouterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationRouterSpec) DeepCopyInto(out *FederatedNotificationRouterSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]GenericOverrideItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationRouterSpec.
func (in *FederatedNotificationRouterSpec) DeepCopy() *FederatedNotificationRouterSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationRouterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationSilence) DeepCopyInto(out *FederatedNotificationSilence) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(GenericFederatedStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationSilence.
func (in *FederatedNotificationSilence) DeepCopy() *FederatedNotificationSilence {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationSilence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedNotificationSilence) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationSilenceList) DeepCopyInto(out *FederatedNotificationSilenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedNotificationSilence, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationSilenceList.
func (in *FederatedNotificationSilenceList) DeepCopy() *FederatedNotificationSilenceList {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationSilenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedNotificationSilenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedNotificationSilenceSpec) DeepCopyInto(out *FederatedNotificationSilenceSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]GenericOverrideItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedNotificationSilenceSpec.
func (in *FederatedNotificationSilenceSpec) DeepCopy() *FederatedNotificationSilenceSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedNotificationSilenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericClusterReference) DeepCopyInto(out *GenericClusterReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericClusterReference.
func (in *GenericClusterReference) DeepCopy() *GenericClusterReference {
	if in == nil {
		return nil
	}
	out := new(GenericClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericClusterStatus) DeepCopyInto(out *GenericClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericClusterStatus.
func (in *GenericClusterStatus) DeepCopy() *GenericClusterStatus {
	if in == nil {
		return nil
	}
	out := new(GenericClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericCondition) DeepCopyInto(out *GenericCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericCondition.
func (in *GenericCondition) DeepCopy() *GenericCondition {
	if in == nil {
		return nil
	}
	out := new(GenericCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericFederatedResource) DeepCopyInto(out *GenericFederatedResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(GenericFederatedStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericFederatedResource.
func (in *GenericFederatedResource) DeepCopy() *GenericFederatedResource {
	if in == nil {
		return nil
	}
	out := new(GenericFederatedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericFederatedStatus) DeepCopyInto(out *GenericFederatedStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*GenericCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(GenericCondition)
				**out = **in
			}
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]GenericClusterStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericFederatedStatus.
func (in *GenericFederatedStatus) DeepCopy() *GenericFederatedStatus {
	if in == nil {
		return nil
	}
	out := new(GenericFederatedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericOverride) DeepCopyInto(out *GenericOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(GenericOverrideSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericOverride.
func (in *GenericOverride) DeepCopy() *GenericOverride {
	if in == nil {
		return nil
	}
	out := new(GenericOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericOverrideItem) DeepCopyInto(out *GenericOverrideItem) {
	*out = *in
	if in.ClusterOverrides != nil {
		in, out := &in.ClusterOverrides, &out.ClusterOverrides
		*out = make([]ClusterOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericOverrideItem.
func (in *GenericOverrideItem) DeepCopy() *GenericOverrideItem {
	if in == nil {
		return nil
	}
	out := new(GenericOverrideItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericOverrideSpec) DeepCopyInto(out *GenericOverrideSpec) {
	*out = *in
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]GenericOverrideItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericOverrideSpec.
func (in *GenericOverrideSpec) DeepCopy() *GenericOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(GenericOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericPlacement) DeepCopyInto(out *GenericPlacement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericPlacement.
func (in *GenericPlacement) DeepCopy() *GenericPlacement {
	if in == nil {
		return nil
	}
	out := new(GenericPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericPlacementFields) DeepCopyInto(out *GenericPlacementFields) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]GenericClusterReference, len(*in))
		copy(*out, *in)
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericPlacementFields.
func (in *GenericPlacementFields) DeepCopy() *GenericPlacementFields {
	if in == nil {
		return nil
	}
	out := new(GenericPlacementFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericPlacementSpec) DeepCopyInto(out *GenericPlacementSpec) {
	*out = *in
	in.Placement.DeepCopyInto(&out.Placement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericPlacementSpec.
func (in *GenericPlacementSpec) DeepCopy() *GenericPlacementSpec {
	if in == nil {
		return nil
	}
	out := new(GenericPlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationConfigTemplate) DeepCopyInto(out *NotificationConfigTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationConfigTemplate.
func (in *NotificationConfigTemplate) DeepCopy() *NotificationConfigTemplate {
	if in == nil {
		return nil
	}
	out := new(NotificationConfigTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationReceiverTemplate) DeepCopyInto(out *NotificationReceiverTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationReceiverTemplate.
func (in *NotificationReceiverTemplate) DeepCopy() *NotificationReceiverTemplate {
	if in == nil {
		return nil
	}
	out := new(NotificationReceiverTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationRouterTemplate) DeepCopyInto(out *NotificationRouterTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationRouterTemplate.
func (in *NotificationRouterTemplate) DeepCopy() *NotificationRouterTemplate {
	if in == nil {
		return nil
	}
	out := new(NotificationRouterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationSilenceTemplate) DeepCopyInto(out *NotificationSilenceTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationSilenceTemplate.
func (in *NotificationSilenceTemplate) DeepCopy() *NotificationSilenceTemplate {
	if in == nil {
		return nil
	}
	out := new(NotificationSilenceTemplate)
	in.DeepCopyInto(out)
	return out
}
