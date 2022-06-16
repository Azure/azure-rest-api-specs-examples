package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/GetAppliedReservations.json
func ExampleAzureReservationAPIClient_GetAppliedReservationList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewAzureReservationAPIClient(cred, nil)
	res, err := client.GetAppliedReservationList(ctx,
		"<subscription-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AzureReservationAPIClientGetAppliedReservationListResult)
}
