Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbiprivatelinks%2Farmpowerbiprivatelinks%2Fv0.4.0/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/README.md) on how to add the SDK to your project and authenticate.

```go
package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateEndpointConnections_Create.json
func ExamplePrivateEndpointConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpowerbiprivatelinks.NewPrivateEndpointConnectionsClient("<subscription-id>",
		"<resource-group-name>",
		"<azure-resource-name>",
		"<private-endpoint-name>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		armpowerbiprivatelinks.PrivateEndpointConnection{
			Properties: &armpowerbiprivatelinks.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armpowerbiprivatelinks.PrivateEndpoint{
					ID: to.Ptr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armpowerbiprivatelinks.ConnectionState{
					Description:     to.Ptr("<description>"),
					ActionsRequired: to.Ptr("<actions-required>"),
					Status:          to.Ptr(armpowerbiprivatelinks.PersistedConnectionStatus("Approved ")),
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
