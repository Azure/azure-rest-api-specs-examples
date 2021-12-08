Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeviceupdate%2Farmdeviceupdate%2Fv0.1.0/sdk/resourcemanager/deviceupdate/armdeviceupdate/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeviceupdate_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2020-03-01-preview/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_CreateOrUpdate.json
func ExamplePrivateEndpointConnectionProxiesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeviceupdate.NewPrivateEndpointConnectionProxiesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<private-endpoint-connection-proxy-id>",
		armdeviceupdate.PrivateEndpointConnectionProxy{
			PrivateEndpointConnectionProxyProperties: armdeviceupdate.PrivateEndpointConnectionProxyProperties{
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
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateEndpointConnectionProxy.ID: %s\n", *res.ID)
}
```
