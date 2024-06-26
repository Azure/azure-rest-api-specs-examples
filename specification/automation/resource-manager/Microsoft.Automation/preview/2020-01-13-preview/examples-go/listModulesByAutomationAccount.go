package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listModulesByAutomationAccount.json
func ExampleModuleClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewModuleClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", nil)
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
		// 			Name: to.Ptr("Azure"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Azure"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2014-03-18T23:02:40.323Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T16:57:48.343Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Azure.Storage"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Azure.Storage"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-30T01:21:44.680Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:04:27.833Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureRM.Automation"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Automation"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-30T01:19:39.427Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:02:24.420Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureRM.Compute"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Compute"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-24T20:24:06.100Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:01:53.810Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureRM.Profile"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-24T20:23:34.723Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:01:22.993Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureRM.Resources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Resources"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-30T01:20:10.367Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:02:55.250Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureRM.Sql"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Sql"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-30T01:20:42.177Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:03:26.080Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureRM.Storage"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Storage"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-30T01:21:13.237Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:03:56.990Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerShell.Core"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Microsoft.PowerShell.Core"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-17T00:12:20.897Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T16:58:19.017Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerShell.Diagnostics"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Microsoft.PowerShell.Diagnostics"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-17T00:12:22.817Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T16:58:49.737Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerShell.Management"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Microsoft.PowerShell.Management"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-17T00:12:24.967Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T16:59:20.380Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerShell.Security"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Microsoft.PowerShell.Security"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-17T00:12:26.753Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T16:59:51.007Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerShell.Utility"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Microsoft.PowerShell.Utility"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-17T00:12:28.643Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:00:21.647Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.WSMan.Management"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Microsoft.WSMan.Management"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-17T00:17:15.003Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:00:52.197Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("OmsCompositeResources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/OmsCompositeResources"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T15:41:47.003Z"); return t}()),
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T15:42:10.567Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Orchestrator.AssetManagement.Cmdlets"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Orchestrator.AssetManagement.Cmdlets"),
		// 			Properties: &armautomation.ModuleProperties{
		// 				ActivityCount: to.Ptr[int32](0),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-12T00:45:12.897Z"); return t}()),
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:01.570Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateSucceeded),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 			},
		// 	}},
		// }
	}
}
