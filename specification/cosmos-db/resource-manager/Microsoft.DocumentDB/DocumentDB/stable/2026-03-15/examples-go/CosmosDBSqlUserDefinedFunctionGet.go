package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBSqlUserDefinedFunctionGet.json
func ExampleSQLResourcesClient_GetSQLUserDefinedFunction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLResourcesClient().GetSQLUserDefinedFunction(ctx, "rgName", "ddb1", "databaseName", "containerName", "userDefinedFunctionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.SQLResourcesClientGetSQLUserDefinedFunctionResponse{
	// 	SQLUserDefinedFunctionGetResults: armcosmos.SQLUserDefinedFunctionGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/sqlContainers/containerName/sqlUserDefinedFunctions/userDefinedFunctionName"),
	// 		Name: to.Ptr("userDefinedFunctionName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers/sqlUserDefinedFunctions"),
	// 		Properties: &armcosmos.SQLUserDefinedFunctionGetProperties{
	// 			Resource: &armcosmos.SQLUserDefinedFunctionGetPropertiesResource{
	// 				ID: to.Ptr("userDefinedFunctionName"),
	// 				Body: to.Ptr("body"),
	// 				Rid: to.Ptr("PD5DALigDgw="),
	// 				Ts: to.Ptr[float32](1459200611),
	// 				Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			},
	// 		},
	// 	},
	// }
}
