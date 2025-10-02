package armstoragediscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagediscovery/armstoragediscovery"
)

// Generated from example definition: 2025-09-01/StorageDiscoveryWorkspaces_Update.json
func ExampleWorkspacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragediscovery.NewClientFactory("b79cb3ba-745e-5d9a-8903-4a02327a7e09", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Update(ctx, "sample-rg", "Sample-Storage-Workspace", armstoragediscovery.WorkspaceUpdate{
		Properties: &armstoragediscovery.WorkspacePropertiesUpdate{
			SKU: to.Ptr(armstoragediscovery.SKUFree),
			WorkspaceRoots: []*string{
				to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09"),
			},
			Description: to.Ptr("Updated Sample Storage Discovery Workspace"),
			Scopes: []*armstoragediscovery.Scope{
				{
					DisplayName: to.Ptr("Updated-Sample-Collection"),
					ResourceTypes: []*armstoragediscovery.ResourceType{
						to.Ptr(armstoragediscovery.ResourceType("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/updated-sample-storageAccount")),
					},
					TagKeysOnly: []*string{
						to.Ptr("updated-filtertag1"),
						to.Ptr("updated-filtertag2"),
					},
					Tags: map[string]*string{
						"updated-filtertag3": to.Ptr("updated-value3"),
						"updated-filtertag4": to.Ptr("updated-value4"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragediscovery.WorkspacesClientUpdateResponse{
	// 	Workspace: &armstoragediscovery.Workspace{
	// 		ID: to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.StorageDiscovery/storageDiscoveryWorkspaces/sampleworkspace"),
	// 		Name: to.Ptr("sampleworkspace"),
	// 		Type: to.Ptr("Microsoft.StorageDiscovery/storageDiscoveryWorkspaces"),
	// 		Location: to.Ptr("westeurope"),
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 		Properties: &armstoragediscovery.WorkspaceProperties{
	// 			ProvisioningState: to.Ptr(armstoragediscovery.ResourceProvisioningStateSucceeded),
	// 			SKU: to.Ptr(armstoragediscovery.SKUFree),
	// 			WorkspaceRoots: []*string{
	// 				to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09"),
	// 			},
	// 			Description: to.Ptr("Updated Sample Storage Discovery Workspace"),
	// 			Scopes: []*armstoragediscovery.Scope{
	// 				{
	// 					DisplayName: to.Ptr("Updated-Sample-Collection"),
	// 					ResourceTypes: []*armstoragediscovery.ResourceType{
	// 						to.Ptr(armstoragediscovery.ResourceType("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/updated-sample-storageAccount")),
	// 					},
	// 					TagKeysOnly: []*string{
	// 						to.Ptr("updated-filtertag1"),
	// 						to.Ptr("updated-filtertag2"),
	// 					},
	// 					Tags: map[string]*string{
	// 						"updated-filtertag3": to.Ptr("updated-value3"),
	// 						"updated-filtertag4": to.Ptr("updated-value4"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
