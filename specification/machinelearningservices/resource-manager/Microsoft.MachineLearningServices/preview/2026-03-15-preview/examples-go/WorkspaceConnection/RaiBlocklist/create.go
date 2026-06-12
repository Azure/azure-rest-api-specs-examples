package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/WorkspaceConnection/RaiBlocklist/create.json
func ExampleConnectionRaiBlocklistClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionRaiBlocklistClient().BeginCreate(ctx, "test-rg", "aml-workspace-name", "testConnection", "raiBlocklistName", armmachinelearning.RaiBlocklistPropertiesBasicResource{
		Properties: &armmachinelearning.RaiBlocklistProperties{
			Description: to.Ptr("Basic blocklist description"),
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
	// res = armmachinelearning.ConnectionRaiBlocklistClientCreateResponse{
	// 	RaiBlocklistPropertiesBasicResource: armmachinelearning.RaiBlocklistPropertiesBasicResource{
	// 		Name: to.Ptr("raiBlocklistName"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections/raiBlocklists"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.MachineLearningServices/workspaces/aml-workspace-name/connections/testConnection/raiBlocklists/raiBlocklistName"),
	// 		Properties: &armmachinelearning.RaiBlocklistProperties{
	// 			Description: to.Ptr("Basic blocklist description"),
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
	// 			CreatedBy: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
	// 		},
	// 	},
	// }
}
