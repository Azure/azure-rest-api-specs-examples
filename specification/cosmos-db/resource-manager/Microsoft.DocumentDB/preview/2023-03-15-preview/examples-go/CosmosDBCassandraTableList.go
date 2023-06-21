package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBCassandraTableList.json
func ExampleCassandraResourcesClient_NewListCassandraTablesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraResourcesClient().NewListCassandraTablesPager("rgName", "ddb1", "keyspaceName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CassandraTableListResult = armcosmos.CassandraTableListResult{
		// 	Value: []*armcosmos.CassandraTableGetResults{
		// 		{
		// 			Name: to.Ptr("tableName"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/cassandraKeyspaces/cassandraTables"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/cassandraKeyspaces/keyspaceName/cassandraTables/tableName"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.CassandraTableGetProperties{
		// 				Resource: &armcosmos.CassandraTableGetPropertiesResource{
		// 					Schema: &armcosmos.CassandraSchema{
		// 						ClusterKeys: []*armcosmos.ClusterKey{
		// 							{
		// 								Name: to.Ptr("columnA"),
		// 								OrderBy: to.Ptr("Asc"),
		// 						}},
		// 						Columns: []*armcosmos.Column{
		// 							{
		// 								Name: to.Ptr("columnA"),
		// 								Type: to.Ptr("Ascii"),
		// 						}},
		// 						PartitionKeys: []*armcosmos.CassandraPartitionKey{
		// 							{
		// 								Name: to.Ptr("columnA"),
		// 						}},
		// 					},
		// 					AnalyticalStorageTTL: to.Ptr[int32](500),
		// 					DefaultTTL: to.Ptr[int32](100),
		// 					ID: to.Ptr("tableName"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
