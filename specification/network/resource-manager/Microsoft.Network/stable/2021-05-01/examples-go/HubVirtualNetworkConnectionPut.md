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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/HubVirtualNetworkConnectionPut.json
func ExampleHubVirtualNetworkConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewHubVirtualNetworkConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-hub-name>",
		"<connection-name>",
		armnetwork.HubVirtualNetworkConnection{
			Properties: &armnetwork.HubVirtualNetworkConnectionProperties{
				EnableInternetSecurity: to.BoolPtr(false),
				RemoteVirtualNetwork: &armnetwork.SubResource{
					ID: to.StringPtr("<id>"),
				},
				RoutingConfiguration: &armnetwork.RoutingConfiguration{
					AssociatedRouteTable: &armnetwork.SubResource{
						ID: to.StringPtr("<id>"),
					},
					PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
						IDs: []*armnetwork.SubResource{
							{
								ID: to.StringPtr("<id>"),
							}},
						Labels: []*string{
							to.StringPtr("label1"),
							to.StringPtr("label2")},
					},
					VnetRoutes: &armnetwork.VnetRoute{
						StaticRoutes: []*armnetwork.StaticRoute{
							{
								Name: to.StringPtr("<name>"),
								AddressPrefixes: []*string{
									to.StringPtr("10.1.0.0/16"),
									to.StringPtr("10.2.0.0/16")},
								NextHopIPAddress: to.StringPtr("<next-hop-ipaddress>"),
							},
							{
								Name: to.StringPtr("<name>"),
								AddressPrefixes: []*string{
									to.StringPtr("10.3.0.0/16"),
									to.StringPtr("10.4.0.0/16")},
								NextHopIPAddress: to.StringPtr("<next-hop-ipaddress>"),
							}},
					},
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
	log.Printf("Response result: %#v\n", res.HubVirtualNetworkConnectionsClientCreateOrUpdateResult)
}
```
