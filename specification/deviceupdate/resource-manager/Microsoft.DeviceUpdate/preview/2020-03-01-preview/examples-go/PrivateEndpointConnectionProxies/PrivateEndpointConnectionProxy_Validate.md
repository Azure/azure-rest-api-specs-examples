Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeviceupdate%2Farmdeviceupdate%2Fv0.2.1/sdk/resourcemanager/deviceupdate/armdeviceupdate/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2020-03-01-preview/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Validate.json
func ExamplePrivateEndpointConnectionProxiesClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeviceupdate.NewPrivateEndpointConnectionProxiesClient("<subscription-id>", cred, nil)
	_, err = client.Validate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<private-endpoint-connection-proxy-id>",
		armdeviceupdate.PrivateEndpointConnectionProxy{
			RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
				ID: to.StringPtr("<id>"),
				ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
					{
						Name: to.StringPtr("<name>"),
						GroupIDs: []*string{
							to.StringPtr("DeviceUpdate")},
						RequestMessage: to.StringPtr("<request-message>"),
					}},
				PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
					{
						GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{},
						ID:                           to.StringPtr("<id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
