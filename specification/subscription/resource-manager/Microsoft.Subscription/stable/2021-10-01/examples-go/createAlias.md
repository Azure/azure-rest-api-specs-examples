Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsubscription%2Farmsubscription%2Fv0.4.0/sdk/resourcemanager/subscription/armsubscription/README.md) on how to add the SDK to your project and authenticate.

```go
package armsubscription_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/createAlias.json
func ExampleAliasClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsubscription.NewAliasClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<alias-name>",
		armsubscription.PutAliasRequest{
			Properties: &armsubscription.PutAliasRequestProperties{
				AdditionalProperties: &armsubscription.PutAliasRequestAdditionalProperties{
					SubscriptionOwnerID:  to.Ptr("<subscription-owner-id>"),
					SubscriptionTenantID: to.Ptr("<subscription-tenant-id>"),
					Tags: map[string]*string{
						"tag1": to.Ptr("Messi"),
						"tag2": to.Ptr("Ronaldo"),
						"tag3": to.Ptr("Lebron"),
					},
				},
				BillingScope: to.Ptr("<billing-scope>"),
				DisplayName:  to.Ptr("<display-name>"),
				Workload:     to.Ptr(armsubscription.WorkloadProduction),
			},
		},
		&armsubscription.AliasClientBeginCreateOptions{ResumeToken: ""})
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
