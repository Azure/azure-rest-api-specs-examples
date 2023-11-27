package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesDailyWithReservationId.json
func ExampleReservationsSummariesClient_NewListByReservationOrderAndReservationPager_reservationSummariesDailyWithReservationId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsSummariesClient().NewListByReservationOrderAndReservationPager("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000", armconsumption.DatagrainDailyGrain, &armconsumption.ReservationsSummariesClientListByReservationOrderAndReservationOptions{Filter: to.Ptr("properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-11-20")})
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
		// page.ReservationSummariesListResult = armconsumption.ReservationSummariesListResult{
		// 	Value: []*armconsumption.ReservationSummary{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000_20171001"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationSummaries"),
		// 			ID: to.Ptr("providers/Microsoft.Capacity/reservationOrders/00000000-0000-0000-0000-000000000000/reservations/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/reservationSummaries/20171001"),
		// 			Tags: map[string]*string{
		// 				"dev": to.Ptr("tools"),
		// 				"env": to.Ptr("newcrp"),
		// 			},
		// 			Properties: &armconsumption.ReservationSummaryProperties{
		// 				AvgUtilizationPercentage: to.Ptr[float64](0),
		// 				Kind: to.Ptr("Reservation"),
		// 				MaxUtilizationPercentage: to.Ptr[float64](0),
		// 				MinUtilizationPercentage: to.Ptr[float64](0),
		// 				PurchasedQuantity: to.Ptr[float64](0),
		// 				RemainingQuantity: to.Ptr[float64](0),
		// 				ReservationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ReservationOrderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ReservedHours: to.Ptr[float64](0),
		// 				SKUName: to.Ptr("Standard_D8s_v3"),
		// 				TotalReservedQuantity: to.Ptr[float64](155),
		// 				UsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-01T00:00:00.000Z"); return t}()),
		// 				UsedHours: to.Ptr[float64](0),
		// 				UsedQuantity: to.Ptr[float64](0),
		// 				UtilizedPercentage: to.Ptr[float64](0),
		// 			},
		// 	}},
		// }
	}
}
