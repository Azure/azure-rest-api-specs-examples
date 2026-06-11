package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/workspaceManagerAssignments/GetAllWorkspaceManagerAssignments.json
func ExampleWorkspaceManagerAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceManagerAssignmentsClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page = armsecurityinsights.WorkspaceManagerAssignmentsClientListResponse{
		// 	WorkspaceManagerAssignmentList: armsecurityinsights.WorkspaceManagerAssignmentList{
		// 		Value: []*armsecurityinsights.WorkspaceManagerAssignment{
		// 			{
		// 				Name: to.Ptr("47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58"),
		// 				Type: to.Ptr("Microsoft.SecurityInsights/workspaceManagerAssignments"),
		// 				Etag: to.Ptr("\"190057d0-0000-0d00-0000-5c6f5adb0000\""),
		// 				ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58"),
		// 				Properties: &armsecurityinsights.WorkspaceManagerAssignmentProperties{
		// 					Items: []*armsecurityinsights.AssignmentItem{
		// 						{
		// 							ResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspac-es/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExampleOne"),
		// 						},
		// 						{
		// 							ResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspac-es/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExampleTwo"),
		// 						},
		// 					},
		// 					LastJobEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-14T04:52:52.9614956Z"); return t}()),
		// 					LastJobProvisioningState: to.Ptr(armsecurityinsights.ProvisioningStateFailed),
		// 					TargetResourceName: to.Ptr("37207a7a-3b8a-438f-a559-c7df400e1b96"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
