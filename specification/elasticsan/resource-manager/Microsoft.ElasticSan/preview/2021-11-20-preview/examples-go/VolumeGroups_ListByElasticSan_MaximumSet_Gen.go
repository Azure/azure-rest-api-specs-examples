package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1da7cbab8d4f554484dedb676ba7bdbdf6cdf78/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/VolumeGroups_ListByElasticSan_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_NewListByElasticSanPager_volumeGroupsListByElasticSanMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeGroupsClient().NewListByElasticSanPager("rgelasticsan", "ti7q-k952-1qB3J_5", nil)
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
		// page.VolumeGroupList = armelasticsan.VolumeGroupList{
		// 	Value: []*armelasticsan.VolumeGroup{
		// 		{
		// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 			Type: to.Ptr("aaaaa"),
		// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Tags: map[string]*string{
		// 				"key5933": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 			Properties: &armelasticsan.VolumeGroupProperties{
		// 				Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 				NetworkACLs: &armelasticsan.NetworkRuleSet{
		// 					VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
		// 						{
		// 							Action: to.Ptr("Allow"),
		// 							VirtualNetworkResourceID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 							State: to.Ptr(armelasticsan.StateProvisioning),
		// 					}},
		// 				},
		// 				ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
		// 				ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 			},
		// 			SystemData: &armelasticsan.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-08T10:39:37.620Z"); return t}()),
		// 				CreatedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-08T10:39:37.620Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
