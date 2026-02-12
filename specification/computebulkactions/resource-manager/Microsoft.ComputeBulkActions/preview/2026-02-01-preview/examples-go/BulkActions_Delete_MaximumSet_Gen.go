package armcomputebulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computebulkactions/armcomputebulkactions"
)

// Generated from example definition: 2026-02-01-preview/BulkActions_Delete_MaximumSet_Gen.json
func ExampleBulkActionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputebulkactions.NewClientFactory("50352BBD-59F1-4EE2-BA9C-A6E51B0B1B2B", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBulkActionsClient().BeginDelete(ctx, "rgcomputebulkactions", "eastus2euap", "3ec2ab23-9f13-4328-85c8-21928acbc7b8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
