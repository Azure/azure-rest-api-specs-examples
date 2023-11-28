package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/listJobsByAutomationAccount.json
func ExampleJobClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobClient().NewListByAutomationAccountPager("mygroup", "ContoseAutomationAccount", &armautomation.JobClientListByAutomationAccountOptions{Filter: nil,
		ClientRequestID: nil,
	})
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
		// page.JobListResultV2 = armautomation.JobListResultV2{
		// 	Value: []*armautomation.JobCollectionItem{
		// 		{
		// 			Name: to.Ptr("job1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Jobs"),
		// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobs/job1"),
		// 			Properties: &armautomation.JobCollectionItemProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-30T05:53:30.243Z"); return t}()),
		// 				JobID: to.Ptr("45203a94-a8cb-47c3-8ce4-4dcc3a5f7d23"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-30T05:53:30.243Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Runbook: &armautomation.RunbookAssociationProperty{
		// 					Name: to.Ptr("TestRunbook"),
		// 				},
		// 				Status: to.Ptr(armautomation.JobStatusNew),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("job2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Jobs"),
		// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobs/job2"),
		// 			Properties: &armautomation.JobCollectionItemProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-30T05:46:49.370Z"); return t}()),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-30T05:48:38.857Z"); return t}()),
		// 				JobID: to.Ptr("7584055f-5118-460a-a2dd-5176c9c8efe9"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-30T05:48:38.857Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Runbook: &armautomation.RunbookAssociationProperty{
		// 					Name: to.Ptr("TestRunbook"),
		// 				},
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-30T05:47:35.200Z"); return t}()),
		// 				Status: to.Ptr(armautomation.JobStatusCompleted),
		// 			},
		// 	}},
		// }
	}
}
