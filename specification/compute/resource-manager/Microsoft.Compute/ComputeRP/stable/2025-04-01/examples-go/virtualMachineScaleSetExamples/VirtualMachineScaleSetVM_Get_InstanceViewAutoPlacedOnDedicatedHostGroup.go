package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_InstanceViewAutoPlacedOnDedicatedHostGroup.json
func ExampleVirtualMachineScaleSetVMsClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMsClient().GetInstanceView(ctx, "myResourceGroup", "myVirtualMachineScaleSet", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSetVMInstanceView = armcompute.VirtualMachineScaleSetVMInstanceView{
	// 	AssignedHost: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/hostGroups/myHostGroup/hosts/myHost"),
	// 	Disks: []*armcompute.DiskInstanceView{
	// 		{
	// 			Name: to.Ptr("myOSDisk"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T04:58:58.088Z"); return t}()),
	// 			}},
	// 	}},
	// 	PlatformFaultDomain: to.Ptr[int32](0),
	// 	PlatformUpdateDomain: to.Ptr[int32](0),
	// 	Statuses: []*armcompute.InstanceViewStatus{
	// 		{
	// 			Code: to.Ptr("ProvisioningState/succeeded"),
	// 			DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-05T04:59:58.185Z"); return t}()),
	// 		},
	// 		{
	// 			Code: to.Ptr("PowerState/running"),
	// 			DisplayStatus: to.Ptr("VM running"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 	}},
	// 	VMAgent: &armcompute.VirtualMachineAgentInstanceView{
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("ProvisioningState/Unavailable"),
	// 				DisplayStatus: to.Ptr("Not Ready"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesWarning),
	// 				Message: to.Ptr("VM status blob is found but not yet populated."),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T05:00:32.000Z"); return t}()),
	// 		}},
	// 		VMAgentVersion: to.Ptr("Unknown"),
	// 	},
	// }
}
