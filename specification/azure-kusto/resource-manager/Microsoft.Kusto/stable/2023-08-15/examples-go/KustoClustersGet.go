package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClustersGet.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "kustorptest", "kustoCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armkusto.Cluster{
	// 	Name: to.Ptr("kustoCluster"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("abcd123"),
	// 	Identity: &armkusto.Identity{
	// 		Type: to.Ptr(armkusto.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("faabad1f-4876-463c-af9d-6ba2d2d2394c"),
	// 		TenantID: to.Ptr("b932977f-6277-4ab7-a2cd-5bd21f07aaf4"),
	// 		UserAssignedIdentities: map[string]*armkusto.ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 		},
	// 	},
	// 	Properties: &armkusto.ClusterProperties{
	// 		AllowedFqdnList: []*string{
	// 			to.Ptr("my-stroage.blob.core.windows.net")},
	// 			AllowedIPRangeList: []*string{
	// 			},
	// 			EnableAutoStop: to.Ptr(true),
	// 			EnableDiskEncryption: to.Ptr(false),
	// 			EnablePurge: to.Ptr(false),
	// 			EnableStreamingIngest: to.Ptr(true),
	// 			EngineType: to.Ptr(armkusto.EngineTypeV3),
	// 			KeyVaultProperties: &armkusto.KeyVaultProperties{
	// 				KeyName: to.Ptr("keyName"),
	// 				KeyVaultURI: to.Ptr("https://dummy.keyvault.com"),
	// 				KeyVersion: to.Ptr("keyVersion"),
	// 			},
	// 			MigrationCluster: &armkusto.MigrationClusterProperties{
	// 				DataIngestionURI: to.Ptr("https://ingest-kustocluster2.westus.kusto.windows.net"),
	// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
	// 				Role: to.Ptr(armkusto.MigrationClusterRoleDestination),
	// 				URI: to.Ptr("https://kustocluster2.westus.kusto.windows.net"),
	// 			},
	// 			PrivateEndpointConnections: []*armkusto.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("privateEndpointTest"),
	// 					Type: to.Ptr("Microsoft.Kusto/Clusters/PrivateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/privateEndpointConnections/privateEndpointTest"),
	// 					Properties: &armkusto.PrivateEndpointConnectionProperties{
	// 						GroupID: to.Ptr("cluster"),
	// 						PrivateEndpoint: &armkusto.PrivateEndpointProperty{
	// 							ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armkusto.PrivateLinkServiceConnectionStateProperty{
	// 							Description: to.Ptr("Auto-approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr("Approved"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 					},
	// 			}},
	// 			ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 			PublicIPType: to.Ptr(armkusto.PublicIPTypeIPv4),
	// 			PublicNetworkAccess: to.Ptr(armkusto.PublicNetworkAccessEnabled),
	// 			RestrictOutboundNetworkAccess: to.Ptr(armkusto.ClusterNetworkAccessFlagEnabled),
	// 		},
	// 		SKU: &armkusto.AzureSKU{
	// 			Name: to.Ptr(armkusto.AzureSKUNameStandardL16AsV3),
	// 			Capacity: to.Ptr[int32](2),
	// 			Tier: to.Ptr(armkusto.AzureSKUTierStandard),
	// 		},
	// 		SystemData: &armkusto.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 		},
	// 	}
}
