package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5d2adf9b7fda669b4a2538c65e937ee74fe3f966/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/restorePointExamples/RestorePointCollection_ListByResourceGroup.json
func ExampleRestorePointCollectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorePointCollectionsClient().NewListPager("myResourceGroup", nil)
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
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/restorePointCollection1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.RestorePointCollectionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointCollectionID: to.Ptr("59f04a5d-f783-4200-a1bd-d3f464e8c4b4"),
		// 				Source: &armcompute.RestorePointCollectionSourceProperties{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/restorePointCollection1"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("restorePointCollection2"),
		// 			Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/restorePointCollection2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.RestorePointCollectionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RestorePointCollectionID: to.Ptr("2875c590-e337-4102-9668-4f5b7e3f98a4"),
		// 				Source: &armcompute.RestorePointCollectionSourceProperties{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/restorePointCollection2"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
