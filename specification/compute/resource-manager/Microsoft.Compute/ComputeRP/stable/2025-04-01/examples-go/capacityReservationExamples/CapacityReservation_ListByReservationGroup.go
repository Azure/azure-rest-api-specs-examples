package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/capacityReservationExamples/CapacityReservation_ListByReservationGroup.json
func ExampleCapacityReservationsClient_NewListByCapacityReservationGroupPager_listCapacityReservationsInReservationGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacityReservationsClient().NewListByCapacityReservationGroupPager("myResourceGroup", "myCapacityReservationGroup", nil)
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
		// page.CapacityReservationListResult = armcompute.CapacityReservationListResult{
		// 	Value: []*armcompute.CapacityReservation{
		// 		{
		// 			Name: to.Ptr("{capacityReservationName}"),
		// 			Type: to.Ptr("Microsoft.Compute/CapacityReservations"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/CapacityReservation/{capacityReservationName}"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("HR"),
		// 			},
		// 			Properties: &armcompute.CapacityReservationProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
		// 				ReservationID: to.Ptr("{GUID}"),
		// 				VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Standard_DS1_v2"),
		// 				Capacity: to.Ptr[int64](4),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1")},
		// 			},
		// 			{
		// 				Name: to.Ptr("{capacityReservationName}"),
		// 				Type: to.Ptr("Microsoft.Compute/CapacityReservations"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/CapacityReservation/{capacityReservationName}"),
		// 				Location: to.Ptr("West US"),
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("HR"),
		// 				},
		// 				Properties: &armcompute.CapacityReservationProperties{
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
		// 					ReservationID: to.Ptr("{GUID}"),
		// 					VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM4"),
		// 					}},
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_A1_v2"),
		// 					Capacity: to.Ptr[int64](4),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1")},
		// 			}},
		// 		}
	}
}
