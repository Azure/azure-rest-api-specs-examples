Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicebus%2Farmservicebus%2Fv0.3.1/sdk/resourcemanager/servicebus/armservicebus/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Subscriptions/SBSubscriptionDelete.json
func ExampleSubscriptionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicebus.NewSubscriptionsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		"<topic-name>",
		"<subscription-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
