package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/ElasticSans_Update_MinimumSet_Gen.json
func ExampleElasticSansClient_BeginUpdate_elasticSansUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewElasticSansClient().BeginUpdate(ctx, "resourcegroupname", "elasticsanname", armelasticsan.Update{}, nil)
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
	// res.ElasticSan = armelasticsan.ElasticSan{
	// 	Name: to.Ptr("vfoatmakv"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/ElasticSans"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 		CreatedBy: to.Ptr("otfifnrahdshqombvtg"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("jnaxavnlhrboshtidtib"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("France Central"),
	// 	Tags: map[string]*string{
	// 		"key5002": to.Ptr("lhag"),
	// 	},
	// 	Properties: &armelasticsan.Properties{
	// 		AvailabilityZones: []*string{
	// 			to.Ptr("1")},
	// 			BaseSizeTiB: to.Ptr[int64](15),
	// 			ExtendedCapacitySizeTiB: to.Ptr[int64](6),
	// 			PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("{privateEndpointConnectionName}"),
	// 					Type: to.Ptr("Microsoft.ElasticSan/elasticSans/privateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 					SystemData: &armelasticsan.SystemData{
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 						CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 						CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("bcclmbseed"),
	// 						LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 					},
	// 					Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 						GroupIDs: []*string{
	// 							to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}")},
	// 							PrivateEndpoint: &armelasticsan.PrivateEndpoint{
	// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"),
	// 							},
	// 							PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
	// 								Description: to.Ptr("Auto-Approved"),
	// 								ActionsRequired: to.Ptr("None"),
	// 								Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
	// 							},
	// 							ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 						},
	// 				}},
	// 				ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 				PublicNetworkAccess: to.Ptr(armelasticsan.PublicNetworkAccessEnabled),
	// 				SKU: &armelasticsan.SKU{
	// 					Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
	// 					Tier: to.Ptr(armelasticsan.SKUTierPremium),
	// 				},
	// 				TotalIops: to.Ptr[int64](22),
	// 				TotalMBps: to.Ptr[int64](4),
	// 				TotalSizeTiB: to.Ptr[int64](27),
	// 				TotalVolumeSizeGiB: to.Ptr[int64](15),
	// 				VolumeGroupCount: to.Ptr[int64](24),
	// 			},
	// 		}
}
