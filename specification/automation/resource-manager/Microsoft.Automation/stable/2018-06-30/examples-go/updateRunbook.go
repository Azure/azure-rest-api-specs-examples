package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/updateRunbook.json
func ExampleRunbookClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRunbookClient().Update(ctx, "rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", armautomation.RunbookUpdateParameters{
		Properties: &armautomation.RunbookUpdateProperties{
			Description:      to.Ptr("Updated Description of the Runbook"),
			LogActivityTrace: to.Ptr[int32](1),
			LogProgress:      to.Ptr(true),
			LogVerbose:       to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
