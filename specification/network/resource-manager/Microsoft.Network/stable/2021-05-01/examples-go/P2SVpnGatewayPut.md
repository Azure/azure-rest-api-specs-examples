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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/P2SVpnGatewayPut.json
func ExampleP2SVPNGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewP2SVPNGatewaysClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gateway-name>",
		armnetwork.P2SVPNGateway{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Properties: &armnetwork.P2SVPNGatewayProperties{
				CustomDNSServers: []*string{
					to.StringPtr("1.1.1.1"),
					to.StringPtr("2.2.2.2")},
				IsRoutingPreferenceInternet: to.BoolPtr(false),
				P2SConnectionConfigurations: []*armnetwork.P2SConnectionConfiguration{
					{
						ID:   to.StringPtr("<id>"),
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.P2SConnectionConfigurationProperties{
							RoutingConfiguration: &armnetwork.RoutingConfiguration{
								AssociatedRouteTable: &armnetwork.SubResource{
									ID: to.StringPtr("<id>"),
								},
								PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
									IDs: []*armnetwork.SubResource{
										{
											ID: to.StringPtr("<id>"),
										},
										{
											ID: to.StringPtr("<id>"),
										},
										{
											ID: to.StringPtr("<id>"),
										}},
									Labels: []*string{
										to.StringPtr("label1"),
										to.StringPtr("label2")},
								},
								VnetRoutes: &armnetwork.VnetRoute{
									StaticRoutes: []*armnetwork.StaticRoute{},
								},
							},
							VPNClientAddressPool: &armnetwork.AddressSpace{
								AddressPrefixes: []*string{
									to.StringPtr("101.3.0.0/16")},
							},
						},
					}},
				VirtualHub: &armnetwork.SubResource{
					ID: to.StringPtr("<id>"),
				},
				VPNGatewayScaleUnit: to.Int32Ptr(1),
				VPNServerConfiguration: &armnetwork.SubResource{
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
	log.Printf("Response result: %#v\n", res.P2SVPNGatewaysClientCreateOrUpdateResult)
}
```
