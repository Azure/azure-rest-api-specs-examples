package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listCompilationJobsByAutomationAccount.json
func ExampleDscCompilationJobClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDscCompilationJobClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", &armautomation.DscCompilationJobClientListByAutomationAccountOptions{Filter: nil})
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
		// page.DscCompilationJobListResult = armautomation.DscCompilationJobListResult{
		// 	Value: []*armautomation.DscCompilationJob{
		// 		{
		// 			Name: to.Ptr("CompilationConfiguration1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/compilationjobs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/"),
		// 			Properties: &armautomation.DscCompilationJobProperties{
		// 				Configuration: &armautomation.DscConfigurationAssociationProperty{
		// 					Name: to.Ptr("TestDscConfiguration"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T19:45:24.590Z"); return t}()),
		// 				JobID: to.Ptr("e6e7fbab-183c-405a-afe6-9eb5db97921a"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T19:45:58.593Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T19:45:52.983Z"); return t}()),
		// 				Status: to.Ptr(armautomation.JobStatusSuspended),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CompilationConfiguration2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/compilationjobs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/"),
		// 			Properties: &armautomation.DscCompilationJobProperties{
		// 				Configuration: &armautomation.DscConfigurationAssociationProperty{
		// 					Name: to.Ptr("NewDscConfiguration"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:29:07.740Z"); return t}()),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:42.600Z"); return t}()),
		// 				JobID: to.Ptr("111d4e06-2d88-46b4-8500-7febd4906838"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:42.600Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:26.480Z"); return t}()),
		// 				Status: to.Ptr(armautomation.JobStatusCompleted),
		// 			},
		// 	}},
		// }
	}
}
