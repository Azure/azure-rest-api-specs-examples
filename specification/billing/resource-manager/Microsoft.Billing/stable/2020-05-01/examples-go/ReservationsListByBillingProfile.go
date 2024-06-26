package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ReservationsListByBillingProfile.json
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
	pager := clientFactory.NewReservationsClient().NewListByBillingProfilePager("{billingAccountName}", "{billingProfileName}", &armbilling.ReservationsClientListByBillingProfileOptions{Filter: to.Ptr("properties/reservedResourceType eq 'VirtualMachines'"),
		Orderby:        to.Ptr("properties/userFriendlyAppliedScopeType asc"),
		RefreshSummary: nil,
		SelectedState:  to.Ptr("Succeeded"),
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
		// 		PendingCount: to.Ptr[float32](0),
		// 		SucceededCount: to.Ptr[float32](1),
		// 	},
		// 	Value: []*armbilling.Reservation{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000001/00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations"),
		// 			ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/00000000-0000-0000-0000-000000000001/reservations/00000000-0000-0000-0000-000000000000"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armbilling.ReservationProperty{
		// 				AppliedScopeType: to.Ptr("Shared"),
		// 				DisplayName: to.Ptr("VM_RI_07-21-2020_12-06"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr("0001-01-01T00:00:00"),
		// 				ExpiryDate: to.Ptr("2023-07-21"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Quantity: to.Ptr[float32](2),
		// 				Renew: to.Ptr(false),
		// 				RenewSource: to.Ptr("/providers/Microsoft.Capacity/reservationorders/00000000-0000-0000-0000-000000000002/reservations/00000000-0000-0000-0000-000000000003"),
		// 				ReservedResourceType: to.Ptr("VirtualMachines"),
		// 				UserFriendlyAppliedScopeType: to.Ptr("Shared"),
		// 				UserFriendlyRenewState: to.Ptr("Off"),
		// 				Utilization: &armbilling.ReservationPropertyUtilization{
		// 					Aggregates: []*armbilling.ReservationUtilizationAggregates{
		// 						{
		// 							Grain: to.Ptr[float32](1),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0.05),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](7),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0.05),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](30),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0.05),
		// 							ValueUnit: to.Ptr("percentage"),
		// 					}},
		// 					Trend: to.Ptr("UP"),
		// 				},
		// 			},
		// 			SKU: &armbilling.ReservationSKUProperty{
		// 				Name: to.Ptr("Standard_D1"),
		// 			},
		// 	}},
		// }
	}
}
