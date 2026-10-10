package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan/v2"
)

// Generated from example definition: 2026-05-01-preview/Volumes_ListByVolumeGroup_MaximumSet_Gen.json
func ExampleVolumesClient_NewListByVolumeGroupPager_volumesListByVolumeGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumesClient().NewListByVolumeGroupPager("resourcegroupname", "elasticsanname", "volumegroupname", &armelasticsan.VolumesClientListByVolumeGroupOptions{
		XMSAccessSoftDeletedResources: to.Ptr(armelasticsan.XMSAccessSoftDeletedResourcesTrue)})
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
		// page = armelasticsan.VolumesClientListByVolumeGroupResponse{
		// 	VolumeList: armelasticsan.VolumeList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionid/resourceGroups/resourcegroupname/providers/Microsoft.ElasticSan/elasticSans/elasticsanname/volumegroups/volumegroupname/volumes?api-version=2026-05-01-preview&%24skiptoken=jkl789mno012"),
		// 		Value: []*armelasticsan.Volume{
		// 			{
		// 				Name: to.Ptr("hxm"),
		// 				Type: to.Ptr("sxsbsdhngvbcpxxcuvyt"),
		// 				ID: to.Ptr("sipsx"),
		// 				Properties: &armelasticsan.VolumeProperties{
		// 					CreationData: &armelasticsan.SourceCreationData{
		// 						CreateSource: to.Ptr(armelasticsan.VolumeCreateOptionNone),
		// 						SourceID: to.Ptr("mdonegivjquite"),
		// 					},
		// 					ManagedBy: []*armelasticsan.ManagedByResources{
		// 						{
		// 							ClientID: to.Ptr("pclpkrpkpmvcsegcubrakcoodrubo"),
		// 							Version: to.Ptr[int32](1),
		// 							ResourceIDs: []*string{
		// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.SomeProvider/someResource/myResource"),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
		// 					SizeGiB: to.Ptr[int64](23),
		// 					StorageTarget: &armelasticsan.IscsiTargetInfo{
		// 						ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
		// 						Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
		// 						TargetIqn: to.Ptr("us"),
		// 						TargetPortalHostname: to.Ptr("oqtcavqpjaonaudz"),
		// 						TargetPortalPort: to.Ptr[int32](1),
		// 					},
		// 					VolumeID: to.Ptr("tbwshhvrbqyseonkeztlbzhrckd"),
		// 				},
		// 				SystemData: &armelasticsan.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2024, time.April, 29, 14, 22, 13, 546000000, time.UTC)),
		// 					CreatedBy: to.Ptr("bpuxtfzqwdhifevjtucoc"),
		// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.April, 29, 14, 22, 13, 547000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("ourjjlolgugpxnkbiegumkicksibep"),
		// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
