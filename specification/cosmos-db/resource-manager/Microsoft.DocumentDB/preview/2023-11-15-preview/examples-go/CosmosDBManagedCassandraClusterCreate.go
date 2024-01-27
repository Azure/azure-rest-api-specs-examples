package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6f8faf5da91b5b9af5f3512fe609e22e99383d41/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/CosmosDBManagedCassandraClusterCreate.json
func ExampleCassandraClustersClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginCreateUpdate(ctx, "cassandra-prod-rg", "cassandra-prod", armcosmos.ClusterResource{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ClusterResourceProperties{
			AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
			CassandraVersion:     to.Ptr("3.11"),
			ClientCertificates: []*armcosmos.Certificate{
				{
					Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
				}},
			ClusterNameOverride:         to.Ptr("ClusterNameIllegalForAzureResource"),
			DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
			ExternalGossipCertificates: []*armcosmos.Certificate{
				{
					Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
				}},
			ExternalSeedNodes: []*armcosmos.SeedNode{
				{
					IPAddress: to.Ptr("10.52.221.2"),
				},
				{
					IPAddress: to.Ptr("10.52.221.3"),
				},
				{
					IPAddress: to.Ptr("10.52.221.4"),
				}},
			HoursBetweenBackups:           to.Ptr[int32](24),
			InitialCassandraAdminPassword: to.Ptr("mypassword"),
		},
	}, nil)
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
	// res.ClusterResource = armcosmos.ClusterResource{
	// 	Name: to.Ptr("cassandra-prod"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ClusterResourceProperties{
	// 		AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
	// 		CassandraVersion: to.Ptr("3.11"),
	// 		ClientCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		ClusterNameOverride: to.Ptr("ClusterNameIllegalForAzureResource"),
	// 		DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
	// 		ExternalGossipCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		ExternalSeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.4"),
	// 		}},
	// 		GossipCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		HoursBetweenBackups: to.Ptr[int32](24),
	// 		ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
	// 		SeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.4"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.4"),
	// 		}},
	// 	},
	// }
}
