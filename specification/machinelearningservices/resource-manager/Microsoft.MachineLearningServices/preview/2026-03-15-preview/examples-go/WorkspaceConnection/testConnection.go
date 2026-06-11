package armmachinelearning_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/WorkspaceConnection/testConnection.json
func ExampleWorkspaceConnectionsClient_BeginTestConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceConnectionsClient().BeginTestConnection(ctx, "resourceGroup-1", "workspace-1", "connection-1", &armmachinelearning.WorkspaceConnectionsClientBeginTestConnectionOptions{
		Body: &armmachinelearning.WorkspaceConnectionPropertiesV2BasicResource{
			Properties: &armmachinelearning.NoneAuthTypeWorkspaceConnectionProperties{
				AuthType:   to.Ptr(armmachinelearning.ConnectionAuthTypeNone),
				Category:   to.Ptr(armmachinelearning.ConnectionCategoryContainerRegistry),
				ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-15T14:30:00Z"); return t }()),
				Target:     to.Ptr("target_url"),
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
