Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridnetwork%2Farmhybridnetwork%2Fv0.2.0/sdk/resourcemanager/hybridnetwork/armhybridnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorSkuPreviewGet.json
func ExampleVendorSKUPreviewClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridnetwork.NewVendorSKUPreviewClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<vendor-name>",
		"<sku-name>",
		"<preview-subscription>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VendorSKUPreviewClientGetResult)
}
```
