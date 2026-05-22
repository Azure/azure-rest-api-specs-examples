package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager/v2"
)

// Generated from example definition: 2024-04-01-preview/TrafficManagerUserMetricsKeys-PUT.json
func ExampleUserMetricsKeysClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewUserMetricsKeysClient().CreateOrUpdate(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
