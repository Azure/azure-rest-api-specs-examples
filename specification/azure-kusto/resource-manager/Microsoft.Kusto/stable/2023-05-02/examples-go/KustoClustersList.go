package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoClustersList.json
func ExampleClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListPager(nil)
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
		// page.ClusterListResult = armkusto.ClusterListResult{
		// 	Value: []*armkusto.Cluster{
		// 		{
		// 			Name: to.Ptr("kustoCluster"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("abcd123"),
		// 			Properties: &armkusto.ClusterProperties{
		// 				AllowedFqdnList: []*string{
		// 					to.Ptr("my-stroage.blob.core.windows.net")},
		// 					AllowedIPRangeList: []*string{
		// 					},
		// 					EnableAutoStop: to.Ptr(true),
		// 					EnableDiskEncryption: to.Ptr(false),
		// 					EnableStreamingIngest: to.Ptr(true),
		// 					EngineType: to.Ptr(armkusto.EngineTypeV3),
		// 					KeyVaultProperties: &armkusto.KeyVaultProperties{
		// 						KeyName: to.Ptr("keyName"),
		// 						KeyVaultURI: to.Ptr("https://dummy.keyvault.com"),
		// 						KeyVersion: to.Ptr("keyVersion"),
		// 					},
		// 					MigrationCluster: &armkusto.MigrationClusterProperties{
		// 						DataIngestionURI: to.Ptr("https://ingest-kustocluster2.westus.kusto.windows.net"),
		// 						ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
		// 						Role: to.Ptr(armkusto.MigrationClusterRoleDestination),
		// 						URI: to.Ptr("https://kustocluster2.westus.kusto.windows.net"),
		// 					},
		// 					PrivateEndpointConnections: []*armkusto.PrivateEndpointConnection{
		// 						{
		// 							Name: to.Ptr("privateEndpointTest"),
		// 							Type: to.Ptr("Microsoft.Kusto/Clusters/PrivateEndpointConnections"),
		// 							ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/privateEndpointConnections/privateEndpointTest"),
		// 							Properties: &armkusto.PrivateEndpointConnectionProperties{
		// 								GroupID: to.Ptr("cluster"),
		// 								PrivateEndpoint: &armkusto.PrivateEndpointProperty{
		// 									ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armkusto.PrivateLinkServiceConnectionStateProperty{
		// 									Description: to.Ptr("Auto-approved"),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Status: to.Ptr("Approved"),
		// 								},
		// 								ProvisioningState: to.Ptr("Succeeded"),
		// 							},
		// 					}},
		// 					ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 					PublicIPType: to.Ptr(armkusto.PublicIPTypeDualStack),
		// 					PublicNetworkAccess: to.Ptr(armkusto.PublicNetworkAccessEnabled),
		// 					RestrictOutboundNetworkAccess: to.Ptr(armkusto.ClusterNetworkAccessFlagEnabled),
		// 				},
		// 				SKU: &armkusto.AzureSKU{
		// 					Name: to.Ptr(armkusto.AzureSKUNameStandardL16AsV3),
		// 					Capacity: to.Ptr[int32](2),
		// 					Tier: to.Ptr(armkusto.AzureSKUTierStandard),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("KustoClusterRPTest3"),
		// 				Type: to.Ptr("Microsoft.Kusto/Clusters"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/KustoClusterRPTest3"),
		// 				Location: to.Ptr("westus"),
		// 				Etag: to.Ptr("abcd123"),
		// 				Properties: &armkusto.ClusterProperties{
		// 					AllowedIPRangeList: []*string{
		// 						to.Ptr("0.0.0.0/0")},
		// 						EnableDiskEncryption: to.Ptr(true),
		// 						EnableStreamingIngest: to.Ptr(true),
		// 						EngineType: to.Ptr(armkusto.EngineTypeV2),
		// 						PrivateEndpointConnections: []*armkusto.PrivateEndpointConnection{
		// 						},
		// 						ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 						PublicIPType: to.Ptr(armkusto.PublicIPTypeIPv4),
		// 						PublicNetworkAccess: to.Ptr(armkusto.PublicNetworkAccessEnabled),
		// 						RestrictOutboundNetworkAccess: to.Ptr(armkusto.ClusterNetworkAccessFlagDisabled),
		// 					},
		// 					SKU: &armkusto.AzureSKU{
		// 						Name: to.Ptr(armkusto.AzureSKUNameStandardL16AsV3),
		// 						Capacity: to.Ptr[int32](2),
		// 						Tier: to.Ptr(armkusto.AzureSKUTierStandard),
		// 					},
		// 			}},
		// 		}
	}
}
