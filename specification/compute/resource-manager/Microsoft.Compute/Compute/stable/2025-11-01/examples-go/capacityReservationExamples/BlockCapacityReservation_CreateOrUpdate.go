package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/capacityReservationExamples/BlockCapacityReservation_CreateOrUpdate.json
func ExampleCapacityReservationsClient_BeginCreateOrUpdate_createOrUpdateABlockCapacityReservation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "blockCapacityReservationGroup", "blockCapacityReservation", armcompute.CapacityReservation{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("Standard_ND96isr_H100_v5"),
			Capacity: to.Ptr[int64](1),
		},
		Properties: &armcompute.CapacityReservationProperties{
			ScheduleProfile: &armcompute.ScheduleProfile{
				Start: to.Ptr("2025-08-01"),
				End:   to.Ptr("2025-08-02"),
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CapacityReservationsClientCreateOrUpdateResponse{
	// 	CapacityReservation: &armcompute.CapacityReservation{
	// 		Name: to.Ptr("blockCapacityReservation"),
	// 		Location: to.Ptr("westus"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/blockCapacityReservationGroup/capacityReservations/blockCapacityReservation"),
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
	// 		Properties: &armcompute.CapacityReservationProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](3),
	// 			ReservationID: to.Ptr("{GUID}"),
	// 			ProvisioningState: to.Ptr("Creating"),
	// 			ScheduleProfile: &armcompute.ScheduleProfile{
	// 				Start: to.Ptr("2025-08-01"),
	// 				End: to.Ptr("2025-08-02"),
	// 			},
	// 			ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.3138469+00:00"); return t}()),
	// 		},
	// 	},
	// }
}
