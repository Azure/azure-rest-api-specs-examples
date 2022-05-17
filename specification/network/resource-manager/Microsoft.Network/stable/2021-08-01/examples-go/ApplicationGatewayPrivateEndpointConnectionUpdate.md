Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv1.0.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ApplicationGatewayPrivateEndpointConnectionUpdate.json
func ExampleApplicationGatewayPrivateEndpointConnectionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewApplicationGatewayPrivateEndpointConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"rg1",
		"appgw",
		"connection1",
		armnetwork.ApplicationGatewayPrivateEndpointConnection{
			Name: to.Ptr("connection1"),
			Properties: &armnetwork.ApplicationGatewayPrivateEndpointConnectionProperties{
				PrivateEndpoint: &armnetwork.PrivateEndpoint{
					ID: to.Ptr("/subscriptions/subId2/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
				},
				PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
					Description: to.Ptr("approved it for some reason."),
					Status:      to.Ptr("Approved"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
