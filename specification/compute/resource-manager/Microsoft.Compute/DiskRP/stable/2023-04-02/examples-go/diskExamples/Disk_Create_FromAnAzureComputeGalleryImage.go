package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskExamples/Disk_Create_FromAnAzureComputeGalleryImage.json
func ExampleDisksClient_BeginCreateOrUpdate_createAManagedDiskFromAnAzureComputeGalleryImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDisksClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDisk", armcompute.Disk{
		Location: to.Ptr("West US"),
		Properties: &armcompute.DiskProperties{
			CreationData: &armcompute.CreationData{
				CreateOption: to.Ptr(armcompute.DiskCreateOptionFromImage),
				GalleryImageReference: &armcompute.ImageDiskReference{
					ID: to.Ptr("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Providers/Microsoft.Compute/Galleries/{galleryName}/Images/{imageName}/Versions/1.0.0"),
				},
			},
			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
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
	// res.Disk = armcompute.Disk{
	// 	Name: to.Ptr("myDisk"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.DiskProperties{
	// 		CreationData: &armcompute.CreationData{
	// 			CreateOption: to.Ptr(armcompute.DiskCreateOptionFromImage),
	// 			GalleryImageReference: &armcompute.ImageDiskReference{
	// 				ID: to.Ptr("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Providers/Microsoft.Compute/Galleries/{galleryName}/Images/{imageName}/Versions/1.0.0"),
	// 			},
	// 		},
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SupportedCapabilities: &armcompute.SupportedCapabilities{
	// 			AcceleratedNetwork: to.Ptr(true),
	// 		},
	// 	},
	// }
}
