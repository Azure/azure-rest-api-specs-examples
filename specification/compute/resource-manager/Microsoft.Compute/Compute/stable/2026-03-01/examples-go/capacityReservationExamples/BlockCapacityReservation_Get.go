package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/capacityReservationExamples/BlockCapacityReservation_Get.json
func ExampleCapacityReservationsClient_Get_getABlockCapacityReservation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationsClient().Get(ctx, "myResourceGroup", "blockCapacityReservationGroup", "blockCapacityReservation", &armcompute.CapacityReservationsClientGetOptions{
		Expand: to.Ptr(armcompute.CapacityReservationInstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CapacityReservationsClientGetResponse{
	// 	CapacityReservation: armcompute.CapacityReservation{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/blockCapacityReservationGroup/capacityReservations/blockCapacityReservation"),
	// 		Properties: &armcompute.CapacityReservationProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](3),
	// 			ReservationID: to.Ptr("{GUID}"),
	// 			ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.3138469+00:00"); return t}()),
	// 			VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			InstanceView: &armcompute.CapacityReservationInstanceView{
	// 				UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 					CurrentCapacity: to.Ptr[int32](5),
	// 					VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 						{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 						},
	// 						{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 						},
	// 					},
	// 				},
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					},
	// 				},
	// 			},
	// 			ScheduleProfile: &armcompute.ScheduleProfile{
	// 				Start: to.Ptr("2025-08-01T12:00:00Z"),
	// 				End: to.Ptr("2025-08-02T11:30:00Z"),
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.3138469+00:00"); return t}()),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("HR"),
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Standard_ND96isr_H100_v5"),
	// 			Capacity: to.Ptr[int64](1),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 		Name: to.Ptr("blockCapacityReservation"),
	// 	},
	// }
}
