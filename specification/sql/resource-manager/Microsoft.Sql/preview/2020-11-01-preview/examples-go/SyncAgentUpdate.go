package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncAgentUpdate.json
func ExampleSyncAgentsClient_BeginCreateOrUpdate_updateASyncAgent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncAgentsClient().BeginCreateOrUpdate(ctx, "syncagentcrud-65440", "syncagentcrud-8475", "syncagentcrud-3187", armsql.SyncAgent{
		Properties: &armsql.SyncAgentProperties{
			SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/syncagentcrud-8475/databases/sync"),
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
	// res.SyncAgent = armsql.SyncAgent{
	// 	Name: to.Ptr("syncagent"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/syncAgents"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/syncagentcrud-8475/syncAgents/syncagentcrud-3187"),
	// 	Properties: &armsql.SyncAgentProperties{
	// 		ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.9999999Z"); return t}()),
	// 		IsUpToDate: to.Ptr(true),
	// 		State: to.Ptr(armsql.SyncAgentStateNeverConnected),
	// 		SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/syncagentcrud-8475/databases/sync"),
	// 		Version: to.Ptr("4.2.0.0"),
	// 	},
	// }
}
