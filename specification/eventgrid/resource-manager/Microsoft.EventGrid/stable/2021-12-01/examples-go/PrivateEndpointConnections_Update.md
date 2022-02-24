Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventgrid%2Farmeventgrid%2Fv0.3.1/sdk/resourcemanager/eventgrid/armeventgrid/README.md) on how to add the SDK to your project and authenticate.

```go
package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/PrivateEndpointConnections_Update.json
func ExamplePrivateEndpointConnectionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armeventgrid.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		armeventgrid.Enum19("topics"),
		"<parent-name>",
		"<private-endpoint-connection-name>",
		armeventgrid.PrivateEndpointConnection{
			Properties: &armeventgrid.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armeventgrid.ConnectionState{
					Description:     to.StringPtr("<description>"),
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          armeventgrid.PersistedConnectionStatus("Approved").ToPtr(),
				},
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
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientUpdateResult)
}
```
