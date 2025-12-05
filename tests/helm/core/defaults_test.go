package core

import (
	"strings"
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestDefaults(t *testing.T) {
	testCases := []testutil.PrivateCaHelmTestCase{
		{
			Name:   "default values validation",
			Values: map[string]interface{}{},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				resources := testutil.PrivateCaIssuerResources{Release: releaseName}

				// Validate deployment defaults
				deployment := h.GetPrivateCaDeployment(resources.Deployment())
				testutil.ValidatePrivateCaDeploymentReplicas(t, deployment, 2)
				assert.Equal(t, int32(10), *deployment.Spec.RevisionHistoryLimit)

				// Validate image based on test mode
				container := deployment.Spec.Template.Spec.Containers[0]
				mode := testutil.GetTestMode()
				if mode == testutil.BetaMode {
					registry, repoName := testutil.GetBetaDefaults()
					expectedRepo := registry + "/" + strings.ToLower(strings.ReplaceAll(repoName, "/", "-")) + "-test"
					assert.Contains(t, container.Image, expectedRepo)
					assert.Equal(t, "Always", string(container.ImagePullPolicy))
				} else if mode == testutil.ProdMode {
					registry, repoName := testutil.GetProdDefaults()
					expectedRepo := registry + "/" + strings.ToLower(strings.ReplaceAll(repoName, "/", "-"))
					assert.Contains(t, container.Image, expectedRepo)
					assert.Equal(t, "IfNotPresent", string(container.ImagePullPolicy))
				} else {
					assert.Contains(t, container.Image, "localhost:5000/aws-privateca-issuer")
					assert.Equal(t, "IfNotPresent", string(container.ImagePullPolicy))
				}

				// Validate service defaults
				service := h.GetPrivateCaService(resources.Service())
				testutil.ValidatePrivateCaServiceType(t, service, corev1.ServiceTypeClusterIP)
				testutil.ValidatePrivateCaServicePort(t, service, 8080)
			},
		},
	}

	testutil.RunPrivateCaHelmTests(t, testCases)
}
