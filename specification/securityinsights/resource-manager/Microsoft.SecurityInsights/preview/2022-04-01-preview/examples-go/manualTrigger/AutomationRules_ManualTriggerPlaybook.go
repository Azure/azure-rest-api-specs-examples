package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/manualTrigger/AutomationRules_ManualTriggerPlaybook.json
func ExampleIncidentsClient_RunPlaybook() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewIncidentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.RunPlaybook(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<incident-identifier>",
		&armsecurityinsights.IncidentsClientRunPlaybookOptions{RequestBody: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
