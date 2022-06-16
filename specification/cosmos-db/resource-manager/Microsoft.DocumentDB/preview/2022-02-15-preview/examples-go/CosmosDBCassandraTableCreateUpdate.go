package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBCassandraTableCreateUpdate.json
func ExampleCassandraResourcesClient_BeginCreateUpdateCassandraTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewCassandraResourcesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdateCassandraTable(ctx,
		"rg1",
		"ddb1",
		"keyspaceName",
		"tableName",
		armcosmos.CassandraTableCreateUpdateParameters{
			Location: to.Ptr("West US"),
			Tags:     map[string]*string{},
			Properties: &armcosmos.CassandraTableCreateUpdateProperties{
				Options: &armcosmos.CreateUpdateOptions{},
				Resource: &armcosmos.CassandraTableResource{
					Schema: &armcosmos.CassandraSchema{
						ClusterKeys: []*armcosmos.ClusterKey{
							{
								Name:    to.Ptr("columnA"),
								OrderBy: to.Ptr("Asc"),
							}},
						Columns: []*armcosmos.Column{
							{
								Name: to.Ptr("columnA"),
								Type: to.Ptr("Ascii"),
							}},
						PartitionKeys: []*armcosmos.CassandraPartitionKey{
							{
								Name: to.Ptr("columnA"),
							}},
					},
					AnalyticalStorageTTL: to.Ptr[int32](500),
					DefaultTTL:           to.Ptr[int32](100),
					ID:                   to.Ptr("tableName"),
				},
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
