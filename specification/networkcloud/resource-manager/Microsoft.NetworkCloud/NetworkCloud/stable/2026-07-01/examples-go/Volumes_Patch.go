package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud/v2"
)

// Generated from example definition: 2026-07-01/Volumes_Patch.json
func ExampleVolumesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumesClient().Update(ctx, "resourceGroupName", "volumeName", armnetworkcloud.VolumePatchParameters{
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworkcloud.VolumesClientUpdateResponse{
	// 	Volume: armnetworkcloud.Volume{
	// 		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 			Type: to.Ptr("CustomLocation"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName"),
	// 		Location: to.Ptr("location"),
	// 		Name: to.Ptr("volumeName"),
	// 		Properties: &armnetworkcloud.VolumeProperties{
	// 			AllocatedSizeMiB: to.Ptr[int64](1048576),
	// 			AssignedStorageApplianceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/storageAppliances/storageApplianceName"),
	// 			AttachedTo: []*string{
	// 				to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName"),
	// 			},
	// 			DetailedStatus: to.Ptr(armnetworkcloud.VolumeDetailedStatusActive),
	// 			DetailedStatusMessage: to.Ptr("Volume is available"),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.VolumeProvisioningStateSucceeded),
	// 			SerialNumber: to.Ptr("XXXX-XXXX-XXXX-XXXX"),
	// 			SizeMiB: to.Ptr[int64](1048576),
	// 			StorageApplianceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/storageAppliances/storageApplianceName"),
	// 		},
	// 		SystemData: &armnetworkcloud.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2021, time.January, 22, 13, 27, 3, 8000000, time.UTC)),
	// 			CreatedBy: to.Ptr("identityA"),
	// 			CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(time.Date(2021, time.January, 22, 13, 29, 3, 1000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("identityB"),
	// 			LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("myvalue1"),
	// 			"key2": to.Ptr("myvalue2"),
	// 		},
	// 		Type: to.Ptr("Microsoft.NetworkCloud/volumes"),
	// 	},
	// }
}
