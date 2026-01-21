package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/capacityReservationExamples/TargetedCapacityReservation_Get.json
func ExampleCapacityReservationsClient_Get_getATargetedCapacityReservation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationsClient().Get(ctx, "myResourceGroup", "targetedCapacityReservationGroup", "targetedCapacityReservation", &armcompute.CapacityReservationsClientGetOptions{Expand: to.Ptr(armcompute.CapacityReservationInstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapacityReservation = armcompute.CapacityReservation{
	// 	Name: to.Ptr("targetedCapacityReservation"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/targetedCapacityReservationGroup/capacityReservations/targetedCapacityReservation"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.CapacityReservationProperties{
	// 		InstanceView: &armcompute.CapacityReservationInstanceView{
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			}},
	// 			UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 				CurrentCapacity: to.Ptr[int32](5),
	// 				VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 				}},
	// 			},
	// 		},
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
	// 		ReservationID: to.Ptr("{GUID}"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
	// 		VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_DS1_v2"),
	// 		Capacity: to.Ptr[int64](4),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
