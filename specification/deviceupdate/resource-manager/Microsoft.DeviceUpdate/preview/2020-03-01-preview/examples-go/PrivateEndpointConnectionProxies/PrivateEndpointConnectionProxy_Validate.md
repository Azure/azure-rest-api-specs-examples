Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeviceupdate%2Farmdeviceupdate%2Fv0.4.0/sdk/resourcemanager/deviceupdate/armdeviceupdate/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2020-03-01-preview/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Validate.json
func ExamplePrivateEndpointConnectionProxiesClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdeviceupdate.NewPrivateEndpointConnectionProxiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Validate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<private-endpoint-connection-proxy-id>",
		armdeviceupdate.PrivateEndpointConnectionProxy{
			RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
				ID:                      to.Ptr("<id>"),
				ImmutableResourceID:     to.Ptr("<immutable-resource-id>"),
				ImmutableSubscriptionID: to.Ptr("<immutable-subscription-id>"),
				Location:                to.Ptr("<location>"),
				ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
					{
						Name: to.Ptr("<name>"),
						GroupIDs: []*string{
							to.Ptr("DeviceUpdate")},
						RequestMessage: to.Ptr("<request-message>"),
					}},
				PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
					{
						GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{},
						ID:                           to.Ptr("<id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
