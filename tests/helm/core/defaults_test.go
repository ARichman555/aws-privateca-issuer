package core

import (
	"strings"
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestDefaults(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name:   "default values validation",
			Values: map[string]interface{}{},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				names := testutil.ResourceNames{Release: releaseName}
				
				// Validate deployment defaults
				deployment := h.GetDeployment(names.Deployment())
				testutil.ValidateDeploymentReplicas(t, deployment, 2)
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
				service := h.GetService(names.Service())
				testutil.ValidateServiceType(t, service, corev1.ServiceTypeClusterIP)
				testutil.ValidateServicePort(t, service, 8080)
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
