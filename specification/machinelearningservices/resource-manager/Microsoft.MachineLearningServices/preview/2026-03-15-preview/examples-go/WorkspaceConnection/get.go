package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/WorkspaceConnection/get.json
func ExampleWorkspaceConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceConnectionsClient().Get(ctx, "resourceGroup-1", "workspace-1", "connection-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.WorkspaceConnectionsClientGetResponse{
	// 	WorkspaceConnectionPropertiesV2BasicResource: armmachinelearning.WorkspaceConnectionPropertiesV2BasicResource{
	// 		Name: to.Ptr("connection-1"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/connections/connection-1"),
	// 		Properties: &armmachinelearning.NoneAuthTypeWorkspaceConnectionProperties{
	// 			AuthType: to.Ptr(armmachinelearning.ConnectionAuthTypeNone),
	// 			Category: to.Ptr(armmachinelearning.ConnectionCategoryContainerRegistry),
	// 			ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-15T14:30:00Z"); return t}()),
	// 			Target: to.Ptr("www.facebook.com"),
	// 		},
	// 	},
	// }
}
