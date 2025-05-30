package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/27046dbff974e3901970aa53b29cec6d8ec1342a/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/VolumeSnapshots_Get_MinimumSet_Gen.json
func ExampleVolumeSnapshotsClient_Get_volumeSnapshotsGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeSnapshotsClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "snapshotname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Snapshot = armelasticsan.Snapshot{
	// 	Name: to.Ptr("qukfugetqthsufp"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/snapshots"),
	// 	ID: to.Ptr("bbqqgzxagggqgkdgjqq"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		CreatedBy: to.Ptr("f"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("wifrlzeszzcckvdzdwxhvservlvarp"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.SnapshotProperties{
	// 		CreationData: &armelasticsan.SnapshotCreationData{
	// 			SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SourceVolumeSizeGiB: to.Ptr[int64](28),
	// 		VolumeName: to.Ptr("volumename"),
	// 	},
	// }
}
