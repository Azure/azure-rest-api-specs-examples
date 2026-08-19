package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_InstanceViewWithCapacityReservationType.json
func ExampleVirtualMachineScaleSetVMsClient_GetInstanceView_getInstanceViewOfAVirtualMachineFromAVMScaleSetThatIsEligibleForAndConsumingAnOpenCapacityReservation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
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
	// res = armcompute.VirtualMachineScaleSetVMsClientGetInstanceViewResponse{
	// 	VirtualMachineScaleSetVMInstanceView: armcompute.VirtualMachineScaleSetVMInstanceView{
	// 		PlatformUpdateDomain: to.Ptr[int32](0),
	// 		PlatformFaultDomain: to.Ptr[int32](0),
	// 		VMAgent: &armcompute.VirtualMachineAgentInstanceView{
	// 			VMAgentVersion: to.Ptr("2.7.41491.1010"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					DisplayStatus: to.Ptr("Ready"),
	// 					Message: to.Ptr("GuestAgent is running and processing the extensions."),
	// 					Time: to.Ptr(time.Date(2026, time.April, 1, 5, 0, 32, 0, time.UTC)),
	// 				},
	// 			},
	// 		},
	// 		Disks: []*armcompute.DiskInstanceView{
	// 			{
	// 				Name: to.Ptr("myOSDisk"),
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 						Time: to.Ptr(time.Date(2026, time.April, 1, 4, 58, 58, 88281500, time.UTC)),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		CapacityReservationType: to.Ptr(armcompute.CapacityReservationTypeOpen),
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("ProvisioningState/succeeded"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 				Time: to.Ptr(time.Date(2026, time.April, 1, 4, 59, 58, 185296600, time.UTC)),
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
