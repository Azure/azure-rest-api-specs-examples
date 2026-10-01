package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: 2026-08-01/SpaceAllocation_Post.json
func ExampleCachesClient_BeginSpaceAllocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCachesClient().BeginSpaceAllocation(ctx, "scgroup", "sc1", &armstoragecache.CachesClientBeginSpaceAllocationOptions{
		SpaceAllocation: []*armstoragecache.StorageTargetSpaceAllocation{
			{
				Name:                 to.Ptr("st1"),
				AllocationPercentage: to.Ptr[int32](25),
			},
			{
				Name:                 to.Ptr("st2"),
				AllocationPercentage: to.Ptr[int32](50),
			},
			{
				Name:                 to.Ptr("st3"),
				AllocationPercentage: to.Ptr[int32](25),
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
