package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/ChangeDirectoryReservationOrder.json
func ExampleReservationOrderClient_ChangeDirectory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewReservationOrderClient(cred, nil)
	res, err := client.ChangeDirectory(ctx,
		"<reservation-order-id>",
		armreservations.ChangeDirectoryRequest{
			DestinationTenantID: to.StringPtr("<destination-tenant-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReservationOrderClientChangeDirectoryResult)
}
