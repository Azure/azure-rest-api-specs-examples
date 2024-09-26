package armfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/fabric/resource-manager/Microsoft.Fabric/stable/2023-11-01/examples/FabricCapacities_ListBySubscription.json
func ExampleCapacitiesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacitiesClient().NewListBySubscriptionPager(nil)
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
		// page.CapacityListResult = armfabric.CapacityListResult{
		// 	Value: []*armfabric.Capacity{
		// 		{
		// 			Name: to.Ptr("azsdktest"),
		// 			Type: to.Ptr("Microsoft.Fabric/capacities"),
		// 			ID: to.Ptr("/subscriptions/548B7FB7-3B2A-4F46-BB02-66473F1FC22C/resourceGroups/TestRG/providers/Microsoft.Fabric/capacities/azsdktest"),
		// 			Location: to.Ptr("West Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armfabric.CapacityProperties{
		// 				Administration: &armfabric.CapacityAdministration{
		// 					Members: []*string{
		// 						to.Ptr("azsdktest@microsoft.com")},
		// 					},
		// 					ProvisioningState: to.Ptr(armfabric.ProvisioningStateSucceeded),
		// 					State: to.Ptr(armfabric.ResourceStateActive),
		// 				},
		// 				SKU: &armfabric.RpSKU{
		// 					Name: to.Ptr("F2"),
		// 					Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("azsdktest2"),
		// 				Type: to.Ptr("Microsoft.Fabric/capacities"),
		// 				ID: to.Ptr("/subscriptions/548B7FB7-3B2A-4F46-BB02-66473F1FC22C/resourceGroups/TestRG/providers/Microsoft.Fabric/capacities/azsdktest2"),
		// 				Location: to.Ptr("West Central US"),
		// 				Tags: map[string]*string{
		// 					"testKey": to.Ptr("testValue"),
		// 				},
		// 				Properties: &armfabric.CapacityProperties{
		// 					Administration: &armfabric.CapacityAdministration{
		// 						Members: []*string{
		// 							to.Ptr("azsdktest@microsoft.com")},
		// 						},
		// 						ProvisioningState: to.Ptr(armfabric.ProvisioningStateSucceeded),
		// 						State: to.Ptr(armfabric.ResourceStateProvisioning),
		// 					},
		// 					SKU: &armfabric.RpSKU{
		// 						Name: to.Ptr("F4"),
		// 						Tier: to.Ptr(armfabric.RpSKUTierFabric),
		// 					},
		// 			}},
		// 		}
	}
}
