package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/runbook/getRunbook.json
func ExampleRunbookClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRunbookClient().Get(ctx, "rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.RunbookClientGetResponse{
	// 	Runbook: armautomation.Runbook{
	// 		Name: to.Ptr("Get-AzureVMTutorial"),
	// 		Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Runbooks"),
	// 		Etag: to.Ptr("\"636263335437500000\""),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/Get-AzureVMTutorial"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armautomation.RunbookProperties{
	// 			Description: to.Ptr("Description of the Runbook"),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:23.75+00:00"); return t}()),
	// 			JobCount: to.Ptr[int32](0),
	// 			LastModifiedBy: to.Ptr("myEmaild@microsoft.com"),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T21:32:23.75+00:00"); return t}()),
	// 			LogActivityTrace: to.Ptr[int32](1),
	// 			LogProgress: to.Ptr(true),
	// 			LogVerbose: to.Ptr(false),
	// 			OutputTypes: []*string{
	// 			},
	// 			Parameters: map[string]*armautomation.RunbookParameter{
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShellWorkflow),
	// 			RuntimeEnvironment: to.Ptr("environmentName"),
	// 			State: to.Ptr(armautomation.RunbookStatePublished),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag01": to.Ptr("value01"),
	// 			"tag02": to.Ptr("value02"),
	// 		},
	// 	},
	// }
}
