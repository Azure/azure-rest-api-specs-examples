package armdellstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dell/armdellstorage"
)

// Generated from example definition: 2025-03-21/FileSystems_Delete_MaximumSet_Gen.json
func ExampleFileSystemsClient_BeginDelete_fileSystemsDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdellstorage.NewClientFactory("4B6E265D-57CF-4A9D-8B35-3CC68ED9D208", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSystemsClient().BeginDelete(ctx, "rgDell", "abcd", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
