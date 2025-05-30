package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/27046dbff974e3901970aa53b29cec6d8ec1342a/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/Volumes_Create_MaximumSet_Gen.json
func ExampleVolumesClient_BeginCreate_volumesCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", armelasticsan.Volume{
		Properties: &armelasticsan.VolumeProperties{
			CreationData: &armelasticsan.SourceCreationData{
				CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
				SourceID:     to.Ptr("mdonegivjquite"),
			},
			ManagedBy: &armelasticsan.ManagedByInfo{
				ResourceID: to.Ptr("pclpkrpkpmvcsegcubrakcoodrubo"),
			},
			SizeGiB: to.Ptr[int64](23),
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
	// 	Name: to.Ptr("hxm"),
	// 	Type: to.Ptr("sxsbsdhngvbcpxxcuvyt"),
	// 	ID: to.Ptr("sipsx"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T14:22:13.546Z"); return t}()),
	// 		CreatedBy: to.Ptr("bpuxtfzqwdhifevjtucoc"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T14:22:13.547Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("ourjjlolgugpxnkbiegumkicksibep"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
	// 			SourceID: to.Ptr("mdonegivjquite"),
	// 		},
	// 		ManagedBy: &armelasticsan.ManagedByInfo{
	// 			ResourceID: to.Ptr("pclpkrpkpmvcsegcubrakcoodrubo"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SizeGiB: to.Ptr[int64](23),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("us"),
	// 			TargetPortalHostname: to.Ptr("oqtcavqpjaonaudz"),
	// 			TargetPortalPort: to.Ptr[int32](1),
	// 		},
	// 		VolumeID: to.Ptr("tbwshhvrbqyseonkeztlbzhrckd"),
	// 	},
	// }
}
