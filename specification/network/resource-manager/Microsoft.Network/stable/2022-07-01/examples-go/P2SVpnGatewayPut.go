package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/P2SVpnGatewayPut.json
func ExampleP2SVPNGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewP2SVPNGatewaysClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "p2sVpnGateway1", armnetwork.P2SVPNGateway{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.P2SVPNGatewayProperties{
			CustomDNSServers: []*string{
				to.Ptr("1.1.1.1"),
				to.Ptr("2.2.2.2")},
			IsRoutingPreferenceInternet: to.Ptr(false),
			P2SConnectionConfigurations: []*armnetwork.P2SConnectionConfiguration{
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1"),
					Name: to.Ptr("P2SConnectionConfig1"),
					Properties: &armnetwork.P2SConnectionConfigurationProperties{
						RoutingConfiguration: &armnetwork.RoutingConfiguration{
							AssociatedRouteTable: &armnetwork.SubResource{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
							},
							PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
								IDs: []*armnetwork.SubResource{
									{
										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
									},
									{
										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
									},
									{
										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"),
									}},
								Labels: []*string{
									to.Ptr("label1"),
									to.Ptr("label2")},
							},
							VnetRoutes: &armnetwork.VnetRoute{
								StaticRoutes: []*armnetwork.StaticRoute{},
							},
						},
						VPNClientAddressPool: &armnetwork.AddressSpace{
							AddressPrefixes: []*string{
								to.Ptr("101.3.0.0/16")},
						},
					},
				}},
			VirtualHub: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
			},
			VPNGatewayScaleUnit: to.Ptr[int32](1),
			VPNServerConfiguration: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
