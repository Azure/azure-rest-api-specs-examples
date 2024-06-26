package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e2749bb2cbee0b4c447a9d6c1d7cbce3d415abd4/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-03-01-preview/examples/SpaceAllocation_Post.json
func ExampleCachesClient_BeginSpaceAllocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCachesClient().BeginSpaceAllocation(ctx, "scgroup", "sc1", &armstoragecache.CachesClientBeginSpaceAllocationOptions{SpaceAllocation: []*armstoragecache.StorageTargetSpaceAllocation{
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
		}},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
