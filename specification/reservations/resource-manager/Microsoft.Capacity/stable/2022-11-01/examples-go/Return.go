package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/Return.json
func ExampleReturnClient_BeginPost() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReturnClient().BeginPost(ctx, "50000000-aaaa-bbbb-cccc-100000000004", armreservations.RefundRequest{
		Properties: &armreservations.RefundRequestProperties{
			ReservationToReturn: &armreservations.ReservationToReturn{
				Quantity:      to.Ptr[int32](1),
				ReservationID: to.Ptr("/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004/reservations/40000000-aaaa-bbbb-cccc-100000000000"),
			},
			ReturnReason: to.Ptr("PurchasedWrongProduct"),
			Scope:        to.Ptr("Reservation"),
			SessionID:    to.Ptr("10000000-aaaa-bbbb-cccc-200000000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReservationOrderResponse = armreservations.ReservationOrderResponse{
	// 	Name: to.Ptr("50000000-aaaa-bbbb-cccc-100000000004"),
	// 	Type: to.Ptr("Microsoft.Capacity/reservationOrders"),
	// 	Etag: to.Ptr[int32](7),
	// 	ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004"),
	// 	Properties: &armreservations.ReservationOrderProperties{
	// 		BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T03:51:49.808Z"); return t}()),
	// 		DisplayName: to.Ptr("cabri"),
	// 		ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-08-30"); return t}()),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T03:51:49.808Z"); return t}()),
	// 		OriginalQuantity: to.Ptr[int32](7),
	// 		ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
	// 		RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T03:49:19.025Z"); return t}()),
	// 		Reservations: []*armreservations.ReservationResponse{
	// 			{
	// 				ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004/reservations/40000000-aaaa-bbbb-cccc-100000000000"),
	// 		}},
	// 		Term: to.Ptr(armreservations.ReservationTermP1Y),
	// 	},
	// }
}
