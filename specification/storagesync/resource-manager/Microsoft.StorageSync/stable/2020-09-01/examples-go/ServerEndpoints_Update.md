Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragesync%2Farmstoragesync%2Fv0.2.0/sdk/resourcemanager/storagesync/armstoragesync/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragesync_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Update.json
func ExampleServerEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragesync.NewServerEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<storage-sync-service-name>",
		"<sync-group-name>",
		"<server-endpoint-name>",
		&armstoragesync.ServerEndpointsClientBeginUpdateOptions{Parameters: &armstoragesync.ServerEndpointUpdateParameters{
			Properties: &armstoragesync.ServerEndpointUpdateProperties{
				CloudTiering:           armstoragesync.FeatureStatus("off").ToPtr(),
				LocalCacheMode:         armstoragesync.LocalCacheMode("UpdateLocallyCachedFiles").ToPtr(),
				OfflineDataTransfer:    armstoragesync.FeatureStatus("off").ToPtr(),
				TierFilesOlderThanDays: to.Int32Ptr(0),
				VolumeFreeSpacePercent: to.Int32Ptr(100),
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
	log.Printf("Response result: %#v\n", res.ServerEndpointsClientUpdateResult)
}
```
