Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/UpdateVMExtension.json
func ExampleVirtualMachineExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-name>",
		"<vm-extension-name>",
		armcompute.VirtualMachineExtensionUpdate{
			Properties: &armcompute.VirtualMachineExtensionUpdateProperties{
				Type:                    to.StringPtr("<type>"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				ProtectedSettingsFromKeyVault: map[string]interface{}{
					"secretUrl": "https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e",
					"sourceVault": map[string]interface{}{
						"id": "/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName",
					},
				},
				Publisher: to.StringPtr("<publisher>"),
				Settings: map[string]interface{}{
					"UserName": "xyz@microsoft.com",
				},
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
	log.Printf("Response result: %#v\n", res.VirtualMachineExtensionsClientUpdateResult)
}
```
