Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.5.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/AddToSharingProfileInAGallery.json
func ExampleGallerySharingProfileClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGallerySharingProfileClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		armcompute.SharingUpdate{
			Groups: []*armcompute.SharingProfileGroup{
				{
					Type: armcompute.SharingProfileGroupTypes("Subscriptions").ToPtr(),
					IDs: []*string{
						to.StringPtr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
						to.StringPtr("380fd389-260b-41aa-bad9-0a83108c370b")},
				},
				{
					Type: armcompute.SharingProfileGroupTypes("AADTenants").ToPtr(),
					IDs: []*string{
						to.StringPtr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
				}},
			OperationType: armcompute.SharingUpdateOperationTypes("Add").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.GallerySharingProfileClientUpdateResult)
}
```
