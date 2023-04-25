package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b065afacb99e7ec65787383550ee233e0e02a6e6/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBSqlTriggerGet.json
func ExampleSQLResourcesClient_GetSQLTrigger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLResourcesClient().GetSQLTrigger(ctx, "rgName", "ddb1", "databaseName", "containerName", "triggerName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLTriggerGetResults = armcosmos.SQLTriggerGetResults{
	// 	Name: to.Ptr("triggerName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers/sqlTriggers"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/sqlContainers/containerName/sqlTriggers/triggerName"),
	// 	Properties: &armcosmos.SQLTriggerGetProperties{
	// 		Resource: &armcosmos.SQLTriggerGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			Body: to.Ptr("body"),
	// 			ID: to.Ptr("triggerName"),
	// 			TriggerOperation: to.Ptr(armcosmos.TriggerOperation("triggerOperation")),
	// 			TriggerType: to.Ptr(armcosmos.TriggerType("triggerType")),
	// 		},
	// 	},
	// }
}
