package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/manualTrigger/Incidents_RunPlaybook.json
func ExampleIncidentsClient_RunPlaybook() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewIncidentsClient().RunPlaybook(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ar4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
