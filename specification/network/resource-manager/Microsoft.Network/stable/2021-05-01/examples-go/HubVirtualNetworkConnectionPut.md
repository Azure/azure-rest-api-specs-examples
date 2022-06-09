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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/HubVirtualNetworkConnectionPut.json
func ExampleHubVirtualNetworkConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewHubVirtualNetworkConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-hub-name>",
		"<connection-name>",
		armnetwork.HubVirtualNetworkConnection{
			Properties: &armnetwork.HubVirtualNetworkConnectionProperties{
				EnableInternetSecurity: to.Ptr(false),
				RemoteVirtualNetwork: &armnetwork.SubResource{
					ID: to.Ptr("<id>"),
				},
				RoutingConfiguration: &armnetwork.RoutingConfiguration{
					AssociatedRouteTable: &armnetwork.SubResource{
						ID: to.Ptr("<id>"),
					},
					PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
						IDs: []*armnetwork.SubResource{
							{
								ID: to.Ptr("<id>"),
							}},
						Labels: []*string{
							to.Ptr("label1"),
							to.Ptr("label2")},
					},
					VnetRoutes: &armnetwork.VnetRoute{
						StaticRoutes: []*armnetwork.StaticRoute{
							{
								Name: to.Ptr("<name>"),
								AddressPrefixes: []*string{
									to.Ptr("10.1.0.0/16"),
									to.Ptr("10.2.0.0/16")},
								NextHopIPAddress: to.Ptr("<next-hop-ipaddress>"),
							},
							{
								Name: to.Ptr("<name>"),
								AddressPrefixes: []*string{
									to.Ptr("10.3.0.0/16"),
									to.Ptr("10.4.0.0/16")},
								NextHopIPAddress: to.Ptr("<next-hop-ipaddress>"),
							}},
					},
				},
			},
		},
		&armnetwork.HubVirtualNetworkConnectionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.
