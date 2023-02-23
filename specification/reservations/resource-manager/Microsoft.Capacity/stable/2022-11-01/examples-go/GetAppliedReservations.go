package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetAppliedReservations.json
func ExampleAzureReservationAPIClient_GetAppliedReservationList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewAzureReservationAPIClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetAppliedReservationList(ctx, "23bc208b-083f-4901-ae85-4f98c0c3b4b6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppliedReservations = armreservations.AppliedReservations{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Capacity/AppliedReservations"),
	// 	ID: to.Ptr("/subscriptions/23bc208b-083f-4901-ae85-4f98c0c3b4b6/providers/microsoft.capacity/AppliedReservations/default"),
	// 	Properties: &armreservations.AppliedReservationsProperties{
	// 		ReservationOrderIDs: &armreservations.AppliedReservationList{
	// 			Value: []*string{
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/e1eccf0b-2db4-4e84-97e7-98b50e9d46f7"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/741a32eb-6158-4cee-9642-a0243ae79fac"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/e061223d-fcff-4d10-bd49-56a740cfb96a"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/2eeb7234-970e-4663-b60b-85241b515901"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/9db2f4c5-b1c5-42a8-bd79-ee56cdde2c7f"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/5da7a877-6d6e-44af-8880-ed3f533bf928"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/f65b0d0a-f945-4105-821c-d00bc8bacde8"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/51304124-e477-4b07-b9fa-03b05c8b924b"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/f5409b98-8a42-4dc6-be0a-cc59bef4d0db"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/a495550a-80a4-46f8-8843-34d4df46f9a6"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/1a966e18-c272-4ce1-a0c2-d4e1039023c3"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/939310b4-f9de-4645-9569-ab5b6cfe958e"),
	// 				to.Ptr("/providers/Microsoft.Capacity/reservationorders/4193a889-7c3b-44dc-8b7b-bfd7aad6c723")},
	// 			},
	// 		},
	// 	}
}
