package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservations.json
func ExampleReservationClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationClient().NewListAllPager(&armreservations.ReservationClientListAllOptions{Filter: to.Ptr("(properties%2farchived+eq+false)"),
		Orderby:        to.Ptr("properties/displayName asc"),
		RefreshSummary: nil,
		Skiptoken:      to.Ptr[float32](50),
		SelectedState:  nil,
		Take:           to.Ptr[float32](1),
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
		// page.ListResult = armreservations.ListResult{
		// 	Summary: &armreservations.ReservationSummary{
		// 		CancelledCount: to.Ptr[float32](0),
		// 		ExpiredCount: to.Ptr[float32](0),
		// 		ExpiringCount: to.Ptr[float32](0),
		// 		FailedCount: to.Ptr[float32](0),
		// 		PendingCount: to.Ptr[float32](0),
		// 		SucceededCount: to.Ptr[float32](1),
		// 	},
		// 	Value: []*armreservations.ReservationResponse{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000001/00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations"),
		// 			ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/00000000-0000-0000-0000-000000000001/reservations/00000000-0000-0000-0000-000000000000"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armreservations.Properties{
		// 				AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeShared),
		// 				DisplayName: to.Ptr("VM_RI_07-21-2020_12-06"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 				ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-07-21"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-21T22:46:32.763Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 				PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-04-22"); return t}()),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-22T22:46:32.763Z"); return t}()),
		// 				Quantity: to.Ptr[int32](2),
		// 				Renew: to.Ptr(false),
		// 				RenewSource: to.Ptr("/providers/Microsoft.Capacity/reservationorders/00000000-0000-0000-0000-000000000002/reservations/00000000-0000-0000-0000-000000000003"),
		// 				ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 				Term: to.Ptr(armreservations.ReservationTermP3Y),
		// 				UserFriendlyAppliedScopeType: to.Ptr("Shared"),
		// 				UserFriendlyRenewState: to.Ptr("Off"),
		// 				Utilization: &armreservations.PropertiesUtilization{
		// 					Aggregates: []*armreservations.ReservationUtilizationAggregates{
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
		// 			SKU: &armreservations.SKUName{
		// 				Name: to.Ptr("Standard_D1"),
		// 			},
		// 	}},
		// }
	}
}
