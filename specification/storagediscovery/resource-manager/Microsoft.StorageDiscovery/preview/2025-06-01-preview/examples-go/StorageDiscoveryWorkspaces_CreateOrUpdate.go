package armstoragediscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagediscovery/armstoragediscovery"
)

// Generated from example definition: 2025-06-01-preview/StorageDiscoveryWorkspaces_CreateOrUpdate.json
func ExampleWorkspacesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragediscovery.NewClientFactory("b79cb3ba-745e-5d9a-8903-4a02327a7e09", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().CreateOrUpdate(ctx, "sample-rg", "Sample-Storage-Workspace", armstoragediscovery.Workspace{
		Location: to.Ptr("westeurope"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		Properties: &armstoragediscovery.WorkspaceProperties{
			WorkspaceRoots: []*string{
				to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09"),
			},
			Description: to.Ptr("Sample Storage Discovery Workspace"),
			Scopes: []*armstoragediscovery.Scope{
				{
					DisplayName: to.Ptr("Sample-Collection"),
					ResourceTypes: []*armstoragediscovery.ResourceType{
						to.Ptr(armstoragediscovery.ResourceType("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/sample-storageAccount")),
					},
					TagKeysOnly: []*string{
						to.Ptr("filterTag1"),
						to.Ptr("filterTag2"),
					},
					Tags: map[string]*string{
						"filterTag3": to.Ptr("value3"),
						"filterTag4": to.Ptr("value4"),
					},
				},
				{
					DisplayName: to.Ptr("Sample-Collection-2"),
					ResourceTypes: []*armstoragediscovery.ResourceType{
						to.Ptr(armstoragediscovery.ResourceType("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/sample-storageAccount")),
					},
					TagKeysOnly: []*string{
						to.Ptr("filterTag5"),
					},
					Tags: map[string]*string{
						"filterTag6": to.Ptr("value6"),
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
	// res = armstoragediscovery.WorkspacesClientCreateOrUpdateResponse{
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
	// 			SKU: to.Ptr(armstoragediscovery.SKUStandard),
	// 			WorkspaceRoots: []*string{
	// 				to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09"),
	// 			},
	// 			Description: to.Ptr("Sample Storage Discovery Workspace"),
	// 			Scopes: []*armstoragediscovery.Scope{
	// 				{
	// 					DisplayName: to.Ptr("Sample-Collection"),
	// 					ResourceTypes: []*armstoragediscovery.ResourceType{
	// 						to.Ptr(armstoragediscovery.ResourceType("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/sample-storageAccount")),
	// 					},
	// 					TagKeysOnly: []*string{
	// 						to.Ptr("filterTag1"),
	// 						to.Ptr("filterTag2"),
	// 					},
	// 					Tags: map[string]*string{
	// 						"filterTag3": to.Ptr("value3"),
	// 						"filterTag4": to.Ptr("value4"),
	// 					},
	// 				},
	// 				{
	// 					DisplayName: to.Ptr("Sample-Collection-2"),
	// 					ResourceTypes: []*armstoragediscovery.ResourceType{
	// 						to.Ptr(armstoragediscovery.ResourceType("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/sample-storageAccount")),
	// 					},
	// 					TagKeysOnly: []*string{
	// 						to.Ptr("filterTag5"),
	// 					},
	// 					Tags: map[string]*string{
	// 						"filterTag6": to.Ptr("value6"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armstoragediscovery.SystemData{
	// 			CreatedBy: to.Ptr("iewyxsnriqktsvp"),
	// 			CreatedByType: to.Ptr(armstoragediscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xrchbnnuzierzpxw"),
	// 			LastModifiedByType: to.Ptr(armstoragediscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
	// 		},
	// 	},
	// }
}
