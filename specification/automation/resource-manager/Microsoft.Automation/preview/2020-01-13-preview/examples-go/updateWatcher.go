package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateWatcher.json
func ExampleWatcherClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatcherClient().Update(ctx, "rg", "MyTestAutomationAccount", "MyTestWatcher", armautomation.WatcherUpdateParameters{
		Name: to.Ptr("MyTestWatcher"),
		Properties: &armautomation.WatcherUpdateProperties{
			ExecutionFrequencyInSeconds: to.Ptr[int64](600),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Watcher = armautomation.Watcher{
	// 	Name: to.Ptr("MyTestWatcher"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/MyTestAutomationAccount/watchers/MyTestWatcher"),
	// 	Properties: &armautomation.WatcherProperties{
	// 		Description: to.Ptr("This is a test watcher."),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-14T21:33:02.197Z"); return t}()),
	// 		ExecutionFrequencyInSeconds: to.Ptr[int64](600),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-14T22:29:57.650Z"); return t}()),
	// 		ScriptName: to.Ptr("MyTestWatcherRunbook"),
	// 		ScriptParameters: map[string]*string{
	// 		},
	// 		ScriptRunOn: to.Ptr("MyTestHybridWorkerGroup"),
	// 		Status: to.Ptr("New"),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// }
}
