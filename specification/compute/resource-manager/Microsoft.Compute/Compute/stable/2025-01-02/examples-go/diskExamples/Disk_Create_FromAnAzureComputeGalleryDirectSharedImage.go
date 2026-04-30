package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-01-02/diskExamples/Disk_Create_FromAnAzureComputeGalleryDirectSharedImage.json
func ExampleDisksClient_BeginCreateOrUpdate_createAManagedDiskFromAnAzureComputeGalleryDirectSharedImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDisksClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDisk", armcompute.Disk{
		Location: to.Ptr("West US"),
		Properties: &armcompute.DiskProperties{
			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
			CreationData: &armcompute.CreationData{
				CreateOption: to.Ptr(armcompute.DiskCreateOptionFromImage),
				GalleryImageReference: &armcompute.ImageDiskReference{
					SharedGalleryImageID: to.Ptr("/SharedGalleries/{sharedGalleryUniqueName}/Images/{imageName}/Versions/1.0.0"),
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
	// res = armcompute.DisksClientCreateOrUpdateResponse{
	// 	Disk: &armcompute.Disk{
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk"),
	// 		Name: to.Ptr("myDisk"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armcompute.DiskProperties{
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 			SupportedCapabilities: &armcompute.SupportedCapabilities{
	// 				AcceleratedNetwork: to.Ptr(true),
	// 			},
	// 			CreationData: &armcompute.CreationData{
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionFromImage),
	// 				GalleryImageReference: &armcompute.ImageDiskReference{
	// 					SharedGalleryImageID: to.Ptr("/SharedGalleries/{sharedGalleryUniqueName}/Images/{imageName}/Versions/1.0.0"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
