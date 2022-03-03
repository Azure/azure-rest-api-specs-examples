Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_GetInstanceView_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	res, err := client.GetInstanceView(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientGetInstanceViewResult)
}
```
