package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/ElasticSans_ListBySubscription_MaximumSet_Gen.json
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
		// 			Name: to.Ptr("dnwirxrikyzsoqmf"),
		// 			Type: to.Ptr("uocreclqzfafzwwbzozgryxurwhf"),
		// 			ID: to.Ptr("wthgilvis"),
		// 			SystemData: &armelasticsan.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-24T09:15:59.688Z"); return t}()),
		// 				CreatedBy: to.Ptr("bcgpoousvimjrhlgweigyeazr"),
		// 				CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-24T09:15:59.688Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("jflskmaiembmcpdrurolsi"),
		// 				LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eupkimcjwiphitotibv"),
		// 			Tags: map[string]*string{
		// 				"key9706": to.Ptr("haitqqakcntcpalkzqmjmcnifnhd"),
		// 			},
		// 			Properties: &armelasticsan.Properties{
		// 				AutoScaleProperties: &armelasticsan.AutoScaleProperties{
		// 					ScaleUpProperties: &armelasticsan.ScaleUpProperties{
		// 						AutoScalePolicyEnforcement: to.Ptr(armelasticsan.AutoScalePolicyEnforcementNone),
		// 						CapacityUnitScaleUpLimitTiB: to.Ptr[int64](17),
		// 						IncreaseCapacityUnitByTiB: to.Ptr[int64](4),
		// 						UnusedSizeTiB: to.Ptr[int64](24),
		// 					},
		// 				},
		// 				AvailabilityZones: []*string{
		// 					to.Ptr("xoz")},
		// 					BaseSizeTiB: to.Ptr[int64](1),
		// 					ExtendedCapacitySizeTiB: to.Ptr[int64](3),
		// 					PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
		// 						{
		// 							Name: to.Ptr("jgjelkmzzuu"),
		// 							Type: to.Ptr("emgnetaviizobnwm"),
		// 							ID: to.Ptr("xnaxhlih"),
		// 							SystemData: &armelasticsan.SystemData{
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-24T09:15:59.688Z"); return t}()),
		// 								CreatedBy: to.Ptr("bcgpoousvimjrhlgweigyeazr"),
		// 								CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-24T09:15:59.688Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("jflskmaiembmcpdrurolsi"),
		// 								LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 							},
		// 							Properties: &armelasticsan.PrivateEndpointConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("seuxenjsfpavichmai")},
		// 									PrivateEndpoint: &armelasticsan.PrivateEndpoint{
		// 										ID: to.Ptr("wnlsfvatungpitrlcbhgdtlpl"),
		// 									},
		// 									PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
		// 										Description: to.Ptr("mrcibusihsmwcbmsrukbmrrjdiy"),
		// 										ActionsRequired: to.Ptr("eyhujyohsisgywrozmunjmrtao"),
		// 										Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
		// 									},
		// 									ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 								},
		// 						}},
		// 						ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 						PublicNetworkAccess: to.Ptr(armelasticsan.PublicNetworkAccessEnabled),
		// 						SKU: &armelasticsan.SKU{
		// 							Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
		// 							Tier: to.Ptr(armelasticsan.SKUTierPremium),
		// 						},
		// 						TotalIops: to.Ptr[int64](29),
		// 						TotalMBps: to.Ptr[int64](6),
		// 						TotalSizeTiB: to.Ptr[int64](5),
		// 						TotalVolumeSizeGiB: to.Ptr[int64](12),
		// 						VolumeGroupCount: to.Ptr[int64](25),
		// 					},
		// 			}},
		// 		}
	}
}
