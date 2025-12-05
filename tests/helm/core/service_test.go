package core

import (
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	corev1 "k8s.io/api/core/v1"
)

func TestService(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name:   "custom service configuration",
			Values: testutil.ServiceValues("NodePort", 9090),
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				names := testutil.ResourceNames{Release: releaseName}
				service := h.GetService(names.Service())
				
				testutil.ValidateServiceType(t, service, corev1.ServiceTypeNodePort)
				testutil.ValidateServicePort(t, service, 9090)
			},
		},
		{
			Name:   "nameOverride affects resource names",
			Values: testutil.NameOverrideValues("custom-issuer", ""),
			DeploymentName: func(releaseName string) string {
				return releaseName + "-custom-issuer"
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				deployment := h.GetDeployment(releaseName + "-custom-issuer")
				testutil.ValidateDeploymentReplicas(t, deployment, 2)
			},
		},
		{
			Name:   "fullnameOverride completely overrides resource names",
			Values: testutil.NameOverrideValues("", "completely-custom-name"),
			DeploymentName: func(releaseName string) string {
				return "completely-custom-name"
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				deployment := h.GetDeployment("completely-custom-name")
				testutil.ValidateDeploymentReplicas(t, deployment, 2)
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
