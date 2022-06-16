package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetReservationDetails.json
func ExampleReservationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewReservationClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"6ef59113-3482-40da-8d79-787f823e34bc",
		"276e7ae4-84d0-4da6-ab4b-d6b94f3557da",
		&armreservations.ReservationClientGetOptions{Expand: to.Ptr("renewProperties")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
