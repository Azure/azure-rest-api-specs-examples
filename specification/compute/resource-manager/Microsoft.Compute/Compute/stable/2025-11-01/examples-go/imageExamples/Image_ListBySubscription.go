package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/imageExamples/Image_ListBySubscription.json
func ExampleImagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewImagesClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcompute.ImagesClientListResponse{
		// 	ImageListResult: armcompute.ImageListResult{
		// 		Value: []*armcompute.Image{
		// 			{
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/myImage"),
		// 				Name: to.Ptr("myImage"),
		// 				Type: to.Ptr("Microsoft.Compute/images"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armcompute.ImageProperties{
		// 					StorageProfile: &armcompute.ImageStorageProfile{
		// 						OSDisk: &armcompute.ImageOSDisk{
		// 							OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 							BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
		// 							Snapshot: &armcompute.SubResource{
		// 								ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
		// 							},
		// 							ManagedDisk: &armcompute.SubResource{
		// 								ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk1"),
		// 							},
		// 							OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
		// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 						},
		// 						DataDisks: []*armcompute.ImageDataDisk{
		// 							{
		// 								Lun: to.Ptr[int32](1),
		// 								BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
		// 								Snapshot: &armcompute.SubResource{
		// 									ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2"),
		// 								},
		// 								ManagedDisk: &armcompute.SubResource{
		// 									ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2"),
		// 								},
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("created"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://svchost:99/subscriptions/subscriptionId/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images?$skiptoken=token/Subscriptions/subscriptionId/ResourceGroups/myResourceGroup/UserVMImages/myImageName"),
		// 	},
		// }
	}
}
