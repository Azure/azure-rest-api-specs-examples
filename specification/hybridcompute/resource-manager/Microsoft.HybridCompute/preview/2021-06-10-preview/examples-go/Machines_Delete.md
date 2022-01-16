Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv0.2.0/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Machines_Delete.json
func ExampleMachinesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<machine-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
