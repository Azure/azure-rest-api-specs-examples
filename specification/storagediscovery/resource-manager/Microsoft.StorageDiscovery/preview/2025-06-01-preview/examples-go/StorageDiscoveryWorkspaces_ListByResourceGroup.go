package armstoragediscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagediscovery/armstoragediscovery"
)

// Generated from example definition: 2025-06-01-preview/StorageDiscoveryWorkspaces_ListByResourceGroup.json
func ExampleWorkspacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragediscovery.NewClientFactory("b79cb3ba-745e-5d9a-8903-4a02327a7e09", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListByResourceGroupPager("sample-rg", nil)
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
		// page = armstoragediscovery.WorkspacesClientListByResourceGroupResponse{
		// 	WorkspaceListResult: armstoragediscovery.WorkspaceListResult{
		// 		Value: []*armstoragediscovery.Workspace{
		// 			{
		// 				ID: to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.StorageDiscovery/storageDiscoveryWorkspaces/Primary-Analytics-Workspace"),
		// 				Name: to.Ptr("Primary-Analytics-Workspace"),
		// 				Type: to.Ptr("Microsoft.StorageDiscovery/storageDiscoveryWorkspaces"),
		// 				Location: to.Ptr("westeurope"),
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 				Properties: &armstoragediscovery.WorkspaceProperties{
		// 					ProvisioningState: to.Ptr(armstoragediscovery.ResourceProvisioningStateSucceeded),
		// 					SKU: to.Ptr(armstoragediscovery.SKUStandard),
		// 					WorkspaceRoots: []*string{
		// 						to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09"),
		// 					},
		// 					Description: to.Ptr("Sample Storage Discovery Workspace"),
		// 					Scopes: []*armstoragediscovery.Scope{
		// 						{
		// 							DisplayName: to.Ptr("Sample-Collection-1"),
		// 							ResourceTypes: []*armstoragediscovery.ResourceType{
		// 								to.Ptr(armstoragediscovery.ResourceType("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/sample-storageAccount")),
		// 							},
		// 							TagKeysOnly: []*string{
		// 								to.Ptr("filterTag1"),
		// 								to.Ptr("filterTag2"),
		// 							},
		// 							Tags: map[string]*string{
		// 								"filterTag3": to.Ptr("value3"),
		// 								"filterTag4": to.Ptr("value4"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.StorageDiscovery/storageDiscoveryWorkspaces/Secondary-Analytics-Workspace"),
		// 				Name: to.Ptr("Secondary-Analytics-Workspace"),
		// 				Type: to.Ptr("Microsoft.StorageDiscovery/storageDiscoveryWorkspaces"),
		// 				Location: to.Ptr("westeurope"),
		// 				Tags: map[string]*string{
		// 					"tag3": to.Ptr("value1"),
		// 					"tag4": to.Ptr("value2"),
		// 				},
		// 				Properties: &armstoragediscovery.WorkspaceProperties{
		// 					ProvisioningState: to.Ptr(armstoragediscovery.ResourceProvisioningStateSucceeded),
		// 					SKU: to.Ptr(armstoragediscovery.SKU("Premium")),
		// 					WorkspaceRoots: []*string{
		// 						to.Ptr("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09"),
		// 					},
		// 					Description: to.Ptr("Sample Storage Discovery Workspace"),
		// 					Scopes: []*armstoragediscovery.Scope{
		// 						{
		// 							DisplayName: to.Ptr("Sample-Collection-2"),
		// 							ResourceTypes: []*armstoragediscovery.ResourceType{
		// 								to.Ptr(armstoragediscovery.ResourceType("/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/sample-storageAccount")),
		// 							},
		// 							TagKeysOnly: []*string{
		// 								to.Ptr("filterTag5"),
		// 							},
		// 							Tags: map[string]*string{
		// 								"filterTag6": to.Ptr("value6"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
