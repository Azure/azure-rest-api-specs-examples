package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/restorePointExamples/RestorePointCollection_ListBySubscription.json
func ExampleRestorePointCollectionsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorePointCollectionsClient().NewListAllPager(nil)
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
		// page.RestorePointCollectionListResult = armcompute.RestorePointCollectionListResult{
		// 	Value: []*armcompute.RestorePointCollection{
		// 		{
		// 			Name: to.Ptr("restorePointCollection1"),
		// 			Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resourceGroup1/providers/Microsoft.Compute/restorePointCollections/restorePointCollection1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.RestorePointCollectionProperties{
		// 				InstantAccess: to.Ptr(true),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointCollectionID: to.Ptr("59f04a5d-f783-4200-a1bd-d3f464e8c4b4"),
		// 				Source: &armcompute.RestorePointCollectionSourceProperties{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/VM_Test"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("restorePointCollection2"),
		// 			Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resourceGroup2/providers/Microsoft.Compute/restorePointCollections/restorePointCollection2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.RestorePointCollectionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointCollectionID: to.Ptr("2875c590-e337-4102-9668-4f5b7e3f98a4"),
		// 				Source: &armcompute.RestorePointCollectionSourceProperties{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/VM_Prod"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
