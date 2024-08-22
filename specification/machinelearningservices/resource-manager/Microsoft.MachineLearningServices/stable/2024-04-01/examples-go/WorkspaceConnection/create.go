package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/WorkspaceConnection/create.json
func ExampleWorkspaceConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceConnectionsClient().Create(ctx, "resourceGroup-1", "workspace-1", "connection-1", armmachinelearning.WorkspaceConnectionPropertiesV2BasicResource{
		Properties: &armmachinelearning.NoneAuthTypeWorkspaceConnectionProperties{
			AuthType: to.Ptr(armmachinelearning.ConnectionAuthTypeNone),
			Category: to.Ptr(armmachinelearning.ConnectionCategoryContainerRegistry),
			Target:   to.Ptr("www.facebook.com"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkspaceConnectionPropertiesV2BasicResource = armmachinelearning.WorkspaceConnectionPropertiesV2BasicResource{
	// 	Name: to.Ptr("connection-1"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/connections/connection-1"),
	// 	Properties: &armmachinelearning.NoneAuthTypeWorkspaceConnectionProperties{
	// 		AuthType: to.Ptr(armmachinelearning.ConnectionAuthTypeNone),
	// 		Category: to.Ptr(armmachinelearning.ConnectionCategoryContainerRegistry),
	// 		Target: to.Ptr("www.facebook.com"),
	// 	},
	// }
}
