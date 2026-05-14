package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ReservationSummariesMonthly.json
func ExampleReservationsSummariesClient_NewListByReservationOrderPager_reservationSummariesMonthly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsSummariesClient().NewListByReservationOrderPager("00000000-0000-0000-0000-000000000000", armconsumption.DatagrainMonthlyGrain, nil)
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
		// page = armconsumption.ReservationsSummariesClientListByReservationOrderResponse{
		// 	ReservationSummariesListResult: armconsumption.ReservationSummariesListResult{
		// 		Value: []*armconsumption.ReservationSummary{
		// 			{
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000_20171001"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationSummaries"),
		// 				ID: to.Ptr("providers/Microsoft.Capacity/reservationOrders/00000000-0000-0000-0000-000000000000/reservations/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/reservationSummaries/20171001"),
		// 				Properties: &armconsumption.ReservationSummaryProperties{
		// 					Kind: to.Ptr("Reservation"),
		// 					ReservationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ReservationOrderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					SKUName: to.Ptr("Standard_D8s_v3"),
		// 					UsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-01T00:00:00Z"); return t}()),
		// 				},
		// 				Tags: map[string]*string{
		// 					"dev": to.Ptr("tools"),
		// 					"env": to.Ptr("newcrp"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
