package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/listHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHybridRunbookWorkerGroupClient().NewListByAutomationAccountPager("rg", "testaccount", &armautomation.HybridRunbookWorkerGroupClientListByAutomationAccountOptions{Filter: nil})
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
		// page.HybridRunbookWorkerGroupsListResult = armautomation.HybridRunbookWorkerGroupsListResult{
		// 	Value: []*armautomation.HybridRunbookWorkerGroup{
		// 		{
		// 			Name: to.Ptr("TestHybridGroup"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/testaccount/hybridRunbookWorkerGroups/TestHybridGroup"),
		// 			Properties: &armautomation.HybridRunbookWorkerGroupProperties{
		// 				Credential: &armautomation.RunAsCredentialAssociationProperty{
		// 					Name: to.Ptr("myRunAsCredentialName"),
		// 				},
		// 				GroupType: to.Ptr(armautomation.GroupTypeEnumUser),
		// 			},
		// 			SystemData: &armautomation.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("foo@contoso.com"),
		// 				CreatedByType: to.Ptr(armautomation.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armautomation.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
