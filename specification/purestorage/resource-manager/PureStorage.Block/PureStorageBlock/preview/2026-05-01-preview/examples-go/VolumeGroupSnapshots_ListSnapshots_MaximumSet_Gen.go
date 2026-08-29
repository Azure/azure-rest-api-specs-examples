package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-05-01-preview/VolumeGroupSnapshots_ListSnapshots_MaximumSet_Gen.json
func ExampleVolumeGroupSnapshotsClient_ListSnapshots() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeGroupSnapshotsClient().ListSnapshots(ctx, "rgpurestorage", "storagepool-01", "volumegroup-01", armpurestorageblock.VolumeGroupSnapshotListRequest{
		Filter:  to.Ptr("substringof('snapshot', name)"),
		Orderby: to.Ptr("name asc"),
		Top:     to.Ptr[int32](10),
		Skip:    to.Ptr[int32](0),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.VolumeGroupSnapshotsClientListSnapshotsResponse{
	// 	VolumeGroupSnapshotPostListResult: armpurestorageblock.VolumeGroupSnapshotPostListResult{
	// 		Value: []*armpurestorageblock.VolumeGroupSnapshot{
	// 			{
	// 				Properties: &armpurestorageblock.VolumeGroupSnapshotProperties{
	// 					CreatedAt: to.Ptr(time.Date(2026, time.February, 10, 8, 0, 0, 0, time.UTC)),
	// 					CreatedByPolicy: to.Ptr(false),
	// 					SoftDeletion: &armpurestorageblock.DestroyedStateProperties{
	// 						Destroyed: to.Ptr(false),
	// 					},
	// 					Space: &armpurestorageblock.Space{
	// 						TotalUsed: to.Ptr[int64](524288000),
	// 						Unique: to.Ptr[int64](262144000),
	// 						Snapshots: to.Ptr[int64](131072000),
	// 						Shared: to.Ptr[int64](131072000),
	// 					},
	// 					VolumeSnapshots: []*armpurestorageblock.VolumeSnapshotInfo{
	// 					},
	// 					ProvisioningState: to.Ptr(armpurestorageblock.ProvisioningStateSucceeded),
	// 				},
	// 				ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgpurestorage/providers/PureStorage.Block/storagePools/storagepool-01/volumeGroups/volumegroup-01/snapshots/snapshot-01"),
	// 				Name: to.Ptr("snapshot-01"),
	// 				Type: to.Ptr("PureStorage.Block/storagePools/volumeGroups/snapshots"),
	// 				SystemData: &armpurestorageblock.SystemData{
	// 					CreatedBy: to.Ptr("user@contoso.com"),
	// 					CreatedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 					CreatedAt: to.Ptr(time.Date(2026, time.February, 10, 8, 0, 0, 0, time.UTC)),
	// 					LastModifiedBy: to.Ptr("user@contoso.com"),
	// 					LastModifiedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(time.Date(2026, time.February, 10, 8, 0, 0, 0, time.UTC)),
	// 				},
	// 			},
	// 		},
	// 		Count: to.Ptr[int32](1),
	// 		TotalCount: to.Ptr[int32](23),
	// 	},
	// }
}
