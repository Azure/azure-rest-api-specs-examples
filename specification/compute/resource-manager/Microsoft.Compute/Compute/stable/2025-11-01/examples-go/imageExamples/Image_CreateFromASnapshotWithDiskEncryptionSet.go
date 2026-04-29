package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/imageExamples/Image_CreateFromASnapshotWithDiskEncryptionSet.json
func ExampleImagesClient_BeginCreateOrUpdate_createAVirtualMachineImageFromASnapshotWithDiskEncryptionSetResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewImagesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myImage", armcompute.Image{
		Location: to.Ptr("West US"),
		Properties: &armcompute.ImageProperties{
			StorageProfile: &armcompute.ImageStorageProfile{
				OSDisk: &armcompute.ImageOSDisk{
					OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
					ManagedDisk: &armcompute.SubResource{
						ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
					},
					DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
					},
					OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
				},
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
	// res = armcompute.ImagesClientCreateOrUpdateResponse{
	// 	Image: &armcompute.Image{
	// 		Properties: &armcompute.ImageProperties{
	// 			StorageProfile: &armcompute.ImageStorageProfile{
	// 				OSDisk: &armcompute.ImageOSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
	// 					ManagedDisk: &armcompute.SubResource{
	// 						ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
	// 					},
	// 					DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
	// 					},
	// 					OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				},
	// 				DataDisks: []*armcompute.ImageDataDisk{
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Creating"),
	// 		},
	// 		Type: to.Ptr("Microsoft.Compute/images"),
	// 		Location: to.Ptr("westus"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/disk/providers/Microsoft.Compute/images/myImage"),
	// 		Name: to.Ptr("myImage"),
	// 	},
	// }
}
