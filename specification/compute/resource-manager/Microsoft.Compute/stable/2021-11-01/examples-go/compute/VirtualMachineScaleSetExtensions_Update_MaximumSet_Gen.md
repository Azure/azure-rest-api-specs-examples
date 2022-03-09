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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetExtensions_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<vmss-extension-name>",
		armcompute.VirtualMachineScaleSetExtensionUpdate{
			Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
				Type:                    to.StringPtr("<type>"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				EnableAutomaticUpgrade:  to.BoolPtr(true),
				ForceUpdateTag:          to.StringPtr("<force-update-tag>"),
				ProtectedSettings:       map[string]interface{}{},
				ProvisionAfterExtensions: []*string{
					to.StringPtr("aa")},
				Publisher:          to.StringPtr("<publisher>"),
				Settings:           map[string]interface{}{},
				SuppressFailures:   to.BoolPtr(true),
				TypeHandlerVersion: to.StringPtr("<type-handler-version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetExtensionsClientUpdateResult)
}
```
