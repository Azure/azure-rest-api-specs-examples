package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/capacityReservationExamples/OpenCapacityReservation_Get.json
func ExampleCapacityReservationsClient_Get_getAnOpenCapacityReservationWithInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationsClient().Get(ctx, "myResourceGroup", "openCapacityReservationGroup", "openCapacityReservation", &armcompute.CapacityReservationsClientGetOptions{
		Expand: to.Ptr(armcompute.CapacityReservationInstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CapacityReservationsClientGetResponse{
	// 	CapacityReservation: armcompute.CapacityReservation{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/openCapacityReservationGroup/capacityReservations/openCapacityReservation"),
	// 		Properties: &armcompute.CapacityReservationProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](3),
	// 			ReservationID: to.Ptr("{GUID}"),
	// 			ProvisioningTime: to.Ptr(time.Date(2021, time.June, 27, 1, 2, 38, 313846900, time.UTC)),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			InstanceView: &armcompute.CapacityReservationInstanceView{
	// 				UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 					CurrentCapacity: to.Ptr[int32](10),
	// 					UsedReservedCountBySubscription: map[string]*int32{
	// 						"{subscription-id1}": to.Ptr[int32](3),
	// 						"{subscription-id2}": to.Ptr[int32](2),
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
	// 			TimeCreated: to.Ptr(time.Date(2021, time.June, 27, 1, 2, 38, 313846900, time.UTC)),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"{tagName}": to.Ptr("{tagValue}"),
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Standard_DS1_v2"),
	// 			Capacity: to.Ptr[int64](10),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 		Name: to.Ptr("openCapacityReservation"),
	// 	},
	// }
}
