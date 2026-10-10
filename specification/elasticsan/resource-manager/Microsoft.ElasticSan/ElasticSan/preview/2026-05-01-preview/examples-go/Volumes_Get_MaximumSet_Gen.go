package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan/v2"
)

// Generated from example definition: 2026-05-01-preview/Volumes_Get_MaximumSet_Gen.json
func ExampleVolumesClient_Get_volumesGetMaximumSetGen() {
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
	// 	Volume: armelasticsan.Volume{
	// 		Name: to.Ptr("wrrvufwzhwd"),
	// 		Type: to.Ptr("kzaqluwzgssofiulhkxmzafsdhcno"),
	// 		ID: to.Ptr("xsqrx"),
	// 		Properties: &armelasticsan.VolumeProperties{
	// 			CreationData: &armelasticsan.SourceCreationData{
	// 				CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 				SourceID: to.Ptr("owsp"),
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
	// 			SizeGiB: to.Ptr[int64](25),
	// 			StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 				ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
	// 				Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 				TargetIqn: to.Ptr("t"),
	// 				TargetPortalHostname: to.Ptr("ptulwr"),
	// 				TargetPortalPort: to.Ptr[int32](26),
	// 			},
	// 			VolumeID: to.Ptr("qbowllkrvxsnpplweplsmfgncplpu"),
	// 		},
	// 		SystemData: &armelasticsan.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2024, time.May, 2, 20, 6, 12, 883000000, time.UTC)),
	// 			CreatedBy: to.Ptr("ndexrszfpxfmlbjzollrgzhhae"),
	// 			CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.May, 2, 20, 6, 12, 884000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("hdqkgsdybfsl"),
	// 			LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
