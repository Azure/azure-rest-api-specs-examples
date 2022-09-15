package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBSqlTriggerCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateSQLTrigger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewSQLResourcesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdateSQLTrigger(ctx, "rg1", "ddb1", "databaseName", "containerName", "triggerName", armcosmos.SQLTriggerCreateUpdateParameters{
		Properties: &armcosmos.SQLTriggerCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.SQLTriggerResource{
				Body:             to.Ptr("body"),
				ID:               to.Ptr("triggerName"),
				TriggerOperation: to.Ptr(armcosmos.TriggerOperation("triggerOperation")),
				TriggerType:      to.Ptr(armcosmos.TriggerType("triggerType")),
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
	// TODO: use response item
	_ = res
}
