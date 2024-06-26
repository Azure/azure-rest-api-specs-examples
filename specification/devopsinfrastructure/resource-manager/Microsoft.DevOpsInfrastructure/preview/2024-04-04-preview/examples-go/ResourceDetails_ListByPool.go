package armdevopsinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/ResourceDetails_ListByPool.json
func ExampleResourceDetailsClient_NewListByPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevopsinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceDetailsClient().NewListByPoolPager("my-resource-group", "my-dev-ops-pool", nil)
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
		// page.ResourceDetailsObjectListResult = armdevopsinfrastructure.ResourceDetailsObjectListResult{
		// 	Value: []*armdevopsinfrastructure.ResourceDetailsObject{
		// 		{
		// 			Name: to.Ptr("dd8cc705c000000"),
		// 			ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/resourceGroups/my-resource-group/providers/Microsoft.DevOpsInfrastructure/pools/my-devops-pool/resources/my-dev-ops-pool:04dcde21-626e-5a7e-8659-ce12f9284b29:dd8cc705c_0"),
		// 			Properties: &armdevopsinfrastructure.ResourceDetailsObjectProperties{
		// 				Image: to.Ptr("my-image"),
		// 				ImageVersion: to.Ptr("4.0.0"),
		// 				Status: to.Ptr(armdevopsinfrastructure.ResourceStatusReady),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dd8cc705c000001"),
		// 			ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/resourceGroups/my-resource-group/providers/Microsoft.DevOpsInfrastructure/pools/my-devops-pool/resources/my-dev-ops-pool:04dcde21-626e-5a7e-8659-ce12f9284b29:dd8cc705c_1"),
		// 			Properties: &armdevopsinfrastructure.ResourceDetailsObjectProperties{
		// 				Image: to.Ptr("my-image"),
		// 				ImageVersion: to.Ptr("4.0.0"),
		// 				Status: to.Ptr(armdevopsinfrastructure.ResourceStatusAllocated),
		// 			},
		// 	}},
		// }
	}
}
