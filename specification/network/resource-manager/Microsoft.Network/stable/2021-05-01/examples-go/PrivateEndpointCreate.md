Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointCreate.json
func ExamplePrivateEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewPrivateEndpointsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-endpoint-name>",
		armnetwork.PrivateEndpoint{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.PrivateEndpointProperties{
				CustomNetworkInterfaceName: to.Ptr("<custom-network-interface-name>"),
				IPConfigurations: []*armnetwork.PrivateEndpointIPConfiguration{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
							GroupID:          to.Ptr("<group-id>"),
							MemberName:       to.Ptr("<member-name>"),
							PrivateIPAddress: to.Ptr("<private-ipaddress>"),
						},
					}},
				PrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
					{
						Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
							GroupIDs: []*string{
								to.Ptr("groupIdFromResource")},
							PrivateLinkServiceID: to.Ptr("<private-link-service-id>"),
							RequestMessage:       to.Ptr("<request-message>"),
						},
					}},
				Subnet: &armnetwork.Subnet{
					ID: to.Ptr("<id>"),
				},
			},
		},
		&armnetwork.PrivateEndpointsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
