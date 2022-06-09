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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetExtensions_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetExtensionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<vmss-extension-name>",
		armcompute.VirtualMachineScaleSetExtensionUpdate{
			Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
				Type:                    to.Ptr("<type>"),
				AutoUpgradeMinorVersion: to.Ptr(true),
				EnableAutomaticUpgrade:  to.Ptr(true),
				ForceUpdateTag:          to.Ptr("<force-update-tag>"),
				ProtectedSettings:       map[string]interface{}{},
				ProvisionAfterExtensions: []*string{
					to.Ptr("aa")},
				Publisher:          to.Ptr("<publisher>"),
				Settings:           map[string]interface{}{},
				SuppressFailures:   to.Ptr(true),
				TypeHandlerVersion: to.Ptr("<type-handler-version>"),
			},
		},
		&armcompute.VirtualMachineScaleSetExtensionsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.7.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.
