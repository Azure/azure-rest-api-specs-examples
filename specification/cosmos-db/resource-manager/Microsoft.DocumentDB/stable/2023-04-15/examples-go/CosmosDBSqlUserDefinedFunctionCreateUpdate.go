package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBSqlUserDefinedFunctionCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateSQLUserDefinedFunction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginCreateUpdateSQLUserDefinedFunction(ctx, "rg1", "ddb1", "databaseName", "containerName", "userDefinedFunctionName", armcosmos.SQLUserDefinedFunctionCreateUpdateParameters{
		Properties: &armcosmos.SQLUserDefinedFunctionCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.SQLUserDefinedFunctionResource{
				Body: to.Ptr("body"),
				ID:   to.Ptr("userDefinedFunctionName"),
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
	// res.SQLUserDefinedFunctionGetResults = armcosmos.SQLUserDefinedFunctionGetResults{
	// 	Name: to.Ptr("userDefinedFunctionName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers/sqlUserDefinedFunctions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/sqlContainers/containerName/sqlUserDefinedFunctions/userDefinedFunctionName"),
	// 	Properties: &armcosmos.SQLUserDefinedFunctionGetProperties{
	// 		Resource: &armcosmos.SQLUserDefinedFunctionGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			Body: to.Ptr("body"),
	// 			ID: to.Ptr("userDefinedFunctionName"),
	// 		},
	// 	},
	// }
}
