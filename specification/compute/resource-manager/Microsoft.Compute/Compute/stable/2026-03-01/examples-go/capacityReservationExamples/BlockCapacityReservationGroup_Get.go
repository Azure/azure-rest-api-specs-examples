package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/capacityReservationExamples/BlockCapacityReservationGroup_Get.json
func ExampleCapacityReservationGroupsClient_Get_getABlockCapacityReservationGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationGroupsClient().Get(ctx, "myResourceGroup", "blockCapacityReservationGroup", &armcompute.CapacityReservationGroupsClientGetOptions{
		Expand: to.Ptr(armcompute.CapacityReservationGroupInstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CapacityReservationGroupsClientGetResponse{
	// 	CapacityReservationGroup: armcompute.CapacityReservationGroup{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/CapacityReservationGroups/blockCapacityReservationGroup"),
	// 		Properties: &armcompute.CapacityReservationGroupProperties{
	// 			CapacityReservations: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/blockCapacityReservationGroup/capacityReservations/blockCapacityReservation1"),
	// 				},
	// 				{
	// 					ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/blockCapacityReservationGroup/capacityReservations/blockCapacityReservation2"),
	// 				},
	// 			},
	// 			InstanceView: &armcompute.CapacityReservationGroupInstanceView{
	// 				CapacityReservations: []*armcompute.CapacityReservationInstanceViewWithName{
	// 					{
	// 						Name: to.Ptr("blockCapacityReservation1"),
	// 						UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 							CurrentCapacity: to.Ptr[int32](5),
	// 							VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 								{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 								},
	// 								{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 								},
	// 							},
	// 						},
	// 						Statuses: []*armcompute.InstanceViewStatus{
	// 							{
	// 								Code: to.Ptr("ProvisioningState/succeeded"),
	// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 								DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("blockCapacityReservation2"),
	// 						UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 							CurrentCapacity: to.Ptr[int32](5),
	// 							VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 								{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
	// 								},
	// 								{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM4"),
	// 								},
	// 							},
	// 						},
	// 						Statuses: []*armcompute.InstanceViewStatus{
	// 							{
	// 								Code: to.Ptr("ProvisioningState/succeeded"),
	// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 								DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ReservationType: to.Ptr(armcompute.ReservationTypeBlock),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"{tagName}": to.Ptr("{tagValue}"),
	// 		},
	// 		Name: to.Ptr("blockCapacityReservationGroup"),
	// 		Zones: []*string{
	// 			to.Ptr("3"),
	// 			to.Ptr("1"),
	// 		},
	// 	},
	// }
}
