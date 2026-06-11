package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/workspaceManagerMembers/GetWorkspaceManagerMember.json
func ExampleWorkspaceManagerMembersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceManagerMembersClient().Get(ctx, "myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.WorkspaceManagerMembersClientGetResponse{
	// 	WorkspaceManagerMember: armsecurityinsights.WorkspaceManagerMember{
	// 		Name: to.Ptr("afbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/workspaceManagerMembers"),
	// 		Etag: to.Ptr("\"190057d0-0000-0d00-0000-5c6f5adb0000\""),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/workspaceManagerMembers/afbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 		Properties: &armsecurityinsights.WorkspaceManagerMemberProperties{
	// 			TargetWorkspaceResourceID: to.Ptr("/subscriptions/7aef9d48-814f-45ad-b644-b0343316e174/resourceGroups/otherRg/providers/Microsoft.OperationalInsights/workspaces/Example_Workspace"),
	// 			TargetWorkspaceTenantID: to.Ptr("f676d436-8d16-42db-81b7-ab578e110ccd"),
	// 		},
	// 	},
	// }
}
