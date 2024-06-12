package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/manualTrigger/Incidents_RunPlaybook.json
func ExampleIncidentsClient_RunPlaybook() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewIncidentsClient().RunPlaybook(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ar4", &armsecurityinsights.IncidentsClientRunPlaybookOptions{RequestBody: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
