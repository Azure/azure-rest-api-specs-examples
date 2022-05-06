Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv0.5.0/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.

```go
package armcosmos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBCassandraTableCreateUpdate.json
func ExampleCassandraResourcesClient_BeginCreateUpdateCassandraTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcosmos.NewCassandraResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateUpdateCassandraTable(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<keyspace-name>",
		"<table-name>",
		armcosmos.CassandraTableCreateUpdateParameters{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armcosmos.CassandraTableCreateUpdateProperties{
				Options: &armcosmos.CreateUpdateOptions{},
				Resource: &armcosmos.CassandraTableResource{
					Schema: &armcosmos.CassandraSchema{
						ClusterKeys: []*armcosmos.ClusterKey{
							{
								Name:    to.Ptr("<name>"),
								OrderBy: to.Ptr("<order-by>"),
							}},
						Columns: []*armcosmos.Column{
							{
								Name: to.Ptr("<name>"),
								Type: to.Ptr("<type>"),
							}},
						PartitionKeys: []*armcosmos.CassandraPartitionKey{
							{
								Name: to.Ptr("<name>"),
							}},
					},
					AnalyticalStorageTTL: to.Ptr[int32](500),
					DefaultTTL:           to.Ptr[int32](100),
					ID:                   to.Ptr("<id>"),
				},
			},
		},
		&armcosmos.CassandraResourcesClientBeginCreateUpdateCassandraTableOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
