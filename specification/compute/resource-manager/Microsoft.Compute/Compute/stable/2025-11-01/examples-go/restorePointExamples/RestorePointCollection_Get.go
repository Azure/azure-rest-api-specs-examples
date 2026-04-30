package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/restorePointExamples/RestorePointCollection_Get.json
func ExampleRestorePointCollectionsClient_Get_getARestorePointCollectionButNotTheRestorePointsContainedInTheRestorePointCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().Get(ctx, "myResourceGroup", "myRpc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.RestorePointCollectionsClientGetResponse{
	// 	RestorePointCollection: &armcompute.RestorePointCollection{
	// 		Name: to.Ptr("myRpc"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/myRpc"),
	// 		Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"myTag1": to.Ptr("tagValue1"),
	// 		},
	// 		Properties: &armcompute.RestorePointCollectionProperties{
	// 			Source: &armcompute.RestorePointCollectionSourceProperties{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 				Location: to.Ptr("eastus"),
	// 			},
	// 			InstantAccess: to.Ptr(true),
	// 			RestorePointCollectionID: to.Ptr("59f04a5d-f783-4200-a1bd-d3f464e8c4b4"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
