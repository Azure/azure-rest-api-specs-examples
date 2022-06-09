```go
package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/PutCommitmentPlan.json
func ExampleCommitmentPlansClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"accountName",
		"commitmentPlanName",
		armcognitiveservices.CommitmentPlan{
			Properties: &armcognitiveservices.CommitmentPlanProperties{
				AutoRenew: to.Ptr(true),
				Current: &armcognitiveservices.CommitmentPeriod{
					Tier: to.Ptr("T1"),
				},
				HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
				PlanType:     to.Ptr("Speech2Text"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcognitiveservices%2Farmcognitiveservices%2Fv1.0.0/sdk/resourcemanager/cognitiveservices/armcognitiveservices/README.md) on how to add the SDK to your project and authenticate.
