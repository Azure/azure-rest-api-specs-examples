package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6f8faf5da91b5b9af5f3512fe609e22e99383d41/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/CosmosDBSqlDatabaseGet.json
func ExampleSQLResourcesClient_GetSQLDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLResourcesClient().GetSQLDatabase(ctx, "rg1", "ddb1", "databaseName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLDatabaseGetResults = armcosmos.SQLDatabaseGetResults{
	// 	Name: to.Ptr("databaseName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.SQLDatabaseGetProperties{
	// 		Resource: &armcosmos.SQLDatabaseGetPropertiesResource{
	// 			Etag: to.Ptr("\"00000a00-0000-0000-0000-56672f920000\""),
	// 			Rid: to.Ptr("CqNBAA=="),
	// 			Ts: to.Ptr[float32](1449602962),
	// 			ID: to.Ptr("databaseName"),
	// 			Colls: to.Ptr("colls/"),
	// 			Users: to.Ptr("users/"),
	// 		},
	// 	},
	// }
}
