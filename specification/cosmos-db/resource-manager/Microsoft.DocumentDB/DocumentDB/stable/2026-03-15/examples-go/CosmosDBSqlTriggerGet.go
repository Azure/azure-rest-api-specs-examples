package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBSqlTriggerGet.json
func ExampleSQLResourcesClient_GetSQLTrigger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
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
	// res = armcosmos.SQLResourcesClientGetSQLTriggerResponse{
	// 	SQLTriggerGetResults: armcosmos.SQLTriggerGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/sqlContainers/containerName/sqlTriggers/triggerName"),
	// 		Name: to.Ptr("triggerName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/sqlContainers/sqlTriggers"),
	// 		Properties: &armcosmos.SQLTriggerGetProperties{
	// 			Resource: &armcosmos.SQLTriggerGetPropertiesResource{
	// 				ID: to.Ptr("triggerName"),
	// 				Body: to.Ptr("body"),
	// 				TriggerType: to.Ptr(armcosmos.TriggerType("triggerType")),
	// 				TriggerOperation: to.Ptr(armcosmos.TriggerOperation("triggerOperation")),
	// 				Rid: to.Ptr("PD5DALigDgw="),
	// 				Ts: to.Ptr[float32](1459200611),
	// 				Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			},
	// 		},
	// 	},
	// }
}
