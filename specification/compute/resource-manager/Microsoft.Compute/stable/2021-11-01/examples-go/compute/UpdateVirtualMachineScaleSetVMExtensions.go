package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/UpdateVirtualMachineScaleSetVMExtensions.json
func ExampleVirtualMachineScaleSetVMExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMExtensionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<vm-extension-name>",
		armcompute.VirtualMachineScaleSetVMExtensionUpdate{
			Properties: &armcompute.VirtualMachineExtensionUpdateProperties{
				Type:                    to.Ptr("<type>"),
				AutoUpgradeMinorVersion: to.Ptr(true),
				Publisher:               to.Ptr("<publisher>"),
				Settings: map[string]interface{}{
					"UserName": "xyz@microsoft.com",
				},
				TypeHandlerVersion: to.Ptr("<type-handler-version>"),
			},
		},
		&armcompute.VirtualMachineScaleSetVMExtensionsClientBeginUpdateOptions{ResumeToken: ""})
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
