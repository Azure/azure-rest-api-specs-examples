package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPython2PackagesByAutomationAccount.json
func ExamplePython2PackageClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPython2PackageClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", nil)
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
		// page.ModuleListResult = armautomation.ModuleListResult{
		// 	Value: []*armautomation.Module{
		// 		{
		// 			Name: to.Ptr("configparser"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/python2Packages/configparser"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-17T17:57:01.786Z"); return t}()),
		// 				IsComposite: to.Ptr(false),
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-24T20:30:16.496Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](39573),
		// 				Version: to.Ptr("3.5.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("flask"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/python2Packages/flask"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-09T18:07:44.613Z"); return t}()),
		// 				IsComposite: to.Ptr(false),
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-09T18:09:20.773Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](91364),
		// 				Version: to.Ptr("1.0.2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("numpy"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/python2Packages/numpy"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-14T01:05:54.670Z"); return t}()),
		// 				IsComposite: to.Ptr(false),
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-17T17:58:09.873Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](13345673),
		// 				Version: to.Ptr("1.14.5"),
		// 			},
		// 	}},
		// }
	}
}
