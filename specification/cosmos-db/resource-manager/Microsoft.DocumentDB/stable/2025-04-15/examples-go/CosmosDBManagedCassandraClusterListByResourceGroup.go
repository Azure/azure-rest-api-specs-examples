package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/011ecc5633300a5eefe43dde748f269d39e96458/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBManagedCassandraClusterListByResourceGroup.json
func ExampleCassandraClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraClustersClient().NewListByResourceGroupPager("cassandra-prod-rg", nil)
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
		// page.ListClusters = armcosmos.ListClusters{
		// 	Value: []*armcosmos.ClusterResource{
		// 		{
		// 			Name: to.Ptr("cassandra-prod"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.ClusterResourceProperties{
		// 				AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
		// 				CassandraVersion: to.Ptr("3.11"),
		// 				ClientCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				ClusterNameOverride: to.Ptr("ClusterNameIllegalForAzureResource"),
		// 				DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
		// 				ExternalGossipCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				ExternalSeedNodes: []*armcosmos.SeedNode{
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.4"),
		// 				}},
		// 				GossipCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				HoursBetweenBackups: to.Ptr[int32](24),
		// 				ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
		// 				SeedNodes: []*armcosmos.SeedNode{
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.4"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.4"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
