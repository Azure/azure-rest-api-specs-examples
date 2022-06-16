package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncMemberCreate.json
func ExampleSyncMembersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewSyncMembersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"syncgroupcrud-65440",
		"syncgroupcrud-8475",
		"syncgroupcrud-4328",
		"syncgroupcrud-3187",
		"syncmembercrud-4879",
		armsql.SyncMember{
			Properties: &armsql.SyncMemberProperties{
				DatabaseName:                      to.Ptr("syncgroupcrud-7421"),
				DatabaseType:                      to.Ptr(armsql.SyncMemberDbTypeAzureSQLDatabase),
				ServerName:                        to.Ptr("syncgroupcrud-3379.database.windows.net"),
				SyncDirection:                     to.Ptr(armsql.SyncDirectionBidirectional),
				SyncMemberAzureDatabaseResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
				UsePrivateLinkConnection:          to.Ptr(true),
				UserName:                          to.Ptr("myUser"),
			},
		},
		nil)
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
