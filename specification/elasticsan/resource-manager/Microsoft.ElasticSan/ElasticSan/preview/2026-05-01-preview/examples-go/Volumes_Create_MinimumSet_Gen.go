package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan/v2"
)

// Generated from example definition: 2026-05-01-preview/Volumes_Create_MinimumSet_Gen.json
func ExampleVolumesClient_BeginCreate_volumesCreateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", armelasticsan.Volume{
		Properties: &armelasticsan.VolumeProperties{
			SizeGiB: to.Ptr[int64](9),
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
	// res = armelasticsan.VolumesClientCreateResponse{
	// 	Volume: armelasticsan.Volume{
	// 		Name: to.Ptr("o"),
	// 		Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/volumes"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 		Properties: &armelasticsan.VolumeProperties{
	// 			CreationData: &armelasticsan.SourceCreationData{
	// 				CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 				SourceID: to.Ptr("ARM Id of Resource"),
	// 			},
	// 			ManagedBy: []*armelasticsan.ManagedByResources{
	// 				{
	// 					ClientID: to.Ptr("pclpkrpkpmvcsegcubrakcoodrubo"),
	// 					Version: to.Ptr[int32](1),
	// 					ResourceIDs: []*string{
	// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.SomeProvider/someResource/myResource"),
	// 					},
	// 				},
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
	// 			CreatedAt: to.Ptr(time.Date(2023, time.August, 23, 12, 16, 10, 57000000, time.UTC)),
	// 			CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 			CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2023, time.August, 23, 12, 16, 10, 57000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("bcclmbseed"),
	// 			LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
