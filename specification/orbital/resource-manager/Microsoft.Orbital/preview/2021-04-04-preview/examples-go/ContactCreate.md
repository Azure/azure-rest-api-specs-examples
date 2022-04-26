Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Forbital%2Farmorbital%2Fv0.4.0/sdk/resourcemanager/orbital/armorbital/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/ContactCreate.json
func ExampleContactsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armorbital.NewContactsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		"<contact-name>",
		armorbital.Contact{
			Properties: &armorbital.ContactsProperties{
				ContactProfile: &armorbital.ResourceReference{
					ID: to.Ptr("<id>"),
				},
				GroundStationName:    to.Ptr("<ground-station-name>"),
				ReservationEndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-16T20:55:00.00Z"); return t }()),
				ReservationStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-16T20:35:00.00Z"); return t }()),
			},
		},
		&armorbital.ContactsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
