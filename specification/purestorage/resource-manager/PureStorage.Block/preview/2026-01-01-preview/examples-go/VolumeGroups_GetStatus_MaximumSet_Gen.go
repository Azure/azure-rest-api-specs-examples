package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-01-01-preview/VolumeGroups_GetStatus_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_GetStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeGroupsClient().GetStatus(ctx, "rgpurestorage", "storagepool-01", "volumegroup-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.VolumeGroupsClientGetStatusResponse{
	// 	VolumeGroupStatus: armpurestorageblock.VolumeGroupStatus{
	// 		Space: &armpurestorageblock.Space{
	// 			TotalUsed: to.Ptr[int64](5368709120),
	// 			Unique: to.Ptr[int64](2684354560),
	// 			Snapshots: to.Ptr[int64](1342177280),
	// 			Shared: to.Ptr[int64](1342177280),
	// 		},
	// 		ConnectedHostCount: to.Ptr[int32](3),
	// 	},
	// }
}
