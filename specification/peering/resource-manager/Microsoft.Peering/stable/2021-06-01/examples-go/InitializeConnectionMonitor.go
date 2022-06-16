package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/InitializeConnectionMonitor.json
func ExampleServicesClient_InitializeConnectionMonitor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewServicesClient("<subscription-id>", cred, nil)
	_, err = client.InitializeConnectionMonitor(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
