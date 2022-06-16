package armstoragecache_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"
)

// x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/StorageTargets_DnsRefresh.json
func ExampleStorageTargetsClient_BeginDNSRefresh() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragecache.NewStorageTargetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDNSRefresh(ctx,
		"<resource-group-name>",
		"<cache-name>",
		"<storage-target-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
