package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listWatchersByAutomationAccount.json
func ExampleWatcherClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWatcherClient().NewListByAutomationAccountPager("rg", "MyTestAutomationAccount", &armautomation.WatcherClientListByAutomationAccountOptions{Filter: nil})
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
		// page.WatcherListResult = armautomation.WatcherListResult{
		// 	Value: []*armautomation.Watcher{
		// 		{
		// 			Name: to.Ptr("MyTestWatcher"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/MyTestAutomationAccount/watchers/MyTestWatcher"),
		// 			Properties: &armautomation.WatcherProperties{
		// 				Description: to.Ptr("This is a test watcher."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-20T21:36:48.597Z"); return t}()),
		// 				ExecutionFrequencyInSeconds: to.Ptr[int64](60),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-20T21:36:48.597Z"); return t}()),
		// 				ScriptName: to.Ptr("MyTestWatcher"),
		// 				ScriptRunOn: to.Ptr("MyTestHybridWorkerGroup"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MyTestWatcher01"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/MyTestAutomationAccount/watchers/MyTestWatcher01"),
		// 			Properties: &armautomation.WatcherProperties{
		// 				Description: to.Ptr("This is a test watcher."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-20T20:47:24.697Z"); return t}()),
		// 				ExecutionFrequencyInSeconds: to.Ptr[int64](60),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-20T20:47:24.697Z"); return t}()),
		// 				ScriptName: to.Ptr("MyTestWatcher"),
		// 				ScriptRunOn: to.Ptr("MyTestHybridWorkerGroup"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MyTestWatcher02"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/MyTestAutomationAccount/watchers/MyTestWatcher02"),
		// 			Properties: &armautomation.WatcherProperties{
		// 				Description: to.Ptr("This is a test watcher."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-20T21:26:35.647Z"); return t}()),
		// 				ExecutionFrequencyInSeconds: to.Ptr[int64](60),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-20T21:26:35.647Z"); return t}()),
		// 				ScriptName: to.Ptr("MyTestWatcher"),
		// 				ScriptRunOn: to.Ptr("MyTestHybridWorkerGroup"),
		// 			},
		// 	}},
		// }
	}
}
