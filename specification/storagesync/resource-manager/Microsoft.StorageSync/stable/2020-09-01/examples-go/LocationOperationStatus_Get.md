Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragesync%2Farmstoragesync%2Fv0.1.0/sdk/resourcemanager/storagesync/armstoragesync/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/LocationOperationStatus_Get.json
func ExampleMicrosoftStorageSyncClient_LocationOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragesync.NewMicrosoftStorageSyncClient("<subscription-id>", cred, nil)
	res, err := client.LocationOperationStatus(ctx,
		"<location-name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("LocationOperationStatus.ID: %s\n", *res.ID)
}
```
