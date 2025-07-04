package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01/Reservations_GetBillingReport_MaximumSet_Gen.json
func ExampleReservationsClient_GetBillingReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("BC47D6CC-AA80-4374-86F8-19D94EC70666", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationsClient().GetBillingReport(ctx, "rgpurestorage", "reservationname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.ReservationsClientGetBillingReportResponse{
	// 	ReservationBillingUsageReport: &armpurestorageblock.ReservationBillingUsageReport{
	// 		Timestamp: to.Ptr("2024-10-04T05:29:25.345Z"),
	// 		BillingUsageProperties: []*armpurestorageblock.BillingUsageProperty{
	// 			{
	// 				PropertyID: to.Ptr("fknpxmzbrocjevhnuxohiwl"),
	// 				PropertyName: to.Ptr("rznqcuwmulhtvp"),
	// 				CurrentValue: to.Ptr("ndiuedcwtwpedirqdq"),
	// 				PreviousValue: to.Ptr("fqnvq"),
	// 				Severity: to.Ptr(armpurestorageblock.UsageSeverityALERT),
	// 				StatusMessage: to.Ptr("vkce"),
	// 				SubProperties: []*armpurestorageblock.BillingUsageProperty{
	// 				},
	// 			},
	// 		},
	// 		OverallStatusMessage: to.Ptr("aurwogtwwsxjoocpsobslpv"),
	// 	},
	// }
}
