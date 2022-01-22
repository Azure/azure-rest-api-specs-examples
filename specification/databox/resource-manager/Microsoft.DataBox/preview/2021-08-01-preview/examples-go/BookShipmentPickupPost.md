Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv0.2.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/BookShipmentPickupPost.json
func ExampleJobsClient_BookShipmentPickUp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	res, err := client.BookShipmentPickUp(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.ShipmentPickUpRequest{
			EndTime:          to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-22T18:30:00Z"); return t }()),
			ShipmentLocation: to.StringPtr("<shipment-location>"),
			StartTime:        to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-20T18:30:00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobsClientBookShipmentPickUpResult)
}
```
