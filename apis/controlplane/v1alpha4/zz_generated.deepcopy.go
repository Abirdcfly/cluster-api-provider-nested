// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1alpha4

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedAPIServer) DeepCopyInto(out *NestedAPIServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedAPIServer.
func (in *NestedAPIServer) DeepCopy() *NestedAPIServer {
	if in == nil {
		return nil
	}
	out := new(NestedAPIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NestedAPIServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedAPIServerList) DeepCopyInto(out *NestedAPIServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NestedAPIServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedAPIServerList.
func (in *NestedAPIServerList) DeepCopy() *NestedAPIServerList {
	if in == nil {
		return nil
	}
	out := new(NestedAPIServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NestedAPIServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedAPIServerSpec) DeepCopyInto(out *NestedAPIServerSpec) {
	*out = *in
	in.NestedComponentSpec.DeepCopyInto(&out.NestedComponentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedAPIServerSpec.
func (in *NestedAPIServerSpec) DeepCopy() *NestedAPIServerSpec {
	if in == nil {
		return nil
	}
	out := new(NestedAPIServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedAPIServerStatus) DeepCopyInto(out *NestedAPIServerStatus) {
	*out = *in
	if in.APIServerService != nil {
		in, out := &in.APIServerService, &out.APIServerService
		*out = new(v1.ObjectReference)
		**out = **in
	}
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedAPIServerStatus.
func (in *NestedAPIServerStatus) DeepCopy() *NestedAPIServerStatus {
	if in == nil {
		return nil
	}
	out := new(NestedAPIServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedComponentSpec) DeepCopyInto(out *NestedComponentSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	in.PatchSpec.DeepCopyInto(&out.PatchSpec)
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedComponentSpec.
func (in *NestedComponentSpec) DeepCopy() *NestedComponentSpec {
	if in == nil {
		return nil
	}
	out := new(NestedComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControlPlane) DeepCopyInto(out *NestedControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControlPlane.
func (in *NestedControlPlane) DeepCopy() *NestedControlPlane {
	if in == nil {
		return nil
	}
	out := new(NestedControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NestedControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControlPlaneList) DeepCopyInto(out *NestedControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NestedControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControlPlaneList.
func (in *NestedControlPlaneList) DeepCopy() *NestedControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(NestedControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NestedControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControlPlaneSpec) DeepCopyInto(out *NestedControlPlaneSpec) {
	*out = *in
	if in.EtcdRef != nil {
		in, out := &in.EtcdRef, &out.EtcdRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.APIServerRef != nil {
		in, out := &in.APIServerRef, &out.APIServerRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ControllerManagerRef != nil {
		in, out := &in.ControllerManagerRef, &out.ControllerManagerRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControlPlaneSpec.
func (in *NestedControlPlaneSpec) DeepCopy() *NestedControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(NestedControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControlPlaneStatus) DeepCopyInto(out *NestedControlPlaneStatus) {
	*out = *in
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = new(NestedControlPlaneStatusEtcd)
		(*in).DeepCopyInto(*out)
	}
	if in.APIServer != nil {
		in, out := &in.APIServer, &out.APIServer
		*out = new(NestedControlPlaneStatusAPIServer)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControlPlaneStatus.
func (in *NestedControlPlaneStatus) DeepCopy() *NestedControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(NestedControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControlPlaneStatusAPIServer) DeepCopyInto(out *NestedControlPlaneStatusAPIServer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControlPlaneStatusAPIServer.
func (in *NestedControlPlaneStatusAPIServer) DeepCopy() *NestedControlPlaneStatusAPIServer {
	if in == nil {
		return nil
	}
	out := new(NestedControlPlaneStatusAPIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControlPlaneStatusEtcd) DeepCopyInto(out *NestedControlPlaneStatusEtcd) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]NestedEtcdAddress, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControlPlaneStatusEtcd.
func (in *NestedControlPlaneStatusEtcd) DeepCopy() *NestedControlPlaneStatusEtcd {
	if in == nil {
		return nil
	}
	out := new(NestedControlPlaneStatusEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControllerManager) DeepCopyInto(out *NestedControllerManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControllerManager.
func (in *NestedControllerManager) DeepCopy() *NestedControllerManager {
	if in == nil {
		return nil
	}
	out := new(NestedControllerManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NestedControllerManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControllerManagerList) DeepCopyInto(out *NestedControllerManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NestedControllerManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControllerManagerList.
func (in *NestedControllerManagerList) DeepCopy() *NestedControllerManagerList {
	if in == nil {
		return nil
	}
	out := new(NestedControllerManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NestedControllerManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControllerManagerSpec) DeepCopyInto(out *NestedControllerManagerSpec) {
	*out = *in
	in.NestedComponentSpec.DeepCopyInto(&out.NestedComponentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControllerManagerSpec.
func (in *NestedControllerManagerSpec) DeepCopy() *NestedControllerManagerSpec {
	if in == nil {
		return nil
	}
	out := new(NestedControllerManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedControllerManagerStatus) DeepCopyInto(out *NestedControllerManagerStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedControllerManagerStatus.
func (in *NestedControllerManagerStatus) DeepCopy() *NestedControllerManagerStatus {
	if in == nil {
		return nil
	}
	out := new(NestedControllerManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedEtcd) DeepCopyInto(out *NestedEtcd) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedEtcd.
func (in *NestedEtcd) DeepCopy() *NestedEtcd {
	if in == nil {
		return nil
	}
	out := new(NestedEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NestedEtcd) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedEtcdAddress) DeepCopyInto(out *NestedEtcdAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedEtcdAddress.
func (in *NestedEtcdAddress) DeepCopy() *NestedEtcdAddress {
	if in == nil {
		return nil
	}
	out := new(NestedEtcdAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedEtcdList) DeepCopyInto(out *NestedEtcdList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NestedEtcd, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedEtcdList.
func (in *NestedEtcdList) DeepCopy() *NestedEtcdList {
	if in == nil {
		return nil
	}
	out := new(NestedEtcdList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NestedEtcdList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedEtcdSpec) DeepCopyInto(out *NestedEtcdSpec) {
	*out = *in
	in.NestedComponentSpec.DeepCopyInto(&out.NestedComponentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedEtcdSpec.
func (in *NestedEtcdSpec) DeepCopy() *NestedEtcdSpec {
	if in == nil {
		return nil
	}
	out := new(NestedEtcdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedEtcdStatus) DeepCopyInto(out *NestedEtcdStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]NestedEtcdAddress, len(*in))
		copy(*out, *in)
	}
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedEtcdStatus.
func (in *NestedEtcdStatus) DeepCopy() *NestedEtcdStatus {
	if in == nil {
		return nil
	}
	out := new(NestedEtcdStatus)
	in.DeepCopyInto(out)
	return out
}
