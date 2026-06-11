package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-01-01-preview/Reservations_GetBillingStatus_MaximumSet_Gen.json
func ExampleReservationsClient_GetBillingStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationsClient().GetBillingStatus(ctx, "rgpurestorage", "reservationname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.ReservationsClientGetBillingStatusResponse{
	// 	ReservationBillingStatus: armpurestorageblock.ReservationBillingStatus{
	// 		Timestamp: to.Ptr("2026-01-16T07:25:56.721Z"),
	// 		TotalUsedCapacityReported: to.Ptr[int64](8),
	// 		LowDrrPoolCount: to.Ptr[int32](5),
	// 		DrrWeightedAverage: to.Ptr[float64](15),
	// 		TotalNonReducibleReported: to.Ptr[int64](9),
	// 		ExtraUsedCapacityNonReducible: to.Ptr[int64](12),
	// 		ExtraUsedCapacityLowUsageRounding: to.Ptr[int64](9),
	// 		ExtraUsedCapacityNonReduciblePlanDiscount: to.Ptr[int64](19),
	// 		TotalUsedCapacityBilled: to.Ptr[int64](7),
	// 		TotalUsedCapacityIncludedPlan: to.Ptr[int64](18),
	// 		TotalUsedCapacityOverage: to.Ptr[int64](28),
	// 		TotalPerformanceReported: to.Ptr[int64](29),
	// 		TotalPerformanceIncludedPlan: to.Ptr[int64](23),
	// 		TotalPerformanceOverage: to.Ptr[int64](16),
	// 	},
	// }
}
