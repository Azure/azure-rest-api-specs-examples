Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnotificationhubs%2Farmnotificationhubs%2Fv0.3.1/sdk/resourcemanager/notificationhubs/armnotificationhubs/README.md) on how to add the SDK to your project and authenticate.

```go
package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubCheckNameAvailability.json
func ExampleClient_CheckNotificationHubAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnotificationhubs.NewClient("<subscription-id>", cred, nil)
	res, err := client.CheckNotificationHubAvailability(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		armnotificationhubs.CheckAvailabilityParameters{
			Name:     to.StringPtr("<name>"),
			Location: to.StringPtr("<location>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientCheckNotificationHubAvailabilityResult)
}
```
