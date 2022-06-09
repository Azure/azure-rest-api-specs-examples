```go
package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/GetCatalog.json
func ExampleAzureReservationAPIClient_GetCatalog() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewAzureReservationAPIClient(cred, nil)
	res, err := client.GetCatalog(ctx,
		"<subscription-id>",
		&armreservations.AzureReservationAPIClientGetCatalogOptions{ReservedResourceType: to.StringPtr("<reserved-resource-type>"),
			Location: to.StringPtr("<location>"),
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AzureReservationAPIClientGetCatalogResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Freservations%2Farmreservations%2Fv0.2.1/sdk/resourcemanager/reservations/armreservations/README.md) on how to add the SDK to your project and authenticate.
