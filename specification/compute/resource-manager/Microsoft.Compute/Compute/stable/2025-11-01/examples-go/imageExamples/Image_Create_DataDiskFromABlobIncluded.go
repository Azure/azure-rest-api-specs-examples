package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/imageExamples/Image_Create_DataDiskFromABlobIncluded.json
func ExampleImagesClient_BeginCreateOrUpdate_createAVirtualMachineImageThatIncludesADataDiskFromABlob() {
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
					OSType:  to.Ptr(armcompute.OperatingSystemTypesLinux),
					BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
					OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
				},
				DataDisks: []*armcompute.ImageDataDisk{
					{
						Lun:     to.Ptr[int32](1),
						BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
					},
				},
				ZoneResilient: to.Ptr(false),
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
	// 					OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 					BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				},
	// 				DataDisks: []*armcompute.ImageDataDisk{
	// 					{
	// 						Lun: to.Ptr[int32](1),
	// 						BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
	// 					},
	// 				},
	// 				ZoneResilient: to.Ptr(false),
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
