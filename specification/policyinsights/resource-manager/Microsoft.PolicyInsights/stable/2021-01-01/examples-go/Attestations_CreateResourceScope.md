Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpolicyinsights%2Farmpolicyinsights%2Fv0.4.0/sdk/resourcemanager/policyinsights/armpolicyinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_CreateResourceScope.json
func ExampleAttestationsClient_BeginCreateOrUpdateAtResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewAttestationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdateAtResource(ctx,
		"<resource-id>",
		"<attestation-name>",
		armpolicyinsights.Attestation{
			Properties: &armpolicyinsights.AttestationProperties{
				Comments:        to.Ptr("<comments>"),
				ComplianceState: to.Ptr(armpolicyinsights.ComplianceStateCompliant),
				Evidence: []*armpolicyinsights.AttestationEvidence{
					{
						Description: to.Ptr("<description>"),
						SourceURI:   to.Ptr("<source-uri>"),
					}},
				ExpiresOn:                   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T00:00:00Z"); return t }()),
				Owner:                       to.Ptr("<owner>"),
				PolicyAssignmentID:          to.Ptr("<policy-assignment-id>"),
				PolicyDefinitionReferenceID: to.Ptr("<policy-definition-reference-id>"),
			},
		},
		&armpolicyinsights.AttestationsClientBeginCreateOrUpdateAtResourceOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
