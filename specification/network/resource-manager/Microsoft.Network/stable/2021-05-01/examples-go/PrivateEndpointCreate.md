Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointCreate.json
func ExamplePrivateEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPrivateEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-endpoint-name>",
		armnetwork.PrivateEndpoint{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.PrivateEndpointProperties{
				CustomNetworkInterfaceName: to.StringPtr("<custom-network-interface-name>"),
				IPConfigurations: []*armnetwork.PrivateEndpointIPConfiguration{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
							GroupID:          to.StringPtr("<group-id>"),
							MemberName:       to.StringPtr("<member-name>"),
							PrivateIPAddress: to.StringPtr("<private-ipaddress>"),
						},
					}},
				PrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
					{
						Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
							GroupIDs: []*string{
								to.StringPtr("groupIdFromResource")},
							PrivateLinkServiceID: to.StringPtr("<private-link-service-id>"),
							RequestMessage:       to.StringPtr("<request-message>"),
						},
					}},
				Subnet: &armnetwork.Subnet{
					ID: to.StringPtr("<id>"),
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
	log.Printf("Response result: %#v\n", res.PrivateEndpointsClientCreateOrUpdateResult)
}
```
