package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: 2025-09-01/VolumeSnapshots_ListByVolumeGroup_MaximumSet_Gen.json
func ExampleVolumeSnapshotsClient_NewListByVolumeGroupPager_volumeSnapshotsListByVolumeGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeSnapshotsClient().NewListByVolumeGroupPager("resourcegroupname", "elasticsanname", "volumegroupname", &armelasticsan.VolumeSnapshotsClientListByVolumeGroupOptions{
		Filter: to.Ptr("volumeName eq <volume name>")})
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
		// page = armelasticsan.VolumeSnapshotsClientListByVolumeGroupResponse{
		// 	SnapshotList: armelasticsan.SnapshotList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionid/resourceGroups/resourcegroupname/providers/Microsoft.ElasticSan/elasticSans/elasticsanname/volumegroups/volumegroupname/snapshots?api-version=2024-07-01-preview&%24skiptoken=def123ghi456"),
		// 		Value: []*armelasticsan.Snapshot{
		// 			{
		// 				Name: to.Ptr("qukfugetqthsufp"),
		// 				Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/snapshots"),
		// 				ID: to.Ptr("bbqqgzxagggqgkdgjqq"),
		// 				Properties: &armelasticsan.SnapshotProperties{
		// 					CreationData: &armelasticsan.SnapshotCreationData{
		// 						SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
		// 					},
		// 					ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 					SourceVolumeSizeGiB: to.Ptr[int64](28),
		// 					VolumeName: to.Ptr("volumename"),
		// 				},
		// 				SystemData: &armelasticsan.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
		// 					CreatedBy: to.Ptr("f"),
		// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("wifrlzeszzcckvdzdwxhvservlvarp"),
		// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
