Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsignalr%2Farmsignalr%2Fv0.5.0/sdk/resourcemanager/signalr/armsignalr/README.md) on how to add the SDK to your project and authenticate.

```go
package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRPrivateEndpointConnections_Update.json
func ExamplePrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsignalr.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<private-endpoint-connection-name>",
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.PrivateEndpointConnection{
			Properties: &armsignalr.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armsignalr.PrivateEndpoint{
					ID: to.Ptr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armsignalr.PrivateLinkServiceConnectionState{
					ActionsRequired: to.Ptr("<actions-required>"),
					Status:          to.Ptr(armsignalr.PrivateLinkServiceConnectionStatusApproved),
				},
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
