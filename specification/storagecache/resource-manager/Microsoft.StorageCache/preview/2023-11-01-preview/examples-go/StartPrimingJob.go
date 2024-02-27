package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StartPrimingJob.json
func ExampleCachesClient_BeginStartPrimingJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCachesClient().BeginStartPrimingJob(ctx, "scgroup", "sc1", &armstoragecache.CachesClientBeginStartPrimingJobOptions{Primingjob: &armstoragecache.PrimingJob{
		PrimingJobName:     to.Ptr("contosoJob"),
		PrimingManifestURL: to.Ptr("https://contosostorage.blob.core.windows.net/contosoblob/00000000_00000000000000000000000000000000.00000000000.FFFFFFFF.00000000?sp=r&st=2021-08-11T19:33:35Z&se=2021-08-12T03:33:35Z&spr=https&sv=2020-08-04&sr=b&sig=<secret-value-from-key>"),
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
