package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationsListByBillingProfile.json
func ExampleReservationsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsClient().NewListByBillingProfilePager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "AAAA-AAAA-AAA-AAA", &armbilling.ReservationsClientListByBillingProfileOptions{Filter: nil,
		OrderBy:        nil,
		Skiptoken:      nil,
		RefreshSummary: nil,
		SelectedState:  to.Ptr("Succeeded"),
		Take:           nil,
	})
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
		// page.ReservationsListResult = armbilling.ReservationsListResult{
		// 	Summary: &armbilling.ReservationSummary{
		// 		CancelledCount: to.Ptr[float32](0),
		// 		ExpiredCount: to.Ptr[float32](0),
		// 		ExpiringCount: to.Ptr[float32](0),
		// 		FailedCount: to.Ptr[float32](0),
		// 		NoBenefitCount: to.Ptr[float32](0),
		// 		PendingCount: to.Ptr[float32](0),
		// 		ProcessingCount: to.Ptr[float32](0),
		// 		SucceededCount: to.Ptr[float32](1),
		// 		WarningCount: to.Ptr[float32](0),
		// 	},
		// 	Value: []*armbilling.Reservation{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("microsoft.billing/billingAccounts/reservationOrders/reservations"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/reservationOrders/00000000-0000-0000-0000-000000000001/reservations/00000000-0000-0000-0000-000000000000"),
		// 			Location: to.Ptr("global"),
		// 			Properties: &armbilling.ReservationProperty{
		// 				AppliedScopeType: to.Ptr("Shared"),
		// 				Archived: to.Ptr(false),
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-18T21:54:31.074Z"); return t}()),
		// 				BillingPlan: to.Ptr(armbilling.ReservationBillingPlanMonthly),
		// 				BillingScopeID: to.Ptr("/subscriptions/eef82110-c91b-4395-9420-fcfcbefc5a47"),
		// 				DisplayName: to.Ptr("VirtualMachineSoftware_01-18-2022_13-51"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				ExpiryDate: to.Ptr("2025-01-18"),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-18T21:54:31.074Z"); return t}()),
		// 				InstanceFlexibility: to.Ptr(armbilling.InstanceFlexibilityOn),
		// 				LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 				ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-01-18"); return t}()),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-18T21:51:29.906Z"); return t}()),
		// 				Quantity: to.Ptr[float32](1),
		// 				Renew: to.Ptr(false),
		// 				ReservedResourceType: to.Ptr("VirtualMachineSoftware"),
		// 				SKUDescription: to.Ptr("Sku description"),
		// 				Term: to.Ptr("P3Y"),
		// 				UserFriendlyAppliedScopeType: to.Ptr("Shared"),
		// 				UserFriendlyRenewState: to.Ptr("Off"),
		// 				Utilization: &armbilling.ReservationPropertyUtilization{
		// 					Aggregates: []*armbilling.ReservationUtilizationAggregates{
		// 						{
		// 							Grain: to.Ptr[float32](1),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](7),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](30),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 					}},
		// 					Trend: to.Ptr("SAME"),
		// 				},
		// 			},
		// 			SKU: &armbilling.ReservationSKUProperty{
		// 				Name: to.Ptr("mock_sku"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 	}},
		// }
	}
}
