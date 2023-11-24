package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationOrders.json
func ExampleReservationOrderClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationOrderClient().NewListPager(nil)
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
		// page.ReservationOrderList = armreservations.ReservationOrderList{
		// 	Value: []*armreservations.ReservationOrderResponse{
		// 		{
		// 			Name: to.Ptr("1e6407ba-37a5-499f-80ed-a3f0f338e443"),
		// 			Type: to.Ptr("Microsoft.Capacity/reservationOrders"),
		// 			Etag: to.Ptr[int32](7),
		// 			ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1e6407ba-37a5-499f-80ed-a3f0f338e443"),
		// 			Properties: &armreservations.ReservationOrderProperties{
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T21:20:23.813Z"); return t}()),
		// 				BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T21:22:56.854Z"); return t}()),
		// 				DisplayName: to.Ptr("cabri"),
		// 				ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-08-29"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T21:20:23.813Z"); return t}()),
		// 				OriginalQuantity: to.Ptr[int32](1),
		// 				ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 				RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T21:20:23.813Z"); return t}()),
		// 				Reservations: []*armreservations.ReservationResponse{
		// 					{
		// 						ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1e6407ba-37a5-499f-80ed-a3f0f338e443/reservations/cae5924e-7a15-419f-a369-124f52d4a106"),
		// 				}},
		// 				Term: to.Ptr(armreservations.ReservationTermP1Y),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1ea6e203-288e-4732-b9e1-da8bbe10c614"),
		// 			Type: to.Ptr("Microsoft.Capacity/reservationOrders"),
		// 			Etag: to.Ptr[int32](7),
		// 			ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1ea6e203-288e-4732-b9e1-da8bbe10c614"),
		// 			Properties: &armreservations.ReservationOrderProperties{
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T12:55:40.279Z"); return t}()),
		// 				BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T12:58:20.526Z"); return t}()),
		// 				DisplayName: to.Ptr("cabri"),
		// 				ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-08-30"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T12:58:20.526Z"); return t}()),
		// 				OriginalQuantity: to.Ptr[int32](9),
		// 				ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 				RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T12:55:40.279Z"); return t}()),
		// 				Reservations: []*armreservations.ReservationResponse{
		// 					{
		// 						ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1ea6e203-288e-4732-b9e1-da8bbe10c614/reservations/d04fd48d-e3f6-42a3-a8f6-1ad0b7513e48"),
		// 				}},
		// 				Term: to.Ptr(armreservations.ReservationTermP1Y),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1f14354c-dc12-4c8d-8090-6f295a3a34aa"),
		// 			Type: to.Ptr("Microsoft.Capacity/reservationOrders"),
		// 			Etag: to.Ptr[int32](7),
		// 			ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa"),
		// 			Properties: &armreservations.ReservationOrderProperties{
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T03:49:19.025Z"); return t}()),
		// 				BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T03:51:49.808Z"); return t}()),
		// 				DisplayName: to.Ptr("cabri"),
		// 				ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-08-30"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T03:51:49.808Z"); return t}()),
		// 				OriginalQuantity: to.Ptr[int32](7),
		// 				ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 				RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T03:49:19.025Z"); return t}()),
		// 				Reservations: []*armreservations.ReservationResponse{
		// 					{
		// 						ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa/reservations/c8c926bd-fc5d-4e29-9d43-b68340ac23a6"),
		// 				}},
		// 				Term: to.Ptr(armreservations.ReservationTermP1Y),
		// 			},
		// 	}},
		// }
	}
}
