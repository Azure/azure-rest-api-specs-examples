package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/TrafficManagerUserMetricsKeys-DELETE.json
func ExampleUserMetricsKeysClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewUserMetricsKeysClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Delete(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
