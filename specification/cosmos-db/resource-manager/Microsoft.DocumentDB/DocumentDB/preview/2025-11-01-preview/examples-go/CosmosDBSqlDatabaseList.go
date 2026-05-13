package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBSqlDatabaseList.json
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
		// 				Name: to.Ptr("databaseName"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armcosmos.SQLDatabaseGetProperties{
		// 					Resource: &armcosmos.SQLDatabaseGetPropertiesResource{
		// 						Colls: to.Ptr("colls/"),
		// 						Etag: to.Ptr("\"00000a00-0000-0000-0000-56672f920000\""),
		// 						Rid: to.Ptr("CqNBAA=="),
		// 						Ts: to.Ptr[float32](1449602962),
		// 						Users: to.Ptr("users/"),
		// 						ID: to.Ptr("databaseName"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
