package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription/v2"
)

// Generated from example definition: 2025-11-01-preview/acceptTargetDirectory.json
func ExampleSubscriptionsClient_AcceptTargetDirectory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSubscriptionsClient().AcceptTargetDirectory(ctx, "6c3c85bc-5366-4eaa-8055-a10529eafd03", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
