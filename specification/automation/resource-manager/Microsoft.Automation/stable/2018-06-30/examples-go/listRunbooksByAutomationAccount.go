package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/listRunbooksByAutomationAccount.json
func ExampleRunbookClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRunbookClient().NewListByAutomationAccountPager("rg", "ContoseAutomationAccount", nil)
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
		// page.RunbookListResult = armautomation.RunbookListResult{
		// 	Value: []*armautomation.Runbook{
		// 		{
		// 			Name: to.Ptr("ASR-AddPublicIp"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Runbooks"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/ASR-AddPublicIp"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.RunbookProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:25.780Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:25.810Z"); return t}()),
		// 				LogActivityTrace: to.Ptr[int32](1),
		// 				LogProgress: to.Ptr(true),
		// 				LogVerbose: to.Ptr(true),
		// 				RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShell),
		// 				State: to.Ptr(armautomation.RunbookStatePublished),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AutoExport"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Runbooks"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/AutoExport"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.RunbookProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:27.327Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:27.327Z"); return t}()),
		// 				LogActivityTrace: to.Ptr[int32](1),
		// 				LogProgress: to.Ptr(true),
		// 				LogVerbose: to.Ptr(true),
		// 				RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShell),
		// 				State: to.Ptr(armautomation.RunbookStatePublished),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Get-AzureVMTutorial"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Runbooks"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/Get-AzureVMTutorial"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.RunbookProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:23.750Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:23.750Z"); return t}()),
		// 				LogActivityTrace: to.Ptr[int32](1),
		// 				LogProgress: to.Ptr(true),
		// 				LogVerbose: to.Ptr(false),
		// 				RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShellWorkflow),
		// 				State: to.Ptr(armautomation.RunbookStatePublished),
		// 			},
		// 	}},
		// }
	}
}
