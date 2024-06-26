package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getStatisticsOfAutomationAccount.json
func ExampleStatisticsClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStatisticsClient().NewListByAutomationAccountPager("rg", "myAutomationAccount11", &armautomation.StatisticsClientListByAutomationAccountOptions{Filter: nil})
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
		// page.StatisticsListResult = armautomation.StatisticsListResult{
		// 	Value: []*armautomation.Statistics{
		// 		{
		// 			CounterProperty: to.Ptr("New"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/New"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Activating"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Activating"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Running"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Running"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Completed"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Completed"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Failed"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Failed"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Stopped"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Stopped"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Blocked"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Blocked"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Suspended"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Suspended"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Disconnected"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Disconnected"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Suspending"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Suspending"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Stopping"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Stopping"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Resuming"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Resuming"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 		},
		// 		{
		// 			CounterProperty: to.Ptr("Removing"),
		// 			CounterValue: to.Ptr[int64](0),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:49.987Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11/statistics/Removing"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T02:11:49.987Z"); return t}()),
		// 	}},
		// }
	}
}
