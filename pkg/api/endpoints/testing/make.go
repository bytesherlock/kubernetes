/*
Copyright 2021 The Kubernetes Authors.

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

package testing

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	api "k8s.io/kubernetes/pkg/apis/core"
)

// Tweak is a function that modifies a Endpoints.
type Tweak func(*api.Endpoints)

// MakeEndpointAddress helps construct EndpointAddress objects which pass API
// validation.
func MakeEndpointAddress(ip string, pod string) api.EndpointAddress {
	return api.EndpointAddress{
		IP: ip,
		TargetRef: &api.ObjectReference{
			Name:      pod,
			Namespace: metav1.NamespaceDefault,
		},
	}
}

// MakeEndpointPort helps construct EndpointPort objects which pass API
// validation.
func MakeEndpointPort(name string, port int) api.EndpointPort {
	return api.EndpointPort{
		Name: name,
		Port: int32(port),
	}
}
