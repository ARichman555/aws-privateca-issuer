package core

import (
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	corev1 "k8s.io/api/core/v1"
)

func TestService(t *testing.T) {
	testCases := []testutil.PrivateCaHelmTestCase{
		{
			Name:   "custom service configuration",
			Values: testutil.PrivateCaServiceValues("NodePort", 9090),
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				resources := testutil.PrivateCaIssuerResources{Release: releaseName}
				service := h.GetPrivateCaService(resources.Service())

				testutil.ValidatePrivateCaServiceType(t, service, corev1.ServiceTypeNodePort)
				testutil.ValidatePrivateCaServicePort(t, service, 9090)
			},
		},
		{
			Name:   "nameOverride affects resource names",
			Values: testutil.PrivateCaNameOverrideValues("custom-issuer", ""),
			DeploymentName: func(releaseName string) string {
				return releaseName + "-custom-issuer"
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				deployment := h.GetPrivateCaDeployment(releaseName + "-custom-issuer")
				testutil.ValidatePrivateCaDeploymentReplicas(t, deployment, 2)
			},
		},
		{
			Name:   "fullnameOverride completely overrides resource names",
			Values: testutil.PrivateCaNameOverrideValues("", "completely-custom-name"),
			DeploymentName: func(releaseName string) string {
				return "completely-custom-name"
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				deployment := h.GetPrivateCaDeployment("completely-custom-name")
				testutil.ValidatePrivateCaDeploymentReplicas(t, deployment, 2)
			},
		},
	}

	testutil.RunPrivateCaHelmTests(t, testCases)
}
