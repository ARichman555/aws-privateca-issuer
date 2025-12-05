package features

import (
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
)

func TestOptionalFields(t *testing.T) {
	testCases := []testutil.PrivateCaHelmTestCase{
		{
			Name: "optional fields configuration",
			Values: map[string]interface{}{
				"env": map[string]interface{}{
					"LOG_LEVEL": "debug",
				},
				"extraContainers": []interface{}{
					map[string]interface{}{
						"name":    "sidecar",
						"image":   "busybox:latest",
						"command": []interface{}{"sleep", "3600"},
					},
				},
				"imagePullSecrets": []interface{}{
					map[string]interface{}{
						"name": "my-registry-secret",
					},
				},
				"nodeSelector": map[string]interface{}{
					"kubernetes.io/os": "linux",
				},
				"tolerations": []interface{}{
					map[string]interface{}{
						"key":      "node-role.kubernetes.io/master",
						"operator": "Exists",
						"effect":   "NoSchedule",
					},
				},
				"priorityClassName": "high-priority",
				"volumes": []interface{}{
					map[string]interface{}{
						"name": "config-volume",
						"configMap": map[string]interface{}{
							"name": "my-config",
						},
					},
				},
				"volumeMounts": []interface{}{
					map[string]interface{}{
						"name":      "config-volume",
						"mountPath": "/etc/config",
					},
				},
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				// Validate environment variables
				container := h.GetPrivateCaDeployment(testutil.PrivateCaIssuerResources{Release: releaseName}.Deployment()).Spec.Template.Spec.Containers[0]
				found := false
				for _, env := range container.Env {
					if env.Name == "LOG_LEVEL" && env.Value == "debug" {
						found = true
						break
					}
				}
				assert.True(t, found, "LOG_LEVEL environment variable should be set to debug")

				// Validate extra containers
				assert.Len(t, h.GetPrivateCaDeployment(testutil.PrivateCaIssuerResources{Release: releaseName}.Deployment()).Spec.Template.Spec.Containers, 2, "Should have main container plus sidecar")
				sidecar := h.GetPrivateCaDeployment(testutil.PrivateCaIssuerResources{Release: releaseName}.Deployment()).Spec.Template.Spec.Containers[1]
				assert.Equal(t, "sidecar", sidecar.Name)
				assert.Equal(t, "busybox:latest", sidecar.Image)

				// Validate node selector
				assert.Equal(t, "linux", h.GetPrivateCaDeployment(testutil.PrivateCaIssuerResources{Release: releaseName}.Deployment()).Spec.Template.Spec.NodeSelector["kubernetes.io/os"])

				// Validate priority class
				assert.Equal(t, "high-priority", h.GetPrivateCaDeployment(testutil.PrivateCaIssuerResources{Release: releaseName}.Deployment()).Spec.Template.Spec.PriorityClassName)
			},
		},
	}

	testutil.RunPrivateCaHelmTests(t, testCases)
}
