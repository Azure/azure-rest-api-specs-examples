package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/UpdateWorkspaceSetting_example.json
func ExampleWorkspaceSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceSettingsClient().Update(ctx, "default", armsecurity.WorkspaceSetting{
		Name: to.Ptr("default"),
		Type: to.Ptr("Microsoft.Security/workspaceSettings"),
		ID:   to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/workspaceSettings/default"),
		Properties: &armsecurity.WorkspaceSettingProperties{
			WorkspaceID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkspaceSetting = armsecurity.WorkspaceSetting{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Security/workspaceSettings"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/workspaceSettings/default"),
	// 	Properties: &armsecurity.WorkspaceSettingProperties{
	// 		Scope: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
	// 		WorkspaceID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace"),
	// 	},
	// }
}
