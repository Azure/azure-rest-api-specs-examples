Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragesync%2Farmstoragesync%2Fv0.2.1/sdk/resourcemanager/storagesync/armstoragesync/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/SyncGroups_Create.json
func ExampleSyncGroupsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragesync.NewSyncGroupsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<storage-sync-service-name>",
		"<sync-group-name>",
		armstoragesync.SyncGroupCreateParameters{
			Properties: map[string]interface{}{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SyncGroupsClientCreateResult)
}
```
