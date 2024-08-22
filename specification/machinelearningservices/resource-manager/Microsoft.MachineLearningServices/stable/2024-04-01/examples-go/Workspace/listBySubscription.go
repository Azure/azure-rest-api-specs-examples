package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/listBySubscription.json
func ExampleWorkspacesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListBySubscriptionPager(&armmachinelearning.WorkspacesClientListBySubscriptionOptions{Skip: nil})
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
		// page.WorkspaceListResult = armmachinelearning.WorkspaceListResult{
		// 	Value: []*armmachinelearning.Workspace{
		// 		{
		// 			Name: to.Ptr("testworkspace"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armmachinelearning.WorkspaceProperties{
		// 				Description: to.Ptr("test description"),
		// 				ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
		// 				ContainerRegistry: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry"),
		// 				DiscoveryURL: to.Ptr("http://example.com"),
		// 				FriendlyName: to.Ptr("HelloName"),
		// 				KeyVault: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
		// 				StorageAccount: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testworkspace"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-5678/providers/Microsoft.MachineLearningServices/workspaces/testworkspace"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armmachinelearning.WorkspaceProperties{
		// 				Description: to.Ptr("test description"),
		// 				ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
		// 				ContainerRegistry: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistryNew"),
		// 				DiscoveryURL: to.Ptr("http://example.com"),
		// 				FriendlyName: to.Ptr("HelloName"),
		// 				KeyVault: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkvNew"),
		// 				StorageAccount: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccountOld"),
		// 			},
		// 	}},
		// }
	}
}
