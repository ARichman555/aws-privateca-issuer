package core

import (
	"context"
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestService(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name: "custom service configuration",
			Values: map[string]interface{}{
				"service": map[string]interface{}{
					"type": "NodePort",
					"port": 9090,
				},
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				serviceName := releaseName + "-aws-privateca-issuer"
				service, err := h.Clientset.CoreV1().Services(h.Namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
				require.NoError(t, err)

				assert.Equal(t, "NodePort", string(service.Spec.Type))
				assert.Equal(t, int32(9090), service.Spec.Ports[0].Port)
			},
		},
		{
			Name: "nameOverride affects resource names",
			Values: map[string]interface{}{
				"nameOverride": "custom-issuer",
			},
			DeploymentName: func(releaseName string) string {
				return releaseName + "-custom-issuer"
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				deploymentName := releaseName + "-custom-issuer"
				deployment, err := h.Clientset.AppsV1().Deployments(h.Namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
				require.NoError(t, err)
				assert.NotNil(t, deployment)
			},
		},
		{
			Name: "fullnameOverride completely overrides resource names",
			Values: map[string]interface{}{
				"fullnameOverride": "completely-custom-name",
			},
			DeploymentName: func(releaseName string) string {
				return "completely-custom-name"
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				deploymentName := "completely-custom-name"
				deployment, err := h.Clientset.AppsV1().Deployments(h.Namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
				require.NoError(t, err)
				assert.NotNil(t, deployment)
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
