package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterCreate.json
func ExampleCassandraClustersClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewCassandraClustersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdate(ctx,
		"cassandra-prod-rg",
		"cassandra-prod",
		armcosmos.ClusterResource{
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
