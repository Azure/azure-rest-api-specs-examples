package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupTriggerSync.json
func ExampleSyncGroupsClient_TriggerSync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewSyncGroupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.TriggerSync(ctx,
		"syncgroupcrud-65440",
		"syncgroupcrud-8475",
		"syncgroupcrud-4328",
		"syncgroupcrud-3187",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
