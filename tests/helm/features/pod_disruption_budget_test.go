package features

import (
	"context"
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func TestPodDisruptionBudget(t *testing.T) {
	testCases := []testutil.PrivateCaHelmTestCase{
		{
			Name: "podDisruptionBudget with maxUnavailable",
			Values: map[string]interface{}{
				"podDisruptionBudget": map[string]interface{}{
					"maxUnavailable": 1,
				},
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				pdb, err := h.Clientset.PolicyV1().PodDisruptionBudgets(h.Namespace).Get(context.TODO(), testutil.PrivateCaIssuerResources{Release: releaseName}.Deployment(), metav1.GetOptions{})
				require.NoError(t, err)

				expectedMaxUnavailable := intstr.FromInt(1)
				assert.Equal(t, &expectedMaxUnavailable, pdb.Spec.MaxUnavailable)
			},
		},
	}

	testutil.RunPrivateCaHelmTests(t, testCases)
}
