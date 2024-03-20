package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/CosmosDBManagedCassandraCommandAsync.json
func ExampleCassandraClustersClient_BeginInvokeCommandAsync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginInvokeCommandAsync(ctx, "cassandra-prod-rg", "cassandra-prod", armcosmos.CommandPostBody{
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
	// res.CommandPublicResource = armcosmos.CommandPublicResource{
	// 	Arguments: map[string]any{
	// 		"status": "",
	// 	},
	// 	CassandraStopStart: to.Ptr(false),
	// 	Command: to.Ptr("nodetool"),
	// 	Host: to.Ptr("10.0.1.12"),
	// 	IsAdmin: to.Ptr(false),
	// 	OutputFile: to.Ptr("301fc7c1-397f-4e4a-89db-478f61f89d67"),
	// 	ReadWrite: to.Ptr(false),
	// }
}
