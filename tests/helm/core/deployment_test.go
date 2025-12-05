package core

import (
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
)

func TestDeployment(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name: "disableApprovedCheck adds command line flag",
			Values: map[string]interface{}{
				"disableApprovedCheck": true,
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				names := testutil.ResourceNames{Release: releaseName}
				deployment := h.GetDeployment(names.Deployment())
				testutil.ValidateDeploymentArgs(t, deployment, "-disable-approved-check")
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
