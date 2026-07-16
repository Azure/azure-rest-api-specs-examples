package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ReservationDetailsByBillingProfileId.json
func ExampleReservationsDetailsClient_NewListPager_reservationDetailsByBillingProfileId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsDetailsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579", &armconsumption.ReservationsDetailsClientListOptions{
		EndDate:   to.Ptr("2019-10-31"),
		StartDate: to.Ptr("2019-09-01")})
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
		// page = armconsumption.ReservationsDetailsClientListResponse{
		// 	ReservationDetailsListResult: armconsumption.ReservationDetailsListResult{
		// 		Value: []*armconsumption.ReservationDetail{
		// 			{
		// 				Name: to.Ptr("reservationDetails_Id1"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationDetails"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579/providers/Microsoft.Consumption/reservationDetails/reservationDetails_Id1"),
		// 				Properties: &armconsumption.ReservationDetailProperties{
		// 					InstanceFlexibilityGroup: to.Ptr("DSv3 Series"),
		// 					InstanceFlexibilityRatio: to.Ptr("1"),
		// 					InstanceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sqlh1/providers/microsoft.compute/virtualmachines/sqlh1"),
		// 					ReservationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ReservationOrderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ReservedHours: to.Ptr[float64](48),
		// 					SKUName: to.Ptr("Standard_D2s_v3"),
		// 					TotalReservedQuantity: to.Ptr[float64](0),
		// 					UsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-30T00:00:00-08:00"); return t}()),
		// 					UsedHours: to.Ptr[float64](0.6),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
