package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/restorePointExamples/RestorePointCollection_CreateOrUpdate_ForCrossRegionCopy.json
func ExampleRestorePointCollectionsClient_CreateOrUpdate_createOrUpdateARestorePointCollectionForCrossRegionCopy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().CreateOrUpdate(ctx, "myResourceGroup", "myRpc", armcompute.RestorePointCollection{
		Location: to.Ptr("norwayeast"),
		Tags: map[string]*string{
			"myTag1": to.Ptr("tagValue1"),
		},
		Properties: &armcompute.RestorePointCollectionProperties{
			Source: &armcompute.RestorePointCollectionSourceProperties{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePointCollection = armcompute.RestorePointCollection{
	// 	Name: to.Ptr("myRpc"),
	// 	Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/myRpc"),
	// 	Location: to.Ptr("norwayeast"),
	// 	Tags: map[string]*string{
	// 		"myTag1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armcompute.RestorePointCollectionProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RestorePointCollectionID: to.Ptr("638f052b-a7c2-450c-89e7-6a3b8f1d6a7c"),
	// 		Source: &armcompute.RestorePointCollectionSourceProperties{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 			Location: to.Ptr("eastus"),
	// 		},
	// 	},
	// }
}
