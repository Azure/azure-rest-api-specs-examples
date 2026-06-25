package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_List.json
func ExampleVirtualMachineScaleSetVMExtensionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMExtensionsClient().List(ctx, "myResourceGroup", "myvmScaleSet", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetVMExtensionsClientListResponse{
	// 	VirtualMachineScaleSetVMExtensionsListResult: armcompute.VirtualMachineScaleSetVMExtensionsListResult{
	// 		Value: []*armcompute.VirtualMachineScaleSetVMExtension{
	// 			{
	// 				Properties: &armcompute.VirtualMachineExtensionProperties{
	// 					AutoUpgradeMinorVersion: to.Ptr(true),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					Publisher: to.Ptr("extPublisher"),
	// 					Type: to.Ptr("extType"),
	// 					TypeHandlerVersion: to.Ptr("1.2"),
	// 					Settings: map[string]any{
	// 						"UserName": "xyz@microsoft.com",
	// 					},
	// 				},
	// 				Name: to.Ptr("myVMExtension"),
	// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/extensions"),
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/extensions/myVMExtension"),
	// 			},
	// 			{
	// 				Properties: &armcompute.VirtualMachineExtensionProperties{
	// 					AutoUpgradeMinorVersion: to.Ptr(true),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					Publisher: to.Ptr("extPublisher1"),
	// 					Type: to.Ptr("extType1"),
	// 					TypeHandlerVersion: to.Ptr("1.0"),
	// 					Settings: map[string]any{
	// 						"UserName": "xyz@microsoft.com",
	// 					},
	// 				},
	// 				Name: to.Ptr("myVMExtension1"),
	// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/extensions"),
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/extensions/myVMExtension1"),
	// 			},
	// 		},
	// 	},
	// }
}
