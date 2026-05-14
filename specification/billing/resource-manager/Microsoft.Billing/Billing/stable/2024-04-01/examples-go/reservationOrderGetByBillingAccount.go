package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/reservationOrderGetByBillingAccount.json
func ExampleReservationOrdersClient_GetByBillingAccount_reservationOrderGetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationOrdersClient().GetByBillingAccount(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.ReservationOrdersClientGetByBillingAccountResponse{
	// 	ReservationOrder: armbilling.ReservationOrder{
	// 		Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Type: to.Ptr("microsoft.billing/billingAccounts/reservationOrders"),
	// 		Etag: to.Ptr[int32](27),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/reservationOrders/20000000-0000-0000-0000-000000000000"),
	// 		Properties: &armbilling.ReservationOrderProperty{
	// 			BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:34:12.9266987Z"); return t}()),
	// 			BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
	// 			BillingPlan: to.Ptr(armbilling.ReservationBillingPlanMonthly),
	// 			BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-AAAA-AAA-AAA"),
	// 			CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:34:13.9736776Z"); return t}()),
	// 			DisplayName: to.Ptr("VM_RI_11-24-2021_22-30"),
	// 			ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-25T06:34:12.9266987Z"); return t}()),
	// 			OriginalQuantity: to.Ptr[int32](1),
	// 			ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:31:18.3925493Z"); return t}()),
	// 			Reservations: []*armbilling.Reservation{
	// 				{
	// 					ID: to.Ptr("/providers/Microsoft.Capacity/reservationOrders/20000000-0000-0000-0000-000000000000/reservations/20000000-0000-0000-0000-000000000001"),
	// 				},
	// 			},
	// 			Term: to.Ptr("P3Y"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
