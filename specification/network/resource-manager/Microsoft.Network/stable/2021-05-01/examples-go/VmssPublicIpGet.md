Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssPublicIpGet.json
func ExamplePublicIPAddressesClient_GetVirtualMachineScaleSetPublicIPAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	res, err := client.GetVirtualMachineScaleSetPublicIPAddress(ctx,
		"<resource-group-name>",
		"<virtual-machine-scale-set-name>",
		"<virtualmachine-index>",
		"<network-interface-name>",
		"<ip-configuration-name>",
		"<public-ipaddress-name>",
		&armnetwork.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressResult)
}
```
