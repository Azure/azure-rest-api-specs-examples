Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsubscription%2Farmsubscription%2Fv1.0.0/sdk/resourcemanager/subscription/armsubscription/README.md) on how to add the SDK to your project and authenticate.

```go
package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/createAlias.json
func ExampleAliasClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsubscription.NewAliasClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"aliasForNewSub",
		armsubscription.PutAliasRequest{
			Properties: &armsubscription.PutAliasRequestProperties{
				AdditionalProperties: &armsubscription.PutAliasRequestAdditionalProperties{
					SubscriptionOwnerID:  to.Ptr("f09b39eb-c496-482c-9ab9-afd799572f4c"),
					SubscriptionTenantID: to.Ptr("66f6e4d6-07dc-4aea-94ea-e12d3026a3c8"),
					Tags: map[string]*string{
						"tag1": to.Ptr("Messi"),
						"tag2": to.Ptr("Ronaldo"),
						"tag3": to.Ptr("Lebron"),
					},
				},
				BillingScope: to.Ptr("/billingAccounts/af6231a7-7f8d-4fcc-a993-dd8466108d07:c663dac6-a9a5-405a-8938-cd903e12ab5b_2019_05_31/billingProfiles/QWDQ-QWHI-AUW-SJDO-DJH/invoiceSections/FEUF-EUHE-ISJ-SKDW-DJH"),
				DisplayName:  to.Ptr("Test Subscription"),
				Workload:     to.Ptr(armsubscription.WorkloadProduction),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
