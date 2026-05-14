package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ReservationSummariesDailyWithBillingProfileId.json
func ExampleReservationsSummariesClient_NewListPager_reservationSummariesDailyWithBillingProfileId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsSummariesClient().NewListPager("providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579", armconsumption.DatagrainDailyGrain, &armconsumption.ReservationsSummariesClientListOptions{
		EndDate:   to.Ptr("2017-11-20"),
		StartDate: to.Ptr("2017-10-01")})
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
		// page = armconsumption.ReservationsSummariesClientListResponse{
		// 	ReservationSummariesListResult: armconsumption.ReservationSummariesListResult{
		// 		Value: []*armconsumption.ReservationSummary{
		// 			{
		// 				Name: to.Ptr("reservationSummaries_Id1"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationSummaries"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579/providers/Microsoft.Consumption/reservationSummaries/reservationSummaries_Id1"),
		// 				Properties: &armconsumption.ReservationSummaryProperties{
		// 					ReservationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ReservationOrderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					SKUName: to.Ptr("Standard_B1s"),
		// 					UsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-01T00:00:00-07:00"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
