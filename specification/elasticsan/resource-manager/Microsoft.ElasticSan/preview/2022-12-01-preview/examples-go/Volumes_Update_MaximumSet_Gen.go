package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/Volumes_Update_MaximumSet_Gen.json
func ExampleVolumesClient_BeginUpdate_volumesUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginUpdate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", armelasticsan.VolumeUpdate{
		Properties: &armelasticsan.VolumeUpdateProperties{
			SizeGiB: to.Ptr[int64](11),
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
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("ipmipjggodoxjua"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-04T05:52:15.371Z"); return t}()),
	// 		CreatedBy: to.Ptr("zloujisrarsox"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-04T05:52:15.371Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("lmdztjjcmdfzam"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr("None"),
	// 			SourceURI: to.Ptr(" ARM Id of resource"),
	// 		},
	// 		SizeGiB: to.Ptr[int64](23),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("yuuex"),
	// 			TargetPortalHostname: to.Ptr("yktnmynpdjpmztwuwcvipp"),
	// 			TargetPortalPort: to.Ptr[int32](3),
	// 		},
	// 		VolumeID: to.Ptr("hkbhievbyopwjqawbyvqi"),
	// 	},
	// }
}
