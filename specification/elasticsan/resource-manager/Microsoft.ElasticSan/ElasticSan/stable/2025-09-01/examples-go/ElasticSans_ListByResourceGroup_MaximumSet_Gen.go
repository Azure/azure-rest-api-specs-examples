package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: 2025-09-01/ElasticSans_ListByResourceGroup_MaximumSet_Gen.json
func ExampleElasticSansClient_NewListByResourceGroupPager_elasticSansListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticSansClient().NewListByResourceGroupPager("resourcegroupname", nil)
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
		// page = armelasticsan.ElasticSansClientListByResourceGroupResponse{
		// 	List: armelasticsan.List{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionid/resourceGroups/resourcegroupname/providers/Microsoft.ElasticSan/elasticSans?api-version=2024-07-01-preview&%24skiptoken=ghi789jkl012"),
		// 		Value: []*armelasticsan.ElasticSan{
		// 			{
		// 				Name: to.Ptr("vfoatmakv"),
		// 				Type: to.Ptr("Microsoft.ElasticSan/ElasticSans"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}"),
		// 				Location: to.Ptr("France Central"),
		// 				Properties: &armelasticsan.Properties{
		// 					AutoScaleProperties: &armelasticsan.AutoScaleProperties{
		// 						ScaleUpProperties: &armelasticsan.ScaleUpProperties{
		// 							AutoScalePolicyEnforcement: to.Ptr(armelasticsan.AutoScalePolicyEnforcementNone),
		// 							CapacityUnitScaleUpLimitTiB: to.Ptr[int64](17),
		// 							IncreaseCapacityUnitByTiB: to.Ptr[int64](4),
		// 							UnusedSizeTiB: to.Ptr[int64](24),
		// 						},
		// 					},
		// 					AvailabilityZones: []*string{
		// 						to.Ptr("1"),
		// 					},
		// 					BaseSizeTiB: to.Ptr[int64](15),
		// 					ExtendedCapacitySizeTiB: to.Ptr[int64](6),
		// 					PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
		// 						{
		// 							Name: to.Ptr("{privateEndpointConnectionName}"),
		// 							Type: to.Ptr("Microsoft.ElasticSan/elasticSans/privateEndpointConnections"),
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateEndpointConnections/{privateEndpointConnectionName}"),
		// 							Properties: &armelasticsan.PrivateEndpointConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"),
		// 								},
		// 								PrivateEndpoint: &armelasticsan.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
		// 									Description: to.Ptr("Auto-Approved"),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
		// 								},
		// 								ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 							},
		// 							SystemData: &armelasticsan.SystemData{
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
		// 								CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
		// 								CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("bcclmbseed"),
		// 								LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armelasticsan.PublicNetworkAccessEnabled),
		// 					SKU: &armelasticsan.SKU{
		// 						Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
		// 						Tier: to.Ptr(armelasticsan.SKUTierPremium),
		// 					},
		// 					TotalIops: to.Ptr[int64](22),
		// 					TotalMBps: to.Ptr[int64](4),
		// 					TotalSizeTiB: to.Ptr[int64](27),
		// 					TotalVolumeSizeGiB: to.Ptr[int64](15),
		// 					VolumeGroupCount: to.Ptr[int64](24),
		// 				},
		// 				SystemData: &armelasticsan.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
		// 					CreatedBy: to.Ptr("otfifnrahdshqombvtg"),
		// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("jnaxavnlhrboshtidtib"),
		// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key5002": to.Ptr("lhag"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
