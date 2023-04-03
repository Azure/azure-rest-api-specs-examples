package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/listAutomationAccountsByResourceGroup.json
func ExampleAccountClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountClient().NewListByResourceGroupPager("rg", nil)
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
		// page.AccountListResult = armautomation.AccountListResult{
		// 	Value: []*armautomation.Account{
		// 		{
		// 			Name: to.Ptr("myaccount"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myaccount"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-24T00:47:04.227+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.4+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccount123"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAccount123"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-29T00:32:32.52+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.4+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccountasfads"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAccountasfads"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:21:03.27+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:21:03.27+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount1"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:22:33.26+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:22:33.26+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount11"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:10:24.523+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:12.027+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount2"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:24:20.31+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:24:20.31+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount3"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount3"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:24:43.967+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:24:43.967+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount4"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount4"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:04:56.9+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:04:56.9+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount6"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount6"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:10:44.567+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:10:44.567+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount7"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount7"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:19:17.943+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:19:17.943+00:00"); return t}()),
		// 				State: to.Ptr(armautomation.AutomationAccountStateOk),
		// 			},
		// 	}},
		// }
	}
}
