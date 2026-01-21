package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Update.json
func ExampleVirtualMachineScaleSetVMExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMExtensionsClient().BeginUpdate(ctx, "myResourceGroup", "myvmScaleSet", "0", "myVMExtension", armcompute.VirtualMachineScaleSetVMExtensionUpdate{
		Properties: &armcompute.VirtualMachineExtensionUpdateProperties{
			Type:                    to.Ptr("extType"),
			AutoUpgradeMinorVersion: to.Ptr(true),
			Publisher:               to.Ptr("extPublisher"),
			Settings: map[string]any{
				"UserName": "xyz@microsoft.com",
			},
			TypeHandlerVersion: to.Ptr("1.2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSetVMExtension = armcompute.VirtualMachineScaleSetVMExtension{
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/extensions/myVMExtension"),
	// 	Name: to.Ptr("myVMExtension"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/extensions"),
	// 	Properties: &armcompute.VirtualMachineExtensionProperties{
	// 		Type: to.Ptr("extType"),
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Publisher: to.Ptr("extPublisher"),
	// 		Settings: map[string]any{
	// 			"UserName": "xyz@microsoft.com",
	// 		},
	// 		TypeHandlerVersion: to.Ptr("1.2"),
	// 	},
	// }
}
