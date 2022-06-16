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
func ExampleEndpointClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdigitaltwins.NewEndpointClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<endpoint-name>",
		armdigitaltwins.EndpointResource{
			Properties: &armdigitaltwins.ServiceBus{
				AuthenticationType:        armdigitaltwins.AuthenticationType("KeyBased").ToPtr(),
				EndpointType:              armdigitaltwins.EndpointType("ServiceBus").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.EndpointClientCreateOrUpdateResult)
}
