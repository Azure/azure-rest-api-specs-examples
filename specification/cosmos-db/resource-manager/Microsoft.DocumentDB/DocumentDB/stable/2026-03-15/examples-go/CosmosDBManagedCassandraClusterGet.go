package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBManagedCassandraClusterGet.json
func ExampleCassandraClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCassandraClustersClient().Get(ctx, "cassandra-prod-rg", "cassandra-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.CassandraClustersClientGetResponse{
	// 	ClusterResource: armcosmos.ClusterResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod"),
	// 		Name: to.Ptr("cassandra-prod"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcosmos.ClusterResourceProperties{
	// 			ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
	// 			DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
	// 			CassandraVersion: to.Ptr("3.11"),
	// 			HoursBetweenBackups: to.Ptr[int32](24),
	// 			AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
	// 			ExternalSeedNodes: []*armcosmos.SeedNode{
	// 				{
	// 					IPAddress: to.Ptr("10.52.221.2"),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.52.221.3"),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.52.221.4"),
	// 				},
	// 			},
	// 			ClusterNameOverride: to.Ptr("ClusterNameIllegalForAzureResource"),
	// 			SeedNodes: []*armcosmos.SeedNode{
	// 				{
	// 					IPAddress: to.Ptr("10.52.221.2"),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.52.221.3"),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.52.221.4"),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("192.168.12.2"),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("192.168.12.3"),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("192.168.12.4"),
	// 				},
	// 			},
	// 			ClientCertificates: []*armcosmos.Certificate{
	// 				{
	// 					Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 				},
	// 			},
	// 			ExternalGossipCertificates: []*armcosmos.Certificate{
	// 				{
	// 					Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 				},
	// 			},
	// 			GossipCertificates: []*armcosmos.Certificate{
	// 				{
	// 					Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
