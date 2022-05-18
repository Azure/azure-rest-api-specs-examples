Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv1.0.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/StorageTargets_CreateOrUpdate.json
func ExampleStorageTargetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragecache.NewStorageTargetsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"scgroup",
		"sc1",
		"st1",
		&armstoragecache.StorageTargetsClientBeginCreateOrUpdateOptions{Storagetarget: &armstoragecache.StorageTarget{
			Properties: &armstoragecache.StorageTargetProperties{
				Junctions: []*armstoragecache.NamespaceJunction{
					{
						NamespacePath:   to.Ptr("/path/on/cache"),
						NfsAccessPolicy: to.Ptr("default"),
						NfsExport:       to.Ptr("exp1"),
						TargetPath:      to.Ptr("/path/on/exp1"),
					},
					{
						NamespacePath:   to.Ptr("/path2/on/cache"),
						NfsAccessPolicy: to.Ptr("rootSquash"),
						NfsExport:       to.Ptr("exp2"),
						TargetPath:      to.Ptr("/path2/on/exp2"),
					}},
				Nfs3: &armstoragecache.Nfs3Target{
					Target:     to.Ptr("10.0.44.44"),
					UsageModel: to.Ptr("READ_HEAVY_INFREQ"),
				},
				TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
