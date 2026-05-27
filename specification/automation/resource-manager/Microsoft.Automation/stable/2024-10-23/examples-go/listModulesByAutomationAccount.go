package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/listModulesByAutomationAccount.json
func ExampleModuleClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
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
		// page = armautomation.ModuleClientListByAutomationAccountResponse{
		// 	ModuleListResult: armautomation.ModuleListResult{
		// 		Value: []*armautomation.Module{
		// 			{
		// 				Name: to.Ptr("Azure"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Azure"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.ModuleProperties{
		// 					ActivityCount: to.Ptr[int32](0),
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2014-03-18T23:02:40.323+00:00"); return t}()),
		// 					IsGlobal: to.Ptr(true),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T16:57:48.343+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateCreated),
		// 					SizeInBytes: to.Ptr[int64](0),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Azure.Storage"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/Azure.Storage"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.ModuleProperties{
		// 					ActivityCount: to.Ptr[int32](0),
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-30T01:21:44.68+00:00"); return t}()),
		// 					IsGlobal: to.Ptr(true),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:04:27.833+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateCreated),
		// 					SizeInBytes: to.Ptr[int64](0),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AzureRM.Automation"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Automation"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.ModuleProperties{
		// 					ActivityCount: to.Ptr[int32](0),
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-30T01:19:39.427+00:00"); return t}()),
		// 					IsGlobal: to.Ptr(true),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:02:24.42+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateCreated),
		// 					SizeInBytes: to.Ptr[int64](0),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AzureRM.Compute"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Compute"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.ModuleProperties{
		// 					ActivityCount: to.Ptr[int32](0),
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-24T20:24:06.1+00:00"); return t}()),
		// 					IsGlobal: to.Ptr(true),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:01:53.81+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateCreated),
		// 					SizeInBytes: to.Ptr[int64](0),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AzureRM.Profile"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.ModuleProperties{
		// 					ActivityCount: to.Ptr[int32](0),
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-24T20:23:34.723+00:00"); return t}()),
		// 					IsGlobal: to.Ptr(true),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:01:22.993+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ModuleProvisioningStateCreated),
		// 					SizeInBytes: to.Ptr[int64](0),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
