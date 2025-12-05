package features

import (
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestDeploymentConfiguration(t *testing.T) {
	testCases := []testutil.PrivateCaHelmTestCase{
		{
			Name: "custom resources and security context",
			Values: map[string]interface{}{
				"replicaCount": 3,
				"resources": map[string]interface{}{
					"limits": map[string]interface{}{
						"cpu":    "100m",
						"memory": "128Mi",
					},
					"requests": map[string]interface{}{
						"cpu":    "25m",
						"memory": "32Mi",
					},
				},
				"securityContext": map[string]interface{}{
					"allowPrivilegeEscalation": false,
					"runAsNonRoot":             true,
				},
				"podSecurityContext": map[string]interface{}{
					"runAsUser": 1000,
				},
				"revisionHistoryLimit": 5,
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				resources := testutil.PrivateCaIssuerResources{Release: releaseName}
				deployment := h.GetPrivateCaDeployment(resources.Deployment())

				testutil.ValidatePrivateCaDeploymentReplicas(t, deployment, 3)
				assert.Equal(t, int32(5), *deployment.Spec.RevisionHistoryLimit)

				// Validate security context
				container := deployment.Spec.Template.Spec.Containers[0]
				assert.Equal(t, false, *container.SecurityContext.AllowPrivilegeEscalation)
				assert.Equal(t, true, *container.SecurityContext.RunAsNonRoot)

				// Validate pod security context
				assert.Equal(t, int64(1000), *deployment.Spec.Template.Spec.SecurityContext.RunAsUser)
			},
		},
		{
			Name: "disableClientSideRateLimiting adds command line flag",
			Values: map[string]interface{}{
				"disableClientSideRateLimiting": true,
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				resources := testutil.PrivateCaIssuerResources{Release: releaseName}
				deployment := h.GetPrivateCaDeployment(resources.Deployment())
				testutil.ValidatePrivateCaDeploymentArgs(t, deployment, "-disable-client-side-rate-limiting")
			},
		},
		{
			Name:   "custom image configuration",
			Values: testutil.PrivateCaImageValues("custom.registry.com/aws-privateca-issuer", "v1.0.0", "Always"),
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				resources := testutil.PrivateCaIssuerResources{Release: releaseName}
				deployment := h.GetPrivateCaDeployment(resources.Deployment())

				container := deployment.Spec.Template.Spec.Containers[0]
				assert.Equal(t, "custom.registry.com/aws-privateca-issuer:v1.0.0", container.Image)
				assert.Equal(t, corev1.PullAlways, container.ImagePullPolicy)
			},
		},
		{
			Name: "pod annotations and labels",
			Values: map[string]interface{}{
				"podAnnotations": map[string]interface{}{
					"prometheus.io/scrape": "true",
					"prometheus.io/port":   "8080",
				},
				"podLabels": map[string]interface{}{
					"environment": "test",
					"team":        "platform",
				},
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				resources := testutil.PrivateCaIssuerResources{Release: releaseName}
				deployment := h.GetPrivateCaDeployment(resources.Deployment())

				// Validate annotations
				annotations := deployment.Spec.Template.Annotations
				assert.Equal(t, "true", annotations["prometheus.io/scrape"])
				assert.Equal(t, "8080", annotations["prometheus.io/port"])

				// Validate labels
				labels := deployment.Spec.Template.Labels
				assert.Equal(t, "test", labels["environment"])
				assert.Equal(t, "platform", labels["team"])
			},
		},
	}

	testutil.RunPrivateCaHelmTests(t, testCases)
}
