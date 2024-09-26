package armfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/fabric/resource-manager/Microsoft.Fabric/stable/2023-11-01/examples/FabricCapacities_ListSkusForCapacity.json
func ExampleCapacitiesClient_NewListSKUsForCapacityPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacitiesClient().NewListSKUsForCapacityPager("TestRG", "azsdktest", nil)
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
		// page.RpSKUEnumerationForExistingResourceResult = armfabric.RpSKUEnumerationForExistingResourceResult{
		// 	Value: []*armfabric.RpSKUDetailsForExistingResource{
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F16"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F8"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F64"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F1024"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F128"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F2"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F256"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F32"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F4"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F512"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 		},
		// 		{
		// 			ResourceType: to.Ptr("Microsoft.Fabric/capacities"),
		// 			SKU: &armfabric.RpSKU{
		// 				Name: to.Ptr("F2048"),
		// 				Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 			},
		// 	}},
		// }
	}
}
