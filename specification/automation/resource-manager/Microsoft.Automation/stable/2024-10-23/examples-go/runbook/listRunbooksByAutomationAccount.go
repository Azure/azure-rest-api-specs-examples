package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/runbook/listRunbooksByAutomationAccount.json
func ExampleRunbookClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
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
		// page = armautomation.RunbookClientListByAutomationAccountResponse{
		// 	RunbookListResult: armautomation.RunbookListResult{
		// 		Value: []*armautomation.Runbook{
		// 			{
		// 				Name: to.Ptr("ASR-AddPublicIp"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Runbooks"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/ASR-AddPublicIp"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.RunbookProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:25.78+00:00"); return t}()),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:25.81+00:00"); return t}()),
		// 					LogActivityTrace: to.Ptr[int32](1),
		// 					LogProgress: to.Ptr(true),
		// 					LogVerbose: to.Ptr(true),
		// 					RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShellWorkflow),
		// 					RuntimeEnvironment: to.Ptr("environmentName"),
		// 					State: to.Ptr(armautomation.RunbookStatePublished),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AutoExport"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Runbooks"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/AutoExport"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.RunbookProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:27.327+00:00"); return t}()),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:27.327+00:00"); return t}()),
		// 					LogActivityTrace: to.Ptr[int32](1),
		// 					LogProgress: to.Ptr(true),
		// 					LogVerbose: to.Ptr(true),
		// 					RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShellWorkflow),
		// 					RuntimeEnvironment: to.Ptr("environmentName"),
		// 					State: to.Ptr(armautomation.RunbookStatePublished),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Get-AzureVMTutorial"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Runbooks"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/Get-AzureVMTutorial"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.RunbookProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:23.75+00:00"); return t}()),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:23.75+00:00"); return t}()),
		// 					LogActivityTrace: to.Ptr[int32](1),
		// 					LogProgress: to.Ptr(true),
		// 					LogVerbose: to.Ptr(false),
		// 					RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShellWorkflow),
		// 					RuntimeEnvironment: to.Ptr("environmentName"),
		// 					State: to.Ptr(armautomation.RunbookStatePublished),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
