package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBManagedCassandraClusterPatch.json
func ExampleCassandraClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewCassandraClustersClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "cassandra-prod-rg", "cassandra-prod", armcosmos.ClusterResource{
		Tags: map[string]*string{
			"owner": to.Ptr("mike"),
		},
		Properties: &armcosmos.ClusterResourceProperties{
			AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodNone),
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
			HoursBetweenBackups: to.Ptr[int32](12),
		},
	}, nil)
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
