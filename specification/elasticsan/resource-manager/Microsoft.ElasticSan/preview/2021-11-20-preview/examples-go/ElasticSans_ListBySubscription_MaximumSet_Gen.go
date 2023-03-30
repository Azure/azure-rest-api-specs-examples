package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1da7cbab8d4f554484dedb676ba7bdbdf6cdf78/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/ElasticSans_ListBySubscription_MaximumSet_Gen.json
func ExampleElasticSansClient_NewListBySubscriptionPager_elasticSansListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticSansClient().NewListBySubscriptionPager(nil)
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
		// page.List = armelasticsan.List{
		// 	Value: []*armelasticsan.ElasticSan{
		// 		{
		// 			Name: to.Ptr("aaaaaaaaaaa"),
		// 			Type: to.Ptr("aaaaaa"),
		// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Tags: map[string]*string{
		// 				"key896": to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 			},
		// 			Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Properties: &armelasticsan.Properties{
		// 				AvailabilityZones: []*string{
		// 					to.Ptr("aaaaaaaaaaaaaaaaa")},
		// 					BaseSizeTiB: to.Ptr[int64](26),
		// 					ExtendedCapacitySizeTiB: to.Ptr[int64](7),
		// 					ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 					SKU: &armelasticsan.SKU{
		// 						Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
		// 						Tier: to.Ptr(armelasticsan.SKUTierPremium),
		// 					},
		// 					TotalIops: to.Ptr[int64](13),
		// 					TotalMBps: to.Ptr[int64](16),
		// 					TotalSizeTiB: to.Ptr[int64](29),
		// 					TotalVolumeSizeGiB: to.Ptr[int64](21),
		// 					VolumeGroupCount: to.Ptr[int64](24),
		// 				},
		// 				SystemData: &armelasticsan.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-08T10:39:37.620Z"); return t}()),
		// 					CreatedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-08T10:39:37.620Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}
