package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBCassandraTableCreateUpdate.json
func ExampleCassandraResourcesClient_BeginCreateUpdateCassandraTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraResourcesClient().BeginCreateUpdateCassandraTable(ctx, "rg1", "ddb1", "keyspaceName", "tableName", armcosmos.CassandraTableCreateUpdateParameters{
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
				DefaultTTL: to.Ptr[int32](100),
				ID:         to.Ptr("tableName"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CassandraTableGetResults = armcosmos.CassandraTableGetResults{
	// 	Name: to.Ptr("tableName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/cassandraKeyspaces/cassandraTables"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/cassandraKeyspaces/keyspaceName/cassandraTables/tableName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.CassandraTableGetProperties{
	// 		Resource: &armcosmos.CassandraTableGetPropertiesResource{
	// 			Schema: &armcosmos.CassandraSchema{
	// 				ClusterKeys: []*armcosmos.ClusterKey{
	// 					{
	// 						Name: to.Ptr("columnA"),
	// 						OrderBy: to.Ptr("Asc"),
	// 				}},
	// 				Columns: []*armcosmos.Column{
	// 					{
	// 						Name: to.Ptr("columnA"),
	// 						Type: to.Ptr("Ascii"),
	// 				}},
	// 				PartitionKeys: []*armcosmos.CassandraPartitionKey{
	// 					{
	// 						Name: to.Ptr("columnA"),
	// 				}},
	// 			},
	// 			DefaultTTL: to.Ptr[int32](100),
	// 			ID: to.Ptr("tableName"),
	// 		},
	// 	},
	// }
}
