package armprogrammableconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/programmableconnectivity/armprogrammableconnectivity"
)

// Generated from example definition: 2024-01-15-preview/OperatorApiConnections_Delete_MinimumSet_Gen.json
func ExampleOperatorAPIConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprogrammableconnectivity.NewClientFactory("B976474B-99FA-4C25-A3BD-8B05C3C3D07A", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOperatorAPIConnectionsClient().BeginDelete(ctx, "rgopenapi", "dawr", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
