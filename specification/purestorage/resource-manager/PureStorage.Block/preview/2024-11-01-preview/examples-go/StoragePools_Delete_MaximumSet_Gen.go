package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01-preview/StoragePools_Delete_MaximumSet_Gen.json
func ExampleStoragePoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("BC47D6CC-AA80-4374-86F8-19D94EC70666", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStoragePoolsClient().BeginDelete(ctx, "rgpurestorage", "storagePoolname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
