package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: 2025-09-01/Volumes_Get_MinimumSet_Gen.json
func ExampleVolumesClient_Get_volumesGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumesClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armelasticsan.VolumesClientGetResponse{
	// 	Volume: &armelasticsan.Volume{
	// 		Name: to.Ptr("o"),
	// 		Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 		ID: to.Ptr("swkcmwglncgtsnejzvldnbpsifxez"),
	// 		Properties: &armelasticsan.VolumeProperties{
	// 			CreationData: &armelasticsan.SourceCreationData{
	// 				CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 				SourceID: to.Ptr("ARM Id of Resource"),
	// 			},
	// 			ManagedBy: &armelasticsan.ManagedByInfo{
	// 				ResourceID: to.Ptr("mtkeip"),
	// 			},
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
	// 			SizeGiB: to.Ptr[int64](9),
	// 			StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 				ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 				Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 				TargetIqn: to.Ptr("izdwogzjedsfug"),
	// 				TargetPortalHostname: to.Ptr("wyfbjobugmad"),
	// 				TargetPortalPort: to.Ptr[int32](21),
	// 			},
	// 			VolumeID: to.Ptr("umwjlxntntjejiyrywrytkzbfbluhk"),
	// 		},
	// 		SystemData: &armelasticsan.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 			CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 			CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("bcclmbseed"),
	// 			LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
