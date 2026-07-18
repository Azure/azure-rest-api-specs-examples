package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBCassandraTableCreateUpdate.json
func ExampleCassandraResourcesClient_BeginCreateUpdateCassandraTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraResourcesClient().BeginCreateUpdateCassandraTable(ctx, "rg1", "ddb1", "keyspaceName", "tableName", armcosmos.CassandraTableCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.CassandraTableCreateUpdateProperties{
			Resource: &armcosmos.CassandraTableResource{
				ID:         to.Ptr("tableName"),
				DefaultTTL: to.Ptr[int32](100),
				Schema: &armcosmos.CassandraSchema{
					Columns: []*armcosmos.Column{
						{
							Name: to.Ptr("columnA"),
							Type: to.Ptr("Ascii"),
						},
					},
					PartitionKeys: []*armcosmos.CassandraPartitionKey{
						{
							Name: to.Ptr("columnA"),
						},
					},
					ClusterKeys: []*armcosmos.ClusterKey{
						{
							Name:    to.Ptr("columnA"),
							OrderBy: to.Ptr("Asc"),
						},
					},
				},
			},
			Options: &armcosmos.CreateUpdateOptions{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.CassandraResourcesClientCreateUpdateCassandraTableResponse{
	// 	CassandraTableGetResults: armcosmos.CassandraTableGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/cassandraKeyspaces/keyspaceName/cassandraTables/tableName"),
	// 		Name: to.Ptr("tableName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/cassandraKeyspaces/cassandraTables"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcosmos.CassandraTableGetProperties{
	// 			Resource: &armcosmos.CassandraTableGetPropertiesResource{
	// 				ID: to.Ptr("tableName"),
	// 				DefaultTTL: to.Ptr[int32](100),
	// 				Schema: &armcosmos.CassandraSchema{
	// 					Columns: []*armcosmos.Column{
	// 						{
	// 							Name: to.Ptr("columnA"),
	// 							Type: to.Ptr("Ascii"),
	// 						},
	// 					},
	// 					PartitionKeys: []*armcosmos.CassandraPartitionKey{
	// 						{
	// 							Name: to.Ptr("columnA"),
	// 						},
	// 					},
	// 					ClusterKeys: []*armcosmos.ClusterKey{
	// 						{
	// 							Name: to.Ptr("columnA"),
	// 							OrderBy: to.Ptr("Asc"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
