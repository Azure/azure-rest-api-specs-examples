package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-01-01-preview/Volumes_ListByVolumeGroup_MaximumSet_Gen.json
func ExampleVolumesClient_NewListByVolumeGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumesClient().NewListByVolumeGroupPager("rgpurestorage", "storagepool-01", "volumegroup-01", nil)
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
		// page = armpurestorageblock.VolumesClientListByVolumeGroupResponse{
		// 	VolumeListResult: armpurestorageblock.VolumeListResult{
		// 		Value: []*armpurestorageblock.Volume{
		// 			{
		// 				Properties: &armpurestorageblock.AzureVolumeProperties{
		// 					Space: &armpurestorageblock.Space{
		// 						TotalUsed: to.Ptr[int64](1073741824),
		// 						Unique: to.Ptr[int64](536870912),
		// 						Snapshots: to.Ptr[int64](268435456),
		// 						Shared: to.Ptr[int64](268435456),
		// 					},
		// 					ProvisionedSize: to.Ptr[int64](10737418240),
		// 					SerialNumber: to.Ptr("PURE-VOL-ABC123XYZ"),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armpurestorageblock.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgpurestorage/providers/PureStorage.Block/storagePools/storagepool-01/volumeGroups/volumegroup-01/volumes/volume-01"),
		// 				Name: to.Ptr("volume-01"),
		// 				Type: to.Ptr("PureStorage.Block/storagePools/volumeGroups/volumes"),
		// 				SystemData: &armpurestorageblock.SystemData{
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
