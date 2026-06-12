package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/listByResourceGroup.json
func ExampleWorkspacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListByResourceGroupPager("workspace-1234", nil)
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
		// page = armmachinelearning.WorkspacesClientListByResourceGroupResponse{
		// 	WorkspaceListResult: armmachinelearning.WorkspaceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.Workspace{
		// 			{
		// 				Name: to.Ptr("testworkspace"),
		// 				Type: to.Ptr("Microsoft.MachineLearningServices/workspaces"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace"),
		// 				Location: to.Ptr("eastus2euap"),
		// 				Properties: &armmachinelearning.WorkspaceProperties{
		// 					Description: to.Ptr("test description"),
		// 					ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
		// 					ContainerRegistry: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry"),
		// 					DiscoveryURL: to.Ptr("http://example.com"),
		// 					FriendlyName: to.Ptr("HelloName"),
		// 					KeyVault: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
		// 					StorageAccount: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testworkspace1"),
		// 				Type: to.Ptr("Microsoft.MachineLearningServices/workspaces"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace1"),
		// 				Location: to.Ptr("eastus2euap"),
		// 				Properties: &armmachinelearning.WorkspaceProperties{
		// 					Description: to.Ptr("test description"),
		// 					ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
		// 					ContainerRegistry: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistryNew"),
		// 					DiscoveryURL: to.Ptr("http://example.com"),
		// 					FriendlyName: to.Ptr("HelloName 1"),
		// 					KeyVault: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkvNew"),
		// 					StorageAccount: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccountOld"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
