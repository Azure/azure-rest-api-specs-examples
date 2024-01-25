package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Update_MaximumSet_Gen.json
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
	// 	Name: to.Ptr("o"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 	ID: to.Ptr("swkcmwglncgtsnejzvldnbpsifxez"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bcclmbseed"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("ARM Id of Resource"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("mtkeip"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](9),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("izdwogzjedsfug"),
	// 			TargetPortalHostname: to.Ptr("wyfbjobugmad"),
	// 			TargetPortalPort: to.Ptr[int32](21),
	// 		},
	// 		VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
	// 	},
	// }
}
