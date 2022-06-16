package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/manualTrigger/AutomationRules_ManualTriggerPlaybook.json
func ExampleAutomationRulesClient_ManualTriggerPlaybook() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewAutomationRulesClient("<subscription-id>", cred, nil)
	_, err = client.ManualTriggerPlaybook(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<incident-identifier>",
		&armsecurityinsights.AutomationRulesClientManualTriggerPlaybookOptions{RequestBody: nil})
	if err != nil {
		log.Fatal(err)
	}
}
