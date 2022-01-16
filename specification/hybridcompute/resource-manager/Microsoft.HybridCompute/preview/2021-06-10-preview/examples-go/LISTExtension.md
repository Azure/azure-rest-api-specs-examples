Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv0.2.0/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/LISTExtension.json
func ExampleMachineExtensionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachineExtensionsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<machine-name>",
		&armhybridcompute.MachineExtensionsClientListOptions{Expand: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
