package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncAgentGetLinkedDatabases.json
func ExampleSyncAgentsClient_NewListLinkedDatabasesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncAgentsClient().NewListLinkedDatabasesPager("syncagentcrud-65440", "syncagentcrud-8475", "syncagentcrud-3187", nil)
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
		// page.SyncAgentLinkedDatabaseListResult = armsql.SyncAgentLinkedDatabaseListResult{
		// 	Value: []*armsql.SyncAgentLinkedDatabase{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/syncagentcrud-8475/syncAgents/syncagentcrud-3187/linkedDatabases/55555555-6666-7777-8888-999999999999"),
		// 			Properties: &armsql.SyncAgentLinkedDatabaseProperties{
		// 				Description: to.Ptr(""),
		// 				DatabaseID: to.Ptr("55555555-6666-7777-8888-999999999999"),
		// 				DatabaseName: to.Ptr("DummySqlServerDb"),
		// 				DatabaseType: to.Ptr(armsql.SyncMemberDbTypeSQLServerDatabase),
		// 				ServerName: to.Ptr("DummySqlServer"),
		// 				UserName: to.Ptr("DummyUser"),
		// 			},
		// 	}},
		// }
	}
}
