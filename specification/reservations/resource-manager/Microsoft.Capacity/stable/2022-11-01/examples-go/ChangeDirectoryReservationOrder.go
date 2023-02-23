package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/ChangeDirectoryReservationOrder.json
func ExampleReservationOrderClient_ChangeDirectory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewReservationOrderClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ChangeDirectory(ctx, "a075419f-44cc-497f-b68a-14ee811d48b9", armreservations.ChangeDirectoryRequest{
		DestinationTenantID: to.Ptr("906655ea-30be-4587-9d12-b50e077b0f32"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ChangeDirectoryResponse = armreservations.ChangeDirectoryResponse{
	// 	ReservationOrder: &armreservations.ChangeDirectoryResult{
	// 		Name: to.Ptr("VM_RI_10-02-2020_15-21"),
	// 		Error: to.Ptr("error string"),
	// 		ID: to.Ptr("a075419f-44cc-497f-b68a-14ee811d48b9"),
	// 		IsSucceeded: to.Ptr(true),
	// 	},
	// 	Reservations: []*armreservations.ChangeDirectoryResult{
	// 		{
	// 			Name: to.Ptr("VM_RI_10-02-2020_15-21"),
	// 			Error: to.Ptr("error string"),
	// 			ID: to.Ptr("1f14354c-dc12-4c8d-8090-6f295a3a34aa"),
	// 			IsSucceeded: to.Ptr(true),
	// 	}},
	// }
}
