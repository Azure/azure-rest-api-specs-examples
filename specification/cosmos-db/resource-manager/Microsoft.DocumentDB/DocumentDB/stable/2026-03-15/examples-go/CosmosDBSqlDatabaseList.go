package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBSqlDatabaseList.json
func ExampleSQLResourcesClient_NewListSQLDatabasesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLResourcesClient().NewListSQLDatabasesPager("rgName", "ddb1", nil)
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
		// page = armcosmos.SQLResourcesClientListSQLDatabasesResponse{
		// 	SQLDatabaseListResult: armcosmos.SQLDatabaseListResult{
		// 		Value: []*armcosmos.SQLDatabaseGetResults{
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName"),
		// 				Name: to.Ptr("databaseName"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases"),
		// 				Location: to.Ptr("West US"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armcosmos.SQLDatabaseGetProperties{
		// 					Resource: &armcosmos.SQLDatabaseGetPropertiesResource{
		// 						ID: to.Ptr("databaseName"),
		// 						Rid: to.Ptr("CqNBAA=="),
		// 						Ts: to.Ptr[float32](1449602962),
		// 						Etag: to.Ptr("\"00000a00-0000-0000-0000-56672f920000\""),
		// 						Colls: to.Ptr("colls/"),
		// 						Users: to.Ptr("users/"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
