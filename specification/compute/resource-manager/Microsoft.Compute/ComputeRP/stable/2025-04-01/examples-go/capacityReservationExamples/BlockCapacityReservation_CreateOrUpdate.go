package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/capacityReservationExamples/BlockCapacityReservation_CreateOrUpdate.json
func ExampleCapacityReservationsClient_BeginCreateOrUpdate_createOrUpdateABlockCapacityReservation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "blockCapacityReservationGroup", "blockCapacityReservation", armcompute.CapacityReservation{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
		},
		Properties: &armcompute.CapacityReservationProperties{
			ScheduleProfile: &armcompute.ScheduleProfile{
				End:   to.Ptr("2025-08-02"),
				Start: to.Ptr("2025-08-01"),
			},
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("Standard_ND96isr_H100_v5"),
			Capacity: to.Ptr[int64](1),
		},
		Zones: []*string{
			to.Ptr("1")},
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
	// res.CapacityReservation = armcompute.CapacityReservation{
	// 	Name: to.Ptr("blockCapacityReservation"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/blockCapacityReservationGroup/capacityReservations/blockCapacityReservation"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.CapacityReservationProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
	// 		ReservationID: to.Ptr("{GUID}"),
	// 		ScheduleProfile: &armcompute.ScheduleProfile{
	// 			End: to.Ptr("2025-08-02"),
	// 			Start: to.Ptr("2025-08-01"),
	// 		},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_ND96isr_H100_v5"),
	// 		Capacity: to.Ptr[int64](1),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
