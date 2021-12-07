Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Forbital%2Farmorbital%2Fv0.1.0/sdk/resourcemanager/orbital/armorbital/README.md) on how to add the SDK to your project and authenticate.

```go
package armorbital_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital"
)

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/ContactCreate.json
func ExampleContactsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewContactsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		"<contact-name>",
		armorbital.Contact{
			Properties: &armorbital.ContactsProperties{
				ContactProfile: &armorbital.ResourceReference{
					ID: to.StringPtr("<id>"),
				},
				GroundStationName:    to.StringPtr("<ground-station-name>"),
				ReservationEndTime:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-16T20:55:00.00Z"); return t }()),
				ReservationStartTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-16T20:35:00.00Z"); return t }()),
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
	log.Printf("Contact.ID: %s\n", *res.ID)
}
```
