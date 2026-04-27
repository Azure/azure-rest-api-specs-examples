package armplaywright_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwright/armplaywright"
)

// Generated from example definition: 2026-02-01-preview/PlaywrightWorkspaces_ListBySubscription.json
func ExampleWorkspacesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywright.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListBySubscriptionPager(nil)
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
		// page = armplaywright.WorkspacesClientListBySubscriptionResponse{
		// 	WorkspaceListResult: armplaywright.WorkspaceListResult{
		// 		Value: []*armplaywright.Workspace{
		// 			{
		// 				Location: to.Ptr("westus3"),
		// 				Properties: &armplaywright.WorkspaceProperties{
		// 					DataplaneURI: to.Ptr("https://api.dataplane.00000000-0000-0000-0000-000000000000.domain.com"),
		// 					RegionalAffinity: to.Ptr(armplaywright.EnablementStatusEnabled),
		// 					LocalAuth: to.Ptr(armplaywright.EnablementStatusEnabled),
		// 					Reporting: to.Ptr(armplaywright.EnablementStatusEnabled),
		// 					WorkspaceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					StorageURI: to.Ptr("https://examplestorageaccount.blob.core.windows.net"),
		// 					ProvisioningState: to.Ptr(armplaywright.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.LoadTestService/PlaywrightWorkspaces/myWorkspace"),
		// 				Name: to.Ptr("myWorkspace"),
		// 				Type: to.Ptr("Microsoft.LoadTestService/PlaywrightWorkspaces"),
		// 				Tags: map[string]*string{
		// 					"Team": to.Ptr("Dev Exp"),
		// 				},
		// 				SystemData: &armplaywright.SystemData{
		// 					CreatedBy: to.Ptr("userId1001"),
		// 					CreatedByType: to.Ptr(armplaywright.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("userId1001"),
		// 					LastModifiedByType: to.Ptr(armplaywright.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
