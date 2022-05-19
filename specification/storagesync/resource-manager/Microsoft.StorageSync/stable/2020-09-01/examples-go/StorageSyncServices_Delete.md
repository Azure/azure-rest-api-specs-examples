Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragesync%2Farmstoragesync%2Fv1.0.0/sdk/resourcemanager/storagesync/armstoragesync/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/StorageSyncServices_Delete.json
func ExampleServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragesync.NewServicesClient("52b8da2f-61e0-4a1f-8dde-336911f367fb", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```
