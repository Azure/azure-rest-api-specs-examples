package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-01-01-preview/AvsStorageContainerVolumes_Update_MaximumSet_Gen.json
func ExampleAvsStorageContainerVolumesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvsStorageContainerVolumesClient().BeginUpdate(ctx, "rgpurestorage", "storagepool-01", "container-01", "a1b2c3d4-e5f6", armpurestorageblock.AvsStorageContainerVolumeUpdate{
		Properties: &armpurestorageblock.AvsStorageContainerVolumeUpdateProperties{
			SoftDeletion: &armpurestorageblock.SoftDeletion{
				Destroyed: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.AvsStorageContainerVolumesClientUpdateResponse{
	// 	AvsStorageContainerVolume: armpurestorageblock.AvsStorageContainerVolume{
	// 		Properties: &armpurestorageblock.VolumeProperties{
	// 			StoragePoolInternalID: to.Ptr("pool-abc123"),
	// 			StoragePoolResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgpurestorage/providers/PureStorage.Block/storagePools/storagepool-01"),
	// 			VolumeInternalID: to.Ptr("vol-xyz789"),
	// 			DisplayName: to.Ptr("AVS VM Volume 01"),
	// 			Space: &armpurestorageblock.Space{
	// 				TotalUsed: to.Ptr[int64](1073741824),
	// 				Unique: to.Ptr[int64](268435456),
	// 				Snapshots: to.Ptr[int64](53687091236870910),
	// 				Shared: to.Ptr[int64](268435456),
	// 			},
	// 			SoftDeletion: &armpurestorageblock.SoftDeletion{
	// 				Destroyed: to.Ptr(true),
	// 				EradicationTimestamp: to.Ptr("2026-01-23T07:25:56.721Z"),
	// 			},
	// 			CreatedTimestamp: to.Ptr("2026-01-16T07:25:56.721Z"),
	// 			ProvisionedSize: to.Ptr[int64](10737418240),
	// 			VolumeType: to.Ptr(armpurestorageblock.VolumeTypeAVS),
	// 			Avs: &armpurestorageblock.AvsDiskDetails{
	// 				DiskID: to.Ptr("disk-abc123"),
	// 				DiskName: to.Ptr("avs-disk-01"),
	// 				Folder: to.Ptr("/Datacenter/vm"),
	// 				AvsVMInternalID: to.Ptr("vm-internal-123"),
	// 				AvsVMResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgpurestorage/providers/Microsoft.AVS/privateClouds/avs-cloud-01/clusters/cluster-01/virtualMachines/vm-01"),
	// 				AvsVMName: to.Ptr("avs-vm-01"),
	// 				AvsStorageContainerResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgpurestorage/providers/PureStorage.Block/storagePools/storagepool-01/avsStorageContainers/container-01"),
	// 			},
	// 			ProvisioningState: to.Ptr(armpurestorageblock.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("volume-01"),
	// 		Type: to.Ptr("PureStorage.Block/storagePools/avsStorageContainers/volumes"),
	// 		SystemData: &armpurestorageblock.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
	// 		},
	// 	},
	// }
}
