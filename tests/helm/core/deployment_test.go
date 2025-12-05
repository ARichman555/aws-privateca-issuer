package core

import (
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
)

func TestDeployment(t *testing.T) {
	testCases := []testutil.PrivateCaHelmTestCase{
		{
			Name: "disableApprovedCheck adds command line flag",
			Values: map[string]interface{}{
				"disableApprovedCheck": true,
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				resources := testutil.PrivateCaIssuerResources{Release: releaseName}
				deployment := h.GetPrivateCaDeployment(resources.Deployment())
				testutil.ValidatePrivateCaDeploymentArgs(t, deployment, "-disable-approved-check")
			},
		},
	}

	testutil.RunPrivateCaHelmTests(t, testCases)
}
