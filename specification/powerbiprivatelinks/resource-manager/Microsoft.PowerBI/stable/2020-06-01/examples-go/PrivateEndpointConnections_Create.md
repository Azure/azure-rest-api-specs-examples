Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbiprivatelinks%2Farmpowerbiprivatelinks%2Fv0.2.0/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/README.md) on how to add the SDK to your project and authenticate.

```go
package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks"
)

// x-ms-original-file: specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateEndpointConnections_Create.json
func ExamplePrivateEndpointConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpowerbiprivatelinks.NewPrivateEndpointConnectionsClient("<subscription-id>",
		"<resource-group-name>",
		"<azure-resource-name>",
		"<private-endpoint-name>", cred, nil)
	res, err := client.Create(ctx,
		armpowerbiprivatelinks.PrivateEndpointConnection{
			Properties: &armpowerbiprivatelinks.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armpowerbiprivatelinks.PrivateEndpoint{
					ID: to.StringPtr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armpowerbiprivatelinks.ConnectionState{
					Description:     to.StringPtr("<description>"),
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          armpowerbiprivatelinks.PersistedConnectionStatus("Approved ").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientCreateResult)
}
```
