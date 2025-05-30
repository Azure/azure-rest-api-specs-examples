package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/27046dbff974e3901970aa53b29cec6d8ec1342a/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/Volumes_Get_MaximumSet_Gen.json
func ExampleVolumesClient_Get_volumesGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("wrrvufwzhwd"),
	// 	Type: to.Ptr("kzaqluwzgssofiulhkxmzafsdhcno"),
	// 	ID: to.Ptr("xsqrx"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-02T20:06:12.883Z"); return t}()),
	// 		CreatedBy: to.Ptr("ndexrszfpxfmlbjzollrgzhhae"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-02T20:06:12.884Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("hdqkgsdybfsl"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("owsp"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("gyqwvotwkluuzzpuedccamwfvasf"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](25),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("t"),
	// 			TargetPortalHostname: to.Ptr("ptulwr"),
	// 			TargetPortalPort: to.Ptr[int32](26),
	// 		},
	// 		VolumeID: to.Ptr("qbowllkrvxsnpplweplsmfgncplpu"),
	// 	},
	// }
}
