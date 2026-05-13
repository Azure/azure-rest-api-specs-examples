package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBManagedCassandraCommandAsync.json
func ExampleCassandraClustersClient_BeginInvokeCommandAsync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginInvokeCommandAsync(ctx, "cassandra-prod-rg", "cassandra-prod", armcosmos.CommandAsyncPostBody{
		Arguments: map[string]any{
			"status": "",
		},
		Command: to.Ptr("nodetool"),
		Host:    to.Ptr("10.0.1.12"),
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
	// res = armcosmos.CassandraClustersClientInvokeCommandAsyncResponse{
	// 	CommandPublicResource: &armcosmos.CommandPublicResource{
	// 		Arguments: map[string]any{
	// 			"status": "",
	// 		},
	// 		CassandraStopStart: to.Ptr(false),
	// 		Command: to.Ptr("nodetool"),
	// 		Host: to.Ptr("10.0.1.12"),
	// 		IsAdmin: to.Ptr(false),
	// 		OutputFile: to.Ptr("301fc7c1-397f-4e4a-89db-478f61f89d67"),
	// 		ReadWrite: to.Ptr(false),
	// 	},
	// }
}
