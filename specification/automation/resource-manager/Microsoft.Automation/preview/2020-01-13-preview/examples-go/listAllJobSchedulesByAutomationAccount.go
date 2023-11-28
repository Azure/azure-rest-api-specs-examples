package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listAllJobSchedulesByAutomationAccount.json
func ExampleJobScheduleClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobScheduleClient().NewListByAutomationAccountPager("rg", "ContoseAutomationAccount", &armautomation.JobScheduleClientListByAutomationAccountOptions{Filter: nil})
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
		// page.JobScheduleListResult = armautomation.JobScheduleListResult{
		// 	Value: []*armautomation.JobSchedule{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobSchedules/2f4d8f35-ecd5-44ee-a019-2382fec58fb7"),
		// 			Properties: &armautomation.JobScheduleProperties{
		// 				JobScheduleID: to.Ptr("2f4d8f35-ecd5-44ee-a019-2382fec58fb7"),
		// 				Runbook: &armautomation.RunbookAssociationProperty{
		// 					Name: to.Ptr("TestRunbook"),
		// 				},
		// 				Schedule: &armautomation.ScheduleAssociationProperty{
		// 					Name: to.Ptr("JobScheduleforTestRunbook"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobSchedules/446f7a33-86ff-45a1-b71c-f998f701b443"),
		// 			Properties: &armautomation.JobScheduleProperties{
		// 				JobScheduleID: to.Ptr("446f7a33-86ff-45a1-b71c-f998f701b443"),
		// 				Runbook: &armautomation.RunbookAssociationProperty{
		// 					Name: to.Ptr("TestRunbook"),
		// 				},
		// 				Schedule: &armautomation.ScheduleAssociationProperty{
		// 					Name: to.Ptr("TestSchedule"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobSchedules/0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc"),
		// 			Properties: &armautomation.JobScheduleProperties{
		// 				JobScheduleID: to.Ptr("0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc"),
		// 				Runbook: &armautomation.RunbookAssociationProperty{
		// 					Name: to.Ptr("TestRunbook"),
		// 				},
		// 				Schedule: &armautomation.ScheduleAssociationProperty{
		// 					Name: to.Ptr("ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
