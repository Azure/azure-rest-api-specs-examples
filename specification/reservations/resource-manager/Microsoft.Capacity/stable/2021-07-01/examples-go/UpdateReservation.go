package armreservations_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/UpdateReservation.json
func ExampleReservationClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewReservationClient(cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<reservation-order-id>",
		"<reservation-id>",
		armreservations.Patch{
			Properties: &armreservations.PatchProperties{
				AppliedScopeType:    armreservations.AppliedScopeType("Shared").ToPtr(),
				InstanceFlexibility: armreservations.InstanceFlexibility("Off").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReservationClientUpdateResult)
}
