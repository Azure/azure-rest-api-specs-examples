Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/GetAGalleryImageVersionWithReplicationStatus.json
func ExampleGalleryImageVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		&armcompute.GalleryImageVersionsClientGetOptions{Expand: armcompute.ReplicationStatusTypes("ReplicationStatus").ToPtr()})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.GalleryImageVersionsClientGetResult)
}
```
