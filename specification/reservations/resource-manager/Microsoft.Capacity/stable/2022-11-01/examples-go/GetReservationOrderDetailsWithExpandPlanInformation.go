package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationOrderDetailsWithExpandPlanInformation.json
func ExampleReservationOrderClient_Get_getReservationWithExpandPayments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationOrderClient().Get(ctx, "a075419f-44cc-497f-b68a-14ee811d48b9", &armreservations.ReservationOrderClientGetOptions{Expand: to.Ptr("schedule")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReservationOrderResponse = armreservations.ReservationOrderResponse{
	// 	Name: to.Ptr("1f14354c-dc12-4c8d-8090-6f295a3a34aa"),
	// 	Type: to.Ptr("Microsoft.Capacity/reservationOrders"),
	// 	Etag: to.Ptr[int32](7),
	// 	ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa"),
	// 	Properties: &armreservations.ReservationOrderProperties{
	// 		BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T03:51:49.000Z"); return t}()),
	// 		BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T03:51:49.000Z"); return t}()),
	// 		DisplayName: to.Ptr("cabri"),
	// 		ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-08-30"); return t}()),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T03:51:49.000Z"); return t}()),
	// 		OriginalQuantity: to.Ptr[int32](7),
	// 		PlanInformation: &armreservations.ReservationOrderBillingPlanInformation{
	// 			NextPaymentDueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-09-30"); return t}()),
	// 			PricingCurrencyTotal: &armreservations.Price{
	// 				Amount: to.Ptr[float64](466),
	// 				CurrencyCode: to.Ptr("USD"),
	// 			},
	// 			StartDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-08-30"); return t}()),
	// 			Transactions: []*armreservations.PaymentDetail{
	// 				{
	// 					BillingAccount: to.Ptr("23456"),
	// 					BillingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-09-30"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-09-30"); return t}()),
	// 					PricingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armreservations.PaymentStatusSucceeded),
	// 				},
	// 				{
	// 					BillingAccount: to.Ptr("23456"),
	// 					BillingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-10-30"); return t}()),
	// 					ExtendedStatusInfo: &armreservations.ExtendedStatusInfo{
	// 						Message: to.Ptr("Credit card charge failed."),
	// 						StatusCode: to.Ptr(armreservations.ReservationStatusCode("PaymentDeclined")),
	// 					},
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-10-30"); return t}()),
	// 					PricingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armreservations.PaymentStatusFailed),
	// 				},
	// 				{
	// 					BillingAccount: to.Ptr("23456"),
	// 					BillingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-10-31"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-10-31"); return t}()),
	// 					PricingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armreservations.PaymentStatusSucceeded),
	// 				},
	// 				{
	// 					BillingAccount: to.Ptr("23456"),
	// 					BillingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-11-30"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-11-30"); return t}()),
	// 					PricingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armreservations.PaymentStatusSucceeded),
	// 				},
	// 				{
	// 					BillingAccount: to.Ptr("23456"),
	// 					BillingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-07-14"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-06-14"); return t}()),
	// 					PricingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armreservations.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					BillingAccount: to.Ptr("23456"),
	// 					BillingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-06-14"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-08-14"); return t}()),
	// 					PricingCurrencyTotal: &armreservations.Price{
	// 						Amount: to.Ptr[float64](466),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armreservations.PaymentStatusScheduled),
	// 			}},
	// 		},
	// 		ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
	// 		RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T03:49:19.000Z"); return t}()),
	// 		Reservations: []*armreservations.ReservationResponse{
	// 			{
	// 				ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa/reservations/c8c926bd-fc5d-4e29-9d43-b68340ac23a6"),
	// 		}},
	// 		Term: to.Ptr(armreservations.ReservationTermP1Y),
	// 	},
	// }
}
