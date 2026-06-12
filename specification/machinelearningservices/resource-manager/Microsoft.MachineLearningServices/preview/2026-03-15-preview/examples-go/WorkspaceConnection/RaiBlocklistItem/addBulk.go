package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/WorkspaceConnection/RaiBlocklistItem/addBulk.json
func ExampleConnectionRaiBlocklistItemClient_BeginAddBulk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionRaiBlocklistItemClient().BeginAddBulk(ctx, "test-rg", "aml-workspace-name", "testConnection", "raiBlocklistName", []*armmachinelearning.RaiBlocklistItemBulkRequest{
		{
			Name: to.Ptr("myblocklistitem1"),
			Properties: &armmachinelearning.RaiBlocklistItemProperties{
				IsRegex: to.Ptr(true),
				Pattern: to.Ptr("^[a-z0-9_-]{2,16}$"),
			},
		},
		{
			Name: to.Ptr("myblocklistitem2"),
			Properties: &armmachinelearning.RaiBlocklistItemProperties{
				IsRegex: to.Ptr(false),
				Pattern: to.Ptr("blockwords"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.ConnectionRaiBlocklistItemClientAddBulkResponse{
	// }
}
