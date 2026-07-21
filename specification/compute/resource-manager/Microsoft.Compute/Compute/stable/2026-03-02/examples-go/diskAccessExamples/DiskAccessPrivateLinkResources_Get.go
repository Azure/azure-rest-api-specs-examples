package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-02/diskAccessExamples/DiskAccessPrivateLinkResources_Get.json
func ExampleDiskAccessesClient_GetPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiskAccessesClient().GetPrivateLinkResources(ctx, "myResourceGroup", "myDiskAccess", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DiskAccessesClientGetPrivateLinkResourcesResponse{
	// 	PrivateLinkResourceListResult: armcompute.PrivateLinkResourceListResult{
	// 		Value: []*armcompute.PrivateLinkResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess/privateLinkResources/disks"),
	// 				Name: to.Ptr("disks"),
	// 				Type: to.Ptr("Microsoft.Compute/diskAccesses/privateLinkResources"),
	// 				Properties: &armcompute.PrivateLinkResourceProperties{
	// 					GroupID: to.Ptr("disks"),
	// 					RequiredMembers: []*string{
	// 						to.Ptr("diskAccess_1"),
	// 					},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.blob.core.windows.net"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
