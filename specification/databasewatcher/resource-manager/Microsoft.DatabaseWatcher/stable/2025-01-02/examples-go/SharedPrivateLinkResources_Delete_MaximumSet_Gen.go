package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/SharedPrivateLinkResources_Delete_MaximumSet_Gen.json
func ExampleSharedPrivateLinkResourcesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSharedPrivateLinkResourcesClient().BeginDelete(ctx, "apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
