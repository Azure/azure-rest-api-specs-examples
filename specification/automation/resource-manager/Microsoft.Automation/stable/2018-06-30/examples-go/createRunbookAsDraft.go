package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createRunbookAsDraft.json
func ExampleRunbookClient_CreateOrUpdate_createRunbookAsDraft() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRunbookClient().CreateOrUpdate(ctx, "rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", armautomation.RunbookCreateOrUpdateParameters{
		Name:     to.Ptr("Get-AzureVMTutorial"),
		Location: to.Ptr("East US 2"),
		Properties: &armautomation.RunbookCreateOrUpdateProperties{
			Description: to.Ptr("Description of the Runbook"),
			Draft:       &armautomation.RunbookDraft{},
			LogProgress: to.Ptr(false),
			LogVerbose:  to.Ptr(false),
			RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShellWorkflow),
		},
		Tags: map[string]*string{
			"tag01": to.Ptr("value01"),
			"tag02": to.Ptr("value02"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Runbook = armautomation.Runbook{
	// 	Name: to.Ptr("Get-AzureVMTutorial"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Runbooks"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/Get-AzureVMTutorial"),
	// 	Location: to.Ptr("East US 2"),
	// 	Tags: map[string]*string{
	// 		"tag01": to.Ptr("value01"),
	// 		"tag02": to.Ptr("value02"),
	// 	},
	// 	Etag: to.Ptr("\"636263318866000000\""),
	// 	Properties: &armautomation.RunbookProperties{
	// 		Description: to.Ptr("Description of the Runbook"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-09T03:25:59.097Z"); return t}()),
	// 		JobCount: to.Ptr[int32](0),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-09T03:25:59.097Z"); return t}()),
	// 		LogActivityTrace: to.Ptr[int32](0),
	// 		LogProgress: to.Ptr(false),
	// 		LogVerbose: to.Ptr(false),
	// 		Parameters: map[string]*armautomation.RunbookParameter{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunbookType: to.Ptr(armautomation.RunbookTypeEnumScript),
	// 		State: to.Ptr(armautomation.RunbookStateNew),
	// 	},
	// }
}
