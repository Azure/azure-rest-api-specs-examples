Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Freservations%2Farmreservations%2Fv1.0.0/sdk/resourcemanager/reservations/armreservations/README.md) on how to add the SDK to your project and authenticate.

```go
package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetCatalog.json
func ExampleAzureReservationAPIClient_GetCatalog() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewAzureReservationAPIClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetCatalog(ctx,
		"23bc208b-083f-4901-ae85-4f98c0c3b4b6",
		&armreservations.AzureReservationAPIClientGetCatalogOptions{ReservedResourceType: to.Ptr("VirtualMachines"),
			Location:    to.Ptr("eastus"),
			PublisherID: nil,
			OfferID:     nil,
			PlanID:      nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
