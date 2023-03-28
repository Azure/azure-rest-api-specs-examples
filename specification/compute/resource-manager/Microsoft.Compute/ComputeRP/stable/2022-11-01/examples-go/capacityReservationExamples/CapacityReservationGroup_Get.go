package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/19f98c9f526f8db961f172276dd6d6882a86ed86/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/capacityReservationExamples/CapacityReservationGroup_Get.json
func ExampleCapacityReservationGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationGroupsClient().Get(ctx, "myResourceGroup", "myCapacityReservationGroup", &armcompute.CapacityReservationGroupsClientGetOptions{Expand: to.Ptr(armcompute.CapacityReservationGroupInstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapacityReservationGroup = armcompute.CapacityReservationGroup{
	// 	Name: to.Ptr("myCapacityReservationGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"{tagName}": to.Ptr("{tagValue}"),
	// 	},
	// 	Properties: &armcompute.CapacityReservationGroupProperties{
	// 		CapacityReservations: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation1"),
	// 			},
	// 			{
	// 				ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation2"),
	// 		}},
	// 		InstanceView: &armcompute.CapacityReservationGroupInstanceView{
	// 			CapacityReservations: []*armcompute.CapacityReservationInstanceViewWithName{
	// 				{
	// 					Statuses: []*armcompute.InstanceViewStatus{
	// 						{
	// 							Code: to.Ptr("ProvisioningState/succeeded"),
	// 							DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					}},
	// 					UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 						CurrentCapacity: to.Ptr[int32](5),
	// 						VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 							{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 						}},
	// 					},
	// 					Name: to.Ptr("myCapacityReservation1"),
	// 				},
	// 				{
	// 					Statuses: []*armcompute.InstanceViewStatus{
	// 						{
	// 							Code: to.Ptr("ProvisioningState/succeeded"),
	// 							DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					}},
	// 					UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 						CurrentCapacity: to.Ptr[int32](5),
	// 						VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 							{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM4"),
	// 						}},
	// 					},
	// 					Name: to.Ptr("myCapacityReservation2"),
	// 			}},
	// 		},
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("3"),
	// 		to.Ptr("1")},
	// 	}
}
