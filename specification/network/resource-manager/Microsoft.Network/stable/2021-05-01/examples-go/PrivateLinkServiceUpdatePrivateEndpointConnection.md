Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateLinkServiceUpdatePrivateEndpointConnection.json
func ExamplePrivateLinkServicesClient_UpdatePrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewPrivateLinkServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.UpdatePrivateEndpointConnection(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<pe-connection-name>",
		armnetwork.PrivateEndpointConnection{
			Name: to.Ptr("<name>"),
			Properties: &armnetwork.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armnetwork.PrivateEndpoint{
					ID: to.Ptr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
					Description: to.Ptr("<description>"),
					Status:      to.Ptr("<status>"),
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
