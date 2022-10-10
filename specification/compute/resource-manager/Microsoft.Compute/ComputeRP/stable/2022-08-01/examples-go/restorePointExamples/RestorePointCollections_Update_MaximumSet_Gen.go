package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/restorePointExamples/RestorePointCollections_Update_MaximumSet_Gen.json
func ExampleRestorePointCollectionsClient_Update_restorePointCollectionsUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewRestorePointCollectionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaa", armcompute.RestorePointCollectionUpdate{
		Tags: map[string]*string{
			"key8536": to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		},
		Properties: &armcompute.RestorePointCollectionProperties{
			Source: &armcompute.RestorePointCollectionSourceProperties{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
