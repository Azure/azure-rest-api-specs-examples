package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/listHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHybridRunbookWorkerGroupClient().NewListByAutomationAccountPager("rg", "testaccount", nil)
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
		// page = armautomation.HybridRunbookWorkerGroupClientListByAutomationAccountResponse{
		// 	HybridRunbookWorkerGroupsListResult: armautomation.HybridRunbookWorkerGroupsListResult{
		// 		Value: []*armautomation.HybridRunbookWorkerGroup{
		// 			{
		// 				Name: to.Ptr("TestHybridGroup"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/testaccount/hybridRunbookWorkerGroups/TestHybridGroup"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.HybridRunbookWorkerGroupProperties{
		// 					Credential: &armautomation.RunAsCredentialAssociationProperty{
		// 						Name: to.Ptr("myRunAsCredentialName"),
		// 					},
		// 					GroupType: to.Ptr(armautomation.GroupTypeEnumUser),
		// 				},
		// 				SystemData: &armautomation.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					CreatedBy: to.Ptr("foo@contoso.com"),
		// 					CreatedByType: to.Ptr(armautomation.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armautomation.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
