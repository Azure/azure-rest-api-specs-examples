package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBMongoDBDatabaseRetrieveThroughputDistribution.json
func ExampleMongoDBResourcesClient_BeginMongoDBDatabaseRetrieveThroughputDistribution() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewMongoDBResourcesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginMongoDBDatabaseRetrieveThroughputDistribution(ctx, "rg1", "ddb1", "databaseName", armcosmos.RetrieveThroughputParameters{
		Properties: &armcosmos.RetrieveThroughputProperties{
			Resource: &armcosmos.RetrieveThroughputPropertiesResource{
				PhysicalPartitionIDs: []*armcosmos.PhysicalPartitionID{
					{
						ID: to.Ptr("0"),
					},
					{
						ID: to.Ptr("1"),
					}},
			},
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
