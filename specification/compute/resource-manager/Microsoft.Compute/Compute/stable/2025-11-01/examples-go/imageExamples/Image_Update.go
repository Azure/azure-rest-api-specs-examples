package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/imageExamples/Image_Update.json
func ExampleImagesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewImagesClient().BeginUpdate(ctx, "myResourceGroup", "myImage", armcompute.ImageUpdate{
		Properties: &armcompute.ImageProperties{
			SourceVirtualMachine: &armcompute.SubResource{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
			},
			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypesV1),
		},
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
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
	// res = armcompute.ImagesClientUpdateResponse{
	// 	Image: &armcompute.Image{
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/myImage"),
	// 		Name: to.Ptr("myImage"),
	// 		Type: to.Ptr("Microsoft.Compute/images"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("HR"),
	// 		},
	// 		Properties: &armcompute.ImageProperties{
	// 			StorageProfile: &armcompute.ImageStorageProfile{
	// 				OSDisk: &armcompute.ImageOSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 					BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
	// 					Snapshot: &armcompute.SubResource{
	// 						ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
	// 					},
	// 					ManagedDisk: &armcompute.SubResource{
	// 						ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk1"),
	// 					},
	// 					OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					DiskSizeGB: to.Ptr[int32](20),
	// 				},
	// 				DataDisks: []*armcompute.ImageDataDisk{
	// 					{
	// 						Lun: to.Ptr[int32](1),
	// 						BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
	// 						Snapshot: &armcompute.SubResource{
	// 							ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2"),
	// 						},
	// 						ManagedDisk: &armcompute.SubResource{
	// 							ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2"),
	// 						},
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 				},
	// 				ZoneResilient: to.Ptr(true),
	// 			},
	// 			ProvisioningState: to.Ptr("created"),
	// 		},
	// 	},
	// }
}
