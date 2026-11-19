package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/UpdateProjectWithToolControls.json
func ExampleProjectsClient_BeginUpdate_updateAFoundryProjectWithNoExplicitlyAllowedTools() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginUpdate(ctx, "myResourceGroup", "myAccount", "myProject", armcognitiveservices.Project{
		Properties: &armcognitiveservices.ProjectProperties{
			ToolControls: &armcognitiveservices.ToolControls{
				DefaultAction: to.Ptr(armcognitiveservices.ToolControlsDefaultActionDeny),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.ProjectsClientUpdateResponse{
	// 	Project: armcognitiveservices.Project{
	// 		Name: to.Ptr("myProject"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount/projects/myProject"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armcognitiveservices.ProjectProperties{
	// 			DisplayName: to.Ptr("My project"),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			ToolControls: &armcognitiveservices.ToolControls{
	// 				DefaultAction: to.Ptr(armcognitiveservices.ToolControlsDefaultActionDeny),
	// 			},
	// 		},
	// 	},
	// }
}
