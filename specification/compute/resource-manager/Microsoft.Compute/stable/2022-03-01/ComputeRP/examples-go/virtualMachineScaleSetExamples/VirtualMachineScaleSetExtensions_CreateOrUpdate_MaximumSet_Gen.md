```go
package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetExtensionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rgcompute",
		"aaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaa",
		armcompute.VirtualMachineScaleSetExtension{
			Name: to.Ptr("{extension-name}"),
			Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
				Type:                    to.Ptr("{extension-Type}"),
				AutoUpgradeMinorVersion: to.Ptr(true),
				EnableAutomaticUpgrade:  to.Ptr(true),
				ForceUpdateTag:          to.Ptr("aaaaaaaaa"),
				ProtectedSettings:       map[string]interface{}{},
				ProvisionAfterExtensions: []*string{
					to.Ptr("aa")},
				Publisher:          to.Ptr("{extension-Publisher}"),
				Settings:           map[string]interface{}{},
				SuppressFailures:   to.Ptr(true),
				TypeHandlerVersion: to.Ptr("{handler-version}"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv2.0.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.
