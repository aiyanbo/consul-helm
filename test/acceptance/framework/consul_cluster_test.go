package framework

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test that if TestConfig has values that need to provided
// to a helm install, it will respect the helmValues over
// the values from TestConfig.
func TestNewHelmCluster(t *testing.T) {
	helmValues := map[string]string{"global.image": "test-image", "feature.enabled": "true"}
	cluster := NewHelmCluster(t, helmValues, kubernetesContext{}, &TestConfig{ConsulImage: "test-config-image"}, "test")

	require.Equal(t, cluster.(*HelmCluster).helmOptions.SetValues, helmValues)
}