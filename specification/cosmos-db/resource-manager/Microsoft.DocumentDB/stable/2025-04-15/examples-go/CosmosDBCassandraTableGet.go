package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/011ecc5633300a5eefe43dde748f269d39e96458/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBCassandraTableGet.json
func ExampleCassandraResourcesClient_GetCassandraTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCassandraResourcesClient().GetCassandraTable(ctx, "rg1", "ddb1", "keyspaceName", "tableName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 		},
	// 	},
	// }
}
