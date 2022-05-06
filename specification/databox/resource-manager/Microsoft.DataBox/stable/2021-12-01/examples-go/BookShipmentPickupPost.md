Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv0.4.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatabox_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2021-12-01/examples/BookShipmentPickupPost.json
func ExampleJobsClient_BookShipmentPickUp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.BookShipmentPickUp(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.ShipmentPickUpRequest{
			EndTime:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-22T18:30:00Z"); return t }()),
			ShipmentLocation: to.Ptr("<shipment-location>"),
			StartTime:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-20T18:30:00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
