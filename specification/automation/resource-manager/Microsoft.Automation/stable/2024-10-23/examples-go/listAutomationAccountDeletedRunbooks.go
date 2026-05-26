package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/listAutomationAccountDeletedRunbooks.json
func ExampleAccountClient_NewListDeletedRunbooksPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountClient().NewListDeletedRunbooksPager("rg", "MyAutomationAccount", nil)
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
		// page = armautomation.AccountClientListDeletedRunbooksResponse{
		// 	DeletedRunbookListResult: armautomation.DeletedRunbookListResult{
		// 		Value: []*armautomation.DeletedRunbook{
		// 			{
		// 				Name: to.Ptr("myrunbook1"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/MyAutomationAccount/runbooks/myrunbook1"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.DeletedRunbookProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-28T21:32:25.78+00:00"); return t}()),
		// 					DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-30T21:32:25.81+00:00"); return t}()),
		// 					RunbookID: to.Ptr("cb855f13-0223-4fe4-8260-9e6583dfef24"),
		// 					RunbookType: to.Ptr("PowerShell"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myrunbook2"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/MyAutomationAccount/runbooks/myrunbook2"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.DeletedRunbookProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-28T21:32:25.78+00:00"); return t}()),
		// 					DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-30T21:32:25.81+00:00"); return t}()),
		// 					RunbookID: to.Ptr("cb855f13-0223-4fe4-8260-9e6583dfef24"),
		// 					RunbookType: to.Ptr("PowerShell"),
		// 					Runtime: to.Ptr("PowerShell-7.2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myrunbook3"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/MyAutomationAccount/runbooks/myrunbook3"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.DeletedRunbookProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-28T21:32:25.78+00:00"); return t}()),
		// 					DeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-30T21:32:25.81+00:00"); return t}()),
		// 					RunbookID: to.Ptr("cb855f13-0223-4fe4-8260-9e6583dfef24"),
		// 					RunbookType: to.Ptr("PowerShell"),
		// 					RuntimeEnvironment: to.Ptr("environmentName"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
