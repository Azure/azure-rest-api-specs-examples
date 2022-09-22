package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/CalculateRefund.json
func ExampleCalculateRefundClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewCalculateRefundClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Post(ctx, "276e7ae4-84d0-4da6-ab4b-d6b94f3557da", armreservations.CalculateRefundRequest{
		ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004"),
		Properties: &armreservations.CalculateRefundRequestProperties{
			ReservationToReturn: &armreservations.ReservationToReturn{
				Quantity:      to.Ptr[int32](1),
				ReservationID: to.Ptr("/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004/reservations/40000000-aaaa-bbbb-cccc-100000000000"),
			},
			Scope: to.Ptr("Reservation"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
