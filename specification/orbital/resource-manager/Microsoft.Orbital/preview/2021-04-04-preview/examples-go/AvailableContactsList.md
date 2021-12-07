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

// x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/preview/2021-04-04-preview/examples/AvailableContactsList.json
func ExampleSpacecraftsClient_BeginListAvailableContacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armorbital.NewSpacecraftsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginListAvailableContacts(ctx,
		"<resource-group-name>",
		"<spacecraft-name>",
		armorbital.ContactParameters{
			ContactProfile: &armorbital.ResourceReference{
				ID: to.StringPtr("<id>"),
			},
			EndTime:           to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-17T23:49:40.00Z"); return t }()),
			GroundStationName: to.StringPtr("<ground-station-name>"),
			StartTime:         to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-16T05:40:21.00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
