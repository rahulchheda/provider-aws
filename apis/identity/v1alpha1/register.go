/*
Copyright 2019 The Crossplane Authors.

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
// +kubebuilder:object:generate=true
// +groupName=identity.aws.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "identity.aws.crossplane.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// IAMUser type metadata.
var (
	IAMUserKind             = reflect.TypeOf(IAMUser{}).Name()
	IAMUserGroupKind        = schema.GroupKind{Group: Group, Kind: IAMUserKind}.String()
	IAMUserKindAPIVersion   = IAMUserKind + "." + SchemeGroupVersion.String()
	IAMUserGroupVersionKind = SchemeGroupVersion.WithKind(IAMUserKind)
)

// IAMUserPolicyAttachment type metadata.
var (
	IAMUserPolicyAttachmentKind             = reflect.TypeOf(IAMUserPolicyAttachment{}).Name()
	IAMUserPolicyAttachmentGroupKind        = schema.GroupKind{Group: Group, Kind: IAMUserPolicyAttachmentKind}.String()
	IAMUserPolicyAttachmentKindAPIVersion   = IAMUserPolicyAttachmentKind + "." + SchemeGroupVersion.String()
	IAMUserPolicyAttachmentGroupVersionKind = SchemeGroupVersion.WithKind(IAMUserPolicyAttachmentKind)
)

// IAMPolicy type metadata.
var (
	IAMPolicyKind             = reflect.TypeOf(IAMPolicy{}).Name()
	IAMPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: IAMPolicyKind}.String()
	IAMPolicyKindAPIVersion   = IAMPolicyKind + "." + SchemeGroupVersion.String()
	IAMPolicyGroupVersionKind = SchemeGroupVersion.WithKind(IAMPolicyKind)
)

func init() {
	SchemeBuilder.Register(&IAMUser{}, &IAMUserList{})
	SchemeBuilder.Register(&IAMPolicy{}, &IAMPolicyList{})
	SchemeBuilder.Register(&IAMUserPolicyAttachment{}, &IAMUserPolicyAttachmentList{})
}
