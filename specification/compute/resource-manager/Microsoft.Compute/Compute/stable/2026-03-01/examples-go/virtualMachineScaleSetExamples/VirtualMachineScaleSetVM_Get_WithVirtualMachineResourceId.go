package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithVirtualMachineResourceId.json
func ExampleVirtualMachineScaleSetVMsClient_Get_getVMScaleSetFlexVMWithVirtualMachineResourceId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMsClient().Get(ctx, "myResourceGroup", "{vmss-flex-name}", "{vmss-flex-vm-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetVMsClientGetResponse{
	// 	VirtualMachineScaleSetVM: armcompute.VirtualMachineScaleSetVM{
	// 		Name: to.Ptr("{vmss-flex-vm-name}"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-flex-name}/virtualMachines/{vmss-flex-vm-name}"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		InstanceID: to.Ptr("{vmss-flex-vm-name}"),
	// 		Etag: to.Ptr("\"3\""),
	// 		Properties: &armcompute.VirtualMachineScaleSetVMProperties{
	// 			VirtualMachineResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vmss-flex-vm-name}"),
	// 		},
	// 	},
	// }
}
