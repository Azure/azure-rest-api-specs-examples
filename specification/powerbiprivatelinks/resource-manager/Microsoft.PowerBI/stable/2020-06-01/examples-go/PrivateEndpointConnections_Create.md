Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbiprivatelinks%2Farmpowerbiprivatelinks%2Fv1.0.0/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/README.md) on how to add the SDK to your project and authenticate.

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
	}
	ctx := context.Background()
	client, err := armpowerbiprivatelinks.NewPrivateEndpointConnectionsClient("a0020869-4d28-422a-89f4-c2413130d73c",
		"resourceGroup",
		"azureResourceName",
		"myPrivateEndpointName", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		armpowerbiprivatelinks.PrivateEndpointConnection{
			Properties: &armpowerbiprivatelinks.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armpowerbiprivatelinks.PrivateEndpoint{
					ID: to.Ptr("/subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpointName"),
				},
				PrivateLinkServiceConnectionState: &armpowerbiprivatelinks.ConnectionState{
					Description:     to.Ptr(""),
					ActionsRequired: to.Ptr("None"),
					Status:          to.Ptr(armpowerbiprivatelinks.PersistedConnectionStatus("Approved ")),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
