package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskAccessExamples/DiskAccessPrivateLinkResources_Get.json
func ExampleDiskAccessesClient_GetPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.PrivateLinkResourceListResult = armcompute.PrivateLinkResourceListResult{
	// 	Value: []*armcompute.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("disks"),
	// 			Type: to.Ptr("Microsoft.Compute/diskAccesses/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess/privateLinkResources/disks"),
	// 			Properties: &armcompute.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("disks"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("diskAccess_1")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.blob.core.windows.net")},
	// 					},
	// 			}},
	// 		}
}
