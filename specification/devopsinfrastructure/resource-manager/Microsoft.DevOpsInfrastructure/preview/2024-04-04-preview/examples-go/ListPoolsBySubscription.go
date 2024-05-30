package armdevopsinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/ListPoolsBySubscription.json
func ExamplePoolsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevopsinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPoolsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PoolListResult = armdevopsinfrastructure.PoolListResult{
		// 	Value: []*armdevopsinfrastructure.Pool{
		// 		{
		// 			ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/resourceGroups/aoiresourceGroupName/providers/Microsoft.DevOpsInfrastructure/Pools/pool"),
		// 			Location: to.Ptr("eastus"),
		// 	}},
		// }
	}
}
