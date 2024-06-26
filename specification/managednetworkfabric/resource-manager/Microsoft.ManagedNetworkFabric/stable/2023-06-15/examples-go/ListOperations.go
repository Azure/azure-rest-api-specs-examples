package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armmanagednetworkfabric.OperationListResult{
		// 	Value: []*armmanagednetworkfabric.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagedNetworkFabric/NetworkFabricControllers/Read"),
		// 			ActionType: to.Ptr(armmanagednetworkfabric.ActionTypeInternal),
		// 			Display: &armmanagednetworkfabric.OperationDisplay{
		// 				Description: to.Ptr("Gets/List the NetworkFabricController resource data."),
		// 				Operation: to.Ptr("Gets/List NetworkFabricController resources."),
		// 				Provider: to.Ptr("Microsoft.ManagedNetworkFabric resource provider"),
		// 				Resource: to.Ptr("NetworkFabricControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armmanagednetworkfabric.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagedNetworkFabric/NetworkFabricControllers/Write"),
		// 			ActionType: to.Ptr(armmanagednetworkfabric.ActionTypeInternal),
		// 			Display: &armmanagednetworkfabric.OperationDisplay{
		// 				Description: to.Ptr("Create or Update NetworkFabricController resource data."),
		// 				Operation: to.Ptr("Create or Update NetworkFabricController resource."),
		// 				Provider: to.Ptr("Microsoft.ManagedNetworkFabric resource provider"),
		// 				Resource: to.Ptr("NetworkFabricControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armmanagednetworkfabric.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ManagedNetworkFabric/NetworkFabricControllers/Delete"),
		// 			ActionType: to.Ptr(armmanagednetworkfabric.ActionTypeInternal),
		// 			Display: &armmanagednetworkfabric.OperationDisplay{
		// 				Description: to.Ptr("Deletes the NetworkFabricController resource."),
		// 				Operation: to.Ptr("Deletes the NetworkFabricController resource."),
		// 				Provider: to.Ptr("Microsoft.ManagedNetworkFabric resource provider"),
		// 				Resource: to.Ptr("NetworkFabricControllers"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armmanagednetworkfabric.OriginUserSystem),
		// 	}},
		// }
	}
}
