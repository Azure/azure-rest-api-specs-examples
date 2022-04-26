Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv0.4.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragecache_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/StorageTargets_CreateOrUpdate.json
func ExampleStorageTargetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstoragecache.NewStorageTargetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cache-name>",
		"<storage-target-name>",
		&armstoragecache.StorageTargetsClientBeginCreateOrUpdateOptions{Storagetarget: &armstoragecache.StorageTarget{
			Properties: &armstoragecache.StorageTargetProperties{
				Junctions: []*armstoragecache.NamespaceJunction{
					{
						NamespacePath:   to.Ptr("<namespace-path>"),
						NfsAccessPolicy: to.Ptr("<nfs-access-policy>"),
						NfsExport:       to.Ptr("<nfs-export>"),
						TargetPath:      to.Ptr("<target-path>"),
					},
					{
						NamespacePath:   to.Ptr("<namespace-path>"),
						NfsAccessPolicy: to.Ptr("<nfs-access-policy>"),
						NfsExport:       to.Ptr("<nfs-export>"),
						TargetPath:      to.Ptr("<target-path>"),
					}},
				Nfs3: &armstoragecache.Nfs3Target{
					Target:     to.Ptr("<target>"),
					UsageModel: to.Ptr("<usage-model>"),
				},
				TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
			},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
