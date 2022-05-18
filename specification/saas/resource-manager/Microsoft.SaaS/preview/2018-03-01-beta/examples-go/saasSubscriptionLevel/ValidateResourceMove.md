Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsaas%2Farmsaas%2Fv0.5.0/sdk/resourcemanager/saas/armsaas/README.md) on how to add the SDK to your project and authenticate.

```go
package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasSubscriptionLevel/ValidateResourceMove.json
func ExampleSubscriptionLevelClient_ValidateMoveResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsaas.NewSubscriptionLevelClient("c825645b-e31b-9cf4-1cee-2aba9e58bc7c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.ValidateMoveResources(ctx,
		"my-saas-rg",
		armsaas.MoveResource{
			Resources: []*string{
				to.Ptr("/subscriptions/c825645b-e31b-9cf4-1cee-2aba9e58bc7c/resourceGroups/my-saas-rg/providers/Microsoft.SaaS/resources/saas1"),
				to.Ptr("/subscriptions/c825645b-e31b-9cf4-1cee-2aba9e58bc7c/resourceGroups/my-saas-rg/providers/Microsoft.SaaS/resources/saas2"),
				to.Ptr("/subscriptions/c825645b-e31b-9cf4-1cee-2aba9e58bc7c/resourceGroups/my-saas-rg/providers/Microsoft.SaaS/resources/saas3")},
			TargetResourceGroup: to.Ptr("/subscriptions/5122d0a3-1e10-4baf-bdc5-c2a452489525/resourceGroups/new-saas-rg"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```
