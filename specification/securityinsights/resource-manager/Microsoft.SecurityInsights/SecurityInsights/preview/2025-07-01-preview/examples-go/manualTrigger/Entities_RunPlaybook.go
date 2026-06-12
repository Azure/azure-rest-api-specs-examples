package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/manualTrigger/Entities_RunPlaybook.json
func ExampleEntitiesClient_RunPlaybook() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewEntitiesClient().RunPlaybook(ctx, "myRg", "myWorkspace", "72e01a22-5cd2-4139-a149-9f2736ff2ar2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
