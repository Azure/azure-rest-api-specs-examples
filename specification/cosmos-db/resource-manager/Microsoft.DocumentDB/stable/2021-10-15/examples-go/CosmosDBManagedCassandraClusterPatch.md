```go
package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraClusterPatch.json
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
	poller, err := client.BeginUpdate(ctx,
		"cassandra-prod-rg",
		"cassandra-prod",
		armcosmos.ClusterResource{
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv1.0.0/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.
