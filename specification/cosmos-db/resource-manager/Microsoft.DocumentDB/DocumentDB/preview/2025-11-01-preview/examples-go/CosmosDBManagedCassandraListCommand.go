package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBManagedCassandraListCommand.json
func ExampleCassandraClustersClient_NewListCommandPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraClustersClient().NewListCommandPager("cassandra-prod-rg", "cassandra-prod", nil)
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
		// page = armcosmos.CassandraClustersClientListCommandResponse{
		// 	ListCommands: armcosmos.ListCommands{
		// 		Value: []*armcosmos.CommandPublicResource{
		// 			{
		// 				Arguments: map[string]any{
		// 					"status": "",
		// 				},
		// 				CassandraStopStart: to.Ptr(false),
		// 				Command: to.Ptr("nodetool"),
		// 				CommandID: to.Ptr("1234"),
		// 				Host: to.Ptr("10.0.1.12"),
		// 				ReadWrite: to.Ptr(true),
		// 				Status: to.Ptr(armcosmos.CommandStatus("Enqueued")),
		// 			},
		// 		},
		// 	},
		// }
	}
}
