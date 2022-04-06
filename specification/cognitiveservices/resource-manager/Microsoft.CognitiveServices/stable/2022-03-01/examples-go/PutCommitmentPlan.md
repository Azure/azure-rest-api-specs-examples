Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcognitiveservices%2Farmcognitiveservices%2Fv0.4.0/sdk/resourcemanager/cognitiveservices/armcognitiveservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/PutCommitmentPlan.json
func ExampleCommitmentPlansClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewCommitmentPlansClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<commitment-plan-name>",
		armcognitiveservices.CommitmentPlan{
			Properties: &armcognitiveservices.CommitmentPlanProperties{
				AutoRenew: to.BoolPtr(true),
				Current: &armcognitiveservices.CommitmentPeriod{
					Tier: to.StringPtr("<tier>"),
				},
				HostingModel: armcognitiveservices.HostingModel("Web").ToPtr(),
				PlanType:     to.StringPtr("<plan-type>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CommitmentPlansClientCreateOrUpdateResult)
}
```
