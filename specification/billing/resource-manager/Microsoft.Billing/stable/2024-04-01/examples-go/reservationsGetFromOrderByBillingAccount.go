package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationsGetFromOrderByBillingAccount.json
func ExampleReservationsClient_NewListByReservationOrderPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsClient().NewListByReservationOrderPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", nil)
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
		// page.ReservationList = armbilling.ReservationList{
		// 	Value: []*armbilling.Reservation{
		// 		{
		// 			Name: to.Ptr("a7d70646-b848-4498-8093-5938128b225c"),
		// 			Type: to.Ptr("microsoft.billing/billingAccounts/reservationOrders/reservations"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/4973e1de-a829-5c64-4fef-0a692e2b3108:1970c5da-0aa4-46fd-a917-4772f9a17978_2019-05-31/reservationOrders/99f340d1-6db4-41b4-b469-cfc499716973/reservations/a7d70646-b848-4498-8093-5938128b225c"),
		// 			Etag: to.Ptr[int32](15),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armbilling.ReservationProperty{
		// 				AppliedScopeType: to.Ptr("Shared"),
		// 				Archived: to.Ptr(false),
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:34:12.926Z"); return t}()),
		// 				BillingPlan: to.Ptr(armbilling.ReservationBillingPlanMonthly),
		// 				BillingScopeID: to.Ptr("/subscriptions/eef82110-c91b-4395-9420-fcfcbefc5a47"),
		// 				DisplayName: to.Ptr("VM_RI_11-24-2021_22-30"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-24T01:46:05.425Z"); return t}()),
		// 				ExpiryDate: to.Ptr("2024-11-25"),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-25T06:34:12.926Z"); return t}()),
		// 				InstanceFlexibility: to.Ptr(armbilling.InstanceFlexibilityOn),
		// 				LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-24T01:46:05.534Z"); return t}()),
		// 				ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-11-25"); return t}()),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:31:18.392Z"); return t}()),
		// 				Quantity: to.Ptr[float32](1),
		// 				Renew: to.Ptr(true),
		// 				ReservedResourceType: to.Ptr("VirtualMachines"),
		// 				SKUDescription: to.Ptr("Reserved VM Instance, Standard_DS1_v2, US West, 3 Years"),
		// 				Term: to.Ptr("P3Y"),
		// 				UserFriendlyAppliedScopeType: to.Ptr("Shared"),
		// 				UserFriendlyRenewState: to.Ptr("On"),
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
		// 				Name: to.Ptr("Standard_DS1_v2"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 	}},
		// }
	}
}
