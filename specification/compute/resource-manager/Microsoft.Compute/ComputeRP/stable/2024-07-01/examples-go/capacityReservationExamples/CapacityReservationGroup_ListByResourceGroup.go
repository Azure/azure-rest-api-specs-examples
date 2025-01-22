package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservationGroup_ListByResourceGroup.json
func ExampleCapacityReservationGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacityReservationGroupsClient().NewListByResourceGroupPager("myResourceGroup", &armcompute.CapacityReservationGroupsClientListByResourceGroupOptions{Expand: to.Ptr(armcompute.ExpandTypesForGetCapacityReservationGroupsVirtualMachinesRef)})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CapacityReservationGroupListResult = armcompute.CapacityReservationGroupListResult{
		// 	Value: []*armcompute.CapacityReservationGroup{
		// 		{
		// 			Name: to.Ptr("{capacityReservationGroupName}"),
		// 			Type: to.Ptr("Microsoft.Compute/CapacityReservationGroups"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.CapacityReservationGroupProperties{
		// 				CapacityReservations: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation1"),
		// 					},
		// 					{
		// 						ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation2"),
		// 				}},
		// 				VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{capacityReservationGroupName}"),
		// 			Type: to.Ptr("Microsoft.Compute/CapacityReservationGroups"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.CapacityReservationGroupProperties{
		// 				CapacityReservations: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation3"),
		// 					},
		// 					{
		// 						ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation4"),
		// 				}},
		// 				VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
