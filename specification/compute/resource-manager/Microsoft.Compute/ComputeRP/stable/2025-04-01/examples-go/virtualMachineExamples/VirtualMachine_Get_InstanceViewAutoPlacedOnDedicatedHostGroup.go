package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachine_Get_InstanceViewAutoPlacedOnDedicatedHostGroup.json
func ExampleVirtualMachinesClient_InstanceView_getInstanceViewOfAVirtualMachinePlacedOnADedicatedHostGroupThroughAutomaticPlacement() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().InstanceView(ctx, "myResourceGroup", "myVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineInstanceView = armcompute.VirtualMachineInstanceView{
	// 	AssignedHost: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/hostGroups/myHostGroup/hosts/myHost"),
	// 	ComputerName: to.Ptr("myVM"),
	// 	Disks: []*armcompute.DiskInstanceView{
	// 		{
	// 			Name: to.Ptr("myOsDisk"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T21:29:47.477Z"); return t}()),
	// 			}},
	// 	}},
	// 	HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypeV1),
	// 	OSName: to.Ptr("Windows Server 2016 Datacenter"),
	// 	OSVersion: to.Ptr("Microsoft Windows NT 10.0.14393.0"),
	// 	Statuses: []*armcompute.InstanceViewStatus{
	// 		{
	// 			Code: to.Ptr("ProvisioningState/succeeded"),
	// 			DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T21:30:12.805Z"); return t}()),
	// 		},
	// 		{
	// 			Code: to.Ptr("PowerState/running"),
	// 			DisplayStatus: to.Ptr("VM running"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 	}},
	// 	VMAgent: &armcompute.VirtualMachineAgentInstanceView{
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("ProvisioningState/succeeded"),
	// 				DisplayStatus: to.Ptr("Ready"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Message: to.Ptr("GuestAgent is running and accepting new configurations."),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-01T23:11:22.000Z"); return t}()),
	// 		}},
	// 		VMAgentVersion: to.Ptr("2.7.41491.949"),
	// 	},
	// }
}
