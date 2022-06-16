package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/streamingendpoint-update.json
func ExampleStreamingEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewStreamingEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<streaming-endpoint-name>",
		armmediaservices.StreamingEndpoint{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag3": to.StringPtr("value3"),
				"tag5": to.StringPtr("value5"),
			},
			Properties: &armmediaservices.StreamingEndpointProperties{
				Description:         to.StringPtr("<description>"),
				AvailabilitySetName: to.StringPtr("<availability-set-name>"),
				ScaleUnits:          to.Int32Ptr(5),
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
	log.Printf("Response result: %#v\n", res.StreamingEndpointsClientUpdateResult)
}
