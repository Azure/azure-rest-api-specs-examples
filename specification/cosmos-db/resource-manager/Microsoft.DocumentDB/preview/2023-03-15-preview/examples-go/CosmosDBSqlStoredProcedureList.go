package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBSqlStoredProcedureList.json
func ExampleSQLResourcesClient_NewListSQLStoredProceduresPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLResourcesClient().NewListSQLStoredProceduresPager("rgName", "ddb1", "databaseName", "containerName", nil)
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
		// page.SQLStoredProcedureListResult = armcosmos.SQLStoredProcedureListResult{
		// 	Value: []*armcosmos.SQLStoredProcedureGetResults{
		// 		{
		// 			Name: to.Ptr("testctn"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers/sqlStoredProcedures"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/sqlContainers/containerName/sqlStoredProcedures/storedProcedureName"),
		// 			Properties: &armcosmos.SQLStoredProcedureGetProperties{
		// 				Resource: &armcosmos.SQLStoredProcedureGetPropertiesResource{
		// 					Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
		// 					Rid: to.Ptr("PD5DALigDgw="),
		// 					Ts: to.Ptr[float32](1459200611),
		// 					Body: to.Ptr("body"),
		// 					ID: to.Ptr("testctn"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
