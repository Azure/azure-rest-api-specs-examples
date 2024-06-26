package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteCircuitConnectionCreate.json
func ExampleExpressRouteCircuitConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteCircuitConnectionsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "ExpressRouteARMCircuitA", "AzurePrivatePeering", "circuitConnectionUSAUS", armnetwork.ExpressRouteCircuitConnection{
		Properties: &armnetwork.ExpressRouteCircuitConnectionPropertiesFormat{
			AddressPrefix:    to.Ptr("10.0.0.0/29"),
			AuthorizationKey: to.Ptr("946a1918-b7a2-4917-b43c-8c4cdaee006a"),
			ExpressRouteCircuitPeering: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid1/resourceGroups/dedharcktinit/providers/Microsoft.Network/expressRouteCircuits/dedharcktlocal/peerings/AzurePrivatePeering"),
			},
			IPv6CircuitConnectionConfig: &armnetwork.IPv6CircuitConnectionConfig{
				AddressPrefix: to.Ptr("aa:bb::/125"),
			},
			PeerExpressRouteCircuitPeering: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid2/resourceGroups/dedharcktpeer/providers/Microsoft.Network/expressRouteCircuits/dedharcktremote/peerings/AzurePrivatePeering"),
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
