package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubAuthorizationRuleDelete.json
func ExampleEventHubsClient_DeleteAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewEventHubsClient().DeleteAuthorizationRule(ctx, "ArunMonocle", "sdk-Namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
