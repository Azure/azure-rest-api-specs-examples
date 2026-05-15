package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/reservationOrdersListByBillingAccount.json
func ExampleReservationOrdersClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationOrdersClient().NewListByBillingAccountPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", nil)
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
		// page = armbilling.ReservationOrdersClientListByBillingAccountResponse{
		// 	ReservationOrderList: armbilling.ReservationOrderList{
		// 		Value: []*armbilling.ReservationOrder{
		// 			{
		// 				Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
		// 				Type: to.Ptr("microsoft.billing/billingAccounts/reservationOrders"),
		// 				Etag: to.Ptr[int32](10),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/400000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/reservationOrders/20000000-0000-0000-0000-000000000000"),
		// 				Properties: &armbilling.ReservationOrderProperty{
		// 					BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-03T21:26:48.5128969Z"); return t}()),
		// 					BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/400000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 					BillingPlan: to.Ptr(armbilling.ReservationBillingPlanUpfront),
		// 					BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/400000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-AAAA-AAA-AAA"),
		// 					CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-03T21:26:50.7784836Z"); return t}()),
		// 					DisplayName: to.Ptr("SUSE_Plan_08-03-2021_14-22"),
		// 					ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-03T21:26:48.5128969Z"); return t}()),
		// 					OriginalQuantity: to.Ptr[int32](1),
		// 					ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-03T21:23:47.9090964Z"); return t}()),
		// 					Reservations: []*armbilling.Reservation{
		// 						{
		// 							ID: to.Ptr("/providers/Microsoft.Capacity/reservationOrders/20000000-0000-0000-0000-000000000000/reservations/20000000-0000-0000-0000-000000000001"),
		// 						},
		// 					},
		// 					Term: to.Ptr("P3Y"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 					"key2": to.Ptr("value2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
