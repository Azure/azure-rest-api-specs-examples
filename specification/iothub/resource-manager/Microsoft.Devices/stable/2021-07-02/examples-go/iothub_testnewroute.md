Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiothub%2Farmiothub%2Fv0.5.0/sdk/resourcemanager/iothub/armiothub/README.md) on how to add the SDK to your project and authenticate.

```go
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
		return
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.TestRoute(ctx,
		"<iot-hub-name>",
		"<resource-group-name>",
		armiothub.TestRouteInput{
			Message: &armiothub.RoutingMessage{
				AppProperties: map[string]*string{
					"key1": to.Ptr("value1"),
				},
				Body: to.Ptr("<body>"),
				SystemProperties: map[string]*string{
					"key1": to.Ptr("value1"),
				},
			},
			Route: &armiothub.RouteProperties{
				Name: to.Ptr("<name>"),
				EndpointNames: []*string{
					to.Ptr("id1")},
				IsEnabled: to.Ptr(true),
				Source:    to.Ptr(armiothub.RoutingSourceDeviceMessages),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
