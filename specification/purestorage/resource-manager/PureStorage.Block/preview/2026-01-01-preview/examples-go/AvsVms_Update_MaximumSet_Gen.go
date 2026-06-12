package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-01-01-preview/AvsVms_Update_MaximumSet_Gen.json
func ExampleAvsVMsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvsVMsClient().BeginUpdate(ctx, "rgpurestorage", "storagepool-01", "abc123def456", armpurestorageblock.AvsVMUpdate{
		Properties: &armpurestorageblock.AvsVMUpdateProperties{
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
	// res = armpurestorageblock.AvsVMsClientUpdateResponse{
	// 	AvsVM: armpurestorageblock.AvsVM{
	// 		Properties: &armpurestorageblock.AvsVMProperties{
	// 			StoragePoolInternalID: to.Ptr("pool-abc123"),
	// 			StoragePoolResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgpurestorage/providers/PureStorage.Block/storagePools/storagepool-01"),
	// 			DisplayName: to.Ptr("AVS Virtual Machine 01"),
	// 			CreatedTimestamp: to.Ptr("2026-01-16T07:25:56.721Z"),
	// 			SoftDeletion: &armpurestorageblock.SoftDeletion{
	// 				Destroyed: to.Ptr(true),
	// 				EradicationTimestamp: to.Ptr("2026-01-23T07:25:56.721Z"),
	// 			},
	// 			VolumeContainerType: to.Ptr(armpurestorageblock.VolumeContainerTypeAVS),
	// 			Avs: &armpurestorageblock.AvsVMDetails{
	// 				VMID: to.Ptr("vm-123"),
	// 				VMName: to.Ptr("avs-vm-01"),
	// 				VMType: to.Ptr(armpurestorageblock.VMTypeVVol),
	// 				AvsVMInternalID: to.Ptr("vm-internal-123"),
	// 			},
	// 			Space: &armpurestorageblock.Space{
	// 				TotalUsed: to.Ptr[int64](1073741824),
	// 				Unique: to.Ptr[int64](268435456),
	// 				Snapshots: to.Ptr[int64](53687091236870910),
	// 				Shared: to.Ptr[int64](268435456),
	// 			},
	// 			ProvisioningState: to.Ptr(armpurestorageblock.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("avs-vm-01"),
	// 		Type: to.Ptr("PureStorage.Block/storagePools/avsVms"),
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
