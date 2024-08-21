package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/WorkspaceConnection/list.json
func ExampleWorkspaceConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceConnectionsClient().NewListPager("resourceGroup-1", "workspace-1", &armmachinelearning.WorkspaceConnectionsClientListOptions{Target: to.Ptr("www.facebook.com"),
		Category: to.Ptr("ContainerRegistry"),
	})
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
		// page.WorkspaceConnectionPropertiesV2BasicResourceArmPaginatedResult = armmachinelearning.WorkspaceConnectionPropertiesV2BasicResourceArmPaginatedResult{
		// 	Value: []*armmachinelearning.WorkspaceConnectionPropertiesV2BasicResource{
		// 		{
		// 			Name: to.Ptr("connection-1"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/linkedWorkspaces/connection-1"),
		// 			Properties: &armmachinelearning.PATAuthTypeWorkspaceConnectionProperties{
		// 				AuthType: to.Ptr(armmachinelearning.ConnectionAuthTypePAT),
		// 				Category: to.Ptr(armmachinelearning.ConnectionCategoryContainerRegistry),
		// 				Target: to.Ptr("www.facebook.com"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("connection-2"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/linkedWorkspaces/connection-2"),
		// 			Properties: &armmachinelearning.PATAuthTypeWorkspaceConnectionProperties{
		// 				AuthType: to.Ptr(armmachinelearning.ConnectionAuthTypePAT),
		// 				Category: to.Ptr(armmachinelearning.ConnectionCategoryContainerRegistry),
		// 				Target: to.Ptr("www.facebook.com"),
		// 			},
		// 	}},
		// }
	}
}
