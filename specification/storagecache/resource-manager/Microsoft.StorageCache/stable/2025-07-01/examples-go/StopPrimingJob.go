package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d18f451b13796bded0808d508c6297f2227271d5/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/StopPrimingJob.json
func ExampleCachesClient_BeginStopPrimingJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCachesClient().BeginStopPrimingJob(ctx, "scgroup", "sc1", &armstoragecache.CachesClientBeginStopPrimingJobOptions{PrimingJobID: &armstoragecache.PrimingJobIDParameter{
		PrimingJobID: to.Ptr("00000000000_0000000000"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
