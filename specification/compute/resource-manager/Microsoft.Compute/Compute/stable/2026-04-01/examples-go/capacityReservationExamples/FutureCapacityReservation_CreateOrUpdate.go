package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/capacityReservationExamples/FutureCapacityReservation_CreateOrUpdate.json
func ExampleCapacityReservationsClient_BeginCreateOrUpdate_createOrUpdateAFutureCapacityReservation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "futureCapacityReservationGroup", "futureCapacityReservation", armcompute.CapacityReservation{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("Standard_DS1_v2"),
			Capacity: to.Ptr[int64](4),
		},
		Properties: &armcompute.CapacityReservationProperties{
			ScheduleProfile: &armcompute.ScheduleProfile{
				Start:                 to.Ptr("2026-08-01T12:00:00Z"),
				MinimumCommitmentDays: to.Ptr[int32](30),
			},
		},
		Zones: []*string{
			to.Ptr("1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CapacityReservationsClientCreateOrUpdateResponse{
	// 	CapacityReservation: armcompute.CapacityReservation{
	// 		Name: to.Ptr("futureCapacityReservation"),
	// 		Location: to.Ptr("westus"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/futureCapacityReservationGroup/capacityReservations/futureCapacityReservation"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("HR"),
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Standard_DS1_v2"),
	// 			Capacity: to.Ptr[int64](4),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 		Properties: &armcompute.CapacityReservationProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](3),
	// 			ReservationID: to.Ptr("{GUID}"),
	// 			ProvisioningState: to.Ptr("Creating"),
	// 			ScheduleProfile: &armcompute.ScheduleProfile{
	// 				Start: to.Ptr("2026-08-01T12:00:00Z"),
	// 				MinimumCommitmentDays: to.Ptr[int32](30),
	// 				ModifiableUntil: to.Ptr(time.Date(2026, time.June, 6, 12, 0, 0, 0, time.UTC)),
	// 			},
	// 			ProvisioningTime: to.Ptr(time.Date(2021, time.June, 27, 1, 2, 38, 313846900, time.UTC)),
	// 		},
	// 	},
	// }
}
