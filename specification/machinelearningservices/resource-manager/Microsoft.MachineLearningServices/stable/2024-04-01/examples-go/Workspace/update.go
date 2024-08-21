package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/update.json
func ExampleWorkspacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginUpdate(ctx, "workspace-1234", "testworkspace", armmachinelearning.WorkspaceUpdateParameters{
		Properties: &armmachinelearning.WorkspacePropertiesUpdateParameters{
			Description:         to.Ptr("new description"),
			FriendlyName:        to.Ptr("New friendly name"),
			PublicNetworkAccess: to.Ptr(armmachinelearning.PublicNetworkAccessDisabled),
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
	// res.Workspace = armmachinelearning.Workspace{
	// 	Name: to.Ptr("testworkspace"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace"),
	// 	Identity: &armmachinelearning.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 	},
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armmachinelearning.WorkspaceProperties{
	// 		Description: to.Ptr("new description"),
	// 		ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
	// 		ContainerRegistry: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry"),
	// 		DiscoveryURL: to.Ptr("http://example.com"),
	// 		FriendlyName: to.Ptr("New friendly name"),
	// 		KeyVault: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
	// 		PublicNetworkAccess: to.Ptr(armmachinelearning.PublicNetworkAccessDisabled),
	// 		StorageAccount: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
	// 	},
	// }
}
