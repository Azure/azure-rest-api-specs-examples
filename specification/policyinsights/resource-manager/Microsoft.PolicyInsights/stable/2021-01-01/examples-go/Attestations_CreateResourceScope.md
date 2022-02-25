Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpolicyinsights%2Farmpolicyinsights%2Fv0.2.0/sdk/resourcemanager/policyinsights/armpolicyinsights/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_CreateResourceScope.json
func ExampleAttestationsClient_BeginCreateOrUpdateAtResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicyinsights.NewAttestationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdateAtResource(ctx,
		"<resource-id>",
		"<attestation-name>",
		armpolicyinsights.Attestation{
			Properties: &armpolicyinsights.AttestationProperties{
				Comments:        to.StringPtr("<comments>"),
				ComplianceState: armpolicyinsights.ComplianceState("Compliant").ToPtr(),
				Evidence: []*armpolicyinsights.AttestationEvidence{
					{
						Description: to.StringPtr("<description>"),
						SourceURI:   to.StringPtr("<source-uri>"),
					}},
				ExpiresOn:                   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T00:00:00Z"); return t }()),
				Owner:                       to.StringPtr("<owner>"),
				PolicyAssignmentID:          to.StringPtr("<policy-assignment-id>"),
				PolicyDefinitionReferenceID: to.StringPtr("<policy-definition-reference-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AttestationsClientCreateOrUpdateAtResourceResult)
}
```
