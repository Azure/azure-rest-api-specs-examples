Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragesync%2Farmstoragesync%2Fv0.2.1/sdk/resourcemanager/storagesync/armstoragesync/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/CloudEndpoints_PostRestore.json
func ExampleCloudEndpointsClient_BeginPostRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragesync.NewCloudEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPostRestore(ctx,
		"<resource-group-name>",
		"<storage-sync-service-name>",
		"<sync-group-name>",
		"<cloud-endpoint-name>",
		armstoragesync.PostRestoreRequest{
			AzureFileShareURI: to.StringPtr("<azure-file-share-uri>"),
			RestoreFileSpec: []*armstoragesync.RestoreFileSpec{
				{
					Path:  to.StringPtr("<path>"),
					Isdir: to.BoolPtr(false),
				},
				{
					Path:  to.StringPtr("<path>"),
					Isdir: to.BoolPtr(true),
				},
				{
					Path:  to.StringPtr("<path>"),
					Isdir: to.BoolPtr(false),
				},
				{
					Path:  to.StringPtr("<path>"),
					Isdir: to.BoolPtr(false),
				}},
			SourceAzureFileShareURI: to.StringPtr("<source-azure-file-share-uri>"),
			Status:                  to.StringPtr("<status>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
