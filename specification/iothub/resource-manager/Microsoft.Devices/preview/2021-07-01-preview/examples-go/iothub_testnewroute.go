package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2021-07-01-preview/examples/iothub_testnewroute.json
func ExampleResourceClient_TestRoute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiothub.NewResourceClient("<subscription-id>", cred, nil)
	res, err := client.TestRoute(ctx,
		"<iot-hub-name>",
		"<resource-group-name>",
		armiothub.TestRouteInput{
			Message: &armiothub.RoutingMessage{
				AppProperties: map[string]*string{
					"key1": to.StringPtr("value1"),
				},
				Body: to.StringPtr("<body>"),
				SystemProperties: map[string]*string{
					"key1": to.StringPtr("value1"),
				},
			},
			Route: &armiothub.RouteProperties{
				Name: to.StringPtr("<name>"),
				EndpointNames: []*string{
					to.StringPtr("id1")},
				IsEnabled: to.BoolPtr(true),
				Source:    armiothub.RoutingSource("DeviceMessages").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ResourceClientTestRouteResult)
}
