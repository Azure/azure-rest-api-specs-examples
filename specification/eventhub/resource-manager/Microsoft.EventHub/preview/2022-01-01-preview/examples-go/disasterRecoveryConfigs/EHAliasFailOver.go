package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/disasterRecoveryConfigs/EHAliasFailOver.json
func ExampleDisasterRecoveryConfigsClient_FailOver() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventhub.NewDisasterRecoveryConfigsClient("exampleSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.FailOver(ctx,
		"exampleResourceGroup",
		"sdk-Namespace-8859",
		"sdk-DisasterRecovery-3814",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
