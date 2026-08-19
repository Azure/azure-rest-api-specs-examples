package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/virtualMachineExamples/VirtualMachine_Get_InstanceViewAutoPlacedOnDedicatedHostGroup.json
func ExampleVirtualMachinesClient_InstanceView_getInstanceViewOfAVirtualMachinePlacedOnADedicatedHostGroupThroughAutomaticPlacement() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
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
	// res = armcompute.VirtualMachinesClientInstanceViewResponse{
	// 	VirtualMachineInstanceView: armcompute.VirtualMachineInstanceView{
	// 		ComputerName: to.Ptr("myVM"),
	// 		OSName: to.Ptr("Windows Server 2016 Datacenter"),
	// 		OSVersion: to.Ptr("Microsoft Windows NT 10.0.14393.0"),
	// 		VMAgent: &armcompute.VirtualMachineAgentInstanceView{
	// 			VMAgentVersion: to.Ptr("2.7.41491.949"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					DisplayStatus: to.Ptr("Ready"),
	// 					Message: to.Ptr("GuestAgent is running and accepting new configurations."),
	// 					Time: to.Ptr(time.Date(2026, time.April, 1, 23, 11, 22, 0, time.UTC)),
	// 				},
	// 			},
	// 		},
	// 		Disks: []*armcompute.DiskInstanceView{
	// 			{
	// 				Name: to.Ptr("myOsDisk"),
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 						Time: to.Ptr(time.Date(2026, time.April, 1, 21, 29, 47, 477089000, time.UTC)),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypeV1),
	// 		AssignedHost: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/hostGroups/myHostGroup/hosts/myHost"),
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("ProvisioningState/succeeded"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 				Time: to.Ptr(time.Date(2026, time.April, 1, 21, 30, 12, 805191700, time.UTC)),
	// 			},
	// 			{
	// 				Code: to.Ptr("PowerState/running"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				DisplayStatus: to.Ptr("VM running"),
	// 			},
	// 		},
	// 	},
	// }
}
