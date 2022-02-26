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

// x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_CreateSubscriptionScope.json
func ExampleAttestationsClient_BeginCreateOrUpdateAtSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicyinsights.NewAttestationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdateAtSubscription(ctx,
		"<attestation-name>",
		armpolicyinsights.Attestation{
			Properties: &armpolicyinsights.AttestationProperties{
				ComplianceState:    armpolicyinsights.ComplianceState("Compliant").ToPtr(),
				PolicyAssignmentID: to.StringPtr("<policy-assignment-id>"),
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
	log.Printf("Response result: %#v\n", res.AttestationsClientCreateOrUpdateAtSubscriptionResult)
}
```
