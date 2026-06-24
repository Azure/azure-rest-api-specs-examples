package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/capacityReservationExamples/BlockCapacityReservation_ListByReservationGroup.json
func ExampleCapacityReservationsClient_NewListByCapacityReservationGroupPager_listBlockCapacityReservationsInReservationGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacityReservationsClient().NewListByCapacityReservationGroupPager("myResourceGroup", "blockCapacityReservationGroup", &armcompute.CapacityReservationsClientListByCapacityReservationGroupOptions{
		Expand: to.Ptr(armcompute.ExpandTypesForGetCapacityReservationGroupsVirtualMachinesRef)})
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
		// page = armcompute.CapacityReservationsClientListByCapacityReservationGroupResponse{
		// 	CapacityReservationListResult: armcompute.CapacityReservationListResult{
		// 		Value: []*armcompute.CapacityReservation{
		// 			{
		// 				Name: to.Ptr("{capacityReservationName}"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/CapacityReservation/{capacityReservationName}"),
		// 				Type: to.Ptr("Microsoft.Compute/CapacityReservations"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armcompute.CapacityReservationProperties{
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					ReservationID: to.Ptr("{GUID}"),
		// 					ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.3138469+00:00"); return t}()),
		// 					VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
		// 						},
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
		// 						},
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ScheduleProfile: &armcompute.ScheduleProfile{
		// 						Start: to.Ptr("2025-08-01T12:00:00Z"),
		// 						End: to.Ptr("2025-08-02T11:30:00Z"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("HR"),
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_ND96isr_H100_v5"),
		// 					Capacity: to.Ptr[int64](1),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("{capacityReservationName}"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/CapacityReservation/{capacityReservationName}"),
		// 				Type: to.Ptr("Microsoft.Compute/CapacityReservations"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armcompute.CapacityReservationProperties{
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					ReservationID: to.Ptr("{GUID}"),
		// 					ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.3138469+00:00"); return t}()),
		// 					VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM4"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ScheduleProfile: &armcompute.ScheduleProfile{
		// 						Start: to.Ptr("2025-08-01T12:00:00Z"),
		// 						End: to.Ptr("2025-08-02T11:30:00Z"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("HR"),
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_ND96isr_H100_v5"),
		// 					Capacity: to.Ptr[int64](1),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
