Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdigitaltwins%2Farmdigitaltwins%2Fv0.1.0/sdk/resourcemanager/digitaltwins/armdigitaltwins/README.md) on how to add the SDK to your project and authenticate.

```go
package armdigitaltwins_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2020-12-01/examples/DigitalTwinsEndpointPut_example.json
func ExampleDigitalTwinsEndpointClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdigitaltwins.NewDigitalTwinsEndpointClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<endpoint-name>",
		armdigitaltwins.DigitalTwinsEndpointResource{
			Properties: &armdigitaltwins.ServiceBus{
				DigitalTwinsEndpointResourceProperties: armdigitaltwins.DigitalTwinsEndpointResourceProperties{
					AuthenticationType: armdigitaltwins.AuthenticationTypeKeyBased.ToPtr(),
					EndpointType:       armdigitaltwins.EndpointTypeServiceBus.ToPtr(),
				},
				PrimaryConnectionString:   to.StringPtr("<primary-connection-string>"),
				SecondaryConnectionString: to.StringPtr("<secondary-connection-string>"),
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
	log.Printf("DigitalTwinsEndpointResource.ID: %s\n", *res.ID)
}
```
