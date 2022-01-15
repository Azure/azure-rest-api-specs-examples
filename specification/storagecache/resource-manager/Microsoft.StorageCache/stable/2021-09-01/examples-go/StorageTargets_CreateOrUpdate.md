Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv0.2.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/StorageTargets_CreateOrUpdate.json
func ExampleStorageTargetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragecache.NewStorageTargetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cache-name>",
		"<storage-target-name>",
		&armstoragecache.StorageTargetsClientBeginCreateOrUpdateOptions{Storagetarget: &armstoragecache.StorageTarget{
			Properties: &armstoragecache.StorageTargetProperties{
				Junctions: []*armstoragecache.NamespaceJunction{
					{
						NamespacePath:   to.StringPtr("<namespace-path>"),
						NfsAccessPolicy: to.StringPtr("<nfs-access-policy>"),
						NfsExport:       to.StringPtr("<nfs-export>"),
						TargetPath:      to.StringPtr("<target-path>"),
					},
					{
						NamespacePath:   to.StringPtr("<namespace-path>"),
						NfsAccessPolicy: to.StringPtr("<nfs-access-policy>"),
						NfsExport:       to.StringPtr("<nfs-export>"),
						TargetPath:      to.StringPtr("<target-path>"),
					}},
				Nfs3: &armstoragecache.Nfs3Target{
					Target:     to.StringPtr("<target>"),
					UsageModel: to.StringPtr("<usage-model>"),
				},
				TargetType: armstoragecache.StorageTargetType("nfs3").ToPtr(),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.StorageTargetsClientCreateOrUpdateResult)
}
```
