package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/workspaceManagerGroups/GetAllWorkspaceManagerGroups.json
func ExampleWorkspaceManagerGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceManagerGroupsClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page = armsecurityinsights.WorkspaceManagerGroupsClientListResponse{
		// 	WorkspaceManagerGroupList: armsecurityinsights.WorkspaceManagerGroupList{
		// 		Value: []*armsecurityinsights.WorkspaceManagerGroup{
		// 			{
		// 				Name: to.Ptr("37207a7a-3b8a-438f-a559-c7df400e1b96"),
		// 				Type: to.Ptr("Microsoft.SecurityInsights/workspaceManagerGroups"),
		// 				Etag: to.Ptr("\"ac04c9ad-4b3c-4e13-b511-8c2225e46521\""),
		// 				ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/workspaceManagerGroups/37207a7a-3b8a-438f-a559-c7df400e1b96"),
		// 				Properties: &armsecurityinsights.WorkspaceManagerGroupProperties{
		// 					Description: to.Ptr("Group of all financial and banking institutions"),
		// 					DisplayName: to.Ptr("Banks"),
		// 					MemberResourceNames: []*string{
		// 						to.Ptr("afbd324f-6c48-459c-8710-8d1e1cd03812"),
		// 						to.Ptr("f5fa104e-c0e3-4747-9182-d342dc048a9e"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
