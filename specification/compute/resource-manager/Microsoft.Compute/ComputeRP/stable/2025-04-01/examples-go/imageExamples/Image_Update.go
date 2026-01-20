package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/imageExamples/Image_Update.json
func ExampleImagesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewImagesClient().BeginUpdate(ctx, "myResourceGroup", "myImage", armcompute.ImageUpdate{
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
		},
		Properties: &armcompute.ImageProperties{
			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypesV1),
			SourceVirtualMachine: &armcompute.SubResource{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Image = armcompute.Image{
	// 	Name: to.Ptr("myImage"),
	// 	Type: to.Ptr("Microsoft.Compute/images"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/myImage"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.ImageProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		StorageProfile: &armcompute.ImageStorageProfile{
	// 			DataDisks: []*armcompute.ImageDataDisk{
	// 				{
	// 					BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
	// 					ManagedDisk: &armcompute.SubResource{
	// 						ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2"),
	// 					},
	// 					Snapshot: &armcompute.SubResource{
	// 						ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2"),
	// 					},
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					Lun: to.Ptr[int32](1),
	// 			}},
	// 			OSDisk: &armcompute.ImageOSDisk{
	// 				BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
	// 				DiskSizeGB: to.Ptr[int32](20),
	// 				ManagedDisk: &armcompute.SubResource{
	// 					ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk1"),
	// 				},
	// 				Snapshot: &armcompute.SubResource{
	// 					ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
	// 				},
	// 				StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 				OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			},
	// 			ZoneResilient: to.Ptr(true),
	// 		},
	// 	},
	// }
}
