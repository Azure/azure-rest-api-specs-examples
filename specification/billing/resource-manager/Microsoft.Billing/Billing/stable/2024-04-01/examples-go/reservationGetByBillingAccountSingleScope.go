package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/reservationGetByBillingAccountSingleScope.json
func ExampleReservationsClient_GetByReservationOrder_reservationGetByBillingAccountSingleScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationsClient().GetByReservationOrder(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.ReservationsClientGetByReservationOrderResponse{
	// 	Reservation: armbilling.Reservation{
	// 		Name: to.Ptr("30000000-0000-0000-0000-000000000000"),
	// 		Type: to.Ptr("microsoft.billing/billingAccounts/reservationOrders/reservations"),
	// 		Etag: to.Ptr[int32](19),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/reservationOrders/20000000-0000-0000-0000-000000000000/reservations/30000000-0000-0000-0000-000000000000"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armbilling.ReservationProperty{
	// 			AppliedScopeProperties: &armbilling.ReservationAppliedScopeProperties{
	// 				DisplayName: to.Ptr("Azure subscription 1"),
	// 				SubscriptionID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000009"),
	// 			},
	// 			AppliedScopeType: to.Ptr("Single"),
	// 			Archived: to.Ptr(false),
	// 			BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:34:12.9266987Z"); return t}()),
	// 			BillingPlan: to.Ptr(armbilling.ReservationBillingPlanMonthly),
	// 			BillingScopeID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000009"),
	// 			DisplayName: to.Ptr("VM_RI_11-24-2021_22-30"),
	// 			DisplayProvisioningState: to.Ptr("Succeeded"),
	// 			EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-26T01:14:18.558909Z"); return t}()),
	// 			ExpiryDate: to.Ptr("2024-11-25"),
	// 			ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-25T06:34:12.9266987Z"); return t}()),
	// 			InstanceFlexibility: to.Ptr(armbilling.InstanceFlexibilityOff),
	// 			LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-26T01:14:18.6057756Z"); return t}()),
	// 			ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.DateOnly, "2021-11-25"); return t}()),
	// 			PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:31:18.3925493Z"); return t}()),
	// 			Quantity: to.Ptr[float32](1),
	// 			Renew: to.Ptr(true),
	// 			RenewProperties: &armbilling.RenewPropertiesResponse{
	// 				BillingCurrencyTotal: &armbilling.Price{
	// 					Amount: to.Ptr[float64](715.68),
	// 					CurrencyCode: to.Ptr("USD"),
	// 				},
	// 				PricingCurrencyTotal: &armbilling.Price{
	// 					Amount: to.Ptr[float64](715.68),
	// 					CurrencyCode: to.Ptr("USD"),
	// 				},
	// 				PurchaseProperties: &armbilling.ReservationPurchaseRequest{
	// 					Location: to.Ptr("westus"),
	// 					Properties: &armbilling.ReservationPurchaseRequestProperties{
	// 						AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeSingle),
	// 						AppliedScopes: []*string{
	// 							to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000009"),
	// 						},
	// 						BillingPlan: to.Ptr(armbilling.ReservationBillingPlanMonthly),
	// 						BillingScopeID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000009"),
	// 						DisplayName: to.Ptr("VM_RI_11-24-2021_22-30_renewed"),
	// 						InstanceFlexibility: to.Ptr(armbilling.InstanceFlexibilityOff),
	// 						Quantity: to.Ptr[int32](1),
	// 						Renew: to.Ptr(false),
	// 						ReservedResourceProperties: &armbilling.ReservationPurchaseRequestPropertiesReservedResourceProperties{
	// 							InstanceFlexibility: to.Ptr(armbilling.InstanceFlexibilityOff),
	// 						},
	// 						ReservedResourceType: to.Ptr("VirtualMachines"),
	// 						Term: to.Ptr("P3Y"),
	// 					},
	// 					SKU: &armbilling.SKUName{
	// 						Name: to.Ptr("Standard_DS1_v2"),
	// 					},
	// 				},
	// 			},
	// 			ReservedResourceType: to.Ptr("VirtualMachines"),
	// 			SKUDescription: to.Ptr("Reserved VM Instance, Standard_DS1_v2, US West, 3 Years"),
	// 			Term: to.Ptr("P3Y"),
	// 			UserFriendlyAppliedScopeType: to.Ptr("Single"),
	// 			UserFriendlyRenewState: to.Ptr("On"),
	// 			Utilization: &armbilling.ReservationPropertyUtilization{
	// 				Aggregates: []*armbilling.ReservationUtilizationAggregates{
	// 					{
	// 						Grain: to.Ptr[float32](1),
	// 						GrainUnit: to.Ptr("days"),
	// 						Value: to.Ptr[float32](0),
	// 						ValueUnit: to.Ptr("percentage"),
	// 					},
	// 					{
	// 						Grain: to.Ptr[float32](7),
	// 						GrainUnit: to.Ptr("days"),
	// 						Value: to.Ptr[float32](0),
	// 						ValueUnit: to.Ptr("percentage"),
	// 					},
	// 					{
	// 						Grain: to.Ptr[float32](30),
	// 						GrainUnit: to.Ptr("days"),
	// 						Value: to.Ptr[float32](0),
	// 						ValueUnit: to.Ptr("percentage"),
	// 					},
	// 				},
	// 				Trend: to.Ptr("SAME"),
	// 			},
	// 		},
	// 		SKU: &armbilling.ReservationSKUProperty{
	// 			Name: to.Ptr("Standard_DS1_v2"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
