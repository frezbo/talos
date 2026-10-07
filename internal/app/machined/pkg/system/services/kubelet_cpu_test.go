// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package services_test

import (
	"context"
	"testing"

	"github.com/containerd/containerd/v2/core/containers"
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"

	"github.com/siderolabs/talos/internal/app/machined/pkg/system/services"
	"github.com/siderolabs/talos/pkg/machinery/resources/k8s"
)

func TestKubeletCPUArgumentsSameSnapshot(t *testing.T) {
	spec := k8s.NewKubeletSpec(k8s.NamespaceName, k8s.KubeletID)
	spec.TypedSpec().Args = []string{"--reserved-cpus=0-1"}
	spec.Metadata().Annotations().Set(k8s.KubeletCPUManagedAnnotation, "true")
	token := k8s.KubeletSpecToken(spec)
	args, option := services.KubeletCPUArguments(spec)
	spec.TypedSpec().Args[0] = "--reserved-cpus=2-3"
	spec.Metadata().Annotations().Delete(k8s.KubeletCPUManagedAnnotation)

	version, err := resource.ParseVersion("2")
	require.NoError(t, err)
	spec.Metadata().SetVersion(version)

	var ociSpec specs.Spec
	require.NoError(t, option(context.Background(), nil, &containers.Container{}, &ociSpec))
	require.Equal(t, []string{"/usr/local/bin/kubelet", "--reserved-cpus=0-1"}, args)
	require.Equal(t, token, ociSpec.Annotations[k8s.KubeletSpecTokenAnnotation])
	require.Equal(t, "true", ociSpec.Annotations[k8s.KubeletCPUManagedAnnotation])
	require.NotEqual(t, k8s.KubeletSpecToken(spec), token)
}
