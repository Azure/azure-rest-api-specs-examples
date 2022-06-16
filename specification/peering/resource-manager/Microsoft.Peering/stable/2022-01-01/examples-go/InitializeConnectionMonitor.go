package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/InitializeConnectionMonitor.json
func ExampleServicesClient_InitializeConnectionMonitor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpeering.NewServicesClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.InitializeConnectionMonitor(ctx,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
