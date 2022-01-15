Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicebus%2Farmservicebus%2Fv0.3.0/sdk/resourcemanager/servicebus/armservicebus/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/PrivateEndPointConnectionCreate.json
func ExamplePrivateEndpointConnectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicebus.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		"<private-endpoint-connection-name>",
		armservicebus.PrivateEndpointConnection{
			Properties: &armservicebus.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armservicebus.PrivateEndpoint{
					ID: to.StringPtr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armservicebus.ConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armservicebus.PrivateLinkConnectionStatus("Rejected").ToPtr(),
				},
				ProvisioningState: armservicebus.EndPointProvisioningState("Succeeded").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientCreateOrUpdateResult)
}
```
