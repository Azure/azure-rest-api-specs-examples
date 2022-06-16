package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createOrUpdateRunbook.json
func ExampleRunbookClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewRunbookClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg",
		"ContoseAutomationAccount",
		"Get-AzureVMTutorial",
		armautomation.RunbookCreateOrUpdateParameters{
			Name:     to.Ptr("Get-AzureVMTutorial"),
			Location: to.Ptr("East US 2"),
			Properties: &armautomation.RunbookCreateOrUpdateProperties{
				Description:      to.Ptr("Description of the Runbook"),
				LogActivityTrace: to.Ptr[int32](1),
				LogProgress:      to.Ptr(true),
				LogVerbose:       to.Ptr(false),
				PublishContentLink: &armautomation.ContentLink{
					ContentHash: &armautomation.ContentHash{
						Algorithm: to.Ptr("SHA256"),
						Value:     to.Ptr("115775B8FF2BE672D8A946BD0B489918C724DDE15A440373CA54461D53010A80"),
					},
					URI: to.Ptr("https://raw.githubusercontent.com/Azure/azure-quickstart-templates/master/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1"),
				},
				RunbookType: to.Ptr(armautomation.RunbookTypeEnumPowerShellWorkflow),
			},
			Tags: map[string]*string{
				"tag01": to.Ptr("value01"),
				"tag02": to.Ptr("value02"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
