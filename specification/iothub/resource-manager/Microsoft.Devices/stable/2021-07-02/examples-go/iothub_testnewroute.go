package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_testnewroute.json
func ExampleResourceClient_TestRoute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.TestRoute(ctx,
		"testHub",
		"myResourceGroup",
		armiothub.TestRouteInput{
			Message: &armiothub.RoutingMessage{
				AppProperties: map[string]*string{
					"key1": to.Ptr("value1"),
				},
				Body: to.Ptr("Body of message"),
				SystemProperties: map[string]*string{
					"key1": to.Ptr("value1"),
				},
			},
			Route: &armiothub.RouteProperties{
				Name: to.Ptr("Routeid"),
				EndpointNames: []*string{
					to.Ptr("id1")},
				IsEnabled: to.Ptr(true),
				Source:    to.Ptr(armiothub.RoutingSourceDeviceMessages),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
