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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta2 "k8s.io/api/apps/v1beta2"
	appsv1beta2 "k8s.io/client-go/applyconfigurations/apps/v1beta2"
	gentype "k8s.io/client-go/gentype"
	typedappsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
)

// fakeReplicaSets implements ReplicaSetInterface
type fakeReplicaSets struct {
	*gentype.FakeClientWithListAndApply[*v1beta2.ReplicaSet, *v1beta2.ReplicaSetList, *appsv1beta2.ReplicaSetApplyConfiguration]
	Fake *FakeAppsV1beta2
}

func newFakeReplicaSets(fake *FakeAppsV1beta2, namespace string) typedappsv1beta2.ReplicaSetInterface {
	return &fakeReplicaSets{
		gentype.NewFakeClientWithListAndApply[*v1beta2.ReplicaSet, *v1beta2.ReplicaSetList, *appsv1beta2.ReplicaSetApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta2.SchemeGroupVersion.WithResource("replicasets"),
			v1beta2.SchemeGroupVersion.WithKind("ReplicaSet"),
			func() *v1beta2.ReplicaSet { return &v1beta2.ReplicaSet{} },
			func() *v1beta2.ReplicaSetList { return &v1beta2.ReplicaSetList{} },
			func(dst, src *v1beta2.ReplicaSetList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta2.ReplicaSetList) []*v1beta2.ReplicaSet { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta2.ReplicaSetList, items []*v1beta2.ReplicaSet) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
